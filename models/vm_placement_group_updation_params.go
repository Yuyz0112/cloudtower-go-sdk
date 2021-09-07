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

// VMPlacementGroupUpdationParams Vm placement group updation params
//
// swagger:model VmPlacementGroupUpdationParams
type VMPlacementGroupUpdationParams struct {

	// data
	// Required: true
	Data *VMPlacementGroupUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *VMPlacementGroupWhereInput `json:"where"`
}

// Validate validates this Vm placement group updation params
func (m *VMPlacementGroupUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMPlacementGroupUpdationParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *VMPlacementGroupUpdationParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm placement group updation params based on the context it is used
func (m *VMPlacementGroupUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMPlacementGroupUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *VMPlacementGroupUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMPlacementGroupUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMPlacementGroupUpdationParams) UnmarshalBinary(b []byte) error {
	var res VMPlacementGroupUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMPlacementGroupUpdationParamsData VM placement group updation params data
//
// swagger:model VMPlacementGroupUpdationParamsData
type VMPlacementGroupUpdationParamsData struct {

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// must hosts
	MustHosts *HostWhereInput `json:"must_hosts,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// prefer hosts
	PreferHosts *HostWhereInput `json:"prefer_hosts,omitempty"`

	// vm host must enabled
	VMHostMustEnabled bool `json:"vm_host_must_enabled,omitempty"`

	// vm host must policy
	VMHostMustPolicy bool `json:"vm_host_must_policy,omitempty"`

	// vm host prefer enabled
	VMHostPreferEnabled bool `json:"vm_host_prefer_enabled,omitempty"`

	// vm host prefer policy
	VMHostPreferPolicy bool `json:"vm_host_prefer_policy,omitempty"`

	// vm vm policy
	VMVMPolicy VMVMPolicy `json:"vm_vm_policy,omitempty"`

	// vm vm policy enabled
	VMVMPolicyEnabled bool `json:"vm_vm_policy_enabled,omitempty"`

	// vms
	Vms *VMWhereInput `json:"vms,omitempty"`
}

// Validate validates this VM placement group updation params data
func (m *VMPlacementGroupUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMustHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreferHosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVMPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMPlacementGroupUpdationParamsData) validateMustHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.MustHosts) { // not required
		return nil
	}

	if m.MustHosts != nil {
		if err := m.MustHosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "must_hosts")
			}
			return err
		}
	}

	return nil
}

func (m *VMPlacementGroupUpdationParamsData) validatePreferHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.PreferHosts) { // not required
		return nil
	}

	if m.PreferHosts != nil {
		if err := m.PreferHosts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "prefer_hosts")
			}
			return err
		}
	}

	return nil
}

func (m *VMPlacementGroupUpdationParamsData) validateVMVMPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.VMVMPolicy) { // not required
		return nil
	}

	if err := m.VMVMPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "vm_vm_policy")
		}
		return err
	}

	return nil
}

func (m *VMPlacementGroupUpdationParamsData) validateVms(formats strfmt.Registry) error {
	if swag.IsZero(m.Vms) { // not required
		return nil
	}

	if m.Vms != nil {
		if err := m.Vms.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vms")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this VM placement group updation params data based on the context it is used
func (m *VMPlacementGroupUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMustHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreferHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMVMPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMPlacementGroupUpdationParamsData) contextValidateMustHosts(ctx context.Context, formats strfmt.Registry) error {

	if m.MustHosts != nil {
		if err := m.MustHosts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "must_hosts")
			}
			return err
		}
	}

	return nil
}

func (m *VMPlacementGroupUpdationParamsData) contextValidatePreferHosts(ctx context.Context, formats strfmt.Registry) error {

	if m.PreferHosts != nil {
		if err := m.PreferHosts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "prefer_hosts")
			}
			return err
		}
	}

	return nil
}

func (m *VMPlacementGroupUpdationParamsData) contextValidateVMVMPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VMVMPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "vm_vm_policy")
		}
		return err
	}

	return nil
}

func (m *VMPlacementGroupUpdationParamsData) contextValidateVms(ctx context.Context, formats strfmt.Registry) error {

	if m.Vms != nil {
		if err := m.Vms.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "vms")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMPlacementGroupUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMPlacementGroupUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res VMPlacementGroupUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
