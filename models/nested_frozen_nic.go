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

// NestedFrozenNic nested frozen nic
//
// swagger:model NestedFrozenNic
type NestedFrozenNic struct {

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// gateway
	// Required: true
	Gateway *string `json:"gateway"`

	// index
	// Required: true
	Index *int32 `json:"index"`

	// ip address
	// Required: true
	IPAddress *string `json:"ip_address"`

	// mac address
	// Required: true
	MacAddress *string `json:"mac_address"`

	// mirror
	Mirror *bool `json:"mirror,omitempty"`

	// model
	Model *VMNicModel `json:"model,omitempty"`

	// subnet mask
	// Required: true
	SubnetMask *string `json:"subnet_mask"`

	// vlan
	// Required: true
	Vlan *NestedFrozenVlan `json:"vlan"`
}

// Validate validates this nested frozen nic
func (m *NestedFrozenNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetMask(formats); err != nil {
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

func (m *NestedFrozenNic) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenNic) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenNic) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ip_address", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenNic) validateMacAddress(formats strfmt.Registry) error {

	if err := validate.Required("mac_address", "body", m.MacAddress); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenNic) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(m.Model) { // not required
		return nil
	}

	if m.Model != nil {
		if err := m.Model.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFrozenNic) validateSubnetMask(formats strfmt.Registry) error {

	if err := validate.Required("subnet_mask", "body", m.SubnetMask); err != nil {
		return err
	}

	return nil
}

func (m *NestedFrozenNic) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("vlan", "body", m.Vlan); err != nil {
		return err
	}

	if m.Vlan != nil {
		if err := m.Vlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested frozen nic based on the context it is used
func (m *NestedFrozenNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFrozenNic) contextValidateModel(ctx context.Context, formats strfmt.Registry) error {

	if m.Model != nil {
		if err := m.Model.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFrozenNic) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Vlan != nil {
		if err := m.Vlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedFrozenNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedFrozenNic) UnmarshalBinary(b []byte) error {
	var res NestedFrozenNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
