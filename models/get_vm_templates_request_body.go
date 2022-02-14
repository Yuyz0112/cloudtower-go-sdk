// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVMTemplatesRequestBody get Vm templates request body
// Example: {"after":"vmTemplates-id-string","before":"vmTemplates-id-string","first":0,"last":0,"orderBy":"clock_offset_ASC","skip":0,"where":{"AND":"VmTemplateWhereInput[]","NOT":"VmTemplateWhereInput[]","OR":"VmTemplateWhereInput[]","clock_offset":"LOCALTIME","clock_offset_in":["LOCALTIME"],"clock_offset_not":"LOCALTIME","clock_offset_not_in":["LOCALTIME"],"cloud_init_supported":false,"cloud_init_supported_not":false,"cluster":"ClusterWhereInput","content_library_vm_template":"ContentLibraryVmTemplateWhereInput","cpu_model":"string","cpu_model_contains":"string","cpu_model_ends_with":"string","cpu_model_gt":"string","cpu_model_gte":"string","cpu_model_in":["string"],"cpu_model_lt":"string","cpu_model_lte":"string","cpu_model_not":"string","cpu_model_not_contains":"string","cpu_model_not_ends_with":"string","cpu_model_not_in":["string"],"cpu_model_not_starts_with":"string","cpu_model_starts_with":"string","description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"firmware":"BIOS","firmware_in":["BIOS"],"firmware_not":"BIOS","firmware_not_in":["BIOS"],"ha":false,"ha_not":false,"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","io_policy":"RESTRICT_EACH_DISK","io_policy_in":["RESTRICT_EACH_DISK"],"io_policy_not":"RESTRICT_EACH_DISK","io_policy_not_in":["RESTRICT_EACH_DISK"],"labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","max_bandwidth":0,"max_bandwidth_gt":0,"max_bandwidth_gte":0,"max_bandwidth_in":[0],"max_bandwidth_lt":0,"max_bandwidth_lte":0,"max_bandwidth_not":0,"max_bandwidth_not_in":[0],"max_bandwidth_policy":"DYNAMIC","max_bandwidth_policy_in":["DYNAMIC"],"max_bandwidth_policy_not":"DYNAMIC","max_bandwidth_policy_not_in":["DYNAMIC"],"max_iops":0,"max_iops_gt":0,"max_iops_gte":0,"max_iops_in":[0],"max_iops_lt":0,"max_iops_lte":0,"max_iops_not":0,"max_iops_not_in":[0],"max_iops_policy":"DYNAMIC","max_iops_policy_in":["DYNAMIC"],"max_iops_policy_not":"DYNAMIC","max_iops_policy_not_in":["DYNAMIC"],"memory":0,"memory_gt":0,"memory_gte":0,"memory_in":[0],"memory_lt":0,"memory_lte":0,"memory_not":0,"memory_not_in":[0],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","size":0,"size_gt":0,"size_gte":0,"size_in":[0],"size_lt":0,"size_lte":0,"size_not":0,"size_not_in":[0],"vcpu":0,"vcpu_gt":0,"vcpu_gte":0,"vcpu_in":[0],"vcpu_lt":0,"vcpu_lte":0,"vcpu_not":0,"vcpu_not_in":[0],"video_type":"string","video_type_contains":"string","video_type_ends_with":"string","video_type_gt":"string","video_type_gte":"string","video_type_in":["string"],"video_type_lt":"string","video_type_lte":"string","video_type_not":"string","video_type_not_contains":"string","video_type_not_ends_with":"string","video_type_not_in":["string"],"video_type_not_starts_with":"string","video_type_starts_with":"string","win_opt":false,"win_opt_not":false}}
//
// swagger:model GetVmTemplatesRequestBody
type GetVMTemplatesRequestBody struct {

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

// Validate validates this get Vm templates request body
func (m *GetVMTemplatesRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get Vm templates request body based on context it is used
func (m *GetVMTemplatesRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetVMTemplatesRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVMTemplatesRequestBody) UnmarshalBinary(b []byte) error {
	var res GetVMTemplatesRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
