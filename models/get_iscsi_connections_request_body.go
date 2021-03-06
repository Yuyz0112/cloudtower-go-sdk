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

// GetIscsiConnectionsRequestBody get iscsi connections request body
// Example: {"after":"iscsiConnections-id-string","before":"iscsiConnections-id-string","first":0,"last":0,"orderBy":"client_port_ASC","skip":0,"where":{"AND":"IscsiConnectionWhereInput[]","NOT":"IscsiConnectionWhereInput[]","OR":"IscsiConnectionWhereInput[]","client_port":0,"client_port_gt":0,"client_port_gte":0,"client_port_in":[0],"client_port_lt":0,"client_port_lte":0,"client_port_not":0,"client_port_not_in":[0],"cluster":"ClusterWhereInput","host":"HostWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","initiator_ip":"string","initiator_ip_contains":"string","initiator_ip_ends_with":"string","initiator_ip_gt":"string","initiator_ip_gte":"string","initiator_ip_in":["string"],"initiator_ip_lt":"string","initiator_ip_lte":"string","initiator_ip_not":"string","initiator_ip_not_contains":"string","initiator_ip_not_ends_with":"string","initiator_ip_not_in":["string"],"initiator_ip_not_starts_with":"string","initiator_ip_starts_with":"string","iscsi_target":"IscsiTargetWhereInput","nvmf_subsystem":"NvmfSubsystemWhereInput","type":"ISCSI","type_in":["ISCSI"],"type_not":"ISCSI","type_not_in":["ISCSI"]}}
//
// swagger:model GetIscsiConnectionsRequestBody
type GetIscsiConnectionsRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *IscsiConnectionOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *IscsiConnectionWhereInput `json:"where,omitempty"`
}

// Validate validates this get iscsi connections request body
func (m *GetIscsiConnectionsRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetIscsiConnectionsRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetIscsiConnectionsRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get iscsi connections request body based on the context it is used
func (m *GetIscsiConnectionsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetIscsiConnectionsRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetIscsiConnectionsRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetIscsiConnectionsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIscsiConnectionsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetIscsiConnectionsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
