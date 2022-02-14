// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetPmemDimmsRequestBody get pmem dimms request body
// Example: {"after":"pmemDimms-id-string","before":"pmemDimms-id-string","first":0,"last":0,"orderBy":"capacity_ASC","skip":0,"where":{"AND":"PmemDimmWhereInput[]","NOT":"PmemDimmWhereInput[]","OR":"PmemDimmWhereInput[]","capacity":0,"capacity_gt":0,"capacity_gte":0,"capacity_in":[0],"capacity_lt":0,"capacity_lte":0,"capacity_not":0,"capacity_not_in":[0],"device_locator":"string","device_locator_contains":"string","device_locator_ends_with":"string","device_locator_gt":"string","device_locator_gte":"string","device_locator_in":["string"],"device_locator_lt":"string","device_locator_lte":"string","device_locator_not":"string","device_locator_not_contains":"string","device_locator_not_ends_with":"string","device_locator_not_in":["string"],"device_locator_not_starts_with":"string","device_locator_starts_with":"string","disk":"DiskWhereInput","health_status":"HEALTHY","health_status_in":["HEALTHY"],"health_status_not":"HEALTHY","health_status_not_in":["HEALTHY"],"host":"HostWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","numa_node":0,"numa_node_gt":0,"numa_node_gte":0,"numa_node_in":[0],"numa_node_lt":0,"numa_node_lte":0,"numa_node_not":0,"numa_node_not_in":[0],"part_number":"string","part_number_contains":"string","part_number_ends_with":"string","part_number_gt":"string","part_number_gte":"string","part_number_in":["string"],"part_number_lt":"string","part_number_lte":"string","part_number_not":"string","part_number_not_contains":"string","part_number_not_ends_with":"string","part_number_not_in":["string"],"part_number_not_starts_with":"string","part_number_starts_with":"string","remaining_life_percent":0,"remaining_life_percent_gt":0,"remaining_life_percent_gte":0,"remaining_life_percent_in":[0],"remaining_life_percent_lt":0,"remaining_life_percent_lte":0,"remaining_life_percent_not":0,"remaining_life_percent_not_in":[0],"version":"string","version_contains":"string","version_ends_with":"string","version_gt":"string","version_gte":"string","version_in":["string"],"version_lt":"string","version_lte":"string","version_not":"string","version_not_contains":"string","version_not_ends_with":"string","version_not_in":["string"],"version_not_starts_with":"string","version_starts_with":"string"}}
//
// swagger:model GetPmemDimmsRequestBody
type GetPmemDimmsRequestBody struct {

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

// Validate validates this get pmem dimms request body
func (m *GetPmemDimmsRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get pmem dimms request body based on context it is used
func (m *GetPmemDimmsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetPmemDimmsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetPmemDimmsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetPmemDimmsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
