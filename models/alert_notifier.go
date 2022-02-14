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

// AlertNotifier alert notifier
//
// swagger:model AlertNotifier
type AlertNotifier struct {

	// disabled
	// Required: true
	Disabled *bool `json:"disabled"`

	// email from
	EmailFrom *string `json:"email_from,omitempty"`

	// email tos
	// Required: true
	EmailTos []string `json:"email_tos"`

	// id
	// Required: true
	ID *string `json:"id"`

	// language code
	LanguageCode interface{} `json:"language_code,omitempty"`

	// notice severities
	// Required: true
	NoticeSeverities []string `json:"notice_severities"`

	// security mode
	SecurityMode interface{} `json:"security_mode,omitempty"`

	// smtp server host
	SMTPServerHost *string `json:"smtp_server_host,omitempty"`

	// smtp server port
	SMTPServerPort *int32 `json:"smtp_server_port,omitempty"`

	// username
	Username *string `json:"username,omitempty"`
}

// Validate validates this alert notifier
func (m *AlertNotifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailTos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoticeSeverities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertNotifier) validateDisabled(formats strfmt.Registry) error {

	if err := validate.Required("disabled", "body", m.Disabled); err != nil {
		return err
	}

	return nil
}

func (m *AlertNotifier) validateEmailTos(formats strfmt.Registry) error {

	if err := validate.Required("email_tos", "body", m.EmailTos); err != nil {
		return err
	}

	return nil
}

func (m *AlertNotifier) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AlertNotifier) validateNoticeSeverities(formats strfmt.Registry) error {

	if err := validate.Required("notice_severities", "body", m.NoticeSeverities); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this alert notifier based on context it is used
func (m *AlertNotifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertNotifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertNotifier) UnmarshalBinary(b []byte) error {
	var res AlertNotifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
