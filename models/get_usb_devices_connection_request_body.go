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

// GetUsbDevicesConnectionRequestBody get usb devices connection request body
// Example: {"after":"usbDevicesConnection-id-string","before":"usbDevicesConnection-id-string","first":0,"last":0,"orderBy":"binded_ASC","skip":0,"where":{"AND":"UsbDeviceWhereInput[]","NOT":"UsbDeviceWhereInput[]","OR":"UsbDeviceWhereInput[]","binded":false,"binded_not":false,"description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","host":"HostWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","manufacturer":"string","manufacturer_contains":"string","manufacturer_ends_with":"string","manufacturer_gt":"string","manufacturer_gte":"string","manufacturer_in":["string"],"manufacturer_lt":"string","manufacturer_lte":"string","manufacturer_not":"string","manufacturer_not_contains":"string","manufacturer_not_ends_with":"string","manufacturer_not_in":["string"],"manufacturer_not_starts_with":"string","manufacturer_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","size":0,"size_gt":0,"size_gte":0,"size_in":[0],"size_lt":0,"size_lte":0,"size_not":0,"size_not_in":[0],"status":"EJECTED","status_in":["EJECTED"],"status_not":"EJECTED","status_not_in":["EJECTED"],"usb_type":"string","usb_type_contains":"string","usb_type_ends_with":"string","usb_type_gt":"string","usb_type_gte":"string","usb_type_in":["string"],"usb_type_lt":"string","usb_type_lte":"string","usb_type_not":"string","usb_type_not_contains":"string","usb_type_not_ends_with":"string","usb_type_not_in":["string"],"usb_type_not_starts_with":"string","usb_type_starts_with":"string","vm":"VmWhereInput"}}
//
// swagger:model GetUsbDevicesConnectionRequestBody
type GetUsbDevicesConnectionRequestBody struct {

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
		UsbDeviceOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		UsbDeviceWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get usb devices connection request body
func (m *GetUsbDevicesConnectionRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetUsbDevicesConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetUsbDevicesConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get usb devices connection request body based on the context it is used
func (m *GetUsbDevicesConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetUsbDevicesConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetUsbDevicesConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetUsbDevicesConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUsbDevicesConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetUsbDevicesConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
