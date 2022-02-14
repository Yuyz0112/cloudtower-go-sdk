// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetUsersRequestBody get users request body
// Example: {"after":"users-id-string","before":"users-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"UserWhereInput[]","NOT":"UserWhereInput[]","OR":"UserWhereInput[]","email_address":"string","email_address_contains":"string","email_address_ends_with":"string","email_address_gt":"string","email_address_gte":"string","email_address_in":["string"],"email_address_lt":"string","email_address_lte":"string","email_address_not":"string","email_address_not_contains":"string","email_address_not_ends_with":"string","email_address_not_in":["string"],"email_address_not_starts_with":"string","email_address_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"ldap_dn":"string","ldap_dn_contains":"string","ldap_dn_ends_with":"string","ldap_dn_gt":"string","ldap_dn_gte":"string","ldap_dn_in":["string"],"ldap_dn_lt":"string","ldap_dn_lte":"string","ldap_dn_not":"string","ldap_dn_not_contains":"string","ldap_dn_not_ends_with":"string","ldap_dn_not_in":["string"],"ldap_dn_not_starts_with":"string","ldap_dn_starts_with":"string","mobile_phone":"string","mobile_phone_contains":"string","mobile_phone_ends_with":"string","mobile_phone_gt":"string","mobile_phone_gte":"string","mobile_phone_in":["string"],"mobile_phone_lt":"string","mobile_phone_lte":"string","mobile_phone_not":"string","mobile_phone_not_contains":"string","mobile_phone_not_ends_with":"string","mobile_phone_not_in":["string"],"mobile_phone_not_starts_with":"string","mobile_phone_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","password":"string","password_contains":"string","password_ends_with":"string","password_gt":"string","password_gte":"string","password_in":["string"],"password_lt":"string","password_lte":"string","password_not":"string","password_not_contains":"string","password_not_ends_with":"string","password_not_in":["string"],"password_not_starts_with":"string","password_starts_with":"string","role":"ADMIN","role_in":["ADMIN"],"role_not":"ADMIN","role_not_in":["ADMIN"],"roles_every":"UserRoleNextWhereInput","roles_none":"UserRoleNextWhereInput","roles_some":"UserRoleNextWhereInput","source":"LDAP","source_in":["LDAP"],"source_not":"LDAP","source_not_in":["LDAP"],"username":"string","username_contains":"string","username_ends_with":"string","username_gt":"string","username_gte":"string","username_in":["string"],"username_lt":"string","username_lte":"string","username_not":"string","username_not_contains":"string","username_not_ends_with":"string","username_not_in":["string"],"username_not_starts_with":"string","username_starts_with":"string"}}
//
// swagger:model GetUsersRequestBody
type GetUsersRequestBody struct {

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

// Validate validates this get users request body
func (m *GetUsersRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get users request body based on context it is used
func (m *GetUsersRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetUsersRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUsersRequestBody) UnmarshalBinary(b []byte) error {
	var res GetUsersRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
