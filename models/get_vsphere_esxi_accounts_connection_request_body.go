// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVsphereEsxiAccountsConnectionRequestBody get vsphere esxi accounts connection request body
// Example: {"after":"vsphereEsxiAccountsConnection-id-string","before":"vsphereEsxiAccountsConnection-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"VsphereEsxiAccountWhereInput[]","NOT":"VsphereEsxiAccountWhereInput[]","OR":"VsphereEsxiAccountWhereInput[]","host":"HostWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","ip":"string","ip_contains":"string","ip_ends_with":"string","ip_gt":"string","ip_gte":"string","ip_in":["string"],"ip_lt":"string","ip_lte":"string","ip_not":"string","ip_not_contains":"string","ip_not_ends_with":"string","ip_not_in":["string"],"ip_not_starts_with":"string","ip_starts_with":"string","is_valid":false,"is_valid_not":false,"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","password":"string","password_contains":"string","password_ends_with":"string","password_gt":"string","password_gte":"string","password_in":["string"],"password_lt":"string","password_lte":"string","password_not":"string","password_not_contains":"string","password_not_ends_with":"string","password_not_in":["string"],"password_not_starts_with":"string","password_starts_with":"string","port":0,"port_gt":0,"port_gte":0,"port_in":[0],"port_lt":0,"port_lte":0,"port_not":0,"port_not_in":[0],"username":"string","username_contains":"string","username_ends_with":"string","username_gt":"string","username_gte":"string","username_in":["string"],"username_lt":"string","username_lte":"string","username_not":"string","username_not_contains":"string","username_not_ends_with":"string","username_not_in":["string"],"username_not_starts_with":"string","username_starts_with":"string"}}
//
// swagger:model GetVsphereEsxiAccountsConnectionRequestBody
type GetVsphereEsxiAccountsConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy interface{} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where interface{} `json:"where,omitempty"`
}

// Validate validates this get vsphere esxi accounts connection request body
func (m *GetVsphereEsxiAccountsConnectionRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get vsphere esxi accounts connection request body based on context it is used
func (m *GetVsphereEsxiAccountsConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetVsphereEsxiAccountsConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVsphereEsxiAccountsConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetVsphereEsxiAccountsConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
