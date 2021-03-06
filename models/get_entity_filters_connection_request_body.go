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

// GetEntityFiltersConnectionRequestBody get entity filters connection request body
// Example: {"after":"entityFiltersConnection-id-string","before":"entityFiltersConnection-id-string","first":0,"last":0,"orderBy":"apply_to_all_clusters_ASC","skip":0,"where":{"AND":"EntityFilterWhereInput[]","NOT":"EntityFilterWhereInput[]","OR":"EntityFilterWhereInput[]","apply_to_all_clusters":false,"apply_to_all_clusters_not":false,"clusters_every":"ClusterWhereInput","clusters_none":"ClusterWhereInput","clusters_some":"ClusterWhereInput","entity_type":"VM","entity_type_in":["VM"],"entity_type_not":"VM","entity_type_not_in":["VM"],"exec_failed_cluster_every":"ClusterWhereInput","exec_failed_cluster_none":"ClusterWhereInput","exec_failed_cluster_some":"ClusterWhereInput","filter_status":"EXECUTING","filter_status_in":["EXECUTING"],"filter_status_not":"EXECUTING","filter_status_not_in":["EXECUTING"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","last_executed_at":"string","last_executed_at_gt":"string","last_executed_at_gte":"string","last_executed_at_in":["string"],"last_executed_at_lt":"string","last_executed_at_lte":"string","last_executed_at_not":"string","last_executed_at_not_in":["string"],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","preset":"string","preset_contains":"string","preset_ends_with":"string","preset_gt":"string","preset_gte":"string","preset_in":["string"],"preset_lt":"string","preset_lte":"string","preset_not":"string","preset_not_contains":"string","preset_not_ends_with":"string","preset_not_in":["string"],"preset_not_starts_with":"string","preset_starts_with":"string"}}
//
// swagger:model GetEntityFiltersConnectionRequestBody
type GetEntityFiltersConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *EntityFilterOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *EntityFilterWhereInput `json:"where,omitempty"`
}

// Validate validates this get entity filters connection request body
func (m *GetEntityFiltersConnectionRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetEntityFiltersConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetEntityFiltersConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get entity filters connection request body based on the context it is used
func (m *GetEntityFiltersConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetEntityFiltersConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetEntityFiltersConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetEntityFiltersConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEntityFiltersConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetEntityFiltersConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
