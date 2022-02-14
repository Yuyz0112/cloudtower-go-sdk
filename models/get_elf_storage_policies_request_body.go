// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetElfStoragePoliciesRequestBody get elf storage policies request body
// Example: {"after":"elfStoragePolicies-id-string","before":"elfStoragePolicies-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"ElfStoragePolicyWhereInput[]","NOT":"ElfStoragePolicyWhereInput[]","OR":"ElfStoragePolicyWhereInput[]","cluster":"ClusterWhereInput","description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","replica_num":0,"replica_num_gt":0,"replica_num_gte":0,"replica_num_in":[0],"replica_num_lt":0,"replica_num_lte":0,"replica_num_not":0,"replica_num_not_in":[0],"stripe_num":0,"stripe_num_gt":0,"stripe_num_gte":0,"stripe_num_in":[0],"stripe_num_lt":0,"stripe_num_lte":0,"stripe_num_not":0,"stripe_num_not_in":[0],"stripe_size":0,"stripe_size_gt":0,"stripe_size_gte":0,"stripe_size_in":[0],"stripe_size_lt":0,"stripe_size_lte":0,"stripe_size_not":0,"stripe_size_not_in":[0],"thin_provision":false,"thin_provision_not":false}}
//
// swagger:model GetElfStoragePoliciesRequestBody
type GetElfStoragePoliciesRequestBody struct {

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

// Validate validates this get elf storage policies request body
func (m *GetElfStoragePoliciesRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get elf storage policies request body based on context it is used
func (m *GetElfStoragePoliciesRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetElfStoragePoliciesRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetElfStoragePoliciesRequestBody) UnmarshalBinary(b []byte) error {
	var res GetElfStoragePoliciesRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
