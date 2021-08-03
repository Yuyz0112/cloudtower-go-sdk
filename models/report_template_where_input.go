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

// ReportTemplateWhereInput report template where input
//
// swagger:model ReportTemplateWhereInput
type ReportTemplateWhereInput struct {

	// a n d
	AND []*ReportTemplateWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*ReportTemplateWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*ReportTemplateWhereInput `json:"OR,omitempty"`

	// created at
	CreatedAt *string `json:"createdAt,omitempty"`

	// created at gt
	CreatedAtGt *string `json:"createdAt_gt,omitempty"`

	// created at gte
	CreatedAtGte *string `json:"createdAt_gte,omitempty"`

	// created at in
	CreatedAtIn []string `json:"createdAt_in,omitempty"`

	// created at lt
	CreatedAtLt *string `json:"createdAt_lt,omitempty"`

	// created at lte
	CreatedAtLte *string `json:"createdAt_lte,omitempty"`

	// created at not
	CreatedAtNot *string `json:"createdAt_not,omitempty"`

	// created at not in
	CreatedAtNotIn []string `json:"createdAt_not_in,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// description contains
	DescriptionContains *string `json:"description_contains,omitempty"`

	// description ends with
	DescriptionEndsWith *string `json:"description_ends_with,omitempty"`

	// description gt
	DescriptionGt *string `json:"description_gt,omitempty"`

	// description gte
	DescriptionGte *string `json:"description_gte,omitempty"`

	// description in
	DescriptionIn []string `json:"description_in,omitempty"`

	// description lt
	DescriptionLt *string `json:"description_lt,omitempty"`

	// description lte
	DescriptionLte *string `json:"description_lte,omitempty"`

	// description not
	DescriptionNot *string `json:"description_not,omitempty"`

	// description not contains
	DescriptionNotContains *string `json:"description_not_contains,omitempty"`

	// description not ends with
	DescriptionNotEndsWith *string `json:"description_not_ends_with,omitempty"`

	// description not in
	DescriptionNotIn []string `json:"description_not_in,omitempty"`

	// description not starts with
	DescriptionNotStartsWith *string `json:"description_not_starts_with,omitempty"`

	// description starts with
	DescriptionStartsWith *string `json:"description_starts_with,omitempty"`

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

	// task num
	TaskNum *float64 `json:"task_num,omitempty"`

	// task num gt
	TaskNumGt *float64 `json:"task_num_gt,omitempty"`

	// task num gte
	TaskNumGte *float64 `json:"task_num_gte,omitempty"`

	// task num in
	TaskNumIn []float64 `json:"task_num_in,omitempty"`

	// task num lt
	TaskNumLt *float64 `json:"task_num_lt,omitempty"`

	// task num lte
	TaskNumLte *float64 `json:"task_num_lte,omitempty"`

	// task num not
	TaskNumNot *float64 `json:"task_num_not,omitempty"`

	// task num not in
	TaskNumNotIn []float64 `json:"task_num_not_in,omitempty"`

	// tasks every
	TasksEvery interface{} `json:"tasks_every,omitempty"`

	// tasks none
	TasksNone interface{} `json:"tasks_none,omitempty"`

	// tasks some
	TasksSome interface{} `json:"tasks_some,omitempty"`
}

// Validate validates this report template where input
func (m *ReportTemplateWhereInput) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportTemplateWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *ReportTemplateWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *ReportTemplateWhereInput) validateOR(formats strfmt.Registry) error {
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

// ContextValidate validate this report template where input based on the context it is used
func (m *ReportTemplateWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportTemplateWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ReportTemplateWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ReportTemplateWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ReportTemplateWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportTemplateWhereInput) UnmarshalBinary(b []byte) error {
	var res ReportTemplateWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}