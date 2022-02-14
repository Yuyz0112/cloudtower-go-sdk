// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetBackupRestorePointsRequestBody get backup restore points request body
// Example: {"after":"backupRestorePoints-id-string","before":"backupRestorePoints-id-string","first":0,"last":0,"orderBy":"cluster_local_id_ASC","skip":0,"where":{"AND":"BackupRestorePointWhereInput[]","NOT":"BackupRestorePointWhereInput[]","OR":"BackupRestorePointWhereInput[]","backup_plan":"BackupPlanWhereInput","backup_service":"BackupServiceWhereInput","backup_store_repository":"BackupStoreRepositoryWhereInput","backup_target_execution":"BackupTargetExecutionWhereInput","cluster_local_id":"string","cluster_local_id_contains":"string","cluster_local_id_ends_with":"string","cluster_local_id_gt":"string","cluster_local_id_gte":"string","cluster_local_id_in":["string"],"cluster_local_id_lt":"string","cluster_local_id_lte":"string","cluster_local_id_not":"string","cluster_local_id_not_contains":"string","cluster_local_id_not_ends_with":"string","cluster_local_id_not_in":["string"],"cluster_local_id_not_starts_with":"string","cluster_local_id_starts_with":"string","compression_ratio":0,"compression_ratio_gt":0,"compression_ratio_gte":0,"compression_ratio_in":[0],"compression_ratio_lt":0,"compression_ratio_lte":0,"compression_ratio_not":0,"compression_ratio_not_in":[0],"creation":"AUTO","creation_in":["AUTO"],"creation_not":"AUTO","creation_not_in":["AUTO"],"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","logical_size":0,"logical_size_gt":0,"logical_size_gte":0,"logical_size_in":[0],"logical_size_lt":0,"logical_size_lte":0,"logical_size_not":0,"logical_size_not_in":[0],"parent_restore_point":"string","parent_restore_point_contains":"string","parent_restore_point_ends_with":"string","parent_restore_point_gt":"string","parent_restore_point_gte":"string","parent_restore_point_in":["string"],"parent_restore_point_lt":"string","parent_restore_point_lte":"string","parent_restore_point_not":"string","parent_restore_point_not_contains":"string","parent_restore_point_not_ends_with":"string","parent_restore_point_not_in":["string"],"parent_restore_point_not_starts_with":"string","parent_restore_point_starts_with":"string","physical_size":0,"physical_size_gt":0,"physical_size_gte":0,"physical_size_in":[0],"physical_size_lt":0,"physical_size_lte":0,"physical_size_not":0,"physical_size_not_in":[0],"resource_version_gt":0,"resource_version_gte":0,"resource_version_in":[0],"resource_version_lt":0,"resource_version_lte":0,"resource_version_not":0,"resource_version_not_in":[0],"size":0,"size_gt":0,"size_gte":0,"size_in":[0],"size_lt":0,"size_lte":0,"size_not":0,"size_not_in":[0],"slice":"string","slice_contains":"string","slice_ends_with":"string","slice_gt":"string","slice_gte":"string","slice_in":["string"],"slice_lt":"string","slice_lte":"string","slice_not":"string","slice_not_contains":"string","slice_not_ends_with":"string","slice_not_in":["string"],"slice_not_starts_with":"string","slice_starts_with":"string","type":"FULL","type_in":["FULL"],"type_not":"FULL","type_not_in":["FULL"],"valid_capacity":0,"valid_capacity_gt":0,"valid_capacity_gte":0,"valid_capacity_in":[0],"valid_capacity_lt":0,"valid_capacity_lte":0,"valid_capacity_not":0,"valid_capacity_not_in":[0],"valid_size":0,"valid_size_gt":0,"valid_size_gte":0,"valid_size_in":[0],"valid_size_lt":0,"valid_size_lte":0,"valid_size_not":0,"valid_size_not_in":[0],"vm":"VmWhereInput","vm_local_id":"string","vm_local_id_contains":"string","vm_local_id_ends_with":"string","vm_local_id_gt":"string","vm_local_id_gte":"string","vm_local_id_in":["string"],"vm_local_id_lt":"string","vm_local_id_lte":"string","vm_local_id_not":"string","vm_local_id_not_contains":"string","vm_local_id_not_ends_with":"string","vm_local_id_not_in":["string"],"vm_local_id_not_starts_with":"string","vm_local_id_starts_with":"string","vm_name":"string","vm_name_contains":"string","vm_name_ends_with":"string","vm_name_gt":"string","vm_name_gte":"string","vm_name_in":["string"],"vm_name_lt":"string","vm_name_lte":"string","vm_name_not":"string","vm_name_not_contains":"string","vm_name_not_ends_with":"string","vm_name_not_in":["string"],"vm_name_not_starts_with":"string","vm_name_starts_with":"string"}}
//
// swagger:model GetBackupRestorePointsRequestBody
type GetBackupRestorePointsRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy interface{} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where interface{} `json:"where,omitempty"`
}

// Validate validates this get backup restore points request body
func (m *GetBackupRestorePointsRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get backup restore points request body based on context it is used
func (m *GetBackupRestorePointsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBackupRestorePointsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupRestorePointsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetBackupRestorePointsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
