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

// CloudInitNetworkTypeEnum cloud init network type enum
//
// swagger:model CloudInitNetworkTypeEnum
type CloudInitNetworkTypeEnum string

func NewCloudInitNetworkTypeEnum(value CloudInitNetworkTypeEnum) *CloudInitNetworkTypeEnum {
	v := value
	return &v
}

const (

	// CloudInitNetworkTypeEnumIPV4 captures enum value "IPV4"
	CloudInitNetworkTypeEnumIPV4 CloudInitNetworkTypeEnum = "IPV4"

	// CloudInitNetworkTypeEnumIPV4DHCP captures enum value "IPV4_DHCP"
	CloudInitNetworkTypeEnumIPV4DHCP CloudInitNetworkTypeEnum = "IPV4_DHCP"
)

// for schema
var cloudInitNetworkTypeEnumEnum []interface{}

func init() {
	var res []CloudInitNetworkTypeEnum
	if err := json.Unmarshal([]byte(`["IPV4","IPV4_DHCP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudInitNetworkTypeEnumEnum = append(cloudInitNetworkTypeEnumEnum, v)
	}
}

func (m CloudInitNetworkTypeEnum) validateCloudInitNetworkTypeEnumEnum(path, location string, value CloudInitNetworkTypeEnum) error {
	if err := validate.EnumCase(path, location, value, cloudInitNetworkTypeEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cloud init network type enum
func (m CloudInitNetworkTypeEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCloudInitNetworkTypeEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cloud init network type enum based on context it is used
func (m CloudInitNetworkTypeEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}