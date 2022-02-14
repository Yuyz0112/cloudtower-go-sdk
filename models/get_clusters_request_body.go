// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetClustersRequestBody get clusters request body
// Example: {"after":"clusters-id-string","before":"clusters-id-string","first":0,"last":0,"orderBy":"architecture_ASC","skip":0,"where":{"AND":"ClusterWhereInput[]","NOT":"ClusterWhereInput[]","OR":"ClusterWhereInput[]","applications_every":"ApplicationWhereInput","applications_none":"ApplicationWhereInput","applications_some":"ApplicationWhereInput","architecture":"AARCH64","architecture_in":["AARCH64"],"architecture_not":"AARCH64","architecture_not_in":["AARCH64"],"auto_converge":false,"auto_converge_not":false,"backup_by_service":"BackupServiceWhereInput","connect_state":"CONNECTED","connect_state_in":["CONNECTED"],"connect_state_not":"CONNECTED","connect_state_not_in":["CONNECTED"],"consistency_groups_every":"ConsistencyGroupWhereInput","consistency_groups_none":"ConsistencyGroupWhereInput","consistency_groups_some":"ConsistencyGroupWhereInput","current_cpu_model":"string","current_cpu_model_contains":"string","current_cpu_model_ends_with":"string","current_cpu_model_gt":"string","current_cpu_model_gte":"string","current_cpu_model_in":["string"],"current_cpu_model_lt":"string","current_cpu_model_lte":"string","current_cpu_model_not":"string","current_cpu_model_not_contains":"string","current_cpu_model_not_ends_with":"string","current_cpu_model_not_in":["string"],"current_cpu_model_not_starts_with":"string","current_cpu_model_starts_with":"string","datacenters_every":"DatacenterWhereInput","datacenters_none":"DatacenterWhereInput","datacenters_some":"DatacenterWhereInput","disconnected_date":"string","disconnected_date_gt":"string","disconnected_date_gte":"string","disconnected_date_in":["string"],"disconnected_date_lt":"string","disconnected_date_lte":"string","disconnected_date_not":"string","disconnected_date_not_in":["string"],"disconnected_reason":"LOAD_CLUSTER_FAILED","disconnected_reason_in":["LOAD_CLUSTER_FAILED"],"disconnected_reason_not":"LOAD_CLUSTER_FAILED","disconnected_reason_not_in":["LOAD_CLUSTER_FAILED"],"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"everoute_cluster":"EverouteClusterWhereInput","failure_data_space":0,"failure_data_space_gt":0,"failure_data_space_gte":0,"failure_data_space_in":[0],"failure_data_space_lt":0,"failure_data_space_lte":0,"failure_data_space_not":0,"failure_data_space_not_in":[0],"has_metrox":false,"has_metrox_not":false,"has_remote_backup":false,"has_remote_backup_not":false,"host_num":0,"host_num_gt":0,"host_num_gte":0,"host_num_in":[0],"host_num_lt":0,"host_num_lte":0,"host_num_not":0,"host_num_not_in":[0],"hosts_every":"HostWhereInput","hosts_none":"HostWhereInput","hosts_some":"HostWhereInput","hypervisor":"BLUESHARK","hypervisor_in":["BLUESHARK"],"hypervisor_not":"BLUESHARK","hypervisor_not_in":["BLUESHARK"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","ip":"string","ip_contains":"string","ip_ends_with":"string","ip_gt":"string","ip_gte":"string","ip_in":["string"],"ip_lt":"string","ip_lte":"string","ip_not":"string","ip_not_contains":"string","ip_not_ends_with":"string","ip_not_in":["string"],"ip_not_starts_with":"string","ip_starts_with":"string","is_all_flash":false,"is_all_flash_not":false,"iscsi_vip":"string","iscsi_vip_contains":"string","iscsi_vip_ends_with":"string","iscsi_vip_gt":"string","iscsi_vip_gte":"string","iscsi_vip_in":["string"],"iscsi_vip_lt":"string","iscsi_vip_lte":"string","iscsi_vip_not":"string","iscsi_vip_not_contains":"string","iscsi_vip_not_ends_with":"string","iscsi_vip_not_in":["string"],"iscsi_vip_not_starts_with":"string","iscsi_vip_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","license_expire_date":"string","license_expire_date_gt":"string","license_expire_date_gte":"string","license_expire_date_in":["string"],"license_expire_date_lt":"string","license_expire_date_lte":"string","license_expire_date_not":"string","license_expire_date_not_in":["string"],"license_serial":"string","license_serial_contains":"string","license_serial_ends_with":"string","license_serial_gt":"string","license_serial_gte":"string","license_serial_in":["string"],"license_serial_lt":"string","license_serial_lte":"string","license_serial_not":"string","license_serial_not_contains":"string","license_serial_not_ends_with":"string","license_serial_not_in":["string"],"license_serial_not_starts_with":"string","license_serial_starts_with":"string","license_sign_date":"string","license_sign_date_gt":"string","license_sign_date_gte":"string","license_sign_date_in":["string"],"license_sign_date_lt":"string","license_sign_date_lte":"string","license_sign_date_not":"string","license_sign_date_not_in":["string"],"license_type":"PERPETUAL","license_type_in":["PERPETUAL"],"license_type_not":"PERPETUAL","license_type_not_in":["PERPETUAL"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","maintenance_end_date":"string","maintenance_end_date_gt":"string","maintenance_end_date_gte":"string","maintenance_end_date_in":["string"],"maintenance_end_date_lt":"string","maintenance_end_date_lte":"string","maintenance_end_date_not":"string","maintenance_end_date_not_in":["string"],"maintenance_start_date":"string","maintenance_start_date_gt":"string","maintenance_start_date_gte":"string","maintenance_start_date_in":["string"],"maintenance_start_date_lt":"string","maintenance_start_date_lte":"string","maintenance_start_date_not":"string","maintenance_start_date_not_in":["string"],"management_vip":"string","management_vip_contains":"string","management_vip_ends_with":"string","management_vip_gt":"string","management_vip_gte":"string","management_vip_in":["string"],"management_vip_lt":"string","management_vip_lte":"string","management_vip_not":"string","management_vip_not_contains":"string","management_vip_not_ends_with":"string","management_vip_not_in":["string"],"management_vip_not_starts_with":"string","management_vip_starts_with":"string","max_chunk_num":0,"max_chunk_num_gt":0,"max_chunk_num_gte":0,"max_chunk_num_in":[0],"max_chunk_num_lt":0,"max_chunk_num_lte":0,"max_chunk_num_not":0,"max_chunk_num_not_in":[0],"max_physical_data_capacity":0,"max_physical_data_capacity_gt":0,"max_physical_data_capacity_gte":0,"max_physical_data_capacity_in":[0],"max_physical_data_capacity_lt":0,"max_physical_data_capacity_lte":0,"max_physical_data_capacity_not":0,"max_physical_data_capacity_not_in":[0],"max_physical_data_capacity_per_node":0,"max_physical_data_capacity_per_node_gt":0,"max_physical_data_capacity_per_node_gte":0,"max_physical_data_capacity_per_node_in":[0],"max_physical_data_capacity_per_node_lt":0,"max_physical_data_capacity_per_node_lte":0,"max_physical_data_capacity_per_node_not":0,"max_physical_data_capacity_per_node_not_in":[0],"mgt_gateway":"string","mgt_gateway_contains":"string","mgt_gateway_ends_with":"string","mgt_gateway_gt":"string","mgt_gateway_gte":"string","mgt_gateway_in":["string"],"mgt_gateway_lt":"string","mgt_gateway_lte":"string","mgt_gateway_not":"string","mgt_gateway_not_contains":"string","mgt_gateway_not_ends_with":"string","mgt_gateway_not_in":["string"],"mgt_gateway_not_starts_with":"string","mgt_gateway_starts_with":"string","mgt_netmask":"string","mgt_netmask_contains":"string","mgt_netmask_ends_with":"string","mgt_netmask_gt":"string","mgt_netmask_gte":"string","mgt_netmask_in":["string"],"mgt_netmask_lt":"string","mgt_netmask_lte":"string","mgt_netmask_not":"string","mgt_netmask_not_contains":"string","mgt_netmask_not_ends_with":"string","mgt_netmask_not_in":["string"],"mgt_netmask_not_starts_with":"string","mgt_netmask_starts_with":"string","migration_data_size":0,"migration_data_size_gt":0,"migration_data_size_gte":0,"migration_data_size_in":[0],"migration_data_size_lt":0,"migration_data_size_lte":0,"migration_data_size_not":0,"migration_data_size_not_in":[0],"migration_speed":0,"migration_speed_gt":0,"migration_speed_gte":0,"migration_speed_in":[0],"migration_speed_lt":0,"migration_speed_lte":0,"migration_speed_not":0,"migration_speed_not_in":[0],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","ntp_mode":"EXTERNAL","ntp_mode_in":["EXTERNAL"],"ntp_mode_not":"EXTERNAL","ntp_mode_not_in":["EXTERNAL"],"nvmf_enabled":false,"nvmf_enabled_not":false,"password":"string","password_contains":"string","password_ends_with":"string","password_gt":"string","password_gte":"string","password_in":["string"],"password_lt":"string","password_lte":"string","password_not":"string","password_not_contains":"string","password_not_ends_with":"string","password_not_in":["string"],"password_not_starts_with":"string","password_starts_with":"string","pmem_enabled":false,"pmem_enabled_not":false,"provisioned_cpu_cores":0,"provisioned_cpu_cores_for_active_vm":0,"provisioned_cpu_cores_for_active_vm_gt":0,"provisioned_cpu_cores_for_active_vm_gte":0,"provisioned_cpu_cores_for_active_vm_in":[0],"provisioned_cpu_cores_for_active_vm_lt":0,"provisioned_cpu_cores_for_active_vm_lte":0,"provisioned_cpu_cores_for_active_vm_not":0,"provisioned_cpu_cores_for_active_vm_not_in":[0],"provisioned_cpu_cores_gt":0,"provisioned_cpu_cores_gte":0,"provisioned_cpu_cores_in":[0],"provisioned_cpu_cores_lt":0,"provisioned_cpu_cores_lte":0,"provisioned_cpu_cores_not":0,"provisioned_cpu_cores_not_in":[0],"provisioned_for_active_vm_ratio":0,"provisioned_for_active_vm_ratio_gt":0,"provisioned_for_active_vm_ratio_gte":0,"provisioned_for_active_vm_ratio_in":[0],"provisioned_for_active_vm_ratio_lt":0,"provisioned_for_active_vm_ratio_lte":0,"provisioned_for_active_vm_ratio_not":0,"provisioned_for_active_vm_ratio_not_in":[0],"provisioned_memory_bytes":0,"provisioned_memory_bytes_gt":0,"provisioned_memory_bytes_gte":0,"provisioned_memory_bytes_in":[0],"provisioned_memory_bytes_lt":0,"provisioned_memory_bytes_lte":0,"provisioned_memory_bytes_not":0,"provisioned_memory_bytes_not_in":[0],"provisioned_ratio":0,"provisioned_ratio_gt":0,"provisioned_ratio_gte":0,"provisioned_ratio_in":[0],"provisioned_ratio_lt":0,"provisioned_ratio_lte":0,"provisioned_ratio_not":0,"provisioned_ratio_not_in":[0],"rdma_enabled":false,"rdma_enabled_not":false,"recover_data_size":0,"recover_data_size_gt":0,"recover_data_size_gte":0,"recover_data_size_in":[0],"recover_data_size_lt":0,"recover_data_size_lte":0,"recover_data_size_not":0,"recover_data_size_not_in":[0],"recover_speed":0,"recover_speed_gt":0,"recover_speed_gte":0,"recover_speed_in":[0],"recover_speed_lt":0,"recover_speed_lte":0,"recover_speed_not":0,"recover_speed_not_in":[0],"reserved_cpu_cores_for_system_service":0,"reserved_cpu_cores_for_system_service_gt":0,"reserved_cpu_cores_for_system_service_gte":0,"reserved_cpu_cores_for_system_service_in":[0],"reserved_cpu_cores_for_system_service_lt":0,"reserved_cpu_cores_for_system_service_lte":0,"reserved_cpu_cores_for_system_service_not":0,"reserved_cpu_cores_for_system_service_not_in":[0],"running_vm_num":0,"running_vm_num_gt":0,"running_vm_num_gte":0,"running_vm_num_in":[0],"running_vm_num_lt":0,"running_vm_num_lte":0,"running_vm_num_not":0,"running_vm_num_not_in":[0],"settings":"ClusterSettingsWhereInput","software_edition":"COMMUNITY","software_edition_in":["COMMUNITY"],"software_edition_not":"COMMUNITY","software_edition_not_in":["COMMUNITY"],"stopped_vm_num":0,"stopped_vm_num_gt":0,"stopped_vm_num_gte":0,"stopped_vm_num_in":[0],"stopped_vm_num_lt":0,"stopped_vm_num_lte":0,"stopped_vm_num_not":0,"stopped_vm_num_not_in":[0],"stretch":false,"stretch_not":false,"suspended_vm_num":0,"suspended_vm_num_gt":0,"suspended_vm_num_gte":0,"suspended_vm_num_in":[0],"suspended_vm_num_lt":0,"suspended_vm_num_lte":0,"suspended_vm_num_not":0,"suspended_vm_num_not_in":[0],"total_cache_capacity":0,"total_cache_capacity_gt":0,"total_cache_capacity_gte":0,"total_cache_capacity_in":[0],"total_cache_capacity_lt":0,"total_cache_capacity_lte":0,"total_cache_capacity_not":0,"total_cache_capacity_not_in":[0],"total_cpu_cores":0,"total_cpu_cores_gt":0,"total_cpu_cores_gte":0,"total_cpu_cores_in":[0],"total_cpu_cores_lt":0,"total_cpu_cores_lte":0,"total_cpu_cores_not":0,"total_cpu_cores_not_in":[0],"total_cpu_hz":0,"total_cpu_hz_gt":0,"total_cpu_hz_gte":0,"total_cpu_hz_in":[0],"total_cpu_hz_lt":0,"total_cpu_hz_lte":0,"total_cpu_hz_not":0,"total_cpu_hz_not_in":[0],"total_cpu_sockets":0,"total_cpu_sockets_gt":0,"total_cpu_sockets_gte":0,"total_cpu_sockets_in":[0],"total_cpu_sockets_lt":0,"total_cpu_sockets_lte":0,"total_cpu_sockets_not":0,"total_cpu_sockets_not_in":[0],"total_data_capacity":0,"total_data_capacity_gt":0,"total_data_capacity_gte":0,"total_data_capacity_in":[0],"total_data_capacity_lt":0,"total_data_capacity_lte":0,"total_data_capacity_not":0,"total_data_capacity_not_in":[0],"total_memory_bytes":0,"total_memory_bytes_gt":0,"total_memory_bytes_gte":0,"total_memory_bytes_in":[0],"total_memory_bytes_lt":0,"total_memory_bytes_lte":0,"total_memory_bytes_not":0,"total_memory_bytes_not_in":[0],"type":"BLUESHARK","type_in":["BLUESHARK"],"type_not":"BLUESHARK","type_not_in":["BLUESHARK"],"upgrade_tool_version":"string","upgrade_tool_version_contains":"string","upgrade_tool_version_ends_with":"string","upgrade_tool_version_gt":"string","upgrade_tool_version_gte":"string","upgrade_tool_version_in":["string"],"upgrade_tool_version_lt":"string","upgrade_tool_version_lte":"string","upgrade_tool_version_not":"string","upgrade_tool_version_not_contains":"string","upgrade_tool_version_not_ends_with":"string","upgrade_tool_version_not_in":["string"],"upgrade_tool_version_not_starts_with":"string","upgrade_tool_version_starts_with":"string","used_cpu_hz":0,"used_cpu_hz_gt":0,"used_cpu_hz_gte":0,"used_cpu_hz_in":[0],"used_cpu_hz_lt":0,"used_cpu_hz_lte":0,"used_cpu_hz_not":0,"used_cpu_hz_not_in":[0],"used_data_space":0,"used_data_space_gt":0,"used_data_space_gte":0,"used_data_space_in":[0],"used_data_space_lt":0,"used_data_space_lte":0,"used_data_space_not":0,"used_data_space_not_in":[0],"used_memory_bytes":0,"used_memory_bytes_gt":0,"used_memory_bytes_gte":0,"used_memory_bytes_in":[0],"used_memory_bytes_lt":0,"used_memory_bytes_lte":0,"used_memory_bytes_not":0,"used_memory_bytes_not_in":[0],"username":"string","username_contains":"string","username_ends_with":"string","username_gt":"string","username_gte":"string","username_in":["string"],"username_lt":"string","username_lte":"string","username_not":"string","username_not_contains":"string","username_not_ends_with":"string","username_not_in":["string"],"username_not_starts_with":"string","username_starts_with":"string","valid_data_space":0,"valid_data_space_gt":0,"valid_data_space_gte":0,"valid_data_space_in":[0],"valid_data_space_lt":0,"valid_data_space_lte":0,"valid_data_space_not":0,"valid_data_space_not_in":[0],"vcenterAccount":"VcenterAccountWhereInput","vdses_every":"VdsWhereInput","vdses_none":"VdsWhereInput","vdses_some":"VdsWhereInput","version":"string","version_contains":"string","version_ends_with":"string","version_gt":"string","version_gte":"string","version_in":["string"],"version_lt":"string","version_lte":"string","version_not":"string","version_not_contains":"string","version_not_ends_with":"string","version_not_in":["string"],"version_not_starts_with":"string","version_semantic_gt":"string","version_semantic_gte":"string","version_semantic_lt":"string","version_semantic_lte":"string","version_starts_with":"string","vhost_enabled":false,"vhost_enabled_not":false,"vm_folders_every":"VmFolderWhereInput","vm_folders_none":"VmFolderWhereInput","vm_folders_some":"VmFolderWhereInput","vm_num":0,"vm_num_gt":0,"vm_num_gte":0,"vm_num_in":[0],"vm_num_lt":0,"vm_num_lte":0,"vm_num_not":0,"vm_num_not_in":[0],"vm_templates_every":"VmTemplateWhereInput","vm_templates_none":"VmTemplateWhereInput","vm_templates_some":"VmTemplateWhereInput","vms_every":"VmWhereInput","vms_none":"VmWhereInput","vms_some":"VmWhereInput","witness":"WitnessWhereInput","zones_every":"ZoneWhereInput","zones_none":"ZoneWhereInput","zones_some":"ZoneWhereInput"}}
//
// swagger:model GetClustersRequestBody
type GetClustersRequestBody struct {

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

// Validate validates this get clusters request body
func (m *GetClustersRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get clusters request body based on context it is used
func (m *GetClustersRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetClustersRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClustersRequestBody) UnmarshalBinary(b []byte) error {
	var res GetClustersRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
