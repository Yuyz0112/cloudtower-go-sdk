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

// GetBackupLicensesRequestBody get backup licenses request body
// Example: {"after":"backupLicenses-id-string","before":"backupLicenses-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"BackupLicenseWhereInput[]","NOT":"BackupLicenseWhereInput[]","OR":"BackupLicenseWhereInput[]","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"expire_date":"string","expire_date_gt":"string","expire_date_gte":"string","expire_date_in":["string"],"expire_date_lt":"string","expire_date_lte":"string","expire_date_not":"string","expire_date_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","license_serial":"string","license_serial_contains":"string","license_serial_ends_with":"string","license_serial_gt":"string","license_serial_gte":"string","license_serial_in":["string"],"license_serial_lt":"string","license_serial_lte":"string","license_serial_not":"string","license_serial_not_contains":"string","license_serial_not_ends_with":"string","license_serial_not_in":["string"],"license_serial_not_starts_with":"string","license_serial_starts_with":"string","max_capacity":0,"max_capacity_gt":0,"max_capacity_gte":0,"max_capacity_in":[0],"max_capacity_lt":0,"max_capacity_lte":0,"max_capacity_not":0,"max_capacity_not_in":[0],"sign_date":"string","sign_date_gt":"string","sign_date_gte":"string","sign_date_in":["string"],"sign_date_lt":"string","sign_date_lte":"string","sign_date_not":"string","sign_date_not_in":["string"],"software_edition":"COMMUNITY","software_edition_in":["COMMUNITY"],"software_edition_not":"COMMUNITY","software_edition_not_in":["COMMUNITY"],"type":"PERPETUAL","type_in":["PERPETUAL"],"type_not":"PERPETUAL","type_not_in":["PERPETUAL"]}}
//
// swagger:model GetBackupLicensesRequestBody
type GetBackupLicensesRequestBody struct {

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
		BackupLicenseOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		BackupLicenseWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get backup licenses request body
func (m *GetBackupLicensesRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetBackupLicensesRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetBackupLicensesRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get backup licenses request body based on the context it is used
func (m *GetBackupLicensesRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetBackupLicensesRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetBackupLicensesRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetBackupLicensesRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupLicensesRequestBody) UnmarshalBinary(b []byte) error {
	var res GetBackupLicensesRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
