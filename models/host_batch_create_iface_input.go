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

// HostBatchCreateIfaceInput host batch create iface input
//
// swagger:model HostBatchCreateIfaceInput
type HostBatchCreateIfaceInput struct {

	// function
	// Required: true
	Function *HostBatchCreateIfaceFunction `json:"function"`

	// ip
	// Required: true
	IP *string `json:"ip"`

	// name
	// Required: true
	Name []string `json:"name"`

	// netmask
	// Required: true
	Netmask *string `json:"netmask"`
}

// Validate validates this host batch create iface input
func (m *HostBatchCreateIfaceInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostBatchCreateIfaceInput) validateFunction(formats strfmt.Registry) error {

	if err := validate.Required("function", "body", m.Function); err != nil {
		return err
	}

	if err := validate.Required("function", "body", m.Function); err != nil {
		return err
	}

	if m.Function != nil {
		if err := m.Function.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("function")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("function")
			}
			return err
		}
	}

	return nil
}

func (m *HostBatchCreateIfaceInput) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *HostBatchCreateIfaceInput) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *HostBatchCreateIfaceInput) validateNetmask(formats strfmt.Registry) error {

	if err := validate.Required("netmask", "body", m.Netmask); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this host batch create iface input based on the context it is used
func (m *HostBatchCreateIfaceInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFunction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostBatchCreateIfaceInput) contextValidateFunction(ctx context.Context, formats strfmt.Registry) error {

	if m.Function != nil {
		if err := m.Function.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("function")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("function")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostBatchCreateIfaceInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostBatchCreateIfaceInput) UnmarshalBinary(b []byte) error {
	var res HostBatchCreateIfaceInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
