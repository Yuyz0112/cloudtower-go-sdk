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

// GetDatacentersRequestBody get datacenters request body
// Example: {"after":"datacenters-id-string","before":"datacenters-id-string","first":0,"last":0,"orderBy":"cluster_num_ASC","skip":0,"where":{"AND":"DatacenterWhereInput[]","NOT":"DatacenterWhereInput[]","OR":"DatacenterWhereInput[]","cluster_num":0,"cluster_num_gt":0,"cluster_num_gte":0,"cluster_num_in":[0],"cluster_num_lt":0,"cluster_num_lte":0,"cluster_num_not":0,"cluster_num_not_in":[0],"clusters_every":"ClusterWhereInput","clusters_none":"ClusterWhereInput","clusters_some":"ClusterWhereInput","failure_data_space":0,"failure_data_space_gt":0,"failure_data_space_gte":0,"failure_data_space_in":[0],"failure_data_space_lt":0,"failure_data_space_lte":0,"failure_data_space_not":0,"failure_data_space_not_in":[0],"host_num":0,"host_num_gt":0,"host_num_gte":0,"host_num_in":[0],"host_num_lt":0,"host_num_lte":0,"host_num_not":0,"host_num_not_in":[0],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","organization":"OrganizationWhereInput","total_cpu_hz":0,"total_cpu_hz_gt":0,"total_cpu_hz_gte":0,"total_cpu_hz_in":[0],"total_cpu_hz_lt":0,"total_cpu_hz_lte":0,"total_cpu_hz_not":0,"total_cpu_hz_not_in":[0],"total_data_capacity":0,"total_data_capacity_gt":0,"total_data_capacity_gte":0,"total_data_capacity_in":[0],"total_data_capacity_lt":0,"total_data_capacity_lte":0,"total_data_capacity_not":0,"total_data_capacity_not_in":[0],"total_memory_bytes":0,"total_memory_bytes_gt":0,"total_memory_bytes_gte":0,"total_memory_bytes_in":[0],"total_memory_bytes_lt":0,"total_memory_bytes_lte":0,"total_memory_bytes_not":0,"total_memory_bytes_not_in":[0],"used_cpu_hz":0,"used_cpu_hz_gt":0,"used_cpu_hz_gte":0,"used_cpu_hz_in":[0],"used_cpu_hz_lt":0,"used_cpu_hz_lte":0,"used_cpu_hz_not":0,"used_cpu_hz_not_in":[0],"used_data_space":0,"used_data_space_gt":0,"used_data_space_gte":0,"used_data_space_in":[0],"used_data_space_lt":0,"used_data_space_lte":0,"used_data_space_not":0,"used_data_space_not_in":[0],"used_memory_bytes":0,"used_memory_bytes_gt":0,"used_memory_bytes_gte":0,"used_memory_bytes_in":[0],"used_memory_bytes_lt":0,"used_memory_bytes_lte":0,"used_memory_bytes_not":0,"used_memory_bytes_not_in":[0],"vm_num":0,"vm_num_gt":0,"vm_num_gte":0,"vm_num_in":[0],"vm_num_lt":0,"vm_num_lte":0,"vm_num_not":0,"vm_num_not_in":[0]}}
//
// swagger:model GetDatacentersRequestBody
type GetDatacentersRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *DatacenterOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *DatacenterWhereInput `json:"where,omitempty"`
}

// Validate validates this get datacenters request body
func (m *GetDatacentersRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetDatacentersRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	if m.OrderBy != nil {
		if err := m.OrderBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetDatacentersRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get datacenters request body based on the context it is used
func (m *GetDatacentersRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetDatacentersRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderBy != nil {
		if err := m.OrderBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetDatacentersRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDatacentersRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDatacentersRequestBody) UnmarshalBinary(b []byte) error {
	var res GetDatacentersRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
