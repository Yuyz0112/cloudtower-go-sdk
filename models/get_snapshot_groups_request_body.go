// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSnapshotGroupsRequestBody get snapshot groups request body
//
// swagger:model GetSnapshotGroupsRequestBody
type GetSnapshotGroupsRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *float64 `json:"first,omitempty"`

	// last
	Last *float64 `json:"last,omitempty"`

	// order by
	OrderBy interface{} `json:"orderBy,omitempty"`

	// skip
	Skip *float64 `json:"skip,omitempty"`

	// where
	Where interface{} `json:"where,omitempty"`
}

// Validate validates this get snapshot groups request body
func (m *GetSnapshotGroupsRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get snapshot groups request body based on context it is used
func (m *GetSnapshotGroupsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetSnapshotGroupsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSnapshotGroupsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetSnapshotGroupsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
