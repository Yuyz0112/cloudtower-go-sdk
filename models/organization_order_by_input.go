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

// OrganizationOrderByInput organization order by input
//
// swagger:model OrganizationOrderByInput
type OrganizationOrderByInput string

func NewOrganizationOrderByInput(value OrganizationOrderByInput) *OrganizationOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OrganizationOrderByInput.
func (m OrganizationOrderByInput) Pointer() *OrganizationOrderByInput {
	return &m
}

const (

	// OrganizationOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	OrganizationOrderByInputCreatedAtASC OrganizationOrderByInput = "createdAt_ASC"

	// OrganizationOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	OrganizationOrderByInputCreatedAtDESC OrganizationOrderByInput = "createdAt_DESC"

	// OrganizationOrderByInputIDASC captures enum value "id_ASC"
	OrganizationOrderByInputIDASC OrganizationOrderByInput = "id_ASC"

	// OrganizationOrderByInputIDDESC captures enum value "id_DESC"
	OrganizationOrderByInputIDDESC OrganizationOrderByInput = "id_DESC"

	// OrganizationOrderByInputNameASC captures enum value "name_ASC"
	OrganizationOrderByInputNameASC OrganizationOrderByInput = "name_ASC"

	// OrganizationOrderByInputNameDESC captures enum value "name_DESC"
	OrganizationOrderByInputNameDESC OrganizationOrderByInput = "name_DESC"

	// OrganizationOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	OrganizationOrderByInputUpdatedAtASC OrganizationOrderByInput = "updatedAt_ASC"

	// OrganizationOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	OrganizationOrderByInputUpdatedAtDESC OrganizationOrderByInput = "updatedAt_DESC"
)

// for schema
var organizationOrderByInputEnum []interface{}

func init() {
	var res []OrganizationOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","id_ASC","id_DESC","name_ASC","name_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		organizationOrderByInputEnum = append(organizationOrderByInputEnum, v)
	}
}

func (m OrganizationOrderByInput) validateOrganizationOrderByInputEnum(path, location string, value OrganizationOrderByInput) error {
	if err := validate.EnumCase(path, location, value, organizationOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this organization order by input
func (m OrganizationOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOrganizationOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this organization order by input based on context it is used
func (m OrganizationOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
