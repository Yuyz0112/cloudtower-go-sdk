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

// ViewUpdationParams view updation params
//
// swagger:model ViewUpdationParams
type ViewUpdationParams struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	Name string `json:"name,omitempty"`

	// time span
	TimeSpan float64 `json:"time_span,omitempty"`

	// time unit
	TimeUnit TimeUnit `json:"time_unit,omitempty"`
}

// Validate validates this view updation params
func (m *ViewUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewUpdationParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ViewUpdationParams) validateTimeUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeUnit) { // not required
		return nil
	}

	if err := m.TimeUnit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time_unit")
		}
		return err
	}

	return nil
}

// ContextValidate validate this view updation params based on the context it is used
func (m *ViewUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTimeUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ViewUpdationParams) contextValidateTimeUnit(ctx context.Context, formats strfmt.Registry) error {

	if err := m.TimeUnit.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time_unit")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ViewUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ViewUpdationParams) UnmarshalBinary(b []byte) error {
	var res ViewUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
