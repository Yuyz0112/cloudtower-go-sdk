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

// VMUpdateParams Vm update params
//
// swagger:model VmUpdateParams
type VMUpdateParams struct {

	// cpu cores
	CPUCores float64 `json:"cpu_cores,omitempty"`

	// cpu sockets
	CPUSockets float64 `json:"cpu_sockets,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// ha
	Ha bool `json:"ha,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// memory
	Memory float64 `json:"memory,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// vcpu
	Vcpu float64 `json:"vcpu,omitempty"`
}

// Validate validates this Vm update params
func (m *VMUpdateParams) Validate(formats strfmt.Registry) error {
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

func (m *VMUpdateParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMUpdateParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Vm update params based on context it is used
func (m *VMUpdateParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMUpdateParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateParams) UnmarshalBinary(b []byte) error {
	var res VMUpdateParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}