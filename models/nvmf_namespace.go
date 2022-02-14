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

// NvmfNamespace nvmf namespace
//
// swagger:model NvmfNamespace
type NvmfNamespace struct {

	// assigned size
	// Required: true
	AssignedSize *float64 `json:"assigned_size"`

	// bps
	// Required: true
	Bps *float64 `json:"bps"`

	// bps max
	// Required: true
	BpsMax *float64 `json:"bps_max"`

	// bps max length
	// Required: true
	BpsMaxLength *float64 `json:"bps_max_length"`

	// bps rd
	// Required: true
	BpsRd *float64 `json:"bps_rd"`

	// bps rd max
	// Required: true
	BpsRdMax *float64 `json:"bps_rd_max"`

	// bps rd max length
	// Required: true
	BpsRdMaxLength *float64 `json:"bps_rd_max_length"`

	// bps wr
	// Required: true
	BpsWr *float64 `json:"bps_wr"`

	// bps wr max
	// Required: true
	BpsWrMax *float64 `json:"bps_wr_max"`

	// bps wr max length
	// Required: true
	BpsWrMaxLength *float64 `json:"bps_wr_max_length"`

	// consistency group
	ConsistencyGroup interface{} `json:"consistency_group,omitempty"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// io size
	// Required: true
	IoSize *float64 `json:"io_size"`

	// iops
	// Required: true
	Iops *float64 `json:"iops"`

	// iops max
	// Required: true
	IopsMax *float64 `json:"iops_max"`

	// iops max length
	// Required: true
	IopsMaxLength *float64 `json:"iops_max_length"`

	// iops rd
	// Required: true
	IopsRd *float64 `json:"iops_rd"`

	// iops rd max
	// Required: true
	IopsRdMax *float64 `json:"iops_rd_max"`

	// iops rd max length
	// Required: true
	IopsRdMaxLength *float64 `json:"iops_rd_max_length"`

	// iops wr
	// Required: true
	IopsWr *float64 `json:"iops_wr"`

	// iops wr max
	// Required: true
	IopsWrMax *float64 `json:"iops_wr_max"`

	// iops wr max length
	// Required: true
	IopsWrMaxLength *float64 `json:"iops_wr_max_length"`

	// is shared
	// Required: true
	IsShared *bool `json:"is_shared"`

	// labels
	Labels []*NestedLabel `json:"labels,omitempty"`

	// local created at
	// Required: true
	LocalCreatedAt *string `json:"local_created_at"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespace group
	NamespaceGroup interface{} `json:"namespace_group,omitempty"`

	// namespace id
	// Required: true
	NamespaceID *int32 `json:"namespace_id"`

	// nqn whitelist
	// Required: true
	NqnWhitelist *string `json:"nqn_whitelist"`

	// nvmf subsystem
	// Required: true
	NvmfSubsystem *NestedNvmfSubsystem `json:"nvmf_subsystem"`

	// replica num
	// Required: true
	ReplicaNum *int32 `json:"replica_num"`

	// shared size
	// Required: true
	SharedSize *float64 `json:"shared_size"`

	// snapshot num
	// Required: true
	SnapshotNum *int32 `json:"snapshot_num"`

	// stripe num
	// Required: true
	StripeNum *int32 `json:"stripe_num"`

	// stripe size
	// Required: true
	StripeSize *float64 `json:"stripe_size"`

	// thin provision
	// Required: true
	ThinProvision *bool `json:"thin_provision"`

	// unique size
	// Required: true
	UniqueSize *float64 `json:"unique_size"`

	// zbs volume id
	// Required: true
	ZbsVolumeID *string `json:"zbs_volume_id"`
}

