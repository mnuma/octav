// DO NOT EDIT. Automatically generated by hsup at Wed, 09 Mar 2016 12:17:34 JST
package validator

import (
	"github.com/lestrrat/go-jsval"
)

var HTTPCreateConferenceRequest *jsval.JSVal
var HTTPCreateConferenceResponse *jsval.JSVal
var HTTPCreateRoomRequest *jsval.JSVal
var HTTPCreateRoomResponse *jsval.JSVal
var HTTPCreateSessionRequest *jsval.JSVal
var HTTPCreateSessionResponse *jsval.JSVal
var HTTPCreateUserRequest *jsval.JSVal
var HTTPCreateUserResponse *jsval.JSVal
var HTTPCreateVenueRequest *jsval.JSVal
var HTTPCreateVenueResponse *jsval.JSVal
var HTTPDeleteConferenceRequest *jsval.JSVal
var HTTPDeleteRoomRequest *jsval.JSVal
var HTTPDeleteUserRequest *jsval.JSVal
var HTTPDeleteVenueRequest *jsval.JSVal
var HTTPListRoomsRequest *jsval.JSVal
var HTTPListRoomsResponse *jsval.JSVal
var HTTPListSessionsByConferenceRequest *jsval.JSVal
var HTTPListSessionsByConferenceResponse *jsval.JSVal
var HTTPListVenuesRequest *jsval.JSVal
var HTTPListVenuesResponse *jsval.JSVal
var HTTPLookupConferenceRequest *jsval.JSVal
var HTTPLookupConferenceResponse *jsval.JSVal
var HTTPLookupRoomRequest *jsval.JSVal
var HTTPLookupRoomResponse *jsval.JSVal
var HTTPLookupSessionRequest *jsval.JSVal
var HTTPLookupSessionResponse *jsval.JSVal
var HTTPLookupUserRequest *jsval.JSVal
var HTTPLookupUserResponse *jsval.JSVal
var HTTPLookupVenueRequest *jsval.JSVal
var HTTPLookupVenueResponse *jsval.JSVal
var M *jsval.ConstraintMap
var R0 jsval.Constraint
var R1 jsval.Constraint
var R2 jsval.Constraint
var R3 jsval.Constraint
var R4 jsval.Constraint
var R5 jsval.Constraint
var R6 jsval.Constraint
var R7 jsval.Constraint
var R8 jsval.Constraint
var R9 jsval.Constraint
var R10 jsval.Constraint
var R11 jsval.Constraint
var R12 jsval.Constraint
var R13 jsval.Constraint
var R14 jsval.Constraint
var R15 jsval.Constraint
var R16 jsval.Constraint
var R17 jsval.Constraint
var R18 jsval.Constraint
var R19 jsval.Constraint
var R20 jsval.Constraint
var R21 jsval.Constraint
var R22 jsval.Constraint
var R23 jsval.Constraint
var R24 jsval.Constraint
var R25 jsval.Constraint
var R26 jsval.Constraint
var R27 jsval.Constraint
var R28 jsval.Constraint
var R29 jsval.Constraint

