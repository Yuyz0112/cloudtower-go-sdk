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

// UserLoginInfoOrderByInput user login info order by input
//
// swagger:model UserLoginInfoOrderByInput
type UserLoginInfoOrderByInput string

func NewUserLoginInfoOrderByInput(value UserLoginInfoOrderByInput) *UserLoginInfoOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UserLoginInfoOrderByInput.
func (m UserLoginInfoOrderByInput) Pointer() *UserLoginInfoOrderByInput {
	return &m
}

const (

	// UserLoginInfoOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	UserLoginInfoOrderByInputCreatedAtASC UserLoginInfoOrderByInput = "createdAt_ASC"

	// UserLoginInfoOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	UserLoginInfoOrderByInputCreatedAtDESC UserLoginInfoOrderByInput = "createdAt_DESC"

	// UserLoginInfoOrderByInputIDASC captures enum value "id_ASC"
	UserLoginInfoOrderByInputIDASC UserLoginInfoOrderByInput = "id_ASC"

	// UserLoginInfoOrderByInputIDDESC captures enum value "id_DESC"
	UserLoginInfoOrderByInputIDDESC UserLoginInfoOrderByInput = "id_DESC"

	// UserLoginInfoOrderByInputLockedAtASC captures enum value "locked_at_ASC"
	UserLoginInfoOrderByInputLockedAtASC UserLoginInfoOrderByInput = "locked_at_ASC"

	// UserLoginInfoOrderByInputLockedAtDESC captures enum value "locked_at_DESC"
	UserLoginInfoOrderByInputLockedAtDESC UserLoginInfoOrderByInput = "locked_at_DESC"

	// UserLoginInfoOrderByInputMissNumASC captures enum value "miss_num_ASC"
	UserLoginInfoOrderByInputMissNumASC UserLoginInfoOrderByInput = "miss_num_ASC"

	// UserLoginInfoOrderByInputMissNumDESC captures enum value "miss_num_DESC"
	UserLoginInfoOrderByInputMissNumDESC UserLoginInfoOrderByInput = "miss_num_DESC"

	// UserLoginInfoOrderByInputMissedAtASC captures enum value "missed_at_ASC"
	UserLoginInfoOrderByInputMissedAtASC UserLoginInfoOrderByInput = "missed_at_ASC"

	// UserLoginInfoOrderByInputMissedAtDESC captures enum value "missed_at_DESC"
	UserLoginInfoOrderByInputMissedAtDESC UserLoginInfoOrderByInput = "missed_at_DESC"

	// UserLoginInfoOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	UserLoginInfoOrderByInputUpdatedAtASC UserLoginInfoOrderByInput = "updatedAt_ASC"

	// UserLoginInfoOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	UserLoginInfoOrderByInputUpdatedAtDESC UserLoginInfoOrderByInput = "updatedAt_DESC"
)

// for schema
var userLoginInfoOrderByInputEnum []interface{}

func init() {
	var res []UserLoginInfoOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","id_ASC","id_DESC","locked_at_ASC","locked_at_DESC","miss_num_ASC","miss_num_DESC","missed_at_ASC","missed_at_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userLoginInfoOrderByInputEnum = append(userLoginInfoOrderByInputEnum, v)
	}
}

func (m UserLoginInfoOrderByInput) validateUserLoginInfoOrderByInputEnum(path, location string, value UserLoginInfoOrderByInput) error {
	if err := validate.EnumCase(path, location, value, userLoginInfoOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user login info order by input
func (m UserLoginInfoOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserLoginInfoOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this user login info order by input based on context it is used
func (m UserLoginInfoOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
