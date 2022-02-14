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

// BackupRestoreExecution backup restore execution
//
// swagger:model BackupRestoreExecution
type BackupRestoreExecution struct {

	// backup restore point
	BackupRestorePoint *NestedBackupRestorePoint `json:"backup_restore_point,omitempty"`

	// duration
	Duration *int32 `json:"duration,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// executed at
	// Required: true
	ExecutedAt *string `json:"executed_at"`

	// id
	// Required: true
	ID *string `json:"id"`

	// mode
	// Required: true
	Mode *BackupRestoreExecutionMode `json:"mode"`

	// name
	// Required: true
	Name *string `json:"name"`

	// read bytes
	ReadBytes *float64 `json:"read_bytes,omitempty"`

	// rebuild name
	RebuildName *string `json:"rebuild_name,omitempty"`

	// rebuild network mapping
	// Required: true
	RebuildNetworkMapping []*NestedBackupRestoreExecutionNetworkMapping `json:"rebuild_network_mapping"`

	// rebuild target cluster
	RebuildTargetCluster *string `json:"rebuild_target_cluster,omitempty"`

	// rebuild target host
	RebuildTargetHost *string `json:"rebuild_target_host,omitempty"`

	// startup after restore
	// Required: true
	StartupAfterRestore *bool `json:"startup_after_restore"`

	// status
	// Required: true
	Status *BackupExecutionStatus `json:"status"`

	// total bytes
	TotalBytes *float64 `json:"total_bytes,omitempty"`
}

// Validate validates this backup restore execution
func (m *BackupRestoreExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupRestorePoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRebuildNetworkMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartupAfterRestore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRestoreExecution) validateBackupRestorePoint(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRestorePoint) { // not required
		return nil
	}

	if m.BackupRestorePoint != nil {
		if err := m.BackupRestorePoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_restore_point")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_restore_point")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestoreExecution) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestoreExecution) validateExecutedAt(formats strfmt.Registry) error {

	if err := validate.Required("executed_at", "body", m.ExecutedAt); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestoreExecution) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestoreExecution) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestoreExecution) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestoreExecution) validateRebuildNetworkMapping(formats strfmt.Registry) error {

	if err := validate.Required("rebuild_network_mapping", "body", m.RebuildNetworkMapping); err != nil {
		return err
	}

	for i := 0; i < len(m.RebuildNetworkMapping); i++ {
		if swag.IsZero(m.RebuildNetworkMapping[i]) { // not required
			continue
		}

		if m.RebuildNetworkMapping[i] != nil {
			if err := m.RebuildNetworkMapping[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rebuild_network_mapping" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rebuild_network_mapping" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupRestoreExecution) validateStartupAfterRestore(formats strfmt.Registry) error {

	if err := validate.Required("startup_after_restore", "body", m.StartupAfterRestore); err != nil {
		return err
	}

	return nil
}

func (m *BackupRestoreExecution) validateStatus(formats strfmt.Registry) error {

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

// ContextValidate validate this backup restore execution based on the context it is used
func (m *BackupRestoreExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupRestorePoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRebuildNetworkMapping(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRestoreExecution) contextValidateBackupRestorePoint(ctx context.Context, formats strfmt.Registry) error {

	if m.BackupRestorePoint != nil {
		if err := m.BackupRestorePoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backup_restore_point")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backup_restore_point")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestoreExecution) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestoreExecution) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if m.Mode != nil {
		if err := m.Mode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRestoreExecution) contextValidateRebuildNetworkMapping(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RebuildNetworkMapping); i++ {

		if m.RebuildNetworkMapping[i] != nil {
			if err := m.RebuildNetworkMapping[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rebuild_network_mapping" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rebuild_network_mapping" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupRestoreExecution) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *BackupRestoreExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRestoreExecution) UnmarshalBinary(b []byte) error {
	var res BackupRestoreExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
