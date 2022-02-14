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

// GlobalSettings global settings
//
// swagger:model GlobalSettings
type GlobalSettings struct {

	// auth
	Auth struct {
		NestedAuthSettings
	} `json:"auth,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// vm recycle bin
	// Required: true
	VMRecycleBin *NestedVMRecycleBin `json:"vm_recycle_bin"`
}

// Validate validates this global settings
func (m *GlobalSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMRecycleBin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalSettings) validateAuth(formats strfmt.Registry) error {
	if swag.IsZero(m.Auth) { // not required
		return nil
	}

	return nil
}

func (m *GlobalSettings) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GlobalSettings) validateVMRecycleBin(formats strfmt.Registry) error {

	if err := validate.Required("vm_recycle_bin", "body", m.VMRecycleBin); err != nil {
		return err
	}

	if m.VMRecycleBin != nil {
		if err := m.VMRecycleBin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_recycle_bin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_recycle_bin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this global settings based on the context it is used
func (m *GlobalSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMRecycleBin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GlobalSettings) contextValidateAuth(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GlobalSettings) contextValidateVMRecycleBin(ctx context.Context, formats strfmt.Registry) error {

	if m.VMRecycleBin != nil {
		if err := m.VMRecycleBin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_recycle_bin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_recycle_bin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GlobalSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GlobalSettings) UnmarshalBinary(b []byte) error {
	var res GlobalSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
