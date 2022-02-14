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

// TaskStatus task status
//
// swagger:model TaskStatus
type TaskStatus string

func NewTaskStatus(value TaskStatus) *TaskStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TaskStatus.
func (m TaskStatus) Pointer() *TaskStatus {
	return &m
}

const (

	// TaskStatusEXECUTING captures enum value "EXECUTING"
	TaskStatusEXECUTING TaskStatus = "EXECUTING"

	// TaskStatusFAILED captures enum value "FAILED"
	TaskStatusFAILED TaskStatus = "FAILED"

	// TaskStatusPENDING captures enum value "PENDING"
	TaskStatusPENDING TaskStatus = "PENDING"

	// TaskStatusSUCCESSED captures enum value "SUCCESSED"
	TaskStatusSUCCESSED TaskStatus = "SUCCESSED"
)

// for schema
var taskStatusEnum []interface{}

func init() {
	var res []TaskStatus
	if err := json.Unmarshal([]byte(`["EXECUTING","FAILED","PENDING","SUCCESSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskStatusEnum = append(taskStatusEnum, v)
	}
}

func (m TaskStatus) validateTaskStatusEnum(path, location string, value TaskStatus) error {
	if err := validate.EnumCase(path, location, value, taskStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this task status
func (m TaskStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTaskStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this task status based on context it is used
func (m TaskStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
