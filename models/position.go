// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Position position
//
// swagger:model Position
type Position struct {

	// typename
	// Enum: [position]
	Typename string `json:"__typename,omitempty"`

	// column
	Column *MaybeScalarsAtInt `json:"column,omitempty"`

	// row
	Row *MaybeScalarsAtInt `json:"row,omitempty"`
}

// Validate validates this position
func (m *Position) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateColumn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var positionTypeTypenamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["position"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		positionTypeTypenamePropEnum = append(positionTypeTypenamePropEnum, v)
	}
}

const (

	// PositionTypenamePosition captures enum value "position"
	PositionTypenamePosition string = "position"
)

// prop value enum
func (m *Position) validateTypenameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, positionTypeTypenamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Position) validateTypename(formats strfmt.Registry) error {
	if swag.IsZero(m.Typename) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypenameEnum("__typename", "body", m.Typename); err != nil {
		return err
	}

	return nil
}

func (m *Position) validateColumn(formats strfmt.Registry) error {
	if swag.IsZero(m.Column) { // not required
		return nil
	}

	if m.Column != nil {
		if err := m.Column.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("column")
			}
			return err
		}
	}

	return nil
}

func (m *Position) validateRow(formats strfmt.Registry) error {
	if swag.IsZero(m.Row) { // not required
		return nil
	}

	if m.Row != nil {
		if err := m.Row.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("row")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this position based on the context it is used
func (m *Position) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateColumn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Position) contextValidateColumn(ctx context.Context, formats strfmt.Registry) error {

	if m.Column != nil {
		if err := m.Column.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("column")
			}
			return err
		}
	}

	return nil
}

func (m *Position) contextValidateRow(ctx context.Context, formats strfmt.Registry) error {

	if m.Row != nil {
		if err := m.Row.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("row")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Position) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Position) UnmarshalBinary(b []byte) error {
	var res Position
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