// Validate validates this nvmf namespace
func (m *NvmfNamespace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsRd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsRdMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsRdMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsWr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsWrMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpsWrMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsRd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsRdMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsRdMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsWr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsWrMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsWrMaxLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsShared(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNqnWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThinProvision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniqueSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZbsVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmfNamespace) validateAssignedSize(formats strfmt.Registry) error {

	if err := validate.Required("assigned_size", "body", m.AssignedSize); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBps(formats strfmt.Registry) error {

	if err := validate.Required("bps", "body", m.Bps); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBpsMax(formats strfmt.Registry) error {

	if err := validate.Required("bps_max", "body", m.BpsMax); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBpsMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("bps_max_length", "body", m.BpsMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBpsRd(formats strfmt.Registry) error {

	if err := validate.Required("bps_rd", "body", m.BpsRd); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBpsRdMax(formats strfmt.Registry) error {

	if err := validate.Required("bps_rd_max", "body", m.BpsRdMax); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBpsRdMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("bps_rd_max_length", "body", m.BpsRdMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBpsWr(formats strfmt.Registry) error {

	if err := validate.Required("bps_wr", "body", m.BpsWr); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBpsWrMax(formats strfmt.Registry) error {

	if err := validate.Required("bps_wr_max", "body", m.BpsWrMax); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateBpsWrMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("bps_wr_max_length", "body", m.BpsWrMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIoSize(formats strfmt.Registry) error {

	if err := validate.Required("io_size", "body", m.IoSize); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIops(formats strfmt.Registry) error {

	if err := validate.Required("iops", "body", m.Iops); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIopsMax(formats strfmt.Registry) error {

	if err := validate.Required("iops_max", "body", m.IopsMax); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIopsMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("iops_max_length", "body", m.IopsMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIopsRd(formats strfmt.Registry) error {

	if err := validate.Required("iops_rd", "body", m.IopsRd); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIopsRdMax(formats strfmt.Registry) error {

	if err := validate.Required("iops_rd_max", "body", m.IopsRdMax); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIopsRdMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("iops_rd_max_length", "body", m.IopsRdMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIopsWr(formats strfmt.Registry) error {

	if err := validate.Required("iops_wr", "body", m.IopsWr); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIopsWrMax(formats strfmt.Registry) error {

	if err := validate.Required("iops_wr_max", "body", m.IopsWrMax); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIopsWrMaxLength(formats strfmt.Registry) error {

	if err := validate.Required("iops_wr_max_length", "body", m.IopsWrMaxLength); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateIsShared(formats strfmt.Registry) error {

	if err := validate.Required("is_shared", "body", m.IsShared); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateLabels(formats strfmt.Registry) error {
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

func (m *NvmfNamespace) validateLocalCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_created_at", "body", m.LocalCreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateNamespaceID(formats strfmt.Registry) error {

	if err := validate.Required("namespace_id", "body", m.NamespaceID); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateNqnWhitelist(formats strfmt.Registry) error {

	if err := validate.Required("nqn_whitelist", "body", m.NqnWhitelist); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateNvmfSubsystem(formats strfmt.Registry) error {

	if err := validate.Required("nvmf_subsystem", "body", m.NvmfSubsystem); err != nil {
		return err
	}

	if m.NvmfSubsystem != nil {
		if err := m.NvmfSubsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_subsystem")
			}
			return err
		}
	}

	return nil
}

func (m *NvmfNamespace) validateReplicaNum(formats strfmt.Registry) error {

	if err := validate.Required("replica_num", "body", m.ReplicaNum); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateSharedSize(formats strfmt.Registry) error {

	if err := validate.Required("shared_size", "body", m.SharedSize); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateSnapshotNum(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_num", "body", m.SnapshotNum); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateStripeNum(formats strfmt.Registry) error {

	if err := validate.Required("stripe_num", "body", m.StripeNum); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateStripeSize(formats strfmt.Registry) error {

	if err := validate.Required("stripe_size", "body", m.StripeSize); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateThinProvision(formats strfmt.Registry) error {

	if err := validate.Required("thin_provision", "body", m.ThinProvision); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateUniqueSize(formats strfmt.Registry) error {

	if err := validate.Required("unique_size", "body", m.UniqueSize); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespace) validateZbsVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("zbs_volume_id", "body", m.ZbsVolumeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nvmf namespace based on the context it is used
func (m *NvmfNamespace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmfNamespace) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

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

func (m *NvmfNamespace) contextValidateNvmfSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.NvmfSubsystem != nil {
		if err := m.NvmfSubsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_subsystem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NvmfNamespace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmfNamespace) UnmarshalBinary(b []byte) error {
	var res NvmfNamespace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
