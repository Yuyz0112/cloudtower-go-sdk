// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedIscsiLunSnapshot nested iscsi lun snapshot
//
// swagger:model NestedIscsiLunSnapshot
type NestedIscsiLunSnapshot struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this nested iscsi lun snapshot
func (m *NestedIscsiLunSnapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedIscsiLunSnapshot) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NestedIscsiLunSnapshot) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested iscsi lun snapshot based on context it is used
func (m *NestedIscsiLunSnapshot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedIscsiLunSnapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedIscsiLunSnapshot) UnmarshalBinary(b []byte) error {
	var res NestedIscsiLunSnapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
