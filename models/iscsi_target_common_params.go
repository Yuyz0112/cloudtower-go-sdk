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

// IscsiTargetCommonParams iscsi target common params
//
// swagger:model IscsiTargetCommonParams
type IscsiTargetCommonParams struct {

	// bps
	Bps float64 `json:"bps,omitempty"`

	// bps max
	BpsMax float64 `json:"bps_max,omitempty"`

	// bps max length
	BpsMaxLength float64 `json:"bps_max_length,omitempty"`

	// bps rd
	BpsRd float64 `json:"bps_rd,omitempty"`

	// bps rd max
	BpsRdMax float64 `json:"bps_rd_max,omitempty"`

	// bps rd max length
	BpsRdMaxLength float64 `json:"bps_rd_max_length,omitempty"`

	// bps wr
	BpsWr float64 `json:"bps_wr,omitempty"`

	// bps wr max
	BpsWrMax float64 `json:"bps_wr_max,omitempty"`

	// bps wr max length
	BpsWrMaxLength float64 `json:"bps_wr_max_length,omitempty"`

	// chap enabled
	ChapEnabled bool `json:"chap_enabled,omitempty"`

	// chap name
	ChapName string `json:"chap_name,omitempty"`

	// chap secret
	ChapSecret string `json:"chap_secret,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// initiator chaps
	InitiatorChaps []*IscsiTargetCommonParamsInitiatorChapsItems0 `json:"initiator_chaps"`

	// iops
	Iops float64 `json:"iops,omitempty"`

	// iops max
	IopsMax float64 `json:"iops_max,omitempty"`

	// iops max length
	IopsMaxLength float64 `json:"iops_max_length,omitempty"`

	// iops rd
	IopsRd float64 `json:"iops_rd,omitempty"`

	// iops rd max
	IopsRdMax float64 `json:"iops_rd_max,omitempty"`

	// iops rd max length
	IopsRdMaxLength float64 `json:"iops_rd_max_length,omitempty"`

	// iops wr
	IopsWr float64 `json:"iops_wr,omitempty"`

	// iops wr max
	IopsWrMax float64 `json:"iops_wr_max,omitempty"`

	// iops wr max length
	IopsWrMaxLength float64 `json:"iops_wr_max_length,omitempty"`

	// ip whitelist
	IPWhitelist string `json:"ip_whitelist,omitempty"`

	// iqn whitelist
	IqnWhitelist string `json:"iqn_whitelist,omitempty"`
}

// Validate validates this iscsi target common params
func (m *IscsiTargetCommonParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInitiatorChaps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiTargetCommonParams) validateInitiatorChaps(formats strfmt.Registry) error {
	if swag.IsZero(m.InitiatorChaps) { // not required
		return nil
	}

	for i := 0; i < len(m.InitiatorChaps); i++ {
		if swag.IsZero(m.InitiatorChaps[i]) { // not required
			continue
		}

		if m.InitiatorChaps[i] != nil {
			if err := m.InitiatorChaps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_chaps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this iscsi target common params based on the context it is used
func (m *IscsiTargetCommonParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInitiatorChaps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiTargetCommonParams) contextValidateInitiatorChaps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InitiatorChaps); i++ {

		if m.InitiatorChaps[i] != nil {
			if err := m.InitiatorChaps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_chaps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IscsiTargetCommonParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiTargetCommonParams) UnmarshalBinary(b []byte) error {
	var res IscsiTargetCommonParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IscsiTargetCommonParamsInitiatorChapsItems0 iscsi target common params initiator chaps items0
//
// swagger:model IscsiTargetCommonParamsInitiatorChapsItems0
type IscsiTargetCommonParamsInitiatorChapsItems0 struct {

	// chap name
	// Required: true
	ChapName *string `json:"chap_name"`

	// chap secret
	// Required: true
	ChapSecret *string `json:"chap_secret"`

	// initiator iqn
	// Required: true
	InitiatorIqn *string `json:"initiator_iqn"`
}

// Validate validates this iscsi target common params initiator chaps items0
func (m *IscsiTargetCommonParamsInitiatorChapsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChapName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChapSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatorIqn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiTargetCommonParamsInitiatorChapsItems0) validateChapName(formats strfmt.Registry) error {

	if err := validate.Required("chap_name", "body", m.ChapName); err != nil {
		return err
	}

	return nil
}

func (m *IscsiTargetCommonParamsInitiatorChapsItems0) validateChapSecret(formats strfmt.Registry) error {

	if err := validate.Required("chap_secret", "body", m.ChapSecret); err != nil {
		return err
	}

	return nil
}

func (m *IscsiTargetCommonParamsInitiatorChapsItems0) validateInitiatorIqn(formats strfmt.Registry) error {

	if err := validate.Required("initiator_iqn", "body", m.InitiatorIqn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi target common params initiator chaps items0 based on context it is used
func (m *IscsiTargetCommonParamsInitiatorChapsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IscsiTargetCommonParamsInitiatorChapsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiTargetCommonParamsInitiatorChapsItems0) UnmarshalBinary(b []byte) error {
	var res IscsiTargetCommonParamsInitiatorChapsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
