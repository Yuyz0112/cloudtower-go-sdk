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

// VMNic Vm nic
//
// swagger:model VmNic
type VMNic struct {

	// enabled
	Enabled *bool `json:"enabled,omitempty"`

	// gateway
	Gateway *string `json:"gateway,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// interface id
	InterfaceID *string `json:"interface_id,omitempty"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// mac address
	MacAddress *string `json:"mac_address,omitempty"`

	// mirror
	Mirror *bool `json:"mirror,omitempty"`

	// model
	Model struct {
		VMNicModel
	} `json:"model,omitempty"`

	// nic
	Nic struct {
		NestedNic
	} `json:"nic,omitempty"`

	// order
	Order *int32 `json:"order,omitempty"`

	// subnet mask
	SubnetMask *string `json:"subnet_mask,omitempty"`

	// vlan
	Vlan struct {
		NestedVlan
	} `json:"vlan,omitempty"`

	// vm
	// Required: true
	VM *NestedVM `json:"vm"`
}

// Validate validates this Vm nic
func (m *VMNic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMNic) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMNic) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VMNic) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(m.Model) { // not required
		return nil
	}

	return nil
}

func (m *VMNic) validateNic(formats strfmt.Registry) error {
	if swag.IsZero(m.Nic) { // not required
		return nil
	}

	return nil
}

func (m *VMNic) validateVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.Vlan) { // not required
		return nil
	}

	return nil
}

func (m *VMNic) validateVM(formats strfmt.Registry) error {

	if err := validate.Required("vm", "body", m.VM); err != nil {
		return err
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm nic based on the context it is used
func (m *VMNic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMNic) contextValidateModel(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *VMNic) contextValidateNic(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *VMNic) contextValidateVlan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *VMNic) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMNic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNic) UnmarshalBinary(b []byte) error {
	var res VMNic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
