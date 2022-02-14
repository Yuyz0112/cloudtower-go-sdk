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

// SnmpTrapReceiverWhereInput snmp trap receiver where input
// Example: {"AND":"SnmpTrapReceiverWhereInput[]","NOT":"SnmpTrapReceiverWhereInput[]","OR":"SnmpTrapReceiverWhereInput[]","auth_pass_phrase":"string","auth_pass_phrase_contains":"string","auth_pass_phrase_ends_with":"string","auth_pass_phrase_gt":"string","auth_pass_phrase_gte":"string","auth_pass_phrase_in":["string"],"auth_pass_phrase_lt":"string","auth_pass_phrase_lte":"string","auth_pass_phrase_not":"string","auth_pass_phrase_not_contains":"string","auth_pass_phrase_not_ends_with":"string","auth_pass_phrase_not_in":["string"],"auth_pass_phrase_not_starts_with":"string","auth_pass_phrase_starts_with":"string","auth_protocol":"MD5","auth_protocol_in":["MD5"],"auth_protocol_not":"MD5","auth_protocol_not_in":["MD5"],"cluster":"ClusterWhereInput","community":"string","community_contains":"string","community_ends_with":"string","community_gt":"string","community_gte":"string","community_in":["string"],"community_lt":"string","community_lte":"string","community_not":"string","community_not_contains":"string","community_not_ends_with":"string","community_not_in":["string"],"community_not_starts_with":"string","community_starts_with":"string","disabled":false,"disabled_not":false,"engine_id":"string","engine_id_contains":"string","engine_id_ends_with":"string","engine_id_gt":"string","engine_id_gte":"string","engine_id_in":["string"],"engine_id_lt":"string","engine_id_lte":"string","engine_id_not":"string","engine_id_not_contains":"string","engine_id_not_ends_with":"string","engine_id_not_in":["string"],"engine_id_not_starts_with":"string","engine_id_starts_with":"string","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"host":"string","host_contains":"string","host_ends_with":"string","host_gt":"string","host_gte":"string","host_in":["string"],"host_lt":"string","host_lte":"string","host_not":"string","host_not_contains":"string","host_not_ends_with":"string","host_not_in":["string"],"host_not_starts_with":"string","host_starts_with":"string","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","inform":false,"inform_not":false,"language_code":"EN_US","language_code_in":["EN_US"],"language_code_not":"EN_US","language_code_not_in":["EN_US"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","port":0,"port_gt":0,"port_gte":0,"port_in":[0],"port_lt":0,"port_lte":0,"port_not":0,"port_not_in":[0],"privacy_pass_phrase":"string","privacy_pass_phrase_contains":"string","privacy_pass_phrase_ends_with":"string","privacy_pass_phrase_gt":"string","privacy_pass_phrase_gte":"string","privacy_pass_phrase_in":["string"],"privacy_pass_phrase_lt":"string","privacy_pass_phrase_lte":"string","privacy_pass_phrase_not":"string","privacy_pass_phrase_not_contains":"string","privacy_pass_phrase_not_ends_with":"string","privacy_pass_phrase_not_in":["string"],"privacy_pass_phrase_not_starts_with":"string","privacy_pass_phrase_starts_with":"string","privacy_protocol":"AES","privacy_protocol_in":["AES"],"privacy_protocol_not":"AES","privacy_protocol_not_in":["AES"],"protocol":"TCP","protocol_in":["TCP"],"protocol_not":"TCP","protocol_not_in":["TCP"],"username":"string","username_contains":"string","username_ends_with":"string","username_gt":"string","username_gte":"string","username_in":["string"],"username_lt":"string","username_lte":"string","username_not":"string","username_not_contains":"string","username_not_ends_with":"string","username_not_in":["string"],"username_not_starts_with":"string","username_starts_with":"string","version":"V2C","version_in":["V2C"],"version_not":"V2C","version_not_in":["V2C"]}
//
// swagger:model SnmpTrapReceiverWhereInput
type SnmpTrapReceiverWhereInput struct {

	// a n d
	AND []*SnmpTrapReceiverWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*SnmpTrapReceiverWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*SnmpTrapReceiverWhereInput `json:"OR,omitempty"`

	// auth pass phrase
	AuthPassPhrase *string `json:"auth_pass_phrase,omitempty"`

	// auth pass phrase contains
	AuthPassPhraseContains *string `json:"auth_pass_phrase_contains,omitempty"`

	// auth pass phrase ends with
	AuthPassPhraseEndsWith *string `json:"auth_pass_phrase_ends_with,omitempty"`

	// auth pass phrase gt
	AuthPassPhraseGt *string `json:"auth_pass_phrase_gt,omitempty"`

	// auth pass phrase gte
	AuthPassPhraseGte *string `json:"auth_pass_phrase_gte,omitempty"`

	// auth pass phrase in
	AuthPassPhraseIn []string `json:"auth_pass_phrase_in,omitempty"`

	// auth pass phrase lt
	AuthPassPhraseLt *string `json:"auth_pass_phrase_lt,omitempty"`

	// auth pass phrase lte
	AuthPassPhraseLte *string `json:"auth_pass_phrase_lte,omitempty"`

	// auth pass phrase not
	AuthPassPhraseNot *string `json:"auth_pass_phrase_not,omitempty"`

	// auth pass phrase not contains
	AuthPassPhraseNotContains *string `json:"auth_pass_phrase_not_contains,omitempty"`

	// auth pass phrase not ends with
	AuthPassPhraseNotEndsWith *string `json:"auth_pass_phrase_not_ends_with,omitempty"`

	// auth pass phrase not in
	AuthPassPhraseNotIn []string `json:"auth_pass_phrase_not_in,omitempty"`

	// auth pass phrase not starts with
	AuthPassPhraseNotStartsWith *string `json:"auth_pass_phrase_not_starts_with,omitempty"`

	// auth pass phrase starts with
	AuthPassPhraseStartsWith *string `json:"auth_pass_phrase_starts_with,omitempty"`

	// auth protocol
	AuthProtocol interface{} `json:"auth_protocol,omitempty"`

	// auth protocol in
	AuthProtocolIn []SnmpAuthProtocol `json:"auth_protocol_in,omitempty"`

	// auth protocol not
	AuthProtocolNot interface{} `json:"auth_protocol_not,omitempty"`

	// auth protocol not in
	AuthProtocolNotIn []SnmpAuthProtocol `json:"auth_protocol_not_in,omitempty"`

	// cluster
	Cluster interface{} `json:"cluster,omitempty"`

	// community
	Community *string `json:"community,omitempty"`

	// community contains
	CommunityContains *string `json:"community_contains,omitempty"`

	// community ends with
	CommunityEndsWith *string `json:"community_ends_with,omitempty"`

	// community gt
	CommunityGt *string `json:"community_gt,omitempty"`

	// community gte
	CommunityGte *string `json:"community_gte,omitempty"`

	// community in
	CommunityIn []string `json:"community_in,omitempty"`

	// community lt
	CommunityLt *string `json:"community_lt,omitempty"`

	// community lte
	CommunityLte *string `json:"community_lte,omitempty"`

	// community not
	CommunityNot *string `json:"community_not,omitempty"`

	// community not contains
	CommunityNotContains *string `json:"community_not_contains,omitempty"`

	// community not ends with
	CommunityNotEndsWith *string `json:"community_not_ends_with,omitempty"`

	// community not in
	CommunityNotIn []string `json:"community_not_in,omitempty"`

	// community not starts with
	CommunityNotStartsWith *string `json:"community_not_starts_with,omitempty"`

	// community starts with
	CommunityStartsWith *string `json:"community_starts_with,omitempty"`

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// disabled not
	DisabledNot *bool `json:"disabled_not,omitempty"`

	// engine id
	EngineID *string `json:"engine_id,omitempty"`

	// engine id contains
	EngineIDContains *string `json:"engine_id_contains,omitempty"`

	// engine id ends with
	EngineIDEndsWith *string `json:"engine_id_ends_with,omitempty"`

	// engine id gt
	EngineIDGt *string `json:"engine_id_gt,omitempty"`

	// engine id gte
	EngineIDGte *string `json:"engine_id_gte,omitempty"`

	// engine id in
	EngineIDIn []string `json:"engine_id_in,omitempty"`

	// engine id lt
	EngineIDLt *string `json:"engine_id_lt,omitempty"`

	// engine id lte
	EngineIDLte *string `json:"engine_id_lte,omitempty"`

	// engine id not
	EngineIDNot *string `json:"engine_id_not,omitempty"`

	// engine id not contains
	EngineIDNotContains *string `json:"engine_id_not_contains,omitempty"`

	// engine id not ends with
	EngineIDNotEndsWith *string `json:"engine_id_not_ends_with,omitempty"`

	// engine id not in
	EngineIDNotIn []string `json:"engine_id_not_in,omitempty"`

	// engine id not starts with
	EngineIDNotStartsWith *string `json:"engine_id_not_starts_with,omitempty"`

	// engine id starts with
	EngineIDStartsWith *string `json:"engine_id_starts_with,omitempty"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot interface{} `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// host
	Host *string `json:"host,omitempty"`

	// host contains
	HostContains *string `json:"host_contains,omitempty"`

	// host ends with
	HostEndsWith *string `json:"host_ends_with,omitempty"`

	// host gt
	HostGt *string `json:"host_gt,omitempty"`

	// host gte
	HostGte *string `json:"host_gte,omitempty"`

	// host in
	HostIn []string `json:"host_in,omitempty"`

	// host lt
	HostLt *string `json:"host_lt,omitempty"`

	// host lte
	HostLte *string `json:"host_lte,omitempty"`

	// host not
	HostNot *string `json:"host_not,omitempty"`

	// host not contains
	HostNotContains *string `json:"host_not_contains,omitempty"`

	// host not ends with
	HostNotEndsWith *string `json:"host_not_ends_with,omitempty"`

	// host not in
	HostNotIn []string `json:"host_not_in,omitempty"`

	// host not starts with
	HostNotStartsWith *string `json:"host_not_starts_with,omitempty"`

	// host starts with
	HostStartsWith *string `json:"host_starts_with,omitempty"`

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

	// inform
	Inform *bool `json:"inform,omitempty"`

	// inform not
	InformNot *bool `json:"inform_not,omitempty"`

	// language code
	LanguageCode interface{} `json:"language_code,omitempty"`

	// language code in
	LanguageCodeIn []SnmpLanguageCode `json:"language_code_in,omitempty"`

	// language code not
	LanguageCodeNot interface{} `json:"language_code_not,omitempty"`

	// language code not in
	LanguageCodeNotIn []SnmpLanguageCode `json:"language_code_not_in,omitempty"`

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// port
	Port *int32 `json:"port,omitempty"`

	// port gt
	PortGt *int32 `json:"port_gt,omitempty"`

	// port gte
	PortGte *int32 `json:"port_gte,omitempty"`

	// port in
	PortIn []int32 `json:"port_in,omitempty"`

	// port lt
	PortLt *int32 `json:"port_lt,omitempty"`

	// port lte
	PortLte *int32 `json:"port_lte,omitempty"`

	// port not
	PortNot *int32 `json:"port_not,omitempty"`

	// port not in
	PortNotIn []int32 `json:"port_not_in,omitempty"`

	// privacy pass phrase
	PrivacyPassPhrase *string `json:"privacy_pass_phrase,omitempty"`

	// privacy pass phrase contains
	PrivacyPassPhraseContains *string `json:"privacy_pass_phrase_contains,omitempty"`

	// privacy pass phrase ends with
	PrivacyPassPhraseEndsWith *string `json:"privacy_pass_phrase_ends_with,omitempty"`

	// privacy pass phrase gt
	PrivacyPassPhraseGt *string `json:"privacy_pass_phrase_gt,omitempty"`

	// privacy pass phrase gte
	PrivacyPassPhraseGte *string `json:"privacy_pass_phrase_gte,omitempty"`

	// privacy pass phrase in
	PrivacyPassPhraseIn []string `json:"privacy_pass_phrase_in,omitempty"`

	// privacy pass phrase lt
	PrivacyPassPhraseLt *string `json:"privacy_pass_phrase_lt,omitempty"`

	// privacy pass phrase lte
	PrivacyPassPhraseLte *string `json:"privacy_pass_phrase_lte,omitempty"`

	// privacy pass phrase not
	PrivacyPassPhraseNot *string `json:"privacy_pass_phrase_not,omitempty"`

	// privacy pass phrase not contains
	PrivacyPassPhraseNotContains *string `json:"privacy_pass_phrase_not_contains,omitempty"`

	// privacy pass phrase not ends with
	PrivacyPassPhraseNotEndsWith *string `json:"privacy_pass_phrase_not_ends_with,omitempty"`

	// privacy pass phrase not in
	PrivacyPassPhraseNotIn []string `json:"privacy_pass_phrase_not_in,omitempty"`

	// privacy pass phrase not starts with
	PrivacyPassPhraseNotStartsWith *string `json:"privacy_pass_phrase_not_starts_with,omitempty"`

	// privacy pass phrase starts with
	PrivacyPassPhraseStartsWith *string `json:"privacy_pass_phrase_starts_with,omitempty"`

	// privacy protocol
	PrivacyProtocol interface{} `json:"privacy_protocol,omitempty"`

	// privacy protocol in
	PrivacyProtocolIn []SnmpPrivacyProtocol `json:"privacy_protocol_in,omitempty"`

	// privacy protocol not
	PrivacyProtocolNot interface{} `json:"privacy_protocol_not,omitempty"`

	// privacy protocol not in
	PrivacyProtocolNotIn []SnmpPrivacyProtocol `json:"privacy_protocol_not_in,omitempty"`

	// protocol
	Protocol interface{} `json:"protocol,omitempty"`

	// protocol in
	ProtocolIn []SnmpProtocol `json:"protocol_in,omitempty"`

	// protocol not
	ProtocolNot interface{} `json:"protocol_not,omitempty"`

	// protocol not in
	ProtocolNotIn []SnmpProtocol `json:"protocol_not_in,omitempty"`

	// username
	Username *string `json:"username,omitempty"`

	// username contains
	UsernameContains *string `json:"username_contains,omitempty"`

	// username ends with
	UsernameEndsWith *string `json:"username_ends_with,omitempty"`

	// username gt
	UsernameGt *string `json:"username_gt,omitempty"`

	// username gte
	UsernameGte *string `json:"username_gte,omitempty"`

	// username in
	UsernameIn []string `json:"username_in,omitempty"`

	// username lt
	UsernameLt *string `json:"username_lt,omitempty"`

	// username lte
	UsernameLte *string `json:"username_lte,omitempty"`

	// username not
	UsernameNot *string `json:"username_not,omitempty"`

	// username not contains
	UsernameNotContains *string `json:"username_not_contains,omitempty"`

	// username not ends with
	UsernameNotEndsWith *string `json:"username_not_ends_with,omitempty"`

	// username not in
	UsernameNotIn []string `json:"username_not_in,omitempty"`

	// username not starts with
	UsernameNotStartsWith *string `json:"username_not_starts_with,omitempty"`

	// username starts with
	UsernameStartsWith *string `json:"username_starts_with,omitempty"`

	// version
	Version interface{} `json:"version,omitempty"`

	// version in
	VersionIn []SnmpVersion `json:"version_in,omitempty"`

	// version not
	VersionNot interface{} `json:"version_not,omitempty"`

	// version not in
	VersionNotIn []SnmpVersion `json:"version_not_in,omitempty"`
}

