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

// VsphereEsxiAccount vsphere esxi account
//
// swagger:model VsphereEsxiAccount
type VsphereEsxiAccount struct {

	// host
	// Required: true
	Host *VsphereEsxiAccountHost `json:"host"`

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
	Port *float64 `json:"port"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this vsphere esxi account
func (m *VsphereEsxiAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

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

func (m *VsphereEsxiAccount) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *VsphereEsxiAccount) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VsphereEsxiAccount) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *VsphereEsxiAccount) validateIsValid(formats strfmt.Registry) error {

	if err := validate.Required("is_valid", "body", m.IsValid); err != nil {
		return err
	}

	return nil
}

func (m *VsphereEsxiAccount) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VsphereEsxiAccount) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *VsphereEsxiAccount) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *VsphereEsxiAccount) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vsphere esxi account based on the context it is used
func (m *VsphereEsxiAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VsphereEsxiAccount) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VsphereEsxiAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsphereEsxiAccount) UnmarshalBinary(b []byte) error {
	var res VsphereEsxiAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VsphereEsxiAccountHost vsphere esxi account host
//
// swagger:model VsphereEsxiAccountHost
type VsphereEsxiAccountHost struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this vsphere esxi account host
func (m *VsphereEsxiAccountHost) Validate(formats strfmt.Registry) error {
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

func (m *VsphereEsxiAccountHost) validateID(formats strfmt.Registry) error {

	if err := validate.Required("host"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VsphereEsxiAccountHost) validateName(formats strfmt.Registry) error {

	if err := validate.Required("host"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this vsphere esxi account host based on context it is used
func (m *VsphereEsxiAccountHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VsphereEsxiAccountHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VsphereEsxiAccountHost) UnmarshalBinary(b []byte) error {
	var res VsphereEsxiAccountHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
