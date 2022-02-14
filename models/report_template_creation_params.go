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

// ReportTemplateCreationParams report template creation params
//
// swagger:model ReportTemplateCreationParams
type ReportTemplateCreationParams struct {

	// description
	Description string `json:"description,omitempty"`

	// execute plan
	// Required: true
	ExecutePlan []*ExecutePlan `json:"execute_plan"`

	// name
	// Required: true
	Name *string `json:"name"`

	// resource meta
	// Required: true
	ResourceMeta []*ResourceMeta `json:"resource_meta"`
}

// Validate validates this report template creation params
func (m *ReportTemplateCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *ReportTemplateCreationParams) validateExecutePlan(formats strfmt.Registry) error {

	if err := validate.Required("execute_plan", "body", m.ExecutePlan); err != nil {
		return err
	}

	for i := 0; i < len(m.ExecutePlan); i++ {
		if swag.IsZero(m.ExecutePlan[i]) { // not required
			continue
		}

		if m.ExecutePlan[i] != nil {
			if err := m.ExecutePlan[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("execute_plan" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("execute_plan" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportTemplateCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ReportTemplateCreationParams) validateResourceMeta(formats strfmt.Registry) error {

	if err := validate.Required("resource_meta", "body", m.ResourceMeta); err != nil {
		return err
	}

	for i := 0; i < len(m.ResourceMeta); i++ {
		if swag.IsZero(m.ResourceMeta[i]) { // not required
			continue
		}

		if m.ResourceMeta[i] != nil {
			if err := m.ResourceMeta[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resource_meta" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resource_meta" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this report template creation params based on the context it is used
func (m *ReportTemplateCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *ReportTemplateCreationParams) contextValidateExecutePlan(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExecutePlan); i++ {

		if m.ExecutePlan[i] != nil {
			if err := m.ExecutePlan[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("execute_plan" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("execute_plan" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportTemplateCreationParams) contextValidateResourceMeta(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourceMeta); i++ {

		if m.ResourceMeta[i] != nil {
			if err := m.ResourceMeta[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resource_meta" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("resource_meta" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportTemplateCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportTemplateCreationParams) UnmarshalBinary(b []byte) error {
	var res ReportTemplateCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
