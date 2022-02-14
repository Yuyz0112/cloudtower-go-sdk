// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMTemplate Vm template
//
// swagger:model VmTemplate
type VMTemplate struct {

	// clock offset
	// Required: true
	ClockOffset *VMClockOffset `json:"clock_offset"`

	// cloud init supported
	// Required: true
	CloudInitSupported *bool `json:"cloud_init_supported"`

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// content library vm template
	ContentLibraryVMTemplate interface{} `json:"content_library_vm_template,omitempty"`

	// cpu
	// Required: true
	CPU *NestedCPU `json:"cpu"`

	// cpu model
	// Required: true
	CPUModel *string `json:"cpu_model"`

	// description
	// Required: true
	Description *string `json:"description"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// firmware
	// Required: true
	Firmware *VMFirmware `json:"firmware"`

	// ha
	// Required: true
	Ha *bool `json:"ha"`

	// id
	// Required: true
	ID *string `json:"id"`

	// io policy
	IoPolicy interface{} `json:"io_policy,omitempty"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// max bandwidth
	MaxBandwidth *float64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy interface{} `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops *int32 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy interface{} `json:"max_iops_policy,omitempty"`

	// memory
	// Required: true
	Memory *float64 `json:"memory"`

	// name
	// Required: true
	Name *string `json:"name"`

	// size
	// Required: true
	Size *float64 `json:"size"`

	// vcpu
	// Required: true
	Vcpu *int32 `json:"vcpu"`

	// video type
	VideoType *string `json:"video_type,omitempty"`

	// vm disks
	VMDisks []*NestedFrozenDisks `json:"vm_disks,omitempty"`

	// vm nics
	VMNics []*NestedTemplateNic `json:"vm_nics,omitempty"`

	// win opt
	// Required: true
	WinOpt *bool `json:"win_opt"`
}

// Validate validates this Vm template
func (m *VMTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClockOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudInitSupported(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWinOpt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplate) validateClockOffset(formats strfmt.Registry) error {

	if err := validate.Required("clock_offset", "body", m.ClockOffset); err != nil {
		return err
	}

	if err := validate.Required("clock_offset", "body", m.ClockOffset); err != nil {
		return err
	}

	if m.ClockOffset != nil {
		if err := m.ClockOffset.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateCloudInitSupported(formats strfmt.Registry) error {

	if err := validate.Required("cloud_init_supported", "body", m.CloudInitSupported); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateCPUModel(formats strfmt.Registry) error {

	if err := validate.Required("cpu_model", "body", m.CPUModel); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateFirmware(formats strfmt.Registry) error {

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
	}

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
	}

	if m.Firmware != nil {
		if err := m.Firmware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) validateHa(formats strfmt.Registry) error {

	if err := validate.Required("ha", "body", m.Ha); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateVcpu(formats strfmt.Registry) error {

	if err := validate.Required("vcpu", "body", m.Vcpu); err != nil {
		return err
	}

	return nil
}

func (m *VMTemplate) validateVMDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.VMDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.VMDisks); i++ {
		if swag.IsZero(m.VMDisks[i]) { // not required
			continue
		}

		if m.VMDisks[i] != nil {
			if err := m.VMDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	for i := 0; i < len(m.VMNics); i++ {
		if swag.IsZero(m.VMNics[i]) { // not required
			continue
		}

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) validateWinOpt(formats strfmt.Registry) error {

	if err := validate.Required("win_opt", "body", m.WinOpt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this Vm template based on the context it is used
func (m *VMTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClockOffset(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMTemplate) contextValidateClockOffset(ctx context.Context, formats strfmt.Registry) error {

	if m.ClockOffset != nil {
		if err := m.ClockOffset.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock_offset")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.CPU != nil {
		if err := m.CPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

	if m.Firmware != nil {
		if err := m.Firmware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMTemplate) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {
			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMDisks); i++ {

		if m.VMDisks[i] != nil {
			if err := m.VMDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMTemplate) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMNics); i++ {

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMTemplate) UnmarshalBinary(b []byte) error {
	var res VMTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
