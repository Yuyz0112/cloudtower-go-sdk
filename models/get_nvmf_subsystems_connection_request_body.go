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

// GetNvmfSubsystemsConnectionRequestBody get nvmf subsystems connection request body
// Example: {"after":"nvmfSubsystemsConnection-id-string","before":"nvmfSubsystemsConnection-id-string","first":0,"last":0,"orderBy":"bps_ASC","skip":0,"where":{"AND":"NvmfSubsystemWhereInput[]","NOT":"NvmfSubsystemWhereInput[]","OR":"NvmfSubsystemWhereInput[]","bps":0,"bps_gt":0,"bps_gte":0,"bps_in":[0],"bps_lt":0,"bps_lte":0,"bps_max":0,"bps_max_gt":0,"bps_max_gte":0,"bps_max_in":[0],"bps_max_length":0,"bps_max_length_gt":0,"bps_max_length_gte":0,"bps_max_length_in":[0],"bps_max_length_lt":0,"bps_max_length_lte":0,"bps_max_length_not":0,"bps_max_length_not_in":[0],"bps_max_lt":0,"bps_max_lte":0,"bps_max_not":0,"bps_max_not_in":[0],"bps_not":0,"bps_not_in":[0],"bps_rd":0,"bps_rd_gt":0,"bps_rd_gte":0,"bps_rd_in":[0],"bps_rd_lt":0,"bps_rd_lte":0,"bps_rd_max":0,"bps_rd_max_gt":0,"bps_rd_max_gte":0,"bps_rd_max_in":[0],"bps_rd_max_length":0,"bps_rd_max_length_gt":0,"bps_rd_max_length_gte":0,"bps_rd_max_length_in":[0],"bps_rd_max_length_lt":0,"bps_rd_max_length_lte":0,"bps_rd_max_length_not":0,"bps_rd_max_length_not_in":[0],"bps_rd_max_lt":0,"bps_rd_max_lte":0,"bps_rd_max_not":0,"bps_rd_max_not_in":[0],"bps_rd_not":0,"bps_rd_not_in":[0],"bps_wr":0,"bps_wr_gt":0,"bps_wr_gte":0,"bps_wr_in":[0],"bps_wr_lt":0,"bps_wr_lte":0,"bps_wr_max":0,"bps_wr_max_gt":0,"bps_wr_max_gte":0,"bps_wr_max_in":[0],"bps_wr_max_length":0,"bps_wr_max_length_gt":0,"bps_wr_max_length_gte":0,"bps_wr_max_length_in":[0],"bps_wr_max_length_lt":0,"bps_wr_max_length_lte":0,"bps_wr_max_length_not":0,"bps_wr_max_length_not_in":[0],"bps_wr_max_lt":0,"bps_wr_max_lte":0,"bps_wr_max_not":0,"bps_wr_max_not_in":[0],"bps_wr_not":0,"bps_wr_not_in":[0],"cluster":"ClusterWhereInput","description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"external_use":false,"external_use_not":false,"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"io_size":0,"io_size_gt":0,"io_size_gte":0,"io_size_in":[0],"io_size_lt":0,"io_size_lte":0,"io_size_not":0,"io_size_not_in":[0],"iops":0,"iops_gt":0,"iops_gte":0,"iops_in":[0],"iops_lt":0,"iops_lte":0,"iops_max":0,"iops_max_gt":0,"iops_max_gte":0,"iops_max_in":[0],"iops_max_length":0,"iops_max_length_gt":0,"iops_max_length_gte":0,"iops_max_length_in":[0],"iops_max_length_lt":0,"iops_max_length_lte":0,"iops_max_length_not":0,"iops_max_length_not_in":[0],"iops_max_lt":0,"iops_max_lte":0,"iops_max_not":0,"iops_max_not_in":[0],"iops_not":0,"iops_not_in":[0],"iops_rd":0,"iops_rd_gt":0,"iops_rd_gte":0,"iops_rd_in":[0],"iops_rd_lt":0,"iops_rd_lte":0,"iops_rd_max":0,"iops_rd_max_gt":0,"iops_rd_max_gte":0,"iops_rd_max_in":[0],"iops_rd_max_length":0,"iops_rd_max_length_gt":0,"iops_rd_max_length_gte":0,"iops_rd_max_length_in":[0],"iops_rd_max_length_lt":0,"iops_rd_max_length_lte":0,"iops_rd_max_length_not":0,"iops_rd_max_length_not_in":[0],"iops_rd_max_lt":0,"iops_rd_max_lte":0,"iops_rd_max_not":0,"iops_rd_max_not_in":[0],"iops_rd_not":0,"iops_rd_not_in":[0],"iops_wr":0,"iops_wr_gt":0,"iops_wr_gte":0,"iops_wr_in":[0],"iops_wr_lt":0,"iops_wr_lte":0,"iops_wr_max":0,"iops_wr_max_gt":0,"iops_wr_max_gte":0,"iops_wr_max_in":[0],"iops_wr_max_length":0,"iops_wr_max_length_gt":0,"iops_wr_max_length_gte":0,"iops_wr_max_length_in":[0],"iops_wr_max_length_lt":0,"iops_wr_max_length_lte":0,"iops_wr_max_length_not":0,"iops_wr_max_length_not_in":[0],"iops_wr_max_lt":0,"iops_wr_max_lte":0,"iops_wr_max_not":0,"iops_wr_max_not_in":[0],"iops_wr_not":0,"iops_wr_not_in":[0],"ip_whitelist":"string","ip_whitelist_contains":"string","ip_whitelist_ends_with":"string","ip_whitelist_gt":"string","ip_whitelist_gte":"string","ip_whitelist_in":["string"],"ip_whitelist_lt":"string","ip_whitelist_lte":"string","ip_whitelist_not":"string","ip_whitelist_not_contains":"string","ip_whitelist_not_ends_with":"string","ip_whitelist_not_in":["string"],"ip_whitelist_not_starts_with":"string","ip_whitelist_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","namespace_groups_every":"NamespaceGroupWhereInput","namespace_groups_none":"NamespaceGroupWhereInput","namespace_groups_some":"NamespaceGroupWhereInput","namespaces_every":"NvmfNamespaceWhereInput","namespaces_none":"NvmfNamespaceWhereInput","namespaces_some":"NvmfNamespaceWhereInput","nqn_name":"string","nqn_name_contains":"string","nqn_name_ends_with":"string","nqn_name_gt":"string","nqn_name_gte":"string","nqn_name_in":["string"],"nqn_name_lt":"string","nqn_name_lte":"string","nqn_name_not":"string","nqn_name_not_contains":"string","nqn_name_not_ends_with":"string","nqn_name_not_in":["string"],"nqn_name_not_starts_with":"string","nqn_name_starts_with":"string","nqn_whitelist":"string","nqn_whitelist_contains":"string","nqn_whitelist_ends_with":"string","nqn_whitelist_gt":"string","nqn_whitelist_gte":"string","nqn_whitelist_in":["string"],"nqn_whitelist_lt":"string","nqn_whitelist_lte":"string","nqn_whitelist_not":"string","nqn_whitelist_not_contains":"string","nqn_whitelist_not_ends_with":"string","nqn_whitelist_not_in":["string"],"nqn_whitelist_not_starts_with":"string","nqn_whitelist_starts_with":"string","policy":"BALANCE","policy_in":["BALANCE"],"policy_not":"BALANCE","policy_not_in":["BALANCE"],"replica_num":0,"replica_num_gt":0,"replica_num_gte":0,"replica_num_in":[0],"replica_num_lt":0,"replica_num_lte":0,"replica_num_not":0,"replica_num_not_in":[0],"stripe_num":0,"stripe_num_gt":0,"stripe_num_gte":0,"stripe_num_in":[0],"stripe_num_lt":0,"stripe_num_lte":0,"stripe_num_not":0,"stripe_num_not_in":[0],"stripe_size":0,"stripe_size_gt":0,"stripe_size_gte":0,"stripe_size_in":[0],"stripe_size_lt":0,"stripe_size_lte":0,"stripe_size_not":0,"stripe_size_not_in":[0],"thin_provision":false,"thin_provision_not":false}}
//
// swagger:model GetNvmfSubsystemsConnectionRequestBody
type GetNvmfSubsystemsConnectionRequestBody struct {

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
		NvmfSubsystemOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		NvmfSubsystemWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get nvmf subsystems connection request body
func (m *GetNvmfSubsystemsConnectionRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetNvmfSubsystemsConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetNvmfSubsystemsConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get nvmf subsystems connection request body based on the context it is used
func (m *GetNvmfSubsystemsConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetNvmfSubsystemsConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetNvmfSubsystemsConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetNvmfSubsystemsConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNvmfSubsystemsConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetNvmfSubsystemsConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
