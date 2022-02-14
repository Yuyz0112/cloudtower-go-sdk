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

// GetVdsesRequestBody get vdses request body
// Example: {"after":"vdses-id-string","before":"vdses-id-string","first":0,"last":0,"orderBy":"bond_mode_ASC","skip":0,"where":{"AND":"VdsWhereInput[]","NOT":"VdsWhereInput[]","OR":"VdsWhereInput[]","bond_mode":"string","bond_mode_contains":"string","bond_mode_ends_with":"string","bond_mode_gt":"string","bond_mode_gte":"string","bond_mode_in":["string"],"bond_mode_lt":"string","bond_mode_lte":"string","bond_mode_not":"string","bond_mode_not_contains":"string","bond_mode_not_ends_with":"string","bond_mode_not_in":["string"],"bond_mode_not_starts_with":"string","bond_mode_starts_with":"string","cluster":"ClusterWhereInput","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"everoute_cluster":"EverouteClusterWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","nics_every":"NicWhereInput","nics_none":"NicWhereInput","nics_some":"NicWhereInput","ovsbr_name":"string","ovsbr_name_contains":"string","ovsbr_name_ends_with":"string","ovsbr_name_gt":"string","ovsbr_name_gte":"string","ovsbr_name_in":["string"],"ovsbr_name_lt":"string","ovsbr_name_lte":"string","ovsbr_name_not":"string","ovsbr_name_not_contains":"string","ovsbr_name_not_ends_with":"string","ovsbr_name_not_in":["string"],"ovsbr_name_not_starts_with":"string","ovsbr_name_starts_with":"string","type":"ACCESS","type_in":["ACCESS"],"type_not":"ACCESS","type_not_in":["ACCESS"],"vlans_every":"VlanWhereInput","vlans_none":"VlanWhereInput","vlans_num":0,"vlans_num_gt":0,"vlans_num_gte":0,"vlans_num_in":[0],"vlans_num_lt":0,"vlans_num_lte":0,"vlans_num_not":0,"vlans_num_not_in":[0],"vlans_some":"VlanWhereInput","work_mode":"string","work_mode_contains":"string","work_mode_ends_with":"string","work_mode_gt":"string","work_mode_gte":"string","work_mode_in":["string"],"work_mode_lt":"string","work_mode_lte":"string","work_mode_not":"string","work_mode_not_contains":"string","work_mode_not_ends_with":"string","work_mode_not_in":["string"],"work_mode_not_starts_with":"string","work_mode_starts_with":"string"}}
//
// swagger:model GetVdsesRequestBody
type GetVdsesRequestBody struct {

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
		VdsOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		VdsWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get vdses request body
func (m *GetVdsesRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetVdsesRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetVdsesRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get vdses request body based on the context it is used
func (m *GetVdsesRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetVdsesRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetVdsesRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetVdsesRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVdsesRequestBody) UnmarshalBinary(b []byte) error {
	var res GetVdsesRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
