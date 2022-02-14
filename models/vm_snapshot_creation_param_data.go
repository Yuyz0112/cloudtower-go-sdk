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

// VMSnapshotCreationParamData Vm snapshot creation param data
//
// swagger:model VmSnapshotCreationParamData
type VMSnapshotCreationParamData struct {

	// consistent type
	ConsistentType ConsistentType `json:"consistent_type,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// vm id
	// Required: true
	VMID *string `json:"vm_id"`
}

// Validate validates this Vm snapshot creation param data
func (m *VMSnapshotCreationParamData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsistentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMSnapshotCreationParamData) validateConsistentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistentType) { // not required
		return nil
	}

	if err := m.ConsistentType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("consistent_type")
		}
		return err
	}

	return nil
}

func (m *VMSnapshotCreationParamData) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMSnapshotCreationParamData) validateVMID(formats strfmt.Registry) error {

	if err := validate.Required("vm_id", "body", m.VMID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Vm snapshot creation param data based on the context it is used
func (m *VMSnapshotCreationParamData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsistentType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMSnapshotCreationParamData) contextValidateConsistentType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ConsistentType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("consistent_type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMSnapshotCreationParamData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMSnapshotCreationParamData) UnmarshalBinary(b []byte) error {
	var res VMSnapshotCreationParamData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
