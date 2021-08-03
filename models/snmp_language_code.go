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

// SnmpLanguageCode snmp language code
//
// swagger:model SnmpLanguageCode
type SnmpLanguageCode string

func NewSnmpLanguageCode(value SnmpLanguageCode) *SnmpLanguageCode {
	v := value
	return &v
}

const (

	// SnmpLanguageCodeENUS captures enum value "EN_US"
	SnmpLanguageCodeENUS SnmpLanguageCode = "EN_US"

	// SnmpLanguageCodeZHCN captures enum value "ZH_CN"
	SnmpLanguageCodeZHCN SnmpLanguageCode = "ZH_CN"
)

// for schema
var snmpLanguageCodeEnum []interface{}

func init() {
	var res []SnmpLanguageCode
	if err := json.Unmarshal([]byte(`["EN_US","ZH_CN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snmpLanguageCodeEnum = append(snmpLanguageCodeEnum, v)
	}
}

func (m SnmpLanguageCode) validateSnmpLanguageCodeEnum(path, location string, value SnmpLanguageCode) error {
	if err := validate.EnumCase(path, location, value, snmpLanguageCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this snmp language code
func (m SnmpLanguageCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSnmpLanguageCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this snmp language code based on context it is used
func (m SnmpLanguageCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
