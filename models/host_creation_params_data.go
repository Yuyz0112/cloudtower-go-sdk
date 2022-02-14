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

// HostCreationParamsData host creation params data
//
// swagger:model HostCreationParamsData
type HostCreationParamsData struct {

	// disks
	// Required: true
	Disks []*HostBatchCreateDiskInput `json:"disks"`

	// host ip
	// Required: true
	HostIP *string `json:"host_ip"`

	// host uuid
	// Required: true
	HostUUID *string `json:"host_uuid"`

	// hostname
	// Required: true
	Hostname *string `json:"hostname"`

	// ifaces
	// Required: true
	Ifaces []*HostBatchCreateIfaceInput `json:"ifaces"`

	// ipmi
	Ipmi *HostBatchCreateIpmiInput `json:"ipmi,omitempty"`

	// platform ip
	PlatformIP string `json:"platform_ip,omitempty"`

	// platform password
	PlatformPassword string `json:"platform_password,omitempty"`

	// platform username
	PlatformUsername string `json:"platform_username,omitempty"`
}

// Validate validates this host creation params data
func (m *HostCreationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpmi(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostCreationParamsData) validateDisks(formats strfmt.Registry) error {

	if err := validate.Required("disks", "body", m.Disks); err != nil {
		return err
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HostCreationParamsData) validateHostIP(formats strfmt.Registry) error {

	if err := validate.Required("host_ip", "body", m.HostIP); err != nil {
		return err
	}

	return nil
}

func (m *HostCreationParamsData) validateHostUUID(formats strfmt.Registry) error {

	if err := validate.Required("host_uuid", "body", m.HostUUID); err != nil {
		return err
	}

	return nil
}

func (m *HostCreationParamsData) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *HostCreationParamsData) validateIfaces(formats strfmt.Registry) error {

	if err := validate.Required("ifaces", "body", m.Ifaces); err != nil {
		return err
	}

	for i := 0; i < len(m.Ifaces); i++ {
		if swag.IsZero(m.Ifaces[i]) { // not required
			continue
		}

		if m.Ifaces[i] != nil {
			if err := m.Ifaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ifaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ifaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HostCreationParamsData) validateIpmi(formats strfmt.Registry) error {
	if swag.IsZero(m.Ipmi) { // not required
		return nil
	}

	if m.Ipmi != nil {
		if err := m.Ipmi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipmi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipmi")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this host creation params data based on the context it is used
func (m *HostCreationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpmi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostCreationParamsData) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HostCreationParamsData) contextValidateIfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ifaces); i++ {

		if m.Ifaces[i] != nil {
			if err := m.Ifaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ifaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ifaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HostCreationParamsData) contextValidateIpmi(ctx context.Context, formats strfmt.Registry) error {

	if m.Ipmi != nil {
		if err := m.Ipmi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipmi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipmi")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostCreationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostCreationParamsData) UnmarshalBinary(b []byte) error {
	var res HostCreationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
