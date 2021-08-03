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

// EntityFilterWhereInput entity filter where input
//
// swagger:model EntityFilterWhereInput
type EntityFilterWhereInput struct {

	// a n d
	AND []*EntityFilterWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*EntityFilterWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*EntityFilterWhereInput `json:"OR,omitempty"`

	// clusters every
	ClustersEvery interface{} `json:"clusters_every,omitempty"`

	// clusters none
	ClustersNone interface{} `json:"clusters_none,omitempty"`

	// clusters some
	ClustersSome interface{} `json:"clusters_some,omitempty"`

	// entity type
	EntityType interface{} `json:"entity_type,omitempty"`

	// entity type in
	EntityTypeIn []EntityType `json:"entity_type_in,omitempty"`

	// entity type not
	EntityTypeNot interface{} `json:"entity_type_not,omitempty"`

	// entity type not in
	EntityTypeNotIn []EntityType `json:"entity_type_not_in,omitempty"`

	// exec failed cluster every
	ExecFailedClusterEvery interface{} `json:"exec_failed_cluster_every,omitempty"`

	// exec failed cluster none
	ExecFailedClusterNone interface{} `json:"exec_failed_cluster_none,omitempty"`

	// exec failed cluster some
	ExecFailedClusterSome interface{} `json:"exec_failed_cluster_some,omitempty"`

	// filter status
	FilterStatus interface{} `json:"filter_status,omitempty"`

	// filter status in
	FilterStatusIn []FilterStatus `json:"filter_status_in,omitempty"`

	// filter status not
	FilterStatusNot interface{} `json:"filter_status_not,omitempty"`

	// filter status not in
	FilterStatusNotIn []FilterStatus `json:"filter_status_not_in,omitempty"`

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

	// last executed at
	LastExecutedAt *string `json:"last_executed_at,omitempty"`

	// last executed at gt
	LastExecutedAtGt *string `json:"last_executed_at_gt,omitempty"`

	// last executed at gte
	LastExecutedAtGte *string `json:"last_executed_at_gte,omitempty"`

	// last executed at in
	LastExecutedAtIn []string `json:"last_executed_at_in,omitempty"`

	// last executed at lt
	LastExecutedAtLt *string `json:"last_executed_at_lt,omitempty"`

	// last executed at lte
	LastExecutedAtLte *string `json:"last_executed_at_lte,omitempty"`

	// last executed at not
	LastExecutedAtNot *string `json:"last_executed_at_not,omitempty"`

	// last executed at not in
	LastExecutedAtNotIn []string `json:"last_executed_at_not_in,omitempty"`

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

	// preset
	Preset *string `json:"preset,omitempty"`

	// preset contains
	PresetContains *string `json:"preset_contains,omitempty"`

	// preset ends with
	PresetEndsWith *string `json:"preset_ends_with,omitempty"`

	// preset gt
	PresetGt *string `json:"preset_gt,omitempty"`

	// preset gte
	PresetGte *string `json:"preset_gte,omitempty"`

	// preset in
	PresetIn []string `json:"preset_in,omitempty"`

	// preset lt
	PresetLt *string `json:"preset_lt,omitempty"`

	// preset lte
	PresetLte *string `json:"preset_lte,omitempty"`

	// preset not
	PresetNot *string `json:"preset_not,omitempty"`

	// preset not contains
	PresetNotContains *string `json:"preset_not_contains,omitempty"`

	// preset not ends with
	PresetNotEndsWith *string `json:"preset_not_ends_with,omitempty"`

	// preset not in
	PresetNotIn []string `json:"preset_not_in,omitempty"`

	// preset not starts with
	PresetNotStartsWith *string `json:"preset_not_starts_with,omitempty"`

	// preset starts with
	PresetStartsWith *string `json:"preset_starts_with,omitempty"`
}

// Validate validates this entity filter where input
func (m *EntityFilterWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEntityTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilterStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityFilterWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *EntityFilterWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *EntityFilterWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *EntityFilterWhereInput) validateEntityTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityTypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityTypeIn); i++ {

		if err := m.EntityTypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EntityFilterWhereInput) validateEntityTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityTypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityTypeNotIn); i++ {

		if err := m.EntityTypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EntityFilterWhereInput) validateFilterStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.FilterStatusIn); i++ {

		if err := m.FilterStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EntityFilterWhereInput) validateFilterStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.FilterStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.FilterStatusNotIn); i++ {

		if err := m.FilterStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this entity filter where input based on the context it is used
func (m *EntityFilterWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEntityTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilterStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilterStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityFilterWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EntityFilterWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EntityFilterWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EntityFilterWhereInput) contextValidateEntityTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityTypeIn); i++ {

		if err := m.EntityTypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity_type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EntityFilterWhereInput) contextValidateEntityTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityTypeNotIn); i++ {

		if err := m.EntityTypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity_type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EntityFilterWhereInput) contextValidateFilterStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FilterStatusIn); i++ {

		if err := m.FilterStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *EntityFilterWhereInput) contextValidateFilterStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FilterStatusNotIn); i++ {

		if err := m.FilterStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter_status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityFilterWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityFilterWhereInput) UnmarshalBinary(b []byte) error {
	var res EntityFilterWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
