// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// VlanOrderByInput vlan order by input
//
// swagger:model VlanOrderByInput
type VlanOrderByInput string

func NewVlanOrderByInput(value VlanOrderByInput) *VlanOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VlanOrderByInput.
func (m VlanOrderByInput) Pointer() *VlanOrderByInput {
	return &m
}

const (

	// VlanOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	VlanOrderByInputCreatedAtASC VlanOrderByInput = "createdAt_ASC"

	// VlanOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	VlanOrderByInputCreatedAtDESC VlanOrderByInput = "createdAt_DESC"

	// VlanOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	VlanOrderByInputEntityAsyncStatusASC VlanOrderByInput = "entityAsyncStatus_ASC"

	// VlanOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	VlanOrderByInputEntityAsyncStatusDESC VlanOrderByInput = "entityAsyncStatus_DESC"

	// VlanOrderByInputGatewayIPASC captures enum value "gateway_ip_ASC"
	VlanOrderByInputGatewayIPASC VlanOrderByInput = "gateway_ip_ASC"

	// VlanOrderByInputGatewayIPDESC captures enum value "gateway_ip_DESC"
	VlanOrderByInputGatewayIPDESC VlanOrderByInput = "gateway_ip_DESC"

	// VlanOrderByInputGatewaySubnetmaskASC captures enum value "gateway_subnetmask_ASC"
	VlanOrderByInputGatewaySubnetmaskASC VlanOrderByInput = "gateway_subnetmask_ASC"

	// VlanOrderByInputGatewaySubnetmaskDESC captures enum value "gateway_subnetmask_DESC"
	VlanOrderByInputGatewaySubnetmaskDESC VlanOrderByInput = "gateway_subnetmask_DESC"

	// VlanOrderByInputIDASC captures enum value "id_ASC"
	VlanOrderByInputIDASC VlanOrderByInput = "id_ASC"

	// VlanOrderByInputIDDESC captures enum value "id_DESC"
	VlanOrderByInputIDDESC VlanOrderByInput = "id_DESC"

	// VlanOrderByInputLocalIDASC captures enum value "local_id_ASC"
	VlanOrderByInputLocalIDASC VlanOrderByInput = "local_id_ASC"

	// VlanOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	VlanOrderByInputLocalIDDESC VlanOrderByInput = "local_id_DESC"

	// VlanOrderByInputNameASC captures enum value "name_ASC"
	VlanOrderByInputNameASC VlanOrderByInput = "name_ASC"

	// VlanOrderByInputNameDESC captures enum value "name_DESC"
	VlanOrderByInputNameDESC VlanOrderByInput = "name_DESC"

	// VlanOrderByInputSubnetmaskASC captures enum value "subnetmask_ASC"
	VlanOrderByInputSubnetmaskASC VlanOrderByInput = "subnetmask_ASC"

	// VlanOrderByInputSubnetmaskDESC captures enum value "subnetmask_DESC"
	VlanOrderByInputSubnetmaskDESC VlanOrderByInput = "subnetmask_DESC"

	// VlanOrderByInputTypeASC captures enum value "type_ASC"
	VlanOrderByInputTypeASC VlanOrderByInput = "type_ASC"

	// VlanOrderByInputTypeDESC captures enum value "type_DESC"
	VlanOrderByInputTypeDESC VlanOrderByInput = "type_DESC"

	// VlanOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	VlanOrderByInputUpdatedAtASC VlanOrderByInput = "updatedAt_ASC"

	// VlanOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	VlanOrderByInputUpdatedAtDESC VlanOrderByInput = "updatedAt_DESC"

	// VlanOrderByInputVlanIDASC captures enum value "vlan_id_ASC"
	VlanOrderByInputVlanIDASC VlanOrderByInput = "vlan_id_ASC"

	// VlanOrderByInputVlanIDDESC captures enum value "vlan_id_DESC"
	VlanOrderByInputVlanIDDESC VlanOrderByInput = "vlan_id_DESC"
)

// for schema
var vlanOrderByInputEnum []interface{}

func init() {
	var res []VlanOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","gateway_ip_ASC","gateway_ip_DESC","gateway_subnetmask_ASC","gateway_subnetmask_DESC","id_ASC","id_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","subnetmask_ASC","subnetmask_DESC","type_ASC","type_DESC","updatedAt_ASC","updatedAt_DESC","vlan_id_ASC","vlan_id_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vlanOrderByInputEnum = append(vlanOrderByInputEnum, v)
	}
}

func (m VlanOrderByInput) validateVlanOrderByInputEnum(path, location string, value VlanOrderByInput) error {
	if err := validate.EnumCase(path, location, value, vlanOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this vlan order by input
func (m VlanOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVlanOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this vlan order by input based on context it is used
func (m VlanOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
