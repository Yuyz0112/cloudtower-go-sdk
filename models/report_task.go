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

// ReportTask report task
//
// swagger:model ReportTask
type ReportTask struct {

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// id
	// Required: true
	ID *string `json:"id"`

	// internal
	// Required: true
	Internal *bool `json:"internal"`

	// name
	// Required: true
	Name *string `json:"name"`

	// plan id
	PlanID *string `json:"plan_id,omitempty"`

	// status
	// Required: true
	Status *TaskStatus `json:"status"`

	// template
	// Required: true
	Template *ReportTaskTemplate `json:"template"`
}

// Validate validates this report task
func (m *ReportTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportTask) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ReportTask) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ReportTask) validateInternal(formats strfmt.Registry) error {

	if err := validate.Required("internal", "body", m.Internal); err != nil {
		return err
	}

	return nil
}

func (m *ReportTask) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ReportTask) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *ReportTask) validateTemplate(formats strfmt.Registry) error {

	if err := validate.Required("template", "body", m.Template); err != nil {
		return err
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this report task based on the context it is used
func (m *ReportTask) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportTask) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *ReportTask) contextValidateTemplate(ctx context.Context, formats strfmt.Registry) error {

	if m.Template != nil {
		if err := m.Template.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportTask) UnmarshalBinary(b []byte) error {
	var res ReportTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ReportTaskTemplate report task template
//
// swagger:model ReportTaskTemplate
type ReportTaskTemplate struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this report task template
func (m *ReportTaskTemplate) Validate(formats strfmt.Registry) error {
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

func (m *ReportTaskTemplate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("template"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ReportTaskTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("template"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this report task template based on context it is used
func (m *ReportTaskTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportTaskTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportTaskTemplate) UnmarshalBinary(b []byte) error {
	var res ReportTaskTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
