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

// NestedEverouteControllerStatus nested everoute controller status
//
// swagger:model NestedEverouteControllerStatus
type NestedEverouteControllerStatus struct {

	// ip addr
	// Required: true
	IPAddr *string `json:"ipAddr"`

	// is health
	// Required: true
	IsHealth *bool `json:"isHealth"`

	// message
	// Required: true
	Message *string `json:"message"`

	// metrics
	Metrics interface{} `json:"metrics,omitempty"`

	// phase
	Phase interface{} `json:"phase,omitempty"`

	// reason
	// Required: true
	Reason *string `json:"reason"`

	// vm
	VM interface{} `json:"vm,omitempty"`

	// vm ID
	// Required: true
	VMID *string `json:"vmID"`
}

// Validate validates this nested everoute controller status
func (m *NestedEverouteControllerStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedEverouteControllerStatus) validateIPAddr(formats strfmt.Registry) error {

	if err := validate.Required("ipAddr", "body", m.IPAddr); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerStatus) validateIsHealth(formats strfmt.Registry) error {

	if err := validate.Required("isHealth", "body", m.IsHealth); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerStatus) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerStatus) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *NestedEverouteControllerStatus) validateVMID(formats strfmt.Registry) error {

	if err := validate.Required("vmID", "body", m.VMID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nested everoute controller status based on context it is used
func (m *NestedEverouteControllerStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NestedEverouteControllerStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedEverouteControllerStatus) UnmarshalBinary(b []byte) error {
	var res NestedEverouteControllerStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
