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

// NestedBackupRestoreExecutionNetworkMapping nested backup restore execution network mapping
//
// swagger:model NestedBackupRestoreExecutionNetworkMapping
type NestedBackupRestoreExecutionNetworkMapping struct {

	// dst vlan id
	// Required: true
	DstVlanID *string `json:"dst_vlan_id"`

	// src vlan id
	// Required: true
	SrcVlanID *string `json:"src_vlan_id"`
}

// Validate validates this nested backup restore execution network mapping
func (m *NestedBackupRestoreExecutionNetworkMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDstVlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrcVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedBackupRestoreExecutionNetworkMapping) validateDstVlanID(formats strfmt.Registry) error {

	if err := validate.Required("dst_vlan_id", "body", m.DstVlanID); err != nil {
		return err
	}

	return nil
}

func (m *NestedBackupRestoreExecutionNetworkMapping) validateSrcVlanID(formats strfmt.Registry) error {

	if err := validate.Required("src_vlan_id", "body", m.SrcVlanID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested backup restore execution network mapping based on context it is used
func (m *NestedBackupRestoreExecutionNetworkMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedBackupRestoreExecutionNetworkMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedBackupRestoreExecutionNetworkMapping) UnmarshalBinary(b []byte) error {
	var res NestedBackupRestoreExecutionNetworkMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}