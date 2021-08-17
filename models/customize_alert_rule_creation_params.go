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

// CustomizeAlertRuleCreationParams customize alert rule creation params
//
// swagger:model CustomizeAlertRuleCreationParams
type CustomizeAlertRuleCreationParams struct {

	// data
	// Required: true
	Data *CustomizeAlertRuleCreationParamsData `json:"data"`

	// where
	// Required: true
	Where *AlertRuleWhereInput `json:"where"`
}

// Validate validates this customize alert rule creation params
func (m *CustomizeAlertRuleCreationParams) Validate(formats strfmt.Registry) error {
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

func (m *CustomizeAlertRuleCreationParams) validateData(formats strfmt.Registry) error {

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

func (m *CustomizeAlertRuleCreationParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this customize alert rule creation params based on the context it is used
func (m *CustomizeAlertRuleCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *CustomizeAlertRuleCreationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CustomizeAlertRuleCreationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CustomizeAlertRuleCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomizeAlertRuleCreationParams) UnmarshalBinary(b []byte) error {
	var res CustomizeAlertRuleCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CustomizeAlertRuleCreationParamsData customize alert rule creation params data
//
// swagger:model CustomizeAlertRuleCreationParamsData
type CustomizeAlertRuleCreationParamsData struct {

	// disabled
	// Required: true
	Disabled *bool `json:"disabled"`

	// thresholds
	// Required: true
	Thresholds []*AlertRuleThresholds `json:"thresholds"`
}

// Validate validates this customize alert rule creation params data
func (m *CustomizeAlertRuleCreationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThresholds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomizeAlertRuleCreationParamsData) validateDisabled(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"disabled", "body", m.Disabled); err != nil {
		return err
	}

	return nil
}

func (m *CustomizeAlertRuleCreationParamsData) validateThresholds(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"thresholds", "body", m.Thresholds); err != nil {
		return err
	}

	for i := 0; i < len(m.Thresholds); i++ {
		if swag.IsZero(m.Thresholds[i]) { // not required
			continue
		}

		if m.Thresholds[i] != nil {
			if err := m.Thresholds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "thresholds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this customize alert rule creation params data based on the context it is used
func (m *CustomizeAlertRuleCreationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateThresholds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomizeAlertRuleCreationParamsData) contextValidateThresholds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Thresholds); i++ {

		if m.Thresholds[i] != nil {
			if err := m.Thresholds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "thresholds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomizeAlertRuleCreationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomizeAlertRuleCreationParamsData) UnmarshalBinary(b []byte) error {
	var res CustomizeAlertRuleCreationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
