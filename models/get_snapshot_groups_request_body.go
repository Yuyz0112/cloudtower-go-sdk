// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSnapshotGroupsRequestBody get snapshot groups request body
// Example: {"after":"snapshotGroups-id-string","before":"snapshotGroups-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"SnapshotGroupWhereInput[]","NOT":"SnapshotGroupWhereInput[]","OR":"SnapshotGroupWhereInput[]","cluster":"ClusterWhereInput","deleted":false,"deleted_not":false,"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"estimated_recycling_time":"string","estimated_recycling_time_gt":"string","estimated_recycling_time_gte":"string","estimated_recycling_time_in":["string"],"estimated_recycling_time_lt":"string","estimated_recycling_time_lte":"string","estimated_recycling_time_not":"string","estimated_recycling_time_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"keep":false,"keep_not":false,"local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","logical_size_bytes":0,"logical_size_bytes_gt":0,"logical_size_bytes_gte":0,"logical_size_bytes_in":[0],"logical_size_bytes_lt":0,"logical_size_bytes_lte":0,"logical_size_bytes_not":0,"logical_size_bytes_not_in":[0],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","object_num":0,"object_num_gt":0,"object_num_gte":0,"object_num_in":[0],"object_num_lt":0,"object_num_lte":0,"object_num_not":0,"object_num_not_in":[0],"snapshotPlanTask":"SnapshotPlanTaskWhereInput","vm_snapshots_every":"VmSnapshotWhereInput","vm_snapshots_none":"VmSnapshotWhereInput","vm_snapshots_some":"VmSnapshotWhereInput"}}
//
// swagger:model GetSnapshotGroupsRequestBody
type GetSnapshotGroupsRequestBody struct {

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

// Validate validates this get snapshot groups request body
func (m *GetSnapshotGroupsRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get snapshot groups request body based on context it is used
func (m *GetSnapshotGroupsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSnapshotGroupsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSnapshotGroupsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetSnapshotGroupsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
