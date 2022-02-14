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

// NestedDiscoveredHostDisk nested discovered host disk
//
// swagger:model NestedDiscoveredHostDisk
type NestedDiscoveredHostDisk struct {

	// dimm ids
	DimmIds []string `json:"dimm_ids,omitempty"`

	// drive
	// Required: true
	Drive *string `json:"drive"`

	// function
	Function interface{} `json:"function,omitempty"`

	// model
	// Required: true
	Model *string `json:"model"`

	// numa node
	NumaNode *int32 `json:"numa_node,omitempty"`

	// persistent memory type
	PersistentMemoryType *string `json:"persistent_memory_type,omitempty"`

	// serial
	// Required: true
	Serial *string `json:"serial"`

	// size
	// Required: true
	Size *float64 `json:"size"`

	// type
	// Required: true
	Type *DiskType `json:"type"`
}

// Validate validates this nested discovered host disk
func (m *NestedDiscoveredHostDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDrive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedDiscoveredHostDisk) validateDrive(formats strfmt.Registry) error {

	if err := validate.Required("drive", "body", m.Drive); err != nil {
		return err
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateSerial(formats strfmt.Registry) error {

	if err := validate.Required("serial", "body", m.Serial); err != nil {
		return err
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *NestedDiscoveredHostDisk) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested discovered host disk based on the context it is used
func (m *NestedDiscoveredHostDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedDiscoveredHostDisk) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedDiscoveredHostDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedDiscoveredHostDisk) UnmarshalBinary(b []byte) error {
	var res NestedDiscoveredHostDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
