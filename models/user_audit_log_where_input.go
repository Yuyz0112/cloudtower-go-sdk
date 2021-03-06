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
)

// UserAuditLogWhereInput user audit log where input
// Example: {"AND":"UserAuditLogWhereInput[]","NOT":"UserAuditLogWhereInput[]","OR":"UserAuditLogWhereInput[]","action":"string","action_contains":"string","action_ends_with":"string","action_gt":"string","action_gte":"string","action_in":["string"],"action_lt":"string","action_lte":"string","action_not":"string","action_not_contains":"string","action_not_ends_with":"string","action_not_in":["string"],"action_not_starts_with":"string","action_starts_with":"string","cluster":"ClusterWhereInput","createdAt":"string","createdAt_gt":"string","createdAt_gte":"string","createdAt_in":["string"],"createdAt_lt":"string","createdAt_lte":"string","createdAt_not":"string","createdAt_not_in":["string"],"finished_at":"string","finished_at_gt":"string","finished_at_gte":"string","finished_at_in":["string"],"finished_at_lt":"string","finished_at_lte":"string","finished_at_not":"string","finished_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","ip_address":"string","ip_address_contains":"string","ip_address_ends_with":"string","ip_address_gt":"string","ip_address_gte":"string","ip_address_in":["string"],"ip_address_lt":"string","ip_address_lte":"string","ip_address_not":"string","ip_address_not_contains":"string","ip_address_not_ends_with":"string","ip_address_not_in":["string"],"ip_address_not_starts_with":"string","ip_address_starts_with":"string","message":"string","message_contains":"string","message_ends_with":"string","message_gt":"string","message_gte":"string","message_in":["string"],"message_lt":"string","message_lte":"string","message_not":"string","message_not_contains":"string","message_not_ends_with":"string","message_not_in":["string"],"message_not_starts_with":"string","message_starts_with":"string","resource_id":"string","resource_id_contains":"string","resource_id_ends_with":"string","resource_id_gt":"string","resource_id_gte":"string","resource_id_in":["string"],"resource_id_lt":"string","resource_id_lte":"string","resource_id_not":"string","resource_id_not_contains":"string","resource_id_not_ends_with":"string","resource_id_not_in":["string"],"resource_id_not_starts_with":"string","resource_id_starts_with":"string","resource_type":"string","resource_type_contains":"string","resource_type_ends_with":"string","resource_type_gt":"string","resource_type_gte":"string","resource_type_in":["string"],"resource_type_lt":"string","resource_type_lte":"string","resource_type_not":"string","resource_type_not_contains":"string","resource_type_not_ends_with":"string","resource_type_not_in":["string"],"resource_type_not_starts_with":"string","resource_type_starts_with":"string","status":"FAILED","status_in":["FAILED"],"status_not":"FAILED","status_not_in":["FAILED"],"user":"UserWhereInput"}
//
// swagger:model UserAuditLogWhereInput
type UserAuditLogWhereInput struct {

	// a n d
	AND []*UserAuditLogWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*UserAuditLogWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*UserAuditLogWhereInput `json:"OR,omitempty"`

	// action
	Action *string `json:"action,omitempty"`

	// action contains
	ActionContains *string `json:"action_contains,omitempty"`

	// action ends with
	ActionEndsWith *string `json:"action_ends_with,omitempty"`

	// action gt
	ActionGt *string `json:"action_gt,omitempty"`

	// action gte
	ActionGte *string `json:"action_gte,omitempty"`

	// action in
	ActionIn []string `json:"action_in,omitempty"`

	// action lt
	ActionLt *string `json:"action_lt,omitempty"`

	// action lte
	ActionLte *string `json:"action_lte,omitempty"`

	// action not
	ActionNot *string `json:"action_not,omitempty"`

	// action not contains
	ActionNotContains *string `json:"action_not_contains,omitempty"`

	// action not ends with
	ActionNotEndsWith *string `json:"action_not_ends_with,omitempty"`

	// action not in
	ActionNotIn []string `json:"action_not_in,omitempty"`

	// action not starts with
	ActionNotStartsWith *string `json:"action_not_starts_with,omitempty"`

	// action starts with
	ActionStartsWith *string `json:"action_starts_with,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// created at
	CreatedAt *string `json:"createdAt,omitempty"`

	// created at gt
	CreatedAtGt *string `json:"createdAt_gt,omitempty"`

	// created at gte
	CreatedAtGte *string `json:"createdAt_gte,omitempty"`

	// created at in
	CreatedAtIn []string `json:"createdAt_in,omitempty"`

	// created at lt
	CreatedAtLt *string `json:"createdAt_lt,omitempty"`

	// created at lte
	CreatedAtLte *string `json:"createdAt_lte,omitempty"`

	// created at not
	CreatedAtNot *string `json:"createdAt_not,omitempty"`

	// created at not in
	CreatedAtNotIn []string `json:"createdAt_not_in,omitempty"`

	// finished at
	FinishedAt *string `json:"finished_at,omitempty"`

	// finished at gt
	FinishedAtGt *string `json:"finished_at_gt,omitempty"`

	// finished at gte
	FinishedAtGte *string `json:"finished_at_gte,omitempty"`

	// finished at in
	FinishedAtIn []string `json:"finished_at_in,omitempty"`

	// finished at lt
	FinishedAtLt *string `json:"finished_at_lt,omitempty"`

	// finished at lte
	FinishedAtLte *string `json:"finished_at_lte,omitempty"`

	// finished at not
	FinishedAtNot *string `json:"finished_at_not,omitempty"`

	// finished at not in
	FinishedAtNotIn []string `json:"finished_at_not_in,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// ip address
	IPAddress *string `json:"ip_address,omitempty"`

	// ip address contains
	IPAddressContains *string `json:"ip_address_contains,omitempty"`

	// ip address ends with
	IPAddressEndsWith *string `json:"ip_address_ends_with,omitempty"`

	// ip address gt
	IPAddressGt *string `json:"ip_address_gt,omitempty"`

	// ip address gte
	IPAddressGte *string `json:"ip_address_gte,omitempty"`

	// ip address in
	IPAddressIn []string `json:"ip_address_in,omitempty"`

	// ip address lt
	IPAddressLt *string `json:"ip_address_lt,omitempty"`

	// ip address lte
	IPAddressLte *string `json:"ip_address_lte,omitempty"`

	// ip address not
	IPAddressNot *string `json:"ip_address_not,omitempty"`

	// ip address not contains
	IPAddressNotContains *string `json:"ip_address_not_contains,omitempty"`

	// ip address not ends with
	IPAddressNotEndsWith *string `json:"ip_address_not_ends_with,omitempty"`

	// ip address not in
	IPAddressNotIn []string `json:"ip_address_not_in,omitempty"`

	// ip address not starts with
	IPAddressNotStartsWith *string `json:"ip_address_not_starts_with,omitempty"`

	// ip address starts with
	IPAddressStartsWith *string `json:"ip_address_starts_with,omitempty"`

	// message
	Message *string `json:"message,omitempty"`

	// message contains
	MessageContains *string `json:"message_contains,omitempty"`

	// message ends with
	MessageEndsWith *string `json:"message_ends_with,omitempty"`

	// message gt
	MessageGt *string `json:"message_gt,omitempty"`

	// message gte
	MessageGte *string `json:"message_gte,omitempty"`

	// message in
	MessageIn []string `json:"message_in,omitempty"`

	// message lt
	MessageLt *string `json:"message_lt,omitempty"`

	// message lte
	MessageLte *string `json:"message_lte,omitempty"`

	// message not
	MessageNot *string `json:"message_not,omitempty"`

	// message not contains
	MessageNotContains *string `json:"message_not_contains,omitempty"`

	// message not ends with
	MessageNotEndsWith *string `json:"message_not_ends_with,omitempty"`

	// message not in
	MessageNotIn []string `json:"message_not_in,omitempty"`

	// message not starts with
	MessageNotStartsWith *string `json:"message_not_starts_with,omitempty"`

	// message starts with
	MessageStartsWith *string `json:"message_starts_with,omitempty"`

	// resource id
	ResourceID *string `json:"resource_id,omitempty"`

	// resource id contains
	ResourceIDContains *string `json:"resource_id_contains,omitempty"`

	// resource id ends with
	ResourceIDEndsWith *string `json:"resource_id_ends_with,omitempty"`

	// resource id gt
	ResourceIDGt *string `json:"resource_id_gt,omitempty"`

	// resource id gte
	ResourceIDGte *string `json:"resource_id_gte,omitempty"`

	// resource id in
	ResourceIDIn []string `json:"resource_id_in,omitempty"`

	// resource id lt
	ResourceIDLt *string `json:"resource_id_lt,omitempty"`

	// resource id lte
	ResourceIDLte *string `json:"resource_id_lte,omitempty"`

	// resource id not
	ResourceIDNot *string `json:"resource_id_not,omitempty"`

	// resource id not contains
	ResourceIDNotContains *string `json:"resource_id_not_contains,omitempty"`

	// resource id not ends with
	ResourceIDNotEndsWith *string `json:"resource_id_not_ends_with,omitempty"`

	// resource id not in
	ResourceIDNotIn []string `json:"resource_id_not_in,omitempty"`

	// resource id not starts with
	ResourceIDNotStartsWith *string `json:"resource_id_not_starts_with,omitempty"`

	// resource id starts with
	ResourceIDStartsWith *string `json:"resource_id_starts_with,omitempty"`

	// resource type
	ResourceType *string `json:"resource_type,omitempty"`

	// resource type contains
	ResourceTypeContains *string `json:"resource_type_contains,omitempty"`

	// resource type ends with
	ResourceTypeEndsWith *string `json:"resource_type_ends_with,omitempty"`

	// resource type gt
	ResourceTypeGt *string `json:"resource_type_gt,omitempty"`

	// resource type gte
	ResourceTypeGte *string `json:"resource_type_gte,omitempty"`

	// resource type in
	ResourceTypeIn []string `json:"resource_type_in,omitempty"`

	// resource type lt
	ResourceTypeLt *string `json:"resource_type_lt,omitempty"`

	// resource type lte
	ResourceTypeLte *string `json:"resource_type_lte,omitempty"`

	// resource type not
	ResourceTypeNot *string `json:"resource_type_not,omitempty"`

	// resource type not contains
	ResourceTypeNotContains *string `json:"resource_type_not_contains,omitempty"`

	// resource type not ends with
	ResourceTypeNotEndsWith *string `json:"resource_type_not_ends_with,omitempty"`

	// resource type not in
	ResourceTypeNotIn []string `json:"resource_type_not_in,omitempty"`

	// resource type not starts with
	ResourceTypeNotStartsWith *string `json:"resource_type_not_starts_with,omitempty"`

	// resource type starts with
	ResourceTypeStartsWith *string `json:"resource_type_starts_with,omitempty"`

	// status
	Status *UserAuditLogStatus `json:"status,omitempty"`

	// status in
	StatusIn []UserAuditLogStatus `json:"status_in,omitempty"`

	// status not
	StatusNot *UserAuditLogStatus `json:"status_not,omitempty"`

	// status not in
	StatusNotIn []UserAuditLogStatus `json:"status_not_in,omitempty"`

	// user
	User *UserWhereInput `json:"user,omitempty"`
}

// Validate validates this user audit log where input
func (m *UserAuditLogWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAuditLogWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
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

func (m *UserAuditLogWhereInput) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
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

func (m *UserAuditLogWhereInput) validateStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) validateStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNot) { // not required
		return nil
	}

	if m.StatusNot != nil {
		if err := m.StatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserAuditLogWhereInput) validateStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this user audit log where input based on the context it is used
func (m *UserAuditLogWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserAuditLogWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserAuditLogWhereInput) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UserAuditLogWhereInput) contextValidateStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusIn); i++ {

		if err := m.StatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) contextValidateStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusNot != nil {
		if err := m.StatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not")
			}
			return err
		}
	}

	return nil
}

func (m *UserAuditLogWhereInput) contextValidateStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatusNotIn); i++ {

		if err := m.StatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *UserAuditLogWhereInput) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserAuditLogWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserAuditLogWhereInput) UnmarshalBinary(b []byte) error {
	var res UserAuditLogWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
