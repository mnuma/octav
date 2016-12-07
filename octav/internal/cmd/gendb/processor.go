package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lestrrat/go-pdebug"
)

var ErrAnnotatedStructNotFound = errors.New("annotated struct was not found")

func snakeCase(s string) string {
	ret := []rune{}
	wasLower := false
	upCount := 0
	for len(s) > 0 {
		r, n := utf8.DecodeRuneInString(s)
		if r == utf8.RuneError {
			panic("yikes")
		}

		s = s[n:]
		if unicode.IsUpper(r) {
			upCount++
			if wasLower {
				ret = append(ret, '_')
			}
			wasLower = false
		} else {
			// consecutive upper cases. naively assume that we only have
			// A-Za-z0-9 in our column names
			if ((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) && upCount > 1 {
				ret = append(ret, ret[len(ret)-1])
				ret[len(ret)-2] = '_'
			}
			upCount = 0
			wasLower = true
		}

		ret = append(ret, unicode.ToLower(r))
	}
	return string(ret)
}


type Processor struct {
	Types []string
	Dir   string
}

func skipGenerated(fi os.FileInfo) bool {
	switch {
	case strings.HasSuffix(fi.Name(), "gen.go"):
		return false
	case strings.HasSuffix(fi.Name(), "_gen.go"):
		return false
	}
	return true
}

func (p *Processor) Do() error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, p.Dir, skipGenerated, parser.ParseComments)
	if err != nil {
		return err
	}

	if len(pkgs) == 0 {
		return errors.New("no packages to process...")
	}

	for _, pkg := range pkgs {
		if err := p.ProcessPkg(pkg); err != nil {
			return err
		}
	}

	return nil
}

func (p *Processor) ShouldProceed(s Struct) bool {
	if len(p.Types) == 0 {
		return true
	}

	for _, t := range p.Types {
		if t == s.Name {
			return true
		}
	}
	return false
}

