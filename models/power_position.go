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

// PowerPosition power position
//
// swagger:model PowerPosition
type PowerPosition string

func NewPowerPosition(value PowerPosition) *PowerPosition {
	v := value
	return &v
}

const (

	// PowerPositionLEFT captures enum value "LEFT"
	PowerPositionLEFT PowerPosition = "LEFT"

	// PowerPositionMIDDLE captures enum value "MIDDLE"
	PowerPositionMIDDLE PowerPosition = "MIDDLE"

	// PowerPositionRIGHT captures enum value "RIGHT"
	PowerPositionRIGHT PowerPosition = "RIGHT"
)

// for schema
var powerPositionEnum []interface{}

func init() {
	var res []PowerPosition
	if err := json.Unmarshal([]byte(`["LEFT","MIDDLE","RIGHT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerPositionEnum = append(powerPositionEnum, v)
	}
}

func (m PowerPosition) validatePowerPositionEnum(path, location string, value PowerPosition) error {
	if err := validate.EnumCase(path, location, value, powerPositionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this power position
func (m PowerPosition) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePowerPositionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this power position based on context it is used
func (m PowerPosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
