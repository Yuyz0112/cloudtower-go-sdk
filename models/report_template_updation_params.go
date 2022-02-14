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

// ReportTemplateUpdationParams report template updation params
//
// swagger:model ReportTemplateUpdationParams
type ReportTemplateUpdationParams struct {

	// data
	// Required: true
	Data *ReportTemplateUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *ReportTemplateWhereInput `json:"where"`
}

// Validate validates this report template updation params
func (m *ReportTemplateUpdationParams) Validate(formats strfmt.Registry) error {
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

func (m *ReportTemplateUpdationParams) validateData(formats strfmt.Registry) error {

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

func (m *ReportTemplateUpdationParams) validateWhere(formats strfmt.Registry) error {

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

// ContextValidate validate this report template updation params based on the context it is used
func (m *ReportTemplateUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ReportTemplateUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ReportTemplateUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ReportTemplateUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportTemplateUpdationParams) UnmarshalBinary(b []byte) error {
	var res ReportTemplateUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ReportTemplateUpdationParamsData report template updation params data
//
// swagger:model ReportTemplateUpdationParamsData
type ReportTemplateUpdationParamsData struct {

	// description
	Description string `json:"description,omitempty"`

	// execute plan
	ExecutePlan []*ExecutePlan `json:"execute_plan"`

	// name
	Name string `json:"name,omitempty"`

	// resource meta
	ResourceMeta []*ResourceMeta `json:"resource_meta"`
}

// Validate validates this report template updation params data
func (m *ReportTemplateUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportTemplateUpdationParamsData) validateExecutePlan(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecutePlan) { // not required
		return nil
	}

	for i := 0; i < len(m.ExecutePlan); i++ {
		if swag.IsZero(m.ExecutePlan[i]) { // not required
			continue
		}

		if m.ExecutePlan[i] != nil {
			if err := m.ExecutePlan[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "execute_plan" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportTemplateUpdationParamsData) validateResourceMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceMeta) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourceMeta); i++ {
		if swag.IsZero(m.ResourceMeta[i]) { // not required
			continue
		}

		if m.ResourceMeta[i] != nil {
			if err := m.ResourceMeta[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "resource_meta" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this report template updation params data based on the context it is used
func (m *ReportTemplateUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExecutePlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportTemplateUpdationParamsData) contextValidateExecutePlan(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExecutePlan); i++ {

		if m.ExecutePlan[i] != nil {
			if err := m.ExecutePlan[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "execute_plan" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportTemplateUpdationParamsData) contextValidateResourceMeta(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceMeta); i++ {

		if m.ResourceMeta[i] != nil {
			if err := m.ResourceMeta[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "resource_meta" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportTemplateUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportTemplateUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res ReportTemplateUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
