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

// SnapshotGroupWhereInput snapshot group where input
// Example: {"AND":"SnapshotGroupWhereInput[]","NOT":"SnapshotGroupWhereInput[]","OR":"SnapshotGroupWhereInput[]","cluster":"ClusterWhereInput","deleted":false,"deleted_not":false,"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"estimated_recycling_time":"string","estimated_recycling_time_gt":"string","estimated_recycling_time_gte":"string","estimated_recycling_time_in":["string"],"estimated_recycling_time_lt":"string","estimated_recycling_time_lte":"string","estimated_recycling_time_not":"string","estimated_recycling_time_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"keep":false,"keep_not":false,"local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","logical_size_bytes":0,"logical_size_bytes_gt":0,"logical_size_bytes_gte":0,"logical_size_bytes_in":[0],"logical_size_bytes_lt":0,"logical_size_bytes_lte":0,"logical_size_bytes_not":0,"logical_size_bytes_not_in":[0],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","object_num":0,"object_num_gt":0,"object_num_gte":0,"object_num_in":[0],"object_num_lt":0,"object_num_lte":0,"object_num_not":0,"object_num_not_in":[0],"snapshotPlanTask":"SnapshotPlanTaskWhereInput","vm_snapshots_every":"VmSnapshotWhereInput","vm_snapshots_none":"VmSnapshotWhereInput","vm_snapshots_some":"VmSnapshotWhereInput"}
//
// swagger:model SnapshotGroupWhereInput
type SnapshotGroupWhereInput struct {

	// a n d
	AND []*SnapshotGroupWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*SnapshotGroupWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*SnapshotGroupWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// deleted
	Deleted *bool `json:"deleted,omitempty"`

	// deleted not
	DeletedNot *bool `json:"deleted_not,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// estimated recycling time
	EstimatedRecyclingTime *string `json:"estimated_recycling_time,omitempty"`

	// estimated recycling time gt
	EstimatedRecyclingTimeGt *string `json:"estimated_recycling_time_gt,omitempty"`

	// estimated recycling time gte
	EstimatedRecyclingTimeGte *string `json:"estimated_recycling_time_gte,omitempty"`

	// estimated recycling time in
	EstimatedRecyclingTimeIn []string `json:"estimated_recycling_time_in,omitempty"`

	// estimated recycling time lt
	EstimatedRecyclingTimeLt *string `json:"estimated_recycling_time_lt,omitempty"`

	// estimated recycling time lte
	EstimatedRecyclingTimeLte *string `json:"estimated_recycling_time_lte,omitempty"`

	// estimated recycling time not
	EstimatedRecyclingTimeNot *string `json:"estimated_recycling_time_not,omitempty"`

	// estimated recycling time not in
	EstimatedRecyclingTimeNotIn []string `json:"estimated_recycling_time_not_in,omitempty"`

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

	// internal
	Internal *bool `json:"internal,omitempty"`

	// internal not
	InternalNot *bool `json:"internal_not,omitempty"`

	// keep
	Keep *bool `json:"keep,omitempty"`

	// keep not
	KeepNot *bool `json:"keep_not,omitempty"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local created at gt
	LocalCreatedAtGt *string `json:"local_created_at_gt,omitempty"`

	// local created at gte
	LocalCreatedAtGte *string `json:"local_created_at_gte,omitempty"`

	// local created at in
	LocalCreatedAtIn []string `json:"local_created_at_in,omitempty"`

	// local created at lt
	LocalCreatedAtLt *string `json:"local_created_at_lt,omitempty"`

	// local created at lte
	LocalCreatedAtLte *string `json:"local_created_at_lte,omitempty"`

	// local created at not
	LocalCreatedAtNot *string `json:"local_created_at_not,omitempty"`

	// local created at not in
	LocalCreatedAtNotIn []string `json:"local_created_at_not_in,omitempty"`

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

	// logical size bytes
	LogicalSizeBytes *float64 `json:"logical_size_bytes,omitempty"`

	// logical size bytes gt
	LogicalSizeBytesGt *float64 `json:"logical_size_bytes_gt,omitempty"`

	// logical size bytes gte
	LogicalSizeBytesGte *float64 `json:"logical_size_bytes_gte,omitempty"`

	// logical size bytes in
	LogicalSizeBytesIn []float64 `json:"logical_size_bytes_in,omitempty"`

	// logical size bytes lt
	LogicalSizeBytesLt *float64 `json:"logical_size_bytes_lt,omitempty"`

	// logical size bytes lte
	LogicalSizeBytesLte *float64 `json:"logical_size_bytes_lte,omitempty"`

	// logical size bytes not
	LogicalSizeBytesNot *float64 `json:"logical_size_bytes_not,omitempty"`

	// logical size bytes not in
	LogicalSizeBytesNotIn []float64 `json:"logical_size_bytes_not_in,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// object num
	ObjectNum *int32 `json:"object_num,omitempty"`

	// object num gt
	ObjectNumGt *int32 `json:"object_num_gt,omitempty"`

	// object num gte
	ObjectNumGte *int32 `json:"object_num_gte,omitempty"`

	// object num in
	ObjectNumIn []int32 `json:"object_num_in,omitempty"`

	// object num lt
	ObjectNumLt *int32 `json:"object_num_lt,omitempty"`

	// object num lte
	ObjectNumLte *int32 `json:"object_num_lte,omitempty"`

	// object num not
	ObjectNumNot *int32 `json:"object_num_not,omitempty"`

	// object num not in
	ObjectNumNotIn []int32 `json:"object_num_not_in,omitempty"`

	// snapshot plan task
	SnapshotPlanTask *SnapshotPlanTaskWhereInput `json:"snapshotPlanTask,omitempty"`

	// vm snapshots every
	VMSnapshotsEvery *VMSnapshotWhereInput `json:"vm_snapshots_every,omitempty"`

	// vm snapshots none
	VMSnapshotsNone *VMSnapshotWhereInput `json:"vm_snapshots_none,omitempty"`

	// vm snapshots some
	VMSnapshotsSome *VMSnapshotWhereInput `json:"vm_snapshots_some,omitempty"`
}

// Validate validates this snapshot group where input
func (m *SnapshotGroupWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotPlanTask(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshotsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshotsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshotsSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotGroupWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *SnapshotGroupWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *SnapshotGroupWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *SnapshotGroupWhereInput) validateCluster(formats strfmt.Registry) error {
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

func (m *SnapshotGroupWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotGroupWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotGroupWhereInput) validateSnapshotPlanTask(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotPlanTask) { // not required
		return nil
	}

	if m.SnapshotPlanTask != nil {
		if err := m.SnapshotPlanTask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotPlanTask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotPlanTask")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) validateVMSnapshotsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshotsEvery) { // not required
		return nil
	}

	if m.VMSnapshotsEvery != nil {
		if err := m.VMSnapshotsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_snapshots_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_snapshots_every")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) validateVMSnapshotsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshotsNone) { // not required
		return nil
	}

	if m.VMSnapshotsNone != nil {
		if err := m.VMSnapshotsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_snapshots_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_snapshots_none")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) validateVMSnapshotsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshotsSome) { // not required
		return nil
	}

	if m.VMSnapshotsSome != nil {
		if err := m.VMSnapshotsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_snapshots_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_snapshots_some")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapshot group where input based on the context it is used
func (m *SnapshotGroupWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotPlanTask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshotsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshotsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshotsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotGroupWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapshotGroupWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapshotGroupWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapshotGroupWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SnapshotGroupWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotGroupWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnapshotGroupWhereInput) contextValidateSnapshotPlanTask(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotPlanTask != nil {
		if err := m.SnapshotPlanTask.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotPlanTask")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotPlanTask")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) contextValidateVMSnapshotsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.VMSnapshotsEvery != nil {
		if err := m.VMSnapshotsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_snapshots_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_snapshots_every")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) contextValidateVMSnapshotsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.VMSnapshotsNone != nil {
		if err := m.VMSnapshotsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_snapshots_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_snapshots_none")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotGroupWhereInput) contextValidateVMSnapshotsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.VMSnapshotsSome != nil {
		if err := m.VMSnapshotsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_snapshots_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_snapshots_some")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotGroupWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotGroupWhereInput) UnmarshalBinary(b []byte) error {
	var res SnapshotGroupWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
