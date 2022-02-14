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
	"github.com/go-openapi/validate"
)

// BackupPlanExecution backup plan execution
//
// swagger:model BackupPlanExecution
type BackupPlanExecution struct {

	// backup plan
	BackupPlan struct {
		NestedBackupPlan
	} `json:"backup_plan,omitempty"`

	// backup restore points
	BackupRestorePoints []*NestedBackupRestorePoint `json:"backup_restore_points,omitempty"`

	// backup target executions
	BackupTargetExecutions []*NestedBackupTargetExecution `json:"backup_target_executions,omitempty"`

	// duration
	Duration *int32 `json:"duration,omitempty"`

	// entity async status
	EntityAsyncStatus struct {
		EntityAsyncStatus
	} `json:"entityAsyncStatus,omitempty"`

	// executed at
	ExecutedAt *string `json:"executed_at,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local created at
	// Required: true
	LocalCreatedAt *string `json:"local_created_at"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// method
	// Required: true
	Method *BackupExecutionMethod `json:"method"`

	// state
	State struct {
		BackupPlanExecutionState
	} `json:"state,omitempty"`

	// status
	// Required: true
	Status *BackupPlanExecutionStatus `json:"status"`

	// success job count
	SuccessJobCount *int32 `json:"success_job_count,omitempty"`

	// total job count
	TotalJobCount *int32 `json:"total_job_count,omitempty"`

	// type
	// Required: true
	Type *BackupExecutionType `json:"type"`
}

// Validate validates this backup plan execution
func (m *BackupPlanExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupRestorePoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackupTargetExecutions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanExecution) validateBackupPlan(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupPlan) { // not required
		return nil
	}

	return nil
}

func (m *BackupPlanExecution) validateBackupRestorePoints(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRestorePoints) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupRestorePoints); i++ {
		if swag.IsZero(m.BackupRestorePoints[i]) { // not required
			continue
		}

		if m.BackupRestorePoints[i] != nil {
			if err := m.BackupRestorePoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_restore_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_restore_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlanExecution) validateBackupTargetExecutions(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupTargetExecutions) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupTargetExecutions); i++ {
		if swag.IsZero(m.BackupTargetExecutions[i]) { // not required
			continue
		}

		if m.BackupTargetExecutions[i] != nil {
			if err := m.BackupTargetExecutions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_target_executions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_target_executions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlanExecution) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	return nil
}

func (m *BackupPlanExecution) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlanExecution) validateLocalCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("local_created_at", "body", m.LocalCreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlanExecution) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *BackupPlanExecution) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	if m.Method != nil {
		if err := m.Method.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanExecution) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	return nil
}

func (m *BackupPlanExecution) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanExecution) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this backup plan execution based on the context it is used
func (m *BackupPlanExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupPlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupRestorePoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackupTargetExecutions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupPlanExecution) contextValidateBackupPlan(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupPlanExecution) contextValidateBackupRestorePoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupRestorePoints); i++ {

		if m.BackupRestorePoints[i] != nil {
			if err := m.BackupRestorePoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_restore_points" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_restore_points" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlanExecution) contextValidateBackupTargetExecutions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupTargetExecutions); i++ {

		if m.BackupTargetExecutions[i] != nil {
			if err := m.BackupTargetExecutions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backup_target_executions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backup_target_executions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupPlanExecution) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupPlanExecution) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.Method != nil {
		if err := m.Method.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanExecution) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *BackupPlanExecution) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *BackupPlanExecution) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupPlanExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupPlanExecution) UnmarshalBinary(b []byte) error {
	var res BackupPlanExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
