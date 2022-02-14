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

// VMToolsStatus Vm tools status
//
// swagger:model VmToolsStatus
type VMToolsStatus string

func NewVMToolsStatus(value VMToolsStatus) *VMToolsStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VMToolsStatus.
func (m VMToolsStatus) Pointer() *VMToolsStatus {
	return &m
}

const (

	// VMToolsStatusNOTINSTALLED captures enum value "NOT_INSTALLED"
	VMToolsStatusNOTINSTALLED VMToolsStatus = "NOT_INSTALLED"

	// VMToolsStatusNOTRUNNING captures enum value "NOT_RUNNING"
	VMToolsStatusNOTRUNNING VMToolsStatus = "NOT_RUNNING"

	// VMToolsStatusRESTRICTION captures enum value "RESTRICTION"
	VMToolsStatusRESTRICTION VMToolsStatus = "RESTRICTION"

	// VMToolsStatusRUNNING captures enum value "RUNNING"
	VMToolsStatusRUNNING VMToolsStatus = "RUNNING"
)

// for schema
var vmToolsStatusEnum []interface{}

func init() {
	var res []VMToolsStatus
	if err := json.Unmarshal([]byte(`["NOT_INSTALLED","NOT_RUNNING","RESTRICTION","RUNNING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmToolsStatusEnum = append(vmToolsStatusEnum, v)
	}
}

func (m VMToolsStatus) validateVMToolsStatusEnum(path, location string, value VMToolsStatus) error {
	if err := validate.EnumCase(path, location, value, vmToolsStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm tools status
func (m VMToolsStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMToolsStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm tools status based on context it is used
func (m VMToolsStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
