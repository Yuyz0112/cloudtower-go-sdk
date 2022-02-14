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

// GetUserLoginInfoesRequestBody get user login infoes request body
// Example: {"after":"userLoginInfoes-id-string","before":"userLoginInfoes-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"UserLoginInfoWhereInput[]","NOT":"UserLoginInfoWhereInput[]","OR":"UserLoginInfoWhereInput[]","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","locked_at":"string","locked_at_gt":"string","locked_at_gte":"string","locked_at_in":["string"],"locked_at_lt":"string","locked_at_lte":"string","locked_at_not":"string","locked_at_not_in":["string"],"miss_num":0,"miss_num_gt":0,"miss_num_gte":0,"miss_num_in":[0],"miss_num_lt":0,"miss_num_lte":0,"miss_num_not":0,"miss_num_not_in":[0],"missed_at":"string","missed_at_gt":"string","missed_at_gte":"string","missed_at_in":["string"],"missed_at_lt":"string","missed_at_lte":"string","missed_at_not":"string","missed_at_not_in":["string"],"user":"UserWhereInput"}}
//
// swagger:model GetUserLoginInfoesRequestBody
type GetUserLoginInfoesRequestBody struct {

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
		UserLoginInfoOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		UserLoginInfoWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get user login infoes request body
func (m *GetUserLoginInfoesRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetUserLoginInfoesRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetUserLoginInfoesRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get user login infoes request body based on the context it is used
func (m *GetUserLoginInfoesRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetUserLoginInfoesRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetUserLoginInfoesRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetUserLoginInfoesRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserLoginInfoesRequestBody) UnmarshalBinary(b []byte) error {
	var res GetUserLoginInfoesRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
