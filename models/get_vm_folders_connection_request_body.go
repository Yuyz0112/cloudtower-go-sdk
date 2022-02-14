// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVMFoldersConnectionRequestBody get Vm folders connection request body
// Example: {"after":"vmFoldersConnection-id-string","before":"vmFoldersConnection-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"VmFolderWhereInput[]","NOT":"VmFolderWhereInput[]","OR":"VmFolderWhereInput[]","cluster":"ClusterWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","vm_num":0,"vm_num_gt":0,"vm_num_gte":0,"vm_num_in":[0],"vm_num_lt":0,"vm_num_lte":0,"vm_num_not":0,"vm_num_not_in":[0],"vms_every":"VmWhereInput","vms_none":"VmWhereInput","vms_some":"VmWhereInput"}}
//
// swagger:model GetVmFoldersConnectionRequestBody
type GetVMFoldersConnectionRequestBody struct {

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

// Validate validates this get Vm folders connection request body
func (m *GetVMFoldersConnectionRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get Vm folders connection request body based on context it is used
func (m *GetVMFoldersConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetVMFoldersConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVMFoldersConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetVMFoldersConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
