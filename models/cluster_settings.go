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

// ClusterSettings cluster settings
//
// swagger:model ClusterSettings
type ClusterSettings struct {

	// cluster
	// Required: true
	Cluster *ClusterSettingsCluster `json:"cluster"`

	// id
	// Required: true
	ID *string `json:"id"`

	// vm recycle bin
	VMRecycleBin *ClusterSettingsVMRecycleBin `json:"vm_recycle_bin,omitempty"`
}

// Validate validates this cluster settings
func (m *ClusterSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
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

func (m *ClusterSettings) validateCluster(formats strfmt.Registry) error {

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

func (m *ClusterSettings) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSettings) validateVMRecycleBin(formats strfmt.Registry) error {
	if swag.IsZero(m.VMRecycleBin) { // not required
		return nil
	}

	if m.VMRecycleBin != nil {
		if err := m.VMRecycleBin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_recycle_bin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster settings based on the context it is used
func (m *ClusterSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
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

func (m *ClusterSettings) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterSettings) contextValidateVMRecycleBin(ctx context.Context, formats strfmt.Registry) error {

	if m.VMRecycleBin != nil {
		if err := m.VMRecycleBin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_recycle_bin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSettings) UnmarshalBinary(b []byte) error {
	var res ClusterSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSettingsCluster cluster settings cluster
//
// swagger:model ClusterSettingsCluster
type ClusterSettingsCluster struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this cluster settings cluster
func (m *ClusterSettingsCluster) Validate(formats strfmt.Registry) error {
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

func (m *ClusterSettingsCluster) validateID(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSettingsCluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cluster"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster settings cluster based on context it is used
func (m *ClusterSettingsCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSettingsCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSettingsCluster) UnmarshalBinary(b []byte) error {
	var res ClusterSettingsCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ClusterSettingsVMRecycleBin cluster settings VM recycle bin
//
// swagger:model ClusterSettingsVMRecycleBin
type ClusterSettingsVMRecycleBin struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// retain
	// Required: true
	Retain *float64 `json:"retain"`
}

// Validate validates this cluster settings VM recycle bin
func (m *ClusterSettingsVMRecycleBin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetain(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSettingsVMRecycleBin) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("vm_recycle_bin"+"."+"enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSettingsVMRecycleBin) validateRetain(formats strfmt.Registry) error {

	if err := validate.Required("vm_recycle_bin"+"."+"retain", "body", m.Retain); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster settings VM recycle bin based on context it is used
func (m *ClusterSettingsVMRecycleBin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSettingsVMRecycleBin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSettingsVMRecycleBin) UnmarshalBinary(b []byte) error {
	var res ClusterSettingsVMRecycleBin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
