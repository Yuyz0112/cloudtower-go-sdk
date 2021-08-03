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

// VMSnapshotOrderByInput Vm snapshot order by input
//
// swagger:model VmSnapshotOrderByInput
type VMSnapshotOrderByInput string

func NewVMSnapshotOrderByInput(value VMSnapshotOrderByInput) *VMSnapshotOrderByInput {
	v := value
	return &v
}

const (

	// VMSnapshotOrderByInputClockOffsetASC captures enum value "clock_offset_ASC"
	VMSnapshotOrderByInputClockOffsetASC VMSnapshotOrderByInput = "clock_offset_ASC"

	// VMSnapshotOrderByInputClockOffsetDESC captures enum value "clock_offset_DESC"
	VMSnapshotOrderByInputClockOffsetDESC VMSnapshotOrderByInput = "clock_offset_DESC"

	// VMSnapshotOrderByInputConsistentTypeASC captures enum value "consistent_type_ASC"
	VMSnapshotOrderByInputConsistentTypeASC VMSnapshotOrderByInput = "consistent_type_ASC"

	// VMSnapshotOrderByInputConsistentTypeDESC captures enum value "consistent_type_DESC"
	VMSnapshotOrderByInputConsistentTypeDESC VMSnapshotOrderByInput = "consistent_type_DESC"

	// VMSnapshotOrderByInputCPUASC captures enum value "cpu_ASC"
	VMSnapshotOrderByInputCPUASC VMSnapshotOrderByInput = "cpu_ASC"

	// VMSnapshotOrderByInputCPUDESC captures enum value "cpu_DESC"
	VMSnapshotOrderByInputCPUDESC VMSnapshotOrderByInput = "cpu_DESC"

	// VMSnapshotOrderByInputCPUModelASC captures enum value "cpu_model_ASC"
	VMSnapshotOrderByInputCPUModelASC VMSnapshotOrderByInput = "cpu_model_ASC"

	// VMSnapshotOrderByInputCPUModelDESC captures enum value "cpu_model_DESC"
	VMSnapshotOrderByInputCPUModelDESC VMSnapshotOrderByInput = "cpu_model_DESC"

	// VMSnapshotOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	VMSnapshotOrderByInputCreatedAtASC VMSnapshotOrderByInput = "createdAt_ASC"

	// VMSnapshotOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	VMSnapshotOrderByInputCreatedAtDESC VMSnapshotOrderByInput = "createdAt_DESC"

	// VMSnapshotOrderByInputDescriptionASC captures enum value "description_ASC"
	VMSnapshotOrderByInputDescriptionASC VMSnapshotOrderByInput = "description_ASC"

	// VMSnapshotOrderByInputDescriptionDESC captures enum value "description_DESC"
	VMSnapshotOrderByInputDescriptionDESC VMSnapshotOrderByInput = "description_DESC"

	// VMSnapshotOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	VMSnapshotOrderByInputEntityAsyncStatusASC VMSnapshotOrderByInput = "entityAsyncStatus_ASC"

	// VMSnapshotOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	VMSnapshotOrderByInputEntityAsyncStatusDESC VMSnapshotOrderByInput = "entityAsyncStatus_DESC"

	// VMSnapshotOrderByInputFirmwareASC captures enum value "firmware_ASC"
	VMSnapshotOrderByInputFirmwareASC VMSnapshotOrderByInput = "firmware_ASC"

	// VMSnapshotOrderByInputFirmwareDESC captures enum value "firmware_DESC"
	VMSnapshotOrderByInputFirmwareDESC VMSnapshotOrderByInput = "firmware_DESC"

	// VMSnapshotOrderByInputHaASC captures enum value "ha_ASC"
	VMSnapshotOrderByInputHaASC VMSnapshotOrderByInput = "ha_ASC"

	// VMSnapshotOrderByInputHaDESC captures enum value "ha_DESC"
	VMSnapshotOrderByInputHaDESC VMSnapshotOrderByInput = "ha_DESC"

	// VMSnapshotOrderByInputIDASC captures enum value "id_ASC"
	VMSnapshotOrderByInputIDASC VMSnapshotOrderByInput = "id_ASC"

	// VMSnapshotOrderByInputIDDESC captures enum value "id_DESC"
	VMSnapshotOrderByInputIDDESC VMSnapshotOrderByInput = "id_DESC"

	// VMSnapshotOrderByInputIoPolicyASC captures enum value "io_policy_ASC"
	VMSnapshotOrderByInputIoPolicyASC VMSnapshotOrderByInput = "io_policy_ASC"

	// VMSnapshotOrderByInputIoPolicyDESC captures enum value "io_policy_DESC"
	VMSnapshotOrderByInputIoPolicyDESC VMSnapshotOrderByInput = "io_policy_DESC"

	// VMSnapshotOrderByInputLocalCreatedAtASC captures enum value "local_created_at_ASC"
	VMSnapshotOrderByInputLocalCreatedAtASC VMSnapshotOrderByInput = "local_created_at_ASC"

	// VMSnapshotOrderByInputLocalCreatedAtDESC captures enum value "local_created_at_DESC"
	VMSnapshotOrderByInputLocalCreatedAtDESC VMSnapshotOrderByInput = "local_created_at_DESC"

	// VMSnapshotOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VMSnapshotOrderByInputLocalIDASC VMSnapshotOrderByInput = "local_id_ASC"

	// VMSnapshotOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VMSnapshotOrderByInputLocalIDDESC VMSnapshotOrderByInput = "local_id_DESC"

	// VMSnapshotOrderByInputMaxBandwidthASC captures enum value "max_bandwidth_ASC"
	VMSnapshotOrderByInputMaxBandwidthASC VMSnapshotOrderByInput = "max_bandwidth_ASC"

	// VMSnapshotOrderByInputMaxBandwidthDESC captures enum value "max_bandwidth_DESC"
	VMSnapshotOrderByInputMaxBandwidthDESC VMSnapshotOrderByInput = "max_bandwidth_DESC"

	// VMSnapshotOrderByInputMaxBandwidthPolicyASC captures enum value "max_bandwidth_policy_ASC"
	VMSnapshotOrderByInputMaxBandwidthPolicyASC VMSnapshotOrderByInput = "max_bandwidth_policy_ASC"

	// VMSnapshotOrderByInputMaxBandwidthPolicyDESC captures enum value "max_bandwidth_policy_DESC"
	VMSnapshotOrderByInputMaxBandwidthPolicyDESC VMSnapshotOrderByInput = "max_bandwidth_policy_DESC"

	// VMSnapshotOrderByInputMaxIopsASC captures enum value "max_iops_ASC"
	VMSnapshotOrderByInputMaxIopsASC VMSnapshotOrderByInput = "max_iops_ASC"

	// VMSnapshotOrderByInputMaxIopsDESC captures enum value "max_iops_DESC"
	VMSnapshotOrderByInputMaxIopsDESC VMSnapshotOrderByInput = "max_iops_DESC"

	// VMSnapshotOrderByInputMaxIopsPolicyASC captures enum value "max_iops_policy_ASC"
	VMSnapshotOrderByInputMaxIopsPolicyASC VMSnapshotOrderByInput = "max_iops_policy_ASC"

	// VMSnapshotOrderByInputMaxIopsPolicyDESC captures enum value "max_iops_policy_DESC"
	VMSnapshotOrderByInputMaxIopsPolicyDESC VMSnapshotOrderByInput = "max_iops_policy_DESC"

	// VMSnapshotOrderByInputMemoryASC captures enum value "memory_ASC"
	VMSnapshotOrderByInputMemoryASC VMSnapshotOrderByInput = "memory_ASC"

	// VMSnapshotOrderByInputMemoryDESC captures enum value "memory_DESC"
	VMSnapshotOrderByInputMemoryDESC VMSnapshotOrderByInput = "memory_DESC"

	// VMSnapshotOrderByInputNameASC captures enum value "name_ASC"
	VMSnapshotOrderByInputNameASC VMSnapshotOrderByInput = "name_ASC"

	// VMSnapshotOrderByInputNameDESC captures enum value "name_DESC"
	VMSnapshotOrderByInputNameDESC VMSnapshotOrderByInput = "name_DESC"

	// VMSnapshotOrderByInputSizeASC captures enum value "size_ASC"
	VMSnapshotOrderByInputSizeASC VMSnapshotOrderByInput = "size_ASC"

	// VMSnapshotOrderByInputSizeDESC captures enum value "size_DESC"
	VMSnapshotOrderByInputSizeDESC VMSnapshotOrderByInput = "size_DESC"

	// VMSnapshotOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	VMSnapshotOrderByInputUpdatedAtASC VMSnapshotOrderByInput = "updatedAt_ASC"

	// VMSnapshotOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	VMSnapshotOrderByInputUpdatedAtDESC VMSnapshotOrderByInput = "updatedAt_DESC"

	// VMSnapshotOrderByInputVcpuASC captures enum value "vcpu_ASC"
	VMSnapshotOrderByInputVcpuASC VMSnapshotOrderByInput = "vcpu_ASC"

	// VMSnapshotOrderByInputVcpuDESC captures enum value "vcpu_DESC"
	VMSnapshotOrderByInputVcpuDESC VMSnapshotOrderByInput = "vcpu_DESC"

	// VMSnapshotOrderByInputVMDisksASC captures enum value "vm_disks_ASC"
	VMSnapshotOrderByInputVMDisksASC VMSnapshotOrderByInput = "vm_disks_ASC"

	// VMSnapshotOrderByInputVMDisksDESC captures enum value "vm_disks_DESC"
	VMSnapshotOrderByInputVMDisksDESC VMSnapshotOrderByInput = "vm_disks_DESC"

	// VMSnapshotOrderByInputVMNicsASC captures enum value "vm_nics_ASC"
	VMSnapshotOrderByInputVMNicsASC VMSnapshotOrderByInput = "vm_nics_ASC"

	// VMSnapshotOrderByInputVMNicsDESC captures enum value "vm_nics_DESC"
	VMSnapshotOrderByInputVMNicsDESC VMSnapshotOrderByInput = "vm_nics_DESC"

	// VMSnapshotOrderByInputWinOptASC captures enum value "win_opt_ASC"
	VMSnapshotOrderByInputWinOptASC VMSnapshotOrderByInput = "win_opt_ASC"

	// VMSnapshotOrderByInputWinOptDESC captures enum value "win_opt_DESC"
	VMSnapshotOrderByInputWinOptDESC VMSnapshotOrderByInput = "win_opt_DESC"
)

