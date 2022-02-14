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

// IscsiConnection iscsi connection
//
// swagger:model IscsiConnection
type IscsiConnection struct {

	// client port
	// Required: true
	ClientPort *int32 `json:"client_port"`

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// host
	// Required: true
	Host *NestedHost `json:"host"`

	// id
	// Required: true
	ID *string `json:"id"`

	// initiator ip
	// Required: true
	InitiatorIP *string `json:"initiator_ip"`

	// iscsi target
	IscsiTarget struct {
		NestedIscsiTarget
	} `json:"iscsi_target,omitempty"`

	// nvmf subsystem
	NvmfSubsystem struct {
		NestedNvmfSubsystem
	} `json:"nvmf_subsystem,omitempty"`

	// type
	// Required: true
	Type *StoreConnectionType `json:"type"`
}

// Validate validates this iscsi connection
func (m *IscsiConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnection) validateClientPort(formats strfmt.Registry) error {

	if err := validate.Required("client_port", "body", m.ClientPort); err != nil {
		return err
	}

	return nil
}

func (m *IscsiConnection) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnection) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnection) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiConnection) validateInitiatorIP(formats strfmt.Registry) error {

	if err := validate.Required("initiator_ip", "body", m.InitiatorIP); err != nil {
		return err
	}

	return nil
}

func (m *IscsiConnection) validateIscsiTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiTarget) { // not required
		return nil
	}

	return nil
}

func (m *IscsiConnection) validateNvmfSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.NvmfSubsystem) { // not required
		return nil
	}

	return nil
}

func (m *IscsiConnection) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this iscsi connection based on the context it is used
func (m *IscsiConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiConnection) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnection) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

func (m *IscsiConnection) contextValidateIscsiTarget(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *IscsiConnection) contextValidateNvmfSubsystem(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *IscsiConnection) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiConnection) UnmarshalBinary(b []byte) error {
	var res IscsiConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
