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
)

// ApplicationWhereInput application where input
// Example: {"AND":"ApplicationWhereInput[]","NOT":"ApplicationWhereInput[]","OR":"ApplicationWhereInput[]","cluster":"ClusterWhereInput","error_message":"string","error_message_contains":"string","error_message_ends_with":"string","error_message_gt":"string","error_message_gte":"string","error_message_in":["string"],"error_message_lt":"string","error_message_lte":"string","error_message_not":"string","error_message_not_contains":"string","error_message_not_ends_with":"string","error_message_not_in":["string"],"error_message_not_starts_with":"string","error_message_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","image_name":"string","image_name_contains":"string","image_name_ends_with":"string","image_name_gt":"string","image_name_gte":"string","image_name_in":["string"],"image_name_lt":"string","image_name_lte":"string","image_name_not":"string","image_name_not_contains":"string","image_name_not_ends_with":"string","image_name_not_in":["string"],"image_name_not_starts_with":"string","image_name_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","memory":0,"memory_gt":0,"memory_gte":0,"memory_in":[0],"memory_lt":0,"memory_lte":0,"memory_not":0,"memory_not_in":[0],"state":"DEPLOY_ERROR","state_in":["DEPLOY_ERROR"],"state_not":"DEPLOY_ERROR","state_not_in":["DEPLOY_ERROR"],"storage_ip":"string","storage_ip_contains":"string","storage_ip_ends_with":"string","storage_ip_gt":"string","storage_ip_gte":"string","storage_ip_in":["string"],"storage_ip_lt":"string","storage_ip_lte":"string","storage_ip_not":"string","storage_ip_not_contains":"string","storage_ip_not_ends_with":"string","storage_ip_not_in":["string"],"storage_ip_not_starts_with":"string","storage_ip_starts_with":"string","type":"MONITOR","type_in":["MONITOR"],"type_not":"MONITOR","type_not_in":["MONITOR"],"update_time":"string","update_time_gt":"string","update_time_gte":"string","update_time_in":["string"],"update_time_lt":"string","update_time_lte":"string","update_time_not":"string","update_time_not_in":["string"],"vcpu":0,"vcpu_gt":0,"vcpu_gte":0,"vcpu_in":[0],"vcpu_lt":0,"vcpu_lte":0,"vcpu_not":0,"vcpu_not_in":[0],"version":"string","version_contains":"string","version_ends_with":"string","version_gt":"string","version_gte":"string","version_in":["string"],"version_lt":"string","version_lte":"string","version_not":"string","version_not_contains":"string","version_not_ends_with":"string","version_not_in":["string"],"version_not_starts_with":"string","version_starts_with":"string","vm":"VmWhereInput","volume_size":0,"volume_size_gt":0,"volume_size_gte":0,"volume_size_in":[0],"volume_size_lt":0,"volume_size_lte":0,"volume_size_not":0,"volume_size_not_in":[0]}
//
// swagger:model ApplicationWhereInput
type ApplicationWhereInput struct {

	// a n d
	AND []*ApplicationWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*ApplicationWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*ApplicationWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// error message
	ErrorMessage *string `json:"error_message,omitempty"`

	// error message contains
	ErrorMessageContains *string `json:"error_message_contains,omitempty"`

	// error message ends with
	ErrorMessageEndsWith *string `json:"error_message_ends_with,omitempty"`

	// error message gt
	ErrorMessageGt *string `json:"error_message_gt,omitempty"`

	// error message gte
	ErrorMessageGte *string `json:"error_message_gte,omitempty"`

	// error message in
	ErrorMessageIn []string `json:"error_message_in,omitempty"`

	// error message lt
	ErrorMessageLt *string `json:"error_message_lt,omitempty"`

	// error message lte
	ErrorMessageLte *string `json:"error_message_lte,omitempty"`

	// error message not
	ErrorMessageNot *string `json:"error_message_not,omitempty"`

	// error message not contains
	ErrorMessageNotContains *string `json:"error_message_not_contains,omitempty"`

	// error message not ends with
	ErrorMessageNotEndsWith *string `json:"error_message_not_ends_with,omitempty"`

	// error message not in
	ErrorMessageNotIn []string `json:"error_message_not_in,omitempty"`

	// error message not starts with
	ErrorMessageNotStartsWith *string `json:"error_message_not_starts_with,omitempty"`

	// error message starts with
	ErrorMessageStartsWith *string `json:"error_message_starts_with,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// image name
	ImageName *string `json:"image_name,omitempty"`

	// image name contains
	ImageNameContains *string `json:"image_name_contains,omitempty"`

	// image name ends with
	ImageNameEndsWith *string `json:"image_name_ends_with,omitempty"`

	// image name gt
	ImageNameGt *string `json:"image_name_gt,omitempty"`

	// image name gte
	ImageNameGte *string `json:"image_name_gte,omitempty"`

	// image name in
	ImageNameIn []string `json:"image_name_in,omitempty"`

	// image name lt
	ImageNameLt *string `json:"image_name_lt,omitempty"`

	// image name lte
	ImageNameLte *string `json:"image_name_lte,omitempty"`

	// image name not
	ImageNameNot *string `json:"image_name_not,omitempty"`

	// image name not contains
	ImageNameNotContains *string `json:"image_name_not_contains,omitempty"`

	// image name not ends with
	ImageNameNotEndsWith *string `json:"image_name_not_ends_with,omitempty"`

	// image name not in
	ImageNameNotIn []string `json:"image_name_not_in,omitempty"`

	// image name not starts with
	ImageNameNotStartsWith *string `json:"image_name_not_starts_with,omitempty"`

	// image name starts with
	ImageNameStartsWith *string `json:"image_name_starts_with,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// memory
	Memory *float64 `json:"memory,omitempty"`

	// memory gt
	MemoryGt *float64 `json:"memory_gt,omitempty"`

	// memory gte
	MemoryGte *float64 `json:"memory_gte,omitempty"`

	// memory in
	MemoryIn []float64 `json:"memory_in,omitempty"`

	// memory lt
	MemoryLt *float64 `json:"memory_lt,omitempty"`

	// memory lte
	MemoryLte *float64 `json:"memory_lte,omitempty"`

	// memory not
	MemoryNot *float64 `json:"memory_not,omitempty"`

	// memory not in
	MemoryNotIn []float64 `json:"memory_not_in,omitempty"`

	// state
	State *ApplicationState `json:"state,omitempty"`

	// state in
	StateIn []ApplicationState `json:"state_in,omitempty"`

	// state not
	StateNot *ApplicationState `json:"state_not,omitempty"`

	// state not in
	StateNotIn []ApplicationState `json:"state_not_in,omitempty"`

	// storage ip
	StorageIP *string `json:"storage_ip,omitempty"`

	// storage ip contains
	StorageIPContains *string `json:"storage_ip_contains,omitempty"`

	// storage ip ends with
	StorageIPEndsWith *string `json:"storage_ip_ends_with,omitempty"`

	// storage ip gt
	StorageIPGt *string `json:"storage_ip_gt,omitempty"`

	// storage ip gte
	StorageIPGte *string `json:"storage_ip_gte,omitempty"`

	// storage ip in
	StorageIPIn []string `json:"storage_ip_in,omitempty"`

	// storage ip lt
	StorageIPLt *string `json:"storage_ip_lt,omitempty"`

	// storage ip lte
	StorageIPLte *string `json:"storage_ip_lte,omitempty"`

	// storage ip not
	StorageIPNot *string `json:"storage_ip_not,omitempty"`

	// storage ip not contains
	StorageIPNotContains *string `json:"storage_ip_not_contains,omitempty"`

	// storage ip not ends with
	StorageIPNotEndsWith *string `json:"storage_ip_not_ends_with,omitempty"`

	// storage ip not in
	StorageIPNotIn []string `json:"storage_ip_not_in,omitempty"`

	// storage ip not starts with
	StorageIPNotStartsWith *string `json:"storage_ip_not_starts_with,omitempty"`

	// storage ip starts with
	StorageIPStartsWith *string `json:"storage_ip_starts_with,omitempty"`

	// type
	Type *ApplicationType `json:"type,omitempty"`

	// type in
	TypeIn []ApplicationType `json:"type_in,omitempty"`

	// type not
	TypeNot *ApplicationType `json:"type_not,omitempty"`

	// type not in
	TypeNotIn []ApplicationType `json:"type_not_in,omitempty"`

	// update time
	UpdateTime *string `json:"update_time,omitempty"`

	// update time gt
	UpdateTimeGt *string `json:"update_time_gt,omitempty"`

	// update time gte
	UpdateTimeGte *string `json:"update_time_gte,omitempty"`

	// update time in
	UpdateTimeIn []string `json:"update_time_in,omitempty"`

	// update time lt
	UpdateTimeLt *string `json:"update_time_lt,omitempty"`

	// update time lte
	UpdateTimeLte *string `json:"update_time_lte,omitempty"`

	// update time not
	UpdateTimeNot *string `json:"update_time_not,omitempty"`

	// update time not in
	UpdateTimeNotIn []string `json:"update_time_not_in,omitempty"`

	// vcpu
	Vcpu *int32 `json:"vcpu,omitempty"`

	// vcpu gt
	VcpuGt *int32 `json:"vcpu_gt,omitempty"`

	// vcpu gte
	VcpuGte *int32 `json:"vcpu_gte,omitempty"`

	// vcpu in
	VcpuIn []int32 `json:"vcpu_in,omitempty"`

	// vcpu lt
	VcpuLt *int32 `json:"vcpu_lt,omitempty"`

	// vcpu lte
	VcpuLte *int32 `json:"vcpu_lte,omitempty"`

	// vcpu not
	VcpuNot *int32 `json:"vcpu_not,omitempty"`

	// vcpu not in
	VcpuNotIn []int32 `json:"vcpu_not_in,omitempty"`

	// version
	Version *string `json:"version,omitempty"`

	// version contains
	VersionContains *string `json:"version_contains,omitempty"`

	// version ends with
	VersionEndsWith *string `json:"version_ends_with,omitempty"`

	// version gt
	VersionGt *string `json:"version_gt,omitempty"`

	// version gte
	VersionGte *string `json:"version_gte,omitempty"`

	// version in
	VersionIn []string `json:"version_in,omitempty"`

	// version lt
	VersionLt *string `json:"version_lt,omitempty"`

	// version lte
	VersionLte *string `json:"version_lte,omitempty"`

	// version not
	VersionNot *string `json:"version_not,omitempty"`

	// version not contains
	VersionNotContains *string `json:"version_not_contains,omitempty"`

	// version not ends with
	VersionNotEndsWith *string `json:"version_not_ends_with,omitempty"`

	// version not in
	VersionNotIn []string `json:"version_not_in,omitempty"`

	// version not starts with
	VersionNotStartsWith *string `json:"version_not_starts_with,omitempty"`

	// version starts with
	VersionStartsWith *string `json:"version_starts_with,omitempty"`

	// vm
	VM *VMWhereInput `json:"vm,omitempty"`

	// volume size
	VolumeSize *float64 `json:"volume_size,omitempty"`

	// volume size gt
	VolumeSizeGt *float64 `json:"volume_size_gt,omitempty"`

	// volume size gte
	VolumeSizeGte *float64 `json:"volume_size_gte,omitempty"`

	// volume size in
	VolumeSizeIn []float64 `json:"volume_size_in,omitempty"`

	// volume size lt
	VolumeSizeLt *float64 `json:"volume_size_lt,omitempty"`

	// volume size lte
	VolumeSizeLte *float64 `json:"volume_size_lte,omitempty"`

	// volume size not
	VolumeSizeNot *float64 `json:"volume_size_not,omitempty"`

	// volume size not in
	VolumeSizeNotIn []float64 `json:"volume_size_not_in,omitempty"`
}

