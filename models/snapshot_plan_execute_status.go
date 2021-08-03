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

// SnapshotPlanExecuteStatus snapshot plan execute status
//
// swagger:model SnapshotPlanExecuteStatus
type SnapshotPlanExecuteStatus string

func NewSnapshotPlanExecuteStatus(value SnapshotPlanExecuteStatus) *SnapshotPlanExecuteStatus {
	v := value
	return &v
}

const (

	// SnapshotPlanExecuteStatusCREATED captures enum value "CREATED"
	SnapshotPlanExecuteStatusCREATED SnapshotPlanExecuteStatus = "CREATED"

	// SnapshotPlanExecuteStatusFAILED captures enum value "FAILED"
	SnapshotPlanExecuteStatusFAILED SnapshotPlanExecuteStatus = "FAILED"

	// SnapshotPlanExecuteStatusINPROGRESS captures enum value "IN_PROGRESS"
	SnapshotPlanExecuteStatusINPROGRESS SnapshotPlanExecuteStatus = "IN_PROGRESS"

	// SnapshotPlanExecuteStatusSUCCEED captures enum value "SUCCEED"
	SnapshotPlanExecuteStatusSUCCEED SnapshotPlanExecuteStatus = "SUCCEED"

	// SnapshotPlanExecuteStatusUNSPECIFIED captures enum value "UNSPECIFIED"
	SnapshotPlanExecuteStatusUNSPECIFIED SnapshotPlanExecuteStatus = "UNSPECIFIED"
)

// for schema
var snapshotPlanExecuteStatusEnum []interface{}

func init() {
	var res []SnapshotPlanExecuteStatus
	if err := json.Unmarshal([]byte(`["CREATED","FAILED","IN_PROGRESS","SUCCEED","UNSPECIFIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotPlanExecuteStatusEnum = append(snapshotPlanExecuteStatusEnum, v)
	}
}

func (m SnapshotPlanExecuteStatus) validateSnapshotPlanExecuteStatusEnum(path, location string, value SnapshotPlanExecuteStatus) error {
	if err := validate.EnumCase(path, location, value, snapshotPlanExecuteStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this snapshot plan execute status
func (m SnapshotPlanExecuteStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSnapshotPlanExecuteStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this snapshot plan execute status based on context it is used
func (m SnapshotPlanExecuteStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
