// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetBackupRestoreExecutionsRequestBody get backup restore executions request body
// Example: {"after":"backupRestoreExecutions-id-string","before":"backupRestoreExecutions-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"BackupRestoreExecutionWhereInput[]","NOT":"BackupRestoreExecutionWhereInput[]","OR":"BackupRestoreExecutionWhereInput[]","backup_restore_point":"BackupRestorePointWhereInput","duration":0,"duration_gt":0,"duration_gte":0,"duration_in":[0],"duration_lt":0,"duration_lte":0,"duration_not":0,"duration_not_in":[0],"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"executed_at":"string","executed_at_gt":"string","executed_at_gte":"string","executed_at_in":["string"],"executed_at_lt":"string","executed_at_lte":"string","executed_at_not":"string","executed_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","mode":"INPLACE","mode_in":["INPLACE"],"mode_not":"INPLACE","mode_not_in":["INPLACE"],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","read_bytes":0,"read_bytes_gt":0,"read_bytes_gte":0,"read_bytes_in":[0],"read_bytes_lt":0,"read_bytes_lte":0,"read_bytes_not":0,"read_bytes_not_in":[0],"rebuild_name":"string","rebuild_name_contains":"string","rebuild_name_ends_with":"string","rebuild_name_gt":"string","rebuild_name_gte":"string","rebuild_name_in":["string"],"rebuild_name_lt":"string","rebuild_name_lte":"string","rebuild_name_not":"string","rebuild_name_not_contains":"string","rebuild_name_not_ends_with":"string","rebuild_name_not_in":["string"],"rebuild_name_not_starts_with":"string","rebuild_name_starts_with":"string","rebuild_target_cluster":"string","rebuild_target_cluster_contains":"string","rebuild_target_cluster_ends_with":"string","rebuild_target_cluster_gt":"string","rebuild_target_cluster_gte":"string","rebuild_target_cluster_in":["string"],"rebuild_target_cluster_lt":"string","rebuild_target_cluster_lte":"string","rebuild_target_cluster_not":"string","rebuild_target_cluster_not_contains":"string","rebuild_target_cluster_not_ends_with":"string","rebuild_target_cluster_not_in":["string"],"rebuild_target_cluster_not_starts_with":"string","rebuild_target_cluster_starts_with":"string","rebuild_target_host":"string","rebuild_target_host_contains":"string","rebuild_target_host_ends_with":"string","rebuild_target_host_gt":"string","rebuild_target_host_gte":"string","rebuild_target_host_in":["string"],"rebuild_target_host_lt":"string","rebuild_target_host_lte":"string","rebuild_target_host_not":"string","rebuild_target_host_not_contains":"string","rebuild_target_host_not_ends_with":"string","rebuild_target_host_not_in":["string"],"rebuild_target_host_not_starts_with":"string","rebuild_target_host_starts_with":"string","resource_version_gt":0,"resource_version_gte":0,"resource_version_in":[0],"resource_version_lt":0,"resource_version_lte":0,"resource_version_not":0,"resource_version_not_in":[0],"startup_after_restore":false,"startup_after_restore_not":false,"status":"ABORTED","status_in":["ABORTED"],"status_not":"ABORTED","status_not_in":["ABORTED"],"total_bytes":0,"total_bytes_gt":0,"total_bytes_gte":0,"total_bytes_in":[0],"total_bytes_lt":0,"total_bytes_lte":0,"total_bytes_not":0,"total_bytes_not_in":[0]}}
//
// swagger:model GetBackupRestoreExecutionsRequestBody
type GetBackupRestoreExecutionsRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy struct {
		BackupRestoreExecutionOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		BackupRestoreExecutionWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get backup restore executions request body
func (m *GetBackupRestoreExecutionsRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBackupRestoreExecutionsRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetBackupRestoreExecutionsRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get backup restore executions request body based on the context it is used
func (m *GetBackupRestoreExecutionsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBackupRestoreExecutionsRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetBackupRestoreExecutionsRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetBackupRestoreExecutionsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupRestoreExecutionsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetBackupRestoreExecutionsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
