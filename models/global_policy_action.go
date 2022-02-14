// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// GlobalPolicyAction global policy action
//
// swagger:model GlobalPolicyAction
type GlobalPolicyAction string

func NewGlobalPolicyAction(value GlobalPolicyAction) *GlobalPolicyAction {
	v := value
	return &v
}

const (

	// GlobalPolicyActionALLOW captures enum value "ALLOW"
	GlobalPolicyActionALLOW GlobalPolicyAction = "ALLOW"

	// GlobalPolicyActionDROP captures enum value "DROP"
	GlobalPolicyActionDROP GlobalPolicyAction = "DROP"
)

// for schema
var globalPolicyActionEnum []interface{}

func init() {
	var res []GlobalPolicyAction
	if err := json.Unmarshal([]byte(`["ALLOW","DROP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		globalPolicyActionEnum = append(globalPolicyActionEnum, v)
	}
}

func (m GlobalPolicyAction) validateGlobalPolicyActionEnum(path, location string, value GlobalPolicyAction) error {
	if err := validate.EnumCase(path, location, value, globalPolicyActionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this global policy action
func (m GlobalPolicyAction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGlobalPolicyActionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this global policy action based on context it is used
func (m GlobalPolicyAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
