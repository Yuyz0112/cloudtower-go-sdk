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

// BackupPlanStatus backup plan status
//
// swagger:model BackupPlanStatus
type BackupPlanStatus string

func NewBackupPlanStatus(value BackupPlanStatus) *BackupPlanStatus {
	v := value
	return &v
}

const (

	// BackupPlanStatusPAUSED captures enum value "PAUSED"
	BackupPlanStatusPAUSED BackupPlanStatus = "PAUSED"

	// BackupPlanStatusSTOPPED captures enum value "STOPPED"
	BackupPlanStatusSTOPPED BackupPlanStatus = "STOPPED"

	// BackupPlanStatusWORKING captures enum value "WORKING"
	BackupPlanStatusWORKING BackupPlanStatus = "WORKING"
)

// for schema
var backupPlanStatusEnum []interface{}

func init() {
	var res []BackupPlanStatus
	if err := json.Unmarshal([]byte(`["PAUSED","STOPPED","WORKING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupPlanStatusEnum = append(backupPlanStatusEnum, v)
	}
}

func (m BackupPlanStatus) validateBackupPlanStatusEnum(path, location string, value BackupPlanStatus) error {
	if err := validate.EnumCase(path, location, value, backupPlanStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this backup plan status
func (m BackupPlanStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBackupPlanStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this backup plan status based on context it is used
func (m BackupPlanStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
