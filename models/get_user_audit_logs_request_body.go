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

// GetUserAuditLogsRequestBody get user audit logs request body
// Example: {"after":"userAuditLogs-id-string","before":"userAuditLogs-id-string","first":0,"last":0,"orderBy":"action_ASC","skip":0,"where":{"AND":"UserAuditLogWhereInput[]","NOT":"UserAuditLogWhereInput[]","OR":"UserAuditLogWhereInput[]","action":"string","action_contains":"string","action_ends_with":"string","action_gt":"string","action_gte":"string","action_in":["string"],"action_lt":"string","action_lte":"string","action_not":"string","action_not_contains":"string","action_not_ends_with":"string","action_not_in":["string"],"action_not_starts_with":"string","action_starts_with":"string","cluster":"ClusterWhereInput","createdAt":"string","createdAt_gt":"string","createdAt_gte":"string","createdAt_in":["string"],"createdAt_lt":"string","createdAt_lte":"string","createdAt_not":"string","createdAt_not_in":["string"],"finished_at":"string","finished_at_gt":"string","finished_at_gte":"string","finished_at_in":["string"],"finished_at_lt":"string","finished_at_lte":"string","finished_at_not":"string","finished_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","ip_address":"string","ip_address_contains":"string","ip_address_ends_with":"string","ip_address_gt":"string","ip_address_gte":"string","ip_address_in":["string"],"ip_address_lt":"string","ip_address_lte":"string","ip_address_not":"string","ip_address_not_contains":"string","ip_address_not_ends_with":"string","ip_address_not_in":["string"],"ip_address_not_starts_with":"string","ip_address_starts_with":"string","message":"string","message_contains":"string","message_ends_with":"string","message_gt":"string","message_gte":"string","message_in":["string"],"message_lt":"string","message_lte":"string","message_not":"string","message_not_contains":"string","message_not_ends_with":"string","message_not_in":["string"],"message_not_starts_with":"string","message_starts_with":"string","resource_id":"string","resource_id_contains":"string","resource_id_ends_with":"string","resource_id_gt":"string","resource_id_gte":"string","resource_id_in":["string"],"resource_id_lt":"string","resource_id_lte":"string","resource_id_not":"string","resource_id_not_contains":"string","resource_id_not_ends_with":"string","resource_id_not_in":["string"],"resource_id_not_starts_with":"string","resource_id_starts_with":"string","resource_type":"string","resource_type_contains":"string","resource_type_ends_with":"string","resource_type_gt":"string","resource_type_gte":"string","resource_type_in":["string"],"resource_type_lt":"string","resource_type_lte":"string","resource_type_not":"string","resource_type_not_contains":"string","resource_type_not_ends_with":"string","resource_type_not_in":["string"],"resource_type_not_starts_with":"string","resource_type_starts_with":"string","status":"FAILED","status_in":["FAILED"],"status_not":"FAILED","status_not_in":["FAILED"],"user":"UserWhereInput"}}
//
// swagger:model GetUserAuditLogsRequestBody
type GetUserAuditLogsRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *UserAuditLogOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *UserAuditLogWhereInput `json:"where,omitempty"`
}

// Validate validates this get user audit logs request body
func (m *GetUserAuditLogsRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetUserAuditLogsRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetUserAuditLogsRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get user audit logs request body based on the context it is used
func (m *GetUserAuditLogsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetUserAuditLogsRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetUserAuditLogsRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetUserAuditLogsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserAuditLogsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetUserAuditLogsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
