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

// GetLogCollectionsRequestBody get log collections request body
// Example: {"after":"logCollections-id-string","before":"logCollections-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"LogCollectionWhereInput[]","NOT":"LogCollectionWhereInput[]","OR":"LogCollectionWhereInput[]","cluster":"ClusterWhereInput","hosts_every":"HostWhereInput","hosts_none":"HostWhereInput","hosts_some":"HostWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","log_ended_at":"string","log_ended_at_gt":"string","log_ended_at_gte":"string","log_ended_at_in":["string"],"log_ended_at_lt":"string","log_ended_at_lte":"string","log_ended_at_not":"string","log_ended_at_not_in":["string"],"log_started_at":"string","log_started_at_gt":"string","log_started_at_gte":"string","log_started_at_in":["string"],"log_started_at_lt":"string","log_started_at_lte":"string","log_started_at_not":"string","log_started_at_not_in":["string"],"owner":"string","owner_contains":"string","owner_ends_with":"string","owner_gt":"string","owner_gte":"string","owner_in":["string"],"owner_lt":"string","owner_lte":"string","owner_not":"string","owner_not_contains":"string","owner_not_ends_with":"string","owner_not_in":["string"],"owner_not_starts_with":"string","owner_starts_with":"string","path":"string","path_contains":"string","path_ends_with":"string","path_gt":"string","path_gte":"string","path_in":["string"],"path_lt":"string","path_lte":"string","path_not":"string","path_not_contains":"string","path_not_ends_with":"string","path_not_in":["string"],"path_not_starts_with":"string","path_starts_with":"string","progress":0,"progress_gt":0,"progress_gte":0,"progress_in":[0],"progress_lt":0,"progress_lte":0,"progress_not":0,"progress_not_in":[0],"size":0,"size_gt":0,"size_gte":0,"size_in":[0],"size_lt":0,"size_lte":0,"size_not":0,"size_not_in":[0],"started_at":"string","started_at_gt":"string","started_at_gte":"string","started_at_in":["string"],"started_at_lt":"string","started_at_lte":"string","started_at_not":"string","started_at_not_in":["string"],"status":"EXECUTING","status_in":["EXECUTING"],"status_not":"EXECUTING","status_not_in":["EXECUTING"],"witness":"WitnessWhereInput"}}
//
// swagger:model GetLogCollectionsRequestBody
type GetLogCollectionsRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *LogCollectionOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *LogCollectionWhereInput `json:"where,omitempty"`
}

// Validate validates this get log collections request body
func (m *GetLogCollectionsRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetLogCollectionsRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetLogCollectionsRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get log collections request body based on the context it is used
func (m *GetLogCollectionsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetLogCollectionsRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetLogCollectionsRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetLogCollectionsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLogCollectionsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetLogCollectionsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
