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

// ElfStoragePolicyOrderByInput elf storage policy order by input
//
// swagger:model ElfStoragePolicyOrderByInput
type ElfStoragePolicyOrderByInput string

func NewElfStoragePolicyOrderByInput(value ElfStoragePolicyOrderByInput) *ElfStoragePolicyOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ElfStoragePolicyOrderByInput.
func (m ElfStoragePolicyOrderByInput) Pointer() *ElfStoragePolicyOrderByInput {
	return &m
}

const (

	// ElfStoragePolicyOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ElfStoragePolicyOrderByInputCreatedAtASC ElfStoragePolicyOrderByInput = "createdAt_ASC"

	// ElfStoragePolicyOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ElfStoragePolicyOrderByInputCreatedAtDESC ElfStoragePolicyOrderByInput = "createdAt_DESC"

	// ElfStoragePolicyOrderByInputDescriptionASC captures enum value "description_ASC"
	ElfStoragePolicyOrderByInputDescriptionASC ElfStoragePolicyOrderByInput = "description_ASC"

	// ElfStoragePolicyOrderByInputDescriptionDESC captures enum value "description_DESC"
	ElfStoragePolicyOrderByInputDescriptionDESC ElfStoragePolicyOrderByInput = "description_DESC"

	// ElfStoragePolicyOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	ElfStoragePolicyOrderByInputEntityAsyncStatusASC ElfStoragePolicyOrderByInput = "entityAsyncStatus_ASC"

	// ElfStoragePolicyOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	ElfStoragePolicyOrderByInputEntityAsyncStatusDESC ElfStoragePolicyOrderByInput = "entityAsyncStatus_DESC"

	// ElfStoragePolicyOrderByInputIDASC captures enum value "id_ASC"
	ElfStoragePolicyOrderByInputIDASC ElfStoragePolicyOrderByInput = "id_ASC"

	// ElfStoragePolicyOrderByInputIDDESC captures enum value "id_DESC"
	ElfStoragePolicyOrderByInputIDDESC ElfStoragePolicyOrderByInput = "id_DESC"

	// ElfStoragePolicyOrderByInputLocalIDASC captures enum value "local_id_ASC"
	ElfStoragePolicyOrderByInputLocalIDASC ElfStoragePolicyOrderByInput = "local_id_ASC"

	// ElfStoragePolicyOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	ElfStoragePolicyOrderByInputLocalIDDESC ElfStoragePolicyOrderByInput = "local_id_DESC"

	// ElfStoragePolicyOrderByInputNameASC captures enum value "name_ASC"
	ElfStoragePolicyOrderByInputNameASC ElfStoragePolicyOrderByInput = "name_ASC"

	// ElfStoragePolicyOrderByInputNameDESC captures enum value "name_DESC"
	ElfStoragePolicyOrderByInputNameDESC ElfStoragePolicyOrderByInput = "name_DESC"

	// ElfStoragePolicyOrderByInputReplicaNumASC captures enum value "replica_num_ASC"
	ElfStoragePolicyOrderByInputReplicaNumASC ElfStoragePolicyOrderByInput = "replica_num_ASC"

	// ElfStoragePolicyOrderByInputReplicaNumDESC captures enum value "replica_num_DESC"
	ElfStoragePolicyOrderByInputReplicaNumDESC ElfStoragePolicyOrderByInput = "replica_num_DESC"

	// ElfStoragePolicyOrderByInputStripeNumASC captures enum value "stripe_num_ASC"
	ElfStoragePolicyOrderByInputStripeNumASC ElfStoragePolicyOrderByInput = "stripe_num_ASC"

	// ElfStoragePolicyOrderByInputStripeNumDESC captures enum value "stripe_num_DESC"
	ElfStoragePolicyOrderByInputStripeNumDESC ElfStoragePolicyOrderByInput = "stripe_num_DESC"

	// ElfStoragePolicyOrderByInputStripeSizeASC captures enum value "stripe_size_ASC"
	ElfStoragePolicyOrderByInputStripeSizeASC ElfStoragePolicyOrderByInput = "stripe_size_ASC"

	// ElfStoragePolicyOrderByInputStripeSizeDESC captures enum value "stripe_size_DESC"
	ElfStoragePolicyOrderByInputStripeSizeDESC ElfStoragePolicyOrderByInput = "stripe_size_DESC"

	// ElfStoragePolicyOrderByInputThinProvisionASC captures enum value "thin_provision_ASC"
	ElfStoragePolicyOrderByInputThinProvisionASC ElfStoragePolicyOrderByInput = "thin_provision_ASC"

	// ElfStoragePolicyOrderByInputThinProvisionDESC captures enum value "thin_provision_DESC"
	ElfStoragePolicyOrderByInputThinProvisionDESC ElfStoragePolicyOrderByInput = "thin_provision_DESC"

	// ElfStoragePolicyOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	ElfStoragePolicyOrderByInputUpdatedAtASC ElfStoragePolicyOrderByInput = "updatedAt_ASC"

	// ElfStoragePolicyOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	ElfStoragePolicyOrderByInputUpdatedAtDESC ElfStoragePolicyOrderByInput = "updatedAt_DESC"
)

// for schema
var elfStoragePolicyOrderByInputEnum []interface{}

func init() {
	var res []ElfStoragePolicyOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","replica_num_ASC","replica_num_DESC","stripe_num_ASC","stripe_num_DESC","stripe_size_ASC","stripe_size_DESC","thin_provision_ASC","thin_provision_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		elfStoragePolicyOrderByInputEnum = append(elfStoragePolicyOrderByInputEnum, v)
	}
}

func (m ElfStoragePolicyOrderByInput) validateElfStoragePolicyOrderByInputEnum(path, location string, value ElfStoragePolicyOrderByInput) error {
	if err := validate.EnumCase(path, location, value, elfStoragePolicyOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this elf storage policy order by input
func (m ElfStoragePolicyOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateElfStoragePolicyOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this elf storage policy order by input based on context it is used
func (m ElfStoragePolicyOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
