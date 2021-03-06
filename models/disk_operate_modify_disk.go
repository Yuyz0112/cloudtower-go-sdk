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

// DiskOperateModifyDisk disk operate modify disk
//
// swagger:model DiskOperateModifyDisk
type DiskOperateModifyDisk struct {

	// bus
	Bus Bus `json:"bus,omitempty"`

	// disk index
	// Required: true
	DiskIndex *int32 `json:"disk_index"`

	// vm volume id
	VMVolumeID string `json:"vm_volume_id,omitempty"`
}

// Validate validates this disk operate modify disk
func (m *DiskOperateModifyDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiskOperateModifyDisk) validateBus(formats strfmt.Registry) error {
	if swag.IsZero(m.Bus) { // not required
		return nil
	}

	if err := m.Bus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("bus")
		}
		return err
	}

	return nil
}

func (m *DiskOperateModifyDisk) validateDiskIndex(formats strfmt.Registry) error {

	if err := validate.Required("disk_index", "body", m.DiskIndex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this disk operate modify disk based on the context it is used
func (m *DiskOperateModifyDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiskOperateModifyDisk) contextValidateBus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Bus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("bus")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DiskOperateModifyDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiskOperateModifyDisk) UnmarshalBinary(b []byte) error {
	var res DiskOperateModifyDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