// for schema
var vmSnapshotOrderByInputEnum []interface{}

func init() {
	var res []VMSnapshotOrderByInput
	if err := json.Unmarshal([]byte(`["clock_offset_ASC","clock_offset_DESC","consistent_type_ASC","consistent_type_DESC","cpu_ASC","cpu_DESC","cpu_model_ASC","cpu_model_DESC","createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","firmware_ASC","firmware_DESC","ha_ASC","ha_DESC","id_ASC","id_DESC","io_policy_ASC","io_policy_DESC","local_created_at_ASC","local_created_at_DESC","local_id_ASC","local_id_DESC","max_bandwidth_ASC","max_bandwidth_DESC","max_bandwidth_policy_ASC","max_bandwidth_policy_DESC","max_iops_ASC","max_iops_DESC","max_iops_policy_ASC","max_iops_policy_DESC","memory_ASC","memory_DESC","name_ASC","name_DESC","size_ASC","size_DESC","updatedAt_ASC","updatedAt_DESC","vcpu_ASC","vcpu_DESC","vm_disks_ASC","vm_disks_DESC","vm_nics_ASC","vm_nics_DESC","win_opt_ASC","win_opt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmSnapshotOrderByInputEnum = append(vmSnapshotOrderByInputEnum, v)
	}
}

func (m VMSnapshotOrderByInput) validateVMSnapshotOrderByInputEnum(path, location string, value VMSnapshotOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vmSnapshotOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Vm snapshot order by input
func (m VMSnapshotOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVMSnapshotOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Vm snapshot order by input based on context it is used
func (m VMSnapshotOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}