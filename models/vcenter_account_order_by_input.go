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

// VcenterAccountOrderByInput vcenter account order by input
//
// swagger:model VcenterAccountOrderByInput
type VcenterAccountOrderByInput string

func NewVcenterAccountOrderByInput(value VcenterAccountOrderByInput) *VcenterAccountOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VcenterAccountOrderByInput.
func (m VcenterAccountOrderByInput) Pointer() *VcenterAccountOrderByInput {
	return &m
}

const (

	// VcenterAccountOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	VcenterAccountOrderByInputCreatedAtASC VcenterAccountOrderByInput = "createdAt_ASC"

	// VcenterAccountOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	VcenterAccountOrderByInputCreatedAtDESC VcenterAccountOrderByInput = "createdAt_DESC"

	// VcenterAccountOrderByInputIDASC captures enum value "id_ASC"
	VcenterAccountOrderByInputIDASC VcenterAccountOrderByInput = "id_ASC"

	// VcenterAccountOrderByInputIDDESC captures enum value "id_DESC"
	VcenterAccountOrderByInputIDDESC VcenterAccountOrderByInput = "id_DESC"

	// VcenterAccountOrderByInputIPASC captures enum value "ip_ASC"
	VcenterAccountOrderByInputIPASC VcenterAccountOrderByInput = "ip_ASC"

	// VcenterAccountOrderByInputIPDESC captures enum value "ip_DESC"
	VcenterAccountOrderByInputIPDESC VcenterAccountOrderByInput = "ip_DESC"

	// VcenterAccountOrderByInputIsValidASC captures enum value "is_valid_ASC"
	VcenterAccountOrderByInputIsValidASC VcenterAccountOrderByInput = "is_valid_ASC"

	// VcenterAccountOrderByInputIsValidDESC captures enum value "is_valid_DESC"
	VcenterAccountOrderByInputIsValidDESC VcenterAccountOrderByInput = "is_valid_DESC"

	// VcenterAccountOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VcenterAccountOrderByInputLocalIDASC VcenterAccountOrderByInput = "local_id_ASC"

	// VcenterAccountOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VcenterAccountOrderByInputLocalIDDESC VcenterAccountOrderByInput = "local_id_DESC"

	// VcenterAccountOrderByInputPasswordASC captures enum value "password_ASC"
	VcenterAccountOrderByInputPasswordASC VcenterAccountOrderByInput = "password_ASC"

	// VcenterAccountOrderByInputPasswordDESC captures enum value "password_DESC"
	VcenterAccountOrderByInputPasswordDESC VcenterAccountOrderByInput = "password_DESC"

	// VcenterAccountOrderByInputPortASC captures enum value "port_ASC"
	VcenterAccountOrderByInputPortASC VcenterAccountOrderByInput = "port_ASC"

	// VcenterAccountOrderByInputPortDESC captures enum value "port_DESC"
	VcenterAccountOrderByInputPortDESC VcenterAccountOrderByInput = "port_DESC"

	// VcenterAccountOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	VcenterAccountOrderByInputUpdatedAtASC VcenterAccountOrderByInput = "updatedAt_ASC"

	// VcenterAccountOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	VcenterAccountOrderByInputUpdatedAtDESC VcenterAccountOrderByInput = "updatedAt_DESC"

	// VcenterAccountOrderByInputUsernameASC captures enum value "username_ASC"
	VcenterAccountOrderByInputUsernameASC VcenterAccountOrderByInput = "username_ASC"

	// VcenterAccountOrderByInputUsernameDESC captures enum value "username_DESC"
	VcenterAccountOrderByInputUsernameDESC VcenterAccountOrderByInput = "username_DESC"
)

// for schema
var vcenterAccountOrderByInputEnum []interface{}

func init() {
	var res []VcenterAccountOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","id_ASC","id_DESC","ip_ASC","ip_DESC","is_valid_ASC","is_valid_DESC","local_id_ASC","local_id_DESC","password_ASC","password_DESC","port_ASC","port_DESC","updatedAt_ASC","updatedAt_DESC","username_ASC","username_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vcenterAccountOrderByInputEnum = append(vcenterAccountOrderByInputEnum, v)
	}
}

func (m VcenterAccountOrderByInput) validateVcenterAccountOrderByInputEnum(path, location string, value VcenterAccountOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vcenterAccountOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this vcenter account order by input
func (m VcenterAccountOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVcenterAccountOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this vcenter account order by input based on context it is used
func (m VcenterAccountOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
