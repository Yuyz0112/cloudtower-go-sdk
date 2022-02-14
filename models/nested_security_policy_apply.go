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

// NestedSecurityPolicyApply nested security policy apply
//
// swagger:model NestedSecurityPolicyApply
type NestedSecurityPolicyApply struct {

	// communicable
	// Required: true
	Communicable *bool `json:"communicable"`

	// selector
	// Required: true
	Selector []*NestedLabel `json:"selector"`

	// selector ids
	// Required: true
	SelectorIds []string `json:"selector_ids"`
}

// Validate validates this nested security policy apply
func (m *NestedSecurityPolicyApply) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommunicable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectorIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSecurityPolicyApply) validateCommunicable(formats strfmt.Registry) error {

	if err := validate.Required("communicable", "body", m.Communicable); err != nil {
		return err
	}

	return nil
}

func (m *NestedSecurityPolicyApply) validateSelector(formats strfmt.Registry) error {

	if err := validate.Required("selector", "body", m.Selector); err != nil {
		return err
	}

	for i := 0; i < len(m.Selector); i++ {
		if swag.IsZero(m.Selector[i]) { // not required
			continue
		}

		if m.Selector[i] != nil {
			if err := m.Selector[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selector" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NestedSecurityPolicyApply) validateSelectorIds(formats strfmt.Registry) error {

	if err := validate.Required("selector_ids", "body", m.SelectorIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested security policy apply based on the context it is used
func (m *NestedSecurityPolicyApply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedSecurityPolicyApply) contextValidateSelector(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Selector); i++ {

		if m.Selector[i] != nil {
			if err := m.Selector[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selector" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedSecurityPolicyApply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedSecurityPolicyApply) UnmarshalBinary(b []byte) error {
	var res NestedSecurityPolicyApply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
