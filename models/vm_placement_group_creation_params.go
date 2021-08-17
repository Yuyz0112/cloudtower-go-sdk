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

// VMPlacementGroupCreationParams Vm placement group creation params
//
// swagger:model VmPlacementGroupCreationParams
type VMPlacementGroupCreationParams struct {

	// associated vm ids
	AssociatedVMIds []string `json:"associated_vm_ids"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// is host must policy place
	IsHostMustPolicyPlace bool `json:"is_host_must_policy_place,omitempty"`

	// is host prefer policy place
	IsHostPreferPolicyPlace bool `json:"is_host_prefer_policy_place,omitempty"`

	// is select host must policy
	// Required: true
	IsSelectHostMustPolicy *bool `json:"is_select_host_must_policy"`

	// is select host prefer policy
	// Required: true
	IsSelectHostPreferPolicy *bool `json:"is_select_host_prefer_policy"`

	// is select vm policy
	// Required: true
	IsSelectVMPolicy *bool `json:"is_select_vm_policy"`

	// must policy host ids
	MustPolicyHostIds []string `json:"must_policy_host_ids"`

	// name
	// Required: true
	Name *string `json:"name"`

	// prefer policy host ids
	PreferPolicyHostIds []string `json:"prefer_policy_host_ids"`

	// vm policy
	VMPolicy VMVMPolicy `json:"vm_policy,omitempty"`
}

// Validate validates this Vm placement group creation params
func (m *VMPlacementGroupCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSelectHostMustPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSelectHostPreferPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSelectVMPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMPlacementGroupCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VMPlacementGroupCreationParams) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *VMPlacementGroupCreationParams) validateIsSelectHostMustPolicy(formats strfmt.Registry) error {

	if err := validate.Required("is_select_host_must_policy", "body", m.IsSelectHostMustPolicy); err != nil {
		return err
	}

	return nil
}

func (m *VMPlacementGroupCreationParams) validateIsSelectHostPreferPolicy(formats strfmt.Registry) error {

	if err := validate.Required("is_select_host_prefer_policy", "body", m.IsSelectHostPreferPolicy); err != nil {
		return err
	}

	return nil
}

func (m *VMPlacementGroupCreationParams) validateIsSelectVMPolicy(formats strfmt.Registry) error {

	if err := validate.Required("is_select_vm_policy", "body", m.IsSelectVMPolicy); err != nil {
		return err
	}

	return nil
}

func (m *VMPlacementGroupCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMPlacementGroupCreationParams) validateVMPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.VMPolicy) { // not required
		return nil
	}

	if err := m.VMPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vm_policy")
		}
		return err
	}

	return nil
}

// ContextValidate validate this Vm placement group creation params based on the context it is used
func (m *VMPlacementGroupCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVMPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMPlacementGroupCreationParams) contextValidateVMPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.VMPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("vm_policy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMPlacementGroupCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMPlacementGroupCreationParams) UnmarshalBinary(b []byte) error {
	var res VMPlacementGroupCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
