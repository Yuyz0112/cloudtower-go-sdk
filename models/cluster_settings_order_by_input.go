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

// ClusterSettingsOrderByInput cluster settings order by input
//
// swagger:model ClusterSettingsOrderByInput
type ClusterSettingsOrderByInput string

func NewClusterSettingsOrderByInput(value ClusterSettingsOrderByInput) *ClusterSettingsOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ClusterSettingsOrderByInput.
func (m ClusterSettingsOrderByInput) Pointer() *ClusterSettingsOrderByInput {
	return &m
}

const (

	// ClusterSettingsOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ClusterSettingsOrderByInputCreatedAtASC ClusterSettingsOrderByInput = "createdAt_ASC"

	// ClusterSettingsOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ClusterSettingsOrderByInputCreatedAtDESC ClusterSettingsOrderByInput = "createdAt_DESC"

	// ClusterSettingsOrderByInputDefaultHaASC captures enum value "default_ha_ASC"
	ClusterSettingsOrderByInputDefaultHaASC ClusterSettingsOrderByInput = "default_ha_ASC"

	// ClusterSettingsOrderByInputDefaultHaDESC captures enum value "default_ha_DESC"
	ClusterSettingsOrderByInputDefaultHaDESC ClusterSettingsOrderByInput = "default_ha_DESC"

	// ClusterSettingsOrderByInputIDASC captures enum value "id_ASC"
	ClusterSettingsOrderByInputIDASC ClusterSettingsOrderByInput = "id_ASC"

	// ClusterSettingsOrderByInputIDDESC captures enum value "id_DESC"
	ClusterSettingsOrderByInputIDDESC ClusterSettingsOrderByInput = "id_DESC"

	// ClusterSettingsOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	ClusterSettingsOrderByInputUpdatedAtASC ClusterSettingsOrderByInput = "updatedAt_ASC"

	// ClusterSettingsOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	ClusterSettingsOrderByInputUpdatedAtDESC ClusterSettingsOrderByInput = "updatedAt_DESC"

	// ClusterSettingsOrderByInputVMRecycleBinASC captures enum value "vm_recycle_bin_ASC"
	ClusterSettingsOrderByInputVMRecycleBinASC ClusterSettingsOrderByInput = "vm_recycle_bin_ASC"

	// ClusterSettingsOrderByInputVMRecycleBinDESC captures enum value "vm_recycle_bin_DESC"
	ClusterSettingsOrderByInputVMRecycleBinDESC ClusterSettingsOrderByInput = "vm_recycle_bin_DESC"
)

// for schema
var clusterSettingsOrderByInputEnum []interface{}

func init() {
	var res []ClusterSettingsOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","default_ha_ASC","default_ha_DESC","id_ASC","id_DESC","updatedAt_ASC","updatedAt_DESC","vm_recycle_bin_ASC","vm_recycle_bin_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterSettingsOrderByInputEnum = append(clusterSettingsOrderByInputEnum, v)
	}
}

func (m ClusterSettingsOrderByInput) validateClusterSettingsOrderByInputEnum(path, location string, value ClusterSettingsOrderByInput) error {
	if err := validate.EnumCase(path, location, value, clusterSettingsOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cluster settings order by input
func (m ClusterSettingsOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateClusterSettingsOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cluster settings order by input based on context it is used
func (m ClusterSettingsOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