func (p *Processor) ProcessPkg(pkg *ast.Package) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ProcessPkg %s", pkg.Name)
		defer g.End()
	}
	for fn, f := range pkg.Files {
		pdebug.Printf("Checking file %s", fn)
		for _, s := range p.ExtractStructs(pkg, f) {
			if pdebug.Enabled {
				pdebug.Printf("Checking struct %s", s.Name)
			}
			if !p.ShouldProceed(s) {
				if pdebug.Enabled {
					pdebug.Printf("Skipping struct %s", s.Name)
				}
				continue
			}

			if err := p.ProcessStruct(s); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s Struct) sqlKeyName(txt string) string {
	return `sql` + s.Name + txt + `Key`
}

func (p *Processor) ProcessStruct(s Struct) error {
	if pdebug.Enabled {
		g := pdebug.Marker("ProcessStruct %s", s.Name)
		defer g.End()
	}

	buf := bytes.Buffer{}

	varname := unicode.ToLower(rune(s.Name[0]))

	scols := bytes.Buffer{}
	icols := bytes.Buffer{}
	sfields := bytes.Buffer{}   // Scan fields
	ifields := bytes.Buffer{}   // Scan fields (everything except for OID)
	ipholders := bytes.Buffer{} // Insert place holders (everything except for OID)
	setCols := bytes.Buffer{}
	setColsSansEID := bytes.Buffer{}
	setParams := bytes.Buffer{}
	setParamsSansEID := bytes.Buffer{}

	hasEID := false
	hasOID := true
	hasCreatedOn := false
	sansOidFields := make([]StructField, 0, len(s.Fields)-1)
	updateFields := []StructField{}
	for _, f := range s.Fields {
		switch f.Name {
		case "OID":
			hasOID = true
			continue
		case "EID":
			hasEID = true
		case "CreatedOn":
			hasCreatedOn = true
		}
		sansOidFields = append(sansOidFields, f)

		switch f.Name {
		case "OID", "CreatedOn", "ModifiedOn":
		default:
			updateFields = append(updateFields, f)
		}
	}

	buf.WriteString("package " + s.PackageName)
	buf.WriteString("\n\n// Automatically generated by gendb utility. DO NOT EDIT!")
	buf.WriteString("\n\nimport (")
	buf.WriteString("\n" + `"bytes"`)
	buf.WriteString("\n" + `"database/sql"`)
	if hasEID && hasOID {
		buf.WriteString("\n" + `"strconv"`)
	}
	if hasCreatedOn {
		buf.WriteString("\n" + `"time"`)
	}
	buf.WriteString("\n\n" + `"github.com/lestrrat/go-pdebug"`)
	buf.WriteString("\n" + `"github.com/pkg/errors"`)
	if hasEID || hasOID {
		buf.WriteString("\n" + `"github.com/builderscon/octav/octav/tools"`)
	}
	buf.WriteString("\n)")

	for i, f := range updateFields {
		setCols.WriteString(f.ColumnName)
		setCols.WriteString(" = ?")
		fmt.Fprintf(&setParams, "%c.%s", varname, f.Name)
		if f.Name != "EID" {
			setColsSansEID.WriteString(f.ColumnName)
			setColsSansEID.WriteString(" = ?")
			fmt.Fprintf(&setParamsSansEID, "%c.%s", varname, f.Name)
		}

		if i < len(updateFields)-1 {
			setCols.WriteString(", ")
			setParams.WriteString(", ")
			if f.Name != "EID" {
				setColsSansEID.WriteString(", ")
				setParamsSansEID.WriteString(", ")
			}
		}
	}

	for i, f := range sansOidFields {
		icols.WriteString(f.ColumnName)
		ipholders.WriteRune('?')
		ifields.WriteRune(varname)
		ifields.WriteRune('.')
		ifields.WriteString(f.Name)

		if i < len(sansOidFields)-1 {
			icols.WriteString(", ")
			ipholders.WriteString(", ")
			ifields.WriteString(", ")
		}
	}

	for i, f := range s.Fields {
		scols.WriteString(s.Tablename)
		scols.WriteByte('.')
		scols.WriteString(f.ColumnName)
		sfields.WriteRune('&')
		sfields.WriteRune(varname)
		sfields.WriteRune('.')
		sfields.WriteString(f.Name)
		if i < len(s.Fields)-1 {
			scols.WriteString(", ")
			sfields.WriteString(", ")
		}
	}

	fmt.Fprintf(&buf, "\nconst %sStdSelectColumns = %s", s.Name, strconv.Quote(scols.String()))

	fmt.Fprintf(&buf, "\nconst %sTable = %s", s.Name, strconv.Quote(s.Tablename))
	fmt.Fprintf(&buf, "\ntype %sList []%s", s.Name, s.Name)
	fmt.Fprintf(&buf, "\n\nfunc (%c *%s) Scan(scanner interface { Scan(...interface{}) error }) error {", varname, s.Name)
	fmt.Fprintf(&buf, "\nreturn scanner.Scan(%s)", sfields.String())
	buf.WriteString("\n}\n")

	buf.WriteString("\nfunc init() {")
	buf.WriteString("\nhooks = append(hooks, func() {")
	buf.WriteString("\nstmt := tools.GetBuffer()")
	buf.WriteString("\ndefer tools.ReleaseBuffer(stmt)")

	if hasOID {
		buf.WriteString("\n\nstmt.Reset()")
		buf.WriteString("\nstmt.WriteString(`DELETE FROM `)")
		fmt.Fprintf(&buf, "\nstmt.WriteString(%sTable)", s.Name)
		buf.WriteString("\nstmt.WriteString(` WHERE oid = ?`)")
		fmt.Fprintf(&buf, "\nlibrary.Register(%s, stmt.String())", strconv.Quote(s.sqlKeyName("DeleteByOID")))

		buf.WriteString("\n\nstmt.Reset()")
		buf.WriteString("\nstmt.WriteString(`UPDATE `)")
		fmt.Fprintf(&buf, "\nstmt.WriteString(%sTable)", s.Name)
		fmt.Fprintf(&buf, "\nstmt.WriteString(` SET %s WHERE oid = ?`)", setCols.String())
		fmt.Fprintf(&buf, "\nlibrary.Register(%s, stmt.String())", strconv.Quote(s.sqlKeyName("UpdateByOID")))
	}

	if hasEID {
		buf.WriteString("\n\nstmt.Reset()")
		buf.WriteString("\nstmt.WriteString(`SELECT `)")
		fmt.Fprintf(&buf, "\nstmt.WriteString(%sStdSelectColumns)", s.Name)
		buf.WriteString("\nstmt.WriteString(` FROM `)")
		fmt.Fprintf(&buf, "\nstmt.WriteString(%sTable)", s.Name)
		buf.WriteString("\nstmt.WriteString(` WHERE `)")
		fmt.Fprintf(&buf, "\nstmt.WriteString(%sTable)", s.Name)
		buf.WriteString("\nstmt.WriteString(`.eid = ?`)")
		fmt.Fprintf(&buf, "\nlibrary.Register(%s, stmt.String())", strconv.Quote(s.sqlKeyName("LoadByEID")))

		buf.WriteString("\n\nstmt.Reset()")
		buf.WriteString("\nstmt.WriteString(`DELETE FROM `)")
		fmt.Fprintf(&buf, "\nstmt.WriteString(%sTable)", s.Name)
		buf.WriteString("\nstmt.WriteString(` WHERE eid = ?`)")
		fmt.Fprintf(&buf, "\nlibrary.Register(%s, stmt.String())", strconv.Quote(s.sqlKeyName("DeleteByEID")))

		buf.WriteString("\n\nstmt.Reset()")
		buf.WriteString("\nstmt.WriteString(`UPDATE `)")
		fmt.Fprintf(&buf, "\nstmt.WriteString(%sTable)", s.Name)
		fmt.Fprintf(&buf, "\nstmt.WriteString(` SET %s WHERE eid = ?`)", setCols.String())
		fmt.Fprintf(&buf, "\nlibrary.Register(%s, stmt.String())", strconv.Quote(s.sqlKeyName("UpdateByEID")))
	}

	buf.WriteString("\n})")
	buf.WriteString("\n}")

	if hasEID {
		fmt.Fprintf(&buf, "\n\nfunc (%c *%s) LoadByEID(tx *Tx, eid string) (err error) {", varname, s.Name)
		buf.WriteString("\nif pdebug.Enabled {")
		fmt.Fprintf(&buf, "\ng := pdebug.Marker(`%s.LoadByEID %%s`, eid).BindError(&err)", s.Name)
		buf.WriteString("\ndefer g.End()")
		buf.WriteString("\n}")
		fmt.Fprintf(&buf, "\nstmt, err := library.GetStmt(%s)", strconv.Quote(s.sqlKeyName("LoadByEID")))
		buf.WriteString("\nif err != nil {\nreturn errors.Wrap(err, `failed to get statement`)\n}")
		buf.WriteString("\nrow := tx.Stmt(stmt).QueryRow(eid)")
		fmt.Fprintf(&buf, "\nif err := %c.Scan(row); err != nil {", varname)
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\nreturn nil")
		buf.WriteString("\n}")
	}

	fmt.Fprintf(&buf, "\n\nfunc (%c *%s) Create(tx *Tx, opts ...InsertOption) (err error) {", varname, s.Name)
	buf.WriteString("\nif pdebug.Enabled {")
	fmt.Fprintf(&buf, "\n"+`g := pdebug.Marker("db.%s.Create").BindError(&err)`, s.Name)
	buf.WriteString("\ndefer g.End()")
	fmt.Fprintf(&buf, "\n"+`pdebug.Printf("%%#v", %c)`, varname)
	buf.WriteString("\n}")
	if hasEID {
		fmt.Fprintf(&buf, "\nif %c.EID == "+`""`+" {", varname)
		buf.WriteString("\n" + `return errors.New("create: non-empty EID required")`)
		buf.WriteString("\n}\n\n")
	}
	if hasCreatedOn {
		fmt.Fprintf(&buf, "\n%c.CreatedOn = time.Now()", varname)
	}

	buf.WriteString("\ndoIgnore := false")
	buf.WriteString("\nfor _, opt := range opts {")
	buf.WriteString("\nswitch opt.(type) {")
	buf.WriteString("\ncase insertIgnoreOption:")
	buf.WriteString("\ndoIgnore = true")
	buf.WriteString("\n}\n}")

	buf.WriteString("\n\nstmt := bytes.Buffer{}")
	buf.WriteString("\n" + `stmt.WriteString("INSERT ")`)
	buf.WriteString("\nif doIgnore {")
	buf.WriteString("\n" + `stmt.WriteString("IGNORE ")`)
	buf.WriteString("\n}")
	buf.WriteString("\n" + `stmt.WriteString("INTO ")`)
	fmt.Fprintf(&buf, "\nstmt.WriteString(%sTable)", s.Name)
	fmt.Fprintf(&buf, "\nstmt.WriteString(` (%s) VALUES (%s)`)", icols.String(), ipholders.String())

	fmt.Fprintf(&buf, "\nresult, err := tx.Exec(stmt.String(), %s)", ifields.String())
	buf.WriteString("\nif err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	buf.WriteString("\nlii, err := result.LastInsertId()")
	buf.WriteString("\nif err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	fmt.Fprintf(&buf, "\n%c.OID = lii", varname)
	buf.WriteString("\nreturn nil")
	buf.WriteString("\n}")

	// This is very inefficient, but it's the best we can do for now
	fmt.Fprintf(&buf, "\n\nfunc (%c %s) Update(tx *Tx) (err error) {", varname, s.Name)
	buf.WriteString("\nif pdebug.Enabled {")
	fmt.Fprintf(&buf, "\ng := pdebug.Marker(`%s.Update`).BindError(&err)", s.Name)
	buf.WriteString("\ndefer g.End()")
	buf.WriteString("\n}")
	fmt.Fprintf(&buf, "\nif %c.OID != 0 {", varname)
	buf.WriteString("\nif pdebug.Enabled {")
	fmt.Fprintf(&buf, "\npdebug.Printf(`Using OID (%%d) as key`, %c.OID)", varname)
	buf.WriteString("\n}")
	fmt.Fprintf(&buf, "\nstmt, err := library.GetStmt(%s)", strconv.Quote(s.sqlKeyName("UpdateByOID")))
	buf.WriteString("\nif err != nil {\nreturn errors.Wrap(err, `failed to get statement`)\n}")
	fmt.Fprintf(&buf, "\n_, err = tx.Stmt(stmt).Exec(%s, %c.OID)", setParams.String(), varname)
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}")
	if hasEID {
		fmt.Fprintf(&buf, "\n"+`if %c.EID != "" {`, varname)
	buf.WriteString("\nif pdebug.Enabled {")
	fmt.Fprintf(&buf, "\npdebug.Printf(`Using EID (%%s) as key`, %c.EID)", varname)
	buf.WriteString("\n}")
		fmt.Fprintf(&buf, "\nstmt, err := library.GetStmt(%s)", strconv.Quote(s.sqlKeyName("UpdateByEID")))
		buf.WriteString("\nif err != nil {\nreturn errors.Wrap(err, `failed to get statement`)\n}")
		fmt.Fprintf(&buf, "\n_, err = tx.Stmt(stmt).Exec(%s, %c.EID)", setParams.String(), varname)
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
	}

	if hasEID {
		fmt.Fprintf(&buf, "\nreturn errors.New(%s)", strconv.Quote("either OID/EID must be filled"))
	} else {
		fmt.Fprintf(&buf, "\nreturn errors.New(%s)", strconv.Quote("OID must be filled"))
	}
	buf.WriteString("\n}")

	fmt.Fprintf(&buf, "\n\nfunc (%c %s) Delete(tx *Tx) error {", varname, s.Name)
	fmt.Fprintf(&buf, "\nif %c.OID != 0 {", varname)
	fmt.Fprintf(&buf, "\nstmt, err := library.GetStmt(%s)", strconv.Quote(s.sqlKeyName("DeleteByOID")))
	buf.WriteString("\nif err != nil {\nreturn errors.Wrap(err, `failed to get statement`)\n}")
	fmt.Fprintf(&buf, "\n_, err = tx.Stmt(stmt).Exec(%c.OID)", varname)
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}\n")
	if hasEID {
		fmt.Fprintf(&buf, "\nif %c.EID != %s {", varname, `""`)
		fmt.Fprintf(&buf, "\nstmt, err := library.GetStmt(%s)", strconv.Quote(s.sqlKeyName("DeleteByEID")))
		buf.WriteString("\nif err != nil {\nreturn errors.Wrap(err, `failed to get statement`)\n}")
		fmt.Fprintf(&buf, "\n_, err = tx.Stmt(stmt).Exec(%c.EID)", varname)
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}\n")
	}

	if hasEID {
		fmt.Fprintf(&buf, "\nreturn errors.New(%s)", strconv.Quote("either OID/EID must be filled"))
		buf.WriteString("\n}\n")
	} else {
		fmt.Fprintf(&buf, "\nreturn errors.New(%s)", strconv.Quote("column OID must be filled"))
		buf.WriteString("\n}")
	}

	fmt.Fprintf(&buf, "\n\nfunc (v *%sList) FromRows(rows *sql.Rows, capacity int) error {", s.Name)
	fmt.Fprintf(&buf, "\nvar res []%s", s.Name)
	buf.WriteString("\nif capacity > 0 {")
	fmt.Fprintf(&buf, "\nres = make([]%s, 0, capacity)", s.Name)
	buf.WriteString("\n} else {")
	fmt.Fprintf(&buf, "\nres = []%s{}", s.Name)
	buf.WriteString("\n}")
	buf.WriteString("\n\nfor rows.Next() {")
	fmt.Fprintf(&buf, "\nvdb := %s{}", s.Name)
	buf.WriteString("\nif err := vdb.Scan(rows); err != nil {")
	buf.WriteString("\nreturn err")
	buf.WriteString("\n}")
	buf.WriteString("\nres = append(res, vdb)")
	buf.WriteString("\n}")
	buf.WriteString("\n*v = res")
	buf.WriteString("\nreturn nil")
	buf.WriteString("\n}")

	if hasOID && hasEID {
		fmt.Fprintf(&buf, "\n\nfunc (v *%sList) LoadSinceEID(tx *Tx, since string, limit int) error {", s.Name)
		buf.WriteString("\nvar s int64")
		buf.WriteString("\n" + `if id := since; id != "" {`)
		fmt.Fprintf(&buf, "\nvdb := %s{}", s.Name)
		buf.WriteString("\nif err := vdb.LoadByEID(tx, id); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}\n")
		buf.WriteString("\ns = vdb.OID")
		buf.WriteString("\n}")
		buf.WriteString("\nreturn v.LoadSince(tx, s, limit)")
		buf.WriteString("\n}\n")

		fmt.Fprintf(&buf, "\n\nfunc (v *%sList) LoadSince(tx *Tx, since int64, limit int) error {", s.Name)
		fmt.Fprintf(&buf, "\nrows, err := tx.Query(`SELECT ` + %sStdSelectColumns + ` FROM ` + %sTable + ` WHERE %s.oid > ? ORDER BY oid ASC LIMIT ` + strconv.Itoa(limit), since)", s.Name, s.Name, s.Tablename)
		buf.WriteString("\nif err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\n\nif err := v.FromRows(rows, limit); err != nil {")
		buf.WriteString("\nreturn err")
		buf.WriteString("\n}")
		buf.WriteString("\nreturn nil")
		buf.WriteString("\n}")
	}

	fsrc, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return err
	}

	fn := filepath.Join(p.Dir, snakeCase(s.Name)+"_gen.go")
	if pdebug.Enabled {
		pdebug.Printf("Generating file %s", fn)
	}
	fi, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer fi.Close()

	if _, err := fi.Write(fsrc); err != nil {
		return err
	}

	return nil
}

