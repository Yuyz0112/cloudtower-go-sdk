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

// BackupPlanExecutionStatus backup plan execution status
//
// swagger:model BackupPlanExecutionStatus
type BackupPlanExecutionStatus string

func NewBackupPlanExecutionStatus(value BackupPlanExecutionStatus) *BackupPlanExecutionStatus {
	v := value
	return &v
}

const (

	// BackupPlanExecutionStatusFAILED captures enum value "FAILED"
	BackupPlanExecutionStatusFAILED BackupPlanExecutionStatus = "FAILED"

	// BackupPlanExecutionStatusNEVEREXECUTE captures enum value "NEVER_EXECUTE"
	BackupPlanExecutionStatusNEVEREXECUTE BackupPlanExecutionStatus = "NEVER_EXECUTE"

	// BackupPlanExecutionStatusPARTIALSUCCESS captures enum value "PARTIAL_SUCCESS"
	BackupPlanExecutionStatusPARTIALSUCCESS BackupPlanExecutionStatus = "PARTIAL_SUCCESS"

	// BackupPlanExecutionStatusPAUSED captures enum value "PAUSED"
	BackupPlanExecutionStatusPAUSED BackupPlanExecutionStatus = "PAUSED"

	// BackupPlanExecutionStatusRUNNING captures enum value "RUNNING"
	BackupPlanExecutionStatusRUNNING BackupPlanExecutionStatus = "RUNNING"

	// BackupPlanExecutionStatusSUCCESS captures enum value "SUCCESS"
	BackupPlanExecutionStatusSUCCESS BackupPlanExecutionStatus = "SUCCESS"
)

// for schema
var backupPlanExecutionStatusEnum []interface{}

func init() {
	var res []BackupPlanExecutionStatus
	if err := json.Unmarshal([]byte(`["FAILED","NEVER_EXECUTE","PARTIAL_SUCCESS","PAUSED","RUNNING","SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupPlanExecutionStatusEnum = append(backupPlanExecutionStatusEnum, v)
	}
}

func (m BackupPlanExecutionStatus) validateBackupPlanExecutionStatusEnum(path, location string, value BackupPlanExecutionStatus) error {
	if err := validate.EnumCase(path, location, value, backupPlanExecutionStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this backup plan execution status
func (m BackupPlanExecutionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBackupPlanExecutionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this backup plan execution status based on context it is used
func (m BackupPlanExecutionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
