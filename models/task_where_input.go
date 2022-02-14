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

// TaskWhereInput task where input
// Example: {"AND":"TaskWhereInput[]","NOT":"TaskWhereInput[]","OR":"TaskWhereInput[]","cluster":"ClusterWhereInput","description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","error_code":"string","error_code_contains":"string","error_code_ends_with":"string","error_code_gt":"string","error_code_gte":"string","error_code_in":["string"],"error_code_lt":"string","error_code_lte":"string","error_code_not":"string","error_code_not_contains":"string","error_code_not_ends_with":"string","error_code_not_in":["string"],"error_code_not_starts_with":"string","error_code_starts_with":"string","error_message":"string","error_message_contains":"string","error_message_ends_with":"string","error_message_gt":"string","error_message_gte":"string","error_message_in":["string"],"error_message_lt":"string","error_message_lte":"string","error_message_not":"string","error_message_not_contains":"string","error_message_not_ends_with":"string","error_message_not_in":["string"],"error_message_not_starts_with":"string","error_message_starts_with":"string","finished_at":"string","finished_at_gt":"string","finished_at_gte":"string","finished_at_in":["string"],"finished_at_lt":"string","finished_at_lte":"string","finished_at_not":"string","finished_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"progress":0,"progress_gt":0,"progress_gte":0,"progress_in":[0],"progress_lt":0,"progress_lte":0,"progress_not":0,"progress_not_in":[0],"resource_id":"string","resource_id_contains":"string","resource_id_ends_with":"string","resource_id_gt":"string","resource_id_gte":"string","resource_id_in":["string"],"resource_id_lt":"string","resource_id_lte":"string","resource_id_not":"string","resource_id_not_contains":"string","resource_id_not_ends_with":"string","resource_id_not_in":["string"],"resource_id_not_starts_with":"string","resource_id_starts_with":"string","resource_mutation":"string","resource_mutation_contains":"string","resource_mutation_ends_with":"string","resource_mutation_gt":"string","resource_mutation_gte":"string","resource_mutation_in":["string"],"resource_mutation_lt":"string","resource_mutation_lte":"string","resource_mutation_not":"string","resource_mutation_not_contains":"string","resource_mutation_not_ends_with":"string","resource_mutation_not_in":["string"],"resource_mutation_not_starts_with":"string","resource_mutation_starts_with":"string","resource_rollback_error":"string","resource_rollback_error_contains":"string","resource_rollback_error_ends_with":"string","resource_rollback_error_gt":"string","resource_rollback_error_gte":"string","resource_rollback_error_in":["string"],"resource_rollback_error_lt":"string","resource_rollback_error_lte":"string","resource_rollback_error_not":"string","resource_rollback_error_not_contains":"string","resource_rollback_error_not_ends_with":"string","resource_rollback_error_not_in":["string"],"resource_rollback_error_not_starts_with":"string","resource_rollback_error_starts_with":"string","resource_rollback_retry_count":0,"resource_rollback_retry_count_gt":0,"resource_rollback_retry_count_gte":0,"resource_rollback_retry_count_in":[0],"resource_rollback_retry_count_lt":0,"resource_rollback_retry_count_lte":0,"resource_rollback_retry_count_not":0,"resource_rollback_retry_count_not_in":[0],"resource_rollbacked":false,"resource_rollbacked_not":false,"resource_type":"string","resource_type_contains":"string","resource_type_ends_with":"string","resource_type_gt":"string","resource_type_gte":"string","resource_type_in":["string"],"resource_type_lt":"string","resource_type_lte":"string","resource_type_not":"string","resource_type_not_contains":"string","resource_type_not_ends_with":"string","resource_type_not_in":["string"],"resource_type_not_starts_with":"string","resource_type_starts_with":"string","snapshot":"string","snapshot_contains":"string","snapshot_ends_with":"string","snapshot_gt":"string","snapshot_gte":"string","snapshot_in":["string"],"snapshot_lt":"string","snapshot_lte":"string","snapshot_not":"string","snapshot_not_contains":"string","snapshot_not_ends_with":"string","snapshot_not_in":["string"],"snapshot_not_starts_with":"string","snapshot_starts_with":"string","started_at":"string","started_at_gt":"string","started_at_gte":"string","started_at_in":["string"],"started_at_lt":"string","started_at_lte":"string","started_at_not":"string","started_at_not_in":["string"],"status":"EXECUTING","status_in":["EXECUTING"],"status_not":"EXECUTING","status_not_in":["EXECUTING"],"user":"UserWhereInput"}
//
// swagger:model TaskWhereInput
type TaskWhereInput struct {

	// a n d
	AND []*TaskWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*TaskWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*TaskWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// description contains
	DescriptionContains *string `json:"description_contains,omitempty"`

	// description ends with
	DescriptionEndsWith *string `json:"description_ends_with,omitempty"`

	// description gt
	DescriptionGt *string `json:"description_gt,omitempty"`

	// description gte
	DescriptionGte *string `json:"description_gte,omitempty"`

	// description in
	DescriptionIn []string `json:"description_in,omitempty"`

	// description lt
	DescriptionLt *string `json:"description_lt,omitempty"`

	// description lte
	DescriptionLte *string `json:"description_lte,omitempty"`

	// description not
	DescriptionNot *string `json:"description_not,omitempty"`

	// description not contains
	DescriptionNotContains *string `json:"description_not_contains,omitempty"`

	// description not ends with
	DescriptionNotEndsWith *string `json:"description_not_ends_with,omitempty"`

	// description not in
	DescriptionNotIn []string `json:"description_not_in,omitempty"`

	// description not starts with
	DescriptionNotStartsWith *string `json:"description_not_starts_with,omitempty"`

	// description starts with
	DescriptionStartsWith *string `json:"description_starts_with,omitempty"`

	// error code
	ErrorCode *string `json:"error_code,omitempty"`

	// error code contains
	ErrorCodeContains *string `json:"error_code_contains,omitempty"`

	// error code ends with
	ErrorCodeEndsWith *string `json:"error_code_ends_with,omitempty"`

	// error code gt
	ErrorCodeGt *string `json:"error_code_gt,omitempty"`

	// error code gte
	ErrorCodeGte *string `json:"error_code_gte,omitempty"`

	// error code in
	ErrorCodeIn []string `json:"error_code_in,omitempty"`

	// error code lt
	ErrorCodeLt *string `json:"error_code_lt,omitempty"`

	// error code lte
	ErrorCodeLte *string `json:"error_code_lte,omitempty"`

	// error code not
	ErrorCodeNot *string `json:"error_code_not,omitempty"`

	// error code not contains
	ErrorCodeNotContains *string `json:"error_code_not_contains,omitempty"`

	// error code not ends with
	ErrorCodeNotEndsWith *string `json:"error_code_not_ends_with,omitempty"`

	// error code not in
	ErrorCodeNotIn []string `json:"error_code_not_in,omitempty"`

	// error code not starts with
	ErrorCodeNotStartsWith *string `json:"error_code_not_starts_with,omitempty"`

	// error code starts with
	ErrorCodeStartsWith *string `json:"error_code_starts_with,omitempty"`

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

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// finished at gt
	FinishedAtGt *string `json:"finished_at_gt,omitempty"`

	// finished at gte
	FinishedAtGte *string `json:"finished_at_gte,omitempty"`

	// finished at in
	FinishedAtIn []string `json:"finished_at_in,omitempty"`

	// finished at lt
	FinishedAtLt *string `json:"finished_at_lt,omitempty"`

	// finished at lte
	FinishedAtLte *string `json:"finished_at_lte,omitempty"`

	// finished at not
	FinishedAtNot *string `json:"finished_at_not,omitempty"`

	// finished at not in
	FinishedAtNotIn []string `json:"finished_at_not_in,omitempty"`

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

	// progress
	Progress *float64 `json:"progress,omitempty"`

	// progress gt
	ProgressGt *float64 `json:"progress_gt,omitempty"`

	// progress gte
	ProgressGte *float64 `json:"progress_gte,omitempty"`

	// progress in
	ProgressIn []float64 `json:"progress_in,omitempty"`

	// progress lt
	ProgressLt *float64 `json:"progress_lt,omitempty"`

	// progress lte
	ProgressLte *float64 `json:"progress_lte,omitempty"`

	// progress not
	ProgressNot *float64 `json:"progress_not,omitempty"`

	// progress not in
	ProgressNotIn []float64 `json:"progress_not_in,omitempty"`

	// resource id
	ResourceID *string `json:"resource_id,omitempty"`

	// resource id contains
	ResourceIDContains *string `json:"resource_id_contains,omitempty"`

	// resource id ends with
	ResourceIDEndsWith *string `json:"resource_id_ends_with,omitempty"`

	// resource id gt
	ResourceIDGt *string `json:"resource_id_gt,omitempty"`

	// resource id gte
	ResourceIDGte *string `json:"resource_id_gte,omitempty"`

	// resource id in
	ResourceIDIn []string `json:"resource_id_in,omitempty"`

	// resource id lt
	ResourceIDLt *string `json:"resource_id_lt,omitempty"`

	// resource id lte
	ResourceIDLte *string `json:"resource_id_lte,omitempty"`

	// resource id not
	ResourceIDNot *string `json:"resource_id_not,omitempty"`

	// resource id not contains
	ResourceIDNotContains *string `json:"resource_id_not_contains,omitempty"`

	// resource id not ends with
	ResourceIDNotEndsWith *string `json:"resource_id_not_ends_with,omitempty"`

	// resource id not in
	ResourceIDNotIn []string `json:"resource_id_not_in,omitempty"`

	// resource id not starts with
	ResourceIDNotStartsWith *string `json:"resource_id_not_starts_with,omitempty"`

	// resource id starts with
	ResourceIDStartsWith *string `json:"resource_id_starts_with,omitempty"`

	// resource mutation
	ResourceMutation *string `json:"resource_mutation,omitempty"`

	// resource mutation contains
	ResourceMutationContains *string `json:"resource_mutation_contains,omitempty"`

	// resource mutation ends with
	ResourceMutationEndsWith *string `json:"resource_mutation_ends_with,omitempty"`

	// resource mutation gt
	ResourceMutationGt *string `json:"resource_mutation_gt,omitempty"`

	// resource mutation gte
	ResourceMutationGte *string `json:"resource_mutation_gte,omitempty"`

	// resource mutation in
	ResourceMutationIn []string `json:"resource_mutation_in,omitempty"`

	// resource mutation lt
	ResourceMutationLt *string `json:"resource_mutation_lt,omitempty"`

	// resource mutation lte
	ResourceMutationLte *string `json:"resource_mutation_lte,omitempty"`

	// resource mutation not
	ResourceMutationNot *string `json:"resource_mutation_not,omitempty"`

	// resource mutation not contains
	ResourceMutationNotContains *string `json:"resource_mutation_not_contains,omitempty"`

	// resource mutation not ends with
	ResourceMutationNotEndsWith *string `json:"resource_mutation_not_ends_with,omitempty"`

	// resource mutation not in
	ResourceMutationNotIn []string `json:"resource_mutation_not_in,omitempty"`

	// resource mutation not starts with
	ResourceMutationNotStartsWith *string `json:"resource_mutation_not_starts_with,omitempty"`

	// resource mutation starts with
	ResourceMutationStartsWith *string `json:"resource_mutation_starts_with,omitempty"`

	// resource rollback error
	ResourceRollbackError *string `json:"resource_rollback_error,omitempty"`

	// resource rollback error contains
	ResourceRollbackErrorContains *string `json:"resource_rollback_error_contains,omitempty"`

	// resource rollback error ends with
	ResourceRollbackErrorEndsWith *string `json:"resource_rollback_error_ends_with,omitempty"`

	// resource rollback error gt
	ResourceRollbackErrorGt *string `json:"resource_rollback_error_gt,omitempty"`

	// resource rollback error gte
	ResourceRollbackErrorGte *string `json:"resource_rollback_error_gte,omitempty"`

	// resource rollback error in
	ResourceRollbackErrorIn []string `json:"resource_rollback_error_in,omitempty"`

	// resource rollback error lt
	ResourceRollbackErrorLt *string `json:"resource_rollback_error_lt,omitempty"`

	// resource rollback error lte
	ResourceRollbackErrorLte *string `json:"resource_rollback_error_lte,omitempty"`

	// resource rollback error not
	ResourceRollbackErrorNot *string `json:"resource_rollback_error_not,omitempty"`

	// resource rollback error not contains
	ResourceRollbackErrorNotContains *string `json:"resource_rollback_error_not_contains,omitempty"`

	// resource rollback error not ends with
	ResourceRollbackErrorNotEndsWith *string `json:"resource_rollback_error_not_ends_with,omitempty"`

	// resource rollback error not in
	ResourceRollbackErrorNotIn []string `json:"resource_rollback_error_not_in,omitempty"`

	// resource rollback error not starts with
	ResourceRollbackErrorNotStartsWith *string `json:"resource_rollback_error_not_starts_with,omitempty"`

	// resource rollback error starts with
	ResourceRollbackErrorStartsWith *string `json:"resource_rollback_error_starts_with,omitempty"`

	// resource rollback retry count
	ResourceRollbackRetryCount *int32 `json:"resource_rollback_retry_count,omitempty"`

	// resource rollback retry count gt
	ResourceRollbackRetryCountGt *int32 `json:"resource_rollback_retry_count_gt,omitempty"`

	// resource rollback retry count gte
	ResourceRollbackRetryCountGte *int32 `json:"resource_rollback_retry_count_gte,omitempty"`

	// resource rollback retry count in
	ResourceRollbackRetryCountIn []int32 `json:"resource_rollback_retry_count_in,omitempty"`

	// resource rollback retry count lt
	ResourceRollbackRetryCountLt *int32 `json:"resource_rollback_retry_count_lt,omitempty"`

	// resource rollback retry count lte
	ResourceRollbackRetryCountLte *int32 `json:"resource_rollback_retry_count_lte,omitempty"`

	// resource rollback retry count not
	ResourceRollbackRetryCountNot *int32 `json:"resource_rollback_retry_count_not,omitempty"`

	// resource rollback retry count not in
	ResourceRollbackRetryCountNotIn []int32 `json:"resource_rollback_retry_count_not_in,omitempty"`

	// resource rollbacked
	ResourceRollbacked *bool `json:"resource_rollbacked,omitempty"`

	// resource rollbacked not
	ResourceRollbackedNot *bool `json:"resource_rollbacked_not,omitempty"`

	// resource type
	ResourceType *string `json:"resource_type,omitempty"`

	// resource type contains
	ResourceTypeContains *string `json:"resource_type_contains,omitempty"`

	// resource type ends with
	ResourceTypeEndsWith *string `json:"resource_type_ends_with,omitempty"`

	// resource type gt
	ResourceTypeGt *string `json:"resource_type_gt,omitempty"`

	// resource type gte
	ResourceTypeGte *string `json:"resource_type_gte,omitempty"`

	// resource type in
	ResourceTypeIn []string `json:"resource_type_in,omitempty"`

	// resource type lt
	ResourceTypeLt *string `json:"resource_type_lt,omitempty"`

	// resource type lte
	ResourceTypeLte *string `json:"resource_type_lte,omitempty"`

	// resource type not
	ResourceTypeNot *string `json:"resource_type_not,omitempty"`

	// resource type not contains
	ResourceTypeNotContains *string `json:"resource_type_not_contains,omitempty"`

	// resource type not ends with
	ResourceTypeNotEndsWith *string `json:"resource_type_not_ends_with,omitempty"`

	// resource type not in
	ResourceTypeNotIn []string `json:"resource_type_not_in,omitempty"`

	// resource type not starts with
	ResourceTypeNotStartsWith *string `json:"resource_type_not_starts_with,omitempty"`

	// resource type starts with
	ResourceTypeStartsWith *string `json:"resource_type_starts_with,omitempty"`

	// snapshot
	Snapshot *string `json:"snapshot,omitempty"`

	// snapshot contains
	SnapshotContains *string `json:"snapshot_contains,omitempty"`

	// snapshot ends with
	SnapshotEndsWith *string `json:"snapshot_ends_with,omitempty"`

	// snapshot gt
	SnapshotGt *string `json:"snapshot_gt,omitempty"`

	// snapshot gte
	SnapshotGte *string `json:"snapshot_gte,omitempty"`

	// snapshot in
	SnapshotIn []string `json:"snapshot_in,omitempty"`

	// snapshot lt
	SnapshotLt *string `json:"snapshot_lt,omitempty"`

	// snapshot lte
	SnapshotLte *string `json:"snapshot_lte,omitempty"`

	// snapshot not
	SnapshotNot *string `json:"snapshot_not,omitempty"`

	// snapshot not contains
	SnapshotNotContains *string `json:"snapshot_not_contains,omitempty"`

	// snapshot not ends with
	SnapshotNotEndsWith *string `json:"snapshot_not_ends_with,omitempty"`

	// snapshot not in
	SnapshotNotIn []string `json:"snapshot_not_in,omitempty"`

	// snapshot not starts with
	SnapshotNotStartsWith *string `json:"snapshot_not_starts_with,omitempty"`

	// snapshot starts with
	SnapshotStartsWith *string `json:"snapshot_starts_with,omitempty"`

	// started at
	StartedAt *string `json:"started_at,omitempty"`

	// started at gt
	StartedAtGt *string `json:"started_at_gt,omitempty"`

	// started at gte
	StartedAtGte *string `json:"started_at_gte,omitempty"`

	// started at in
	StartedAtIn []string `json:"started_at_in,omitempty"`

	// started at lt
	StartedAtLt *string `json:"started_at_lt,omitempty"`

	// started at lte
	StartedAtLte *string `json:"started_at_lte,omitempty"`

	// started at not
	StartedAtNot *string `json:"started_at_not,omitempty"`

	// started at not in
	StartedAtNotIn []string `json:"started_at_not_in,omitempty"`

	// status
	Status *TaskStatus `json:"status,omitempty"`

	// status in
	StatusIn []TaskStatus `json:"status_in,omitempty"`

	// status not
	StatusNot *TaskStatus `json:"status_not,omitempty"`

	// status not in
	StatusNotIn []TaskStatus `json:"status_not_in,omitempty"`

	// user
	User *UserWhereInput `json:"user,omitempty"`
}

// Validate validates this task where input
func (m *TaskWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *TaskWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *TaskWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *TaskWhereInput) validateCluster(formats strfmt.Registry) error {
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

func (m *TaskWhereInput) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *TaskWhereInput) validateStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *TaskWhereInput) validateStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNot) { // not required
		return nil
	}

	if m.StatusNot != nil {
		if err := m.StatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not")
			}
			return err
		}
	}

	return nil
}

func (m *TaskWhereInput) validateStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *TaskWhereInput) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this task where input based on the context it is used
func (m *TaskWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TaskWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TaskWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TaskWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *TaskWhereInput) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *TaskWhereInput) contextValidateStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *TaskWhereInput) contextValidateStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusNot != nil {
		if err := m.StatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not")
			}
			return err
		}
	}

	return nil
}

func (m *TaskWhereInput) contextValidateStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *TaskWhereInput) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskWhereInput) UnmarshalBinary(b []byte) error {
	var res TaskWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