// Validate validates this application where input
func (m *ApplicationWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationWhereInput) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) validateStateIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StateIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StateIn); i++ {

		if err := m.StateIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ApplicationWhereInput) validateStateNot(formats strfmt.Registry) error {
	if swag.IsZero(m.StateNot) { // not required
		return nil
	}

	if m.StateNot != nil {
		if err := m.StateNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state_not")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) validateStateNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StateNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StateNotIn); i++ {

		if err := m.StateNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ApplicationWhereInput) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) validateTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ApplicationWhereInput) validateTypeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNot) { // not required
		return nil
	}

	if m.TypeNot != nil {
		if err := m.TypeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) validateTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ApplicationWhereInput) validateVM(formats strfmt.Registry) error {
	if swag.IsZero(m.VM) { // not required
		return nil
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application where input based on the context it is used
func (m *ApplicationWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStateNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateStateIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StateIn); i++ {

		if err := m.StateIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateStateNot(ctx context.Context, formats strfmt.Registry) error {

	if m.StateNot != nil {
		if err := m.StateNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state_not")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateStateNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StateNotIn); i++ {

		if err := m.StateNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateTypeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.TypeNot != nil {
		if err := m.TypeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ApplicationWhereInput) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationWhereInput) UnmarshalBinary(b []byte) error {
	var res ApplicationWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
