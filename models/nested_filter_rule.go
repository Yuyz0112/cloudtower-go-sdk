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

// NestedFilterRule nested filter rule
//
// swagger:model NestedFilterRule
type NestedFilterRule struct {

	// aggregation
	// Required: true
	Aggregation *FilterRuleAggregationEnum `json:"aggregation"`

	// duration
	// Required: true
	Duration *int32 `json:"duration"`

	// metric
	// Required: true
	Metric *FilterRuleMetricEnum `json:"metric"`

	// op
	// Required: true
	Op *FilterRuleOpEnum `json:"op"`

	// quantile
	// Required: true
	Quantile *int32 `json:"quantile"`

	// threshold
	// Required: true
	Threshold *float64 `json:"threshold"`
}

// Validate validates this nested filter rule
func (m *NestedFilterRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFilterRule) validateAggregation(formats strfmt.Registry) error {

	if err := validate.Required("aggregation", "body", m.Aggregation); err != nil {
		return err
	}

	if err := validate.Required("aggregation", "body", m.Aggregation); err != nil {
		return err
	}

	if m.Aggregation != nil {
		if err := m.Aggregation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *NestedFilterRule) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	if m.Op != nil {
		if err := m.Op.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) validateQuantile(formats strfmt.Registry) error {

	if err := validate.Required("quantile", "body", m.Quantile); err != nil {
		return err
	}

	return nil
}

func (m *NestedFilterRule) validateThreshold(formats strfmt.Registry) error {

	if err := validate.Required("threshold", "body", m.Threshold); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested filter rule based on the context it is used
func (m *NestedFilterRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetric(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedFilterRule) contextValidateAggregation(ctx context.Context, formats strfmt.Registry) error {

	if m.Aggregation != nil {
		if err := m.Aggregation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregation")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) contextValidateMetric(ctx context.Context, formats strfmt.Registry) error {

	if m.Metric != nil {
		if err := m.Metric.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *NestedFilterRule) contextValidateOp(ctx context.Context, formats strfmt.Registry) error {

	if m.Op != nil {
		if err := m.Op.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("op")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedFilterRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedFilterRule) UnmarshalBinary(b []byte) error {
	var res NestedFilterRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
