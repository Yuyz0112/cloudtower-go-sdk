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

// IsolationPolicyWhereInput isolation policy where input
// Example: {"AND":"IsolationPolicyWhereInput[]","NOT":"IsolationPolicyWhereInput[]","OR":"IsolationPolicyWhereInput[]","everoute_cluster":"EverouteClusterWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","mode":"ALL","mode_in":["ALL"],"mode_not":"ALL","mode_not_in":["ALL"],"vm":"VmWhereInput"}
//
// swagger:model IsolationPolicyWhereInput
type IsolationPolicyWhereInput struct {

	// a n d
	AND []*IsolationPolicyWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*IsolationPolicyWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*IsolationPolicyWhereInput `json:"OR,omitempty"`

	// everoute cluster
	EverouteCluster interface{} `json:"everoute_cluster,omitempty"`

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

	// labels every
	LabelsEvery interface{} `json:"labels_every,omitempty"`

	// labels none
	LabelsNone interface{} `json:"labels_none,omitempty"`

	// labels some
	LabelsSome interface{} `json:"labels_some,omitempty"`

	// mode
	Mode interface{} `json:"mode,omitempty"`

	// mode in
	ModeIn []IsolationMode `json:"mode_in,omitempty"`

	// mode not
	ModeNot interface{} `json:"mode_not,omitempty"`

	// mode not in
	ModeNotIn []IsolationMode `json:"mode_not_in,omitempty"`

	// vm
	VM interface{} `json:"vm,omitempty"`
}

// Validate validates this isolation policy where input
func (m *IsolationPolicyWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateModeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsolationPolicyWhereInput) validateAND(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *IsolationPolicyWhereInput) validateNOT(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *IsolationPolicyWhereInput) validateOR(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *IsolationPolicyWhereInput) validateModeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ModeIn); i++ {

		if err := m.ModeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IsolationPolicyWhereInput) validateModeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ModeNotIn); i++ {

		if err := m.ModeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this isolation policy where input based on the context it is used
func (m *IsolationPolicyWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateModeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IsolationPolicyWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IsolationPolicyWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IsolationPolicyWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IsolationPolicyWhereInput) contextValidateModeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModeIn); i++ {

		if err := m.ModeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *IsolationPolicyWhereInput) contextValidateModeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModeNotIn); i++ {

		if err := m.ModeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IsolationPolicyWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IsolationPolicyWhereInput) UnmarshalBinary(b []byte) error {
	var res IsolationPolicyWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}