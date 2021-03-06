// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EverouteLicenseWhereInput everoute license where input
// Example: {"AND":"EverouteLicenseWhereInput[]","NOT":"EverouteLicenseWhereInput[]","OR":"EverouteLicenseWhereInput[]","code":"string","code_contains":"string","code_ends_with":"string","code_gt":"string","code_gte":"string","code_in":["string"],"code_lt":"string","code_lte":"string","code_not":"string","code_not_contains":"string","code_not_ends_with":"string","code_not_in":["string"],"code_not_starts_with":"string","code_starts_with":"string","expire_date":"string","expire_date_gt":"string","expire_date_gte":"string","expire_date_in":["string"],"expire_date_lt":"string","expire_date_lte":"string","expire_date_not":"string","expire_date_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","max_socket_num":0,"max_socket_num_gt":0,"max_socket_num_gte":0,"max_socket_num_in":[0],"max_socket_num_lt":0,"max_socket_num_lte":0,"max_socket_num_not":0,"max_socket_num_not_in":[0],"serial":"string","serial_contains":"string","serial_ends_with":"string","serial_gt":"string","serial_gte":"string","serial_in":["string"],"serial_lt":"string","serial_lte":"string","serial_not":"string","serial_not_contains":"string","serial_not_ends_with":"string","serial_not_in":["string"],"serial_not_starts_with":"string","serial_starts_with":"string","sign_date":"string","sign_date_gt":"string","sign_date_gte":"string","sign_date_in":["string"],"sign_date_lt":"string","sign_date_lte":"string","sign_date_not":"string","sign_date_not_in":["string"],"software_edition":"COMMUNITY","software_edition_in":["COMMUNITY"],"software_edition_not":"COMMUNITY","software_edition_not_in":["COMMUNITY"],"type":"PERPETUAL","type_in":["PERPETUAL"],"type_not":"PERPETUAL","type_not_in":["PERPETUAL"],"uid":"string","uid_contains":"string","uid_ends_with":"string","uid_gt":"string","uid_gte":"string","uid_in":["string"],"uid_lt":"string","uid_lte":"string","uid_not":"string","uid_not_contains":"string","uid_not_ends_with":"string","uid_not_in":["string"],"uid_not_starts_with":"string","uid_starts_with":"string"}
//
// swagger:model EverouteLicenseWhereInput
type EverouteLicenseWhereInput struct {

	// a n d
	AND []*EverouteLicenseWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*EverouteLicenseWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*EverouteLicenseWhereInput `json:"OR,omitempty"`

	// code
	Code *string `json:"code,omitempty"`

	// code contains
	CodeContains *string `json:"code_contains,omitempty"`

	// code ends with
	CodeEndsWith *string `json:"code_ends_with,omitempty"`

	// code gt
	CodeGt *string `json:"code_gt,omitempty"`

	// code gte
	CodeGte *string `json:"code_gte,omitempty"`

	// code in
	CodeIn []string `json:"code_in,omitempty"`

	// code lt
	CodeLt *string `json:"code_lt,omitempty"`

	// code lte
	CodeLte *string `json:"code_lte,omitempty"`

	// code not
	CodeNot *string `json:"code_not,omitempty"`

	// code not contains
	CodeNotContains *string `json:"code_not_contains,omitempty"`

	// code not ends with
	CodeNotEndsWith *string `json:"code_not_ends_with,omitempty"`

	// code not in
	CodeNotIn []string `json:"code_not_in,omitempty"`

	// code not starts with
	CodeNotStartsWith *string `json:"code_not_starts_with,omitempty"`

	// code starts with
	CodeStartsWith *string `json:"code_starts_with,omitempty"`

	// expire date
	ExpireDate *string `json:"expire_date,omitempty"`

	// expire date gt
	ExpireDateGt *string `json:"expire_date_gt,omitempty"`

	// expire date gte
	ExpireDateGte *string `json:"expire_date_gte,omitempty"`

	// expire date in
	ExpireDateIn []string `json:"expire_date_in,omitempty"`

	// expire date lt
	ExpireDateLt *string `json:"expire_date_lt,omitempty"`

	// expire date lte
	ExpireDateLte *string `json:"expire_date_lte,omitempty"`

	// expire date not
	ExpireDateNot *string `json:"expire_date_not,omitempty"`

	// expire date not in
	ExpireDateNotIn []string `json:"expire_date_not_in,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// max socket num
	MaxSocketNum *int32 `json:"max_socket_num,omitempty"`

	// max socket num gt
	MaxSocketNumGt *int32 `json:"max_socket_num_gt,omitempty"`

	// max socket num gte
	MaxSocketNumGte *int32 `json:"max_socket_num_gte,omitempty"`

	// max socket num in
	MaxSocketNumIn []int32 `json:"max_socket_num_in,omitempty"`

	// max socket num lt
	MaxSocketNumLt *int32 `json:"max_socket_num_lt,omitempty"`

	// max socket num lte
	MaxSocketNumLte *int32 `json:"max_socket_num_lte,omitempty"`

	// max socket num not
	MaxSocketNumNot *int32 `json:"max_socket_num_not,omitempty"`

	// max socket num not in
	MaxSocketNumNotIn []int32 `json:"max_socket_num_not_in,omitempty"`

	// serial
	Serial *string `json:"serial,omitempty"`

	// serial contains
	SerialContains *string `json:"serial_contains,omitempty"`

	// serial ends with
	SerialEndsWith *string `json:"serial_ends_with,omitempty"`

	// serial gt
	SerialGt *string `json:"serial_gt,omitempty"`

	// serial gte
	SerialGte *string `json:"serial_gte,omitempty"`

	// serial in
	SerialIn []string `json:"serial_in,omitempty"`

	// serial lt
	SerialLt *string `json:"serial_lt,omitempty"`

	// serial lte
	SerialLte *string `json:"serial_lte,omitempty"`

	// serial not
	SerialNot *string `json:"serial_not,omitempty"`

	// serial not contains
	SerialNotContains *string `json:"serial_not_contains,omitempty"`

	// serial not ends with
	SerialNotEndsWith *string `json:"serial_not_ends_with,omitempty"`

	// serial not in
	SerialNotIn []string `json:"serial_not_in,omitempty"`

	// serial not starts with
	SerialNotStartsWith *string `json:"serial_not_starts_with,omitempty"`

	// serial starts with
	SerialStartsWith *string `json:"serial_starts_with,omitempty"`

	// sign date
	SignDate *string `json:"sign_date,omitempty"`

	// sign date gt
	SignDateGt *string `json:"sign_date_gt,omitempty"`

	// sign date gte
	SignDateGte *string `json:"sign_date_gte,omitempty"`

	// sign date in
	SignDateIn []string `json:"sign_date_in,omitempty"`

	// sign date lt
	SignDateLt *string `json:"sign_date_lt,omitempty"`

	// sign date lte
	SignDateLte *string `json:"sign_date_lte,omitempty"`

	// sign date not
	SignDateNot *string `json:"sign_date_not,omitempty"`

	// sign date not in
	SignDateNotIn []string `json:"sign_date_not_in,omitempty"`

	// software edition
	SoftwareEdition *SoftwareEdition `json:"software_edition,omitempty"`

	// software edition in
	SoftwareEditionIn []SoftwareEdition `json:"software_edition_in,omitempty"`

	// software edition not
	SoftwareEditionNot *SoftwareEdition `json:"software_edition_not,omitempty"`

	// software edition not in
	SoftwareEditionNotIn []SoftwareEdition `json:"software_edition_not_in,omitempty"`

	// type
	Type *LicenseType `json:"type,omitempty"`

	// type in
	TypeIn []LicenseType `json:"type_in,omitempty"`

	// type not
	TypeNot *LicenseType `json:"type_not,omitempty"`

	// type not in
	TypeNotIn []LicenseType `json:"type_not_in,omitempty"`

	// uid
	UID *string `json:"uid,omitempty"`

	// uid contains
	UIDContains *string `json:"uid_contains,omitempty"`

	// uid ends with
	UIDEndsWith *string `json:"uid_ends_with,omitempty"`

	// uid gt
	UIDGt *string `json:"uid_gt,omitempty"`

	// uid gte
	UIDGte *string `json:"uid_gte,omitempty"`

	// uid in
	UIDIn []string `json:"uid_in,omitempty"`

	// uid lt
	UIDLt *string `json:"uid_lt,omitempty"`

	// uid lte
	UIDLte *string `json:"uid_lte,omitempty"`

	// uid not
	UIDNot *string `json:"uid_not,omitempty"`

	// uid not contains
	UIDNotContains *string `json:"uid_not_contains,omitempty"`

	// uid not ends with
	UIDNotEndsWith *string `json:"uid_not_ends_with,omitempty"`

	// uid not in
	UIDNotIn []string `json:"uid_not_in,omitempty"`

	// uid not starts with
	UIDNotStartsWith *string `json:"uid_not_starts_with,omitempty"`

	// uid starts with
	UIDStartsWith *string `json:"uid_starts_with,omitempty"`
}

// Validate validates this everoute license where input
func (m *EverouteLicenseWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEdition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEditionIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEditionNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEditionNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EverouteLicenseWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateSoftwareEdition(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEdition) { // not required
		return nil
	}

	if m.SoftwareEdition != nil {
		if err := m.SoftwareEdition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateSoftwareEditionIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEditionIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareEditionIn); i++ {

		if err := m.SoftwareEditionIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateSoftwareEditionNot(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEditionNot) { // not required
		return nil
	}

	if m.SoftwareEditionNot != nil {
		if err := m.SoftwareEditionNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateSoftwareEditionNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEditionNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareEditionNotIn); i++ {

		if err := m.SoftwareEditionNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateTypeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNot) { // not required
		return nil
	}

	if m.TypeNot != nil {
		if err := m.TypeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicenseWhereInput) validateTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this everoute license where input based on the context it is used
func (m *EverouteLicenseWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEdition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEditionIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEditionNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEditionNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateSoftwareEdition(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareEdition != nil {
		if err := m.SoftwareEdition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateSoftwareEditionIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SoftwareEditionIn); i++ {

		if err := m.SoftwareEditionIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateSoftwareEditionNot(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareEditionNot != nil {
		if err := m.SoftwareEditionNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateSoftwareEditionNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SoftwareEditionNotIn); i++ {

		if err := m.SoftwareEditionNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateTypeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.TypeNot != nil {
		if err := m.TypeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not")
			}
			return err
		}
	}

	return nil
}

func (m *EverouteLicenseWhereInput) contextValidateTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EverouteLicenseWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EverouteLicenseWhereInput) UnmarshalBinary(b []byte) error {
	var res EverouteLicenseWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
