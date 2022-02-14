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

// NestedTemplateNic nested template nic
//
// swagger:model NestedTemplateNic
type NestedTemplateNic struct {

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// index
	// Required: true
	Index *int32 `json:"index"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// mac address
	MacAddress *string `json:"mac_address,omitempty"`

	// mirror
	Mirror *bool `json:"mirror,omitempty"`

	// model
	Model interface{} `json:"model,omitempty"`

	// vlan
	// Required: true
	Vlan *NestedFrozenVlan `json:"vlan"`
}

// Validate validates this nested template nic
func (m *NestedTemplateNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedTemplateNic) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *NestedTemplateNic) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("vlan", "body", m.Vlan); err != nil {
		return err
	}

	if m.Vlan != nil {
		if err := m.Vlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested template nic based on the context it is used
func (m *NestedTemplateNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedTemplateNic) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlan != nil {
		if err := m.Vlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedTemplateNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedTemplateNic) UnmarshalBinary(b []byte) error {
	var res NestedTemplateNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
