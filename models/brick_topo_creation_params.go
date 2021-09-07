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

// BrickTopoCreationParams brick topo creation params
//
// swagger:model BrickTopoCreationParams
type BrickTopoCreationParams struct {

	// capacity
	Capacity *BrickTopoCreationParamsCapacity `json:"capacity,omitempty"`

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// height
	// Required: true
	Height *float64 `json:"height"`

	// name
	// Required: true
	Name *string `json:"name"`

	// node topoes
	NodeTopoes *NodeTopoWhereInput `json:"node_topoes,omitempty"`

	// position
	// Required: true
	Position *float64 `json:"position"`

	// rack topo id
	RackTopoID string `json:"rack_topo_id,omitempty"`

	// tag position in brick
	TagPositionInBrick []*BrickTopoCreationParamsTagPositionInBrickItems0 `json:"tag_position_in_brick,omitempty"`
}

// Validate validates this brick topo creation params
func (m *BrickTopoCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeTopoes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagPositionInBrick(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoCreationParams) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParams) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParams) validateNodeTopoes(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeTopoes) { // not required
		return nil
	}

	if m.NodeTopoes != nil {
		if err := m.NodeTopoes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_topoes")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoCreationParams) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParams) validateTagPositionInBrick(formats strfmt.Registry) error {
	if swag.IsZero(m.TagPositionInBrick) { // not required
		return nil
	}

	for i := 0; i < len(m.TagPositionInBrick); i++ {
		if swag.IsZero(m.TagPositionInBrick[i]) { // not required
			continue
		}

		if m.TagPositionInBrick[i] != nil {
			if err := m.TagPositionInBrick[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag_position_in_brick" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this brick topo creation params based on the context it is used
func (m *BrickTopoCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeTopoes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTagPositionInBrick(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoCreationParams) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	if m.Capacity != nil {
		if err := m.Capacity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoCreationParams) contextValidateNodeTopoes(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeTopoes != nil {
		if err := m.NodeTopoes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_topoes")
			}
			return err
		}
	}

	return nil
}

func (m *BrickTopoCreationParams) contextValidateTagPositionInBrick(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TagPositionInBrick); i++ {

		if m.TagPositionInBrick[i] != nil {
			if err := m.TagPositionInBrick[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tag_position_in_brick" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BrickTopoCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrickTopoCreationParams) UnmarshalBinary(b []byte) error {
	var res BrickTopoCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BrickTopoCreationParamsCapacity brick topo creation params capacity
//
// swagger:model BrickTopoCreationParamsCapacity
type BrickTopoCreationParamsCapacity struct {

	// column
	Column *float64 `json:"column,omitempty"`

	// row
	Row *float64 `json:"row,omitempty"`
}

// Validate validates this brick topo creation params capacity
func (m *BrickTopoCreationParamsCapacity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this brick topo creation params capacity based on context it is used
func (m *BrickTopoCreationParamsCapacity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BrickTopoCreationParamsCapacity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrickTopoCreationParamsCapacity) UnmarshalBinary(b []byte) error {
	var res BrickTopoCreationParamsCapacity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BrickTopoCreationParamsTagPositionInBrickItems0 brick topo creation params tag position in brick items0
//
// swagger:model BrickTopoCreationParamsTagPositionInBrickItems0
type BrickTopoCreationParamsTagPositionInBrickItems0 struct {

	// column
	// Required: true
	Column *float64 `json:"column"`

	// row
	// Required: true
	Row *float64 `json:"row"`

	// tag
	// Required: true
	Tag *string `json:"tag"`
}

// Validate validates this brick topo creation params tag position in brick items0
func (m *BrickTopoCreationParamsTagPositionInBrickItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColumn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BrickTopoCreationParamsTagPositionInBrickItems0) validateColumn(formats strfmt.Registry) error {

	if err := validate.Required("column", "body", m.Column); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParamsTagPositionInBrickItems0) validateRow(formats strfmt.Registry) error {

	if err := validate.Required("row", "body", m.Row); err != nil {
		return err
	}

	return nil
}

func (m *BrickTopoCreationParamsTagPositionInBrickItems0) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this brick topo creation params tag position in brick items0 based on context it is used
func (m *BrickTopoCreationParamsTagPositionInBrickItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BrickTopoCreationParamsTagPositionInBrickItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BrickTopoCreationParamsTagPositionInBrickItems0) UnmarshalBinary(b []byte) error {
	var res BrickTopoCreationParamsTagPositionInBrickItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}