// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ElfStoragePolicy elf storage policy
//
// swagger:model ElfStoragePolicy
type ElfStoragePolicy struct {

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// description
	// Required: true
	Description *string `json:"description"`

	// entity async status
	EntityAsyncStatus struct {
		EntityAsyncStatus
	} `json:"entityAsyncStatus,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// local id
	// Required: true
	LocalID *string `json:"local_id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// replica num
	// Required: true
	ReplicaNum *int32 `json:"replica_num"`

	// stripe num
	// Required: true
	StripeNum *int32 `json:"stripe_num"`

	// stripe size
	// Required: true
	StripeSize *float64 `json:"stripe_size"`

	// thin provision
	// Required: true
	ThinProvision *bool `json:"thin_provision"`
}

// Validate validates this elf storage policy
func (m *ElfStoragePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeNum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStripeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThinProvision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfStoragePolicy) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ElfStoragePolicy) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ElfStoragePolicy) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	return nil
}

func (m *ElfStoragePolicy) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ElfStoragePolicy) validateLocalID(formats strfmt.Registry) error {

	if err := validate.Required("local_id", "body", m.LocalID); err != nil {
		return err
	}

	return nil
}

func (m *ElfStoragePolicy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ElfStoragePolicy) validateReplicaNum(formats strfmt.Registry) error {

	if err := validate.Required("replica_num", "body", m.ReplicaNum); err != nil {
		return err
	}

	return nil
}

func (m *ElfStoragePolicy) validateStripeNum(formats strfmt.Registry) error {

	if err := validate.Required("stripe_num", "body", m.StripeNum); err != nil {
		return err
	}

	return nil
}

func (m *ElfStoragePolicy) validateStripeSize(formats strfmt.Registry) error {

	if err := validate.Required("stripe_size", "body", m.StripeSize); err != nil {
		return err
	}

	return nil
}

func (m *ElfStoragePolicy) validateThinProvision(formats strfmt.Registry) error {

	if err := validate.Required("thin_provision", "body", m.ThinProvision); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this elf storage policy based on the context it is used
func (m *ElfStoragePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElfStoragePolicy) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *ElfStoragePolicy) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ElfStoragePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElfStoragePolicy) UnmarshalBinary(b []byte) error {
	var res ElfStoragePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
