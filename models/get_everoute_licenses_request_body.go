// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetEverouteLicensesRequestBody get everoute licenses request body
// Example: {"after":"everouteLicenses-id-string","before":"everouteLicenses-id-string","first":0,"last":0,"orderBy":"code_ASC","skip":0,"where":{"AND":"EverouteLicenseWhereInput[]","NOT":"EverouteLicenseWhereInput[]","OR":"EverouteLicenseWhereInput[]","code":"string","code_contains":"string","code_ends_with":"string","code_gt":"string","code_gte":"string","code_in":["string"],"code_lt":"string","code_lte":"string","code_not":"string","code_not_contains":"string","code_not_ends_with":"string","code_not_in":["string"],"code_not_starts_with":"string","code_starts_with":"string","expire_date":"string","expire_date_gt":"string","expire_date_gte":"string","expire_date_in":["string"],"expire_date_lt":"string","expire_date_lte":"string","expire_date_not":"string","expire_date_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","max_socket_num":0,"max_socket_num_gt":0,"max_socket_num_gte":0,"max_socket_num_in":[0],"max_socket_num_lt":0,"max_socket_num_lte":0,"max_socket_num_not":0,"max_socket_num_not_in":[0],"serial":"string","serial_contains":"string","serial_ends_with":"string","serial_gt":"string","serial_gte":"string","serial_in":["string"],"serial_lt":"string","serial_lte":"string","serial_not":"string","serial_not_contains":"string","serial_not_ends_with":"string","serial_not_in":["string"],"serial_not_starts_with":"string","serial_starts_with":"string","sign_date":"string","sign_date_gt":"string","sign_date_gte":"string","sign_date_in":["string"],"sign_date_lt":"string","sign_date_lte":"string","sign_date_not":"string","sign_date_not_in":["string"],"software_edition":"COMMUNITY","software_edition_in":["COMMUNITY"],"software_edition_not":"COMMUNITY","software_edition_not_in":["COMMUNITY"],"type":"PERPETUAL","type_in":["PERPETUAL"],"type_not":"PERPETUAL","type_not_in":["PERPETUAL"],"uid":"string","uid_contains":"string","uid_ends_with":"string","uid_gt":"string","uid_gte":"string","uid_in":["string"],"uid_lt":"string","uid_lte":"string","uid_not":"string","uid_not_contains":"string","uid_not_ends_with":"string","uid_not_in":["string"],"uid_not_starts_with":"string","uid_starts_with":"string"}}
//
// swagger:model GetEverouteLicensesRequestBody
type GetEverouteLicensesRequestBody struct {

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

// Validate validates this get everoute licenses request body
func (m *GetEverouteLicensesRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get everoute licenses request body based on context it is used
func (m *GetEverouteLicensesRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetEverouteLicensesRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEverouteLicensesRequestBody) UnmarshalBinary(b []byte) error {
	var res GetEverouteLicensesRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
