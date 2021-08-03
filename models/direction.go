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

// Direction direction
//
// swagger:model Direction
type Direction string

func NewDirection(value Direction) *Direction {
	v := value
	return &v
}

const (

	// DirectionHORIZONTAL captures enum value "HORIZONTAL"
	DirectionHORIZONTAL Direction = "HORIZONTAL"

	// DirectionVERTICAL captures enum value "VERTICAL"
	DirectionVERTICAL Direction = "VERTICAL"
)

// for schema
var directionEnum []interface{}

func init() {
	var res []Direction
	if err := json.Unmarshal([]byte(`["HORIZONTAL","VERTICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directionEnum = append(directionEnum, v)
	}
}

func (m Direction) validateDirectionEnum(path, location string, value Direction) error {
	if err := validate.EnumCase(path, location, value, directionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this direction
func (m Direction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this direction based on context it is used
func (m Direction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
