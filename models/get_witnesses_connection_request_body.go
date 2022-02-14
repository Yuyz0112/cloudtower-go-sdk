// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetWitnessesConnectionRequestBody get witnesses connection request body
// Example: {"after":"witnessesConnection-id-string","before":"witnessesConnection-id-string","first":0,"last":0,"orderBy":"cpu_hz_per_core_ASC","skip":0,"where":{"AND":"WitnessWhereInput[]","NOT":"WitnessWhereInput[]","OR":"WitnessWhereInput[]","cluster":"ClusterWhereInput","cpu_hz_per_core":0,"cpu_hz_per_core_gt":0,"cpu_hz_per_core_gte":0,"cpu_hz_per_core_in":[0],"cpu_hz_per_core_lt":0,"cpu_hz_per_core_lte":0,"cpu_hz_per_core_not":0,"cpu_hz_per_core_not_in":[0],"data_ip":"string","data_ip_contains":"string","data_ip_ends_with":"string","data_ip_gt":"string","data_ip_gte":"string","data_ip_in":["string"],"data_ip_lt":"string","data_ip_lte":"string","data_ip_not":"string","data_ip_not_contains":"string","data_ip_not_ends_with":"string","data_ip_not_in":["string"],"data_ip_not_starts_with":"string","data_ip_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","system_data_capacity":0,"system_data_capacity_gt":0,"system_data_capacity_gte":0,"system_data_capacity_in":[0],"system_data_capacity_lt":0,"system_data_capacity_lte":0,"system_data_capacity_not":0,"system_data_capacity_not_in":[0],"system_used_data_space":0,"system_used_data_space_gt":0,"system_used_data_space_gte":0,"system_used_data_space_in":[0],"system_used_data_space_lt":0,"system_used_data_space_lte":0,"system_used_data_space_not":0,"system_used_data_space_not_in":[0],"total_cpu_cores":0,"total_cpu_cores_gt":0,"total_cpu_cores_gte":0,"total_cpu_cores_in":[0],"total_cpu_cores_lt":0,"total_cpu_cores_lte":0,"total_cpu_cores_not":0,"total_cpu_cores_not_in":[0],"total_cpu_hz":0,"total_cpu_hz_gt":0,"total_cpu_hz_gte":0,"total_cpu_hz_in":[0],"total_cpu_hz_lt":0,"total_cpu_hz_lte":0,"total_cpu_hz_not":0,"total_cpu_hz_not_in":[0],"total_memory_bytes":0,"total_memory_bytes_gt":0,"total_memory_bytes_gte":0,"total_memory_bytes_in":[0],"total_memory_bytes_lt":0,"total_memory_bytes_lte":0,"total_memory_bytes_not":0,"total_memory_bytes_not_in":[0]}}
//
// swagger:model GetWitnessesConnectionRequestBody
type GetWitnessesConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy struct {
		WitnessOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		WitnessWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get witnesses connection request body
func (m *GetWitnessesConnectionRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWitnessesConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetWitnessesConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get witnesses connection request body based on the context it is used
func (m *GetWitnessesConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWitnessesConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetWitnessesConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetWitnessesConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWitnessesConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetWitnessesConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