// Validate validates this snmp trap receiver where input
func (m *SnmpTrapReceiverWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateAuthProtocolIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthProtocolNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCodeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguageCodeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacyProtocolIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacyProtocolNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateAND(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateNOT(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateOR(formats strfmt.Registry) error {
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
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateAuthProtocolIn(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthProtocolIn) { // not required
		return nil
	}

	for i := 0; i < len(m.AuthProtocolIn); i++ {

		if err := m.AuthProtocolIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_protocol_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateAuthProtocolNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthProtocolNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.AuthProtocolNotIn); i++ {

		if err := m.AuthProtocolNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_protocol_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateLanguageCodeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCodeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.LanguageCodeIn); i++ {

		if err := m.LanguageCodeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateLanguageCodeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCodeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.LanguageCodeNotIn); i++ {

		if err := m.LanguageCodeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validatePrivacyProtocolIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivacyProtocolIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PrivacyProtocolIn); i++ {

		if err := m.PrivacyProtocolIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy_protocol_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validatePrivacyProtocolNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.PrivacyProtocolNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.PrivacyProtocolNotIn); i++ {

		if err := m.PrivacyProtocolNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy_protocol_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateProtocolIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtocolIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtocolIn); i++ {

		if err := m.ProtocolIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateProtocolNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtocolNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtocolNotIn); i++ {

		if err := m.ProtocolNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateVersionIn(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionIn) { // not required
		return nil
	}

	for i := 0; i < len(m.VersionIn); i++ {

		if err := m.VersionIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) validateVersionNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.VersionNotIn); i++ {

		if err := m.VersionNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this snmp trap receiver where input based on the context it is used
func (m *SnmpTrapReceiverWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateAuthProtocolIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthProtocolNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCodeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguageCodeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivacyProtocolIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivacyProtocolNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocolIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocolNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersionIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersionNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateAuthProtocolIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuthProtocolIn); i++ {

		if err := m.AuthProtocolIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_protocol_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateAuthProtocolNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuthProtocolNotIn); i++ {

		if err := m.AuthProtocolNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_protocol_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateLanguageCodeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LanguageCodeIn); i++ {

		if err := m.LanguageCodeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateLanguageCodeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LanguageCodeNotIn); i++ {

		if err := m.LanguageCodeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidatePrivacyProtocolIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrivacyProtocolIn); i++ {

		if err := m.PrivacyProtocolIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy_protocol_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidatePrivacyProtocolNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PrivacyProtocolNotIn); i++ {

		if err := m.PrivacyProtocolNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy_protocol_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateProtocolIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProtocolIn); i++ {

		if err := m.ProtocolIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateProtocolNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProtocolNotIn); i++ {

		if err := m.ProtocolNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protocol_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateVersionIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VersionIn); i++ {

		if err := m.VersionIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SnmpTrapReceiverWhereInput) contextValidateVersionNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VersionNotIn); i++ {

		if err := m.VersionNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnmpTrapReceiverWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnmpTrapReceiverWhereInput) UnmarshalBinary(b []byte) error {
	var res SnmpTrapReceiverWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
