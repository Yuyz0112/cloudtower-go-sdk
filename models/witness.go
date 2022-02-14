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

// Witness witness
//
// swagger:model Witness
type Witness struct {

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// cpu hz per core
	// Required: true
	CPUHzPerCore *float64 `json:"cpu_hz_per_core"`

	// data ip
	// Required: true
	DataIP *string `json:"data_ip"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// system data capacity
	// Required: true
	SystemDataCapacity *float64 `json:"system_data_capacity"`

	// system used data space
	// Required: true
	SystemUsedDataSpace *float64 `json:"system_used_data_space"`

	// total cpu cores
	// Required: true
	TotalCPUCores *int32 `json:"total_cpu_cores"`

	// total cpu hz
	// Required: true
	TotalCPUHz *float64 `json:"total_cpu_hz"`

	// total memory bytes
	// Required: true
	TotalMemoryBytes *float64 `json:"total_memory_bytes"`
}

// Validate validates this witness
func (m *Witness) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUHzPerCore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemDataCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemUsedDataSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCPUHz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalMemoryBytes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Witness) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *Witness) validateCPUHzPerCore(formats strfmt.Registry) error {

	if err := validate.Required("cpu_hz_per_core", "body", m.CPUHzPerCore); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateDataIP(formats strfmt.Registry) error {

	if err := validate.Required("data_ip", "body", m.DataIP); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateSystemDataCapacity(formats strfmt.Registry) error {

	if err := validate.Required("system_data_capacity", "body", m.SystemDataCapacity); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateSystemUsedDataSpace(formats strfmt.Registry) error {

	if err := validate.Required("system_used_data_space", "body", m.SystemUsedDataSpace); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateTotalCPUCores(formats strfmt.Registry) error {

	if err := validate.Required("total_cpu_cores", "body", m.TotalCPUCores); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateTotalCPUHz(formats strfmt.Registry) error {

	if err := validate.Required("total_cpu_hz", "body", m.TotalCPUHz); err != nil {
		return err
	}

	return nil
}

func (m *Witness) validateTotalMemoryBytes(formats strfmt.Registry) error {

	if err := validate.Required("total_memory_bytes", "body", m.TotalMemoryBytes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this witness based on the context it is used
func (m *Witness) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Witness) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Witness) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Witness) UnmarshalBinary(b []byte) error {
	var res Witness
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
