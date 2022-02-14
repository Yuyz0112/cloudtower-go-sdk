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

// UserLoginInfoWhereInput user login info where input
// Example: {"AND":"UserLoginInfoWhereInput[]","NOT":"UserLoginInfoWhereInput[]","OR":"UserLoginInfoWhereInput[]","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","locked_at":"string","locked_at_gt":"string","locked_at_gte":"string","locked_at_in":["string"],"locked_at_lt":"string","locked_at_lte":"string","locked_at_not":"string","locked_at_not_in":["string"],"miss_num":0,"miss_num_gt":0,"miss_num_gte":0,"miss_num_in":[0],"miss_num_lt":0,"miss_num_lte":0,"miss_num_not":0,"miss_num_not_in":[0],"missed_at":"string","missed_at_gt":"string","missed_at_gte":"string","missed_at_in":["string"],"missed_at_lt":"string","missed_at_lte":"string","missed_at_not":"string","missed_at_not_in":["string"],"user":"UserWhereInput"}
//
// swagger:model UserLoginInfoWhereInput
type UserLoginInfoWhereInput struct {

	// a n d
	AND []*UserLoginInfoWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*UserLoginInfoWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*UserLoginInfoWhereInput `json:"OR,omitempty"`

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

	// locked at
	LockedAt *string `json:"locked_at,omitempty"`

	// locked at gt
	LockedAtGt *string `json:"locked_at_gt,omitempty"`

	// locked at gte
	LockedAtGte *string `json:"locked_at_gte,omitempty"`

	// locked at in
	LockedAtIn []string `json:"locked_at_in,omitempty"`

	// locked at lt
	LockedAtLt *string `json:"locked_at_lt,omitempty"`

	// locked at lte
	LockedAtLte *string `json:"locked_at_lte,omitempty"`

	// locked at not
	LockedAtNot *string `json:"locked_at_not,omitempty"`

	// locked at not in
	LockedAtNotIn []string `json:"locked_at_not_in,omitempty"`

	// miss num
	MissNum *int32 `json:"miss_num,omitempty"`

	// miss num gt
	MissNumGt *int32 `json:"miss_num_gt,omitempty"`

	// miss num gte
	MissNumGte *int32 `json:"miss_num_gte,omitempty"`

	// miss num in
	MissNumIn []int32 `json:"miss_num_in,omitempty"`

	// miss num lt
	MissNumLt *int32 `json:"miss_num_lt,omitempty"`

	// miss num lte
	MissNumLte *int32 `json:"miss_num_lte,omitempty"`

	// miss num not
	MissNumNot *int32 `json:"miss_num_not,omitempty"`

	// miss num not in
	MissNumNotIn []int32 `json:"miss_num_not_in,omitempty"`

	// missed at
	MissedAt *string `json:"missed_at,omitempty"`

	// missed at gt
	MissedAtGt *string `json:"missed_at_gt,omitempty"`

	// missed at gte
	MissedAtGte *string `json:"missed_at_gte,omitempty"`

	// missed at in
	MissedAtIn []string `json:"missed_at_in,omitempty"`

	// missed at lt
	MissedAtLt *string `json:"missed_at_lt,omitempty"`

	// missed at lte
	MissedAtLte *string `json:"missed_at_lte,omitempty"`

	// missed at not
	MissedAtNot *string `json:"missed_at_not,omitempty"`

	// missed at not in
	MissedAtNotIn []string `json:"missed_at_not_in,omitempty"`

	// user
	User *UserWhereInput `json:"user,omitempty"`
}

// Validate validates this user login info where input
func (m *UserLoginInfoWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserLoginInfoWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *UserLoginInfoWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *UserLoginInfoWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *UserLoginInfoWhereInput) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user login info where input based on the context it is used
func (m *UserLoginInfoWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserLoginInfoWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserLoginInfoWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserLoginInfoWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserLoginInfoWhereInput) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserLoginInfoWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserLoginInfoWhereInput) UnmarshalBinary(b []byte) error {
	var res UserLoginInfoWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
