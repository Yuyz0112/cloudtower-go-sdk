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

// GetReportTemplatesRequestBody get report templates request body
// Example: {"after":"reportTemplates-id-string","before":"reportTemplates-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"ReportTemplateWhereInput[]","NOT":"ReportTemplateWhereInput[]","OR":"ReportTemplateWhereInput[]","createdAt":"string","createdAt_gt":"string","createdAt_gte":"string","createdAt_in":["string"],"createdAt_lt":"string","createdAt_lte":"string","createdAt_not":"string","createdAt_not_in":["string"],"description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","preset":"string","preset_contains":"string","preset_ends_with":"string","preset_gt":"string","preset_gte":"string","preset_in":["string"],"preset_lt":"string","preset_lte":"string","preset_not":"string","preset_not_contains":"string","preset_not_ends_with":"string","preset_not_in":["string"],"preset_not_starts_with":"string","preset_starts_with":"string","task_num":0,"task_num_gt":0,"task_num_gte":0,"task_num_in":[0],"task_num_lt":0,"task_num_lte":0,"task_num_not":0,"task_num_not_in":[0],"tasks_every":"ReportTaskWhereInput","tasks_none":"ReportTaskWhereInput","tasks_some":"ReportTaskWhereInput"}}
//
// swagger:model GetReportTemplatesRequestBody
type GetReportTemplatesRequestBody struct {

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
		ReportTemplateOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		ReportTemplateWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get report templates request body
func (m *GetReportTemplatesRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetReportTemplatesRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetReportTemplatesRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get report templates request body based on the context it is used
func (m *GetReportTemplatesRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetReportTemplatesRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetReportTemplatesRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetReportTemplatesRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetReportTemplatesRequestBody) UnmarshalBinary(b []byte) error {
	var res GetReportTemplatesRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
