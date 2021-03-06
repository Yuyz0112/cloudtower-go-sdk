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

// GetTasksRequestBody get tasks request body
// Example: {"after":"tasks-id-string","before":"tasks-id-string","first":0,"last":0,"orderBy":"args_ASC","skip":0,"where":{"AND":"TaskWhereInput[]","NOT":"TaskWhereInput[]","OR":"TaskWhereInput[]","cluster":"ClusterWhereInput","description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","error_code":"string","error_code_contains":"string","error_code_ends_with":"string","error_code_gt":"string","error_code_gte":"string","error_code_in":["string"],"error_code_lt":"string","error_code_lte":"string","error_code_not":"string","error_code_not_contains":"string","error_code_not_ends_with":"string","error_code_not_in":["string"],"error_code_not_starts_with":"string","error_code_starts_with":"string","error_message":"string","error_message_contains":"string","error_message_ends_with":"string","error_message_gt":"string","error_message_gte":"string","error_message_in":["string"],"error_message_lt":"string","error_message_lte":"string","error_message_not":"string","error_message_not_contains":"string","error_message_not_ends_with":"string","error_message_not_in":["string"],"error_message_not_starts_with":"string","error_message_starts_with":"string","finished_at":"string","finished_at_gt":"string","finished_at_gte":"string","finished_at_in":["string"],"finished_at_lt":"string","finished_at_lte":"string","finished_at_not":"string","finished_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"progress":0,"progress_gt":0,"progress_gte":0,"progress_in":[0],"progress_lt":0,"progress_lte":0,"progress_not":0,"progress_not_in":[0],"resource_id":"string","resource_id_contains":"string","resource_id_ends_with":"string","resource_id_gt":"string","resource_id_gte":"string","resource_id_in":["string"],"resource_id_lt":"string","resource_id_lte":"string","resource_id_not":"string","resource_id_not_contains":"string","resource_id_not_ends_with":"string","resource_id_not_in":["string"],"resource_id_not_starts_with":"string","resource_id_starts_with":"string","resource_mutation":"string","resource_mutation_contains":"string","resource_mutation_ends_with":"string","resource_mutation_gt":"string","resource_mutation_gte":"string","resource_mutation_in":["string"],"resource_mutation_lt":"string","resource_mutation_lte":"string","resource_mutation_not":"string","resource_mutation_not_contains":"string","resource_mutation_not_ends_with":"string","resource_mutation_not_in":["string"],"resource_mutation_not_starts_with":"string","resource_mutation_starts_with":"string","resource_rollback_error":"string","resource_rollback_error_contains":"string","resource_rollback_error_ends_with":"string","resource_rollback_error_gt":"string","resource_rollback_error_gte":"string","resource_rollback_error_in":["string"],"resource_rollback_error_lt":"string","resource_rollback_error_lte":"string","resource_rollback_error_not":"string","resource_rollback_error_not_contains":"string","resource_rollback_error_not_ends_with":"string","resource_rollback_error_not_in":["string"],"resource_rollback_error_not_starts_with":"string","resource_rollback_error_starts_with":"string","resource_rollback_retry_count":0,"resource_rollback_retry_count_gt":0,"resource_rollback_retry_count_gte":0,"resource_rollback_retry_count_in":[0],"resource_rollback_retry_count_lt":0,"resource_rollback_retry_count_lte":0,"resource_rollback_retry_count_not":0,"resource_rollback_retry_count_not_in":[0],"resource_rollbacked":false,"resource_rollbacked_not":false,"resource_type":"string","resource_type_contains":"string","resource_type_ends_with":"string","resource_type_gt":"string","resource_type_gte":"string","resource_type_in":["string"],"resource_type_lt":"string","resource_type_lte":"string","resource_type_not":"string","resource_type_not_contains":"string","resource_type_not_ends_with":"string","resource_type_not_in":["string"],"resource_type_not_starts_with":"string","resource_type_starts_with":"string","snapshot":"string","snapshot_contains":"string","snapshot_ends_with":"string","snapshot_gt":"string","snapshot_gte":"string","snapshot_in":["string"],"snapshot_lt":"string","snapshot_lte":"string","snapshot_not":"string","snapshot_not_contains":"string","snapshot_not_ends_with":"string","snapshot_not_in":["string"],"snapshot_not_starts_with":"string","snapshot_starts_with":"string","started_at":"string","started_at_gt":"string","started_at_gte":"string","started_at_in":["string"],"started_at_lt":"string","started_at_lte":"string","started_at_not":"string","started_at_not_in":["string"],"status":"EXECUTING","status_in":["EXECUTING"],"status_not":"EXECUTING","status_not_in":["EXECUTING"],"user":"UserWhereInput"}}
//
// swagger:model GetTasksRequestBody
type GetTasksRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *TaskOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *TaskWhereInput `json:"where,omitempty"`
}

// Validate validates this get tasks request body
func (m *GetTasksRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetTasksRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetTasksRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get tasks request body based on the context it is used
func (m *GetTasksRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetTasksRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetTasksRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetTasksRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTasksRequestBody) UnmarshalBinary(b []byte) error {
	var res GetTasksRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
