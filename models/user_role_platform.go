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

// UserRolePlatform user role platform
//
// swagger:model UserRolePlatform
type UserRolePlatform string

func NewUserRolePlatform(value UserRolePlatform) *UserRolePlatform {
	v := value
	return &v
}

const (

	// UserRolePlatformMANAGEMENT captures enum value "MANAGEMENT"
	UserRolePlatformMANAGEMENT UserRolePlatform = "MANAGEMENT"

	// UserRolePlatformSELFSERVICE captures enum value "SELF_SERVICE"
	UserRolePlatformSELFSERVICE UserRolePlatform = "SELF_SERVICE"
)

// for schema
var userRolePlatformEnum []interface{}

func init() {
	var res []UserRolePlatform
	if err := json.Unmarshal([]byte(`["MANAGEMENT","SELF_SERVICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userRolePlatformEnum = append(userRolePlatformEnum, v)
	}
}

func (m UserRolePlatform) validateUserRolePlatformEnum(path, location string, value UserRolePlatform) error {
	if err := validate.EnumCase(path, location, value, userRolePlatformEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user role platform
func (m UserRolePlatform) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserRolePlatformEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user role platform based on context it is used
func (m UserRolePlatform) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}