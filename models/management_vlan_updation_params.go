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

// ManagementVlanUpdationParams management vlan updation params
//
// swagger:model ManagementVlanUpdationParams
type ManagementVlanUpdationParams struct {

	// extra ip
	ExtraIP []*ManagementVlanUpdationParamsExtraIPItems0 `json:"extra_ip"`

	// gateway ip
	GatewayIP string `json:"gateway_ip,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// subnetmask
	Subnetmask string `json:"subnetmask,omitempty"`

	// vlan id
	VlanID float64 `json:"vlan_id,omitempty"`
}

// Validate validates this management vlan updation params
func (m *ManagementVlanUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtraIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementVlanUpdationParams) validateExtraIP(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraIP) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtraIP); i++ {
		if swag.IsZero(m.ExtraIP[i]) { // not required
			continue
		}

		if m.ExtraIP[i] != nil {
			if err := m.ExtraIP[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extra_ip" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ManagementVlanUpdationParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this management vlan updation params based on the context it is used
func (m *ManagementVlanUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExtraIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementVlanUpdationParams) contextValidateExtraIP(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtraIP); i++ {

		if m.ExtraIP[i] != nil {
			if err := m.ExtraIP[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extra_ip" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementVlanUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementVlanUpdationParams) UnmarshalBinary(b []byte) error {
	var res ManagementVlanUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ManagementVlanUpdationParamsExtraIPItems0 management vlan updation params extra IP items0
//
// swagger:model ManagementVlanUpdationParamsExtraIPItems0
type ManagementVlanUpdationParamsExtraIPItems0 struct {

	// host id
	// Required: true
	HostID *string `json:"host_id"`

	// management ip
	// Required: true
	ManagementIP *string `json:"management_ip"`
}

// Validate validates this management vlan updation params extra IP items0
func (m *ManagementVlanUpdationParamsExtraIPItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementVlanUpdationParamsExtraIPItems0) validateHostID(formats strfmt.Registry) error {

	if err := validate.Required("host_id", "body", m.HostID); err != nil {
		return err
	}

	return nil
}

func (m *ManagementVlanUpdationParamsExtraIPItems0) validateManagementIP(formats strfmt.Registry) error {

	if err := validate.Required("management_ip", "body", m.ManagementIP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this management vlan updation params extra IP items0 based on context it is used
func (m *ManagementVlanUpdationParamsExtraIPItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManagementVlanUpdationParamsExtraIPItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementVlanUpdationParamsExtraIPItems0) UnmarshalBinary(b []byte) error {
	var res ManagementVlanUpdationParamsExtraIPItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
