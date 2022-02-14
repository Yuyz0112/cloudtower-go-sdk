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

// VMVolumeOrderByInput Vm volume order by input
//
// swagger:model VmVolumeOrderByInput
type VMVolumeOrderByInput string

func NewVMVolumeOrderByInput(value VMVolumeOrderByInput) *VMVolumeOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VMVolumeOrderByInput.
func (m VMVolumeOrderByInput) Pointer() *VMVolumeOrderByInput {
	return &m
}

const (

	// VMVolumeOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	VMVolumeOrderByInputCreatedAtASC VMVolumeOrderByInput = "createdAt_ASC"

	// VMVolumeOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	VMVolumeOrderByInputCreatedAtDESC VMVolumeOrderByInput = "createdAt_DESC"

	// VMVolumeOrderByInputDescriptionASC captures enum value "description_ASC"
	VMVolumeOrderByInputDescriptionASC VMVolumeOrderByInput = "description_ASC"

	// VMVolumeOrderByInputDescriptionDESC captures enum value "description_DESC"
	VMVolumeOrderByInputDescriptionDESC VMVolumeOrderByInput = "description_DESC"

	// VMVolumeOrderByInputElfStoragePolicyASC captures enum value "elf_storage_policy_ASC"
	VMVolumeOrderByInputElfStoragePolicyASC VMVolumeOrderByInput = "elf_storage_policy_ASC"

	// VMVolumeOrderByInputElfStoragePolicyDESC captures enum value "elf_storage_policy_DESC"
	VMVolumeOrderByInputElfStoragePolicyDESC VMVolumeOrderByInput = "elf_storage_policy_DESC"

	// VMVolumeOrderByInputGuestSizeUsageASC captures enum value "guest_size_usage_ASC"
	VMVolumeOrderByInputGuestSizeUsageASC VMVolumeOrderByInput = "guest_size_usage_ASC"

	// VMVolumeOrderByInputGuestSizeUsageDESC captures enum value "guest_size_usage_DESC"
	VMVolumeOrderByInputGuestSizeUsageDESC VMVolumeOrderByInput = "guest_size_usage_DESC"

	// VMVolumeOrderByInputGuestUsedSizeASC captures enum value "guest_used_size_ASC"
	VMVolumeOrderByInputGuestUsedSizeASC VMVolumeOrderByInput = "guest_used_size_ASC"

	// VMVolumeOrderByInputGuestUsedSizeDESC captures enum value "guest_used_size_DESC"
	VMVolumeOrderByInputGuestUsedSizeDESC VMVolumeOrderByInput = "guest_used_size_DESC"

	// VMVolumeOrderByInputIDASC captures enum value "id_ASC"
	VMVolumeOrderByInputIDASC VMVolumeOrderByInput = "id_ASC"

	// VMVolumeOrderByInputIDDESC captures enum value "id_DESC"
	VMVolumeOrderByInputIDDESC VMVolumeOrderByInput = "id_DESC"

	// VMVolumeOrderByInputLocalCreatedAtASC captures enum value "local_created_at_ASC"
	VMVolumeOrderByInputLocalCreatedAtASC VMVolumeOrderByInput = "local_created_at_ASC"

	// VMVolumeOrderByInputLocalCreatedAtDESC captures enum value "local_created_at_DESC"
	VMVolumeOrderByInputLocalCreatedAtDESC VMVolumeOrderByInput = "local_created_at_DESC"

	// VMVolumeOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VMVolumeOrderByInputLocalIDASC VMVolumeOrderByInput = "local_id_ASC"

	// VMVolumeOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VMVolumeOrderByInputLocalIDDESC VMVolumeOrderByInput = "local_id_DESC"

	// VMVolumeOrderByInputMountingASC captures enum value "mounting_ASC"
	VMVolumeOrderByInputMountingASC VMVolumeOrderByInput = "mounting_ASC"

	// VMVolumeOrderByInputMountingDESC captures enum value "mounting_DESC"
	VMVolumeOrderByInputMountingDESC VMVolumeOrderByInput = "mounting_DESC"

	// VMVolumeOrderByInputNameASC captures enum value "name_ASC"
	VMVolumeOrderByInputNameASC VMVolumeOrderByInput = "name_ASC"

	// VMVolumeOrderByInputNameDESC captures enum value "name_DESC"
	VMVolumeOrderByInputNameDESC VMVolumeOrderByInput = "name_DESC"

	// VMVolumeOrderByInputPathASC captures enum value "path_ASC"
	VMVolumeOrderByInputPathASC VMVolumeOrderByInput = "path_ASC"

	// VMVolumeOrderByInputPathDESC captures enum value "path_DESC"
	VMVolumeOrderByInputPathDESC VMVolumeOrderByInput = "path_DESC"

	// VMVolumeOrderByInputSharingASC captures enum value "sharing_ASC"
	VMVolumeOrderByInputSharingASC VMVolumeOrderByInput = "sharing_ASC"

	// VMVolumeOrderByInputSharingDESC captures enum value "sharing_DESC"
	VMVolumeOrderByInputSharingDESC VMVolumeOrderByInput = "sharing_DESC"

	// VMVolumeOrderByInputSizeASC captures enum value "size_ASC"
	VMVolumeOrderByInputSizeASC VMVolumeOrderByInput = "size_ASC"

	// VMVolumeOrderByInputSizeDESC captures enum value "size_DESC"
	VMVolumeOrderByInputSizeDESC VMVolumeOrderByInput = "size_DESC"

	// VMVolumeOrderByInputUniqueSizeASC captures enum value "unique_size_ASC"
	VMVolumeOrderByInputUniqueSizeASC VMVolumeOrderByInput = "unique_size_ASC"

	// VMVolumeOrderByInputUniqueSizeDESC captures enum value "unique_size_DESC"
	VMVolumeOrderByInputUniqueSizeDESC VMVolumeOrderByInput = "unique_size_DESC"

	// VMVolumeOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	VMVolumeOrderByInputUpdatedAtASC VMVolumeOrderByInput = "updatedAt_ASC"

	// VMVolumeOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	VMVolumeOrderByInputUpdatedAtDESC VMVolumeOrderByInput = "updatedAt_DESC"
)

// for schema
var vmVolumeOrderByInputEnum []interface{}

func init() {
	var res []VMVolumeOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","elf_storage_policy_ASC","elf_storage_policy_DESC","guest_size_usage_ASC","guest_size_usage_DESC","guest_used_size_ASC","guest_used_size_DESC","id_ASC","id_DESC","local_created_at_ASC","local_created_at_DESC","local_id_ASC","local_id_DESC","mounting_ASC","mounting_DESC","name_ASC","name_DESC","path_ASC","path_DESC","sharing_ASC","sharing_DESC","size_ASC","size_DESC","unique_size_ASC","unique_size_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmVolumeOrderByInputEnum = append(vmVolumeOrderByInputEnum, v)
	}
}

func (m VMVolumeOrderByInput) validateVMVolumeOrderByInputEnum(path, location string, value VMVolumeOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vmVolumeOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm volume order by input
func (m VMVolumeOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMVolumeOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm volume order by input based on context it is used
func (m VMVolumeOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
