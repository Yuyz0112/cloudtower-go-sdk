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

// GetGraphsRequestBody get graphs request body
// Example: {"after":"graphs-id-string","before":"graphs-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"GraphWhereInput[]","NOT":"GraphWhereInput[]","OR":"GraphWhereInput[]","cluster":"ClusterWhereInput","disks_every":"DiskWhereInput","disks_none":"DiskWhereInput","disks_some":"DiskWhereInput","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"hosts_every":"HostWhereInput","hosts_none":"HostWhereInput","hosts_some":"HostWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","luns_every":"IscsiLunWhereInput","luns_none":"IscsiLunWhereInput","luns_some":"IscsiLunWhereInput","metric_count":0,"metric_count_gt":0,"metric_count_gte":0,"metric_count_in":[0],"metric_count_lt":0,"metric_count_lte":0,"metric_count_not":0,"metric_count_not_in":[0],"metric_name":"string","metric_name_contains":"string","metric_name_ends_with":"string","metric_name_gt":"string","metric_name_gte":"string","metric_name_in":["string"],"metric_name_lt":"string","metric_name_lte":"string","metric_name_not":"string","metric_name_not_contains":"string","metric_name_not_ends_with":"string","metric_name_not_in":["string"],"metric_name_not_starts_with":"string","metric_name_starts_with":"string","metric_type":"BOTTOMK","metric_type_in":["BOTTOMK"],"metric_type_not":"BOTTOMK","metric_type_not_in":["BOTTOMK"],"namespaces_every":"NvmfNamespaceWhereInput","namespaces_none":"NvmfNamespaceWhereInput","namespaces_some":"NvmfNamespaceWhereInput","network":"ACCESS","network_in":["ACCESS"],"network_not":"ACCESS","network_not_in":["ACCESS"],"nics_every":"NicWhereInput","nics_none":"NicWhereInput","nics_some":"NicWhereInput","resource_type":"string","resource_type_contains":"string","resource_type_ends_with":"string","resource_type_gt":"string","resource_type_gte":"string","resource_type_in":["string"],"resource_type_lt":"string","resource_type_lte":"string","resource_type_not":"string","resource_type_not_contains":"string","resource_type_not_ends_with":"string","resource_type_not_in":["string"],"resource_type_not_starts_with":"string","resource_type_starts_with":"string","service":"string","service_contains":"string","service_ends_with":"string","service_gt":"string","service_gte":"string","service_in":["string"],"service_lt":"string","service_lte":"string","service_not":"string","service_not_contains":"string","service_not_ends_with":"string","service_not_in":["string"],"service_not_starts_with":"string","service_starts_with":"string","title":"string","title_contains":"string","title_ends_with":"string","title_gt":"string","title_gte":"string","title_in":["string"],"title_lt":"string","title_lte":"string","title_not":"string","title_not_contains":"string","title_not_ends_with":"string","title_not_in":["string"],"title_not_starts_with":"string","title_starts_with":"string","type":"AREA","type_in":["AREA"],"type_not":"AREA","type_not_in":["AREA"],"view":"ViewWhereInput","vmNics_every":"VmNicWhereInput","vmNics_none":"VmNicWhereInput","vmNics_some":"VmNicWhereInput","vmVolumes_every":"VmVolumeWhereInput","vmVolumes_none":"VmVolumeWhereInput","vmVolumes_some":"VmVolumeWhereInput","vms_every":"VmWhereInput","vms_none":"VmWhereInput","vms_some":"VmWhereInput","witnesses_every":"WitnessWhereInput","witnesses_none":"WitnessWhereInput","witnesses_some":"WitnessWhereInput","zones_every":"ZoneWhereInput","zones_none":"ZoneWhereInput","zones_some":"ZoneWhereInput"}}
//
// swagger:model GetGraphsRequestBody
type GetGraphsRequestBody struct {

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
		GraphOrderByInput
	} `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where struct {
		GraphWhereInput
	} `json:"where,omitempty"`
}

// Validate validates this get graphs request body
func (m *GetGraphsRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetGraphsRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	return nil
}

func (m *GetGraphsRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this get graphs request body based on the context it is used
func (m *GetGraphsRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetGraphsRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *GetGraphsRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *GetGraphsRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetGraphsRequestBody) UnmarshalBinary(b []byte) error {
	var res GetGraphsRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
