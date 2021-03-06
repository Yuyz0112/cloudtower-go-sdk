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

// NamespaceGroupOrderByInput namespace group order by input
//
// swagger:model NamespaceGroupOrderByInput
type NamespaceGroupOrderByInput string

func NewNamespaceGroupOrderByInput(value NamespaceGroupOrderByInput) *NamespaceGroupOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NamespaceGroupOrderByInput.
func (m NamespaceGroupOrderByInput) Pointer() *NamespaceGroupOrderByInput {
	return &m
}

const (

	// NamespaceGroupOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	NamespaceGroupOrderByInputCreatedAtASC NamespaceGroupOrderByInput = "createdAt_ASC"

	// NamespaceGroupOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	NamespaceGroupOrderByInputCreatedAtDESC NamespaceGroupOrderByInput = "createdAt_DESC"

	// NamespaceGroupOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	NamespaceGroupOrderByInputEntityAsyncStatusASC NamespaceGroupOrderByInput = "entityAsyncStatus_ASC"

	// NamespaceGroupOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	NamespaceGroupOrderByInputEntityAsyncStatusDESC NamespaceGroupOrderByInput = "entityAsyncStatus_DESC"

	// NamespaceGroupOrderByInputIDASC captures enum value "id_ASC"
	NamespaceGroupOrderByInputIDASC NamespaceGroupOrderByInput = "id_ASC"

	// NamespaceGroupOrderByInputIDDESC captures enum value "id_DESC"
	NamespaceGroupOrderByInputIDDESC NamespaceGroupOrderByInput = "id_DESC"

	// NamespaceGroupOrderByInputLocalCreateTimeASC captures enum value "local_create_time_ASC"
	NamespaceGroupOrderByInputLocalCreateTimeASC NamespaceGroupOrderByInput = "local_create_time_ASC"

	// NamespaceGroupOrderByInputLocalCreateTimeDESC captures enum value "local_create_time_DESC"
	NamespaceGroupOrderByInputLocalCreateTimeDESC NamespaceGroupOrderByInput = "local_create_time_DESC"

	// NamespaceGroupOrderByInputLocalIDASC captures enum value "local_id_ASC"
	NamespaceGroupOrderByInputLocalIDASC NamespaceGroupOrderByInput = "local_id_ASC"

	// NamespaceGroupOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	NamespaceGroupOrderByInputLocalIDDESC NamespaceGroupOrderByInput = "local_id_DESC"

	// NamespaceGroupOrderByInputNameASC captures enum value "name_ASC"
	NamespaceGroupOrderByInputNameASC NamespaceGroupOrderByInput = "name_ASC"

	// NamespaceGroupOrderByInputNameDESC captures enum value "name_DESC"
	NamespaceGroupOrderByInputNameDESC NamespaceGroupOrderByInput = "name_DESC"

	// NamespaceGroupOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	NamespaceGroupOrderByInputUpdatedAtASC NamespaceGroupOrderByInput = "updatedAt_ASC"

	// NamespaceGroupOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	NamespaceGroupOrderByInputUpdatedAtDESC NamespaceGroupOrderByInput = "updatedAt_DESC"
)

// for schema
var namespaceGroupOrderByInputEnum []interface{}

func init() {
	var res []NamespaceGroupOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","local_create_time_ASC","local_create_time_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		namespaceGroupOrderByInputEnum = append(namespaceGroupOrderByInputEnum, v)
	}
}

func (m NamespaceGroupOrderByInput) validateNamespaceGroupOrderByInputEnum(path, location string, value NamespaceGroupOrderByInput) error {
	if err := validate.EnumCase(path, location, value, namespaceGroupOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this namespace group order by input
func (m NamespaceGroupOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNamespaceGroupOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this namespace group order by input based on context it is used
func (m NamespaceGroupOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
