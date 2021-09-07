// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VlanWhereInput vlan where input
//
// swagger:model VlanWhereInput
type VlanWhereInput struct {

	// a n d
	AND []*VlanWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VlanWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VlanWhereInput `json:"OR,omitempty"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot interface{} `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// gateway ip
	GatewayIP *string `json:"gateway_ip,omitempty"`

	// gateway ip contains
	GatewayIPContains *string `json:"gateway_ip_contains,omitempty"`

	// gateway ip ends with
	GatewayIPEndsWith *string `json:"gateway_ip_ends_with,omitempty"`

	// gateway ip gt
	GatewayIPGt *string `json:"gateway_ip_gt,omitempty"`

	// gateway ip gte
	GatewayIPGte *string `json:"gateway_ip_gte,omitempty"`

	// gateway ip in
	GatewayIPIn []string `json:"gateway_ip_in,omitempty"`

	// gateway ip lt
	GatewayIPLt *string `json:"gateway_ip_lt,omitempty"`

	// gateway ip lte
	GatewayIPLte *string `json:"gateway_ip_lte,omitempty"`

	// gateway ip not
	GatewayIPNot *string `json:"gateway_ip_not,omitempty"`

	// gateway ip not contains
	GatewayIPNotContains *string `json:"gateway_ip_not_contains,omitempty"`

	// gateway ip not ends with
	GatewayIPNotEndsWith *string `json:"gateway_ip_not_ends_with,omitempty"`

	// gateway ip not in
	GatewayIPNotIn []string `json:"gateway_ip_not_in,omitempty"`

	// gateway ip not starts with
	GatewayIPNotStartsWith *string `json:"gateway_ip_not_starts_with,omitempty"`

	// gateway ip starts with
	GatewayIPStartsWith *string `json:"gateway_ip_starts_with,omitempty"`

	// gateway subnetmask
	GatewaySubnetmask *string `json:"gateway_subnetmask,omitempty"`

	// gateway subnetmask contains
	GatewaySubnetmaskContains *string `json:"gateway_subnetmask_contains,omitempty"`

	// gateway subnetmask ends with
	GatewaySubnetmaskEndsWith *string `json:"gateway_subnetmask_ends_with,omitempty"`

	// gateway subnetmask gt
	GatewaySubnetmaskGt *string `json:"gateway_subnetmask_gt,omitempty"`

	// gateway subnetmask gte
	GatewaySubnetmaskGte *string `json:"gateway_subnetmask_gte,omitempty"`

	// gateway subnetmask in
	GatewaySubnetmaskIn []string `json:"gateway_subnetmask_in,omitempty"`

	// gateway subnetmask lt
	GatewaySubnetmaskLt *string `json:"gateway_subnetmask_lt,omitempty"`

	// gateway subnetmask lte
	GatewaySubnetmaskLte *string `json:"gateway_subnetmask_lte,omitempty"`

	// gateway subnetmask not
	GatewaySubnetmaskNot *string `json:"gateway_subnetmask_not,omitempty"`

	// gateway subnetmask not contains
	GatewaySubnetmaskNotContains *string `json:"gateway_subnetmask_not_contains,omitempty"`

	// gateway subnetmask not ends with
	GatewaySubnetmaskNotEndsWith *string `json:"gateway_subnetmask_not_ends_with,omitempty"`

	// gateway subnetmask not in
	GatewaySubnetmaskNotIn []string `json:"gateway_subnetmask_not_in,omitempty"`

	// gateway subnetmask not starts with
	GatewaySubnetmaskNotStartsWith *string `json:"gateway_subnetmask_not_starts_with,omitempty"`

	// gateway subnetmask starts with
	GatewaySubnetmaskStartsWith *string `json:"gateway_subnetmask_starts_with,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// labels every
	LabelsEvery interface{} `json:"labels_every,omitempty"`

	// labels none
	LabelsNone interface{} `json:"labels_none,omitempty"`

	// labels some
	LabelsSome interface{} `json:"labels_some,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// subnetmask
	Subnetmask *string `json:"subnetmask,omitempty"`

	// subnetmask contains
	SubnetmaskContains *string `json:"subnetmask_contains,omitempty"`

	// subnetmask ends with
	SubnetmaskEndsWith *string `json:"subnetmask_ends_with,omitempty"`

	// subnetmask gt
	SubnetmaskGt *string `json:"subnetmask_gt,omitempty"`

	// subnetmask gte
	SubnetmaskGte *string `json:"subnetmask_gte,omitempty"`

	// subnetmask in
	SubnetmaskIn []string `json:"subnetmask_in,omitempty"`

	// subnetmask lt
	SubnetmaskLt *string `json:"subnetmask_lt,omitempty"`

	// subnetmask lte
	SubnetmaskLte *string `json:"subnetmask_lte,omitempty"`

	// subnetmask not
	SubnetmaskNot *string `json:"subnetmask_not,omitempty"`

	// subnetmask not contains
	SubnetmaskNotContains *string `json:"subnetmask_not_contains,omitempty"`

	// subnetmask not ends with
	SubnetmaskNotEndsWith *string `json:"subnetmask_not_ends_with,omitempty"`

	// subnetmask not in
	SubnetmaskNotIn []string `json:"subnetmask_not_in,omitempty"`

	// subnetmask not starts with
	SubnetmaskNotStartsWith *string `json:"subnetmask_not_starts_with,omitempty"`

	// subnetmask starts with
	SubnetmaskStartsWith *string `json:"subnetmask_starts_with,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`

	// type in
	TypeIn []NetworkType `json:"type_in,omitempty"`

	// type not
	TypeNot interface{} `json:"type_not,omitempty"`

	// type not in
	TypeNotIn []NetworkType `json:"type_not_in,omitempty"`

	// vds
	Vds interface{} `json:"vds,omitempty"`

	// vlan id
	VlanID *float64 `json:"vlan_id,omitempty"`

	// vlan id gt
	VlanIDGt *float64 `json:"vlan_id_gt,omitempty"`

	// vlan id gte
	VlanIDGte *float64 `json:"vlan_id_gte,omitempty"`

	// vlan id in
	VlanIDIn []float64 `json:"vlan_id_in,omitempty"`

	// vlan id lt
	VlanIDLt *float64 `json:"vlan_id_lt,omitempty"`

	// vlan id lte
	VlanIDLte *float64 `json:"vlan_id_lte,omitempty"`

	// vlan id not
	VlanIDNot *float64 `json:"vlan_id_not,omitempty"`

	// vlan id not in
	VlanIDNotIn []float64 `json:"vlan_id_not_in,omitempty"`

	// vm nics every
	VMNicsEvery interface{} `json:"vm_nics_every,omitempty"`

	// vm nics none
	VMNicsNone interface{} `json:"vm_nics_none,omitempty"`

	// vm nics some
	VMNicsSome interface{} `json:"vm_nics_some,omitempty"`
}

// Validate validates this vlan where input
func (m *VlanWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VlanWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VlanWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VlanWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VlanWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VlanWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VlanWhereInput) validateTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VlanWhereInput) validateTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this vlan where input based on the context it is used
func (m *VlanWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VlanWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VlanWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VlanWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VlanWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VlanWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VlanWhereInput) contextValidateTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *VlanWhereInput) contextValidateTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VlanWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VlanWhereInput) UnmarshalBinary(b []byte) error {
	var res VlanWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
