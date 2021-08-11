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

// VMVolumeCreationParams Vm volume creation params
//
// swagger:model VmVolumeCreationParams
type VMVolumeCreationParams struct {

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// elf storage policy
	ElfStoragePolicy VMVolumeElfStoragePolicyType `json:"elf_storage_policy,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// sharing
	Sharing bool `json:"sharing,omitempty"`

	// size
	Size float64 `json:"size,omitempty"`
}

// Validate validates this Vm volume creation params
func (m *VMVolumeCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfStoragePolicy(formats); err != nil {
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

func (m *VMVolumeCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VMVolumeCreationParams) validateElfStoragePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ElfStoragePolicy) { // not required
		return nil
	}

	if err := m.ElfStoragePolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("elf_storage_policy")
		}
		return err
	}

	return nil
}

func (m *VMVolumeCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Vm volume creation params based on the context it is used
func (m *VMVolumeCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElfStoragePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMVolumeCreationParams) contextValidateElfStoragePolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ElfStoragePolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("elf_storage_policy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMVolumeCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMVolumeCreationParams) UnmarshalBinary(b []byte) error {
	var res VMVolumeCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
