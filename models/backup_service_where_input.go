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

// BackupServiceWhereInput backup service where input
// Example: {"AND":"BackupServiceWhereInput[]","NOT":"BackupServiceWhereInput[]","OR":"BackupServiceWhereInput[]","backup_clusters_every":"ClusterWhereInput","backup_clusters_none":"ClusterWhereInput","backup_clusters_some":"ClusterWhereInput","backup_package":"BackupPackageWhereInput","backup_plans_every":"BackupPlanWhereInput","backup_plans_none":"BackupPlanWhereInput","backup_plans_some":"BackupPlanWhereInput","backup_store_repositories_every":"BackupStoreRepositoryWhereInput","backup_store_repositories_none":"BackupStoreRepositoryWhereInput","backup_store_repositories_some":"BackupStoreRepositoryWhereInput","description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"gateway":"string","gateway_contains":"string","gateway_ends_with":"string","gateway_gt":"string","gateway_gte":"string","gateway_in":["string"],"gateway_lt":"string","gateway_lte":"string","gateway_not":"string","gateway_not_contains":"string","gateway_not_ends_with":"string","gateway_not_in":["string"],"gateway_not_starts_with":"string","gateway_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","iops_limit":0,"iops_limit_gt":0,"iops_limit_gte":0,"iops_limit_in":[0],"iops_limit_lt":0,"iops_limit_lte":0,"iops_limit_not":0,"iops_limit_not_in":[0],"ip":"string","ip_contains":"string","ip_ends_with":"string","ip_gt":"string","ip_gte":"string","ip_in":["string"],"ip_lt":"string","ip_lte":"string","ip_not":"string","ip_not_contains":"string","ip_not_ends_with":"string","ip_not_in":["string"],"ip_not_starts_with":"string","ip_starts_with":"string","kube_config":"string","kube_config_contains":"string","kube_config_ends_with":"string","kube_config_gt":"string","kube_config_gte":"string","kube_config_in":["string"],"kube_config_lt":"string","kube_config_lte":"string","kube_config_not":"string","kube_config_not_contains":"string","kube_config_not_ends_with":"string","kube_config_not_in":["string"],"kube_config_not_starts_with":"string","kube_config_starts_with":"string","max_job_retry_times":0,"max_job_retry_times_gt":0,"max_job_retry_times_gte":0,"max_job_retry_times_in":[0],"max_job_retry_times_lt":0,"max_job_retry_times_lte":0,"max_job_retry_times_not":0,"max_job_retry_times_not_in":[0],"max_parallel_backup_jobs":0,"max_parallel_backup_jobs_gt":0,"max_parallel_backup_jobs_gte":0,"max_parallel_backup_jobs_in":[0],"max_parallel_backup_jobs_lt":0,"max_parallel_backup_jobs_lte":0,"max_parallel_backup_jobs_not":0,"max_parallel_backup_jobs_not_in":[0],"max_parallel_restore_jobs":0,"max_parallel_restore_jobs_gt":0,"max_parallel_restore_jobs_gte":0,"max_parallel_restore_jobs_in":[0],"max_parallel_restore_jobs_lt":0,"max_parallel_restore_jobs_lte":0,"max_parallel_restore_jobs_not":0,"max_parallel_restore_jobs_not_in":[0],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","resource_version_gt":0,"resource_version_gte":0,"resource_version_in":[0],"resource_version_lt":0,"resource_version_lte":0,"resource_version_not":0,"resource_version_not_in":[0],"retry_interval":0,"retry_interval_gt":0,"retry_interval_gte":0,"retry_interval_in":[0],"retry_interval_lt":0,"retry_interval_lte":0,"retry_interval_not":0,"retry_interval_not_in":[0],"running_vm":"VmWhereInput","status":"INSTALLING","status_in":["INSTALLING"],"status_not":"INSTALLING","status_not_in":["INSTALLING"],"subnet_mask":"string","subnet_mask_contains":"string","subnet_mask_ends_with":"string","subnet_mask_gt":"string","subnet_mask_gte":"string","subnet_mask_in":["string"],"subnet_mask_lt":"string","subnet_mask_lte":"string","subnet_mask_not":"string","subnet_mask_not_contains":"string","subnet_mask_not_ends_with":"string","subnet_mask_not_in":["string"],"subnet_mask_not_starts_with":"string","subnet_mask_starts_with":"string"}
//
// swagger:model BackupServiceWhereInput
type BackupServiceWhereInput struct {

	// a n d
	AND []*BackupServiceWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*BackupServiceWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*BackupServiceWhereInput `json:"OR,omitempty"`

	// backup clusters every
	BackupClustersEvery interface{} `json:"backup_clusters_every,omitempty"`

	// backup clusters none
	BackupClustersNone interface{} `json:"backup_clusters_none,omitempty"`

	// backup clusters some
	BackupClustersSome interface{} `json:"backup_clusters_some,omitempty"`

	// backup package
	BackupPackage interface{} `json:"backup_package,omitempty"`

	// backup plans every
	BackupPlansEvery interface{} `json:"backup_plans_every,omitempty"`

	// backup plans none
	BackupPlansNone interface{} `json:"backup_plans_none,omitempty"`

	// backup plans some
	BackupPlansSome interface{} `json:"backup_plans_some,omitempty"`

	// backup store repositories every
	BackupStoreRepositoriesEvery interface{} `json:"backup_store_repositories_every,omitempty"`

	// backup store repositories none
	BackupStoreRepositoriesNone interface{} `json:"backup_store_repositories_none,omitempty"`

	// backup store repositories some
	BackupStoreRepositoriesSome interface{} `json:"backup_store_repositories_some,omitempty"`

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

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot interface{} `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// gateway
	Gateway *string `json:"gateway,omitempty"`

	// gateway contains
	GatewayContains *string `json:"gateway_contains,omitempty"`

	// gateway ends with
	GatewayEndsWith *string `json:"gateway_ends_with,omitempty"`

	// gateway gt
	GatewayGt *string `json:"gateway_gt,omitempty"`

	// gateway gte
	GatewayGte *string `json:"gateway_gte,omitempty"`

	// gateway in
	GatewayIn []string `json:"gateway_in,omitempty"`

	// gateway lt
	GatewayLt *string `json:"gateway_lt,omitempty"`

	// gateway lte
	GatewayLte *string `json:"gateway_lte,omitempty"`

	// gateway not
	GatewayNot *string `json:"gateway_not,omitempty"`

	// gateway not contains
	GatewayNotContains *string `json:"gateway_not_contains,omitempty"`

	// gateway not ends with
	GatewayNotEndsWith *string `json:"gateway_not_ends_with,omitempty"`

	// gateway not in
	GatewayNotIn []string `json:"gateway_not_in,omitempty"`

	// gateway not starts with
	GatewayNotStartsWith *string `json:"gateway_not_starts_with,omitempty"`

	// gateway starts with
	GatewayStartsWith *string `json:"gateway_starts_with,omitempty"`

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

	// iops limit
	IopsLimit *int32 `json:"iops_limit,omitempty"`

	// iops limit gt
	IopsLimitGt *int32 `json:"iops_limit_gt,omitempty"`

	// iops limit gte
	IopsLimitGte *int32 `json:"iops_limit_gte,omitempty"`

	// iops limit in
	IopsLimitIn []int32 `json:"iops_limit_in,omitempty"`

	// iops limit lt
	IopsLimitLt *int32 `json:"iops_limit_lt,omitempty"`

	// iops limit lte
	IopsLimitLte *int32 `json:"iops_limit_lte,omitempty"`

	// iops limit not
	IopsLimitNot *int32 `json:"iops_limit_not,omitempty"`

	// iops limit not in
	IopsLimitNotIn []int32 `json:"iops_limit_not_in,omitempty"`

	// ip
	IP *string `json:"ip,omitempty"`

	// ip contains
	IPContains *string `json:"ip_contains,omitempty"`

	// ip ends with
	IPEndsWith *string `json:"ip_ends_with,omitempty"`

	// ip gt
	IPGt *string `json:"ip_gt,omitempty"`

	// ip gte
	IPGte *string `json:"ip_gte,omitempty"`

	// ip in
	IPIn []string `json:"ip_in,omitempty"`

	// ip lt
	IPLt *string `json:"ip_lt,omitempty"`

	// ip lte
	IPLte *string `json:"ip_lte,omitempty"`

	// ip not
	IPNot *string `json:"ip_not,omitempty"`

	// ip not contains
	IPNotContains *string `json:"ip_not_contains,omitempty"`

	// ip not ends with
	IPNotEndsWith *string `json:"ip_not_ends_with,omitempty"`

	// ip not in
	IPNotIn []string `json:"ip_not_in,omitempty"`

	// ip not starts with
	IPNotStartsWith *string `json:"ip_not_starts_with,omitempty"`

	// ip starts with
	IPStartsWith *string `json:"ip_starts_with,omitempty"`

	// kube config
	KubeConfig *string `json:"kube_config,omitempty"`

	// kube config contains
	KubeConfigContains *string `json:"kube_config_contains,omitempty"`

	// kube config ends with
	KubeConfigEndsWith *string `json:"kube_config_ends_with,omitempty"`

	// kube config gt
	KubeConfigGt *string `json:"kube_config_gt,omitempty"`

	// kube config gte
	KubeConfigGte *string `json:"kube_config_gte,omitempty"`

	// kube config in
	KubeConfigIn []string `json:"kube_config_in,omitempty"`

	// kube config lt
	KubeConfigLt *string `json:"kube_config_lt,omitempty"`

	// kube config lte
	KubeConfigLte *string `json:"kube_config_lte,omitempty"`

	// kube config not
	KubeConfigNot *string `json:"kube_config_not,omitempty"`

	// kube config not contains
	KubeConfigNotContains *string `json:"kube_config_not_contains,omitempty"`

	// kube config not ends with
	KubeConfigNotEndsWith *string `json:"kube_config_not_ends_with,omitempty"`

	// kube config not in
	KubeConfigNotIn []string `json:"kube_config_not_in,omitempty"`

	// kube config not starts with
	KubeConfigNotStartsWith *string `json:"kube_config_not_starts_with,omitempty"`

	// kube config starts with
	KubeConfigStartsWith *string `json:"kube_config_starts_with,omitempty"`

	// max job retry times
	MaxJobRetryTimes *int32 `json:"max_job_retry_times,omitempty"`

	// max job retry times gt
	MaxJobRetryTimesGt *int32 `json:"max_job_retry_times_gt,omitempty"`

	// max job retry times gte
	MaxJobRetryTimesGte *int32 `json:"max_job_retry_times_gte,omitempty"`

	// max job retry times in
	MaxJobRetryTimesIn []int32 `json:"max_job_retry_times_in,omitempty"`

	// max job retry times lt
	MaxJobRetryTimesLt *int32 `json:"max_job_retry_times_lt,omitempty"`

	// max job retry times lte
	MaxJobRetryTimesLte *int32 `json:"max_job_retry_times_lte,omitempty"`

	// max job retry times not
	MaxJobRetryTimesNot *int32 `json:"max_job_retry_times_not,omitempty"`

	// max job retry times not in
	MaxJobRetryTimesNotIn []int32 `json:"max_job_retry_times_not_in,omitempty"`

	// max parallel backup jobs
	MaxParallelBackupJobs *int32 `json:"max_parallel_backup_jobs,omitempty"`

	// max parallel backup jobs gt
	MaxParallelBackupJobsGt *int32 `json:"max_parallel_backup_jobs_gt,omitempty"`

	// max parallel backup jobs gte
	MaxParallelBackupJobsGte *int32 `json:"max_parallel_backup_jobs_gte,omitempty"`

	// max parallel backup jobs in
	MaxParallelBackupJobsIn []int32 `json:"max_parallel_backup_jobs_in,omitempty"`

	// max parallel backup jobs lt
	MaxParallelBackupJobsLt *int32 `json:"max_parallel_backup_jobs_lt,omitempty"`

	// max parallel backup jobs lte
	MaxParallelBackupJobsLte *int32 `json:"max_parallel_backup_jobs_lte,omitempty"`

	// max parallel backup jobs not
	MaxParallelBackupJobsNot *int32 `json:"max_parallel_backup_jobs_not,omitempty"`

	// max parallel backup jobs not in
	MaxParallelBackupJobsNotIn []int32 `json:"max_parallel_backup_jobs_not_in,omitempty"`

	// max parallel restore jobs
	MaxParallelRestoreJobs *int32 `json:"max_parallel_restore_jobs,omitempty"`

	// max parallel restore jobs gt
	MaxParallelRestoreJobsGt *int32 `json:"max_parallel_restore_jobs_gt,omitempty"`

	// max parallel restore jobs gte
	MaxParallelRestoreJobsGte *int32 `json:"max_parallel_restore_jobs_gte,omitempty"`

	// max parallel restore jobs in
	MaxParallelRestoreJobsIn []int32 `json:"max_parallel_restore_jobs_in,omitempty"`

	// max parallel restore jobs lt
	MaxParallelRestoreJobsLt *int32 `json:"max_parallel_restore_jobs_lt,omitempty"`

	// max parallel restore jobs lte
	MaxParallelRestoreJobsLte *int32 `json:"max_parallel_restore_jobs_lte,omitempty"`

	// max parallel restore jobs not
	MaxParallelRestoreJobsNot *int32 `json:"max_parallel_restore_jobs_not,omitempty"`

	// max parallel restore jobs not in
	MaxParallelRestoreJobsNotIn []int32 `json:"max_parallel_restore_jobs_not_in,omitempty"`

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

	// retry interval
	RetryInterval *int32 `json:"retry_interval,omitempty"`

	// retry interval gt
	RetryIntervalGt *int32 `json:"retry_interval_gt,omitempty"`

	// retry interval gte
	RetryIntervalGte *int32 `json:"retry_interval_gte,omitempty"`

	// retry interval in
	RetryIntervalIn []int32 `json:"retry_interval_in,omitempty"`

	// retry interval lt
	RetryIntervalLt *int32 `json:"retry_interval_lt,omitempty"`

	// retry interval lte
	RetryIntervalLte *int32 `json:"retry_interval_lte,omitempty"`

	// retry interval not
	RetryIntervalNot *int32 `json:"retry_interval_not,omitempty"`

	// retry interval not in
	RetryIntervalNotIn []int32 `json:"retry_interval_not_in,omitempty"`

	// running vm
	RunningVM interface{} `json:"running_vm,omitempty"`

	// status
	Status interface{} `json:"status,omitempty"`

	// status in
	StatusIn []BackupServiceStatus `json:"status_in,omitempty"`

	// status not
	StatusNot interface{} `json:"status_not,omitempty"`

	// status not in
	StatusNotIn []BackupServiceStatus `json:"status_not_in,omitempty"`

	// subnet mask
	SubnetMask *string `json:"subnet_mask,omitempty"`

	// subnet mask contains
	SubnetMaskContains *string `json:"subnet_mask_contains,omitempty"`

	// subnet mask ends with
	SubnetMaskEndsWith *string `json:"subnet_mask_ends_with,omitempty"`

	// subnet mask gt
	SubnetMaskGt *string `json:"subnet_mask_gt,omitempty"`

	// subnet mask gte
	SubnetMaskGte *string `json:"subnet_mask_gte,omitempty"`

	// subnet mask in
	SubnetMaskIn []string `json:"subnet_mask_in,omitempty"`

	// subnet mask lt
	SubnetMaskLt *string `json:"subnet_mask_lt,omitempty"`

	// subnet mask lte
	SubnetMaskLte *string `json:"subnet_mask_lte,omitempty"`

	// subnet mask not
	SubnetMaskNot *string `json:"subnet_mask_not,omitempty"`

	// subnet mask not contains
	SubnetMaskNotContains *string `json:"subnet_mask_not_contains,omitempty"`

	// subnet mask not ends with
	SubnetMaskNotEndsWith *string `json:"subnet_mask_not_ends_with,omitempty"`

	// subnet mask not in
	SubnetMaskNotIn []string `json:"subnet_mask_not_in,omitempty"`

	// subnet mask not starts with
	SubnetMaskNotStartsWith *string `json:"subnet_mask_not_starts_with,omitempty"`

	// subnet mask starts with
	SubnetMaskStartsWith *string `json:"subnet_mask_starts_with,omitempty"`
}

// Validate validates this backup service where input
func (m *BackupServiceWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusIn(formats); err != nil {
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

func (m *BackupServiceWhereInput) validateAND(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) validateNOT(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) validateOR(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) validateStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) validateStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this backup service where input based on the context it is used
func (m *BackupServiceWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusIn(ctx, formats); err != nil {
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

func (m *BackupServiceWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) contextValidateStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupServiceWhereInput) contextValidateStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupServiceWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupServiceWhereInput) UnmarshalBinary(b []byte) error {
	var res BackupServiceWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