func (p *Processor) ExtractStructs(pkg *ast.Package, f *ast.File) []Struct {
	ctx := &InspectionCtx{
		Marker:  "+DB",
		Package: pkg.Name,
	}

	ast.Inspect(f, ctx.ExtractStructs)
	return ctx.Structs
}

func (c *InspectionCtx) ExtractStructs(n ast.Node) bool {
	var decl *ast.GenDecl
	var ok bool
	var err error

	if decl, ok = n.(*ast.GenDecl); !ok {
		return true
	}

	if decl.Tok != token.TYPE {
		return true
	}

	if err = c.ExtractStructsFromDecl(decl); err != nil {
		return true
	}

	return false
}

func (ctx *InspectionCtx) ExtractStructsFromDecl(decl *ast.GenDecl) error {
	marked := false
	cacheEnabled := true
	noScanner := false
	tablename := ""
	preCreate := ""
	postCreate := ""
	cacheExpires := "1800"
	if cg := decl.Doc; cg != nil {
		for _, c := range cg.List {
			txt := strings.TrimSpace(strings.TrimLeft(c.Text, "//"))
			if strings.HasPrefix(txt, ctx.Marker) {
				marked = true
				tag := reflect.StructTag(strings.TrimPrefix(txt, ctx.Marker))
				tablename = tag.Get("tablename")
				preCreate = tag.Get("pre_create")
				postCreate = tag.Get("post_create")
				if tag.Get("cache") == "false" {
					cacheEnabled = false
				}
				if ce := tag.Get("cache_expires"); ce != "" {
					cacheExpires = ce
				}
				if tag.Get("scanner") == "false" {
					noScanner = true
				}
				break
			}
		}
	}

	if !marked {
		return ErrAnnotatedStructNotFound
	}

	for _, spec := range decl.Specs {
		var t *ast.TypeSpec
		var s *ast.StructType
		var ok bool
		var ident *ast.Ident

		if t, ok = spec.(*ast.TypeSpec); !ok {
			return ErrAnnotatedStructNotFound
		}

		if s, ok = t.Type.(*ast.StructType); !ok {
			return ErrAnnotatedStructNotFound
		}

		if tablename == "" {
			tablename = fmt.Sprintf("%s_%s",
				ctx.Package,
				snakeCase(t.Name.Name),
			)
		}

		st := Struct{
			PackageName:  ctx.Package,
			CacheEnabled: cacheEnabled,
			CacheExpires: cacheExpires,
			Fields:       make([]StructField, 0, len(s.Fields.List)),
			Name:         t.Name.Name,
			NoScanner:    noScanner,
			PreCreate:    preCreate,
			PostCreate:   postCreate,
			Tablename:    tablename,
		}

	LoopFields:
		for _, f := range s.Fields.List {
			autoIncrement := false
			primaryKey := false
			unique := false
			converter := ""
			columnName := ""

			if f.Tag != nil {
				v := f.Tag.Value
				if len(v) >= 2 {
					if v[0] == '`' {
						v = v[1:]
					}
					if v[len(v)-1] == '`' {
						v = v[:len(v)-1]
					}
				}

				tags := reflect.StructTag(v).Get("db")
				for _, tag := range strings.Split(tags, ",") {
					switch tag {
					case "ignore":
						continue LoopFields
					case "auto_increment":
						autoIncrement = true
					case "pk":
						primaryKey = true
					case "unique":
						unique = true
					default:
						switch {
						case strings.HasPrefix(tag, "converter="):
							converter = tag[10:]
						case strings.HasPrefix(tag, "column="):
							columnName = tag[7:]
						}
					}
				}
			}

			if columnName == "" {
				columnName = snakeCase(f.Names[0].Name)
			}

			switch f.Type.(type) {
			case *ast.Ident:
				ident = f.Type.(*ast.Ident)
			case *ast.SelectorExpr:
				ident = f.Type.(*ast.SelectorExpr).Sel
			case *ast.StarExpr:
				ident = f.Type.(*ast.StarExpr).X.(*ast.SelectorExpr).Sel
			default:
				fmt.Printf("%#v\n", f.Type)
				panic("field type not supported")
			}

			field := StructField{
				AutoIncrement: autoIncrement,
				Converter:     converter,
				ColumnName:    columnName,
				Name:          f.Names[0].Name,
				Type:          ident.Name,
				Unique:        unique,
			}

			if autoIncrement {
				st.AutoIncrementField = &field
			}

			if primaryKey {
				st.PrimaryKey = &field
			}

			// if auto_increment, this key does not need to
			// be included in the st.Fields list
			st.Fields = append(st.Fields, field)
		}
		ctx.Structs = append(ctx.Structs, st)
	}

	return nil
}
