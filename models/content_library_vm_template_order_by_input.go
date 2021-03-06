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

// ContentLibraryVMTemplateOrderByInput content library Vm template order by input
//
// swagger:model ContentLibraryVmTemplateOrderByInput
type ContentLibraryVMTemplateOrderByInput string

func NewContentLibraryVMTemplateOrderByInput(value ContentLibraryVMTemplateOrderByInput) *ContentLibraryVMTemplateOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ContentLibraryVMTemplateOrderByInput.
func (m ContentLibraryVMTemplateOrderByInput) Pointer() *ContentLibraryVMTemplateOrderByInput {
	return &m
}

const (

	// ContentLibraryVMTemplateOrderByInputArchitectureASC captures enum value "architecture_ASC"
	ContentLibraryVMTemplateOrderByInputArchitectureASC ContentLibraryVMTemplateOrderByInput = "architecture_ASC"

	// ContentLibraryVMTemplateOrderByInputArchitectureDESC captures enum value "architecture_DESC"
	ContentLibraryVMTemplateOrderByInputArchitectureDESC ContentLibraryVMTemplateOrderByInput = "architecture_DESC"

	// ContentLibraryVMTemplateOrderByInputCloudInitSupportedASC captures enum value "cloud_init_supported_ASC"
	ContentLibraryVMTemplateOrderByInputCloudInitSupportedASC ContentLibraryVMTemplateOrderByInput = "cloud_init_supported_ASC"

	// ContentLibraryVMTemplateOrderByInputCloudInitSupportedDESC captures enum value "cloud_init_supported_DESC"
	ContentLibraryVMTemplateOrderByInputCloudInitSupportedDESC ContentLibraryVMTemplateOrderByInput = "cloud_init_supported_DESC"

	// ContentLibraryVMTemplateOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ContentLibraryVMTemplateOrderByInputCreatedAtASC ContentLibraryVMTemplateOrderByInput = "createdAt_ASC"

	// ContentLibraryVMTemplateOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ContentLibraryVMTemplateOrderByInputCreatedAtDESC ContentLibraryVMTemplateOrderByInput = "createdAt_DESC"

	// ContentLibraryVMTemplateOrderByInputDescriptionASC captures enum value "description_ASC"
	ContentLibraryVMTemplateOrderByInputDescriptionASC ContentLibraryVMTemplateOrderByInput = "description_ASC"

	// ContentLibraryVMTemplateOrderByInputDescriptionDESC captures enum value "description_DESC"
	ContentLibraryVMTemplateOrderByInputDescriptionDESC ContentLibraryVMTemplateOrderByInput = "description_DESC"

	// ContentLibraryVMTemplateOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	ContentLibraryVMTemplateOrderByInputEntityAsyncStatusASC ContentLibraryVMTemplateOrderByInput = "entityAsyncStatus_ASC"

	// ContentLibraryVMTemplateOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	ContentLibraryVMTemplateOrderByInputEntityAsyncStatusDESC ContentLibraryVMTemplateOrderByInput = "entityAsyncStatus_DESC"

	// ContentLibraryVMTemplateOrderByInputIDASC captures enum value "id_ASC"
	ContentLibraryVMTemplateOrderByInputIDASC ContentLibraryVMTemplateOrderByInput = "id_ASC"

	// ContentLibraryVMTemplateOrderByInputIDDESC captures enum value "id_DESC"
	ContentLibraryVMTemplateOrderByInputIDDESC ContentLibraryVMTemplateOrderByInput = "id_DESC"

	// ContentLibraryVMTemplateOrderByInputMemoryASC captures enum value "memory_ASC"
	ContentLibraryVMTemplateOrderByInputMemoryASC ContentLibraryVMTemplateOrderByInput = "memory_ASC"

	// ContentLibraryVMTemplateOrderByInputMemoryDESC captures enum value "memory_DESC"
	ContentLibraryVMTemplateOrderByInputMemoryDESC ContentLibraryVMTemplateOrderByInput = "memory_DESC"

	// ContentLibraryVMTemplateOrderByInputNameASC captures enum value "name_ASC"
	ContentLibraryVMTemplateOrderByInputNameASC ContentLibraryVMTemplateOrderByInput = "name_ASC"

	// ContentLibraryVMTemplateOrderByInputNameDESC captures enum value "name_DESC"
	ContentLibraryVMTemplateOrderByInputNameDESC ContentLibraryVMTemplateOrderByInput = "name_DESC"

	// ContentLibraryVMTemplateOrderByInputOsASC captures enum value "os_ASC"
	ContentLibraryVMTemplateOrderByInputOsASC ContentLibraryVMTemplateOrderByInput = "os_ASC"

	// ContentLibraryVMTemplateOrderByInputOsDESC captures enum value "os_DESC"
	ContentLibraryVMTemplateOrderByInputOsDESC ContentLibraryVMTemplateOrderByInput = "os_DESC"

	// ContentLibraryVMTemplateOrderByInputSizeASC captures enum value "size_ASC"
	ContentLibraryVMTemplateOrderByInputSizeASC ContentLibraryVMTemplateOrderByInput = "size_ASC"

	// ContentLibraryVMTemplateOrderByInputSizeDESC captures enum value "size_DESC"
	ContentLibraryVMTemplateOrderByInputSizeDESC ContentLibraryVMTemplateOrderByInput = "size_DESC"

	// ContentLibraryVMTemplateOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	ContentLibraryVMTemplateOrderByInputUpdatedAtASC ContentLibraryVMTemplateOrderByInput = "updatedAt_ASC"

	// ContentLibraryVMTemplateOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	ContentLibraryVMTemplateOrderByInputUpdatedAtDESC ContentLibraryVMTemplateOrderByInput = "updatedAt_DESC"

	// ContentLibraryVMTemplateOrderByInputVcpuASC captures enum value "vcpu_ASC"
	ContentLibraryVMTemplateOrderByInputVcpuASC ContentLibraryVMTemplateOrderByInput = "vcpu_ASC"

	// ContentLibraryVMTemplateOrderByInputVcpuDESC captures enum value "vcpu_DESC"
	ContentLibraryVMTemplateOrderByInputVcpuDESC ContentLibraryVMTemplateOrderByInput = "vcpu_DESC"
)

// for schema
var contentLibraryVmTemplateOrderByInputEnum []interface{}

func init() {
	var res []ContentLibraryVMTemplateOrderByInput
	if err := json.Unmarshal([]byte(`["architecture_ASC","architecture_DESC","cloud_init_supported_ASC","cloud_init_supported_DESC","createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","id_ASC","id_DESC","memory_ASC","memory_DESC","name_ASC","name_DESC","os_ASC","os_DESC","size_ASC","size_DESC","updatedAt_ASC","updatedAt_DESC","vcpu_ASC","vcpu_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentLibraryVmTemplateOrderByInputEnum = append(contentLibraryVmTemplateOrderByInputEnum, v)
	}
}

func (m ContentLibraryVMTemplateOrderByInput) validateContentLibraryVMTemplateOrderByInputEnum(path, location string, value ContentLibraryVMTemplateOrderByInput) error {
	if err := validate.EnumCase(path, location, value, contentLibraryVmTemplateOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this content library Vm template order by input
func (m ContentLibraryVMTemplateOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContentLibraryVMTemplateOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this content library Vm template order by input based on context it is used
func (m ContentLibraryVMTemplateOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