func init() {
	M = &jsval.ConstraintMap{}
	R0 = jsval.String().Enum("pending", "accepted", "rejected").Default("pending")
	R1 = jsval.String().Enum("allow", "disallow")
	R2 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		).
		AddProp(
			`dates`,
			jsval.Reference(M).RefersTo(`#/definitions/date-array`),
		).
		AddProp(
			`description`,
			jsval.Reference(M).RefersTo(`#/definitions/string-en`),
		).
		AddProp(
			`id`,
			jsval.Reference(M).RefersTo(`#/definitions/uuid`),
		).
		AddProp(
			`name`,
			jsval.Reference(M).RefersTo(`#/definitions/string-en`),
		).
		AddProp(
			`slug`,
			jsval.Reference(M).RefersTo(`#/definitions/string-en`),
		).
		AddProp(
			`venue`,
			jsval.Reference(M).RefersTo(`#/definitions/venue`),
		).
		PatternPropertiesString(
			"description#[a-z]+",
			jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
		).
		PatternPropertiesString(
			"title#[a-z]+",
			jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
		)
	R3 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		)
	R4 = jsval.Array().
		Items(
			jsval.Reference(M).RefersTo(`#/definitions/date`),
		).
		AdditionalItems(
			jsval.EmptyConstraint,
		)
	R5 = jsval.String().RegexpString(`^\d+-\d+-\d+$`)
	R6 = jsval.String()
	R7 = jsval.String()
	R8 = jsval.String().Format("email")
	R9 = jsval.String().Default("en")
	R10 = jsval.Number()
	R11 = jsval.Number()
	R12 = jsval.String()
	R13 = jsval.String()
	R14 = jsval.String().Enum("beginner", "intermediate", "advanced")
	R15 = jsval.Integer().Minimum(0)
	R16 = jsval.Integer().Minimum(0).Default(10)
	R17 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		).
		AddProp(
			`capcity`,
			jsval.Reference(M).RefersTo(`#/definitions/positiveInteger`),
		).
		AddProp(
			`id`,
			jsval.Reference(M).RefersTo(`#/definitions/uuid`),
		).
		AddProp(
			`name`,
			jsval.String(),
		).
		AddProp(
			`venue_id`,
			jsval.Reference(M).RefersTo(`#/definitions/uuid`),
		).
		PatternPropertiesString(
			"name#[a-z]+",
			jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
		)
	R18 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		).
		AddProp(
			`abstract`,
			jsval.String(),
		).
		AddProp(
			`category`,
			jsval.String(),
		).
		AddProp(
			`conference`,
			jsval.Reference(M).RefersTo(`#/definitions/conference`),
		).
		AddProp(
			`confirmed`,
			jsval.Boolean().Default(false),
		).
		AddProp(
			`duration`,
			jsval.Reference(M).RefersTo(`#/definitions/duration`),
		).
		AddProp(
			`has-interpretation`,
			jsval.Boolean().Default(false),
		).
		AddProp(
			`material-level`,
			jsval.Reference(M).RefersTo(`#/definitions/material-level`),
		).
		AddProp(
			`memo`,
			jsval.String(),
		).
		AddProp(
			`photo-permission`,
			jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
		).
		AddProp(
			`room`,
			jsval.Reference(M).RefersTo(`#/definitions/room`),
		).
		AddProp(
			`slide-language`,
			jsval.Reference(M).RefersTo(`#/definitions/language`),
		).
		AddProp(
			`slide-subtitles`,
			jsval.Reference(M).RefersTo(`#/definitions/language`),
		).
		AddProp(
			`slide-url`,
			jsval.Reference(M).RefersTo(`#/definitions/url`),
		).
		AddProp(
			`sort-order`,
			jsval.Reference(M).RefersTo(`#/definitions/positiveInteger`),
		).
		AddProp(
			`speaker`,
			jsval.Reference(M).RefersTo(`#/definitions/speaker-array`),
		).
		AddProp(
			`spoken-language`,
			jsval.Reference(M).RefersTo(`#/definitions/language`),
		).
		AddProp(
			`starts-on`,
			jsval.Reference(M).RefersTo(`#/definitions/datetime`),
		).
		AddProp(
			`status`,
			jsval.Reference(M).RefersTo(`#/definitions/acceptance-status`),
		).
		AddProp(
			`tags`,
			jsval.Reference(M).RefersTo(`#/definitions/tag-array`),
		).
		AddProp(
			`title`,
			jsval.String(),
		).
		AddProp(
			`video-permission`,
			jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
		).
		AddProp(
			`video-url`,
			jsval.Reference(M).RefersTo(`#/definitions/url`),
		)
	R19 = jsval.Object().
		AdditionalProperties(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				PatternPropertiesString(
					"^[a-z0-9-]+-profile",
					jsval.Object().
						AdditionalProperties(
							jsval.EmptyConstraint,
						),
				),
		).
		AddProp(
			`email`,
			jsval.Reference(M).RefersTo(`#/definitions/email`),
		).
		AddProp(
			`id`,
			jsval.Reference(M).RefersTo(`#/definitions/uuid`),
		).
		AddProp(
			`name`,
			jsval.String(),
		)
	R20 = jsval.Array().
		Items(
			jsval.Reference(M).RefersTo(`#/definitions/speaker`),
		).
		AdditionalItems(
			jsval.EmptyConstraint,
		)
	R21 = jsval.String()
	R22 = jsval.String()
	R23 = jsval.String()
	R24 = jsval.Array().
		Items(
			jsval.Reference(M).RefersTo(`#/definitions/tag`),
		).
		AdditionalItems(
			jsval.EmptyConstraint,
		)
	R25 = jsval.String().Enum("XXXL", "XXL", "XL", "L", "M", "S", "XS")
	R26 = jsval.String().Format("url")
	R27 = jsval.String().RegexpString(`^[a-fA-F0-9-]+$`)
	R28 = jsval.String().RegexpString(`^[a-fA-F0-9-]+$`).Default("")
	R29 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
		).
		AddProp(
			`id`,
			jsval.Reference(M).RefersTo(`#/definitions/uuid`),
		).
		AddProp(
			`name`,
			jsval.String(),
		).
		AddProp(
			`rooms`,
			jsval.Reference(M).RefersTo(`#/definitions/room`),
		)
	M.SetReference(`#/definitions/acceptance-status`, R0)
	M.SetReference(`#/definitions/binary-permission`, R1)
	M.SetReference(`#/definitions/conference`, R2)
	M.SetReference(`#/definitions/date`, R3)
	M.SetReference(`#/definitions/date-array`, R4)
	M.SetReference(`#/definitions/date-string`, R5)
	M.SetReference(`#/definitions/datetime`, R6)
	M.SetReference(`#/definitions/duration`, R7)
	M.SetReference(`#/definitions/email`, R8)
	M.SetReference(`#/definitions/language`, R9)
	M.SetReference(`#/definitions/latitude`, R10)
	M.SetReference(`#/definitions/longitude`, R11)
	M.SetReference(`#/definitions/markdown-en`, R12)
	M.SetReference(`#/definitions/markdown-i18n`, R13)
	M.SetReference(`#/definitions/material-level`, R14)
	M.SetReference(`#/definitions/positiveInteger`, R15)
	M.SetReference(`#/definitions/positiveIntegerDefault10`, R16)
	M.SetReference(`#/definitions/room`, R17)
	M.SetReference(`#/definitions/session`, R18)
	M.SetReference(`#/definitions/speaker`, R19)
	M.SetReference(`#/definitions/speaker-array`, R20)
	M.SetReference(`#/definitions/string-en`, R21)
	M.SetReference(`#/definitions/string-i18n`, R22)
	M.SetReference(`#/definitions/tag`, R23)
	M.SetReference(`#/definitions/tag-array`, R24)
	M.SetReference(`#/definitions/tshirt_size`, R25)
	M.SetReference(`#/definitions/url`, R26)
	M.SetReference(`#/definitions/uuid`, R27)
	M.SetReference(`#/definitions/uuidDefaultEmpty`, R28)
	M.SetReference(`#/definitions/venue`, R29)
	HTTPCreateConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("slug", "title").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`description`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`slug`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`starts-on`,
					jsval.Reference(M).RefersTo(`#/definitions/datetime`),
				).
				AddProp(
					`sub_title`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`title`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				PatternPropertiesString(
					"description#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPCreateConferenceResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`dates`,
					jsval.Reference(M).RefersTo(`#/definitions/date-array`),
				).
				AddProp(
					`description`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`slug`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`venue`,
					jsval.Reference(M).RefersTo(`#/definitions/venue`),
				).
				PatternPropertiesString(
					"description#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPCreateRoomRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("name", "venue_id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`capacity`,
					jsval.Reference(M).RefersTo(`#/definitions/positiveInteger`),
				).
				AddProp(
					`name`,
					jsval.String(),
				).
				AddProp(
					`venue_id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				PatternPropertiesString(
					"name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPCreateRoomResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`capcity`,
					jsval.Reference(M).RefersTo(`#/definitions/positiveInteger`),
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`name`,
					jsval.String(),
				).
				AddProp(
					`venue_id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				PatternPropertiesString(
					"name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPCreateSessionRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`abstract`,
					jsval.Reference(M).RefersTo(`#/definitions/markdown-en`),
				).
				AddProp(
					`category`,
					jsval.String(),
				).
				AddProp(
					`conference_id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`confirmed`,
					jsval.Boolean().Default(false),
				).
				AddProp(
					`duration`,
					jsval.Reference(M).RefersTo(`#/definitions/duration`),
				).
				AddProp(
					`has_interpretation`,
					jsval.Boolean().Default(false),
				).
				AddProp(
					`material_level`,
					jsval.Reference(M).RefersTo(`#/definitions/material-level`),
				).
				AddProp(
					`memo`,
					jsval.String(),
				).
				AddProp(
					`photo_permission`,
					jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
				).
				AddProp(
					`room_id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`slide_language`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`slide_subtitles`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`slide_url`,
					jsval.Reference(M).RefersTo(`#/definitions/url`),
				).
				AddProp(
					`sort_order`,
					jsval.Reference(M).RefersTo(`#/definitions/positiveInteger`),
				).
				AddProp(
					`speaker_id`,
					jsval.Array().
						Items(
							jsval.Reference(M).RefersTo(`#/definitions/uuid`),
						).
						AdditionalItems(
							jsval.EmptyConstraint,
						),
				).
				AddProp(
					`spoken_language`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`starts_on`,
					jsval.Reference(M).RefersTo(`#/definitions/datetime`),
				).
				AddProp(
					`status`,
					jsval.Reference(M).RefersTo(`#/definitions/acceptance-status`),
				).
				AddProp(
					`tag`,
					jsval.Reference(M).RefersTo(`#/definitions/tag-array`),
				).
				AddProp(
					`title`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`video_permission`,
					jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
				).
				AddProp(
					`video_url`,
					jsval.Reference(M).RefersTo(`#/definitions/url`),
				).
				PatternPropertiesString(
					"abstract#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/markdown-i18n`),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPCreateSessionResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`abstract`,
					jsval.String(),
				).
				AddProp(
					`category`,
					jsval.String(),
				).
				AddProp(
					`conference`,
					jsval.Reference(M).RefersTo(`#/definitions/conference`),
				).
				AddProp(
					`confirmed`,
					jsval.Boolean().Default(false),
				).
				AddProp(
					`duration`,
					jsval.Reference(M).RefersTo(`#/definitions/duration`),
				).
				AddProp(
					`has-interpretation`,
					jsval.Boolean().Default(false),
				).
				AddProp(
					`material-level`,
					jsval.Reference(M).RefersTo(`#/definitions/material-level`),
				).
				AddProp(
					`memo`,
					jsval.String(),
				).
				AddProp(
					`photo-permission`,
					jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
				).
				AddProp(
					`room`,
					jsval.Reference(M).RefersTo(`#/definitions/room`),
				).
				AddProp(
					`slide-language`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`slide-subtitles`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`slide-url`,
					jsval.Reference(M).RefersTo(`#/definitions/url`),
				).
				AddProp(
					`sort-order`,
					jsval.Reference(M).RefersTo(`#/definitions/positiveInteger`),
				).
				AddProp(
					`speaker`,
					jsval.Reference(M).RefersTo(`#/definitions/speaker-array`),
				).
				AddProp(
					`spoken-language`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`starts-on`,
					jsval.Reference(M).RefersTo(`#/definitions/datetime`),
				).
				AddProp(
					`status`,
					jsval.Reference(M).RefersTo(`#/definitions/acceptance-status`),
				).
				AddProp(
					`tags`,
					jsval.Reference(M).RefersTo(`#/definitions/tag-array`),
				).
				AddProp(
					`title`,
					jsval.String(),
				).
				AddProp(
					`video-permission`,
					jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
				).
				AddProp(
					`video-url`,
					jsval.Reference(M).RefersTo(`#/definitions/url`),
				),
		)

	HTTPCreateUserRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`email`,
					jsval.Reference(M).RefersTo(`#/definitions/email`),
				).
				AddProp(
					`first_name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`last_name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`nickname`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`tshirt_size`,
					jsval.Reference(M).RefersTo(`#/definitions/tshirt_size`),
				).
				PatternPropertiesString(
					"first_name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				).
				PatternPropertiesString(
					"last_name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPCreateUserResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`email`,
					jsval.Reference(M).RefersTo(`#/definitions/email`),
				).
				AddProp(
					`first_name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`last_name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`nickname`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`tshirt_size`,
					jsval.Reference(M).RefersTo(`#/definitions/tshirt_size`),
				).
				PatternPropertiesString(
					"first_name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				).
				PatternPropertiesString(
					"last_name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPCreateVenueRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("address", "name").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`address`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`latitude`,
					jsval.Reference(M).RefersTo(`#/definitions/latitude`),
				).
				AddProp(
					`longitude`,
					jsval.Reference(M).RefersTo(`#/definitions/longitude`),
				).
				AddProp(
					`name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				PatternPropertiesString(
					"address#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				).
				PatternPropertiesString(
					"name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPCreateVenueResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`name`,
					jsval.String(),
				).
				AddProp(
					`rooms`,
					jsval.Reference(M).RefersTo(`#/definitions/room`),
				),
		)

	HTTPDeleteConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPDeleteRoomRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPDeleteUserRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPDeleteVenueRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPListRoomsRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("venue_id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`lang`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`limit`,
					jsval.Reference(M).RefersTo(`#/definitions/positiveIntegerDefault10`),
				).
				AddProp(
					`venue_id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPListRoomsResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Array().
				Items(
					jsval.Reference(M).RefersTo(`#/definitions/room`),
				).
				AdditionalItems(
					jsval.EmptyConstraint,
				),
		)

	HTTPListSessionsByConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("conference_id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`conference_id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`date`,
					jsval.Reference(M).RefersTo(`#/definitions/date-string`),
				),
		)

	HTTPListSessionsByConferenceResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Array().
				Items(
					jsval.Reference(M).RefersTo(`#/definitions/session`),
				).
				AdditionalItems(
					jsval.EmptyConstraint,
				),
		)

	HTTPListVenuesRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`lang`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`since`,
					jsval.Reference(M).RefersTo(`#/definitions/uuidDefaultEmpty`),
				),
		)

	HTTPListVenuesResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Array().
				Items(
					jsval.Reference(M).RefersTo(`#/definitions/venue`),
				).
				AdditionalItems(
					jsval.EmptyConstraint,
				),
		)

	HTTPLookupConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPLookupConferenceResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`dates`,
					jsval.Reference(M).RefersTo(`#/definitions/date-array`),
				).
				AddProp(
					`description`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`slug`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`venue`,
					jsval.Reference(M).RefersTo(`#/definitions/venue`),
				).
				PatternPropertiesString(
					"description#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				).
				PatternPropertiesString(
					"title#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPLookupRoomRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPLookupRoomResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`capcity`,
					jsval.Reference(M).RefersTo(`#/definitions/positiveInteger`),
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`name`,
					jsval.String(),
				).
				AddProp(
					`venue_id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				PatternPropertiesString(
					"name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPLookupSessionRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPLookupSessionResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`abstract`,
					jsval.String(),
				).
				AddProp(
					`category`,
					jsval.String(),
				).
				AddProp(
					`conference`,
					jsval.Reference(M).RefersTo(`#/definitions/conference`),
				).
				AddProp(
					`confirmed`,
					jsval.Boolean().Default(false),
				).
				AddProp(
					`duration`,
					jsval.Reference(M).RefersTo(`#/definitions/duration`),
				).
				AddProp(
					`has-interpretation`,
					jsval.Boolean().Default(false),
				).
				AddProp(
					`material-level`,
					jsval.Reference(M).RefersTo(`#/definitions/material-level`),
				).
				AddProp(
					`memo`,
					jsval.String(),
				).
				AddProp(
					`photo-permission`,
					jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
				).
				AddProp(
					`room`,
					jsval.Reference(M).RefersTo(`#/definitions/room`),
				).
				AddProp(
					`slide-language`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`slide-subtitles`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`slide-url`,
					jsval.Reference(M).RefersTo(`#/definitions/url`),
				).
				AddProp(
					`sort-order`,
					jsval.Reference(M).RefersTo(`#/definitions/positiveInteger`),
				).
				AddProp(
					`speaker`,
					jsval.Reference(M).RefersTo(`#/definitions/speaker-array`),
				).
				AddProp(
					`spoken-language`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`starts-on`,
					jsval.Reference(M).RefersTo(`#/definitions/datetime`),
				).
				AddProp(
					`status`,
					jsval.Reference(M).RefersTo(`#/definitions/acceptance-status`),
				).
				AddProp(
					`tags`,
					jsval.Reference(M).RefersTo(`#/definitions/tag-array`),
				).
				AddProp(
					`title`,
					jsval.String(),
				).
				AddProp(
					`video-permission`,
					jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
				).
				AddProp(
					`video-url`,
					jsval.Reference(M).RefersTo(`#/definitions/url`),
				),
		)

	HTTPLookupUserRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPLookupUserResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`email`,
					jsval.Reference(M).RefersTo(`#/definitions/email`),
				).
				AddProp(
					`first_name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`last_name`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`nickname`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`tshirt_size`,
					jsval.Reference(M).RefersTo(`#/definitions/tshirt_size`),
				).
				PatternPropertiesString(
					"first_name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				).
				PatternPropertiesString(
					"last_name#[a-z]+",
					jsval.Reference(M).RefersTo(`#/definitions/string-i18n`),
				),
		)

	HTTPLookupVenueRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPLookupVenueResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				).
				AddProp(
					`name`,
					jsval.String(),
				).
				AddProp(
					`rooms`,
					jsval.Reference(M).RefersTo(`#/definitions/room`),
				),
		)

}
