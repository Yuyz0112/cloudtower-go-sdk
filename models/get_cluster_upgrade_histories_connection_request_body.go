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

// GetClusterUpgradeHistoriesConnectionRequestBody get cluster upgrade histories connection request body
// Example: {"after":"clusterUpgradeHistoriesConnection-id-string","before":"clusterUpgradeHistoriesConnection-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"ClusterUpgradeHistoryWhereInput[]","NOT":"ClusterUpgradeHistoryWhereInput[]","OR":"ClusterUpgradeHistoryWhereInput[]","cluster":"ClusterWhereInput","date":"string","date_gt":"string","date_gte":"string","date_in":["string"],"date_lt":"string","date_lte":"string","date_not":"string","date_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","result":"string","result_contains":"string","result_ends_with":"string","result_gt":"string","result_gte":"string","result_in":["string"],"result_lt":"string","result_lte":"string","result_not":"string","result_not_contains":"string","result_not_ends_with":"string","result_not_in":["string"],"result_not_starts_with":"string","result_starts_with":"string","task_uuid":"string","task_uuid_contains":"string","task_uuid_ends_with":"string","task_uuid_gt":"string","task_uuid_gte":"string","task_uuid_in":["string"],"task_uuid_lt":"string","task_uuid_lte":"string","task_uuid_not":"string","task_uuid_not_contains":"string","task_uuid_not_ends_with":"string","task_uuid_not_in":["string"],"task_uuid_not_starts_with":"string","task_uuid_starts_with":"string","version":"string","version_contains":"string","version_ends_with":"string","version_gt":"string","version_gte":"string","version_in":["string"],"version_lt":"string","version_lte":"string","version_not":"string","version_not_contains":"string","version_not_ends_with":"string","version_not_in":["string"],"version_not_starts_with":"string","version_starts_with":"string"}}
//
// swagger:model GetClusterUpgradeHistoriesConnectionRequestBody
type GetClusterUpgradeHistoriesConnectionRequestBody struct {

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
		ClusterUpgradeHistoryOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		ClusterUpgradeHistoryWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get cluster upgrade histories connection request body
func (m *GetClusterUpgradeHistoriesConnectionRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetClusterUpgradeHistoriesConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetClusterUpgradeHistoriesConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get cluster upgrade histories connection request body based on the context it is used
func (m *GetClusterUpgradeHistoriesConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetClusterUpgradeHistoriesConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetClusterUpgradeHistoriesConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetClusterUpgradeHistoriesConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClusterUpgradeHistoriesConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetClusterUpgradeHistoriesConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
