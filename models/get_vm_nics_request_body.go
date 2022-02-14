// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetVMNicsRequestBody get Vm nics request body
// Example: {"after":"vmNics-id-string","before":"vmNics-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"VmNicWhereInput[]","NOT":"VmNicWhereInput[]","OR":"VmNicWhereInput[]","enabled":false,"enabled_not":false,"gateway":"string","gateway_contains":"string","gateway_ends_with":"string","gateway_gt":"string","gateway_gte":"string","gateway_in":["string"],"gateway_lt":"string","gateway_lte":"string","gateway_not":"string","gateway_not_contains":"string","gateway_not_ends_with":"string","gateway_not_in":["string"],"gateway_not_starts_with":"string","gateway_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","interface_id":"string","interface_id_contains":"string","interface_id_ends_with":"string","interface_id_gt":"string","interface_id_gte":"string","interface_id_in":["string"],"interface_id_lt":"string","interface_id_lte":"string","interface_id_not":"string","interface_id_not_contains":"string","interface_id_not_ends_with":"string","interface_id_not_in":["string"],"interface_id_not_starts_with":"string","interface_id_starts_with":"string","ip_address":"string","ip_address_contains":"string","ip_address_ends_with":"string","ip_address_gt":"string","ip_address_gte":"string","ip_address_in":["string"],"ip_address_lt":"string","ip_address_lte":"string","ip_address_not":"string","ip_address_not_contains":"string","ip_address_not_ends_with":"string","ip_address_not_in":["string"],"ip_address_not_starts_with":"string","ip_address_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","mac_address":"string","mac_address_contains":"string","mac_address_ends_with":"string","mac_address_gt":"string","mac_address_gte":"string","mac_address_in":["string"],"mac_address_lt":"string","mac_address_lte":"string","mac_address_not":"string","mac_address_not_contains":"string","mac_address_not_ends_with":"string","mac_address_not_in":["string"],"mac_address_not_starts_with":"string","mac_address_starts_with":"string","mirror":false,"mirror_not":false,"model":"E1000","model_in":["E1000"],"model_not":"E1000","model_not_in":["E1000"],"nic":"NicWhereInput","order":0,"order_gt":0,"order_gte":0,"order_in":[0],"order_lt":0,"order_lte":0,"order_not":0,"order_not_in":[0],"subnet_mask":"string","subnet_mask_contains":"string","subnet_mask_ends_with":"string","subnet_mask_gt":"string","subnet_mask_gte":"string","subnet_mask_in":["string"],"subnet_mask_lt":"string","subnet_mask_lte":"string","subnet_mask_not":"string","subnet_mask_not_contains":"string","subnet_mask_not_ends_with":"string","subnet_mask_not_in":["string"],"subnet_mask_not_starts_with":"string","subnet_mask_starts_with":"string","vlan":"VlanWhereInput","vm":"VmWhereInput"}}
//
// swagger:model GetVmNicsRequestBody
type GetVMNicsRequestBody struct {

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

// Validate validates this get Vm nics request body
func (m *GetVMNicsRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get Vm nics request body based on context it is used
func (m *GetVMNicsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetVMNicsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetVMNicsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetVMNicsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
