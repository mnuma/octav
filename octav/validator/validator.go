// DO NOT EDIT. Automatically generated by hsup at Thu, 03 Mar 2016 12:57:14 JST
package validator

import (
	"github.com/lestrrat/go-jsval"
)

var HTTPCreateRoomRequest *jsval.JSVal
var HTTPListRoomsRequest *jsval.JSVal
var HTTPCreateConferenceRequest *jsval.JSVal
var HTTPCreateSessionRequest *jsval.JSVal
var HTTPCreateUserRequest *jsval.JSVal
var HTTPCreateVenueRequest *jsval.JSVal
var HTTPListVenuesRequest *jsval.JSVal
var HTTPCreateSessionResponse *jsval.JSVal
var HTTPCreateUserResponse *jsval.JSVal
var HTTPCreateVenueResponse *jsval.JSVal
var HTTPListVenuesResponse *jsval.JSVal
var HTTPCreateRoomResponse *jsval.JSVal
var HTTPListRoomsResponse *jsval.JSVal
var HTTPCreateConferenceResponse *jsval.JSVal
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

func init() {
	M = &jsval.ConstraintMap{}
	R0 = jsval.String().Enum("pending", "accepted", "rejected")
	R1 = jsval.String().Enum("allow", "disallow")
	R2 = jsval.Object().
		AdditionalProperties(
			jsval.EmptyConstraint,
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
			`venue`,
			jsval.Reference(M).RefersTo(`#/definitions/venue`),
		)
	R3 = jsval.String()
	R4 = jsval.String()
	R5 = jsval.String().Format("email")
	R6 = jsval.String()
	R7 = jsval.Number()
	R8 = jsval.Number()
	R9 = jsval.String()
	R10 = jsval.String().Enum("beginner", "intermediate", "advanced")
	R11 = jsval.Integer().Minimum(0)
	R12 = jsval.Object().
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
			`venue-id`,
			jsval.Reference(M).RefersTo(`#/definitions/uuid`),
		)
	R13 = jsval.Array().
		Items(
			jsval.Reference(M).RefersTo(`#/definitions/speaker`),
		).
		AdditionalItems(
			jsval.EmptyConstraint,
		)
	R14 = jsval.String()
	R15 = jsval.String()
	R16 = jsval.String()
	R17 = jsval.Array().
		Items(
			jsval.Reference(M).RefersTo(`#/definitions/tag`),
		).
		AdditionalItems(
			jsval.EmptyConstraint,
		)
	R18 = jsval.String().Enum("XXXL", "XXL", "XL", "L", "M", "S", "XS")
	R19 = jsval.String().Format("url")
	R20 = jsval.String().RegexpString(`^[a-fA-F0-9-]+$`)
	R21 = jsval.Object().
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
	M.SetReference(`#/definitions/datetime`, R3)
	M.SetReference(`#/definitions/duration`, R4)
	M.SetReference(`#/definitions/email`, R5)
	M.SetReference(`#/definitions/language`, R6)
	M.SetReference(`#/definitions/latitude`, R7)
	M.SetReference(`#/definitions/longitude`, R8)
	M.SetReference(`#/definitions/markdown-en`, R9)
	M.SetReference(`#/definitions/material-level`, R10)
	M.SetReference(`#/definitions/positiveInteger`, R11)
	M.SetReference(`#/definitions/room`, R12)
	M.SetReference(`#/definitions/speaker-array`, R13)
	M.SetReference(`#/definitions/string-en`, R14)
	M.SetReference(`#/definitions/string-i18n`, R15)
	M.SetReference(`#/definitions/tag`, R16)
	M.SetReference(`#/definitions/tag-array`, R17)
	M.SetReference(`#/definitions/tshirt_size`, R18)
	M.SetReference(`#/definitions/url`, R19)
	M.SetReference(`#/definitions/uuid`, R20)
	M.SetReference(`#/definitions/venue`, R21)
	HTTPCreateRoomRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("name", "venue-id").
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
					`venue-id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPListRoomsRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("venue-id").
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`lang`,
					jsval.Reference(M).RefersTo(`#/definitions/language`),
				).
				AddProp(
					`venue-id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
				),
		)

	HTTPCreateConferenceRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
				).
				AddProp(
					`description`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`duration`,
					jsval.Reference(M).RefersTo(`#/definitions/duration`),
				).
				AddProp(
					`starts-on`,
					jsval.Reference(M).RefersTo(`#/definitions/datetime`),
				).
				AddProp(
					`title`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
				).
				AddProp(
					`venue-id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
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
					`conference-id`,
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
					`patternProperties`,
					jsval.EmptyConstraint,
				).
				AddProp(
					`photo-permission`,
					jsval.Reference(M).RefersTo(`#/definitions/binary-permission`),
				).
				AddProp(
					`room-id`,
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
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
					`speaker-id`,
					jsval.Array().
						Items(
							jsval.Reference(M).RefersTo(`#/definitions/uuid`),
						).
						AdditionalItems(
							jsval.EmptyConstraint,
						),
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
					`tag`,
					jsval.Reference(M).RefersTo(`#/definitions/tag-array`),
				).
				AddProp(
					`title`,
					jsval.Reference(M).RefersTo(`#/definitions/string-en`),
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
				),
		)

	HTTPCreateVenueRequest = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				Required("address", "latitude", "longitude", "name").
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
					jsval.Reference(M).RefersTo(`#/definitions/uuid`),
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
					`tag`,
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
					`venue-id`,
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

	HTTPCreateConferenceResponse = jsval.New().
		SetConstraintMap(M).
		SetRoot(
			jsval.Object().
				AdditionalProperties(
					jsval.EmptyConstraint,
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
					`venue`,
					jsval.Reference(M).RefersTo(`#/definitions/venue`),
				),
		)

}
