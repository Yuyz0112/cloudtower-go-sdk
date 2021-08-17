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

// HostBatchCreateDiskInput host batch create disk input
//
// swagger:model HostBatchCreateDiskInput
type HostBatchCreateDiskInput struct {

	// drive
	// Required: true
	Drive *string `json:"drive"`

	// function
	Function MaybeDiskFunction `json:"function,omitempty"`

	// type
	Type MaybeDiskType `json:"type,omitempty"`
}

// Validate validates this host batch create disk input
func (m *HostBatchCreateDiskInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrive(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostBatchCreateDiskInput) validateDrive(formats strfmt.Registry) error {

	if err := validate.Required("drive", "body", m.Drive); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this host batch create disk input based on context it is used
func (m *HostBatchCreateDiskInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostBatchCreateDiskInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostBatchCreateDiskInput) UnmarshalBinary(b []byte) error {
	var res HostBatchCreateDiskInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
