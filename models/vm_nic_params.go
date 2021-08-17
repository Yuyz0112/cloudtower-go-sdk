// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMNicParams Vm nic params
//
// swagger:model VmNicParams
type VMNicParams []*VMNicParamsItems0

// Validate validates this Vm nic params
func (m VMNicParams) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this Vm nic params based on the context it is used
func (m VMNicParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// VMNicParamsItems0 VM nic params items0
//
// swagger:model VMNicParamsItems0
type VMNicParamsItems0 struct {

	// connect vlan id
	// Required: true
	ConnectVlanID *string `json:"connect_vlan_id"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// ip address
	IPAddress string `json:"ip_address,omitempty"`

	// local id
	LocalID string `json:"local_id,omitempty"`

	// mac address
	MacAddress string `json:"mac_address,omitempty"`

	// mirror
	Mirror bool `json:"mirror,omitempty"`

	// model
	Model VMNicModel `json:"model,omitempty"`

	// nic id
	NicID string `json:"nic_id,omitempty"`

	// subnet mask
	SubnetMask string `json:"subnet_mask,omitempty"`
}

// Validate validates this VM nic params items0
func (m *VMNicParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectVlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMNicParamsItems0) validateConnectVlanID(formats strfmt.Registry) error {

	if err := validate.Required("connect_vlan_id", "body", m.ConnectVlanID); err != nil {
		return err
	}

	return nil
}

func (m *VMNicParamsItems0) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(m.Model) { // not required
		return nil
	}

	if err := m.Model.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("model")
		}
		return err
	}

	return nil
}

// ContextValidate validate this VM nic params items0 based on the context it is used
func (m *VMNicParamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMNicParamsItems0) contextValidateModel(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Model.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("model")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMNicParamsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNicParamsItems0) UnmarshalBinary(b []byte) error {
	var res VMNicParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
