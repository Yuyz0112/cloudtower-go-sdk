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

// BackupRestoreExecutionWhereInput backup restore execution where input
// Example: {"AND":"BackupRestoreExecutionWhereInput[]","NOT":"BackupRestoreExecutionWhereInput[]","OR":"BackupRestoreExecutionWhereInput[]","backup_restore_point":"BackupRestorePointWhereInput","duration":0,"duration_gt":0,"duration_gte":0,"duration_in":[0],"duration_lt":0,"duration_lte":0,"duration_not":0,"duration_not_in":[0],"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"executed_at":"string","executed_at_gt":"string","executed_at_gte":"string","executed_at_in":["string"],"executed_at_lt":"string","executed_at_lte":"string","executed_at_not":"string","executed_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","mode":"INPLACE","mode_in":["INPLACE"],"mode_not":"INPLACE","mode_not_in":["INPLACE"],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","read_bytes":0,"read_bytes_gt":0,"read_bytes_gte":0,"read_bytes_in":[0],"read_bytes_lt":0,"read_bytes_lte":0,"read_bytes_not":0,"read_bytes_not_in":[0],"rebuild_name":"string","rebuild_name_contains":"string","rebuild_name_ends_with":"string","rebuild_name_gt":"string","rebuild_name_gte":"string","rebuild_name_in":["string"],"rebuild_name_lt":"string","rebuild_name_lte":"string","rebuild_name_not":"string","rebuild_name_not_contains":"string","rebuild_name_not_ends_with":"string","rebuild_name_not_in":["string"],"rebuild_name_not_starts_with":"string","rebuild_name_starts_with":"string","rebuild_target_cluster":"string","rebuild_target_cluster_contains":"string","rebuild_target_cluster_ends_with":"string","rebuild_target_cluster_gt":"string","rebuild_target_cluster_gte":"string","rebuild_target_cluster_in":["string"],"rebuild_target_cluster_lt":"string","rebuild_target_cluster_lte":"string","rebuild_target_cluster_not":"string","rebuild_target_cluster_not_contains":"string","rebuild_target_cluster_not_ends_with":"string","rebuild_target_cluster_not_in":["string"],"rebuild_target_cluster_not_starts_with":"string","rebuild_target_cluster_starts_with":"string","rebuild_target_host":"string","rebuild_target_host_contains":"string","rebuild_target_host_ends_with":"string","rebuild_target_host_gt":"string","rebuild_target_host_gte":"string","rebuild_target_host_in":["string"],"rebuild_target_host_lt":"string","rebuild_target_host_lte":"string","rebuild_target_host_not":"string","rebuild_target_host_not_contains":"string","rebuild_target_host_not_ends_with":"string","rebuild_target_host_not_in":["string"],"rebuild_target_host_not_starts_with":"string","rebuild_target_host_starts_with":"string","resource_version_gt":0,"resource_version_gte":0,"resource_version_in":[0],"resource_version_lt":0,"resource_version_lte":0,"resource_version_not":0,"resource_version_not_in":[0],"startup_after_restore":false,"startup_after_restore_not":false,"status":"ABORTED","status_in":["ABORTED"],"status_not":"ABORTED","status_not_in":["ABORTED"],"total_bytes":0,"total_bytes_gt":0,"total_bytes_gte":0,"total_bytes_in":[0],"total_bytes_lt":0,"total_bytes_lte":0,"total_bytes_not":0,"total_bytes_not_in":[0]}
//
// swagger:model BackupRestoreExecutionWhereInput
type BackupRestoreExecutionWhereInput struct {

	// a n d
	AND []*BackupRestoreExecutionWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*BackupRestoreExecutionWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*BackupRestoreExecutionWhereInput `json:"OR,omitempty"`

	// backup restore point
	BackupRestorePoint struct {
		BackupRestorePointWhereInput
	} `json:"backup_restore_point,omitempty"`

	// duration
	Duration *int32 `json:"duration,omitempty"`

	// duration gt
	DurationGt *int32 `json:"duration_gt,omitempty"`

	// duration gte
	DurationGte *int32 `json:"duration_gte,omitempty"`

	// duration in
	DurationIn []int32 `json:"duration_in,omitempty"`

	// duration lt
	DurationLt *int32 `json:"duration_lt,omitempty"`

	// duration lte
	DurationLte *int32 `json:"duration_lte,omitempty"`

	// duration not
	DurationNot *int32 `json:"duration_not,omitempty"`

	// duration not in
	DurationNotIn []int32 `json:"duration_not_in,omitempty"`

	// entity async status
	EntityAsyncStatus struct {
		EntityAsyncStatus
	} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot struct {
		EntityAsyncStatus
	} `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// executed at
	ExecutedAt *string `json:"executed_at,omitempty"`

	// executed at gt
	ExecutedAtGt *string `json:"executed_at_gt,omitempty"`

	// executed at gte
	ExecutedAtGte *string `json:"executed_at_gte,omitempty"`

	// executed at in
	ExecutedAtIn []string `json:"executed_at_in,omitempty"`

	// executed at lt
	ExecutedAtLt *string `json:"executed_at_lt,omitempty"`

	// executed at lte
	ExecutedAtLte *string `json:"executed_at_lte,omitempty"`

	// executed at not
	ExecutedAtNot *string `json:"executed_at_not,omitempty"`

	// executed at not in
	ExecutedAtNotIn []string `json:"executed_at_not_in,omitempty"`

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

	// mode
	Mode struct {
		BackupRestoreExecutionMode
	} `json:"mode,omitempty"`

	// mode in
	ModeIn []BackupRestoreExecutionMode `json:"mode_in,omitempty"`

	// mode not
	ModeNot struct {
		BackupRestoreExecutionMode
	} `json:"mode_not,omitempty"`

	// mode not in
	ModeNotIn []BackupRestoreExecutionMode `json:"mode_not_in,omitempty"`

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

	// read bytes
	ReadBytes *float64 `json:"read_bytes,omitempty"`

	// read bytes gt
	ReadBytesGt *float64 `json:"read_bytes_gt,omitempty"`

	// read bytes gte
	ReadBytesGte *float64 `json:"read_bytes_gte,omitempty"`

	// read bytes in
	ReadBytesIn []float64 `json:"read_bytes_in,omitempty"`

	// read bytes lt
	ReadBytesLt *float64 `json:"read_bytes_lt,omitempty"`

	// read bytes lte
	ReadBytesLte *float64 `json:"read_bytes_lte,omitempty"`

	// read bytes not
	ReadBytesNot *float64 `json:"read_bytes_not,omitempty"`

	// read bytes not in
	ReadBytesNotIn []float64 `json:"read_bytes_not_in,omitempty"`

	// rebuild name
	RebuildName *string `json:"rebuild_name,omitempty"`

	// rebuild name contains
	RebuildNameContains *string `json:"rebuild_name_contains,omitempty"`

	// rebuild name ends with
	RebuildNameEndsWith *string `json:"rebuild_name_ends_with,omitempty"`

	// rebuild name gt
	RebuildNameGt *string `json:"rebuild_name_gt,omitempty"`

	// rebuild name gte
	RebuildNameGte *string `json:"rebuild_name_gte,omitempty"`

	// rebuild name in
	RebuildNameIn []string `json:"rebuild_name_in,omitempty"`

	// rebuild name lt
	RebuildNameLt *string `json:"rebuild_name_lt,omitempty"`

	// rebuild name lte
	RebuildNameLte *string `json:"rebuild_name_lte,omitempty"`

	// rebuild name not
	RebuildNameNot *string `json:"rebuild_name_not,omitempty"`

	// rebuild name not contains
	RebuildNameNotContains *string `json:"rebuild_name_not_contains,omitempty"`

	// rebuild name not ends with
	RebuildNameNotEndsWith *string `json:"rebuild_name_not_ends_with,omitempty"`

	// rebuild name not in
	RebuildNameNotIn []string `json:"rebuild_name_not_in,omitempty"`

	// rebuild name not starts with
	RebuildNameNotStartsWith *string `json:"rebuild_name_not_starts_with,omitempty"`

	// rebuild name starts with
	RebuildNameStartsWith *string `json:"rebuild_name_starts_with,omitempty"`

	// rebuild target cluster
	RebuildTargetCluster *string `json:"rebuild_target_cluster,omitempty"`

	// rebuild target cluster contains
	RebuildTargetClusterContains *string `json:"rebuild_target_cluster_contains,omitempty"`

	// rebuild target cluster ends with
	RebuildTargetClusterEndsWith *string `json:"rebuild_target_cluster_ends_with,omitempty"`

	// rebuild target cluster gt
	RebuildTargetClusterGt *string `json:"rebuild_target_cluster_gt,omitempty"`

	// rebuild target cluster gte
	RebuildTargetClusterGte *string `json:"rebuild_target_cluster_gte,omitempty"`

	// rebuild target cluster in
	RebuildTargetClusterIn []string `json:"rebuild_target_cluster_in,omitempty"`

	// rebuild target cluster lt
	RebuildTargetClusterLt *string `json:"rebuild_target_cluster_lt,omitempty"`

	// rebuild target cluster lte
	RebuildTargetClusterLte *string `json:"rebuild_target_cluster_lte,omitempty"`

	// rebuild target cluster not
	RebuildTargetClusterNot *string `json:"rebuild_target_cluster_not,omitempty"`

	// rebuild target cluster not contains
	RebuildTargetClusterNotContains *string `json:"rebuild_target_cluster_not_contains,omitempty"`

	// rebuild target cluster not ends with
	RebuildTargetClusterNotEndsWith *string `json:"rebuild_target_cluster_not_ends_with,omitempty"`

	// rebuild target cluster not in
	RebuildTargetClusterNotIn []string `json:"rebuild_target_cluster_not_in,omitempty"`

	// rebuild target cluster not starts with
	RebuildTargetClusterNotStartsWith *string `json:"rebuild_target_cluster_not_starts_with,omitempty"`

	// rebuild target cluster starts with
	RebuildTargetClusterStartsWith *string `json:"rebuild_target_cluster_starts_with,omitempty"`

	// rebuild target host
	RebuildTargetHost *string `json:"rebuild_target_host,omitempty"`

	// rebuild target host contains
	RebuildTargetHostContains *string `json:"rebuild_target_host_contains,omitempty"`

	// rebuild target host ends with
	RebuildTargetHostEndsWith *string `json:"rebuild_target_host_ends_with,omitempty"`

	// rebuild target host gt
	RebuildTargetHostGt *string `json:"rebuild_target_host_gt,omitempty"`

	// rebuild target host gte
	RebuildTargetHostGte *string `json:"rebuild_target_host_gte,omitempty"`

	// rebuild target host in
	RebuildTargetHostIn []string `json:"rebuild_target_host_in,omitempty"`

	// rebuild target host lt
	RebuildTargetHostLt *string `json:"rebuild_target_host_lt,omitempty"`

	// rebuild target host lte
	RebuildTargetHostLte *string `json:"rebuild_target_host_lte,omitempty"`

	// rebuild target host not
	RebuildTargetHostNot *string `json:"rebuild_target_host_not,omitempty"`

	// rebuild target host not contains
	RebuildTargetHostNotContains *string `json:"rebuild_target_host_not_contains,omitempty"`

	// rebuild target host not ends with
	RebuildTargetHostNotEndsWith *string `json:"rebuild_target_host_not_ends_with,omitempty"`

	// rebuild target host not in
	RebuildTargetHostNotIn []string `json:"rebuild_target_host_not_in,omitempty"`

	// rebuild target host not starts with
	RebuildTargetHostNotStartsWith *string `json:"rebuild_target_host_not_starts_with,omitempty"`

	// rebuild target host starts with
	RebuildTargetHostStartsWith *string `json:"rebuild_target_host_starts_with,omitempty"`

	// resource version gt
	ResourceVersionGt *int32 `json:"resource_version_gt,omitempty"`

	// resource version gte
	ResourceVersionGte *int32 `json:"resource_version_gte,omitempty"`

	// resource version in
	ResourceVersionIn []int32 `json:"resource_version_in,omitempty"`

	// resource version lt
	ResourceVersionLt *int32 `json:"resource_version_lt,omitempty"`

	// resource version lte
	ResourceVersionLte *int32 `json:"resource_version_lte,omitempty"`

	// resource version not
	ResourceVersionNot *int32 `json:"resource_version_not,omitempty"`

	// resource version not in
	ResourceVersionNotIn []int32 `json:"resource_version_not_in,omitempty"`

	// startup after restore
	StartupAfterRestore *bool `json:"startup_after_restore,omitempty"`

	// startup after restore not
	StartupAfterRestoreNot *bool `json:"startup_after_restore_not,omitempty"`

	// status
	Status struct {
		BackupExecutionStatus
	} `json:"status,omitempty"`

	// status in
	StatusIn []BackupExecutionStatus `json:"status_in,omitempty"`

	// status not
	StatusNot struct {
		BackupExecutionStatus
	} `json:"status_not,omitempty"`

	// status not in
	StatusNotIn []BackupExecutionStatus `json:"status_not_in,omitempty"`

	// total bytes
	TotalBytes *float64 `json:"total_bytes,omitempty"`

	// total bytes gt
	TotalBytesGt *float64 `json:"total_bytes_gt,omitempty"`

	// total bytes gte
	TotalBytesGte *float64 `json:"total_bytes_gte,omitempty"`

	// total bytes in
	TotalBytesIn []float64 `json:"total_bytes_in,omitempty"`

	// total bytes lt
	TotalBytesLt *float64 `json:"total_bytes_lt,omitempty"`

	// total bytes lte
	TotalBytesLte *float64 `json:"total_bytes_lte,omitempty"`

	// total bytes not
	TotalBytesNot *float64 `json:"total_bytes_not,omitempty"`

	// total bytes not in
	TotalBytesNotIn []float64 `json:"total_bytes_not_in,omitempty"`
}

// Validate validates this backup restore execution where input
func (m *BackupRestoreExecutionWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateBackupRestorePoint(formats); err != nil {
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

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModeNotIn(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *BackupRestoreExecutionWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *BackupRestoreExecutionWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *BackupRestoreExecutionWhereInput) validateBackupRestorePoint(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRestorePoint) { // not required
		return nil
	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
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

func (m *BackupRestoreExecutionWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
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

func (m *BackupRestoreExecutionWhereInput) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateModeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ModeIn); i++ {

		if err := m.ModeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateModeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeNot) { // not required
		return nil
	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateModeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ModeNotIn); i++ {

		if err := m.ModeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateStatusIn(formats strfmt.Registry) error {
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

func (m *BackupRestoreExecutionWhereInput) validateStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNot) { // not required
		return nil
	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) validateStatusNotIn(formats strfmt.Registry) error {
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

// ContextValidate validate this backup restore execution where input based on the context it is used
func (m *BackupRestoreExecutionWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateBackupRestorePoint(ctx, formats); err != nil {
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

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModeNotIn(ctx, formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupRestoreExecutionWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupRestoreExecutionWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupRestoreExecutionWhereInput) contextValidateBackupRestorePoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupRestoreExecutionWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupRestoreExecutionWhereInput) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateModeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModeIn); i++ {

		if err := m.ModeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateModeNot(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateModeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModeNotIn); i++ {

		if err := m.ModeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupRestoreExecutionWhereInput) contextValidateStatusNot(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupRestoreExecutionWhereInput) contextValidateStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *BackupRestoreExecutionWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRestoreExecutionWhereInput) UnmarshalBinary(b []byte) error {
	var res BackupRestoreExecutionWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
