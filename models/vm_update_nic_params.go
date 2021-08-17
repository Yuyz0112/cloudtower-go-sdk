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

// VMUpdateNicParams Vm update nic params
//
// swagger:model VmUpdateNicParams
type VMUpdateNicParams struct {

	// data
	// Required: true
	Data *VMUpdateNicParamsData `json:"data"`

	// where
	// Required: true
	Where *VMWhereInput `json:"where"`
}

// Validate validates this Vm update nic params
func (m *VMUpdateNicParams) Validate(formats strfmt.Registry) error {
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

func (m *VMUpdateNicParams) validateData(formats strfmt.Registry) error {

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

func (m *VMUpdateNicParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this Vm update nic params based on the context it is used
func (m *VMUpdateNicParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VMUpdateNicParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMUpdateNicParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMUpdateNicParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateNicParams) UnmarshalBinary(b []byte) error {
	var res VMUpdateNicParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMUpdateNicParamsData VM update nic params data
//
// swagger:model VMUpdateNicParamsData
type VMUpdateNicParamsData struct {

	// connect vlan id
	// Required: true
	ConnectVlanID *string `json:"connect_vlan_id"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// gateway
	Gateway string `json:"gateway,omitempty"`

	// ip address
	IPAddress string `json:"ip_address,omitempty"`

	// mac address
	MacAddress string `json:"mac_address,omitempty"`

	// mirror
	Mirror bool `json:"mirror,omitempty"`

	// model
	Model VMNicModel `json:"model,omitempty"`

	// nic id
	NicID string `json:"nic_id,omitempty"`

	// nic index
	// Required: true
	NicIndex *float64 `json:"nic_index"`

	// subnet mask
	SubnetMask string `json:"subnet_mask,omitempty"`
}

// Validate validates this VM update nic params data
func (m *VMUpdateNicParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectVlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMUpdateNicParamsData) validateConnectVlanID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"connect_vlan_id", "body", m.ConnectVlanID); err != nil {
		return err
	}

	return nil
}

func (m *VMUpdateNicParamsData) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(m.Model) { // not required
		return nil
	}

	if err := m.Model.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "model")
		}
		return err
	}

	return nil
}

func (m *VMUpdateNicParamsData) validateNicIndex(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"nic_index", "body", m.NicIndex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this VM update nic params data based on the context it is used
func (m *VMUpdateNicParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMUpdateNicParamsData) contextValidateModel(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Model.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data" + "." + "model")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMUpdateNicParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMUpdateNicParamsData) UnmarshalBinary(b []byte) error {
	var res VMUpdateNicParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
