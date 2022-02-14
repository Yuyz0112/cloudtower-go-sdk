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

// VcenterAccount vcenter account
//
// swagger:model VcenterAccount
type VcenterAccount struct {

	// cluster
	Cluster interface{} `json:"cluster,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// ip
	// Required: true
	IP *string `json:"ip"`

	// is valid
	// Required: true
	IsValid *bool `json:"is_valid"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// password
	// Required: true
	Password *string `json:"password"`

	// port
	// Required: true
	Port *int32 `json:"port"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this vcenter account
func (m *VcenterAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsValid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcenterAccount) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VcenterAccount) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *VcenterAccount) validateIsValid(formats strfmt.Registry) error {

	if err := validate.Required("is_valid", "body", m.IsValid); err != nil {
		return err
	}

	return nil
}

func (m *VcenterAccount) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VcenterAccount) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *VcenterAccount) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *VcenterAccount) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vcenter account based on context it is used
func (m *VcenterAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VcenterAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcenterAccount) UnmarshalBinary(b []byte) error {
	var res VcenterAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
