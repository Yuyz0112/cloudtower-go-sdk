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

// RackTopoUpdationParams rack topo updation params
//
// swagger:model RackTopoUpdationParams
type RackTopoUpdationParams struct {

	// data
	// Required: true
	Data *RackTopoUpdationParamsData `json:"data"`

	// where
	// Required: true
	Where *RackTopoWhereInput `json:"where"`
}

// Validate validates this rack topo updation params
func (m *RackTopoUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopoUpdationParams) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RackTopoUpdationParams) validateWhere(formats strfmt.Registry) error {

	if err := validate.Required("where", "body", m.Where); err != nil {
		return err
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rack topo updation params based on the context it is used
func (m *RackTopoUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopoUpdationParams) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *RackTopoUpdationParams) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackTopoUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTopoUpdationParams) UnmarshalBinary(b []byte) error {
	var res RackTopoUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RackTopoUpdationParamsData rack topo updation params data
//
// swagger:model RackTopoUpdationParamsData
type RackTopoUpdationParamsData struct {

	// brick topoes
	BrickTopoes *BrickTopoWhereInput `json:"brick_topoes,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// height
	Height int32 `json:"height,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// zone topo id
	ZoneTopoID string `json:"zone_topo_id,omitempty"`
}

// Validate validates this rack topo updation params data
func (m *RackTopoUpdationParamsData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrickTopoes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopoUpdationParamsData) validateBrickTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.BrickTopoes) { // not required
		return nil
	}

	if m.BrickTopoes != nil {
		if err := m.BrickTopoes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "brick_topoes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "brick_topoes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rack topo updation params data based on the context it is used
func (m *RackTopoUpdationParamsData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBrickTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RackTopoUpdationParamsData) contextValidateBrickTopoes(ctx context.Context, formats strfmt.Registry) error {

	if m.BrickTopoes != nil {
		if err := m.BrickTopoes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "brick_topoes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "brick_topoes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RackTopoUpdationParamsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RackTopoUpdationParamsData) UnmarshalBinary(b []byte) error {
	var res RackTopoUpdationParamsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
