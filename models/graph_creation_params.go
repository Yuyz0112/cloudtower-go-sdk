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

// GraphCreationParams graph creation params
//
// swagger:model GraphCreationParams
type GraphCreationParams struct {

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// connect id
	// Required: true
	ConnectID []string `json:"connect_id"`

	// instance ids
	InstanceIds []string `json:"instance_ids"`

	// metric count
	MetricCount int32 `json:"metric_count,omitempty"`

	// metric name
	// Required: true
	MetricName *string `json:"metric_name"`

	// metric type
	MetricType MetricType `json:"metric_type,omitempty"`

	// network
	Network NetworkType `json:"network,omitempty"`

	// resource type
	// Required: true
	ResourceType *string `json:"resource_type"`

	// service
	Service string `json:"service,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// type
	// Required: true
	Type *GraphType `json:"type"`

	// view id
	// Required: true
	ViewID *string `json:"view_id"`
}

// Validate validates this graph creation params
func (m *GraphCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GraphCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *GraphCreationParams) validateConnectID(formats strfmt.Registry) error {

	if err := validate.Required("connect_id", "body", m.ConnectID); err != nil {
		return err
	}

	return nil
}

func (m *GraphCreationParams) validateMetricName(formats strfmt.Registry) error {

	if err := validate.Required("metric_name", "body", m.MetricName); err != nil {
		return err
	}

	return nil
}

func (m *GraphCreationParams) validateMetricType(formats strfmt.Registry) error {
	if swag.IsZero(m.MetricType) { // not required
		return nil
	}

	if err := m.MetricType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metric_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("metric_type")
		}
		return err
	}

	return nil
}

func (m *GraphCreationParams) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if err := m.Network.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("network")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("network")
		}
		return err
	}

	return nil
}

func (m *GraphCreationParams) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *GraphCreationParams) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *GraphCreationParams) validateType(formats strfmt.Registry) error {

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

func (m *GraphCreationParams) validateViewID(formats strfmt.Registry) error {

	if err := validate.Required("view_id", "body", m.ViewID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this graph creation params based on the context it is used
func (m *GraphCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetricType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetwork(ctx, formats); err != nil {
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

func (m *GraphCreationParams) contextValidateMetricType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MetricType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metric_type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("metric_type")
		}
		return err
	}

	return nil
}

func (m *GraphCreationParams) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Network.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("network")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("network")
		}
		return err
	}

	return nil
}

func (m *GraphCreationParams) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GraphCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GraphCreationParams) UnmarshalBinary(b []byte) error {
	var res GraphCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
