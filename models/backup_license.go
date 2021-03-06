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

// BackupLicense backup license
//
// swagger:model BackupLicense
type BackupLicense struct {

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// expire date
	// Required: true
	ExpireDate *string `json:"expire_date"`

	// id
	// Required: true
	ID *string `json:"id"`

	// license serial
	// Required: true
	LicenseSerial *string `json:"license_serial"`

	// max capacity
	// Required: true
	MaxCapacity *float64 `json:"max_capacity"`

	// sign date
	// Required: true
	SignDate *string `json:"sign_date"`

	// software edition
	// Required: true
	SoftwareEdition *SoftwareEdition `json:"software_edition"`

	// type
	// Required: true
	Type *LicenseType `json:"type"`
}

// Validate validates this backup license
func (m *BackupLicense) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpireDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEdition(formats); err != nil {
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

func (m *BackupLicense) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicense) validateExpireDate(formats strfmt.Registry) error {

	if err := validate.Required("expire_date", "body", m.ExpireDate); err != nil {
		return err
	}

	return nil
}

func (m *BackupLicense) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BackupLicense) validateLicenseSerial(formats strfmt.Registry) error {

	if err := validate.Required("license_serial", "body", m.LicenseSerial); err != nil {
		return err
	}

	return nil
}

func (m *BackupLicense) validateMaxCapacity(formats strfmt.Registry) error {

	if err := validate.Required("max_capacity", "body", m.MaxCapacity); err != nil {
		return err
	}

	return nil
}

func (m *BackupLicense) validateSignDate(formats strfmt.Registry) error {

	if err := validate.Required("sign_date", "body", m.SignDate); err != nil {
		return err
	}

	return nil
}

func (m *BackupLicense) validateSoftwareEdition(formats strfmt.Registry) error {

	if err := validate.Required("software_edition", "body", m.SoftwareEdition); err != nil {
		return err
	}

	if err := validate.Required("software_edition", "body", m.SoftwareEdition); err != nil {
		return err
	}

	if m.SoftwareEdition != nil {
		if err := m.SoftwareEdition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicense) validateType(formats strfmt.Registry) error {

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

// ContextValidate validate this backup license based on the context it is used
func (m *BackupLicense) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEdition(ctx, formats); err != nil {
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

func (m *BackupLicense) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicense) contextValidateSoftwareEdition(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareEdition != nil {
		if err := m.SoftwareEdition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicense) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BackupLicense) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupLicense) UnmarshalBinary(b []byte) error {
	var res BackupLicense
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
