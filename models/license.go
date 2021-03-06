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

// License license
//
// swagger:model License
type License struct {

	// expire date
	// Required: true
	ExpireDate *string `json:"expire_date"`

	// id
	// Required: true
	ID *string `json:"id"`

	// license serial
	// Required: true
	LicenseSerial *string `json:"license_serial"`

	// maintenance end date
	MaintenanceEndDate *string `json:"maintenance_end_date,omitempty"`

	// maintenance start date
	MaintenanceStartDate *string `json:"maintenance_start_date,omitempty"`

	// max chunk num
	// Required: true
	MaxChunkNum *int32 `json:"max_chunk_num"`

	// max cluster num
	// Required: true
	MaxClusterNum *int32 `json:"max_cluster_num"`

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

// Validate validates this license
func (m *License) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpireDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxChunkNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxClusterNum(formats); err != nil {
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

func (m *License) validateExpireDate(formats strfmt.Registry) error {

	if err := validate.Required("expire_date", "body", m.ExpireDate); err != nil {
		return err
	}

	return nil
}

func (m *License) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *License) validateLicenseSerial(formats strfmt.Registry) error {

	if err := validate.Required("license_serial", "body", m.LicenseSerial); err != nil {
		return err
	}

	return nil
}

func (m *License) validateMaxChunkNum(formats strfmt.Registry) error {

	if err := validate.Required("max_chunk_num", "body", m.MaxChunkNum); err != nil {
		return err
	}

	return nil
}

func (m *License) validateMaxClusterNum(formats strfmt.Registry) error {

	if err := validate.Required("max_cluster_num", "body", m.MaxClusterNum); err != nil {
		return err
	}

	return nil
}

func (m *License) validateSignDate(formats strfmt.Registry) error {

	if err := validate.Required("sign_date", "body", m.SignDate); err != nil {
		return err
	}

	return nil
}

func (m *License) validateSoftwareEdition(formats strfmt.Registry) error {

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

func (m *License) validateType(formats strfmt.Registry) error {

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

// ContextValidate validate this license based on the context it is used
func (m *License) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *License) contextValidateSoftwareEdition(ctx context.Context, formats strfmt.Registry) error {

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

func (m *License) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *License) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *License) UnmarshalBinary(b []byte) error {
	var res License
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
