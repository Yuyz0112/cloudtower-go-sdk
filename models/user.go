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

// User user
//
// swagger:model User
type User struct {

	// email address
	EmailAddress *string `json:"email_address,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// internal
	// Required: true
	Internal *bool `json:"internal"`

	// ldap dn
	LdapDn *string `json:"ldap_dn,omitempty"`

	// login info
	LoginInfo *NestedUserLoginInfo `json:"login_info,omitempty"`

	// mobile phone
	MobilePhone *string `json:"mobile_phone,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// password expired
	PasswordExpired *bool `json:"password_expired,omitempty"`

	// password recover qa
	PasswordRecoverQa *NestedPasswordRecoverQa `json:"password_recover_qa,omitempty"`

	// password updated at
	PasswordUpdatedAt *string `json:"password_updated_at,omitempty"`

	// role
	Role *UserRole `json:"role,omitempty"`

	// roles
	Roles []*NestedUserRoleNext `json:"roles,omitempty"`

	// source
	// Required: true
	Source *UserSource `json:"source"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLoginInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordRecoverQa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *User) validateInternal(formats strfmt.Registry) error {

	if err := validate.Required("internal", "body", m.Internal); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLoginInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.LoginInfo) { // not required
		return nil
	}

	if m.LoginInfo != nil {
		if err := m.LoginInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("login_info")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *User) validatePasswordRecoverQa(formats strfmt.Registry) error {
	if swag.IsZero(m.PasswordRecoverQa) { // not required
		return nil
	}

	if m.PasswordRecoverQa != nil {
		if err := m.PasswordRecoverQa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("password_recover_qa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("password_recover_qa")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user based on the context it is used
func (m *User) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoginInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePasswordRecoverQa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) contextValidateLoginInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.LoginInfo != nil {
		if err := m.LoginInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("login_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("login_info")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidatePasswordRecoverQa(ctx context.Context, formats strfmt.Registry) error {

	if m.PasswordRecoverQa != nil {
		if err := m.PasswordRecoverQa.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("password_recover_qa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("password_recover_qa")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {
		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *User) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {
			if err := m.Roles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
