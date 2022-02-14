// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetBackupPlansConnectionRequestBody get backup plans connection request body
// Example: {"after":"backupPlansConnection-id-string","before":"backupPlansConnection-id-string","first":0,"last":0,"orderBy":"backup_delay_option_ASC","skip":0,"where":{"AND":"BackupPlanWhereInput[]","NOT":"BackupPlanWhereInput[]","OR":"BackupPlanWhereInput[]","backup_delay_option":"KEEP_GO_ON","backup_delay_option_in":["KEEP_GO_ON"],"backup_delay_option_not":"KEEP_GO_ON","backup_delay_option_not_in":["KEEP_GO_ON"],"backup_plan_executions_every":"BackupPlanExecutionWhereInput","backup_plan_executions_none":"BackupPlanExecutionWhereInput","backup_plan_executions_some":"BackupPlanExecutionWhereInput","backup_restore_point_count":0,"backup_restore_point_count_gt":0,"backup_restore_point_count_gte":0,"backup_restore_point_count_in":[0],"backup_restore_point_count_lt":0,"backup_restore_point_count_lte":0,"backup_restore_point_count_not":0,"backup_restore_point_count_not_in":[0],"backup_restore_points_every":"BackupRestorePointWhereInput","backup_restore_points_none":"BackupRestorePointWhereInput","backup_restore_points_some":"BackupRestorePointWhereInput","backup_service":"BackupServiceWhereInput","backup_store_repository":"BackupStoreRepositoryWhereInput","backup_total_size":0,"backup_total_size_gt":0,"backup_total_size_gte":0,"backup_total_size_in":[0],"backup_total_size_lt":0,"backup_total_size_lte":0,"backup_total_size_not":0,"backup_total_size_not_in":[0],"compression":false,"compression_not":false,"delete_strategy":"DELETE_RESTORE_POINT","delete_strategy_in":["DELETE_RESTORE_POINT"],"delete_strategy_not":"DELETE_RESTORE_POINT","delete_strategy_not_in":["DELETE_RESTORE_POINT"],"description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","enable_window":false,"enable_window_not":false,"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"full_interval":0,"full_interval_gt":0,"full_interval_gte":0,"full_interval_in":[0],"full_interval_lt":0,"full_interval_lte":0,"full_interval_not":0,"full_interval_not_in":[0],"full_period":"DAILY","full_period_in":["DAILY"],"full_period_not":"DAILY","full_period_not_in":["DAILY"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","incremental_period":"DAILY","incremental_period_in":["DAILY"],"incremental_period_not":"DAILY","incremental_period_not_in":["DAILY"],"keep_policy":"COUNT","keep_policy_in":["COUNT"],"keep_policy_not":"COUNT","keep_policy_not_in":["COUNT"],"keep_policy_value":0,"keep_policy_value_gt":0,"keep_policy_value_gte":0,"keep_policy_value_in":[0],"keep_policy_value_lt":0,"keep_policy_value_lte":0,"keep_policy_value_not":0,"keep_policy_value_not_in":[0],"last_execute_status":"FAILED","last_execute_status_in":["FAILED"],"last_execute_status_message":"string","last_execute_status_message_contains":"string","last_execute_status_message_ends_with":"string","last_execute_status_message_gt":"string","last_execute_status_message_gte":"string","last_execute_status_message_in":["string"],"last_execute_status_message_lt":"string","last_execute_status_message_lte":"string","last_execute_status_message_not":"string","last_execute_status_message_not_contains":"string","last_execute_status_message_not_ends_with":"string","last_execute_status_message_not_in":["string"],"last_execute_status_message_not_starts_with":"string","last_execute_status_message_starts_with":"string","last_execute_status_not":"FAILED","last_execute_status_not_in":["FAILED"],"last_execute_success_job_count":0,"last_execute_success_job_count_gt":0,"last_execute_success_job_count_gte":0,"last_execute_success_job_count_in":[0],"last_execute_success_job_count_lt":0,"last_execute_success_job_count_lte":0,"last_execute_success_job_count_not":0,"last_execute_success_job_count_not_in":[0],"last_execute_total_job_count":0,"last_execute_total_job_count_gt":0,"last_execute_total_job_count_gte":0,"last_execute_total_job_count_in":[0],"last_execute_total_job_count_lt":0,"last_execute_total_job_count_lte":0,"last_execute_total_job_count_not":0,"last_execute_total_job_count_not_in":[0],"last_executed_at":"string","last_executed_at_gt":"string","last_executed_at_gte":"string","last_executed_at_in":["string"],"last_executed_at_lt":"string","last_executed_at_lte":"string","last_executed_at_not":"string","last_executed_at_not_in":["string"],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","namespace":"string","namespace_contains":"string","namespace_ends_with":"string","namespace_gt":"string","namespace_gte":"string","namespace_in":["string"],"namespace_lt":"string","namespace_lte":"string","namespace_not":"string","namespace_not_contains":"string","namespace_not_ends_with":"string","namespace_not_in":["string"],"namespace_not_starts_with":"string","namespace_starts_with":"string","next_execute_time":"string","next_execute_time_gt":"string","next_execute_time_gte":"string","next_execute_time_in":["string"],"next_execute_time_lt":"string","next_execute_time_lte":"string","next_execute_time_not":"string","next_execute_time_not_in":["string"],"resource_version_gt":0,"resource_version_gte":0,"resource_version_in":[0],"resource_version_lt":0,"resource_version_lte":0,"resource_version_not":0,"resource_version_not_in":[0],"status":"PAUSED","status_in":["PAUSED"],"status_not":"PAUSED","status_not_in":["PAUSED"],"vms_every":"VmWhereInput","vms_none":"VmWhereInput","vms_some":"VmWhereInput","window_end":"string","window_end_contains":"string","window_end_ends_with":"string","window_end_gt":"string","window_end_gte":"string","window_end_in":["string"],"window_end_lt":"string","window_end_lte":"string","window_end_not":"string","window_end_not_contains":"string","window_end_not_ends_with":"string","window_end_not_in":["string"],"window_end_not_starts_with":"string","window_end_starts_with":"string","window_start":"string","window_start_contains":"string","window_start_ends_with":"string","window_start_gt":"string","window_start_gte":"string","window_start_in":["string"],"window_start_lt":"string","window_start_lte":"string","window_start_not":"string","window_start_not_contains":"string","window_start_not_ends_with":"string","window_start_not_in":["string"],"window_start_not_starts_with":"string","window_start_starts_with":"string"}}
//
// swagger:model GetBackupPlansConnectionRequestBody
type GetBackupPlansConnectionRequestBody struct {

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

// Validate validates this get backup plans connection request body
func (m *GetBackupPlansConnectionRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get backup plans connection request body based on context it is used
func (m *GetBackupPlansConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBackupPlansConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupPlansConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetBackupPlansConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
