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

// SnmpTrapReceiverCreationParams snmp trap receiver creation params
//
// swagger:model SnmpTrapReceiverCreationParams
type SnmpTrapReceiverCreationParams struct {

	// auth pass phrase
	AuthPassPhrase string `json:"auth_pass_phrase,omitempty"`

	// auth protocol
	AuthProtocol SnmpAuthProtocol `json:"auth_protocol,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// community
	Community string `json:"community,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// engine id
	EngineID string `json:"engine_id,omitempty"`

	// host
	// Required: true
	Host *string `json:"host"`

	// inform
	Inform bool `json:"inform,omitempty"`

	// language code
	// Required: true
	LanguageCode *SnmpLanguageCode `json:"language_code"`

	// name
	// Required: true
	Name *string `json:"name"`

	// port
	// Required: true
	Port *int32 `json:"port"`

	// privacy pass phrase
	PrivacyPassPhrase string `json:"privacy_pass_phrase,omitempty"`

	// privacy protocol
	PrivacyProtocol SnmpPrivacyProtocol `json:"privacy_protocol,omitempty"`

	// protocol
	// Required: true
	Protocol *SnmpProtocol `json:"protocol"`

	// username
	Username string `json:"username,omitempty"`

	// version
	// Required: true
	Version *SnmpVersion `json:"version"`
}

// Validate validates this snmp trap receiver creation params
func (m *SnmpTrapReceiverCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacyProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpTrapReceiverCreationParams) validateAuthProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthProtocol) { // not required
		return nil
	}

	if err := m.AuthProtocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("auth_protocol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("auth_protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) validateLanguageCode(formats strfmt.Registry) error {

	if err := validate.Required("language_code", "body", m.LanguageCode); err != nil {
		return err
	}

	if err := validate.Required("language_code", "body", m.LanguageCode); err != nil {
		return err
	}

	if m.LanguageCode != nil {
		if err := m.LanguageCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) validatePrivacyProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivacyProtocol) { // not required
		return nil
	}

	if err := m.PrivacyProtocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("privacy_protocol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("privacy_protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	if m.Protocol != nil {
		if err := m.Protocol.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snmp trap receiver creation params based on the context it is used
func (m *SnmpTrapReceiverCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivacyProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpTrapReceiverCreationParams) contextValidateAuthProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AuthProtocol.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("auth_protocol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("auth_protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) contextValidateLanguageCode(ctx context.Context, formats strfmt.Registry) error {

	if m.LanguageCode != nil {
		if err := m.LanguageCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) contextValidatePrivacyProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PrivacyProtocol.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("privacy_protocol")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("privacy_protocol")
		}
		return err
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if m.Protocol != nil {
		if err := m.Protocol.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("protocol")
			}
			return err
		}
	}

	return nil
}

func (m *SnmpTrapReceiverCreationParams) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpTrapReceiverCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpTrapReceiverCreationParams) UnmarshalBinary(b []byte) error {
	var res SnmpTrapReceiverCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
