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

// GetApplicationsConnectionRequestBody get applications connection request body
// Example: {"after":"applicationsConnection-id-string","before":"applicationsConnection-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"ApplicationWhereInput[]","NOT":"ApplicationWhereInput[]","OR":"ApplicationWhereInput[]","cluster":"ClusterWhereInput","error_message":"string","error_message_contains":"string","error_message_ends_with":"string","error_message_gt":"string","error_message_gte":"string","error_message_in":["string"],"error_message_lt":"string","error_message_lte":"string","error_message_not":"string","error_message_not_contains":"string","error_message_not_ends_with":"string","error_message_not_in":["string"],"error_message_not_starts_with":"string","error_message_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","image_name":"string","image_name_contains":"string","image_name_ends_with":"string","image_name_gt":"string","image_name_gte":"string","image_name_in":["string"],"image_name_lt":"string","image_name_lte":"string","image_name_not":"string","image_name_not_contains":"string","image_name_not_ends_with":"string","image_name_not_in":["string"],"image_name_not_starts_with":"string","image_name_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","memory":0,"memory_gt":0,"memory_gte":0,"memory_in":[0],"memory_lt":0,"memory_lte":0,"memory_not":0,"memory_not_in":[0],"state":"DEPLOY_ERROR","state_in":["DEPLOY_ERROR"],"state_not":"DEPLOY_ERROR","state_not_in":["DEPLOY_ERROR"],"storage_ip":"string","storage_ip_contains":"string","storage_ip_ends_with":"string","storage_ip_gt":"string","storage_ip_gte":"string","storage_ip_in":["string"],"storage_ip_lt":"string","storage_ip_lte":"string","storage_ip_not":"string","storage_ip_not_contains":"string","storage_ip_not_ends_with":"string","storage_ip_not_in":["string"],"storage_ip_not_starts_with":"string","storage_ip_starts_with":"string","type":"MONITOR","type_in":["MONITOR"],"type_not":"MONITOR","type_not_in":["MONITOR"],"update_time":"string","update_time_gt":"string","update_time_gte":"string","update_time_in":["string"],"update_time_lt":"string","update_time_lte":"string","update_time_not":"string","update_time_not_in":["string"],"vcpu":0,"vcpu_gt":0,"vcpu_gte":0,"vcpu_in":[0],"vcpu_lt":0,"vcpu_lte":0,"vcpu_not":0,"vcpu_not_in":[0],"version":"string","version_contains":"string","version_ends_with":"string","version_gt":"string","version_gte":"string","version_in":["string"],"version_lt":"string","version_lte":"string","version_not":"string","version_not_contains":"string","version_not_ends_with":"string","version_not_in":["string"],"version_not_starts_with":"string","version_starts_with":"string","vm":"VmWhereInput","volume_size":0,"volume_size_gt":0,"volume_size_gte":0,"volume_size_in":[0],"volume_size_lt":0,"volume_size_lte":0,"volume_size_not":0,"volume_size_not_in":[0]}}
//
// swagger:model GetApplicationsConnectionRequestBody
type GetApplicationsConnectionRequestBody struct {

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
		ApplicationOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		ApplicationWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get applications connection request body
func (m *GetApplicationsConnectionRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetApplicationsConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetApplicationsConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get applications connection request body based on the context it is used
func (m *GetApplicationsConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetApplicationsConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetApplicationsConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetApplicationsConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetApplicationsConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetApplicationsConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
