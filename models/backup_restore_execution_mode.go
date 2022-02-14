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

// BackupRestoreExecutionMode backup restore execution mode
//
// swagger:model BackupRestoreExecutionMode
type BackupRestoreExecutionMode string

func NewBackupRestoreExecutionMode(value BackupRestoreExecutionMode) *BackupRestoreExecutionMode {
	v := value
	return &v
}

const (

	// BackupRestoreExecutionModeINPLACE captures enum value "INPLACE"
	BackupRestoreExecutionModeINPLACE BackupRestoreExecutionMode = "INPLACE"

	// BackupRestoreExecutionModeREBUILD captures enum value "REBUILD"
	BackupRestoreExecutionModeREBUILD BackupRestoreExecutionMode = "REBUILD"
)

// for schema
var backupRestoreExecutionModeEnum []interface{}

func init() {
	var res []BackupRestoreExecutionMode
	if err := json.Unmarshal([]byte(`["INPLACE","REBUILD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupRestoreExecutionModeEnum = append(backupRestoreExecutionModeEnum, v)
	}
}

func (m BackupRestoreExecutionMode) validateBackupRestoreExecutionModeEnum(path, location string, value BackupRestoreExecutionMode) error {
	if err := validate.EnumCase(path, location, value, backupRestoreExecutionModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this backup restore execution mode
func (m BackupRestoreExecutionMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBackupRestoreExecutionModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this backup restore execution mode based on context it is used
func (m BackupRestoreExecutionMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
