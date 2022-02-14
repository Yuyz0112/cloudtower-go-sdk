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

// IscsiLunWhereInput iscsi lun where input
// Example: {"AND":"IscsiLunWhereInput[]","NOT":"IscsiLunWhereInput[]","OR":"IscsiLunWhereInput[]","allowed_initiators":"string","allowed_initiators_contains":"string","allowed_initiators_ends_with":"string","allowed_initiators_gt":"string","allowed_initiators_gte":"string","allowed_initiators_in":["string"],"allowed_initiators_lt":"string","allowed_initiators_lte":"string","allowed_initiators_not":"string","allowed_initiators_not_contains":"string","allowed_initiators_not_ends_with":"string","allowed_initiators_not_in":["string"],"allowed_initiators_not_starts_with":"string","allowed_initiators_starts_with":"string","assigned_size":0,"assigned_size_gt":0,"assigned_size_gte":0,"assigned_size_in":[0],"assigned_size_lt":0,"assigned_size_lte":0,"assigned_size_not":0,"assigned_size_not_in":[0],"bps":0,"bps_gt":0,"bps_gte":0,"bps_in":[0],"bps_lt":0,"bps_lte":0,"bps_max":0,"bps_max_gt":0,"bps_max_gte":0,"bps_max_in":[0],"bps_max_length":0,"bps_max_length_gt":0,"bps_max_length_gte":0,"bps_max_length_in":[0],"bps_max_length_lt":0,"bps_max_length_lte":0,"bps_max_length_not":0,"bps_max_length_not_in":[0],"bps_max_lt":0,"bps_max_lte":0,"bps_max_not":0,"bps_max_not_in":[0],"bps_not":0,"bps_not_in":[0],"bps_rd":0,"bps_rd_gt":0,"bps_rd_gte":0,"bps_rd_in":[0],"bps_rd_lt":0,"bps_rd_lte":0,"bps_rd_max":0,"bps_rd_max_gt":0,"bps_rd_max_gte":0,"bps_rd_max_in":[0],"bps_rd_max_length":0,"bps_rd_max_length_gt":0,"bps_rd_max_length_gte":0,"bps_rd_max_length_in":[0],"bps_rd_max_length_lt":0,"bps_rd_max_length_lte":0,"bps_rd_max_length_not":0,"bps_rd_max_length_not_in":[0],"bps_rd_max_lt":0,"bps_rd_max_lte":0,"bps_rd_max_not":0,"bps_rd_max_not_in":[0],"bps_rd_not":0,"bps_rd_not_in":[0],"bps_wr":0,"bps_wr_gt":0,"bps_wr_gte":0,"bps_wr_in":[0],"bps_wr_lt":0,"bps_wr_lte":0,"bps_wr_max":0,"bps_wr_max_gt":0,"bps_wr_max_gte":0,"bps_wr_max_in":[0],"bps_wr_max_length":0,"bps_wr_max_length_gt":0,"bps_wr_max_length_gte":0,"bps_wr_max_length_in":[0],"bps_wr_max_length_lt":0,"bps_wr_max_length_lte":0,"bps_wr_max_length_not":0,"bps_wr_max_length_not_in":[0],"bps_wr_max_lt":0,"bps_wr_max_lte":0,"bps_wr_max_not":0,"bps_wr_max_not_in":[0],"bps_wr_not":0,"bps_wr_not_in":[0],"consistency_group":"ConsistencyGroupWhereInput","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","io_size":0,"io_size_gt":0,"io_size_gte":0,"io_size_in":[0],"io_size_lt":0,"io_size_lte":0,"io_size_not":0,"io_size_not_in":[0],"iops":0,"iops_gt":0,"iops_gte":0,"iops_in":[0],"iops_lt":0,"iops_lte":0,"iops_max":0,"iops_max_gt":0,"iops_max_gte":0,"iops_max_in":[0],"iops_max_length":0,"iops_max_length_gt":0,"iops_max_length_gte":0,"iops_max_length_in":[0],"iops_max_length_lt":0,"iops_max_length_lte":0,"iops_max_length_not":0,"iops_max_length_not_in":[0],"iops_max_lt":0,"iops_max_lte":0,"iops_max_not":0,"iops_max_not_in":[0],"iops_not":0,"iops_not_in":[0],"iops_rd":0,"iops_rd_gt":0,"iops_rd_gte":0,"iops_rd_in":[0],"iops_rd_lt":0,"iops_rd_lte":0,"iops_rd_max":0,"iops_rd_max_gt":0,"iops_rd_max_gte":0,"iops_rd_max_in":[0],"iops_rd_max_length":0,"iops_rd_max_length_gt":0,"iops_rd_max_length_gte":0,"iops_rd_max_length_in":[0],"iops_rd_max_length_lt":0,"iops_rd_max_length_lte":0,"iops_rd_max_length_not":0,"iops_rd_max_length_not_in":[0],"iops_rd_max_lt":0,"iops_rd_max_lte":0,"iops_rd_max_not":0,"iops_rd_max_not_in":[0],"iops_rd_not":0,"iops_rd_not_in":[0],"iops_wr":0,"iops_wr_gt":0,"iops_wr_gte":0,"iops_wr_in":[0],"iops_wr_lt":0,"iops_wr_lte":0,"iops_wr_max":0,"iops_wr_max_gt":0,"iops_wr_max_gte":0,"iops_wr_max_in":[0],"iops_wr_max_length":0,"iops_wr_max_length_gt":0,"iops_wr_max_length_gte":0,"iops_wr_max_length_in":[0],"iops_wr_max_length_lt":0,"iops_wr_max_length_lte":0,"iops_wr_max_length_not":0,"iops_wr_max_length_not_in":[0],"iops_wr_max_lt":0,"iops_wr_max_lte":0,"iops_wr_max_not":0,"iops_wr_max_not_in":[0],"iops_wr_not":0,"iops_wr_not_in":[0],"iscsi_target":"IscsiTargetWhereInput","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","lun_id":0,"lun_id_gt":0,"lun_id_gte":0,"lun_id_in":[0],"lun_id_lt":0,"lun_id_lte":0,"lun_id_not":0,"lun_id_not_in":[0],"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","replica_num":0,"replica_num_gt":0,"replica_num_gte":0,"replica_num_in":[0],"replica_num_lt":0,"replica_num_lte":0,"replica_num_not":0,"replica_num_not_in":[0],"shared_size":0,"shared_size_gt":0,"shared_size_gte":0,"shared_size_in":[0],"shared_size_lt":0,"shared_size_lte":0,"shared_size_not":0,"shared_size_not_in":[0],"snapshot_num":0,"snapshot_num_gt":0,"snapshot_num_gte":0,"snapshot_num_in":[0],"snapshot_num_lt":0,"snapshot_num_lte":0,"snapshot_num_not":0,"snapshot_num_not_in":[0],"stripe_num":0,"stripe_num_gt":0,"stripe_num_gte":0,"stripe_num_in":[0],"stripe_num_lt":0,"stripe_num_lte":0,"stripe_num_not":0,"stripe_num_not_in":[0],"stripe_size":0,"stripe_size_gt":0,"stripe_size_gte":0,"stripe_size_in":[0],"stripe_size_lt":0,"stripe_size_lte":0,"stripe_size_not":0,"stripe_size_not_in":[0],"thin_provision":false,"thin_provision_not":false,"unique_size":0,"unique_size_gt":0,"unique_size_gte":0,"unique_size_in":[0],"unique_size_lt":0,"unique_size_lte":0,"unique_size_not":0,"unique_size_not_in":[0],"zbs_volume_id":"string","zbs_volume_id_contains":"string","zbs_volume_id_ends_with":"string","zbs_volume_id_gt":"string","zbs_volume_id_gte":"string","zbs_volume_id_in":["string"],"zbs_volume_id_lt":"string","zbs_volume_id_lte":"string","zbs_volume_id_not":"string","zbs_volume_id_not_contains":"string","zbs_volume_id_not_ends_with":"string","zbs_volume_id_not_in":["string"],"zbs_volume_id_not_starts_with":"string","zbs_volume_id_starts_with":"string"}
//
// swagger:model IscsiLunWhereInput
type IscsiLunWhereInput struct {

	// a n d
	AND []*IscsiLunWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*IscsiLunWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*IscsiLunWhereInput `json:"OR,omitempty"`

	// allowed initiators
	AllowedInitiators *string `json:"allowed_initiators,omitempty"`

	// allowed initiators contains
	AllowedInitiatorsContains *string `json:"allowed_initiators_contains,omitempty"`

	// allowed initiators ends with
	AllowedInitiatorsEndsWith *string `json:"allowed_initiators_ends_with,omitempty"`

	// allowed initiators gt
	AllowedInitiatorsGt *string `json:"allowed_initiators_gt,omitempty"`

	// allowed initiators gte
	AllowedInitiatorsGte *string `json:"allowed_initiators_gte,omitempty"`

	// allowed initiators in
	AllowedInitiatorsIn []string `json:"allowed_initiators_in,omitempty"`

	// allowed initiators lt
	AllowedInitiatorsLt *string `json:"allowed_initiators_lt,omitempty"`

	// allowed initiators lte
	AllowedInitiatorsLte *string `json:"allowed_initiators_lte,omitempty"`

	// allowed initiators not
	AllowedInitiatorsNot *string `json:"allowed_initiators_not,omitempty"`

	// allowed initiators not contains
	AllowedInitiatorsNotContains *string `json:"allowed_initiators_not_contains,omitempty"`

	// allowed initiators not ends with
	AllowedInitiatorsNotEndsWith *string `json:"allowed_initiators_not_ends_with,omitempty"`

	// allowed initiators not in
	AllowedInitiatorsNotIn []string `json:"allowed_initiators_not_in,omitempty"`

	// allowed initiators not starts with
	AllowedInitiatorsNotStartsWith *string `json:"allowed_initiators_not_starts_with,omitempty"`

	// allowed initiators starts with
	AllowedInitiatorsStartsWith *string `json:"allowed_initiators_starts_with,omitempty"`

	// assigned size
	AssignedSize *float64 `json:"assigned_size,omitempty"`

	// assigned size gt
	AssignedSizeGt *float64 `json:"assigned_size_gt,omitempty"`

	// assigned size gte
	AssignedSizeGte *float64 `json:"assigned_size_gte,omitempty"`

	// assigned size in
	AssignedSizeIn []float64 `json:"assigned_size_in,omitempty"`

	// assigned size lt
	AssignedSizeLt *float64 `json:"assigned_size_lt,omitempty"`

	// assigned size lte
	AssignedSizeLte *float64 `json:"assigned_size_lte,omitempty"`

	// assigned size not
	AssignedSizeNot *float64 `json:"assigned_size_not,omitempty"`

	// assigned size not in
	AssignedSizeNotIn []float64 `json:"assigned_size_not_in,omitempty"`

	// bps
	Bps *float64 `json:"bps,omitempty"`

	// bps gt
	BpsGt *float64 `json:"bps_gt,omitempty"`

	// bps gte
	BpsGte *float64 `json:"bps_gte,omitempty"`

	// bps in
	BpsIn []float64 `json:"bps_in,omitempty"`

	// bps lt
	BpsLt *float64 `json:"bps_lt,omitempty"`

	// bps lte
	BpsLte *float64 `json:"bps_lte,omitempty"`

	// bps max
	BpsMax *float64 `json:"bps_max,omitempty"`

	// bps max gt
	BpsMaxGt *float64 `json:"bps_max_gt,omitempty"`

	// bps max gte
	BpsMaxGte *float64 `json:"bps_max_gte,omitempty"`

	// bps max in
	BpsMaxIn []float64 `json:"bps_max_in,omitempty"`

	// bps max length
	BpsMaxLength *float64 `json:"bps_max_length,omitempty"`

	// bps max length gt
	BpsMaxLengthGt *float64 `json:"bps_max_length_gt,omitempty"`

	// bps max length gte
	BpsMaxLengthGte *float64 `json:"bps_max_length_gte,omitempty"`

	// bps max length in
	BpsMaxLengthIn []float64 `json:"bps_max_length_in,omitempty"`

	// bps max length lt
	BpsMaxLengthLt *float64 `json:"bps_max_length_lt,omitempty"`

	// bps max length lte
	BpsMaxLengthLte *float64 `json:"bps_max_length_lte,omitempty"`

	// bps max length not
	BpsMaxLengthNot *float64 `json:"bps_max_length_not,omitempty"`

	// bps max length not in
	BpsMaxLengthNotIn []float64 `json:"bps_max_length_not_in,omitempty"`

	// bps max lt
	BpsMaxLt *float64 `json:"bps_max_lt,omitempty"`

	// bps max lte
	BpsMaxLte *float64 `json:"bps_max_lte,omitempty"`

	// bps max not
	BpsMaxNot *float64 `json:"bps_max_not,omitempty"`

	// bps max not in
	BpsMaxNotIn []float64 `json:"bps_max_not_in,omitempty"`

	// bps not
	BpsNot *float64 `json:"bps_not,omitempty"`

	// bps not in
	BpsNotIn []float64 `json:"bps_not_in,omitempty"`

	// bps rd
	BpsRd *float64 `json:"bps_rd,omitempty"`

	// bps rd gt
	BpsRdGt *float64 `json:"bps_rd_gt,omitempty"`

	// bps rd gte
	BpsRdGte *float64 `json:"bps_rd_gte,omitempty"`

	// bps rd in
	BpsRdIn []float64 `json:"bps_rd_in,omitempty"`

	// bps rd lt
	BpsRdLt *float64 `json:"bps_rd_lt,omitempty"`

	// bps rd lte
	BpsRdLte *float64 `json:"bps_rd_lte,omitempty"`

	// bps rd max
	BpsRdMax *float64 `json:"bps_rd_max,omitempty"`

	// bps rd max gt
	BpsRdMaxGt *float64 `json:"bps_rd_max_gt,omitempty"`

	// bps rd max gte
	BpsRdMaxGte *float64 `json:"bps_rd_max_gte,omitempty"`

	// bps rd max in
	BpsRdMaxIn []float64 `json:"bps_rd_max_in,omitempty"`

	// bps rd max length
	BpsRdMaxLength *float64 `json:"bps_rd_max_length,omitempty"`

	// bps rd max length gt
	BpsRdMaxLengthGt *float64 `json:"bps_rd_max_length_gt,omitempty"`

	// bps rd max length gte
	BpsRdMaxLengthGte *float64 `json:"bps_rd_max_length_gte,omitempty"`

	// bps rd max length in
	BpsRdMaxLengthIn []float64 `json:"bps_rd_max_length_in,omitempty"`

	// bps rd max length lt
	BpsRdMaxLengthLt *float64 `json:"bps_rd_max_length_lt,omitempty"`

	// bps rd max length lte
	BpsRdMaxLengthLte *float64 `json:"bps_rd_max_length_lte,omitempty"`

	// bps rd max length not
	BpsRdMaxLengthNot *float64 `json:"bps_rd_max_length_not,omitempty"`

	// bps rd max length not in
	BpsRdMaxLengthNotIn []float64 `json:"bps_rd_max_length_not_in,omitempty"`

	// bps rd max lt
	BpsRdMaxLt *float64 `json:"bps_rd_max_lt,omitempty"`

	// bps rd max lte
	BpsRdMaxLte *float64 `json:"bps_rd_max_lte,omitempty"`

	// bps rd max not
	BpsRdMaxNot *float64 `json:"bps_rd_max_not,omitempty"`

	// bps rd max not in
	BpsRdMaxNotIn []float64 `json:"bps_rd_max_not_in,omitempty"`

	// bps rd not
	BpsRdNot *float64 `json:"bps_rd_not,omitempty"`

	// bps rd not in
	BpsRdNotIn []float64 `json:"bps_rd_not_in,omitempty"`

	// bps wr
	BpsWr *float64 `json:"bps_wr,omitempty"`

	// bps wr gt
	BpsWrGt *float64 `json:"bps_wr_gt,omitempty"`

	// bps wr gte
	BpsWrGte *float64 `json:"bps_wr_gte,omitempty"`

	// bps wr in
	BpsWrIn []float64 `json:"bps_wr_in,omitempty"`

	// bps wr lt
	BpsWrLt *float64 `json:"bps_wr_lt,omitempty"`

	// bps wr lte
	BpsWrLte *float64 `json:"bps_wr_lte,omitempty"`

	// bps wr max
	BpsWrMax *float64 `json:"bps_wr_max,omitempty"`

	// bps wr max gt
	BpsWrMaxGt *float64 `json:"bps_wr_max_gt,omitempty"`

	// bps wr max gte
	BpsWrMaxGte *float64 `json:"bps_wr_max_gte,omitempty"`

	// bps wr max in
	BpsWrMaxIn []float64 `json:"bps_wr_max_in,omitempty"`

	// bps wr max length
	BpsWrMaxLength *float64 `json:"bps_wr_max_length,omitempty"`

	// bps wr max length gt
	BpsWrMaxLengthGt *float64 `json:"bps_wr_max_length_gt,omitempty"`

	// bps wr max length gte
	BpsWrMaxLengthGte *float64 `json:"bps_wr_max_length_gte,omitempty"`

	// bps wr max length in
	BpsWrMaxLengthIn []float64 `json:"bps_wr_max_length_in,omitempty"`

	// bps wr max length lt
	BpsWrMaxLengthLt *float64 `json:"bps_wr_max_length_lt,omitempty"`

	// bps wr max length lte
	BpsWrMaxLengthLte *float64 `json:"bps_wr_max_length_lte,omitempty"`

	// bps wr max length not
	BpsWrMaxLengthNot *float64 `json:"bps_wr_max_length_not,omitempty"`

	// bps wr max length not in
	BpsWrMaxLengthNotIn []float64 `json:"bps_wr_max_length_not_in,omitempty"`

	// bps wr max lt
	BpsWrMaxLt *float64 `json:"bps_wr_max_lt,omitempty"`

	// bps wr max lte
	BpsWrMaxLte *float64 `json:"bps_wr_max_lte,omitempty"`

	// bps wr max not
	BpsWrMaxNot *float64 `json:"bps_wr_max_not,omitempty"`

	// bps wr max not in
	BpsWrMaxNotIn []float64 `json:"bps_wr_max_not_in,omitempty"`

	// bps wr not
	BpsWrNot *float64 `json:"bps_wr_not,omitempty"`

	// bps wr not in
	BpsWrNotIn []float64 `json:"bps_wr_not_in,omitempty"`

	// consistency group
	ConsistencyGroup interface{} `json:"consistency_group,omitempty"`

	// entity async status
	EntityAsyncStatus interface{} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot interface{} `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

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

	// io size
	IoSize *float64 `json:"io_size,omitempty"`

	// io size gt
	IoSizeGt *float64 `json:"io_size_gt,omitempty"`

	// io size gte
	IoSizeGte *float64 `json:"io_size_gte,omitempty"`

	// io size in
	IoSizeIn []float64 `json:"io_size_in,omitempty"`

	// io size lt
	IoSizeLt *float64 `json:"io_size_lt,omitempty"`

	// io size lte
	IoSizeLte *float64 `json:"io_size_lte,omitempty"`

	// io size not
	IoSizeNot *float64 `json:"io_size_not,omitempty"`

	// io size not in
	IoSizeNotIn []float64 `json:"io_size_not_in,omitempty"`

	// iops
	Iops *float64 `json:"iops,omitempty"`

	// iops gt
	IopsGt *float64 `json:"iops_gt,omitempty"`

	// iops gte
	IopsGte *float64 `json:"iops_gte,omitempty"`

	// iops in
	IopsIn []float64 `json:"iops_in,omitempty"`

	// iops lt
	IopsLt *float64 `json:"iops_lt,omitempty"`

	// iops lte
	IopsLte *float64 `json:"iops_lte,omitempty"`

	// iops max
	IopsMax *float64 `json:"iops_max,omitempty"`

	// iops max gt
	IopsMaxGt *float64 `json:"iops_max_gt,omitempty"`

	// iops max gte
	IopsMaxGte *float64 `json:"iops_max_gte,omitempty"`

	// iops max in
	IopsMaxIn []float64 `json:"iops_max_in,omitempty"`

	// iops max length
	IopsMaxLength *float64 `json:"iops_max_length,omitempty"`

	// iops max length gt
	IopsMaxLengthGt *float64 `json:"iops_max_length_gt,omitempty"`

	// iops max length gte
	IopsMaxLengthGte *float64 `json:"iops_max_length_gte,omitempty"`

	// iops max length in
	IopsMaxLengthIn []float64 `json:"iops_max_length_in,omitempty"`

	// iops max length lt
	IopsMaxLengthLt *float64 `json:"iops_max_length_lt,omitempty"`

	// iops max length lte
	IopsMaxLengthLte *float64 `json:"iops_max_length_lte,omitempty"`

	// iops max length not
	IopsMaxLengthNot *float64 `json:"iops_max_length_not,omitempty"`

	// iops max length not in
	IopsMaxLengthNotIn []float64 `json:"iops_max_length_not_in,omitempty"`

	// iops max lt
	IopsMaxLt *float64 `json:"iops_max_lt,omitempty"`

	// iops max lte
	IopsMaxLte *float64 `json:"iops_max_lte,omitempty"`

	// iops max not
	IopsMaxNot *float64 `json:"iops_max_not,omitempty"`

	// iops max not in
	IopsMaxNotIn []float64 `json:"iops_max_not_in,omitempty"`

	// iops not
	IopsNot *float64 `json:"iops_not,omitempty"`

	// iops not in
	IopsNotIn []float64 `json:"iops_not_in,omitempty"`

	// iops rd
	IopsRd *float64 `json:"iops_rd,omitempty"`

	// iops rd gt
	IopsRdGt *float64 `json:"iops_rd_gt,omitempty"`

	// iops rd gte
	IopsRdGte *float64 `json:"iops_rd_gte,omitempty"`

	// iops rd in
	IopsRdIn []float64 `json:"iops_rd_in,omitempty"`

	// iops rd lt
	IopsRdLt *float64 `json:"iops_rd_lt,omitempty"`

	// iops rd lte
	IopsRdLte *float64 `json:"iops_rd_lte,omitempty"`

	// iops rd max
	IopsRdMax *float64 `json:"iops_rd_max,omitempty"`

	// iops rd max gt
	IopsRdMaxGt *float64 `json:"iops_rd_max_gt,omitempty"`

	// iops rd max gte
	IopsRdMaxGte *float64 `json:"iops_rd_max_gte,omitempty"`

	// iops rd max in
	IopsRdMaxIn []float64 `json:"iops_rd_max_in,omitempty"`

	// iops rd max length
	IopsRdMaxLength *float64 `json:"iops_rd_max_length,omitempty"`

	// iops rd max length gt
	IopsRdMaxLengthGt *float64 `json:"iops_rd_max_length_gt,omitempty"`

	// iops rd max length gte
	IopsRdMaxLengthGte *float64 `json:"iops_rd_max_length_gte,omitempty"`

	// iops rd max length in
	IopsRdMaxLengthIn []float64 `json:"iops_rd_max_length_in,omitempty"`

	// iops rd max length lt
	IopsRdMaxLengthLt *float64 `json:"iops_rd_max_length_lt,omitempty"`

	// iops rd max length lte
	IopsRdMaxLengthLte *float64 `json:"iops_rd_max_length_lte,omitempty"`

	// iops rd max length not
	IopsRdMaxLengthNot *float64 `json:"iops_rd_max_length_not,omitempty"`

	// iops rd max length not in
	IopsRdMaxLengthNotIn []float64 `json:"iops_rd_max_length_not_in,omitempty"`

	// iops rd max lt
	IopsRdMaxLt *float64 `json:"iops_rd_max_lt,omitempty"`

	// iops rd max lte
	IopsRdMaxLte *float64 `json:"iops_rd_max_lte,omitempty"`

	// iops rd max not
	IopsRdMaxNot *float64 `json:"iops_rd_max_not,omitempty"`

	// iops rd max not in
	IopsRdMaxNotIn []float64 `json:"iops_rd_max_not_in,omitempty"`

	// iops rd not
	IopsRdNot *float64 `json:"iops_rd_not,omitempty"`

	// iops rd not in
	IopsRdNotIn []float64 `json:"iops_rd_not_in,omitempty"`

	// iops wr
	IopsWr *float64 `json:"iops_wr,omitempty"`

	// iops wr gt
	IopsWrGt *float64 `json:"iops_wr_gt,omitempty"`

	// iops wr gte
	IopsWrGte *float64 `json:"iops_wr_gte,omitempty"`

	// iops wr in
	IopsWrIn []float64 `json:"iops_wr_in,omitempty"`

	// iops wr lt
	IopsWrLt *float64 `json:"iops_wr_lt,omitempty"`

	// iops wr lte
	IopsWrLte *float64 `json:"iops_wr_lte,omitempty"`

	// iops wr max
	IopsWrMax *float64 `json:"iops_wr_max,omitempty"`

	// iops wr max gt
	IopsWrMaxGt *float64 `json:"iops_wr_max_gt,omitempty"`

	// iops wr max gte
	IopsWrMaxGte *float64 `json:"iops_wr_max_gte,omitempty"`

	// iops wr max in
	IopsWrMaxIn []float64 `json:"iops_wr_max_in,omitempty"`

	// iops wr max length
	IopsWrMaxLength *float64 `json:"iops_wr_max_length,omitempty"`

	// iops wr max length gt
	IopsWrMaxLengthGt *float64 `json:"iops_wr_max_length_gt,omitempty"`

	// iops wr max length gte
	IopsWrMaxLengthGte *float64 `json:"iops_wr_max_length_gte,omitempty"`

	// iops wr max length in
	IopsWrMaxLengthIn []float64 `json:"iops_wr_max_length_in,omitempty"`

	// iops wr max length lt
	IopsWrMaxLengthLt *float64 `json:"iops_wr_max_length_lt,omitempty"`

	// iops wr max length lte
	IopsWrMaxLengthLte *float64 `json:"iops_wr_max_length_lte,omitempty"`

	// iops wr max length not
	IopsWrMaxLengthNot *float64 `json:"iops_wr_max_length_not,omitempty"`

	// iops wr max length not in
	IopsWrMaxLengthNotIn []float64 `json:"iops_wr_max_length_not_in,omitempty"`

	// iops wr max lt
	IopsWrMaxLt *float64 `json:"iops_wr_max_lt,omitempty"`

	// iops wr max lte
	IopsWrMaxLte *float64 `json:"iops_wr_max_lte,omitempty"`

	// iops wr max not
	IopsWrMaxNot *float64 `json:"iops_wr_max_not,omitempty"`

	// iops wr max not in
	IopsWrMaxNotIn []float64 `json:"iops_wr_max_not_in,omitempty"`

	// iops wr not
	IopsWrNot *float64 `json:"iops_wr_not,omitempty"`

	// iops wr not in
	IopsWrNotIn []float64 `json:"iops_wr_not_in,omitempty"`

	// iscsi target
	IscsiTarget interface{} `json:"iscsi_target,omitempty"`

	// labels every
	LabelsEvery interface{} `json:"labels_every,omitempty"`

	// labels none
	LabelsNone interface{} `json:"labels_none,omitempty"`

	// labels some
	LabelsSome interface{} `json:"labels_some,omitempty"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local created at gt
	LocalCreatedAtGt *string `json:"local_created_at_gt,omitempty"`

	// local created at gte
	LocalCreatedAtGte *string `json:"local_created_at_gte,omitempty"`

	// local created at in
	LocalCreatedAtIn []string `json:"local_created_at_in,omitempty"`

	// local created at lt
	LocalCreatedAtLt *string `json:"local_created_at_lt,omitempty"`

	// local created at lte
	LocalCreatedAtLte *string `json:"local_created_at_lte,omitempty"`

	// local created at not
	LocalCreatedAtNot *string `json:"local_created_at_not,omitempty"`

	// local created at not in
	LocalCreatedAtNotIn []string `json:"local_created_at_not_in,omitempty"`

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

	// lun id
	LunID *int32 `json:"lun_id,omitempty"`

	// lun id gt
	LunIDGt *int32 `json:"lun_id_gt,omitempty"`

	// lun id gte
	LunIDGte *int32 `json:"lun_id_gte,omitempty"`

	// lun id in
	LunIDIn []int32 `json:"lun_id_in,omitempty"`

	// lun id lt
	LunIDLt *int32 `json:"lun_id_lt,omitempty"`

	// lun id lte
	LunIDLte *int32 `json:"lun_id_lte,omitempty"`

	// lun id not
	LunIDNot *int32 `json:"lun_id_not,omitempty"`

	// lun id not in
	LunIDNotIn []int32 `json:"lun_id_not_in,omitempty"`

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

	// replica num
	ReplicaNum *int32 `json:"replica_num,omitempty"`

	// replica num gt
	ReplicaNumGt *int32 `json:"replica_num_gt,omitempty"`

	// replica num gte
	ReplicaNumGte *int32 `json:"replica_num_gte,omitempty"`

	// replica num in
	ReplicaNumIn []int32 `json:"replica_num_in,omitempty"`

	// replica num lt
	ReplicaNumLt *int32 `json:"replica_num_lt,omitempty"`

	// replica num lte
	ReplicaNumLte *int32 `json:"replica_num_lte,omitempty"`

	// replica num not
	ReplicaNumNot *int32 `json:"replica_num_not,omitempty"`

	// replica num not in
	ReplicaNumNotIn []int32 `json:"replica_num_not_in,omitempty"`

	// shared size
	SharedSize *float64 `json:"shared_size,omitempty"`

	// shared size gt
	SharedSizeGt *float64 `json:"shared_size_gt,omitempty"`

	// shared size gte
	SharedSizeGte *float64 `json:"shared_size_gte,omitempty"`

	// shared size in
	SharedSizeIn []float64 `json:"shared_size_in,omitempty"`

	// shared size lt
	SharedSizeLt *float64 `json:"shared_size_lt,omitempty"`

	// shared size lte
	SharedSizeLte *float64 `json:"shared_size_lte,omitempty"`

	// shared size not
	SharedSizeNot *float64 `json:"shared_size_not,omitempty"`

	// shared size not in
	SharedSizeNotIn []float64 `json:"shared_size_not_in,omitempty"`

	// snapshot num
	SnapshotNum *int32 `json:"snapshot_num,omitempty"`

	// snapshot num gt
	SnapshotNumGt *int32 `json:"snapshot_num_gt,omitempty"`

	// snapshot num gte
	SnapshotNumGte *int32 `json:"snapshot_num_gte,omitempty"`

	// snapshot num in
	SnapshotNumIn []int32 `json:"snapshot_num_in,omitempty"`

	// snapshot num lt
	SnapshotNumLt *int32 `json:"snapshot_num_lt,omitempty"`

	// snapshot num lte
	SnapshotNumLte *int32 `json:"snapshot_num_lte,omitempty"`

	// snapshot num not
	SnapshotNumNot *int32 `json:"snapshot_num_not,omitempty"`

	// snapshot num not in
	SnapshotNumNotIn []int32 `json:"snapshot_num_not_in,omitempty"`

	// stripe num
	StripeNum *int32 `json:"stripe_num,omitempty"`

	// stripe num gt
	StripeNumGt *int32 `json:"stripe_num_gt,omitempty"`

	// stripe num gte
	StripeNumGte *int32 `json:"stripe_num_gte,omitempty"`

	// stripe num in
	StripeNumIn []int32 `json:"stripe_num_in,omitempty"`

	// stripe num lt
	StripeNumLt *int32 `json:"stripe_num_lt,omitempty"`

	// stripe num lte
	StripeNumLte *int32 `json:"stripe_num_lte,omitempty"`

	// stripe num not
	StripeNumNot *int32 `json:"stripe_num_not,omitempty"`

	// stripe num not in
	StripeNumNotIn []int32 `json:"stripe_num_not_in,omitempty"`

	// stripe size
	StripeSize *float64 `json:"stripe_size,omitempty"`

	// stripe size gt
	StripeSizeGt *float64 `json:"stripe_size_gt,omitempty"`

	// stripe size gte
	StripeSizeGte *float64 `json:"stripe_size_gte,omitempty"`

	// stripe size in
	StripeSizeIn []float64 `json:"stripe_size_in,omitempty"`

	// stripe size lt
	StripeSizeLt *float64 `json:"stripe_size_lt,omitempty"`

	// stripe size lte
	StripeSizeLte *float64 `json:"stripe_size_lte,omitempty"`

	// stripe size not
	StripeSizeNot *float64 `json:"stripe_size_not,omitempty"`

	// stripe size not in
	StripeSizeNotIn []float64 `json:"stripe_size_not_in,omitempty"`

	// thin provision
	ThinProvision *bool `json:"thin_provision,omitempty"`

	// thin provision not
	ThinProvisionNot *bool `json:"thin_provision_not,omitempty"`

	// unique size
	UniqueSize *float64 `json:"unique_size,omitempty"`

	// unique size gt
	UniqueSizeGt *float64 `json:"unique_size_gt,omitempty"`

	// unique size gte
	UniqueSizeGte *float64 `json:"unique_size_gte,omitempty"`

	// unique size in
	UniqueSizeIn []float64 `json:"unique_size_in,omitempty"`

	// unique size lt
	UniqueSizeLt *float64 `json:"unique_size_lt,omitempty"`

	// unique size lte
	UniqueSizeLte *float64 `json:"unique_size_lte,omitempty"`

	// unique size not
	UniqueSizeNot *float64 `json:"unique_size_not,omitempty"`

	// unique size not in
	UniqueSizeNotIn []float64 `json:"unique_size_not_in,omitempty"`

	// zbs volume id
	ZbsVolumeID *string `json:"zbs_volume_id,omitempty"`

	// zbs volume id contains
	ZbsVolumeIDContains *string `json:"zbs_volume_id_contains,omitempty"`

	// zbs volume id ends with
	ZbsVolumeIDEndsWith *string `json:"zbs_volume_id_ends_with,omitempty"`

	// zbs volume id gt
	ZbsVolumeIDGt *string `json:"zbs_volume_id_gt,omitempty"`

	// zbs volume id gte
	ZbsVolumeIDGte *string `json:"zbs_volume_id_gte,omitempty"`

	// zbs volume id in
	ZbsVolumeIDIn []string `json:"zbs_volume_id_in,omitempty"`

	// zbs volume id lt
	ZbsVolumeIDLt *string `json:"zbs_volume_id_lt,omitempty"`

	// zbs volume id lte
	ZbsVolumeIDLte *string `json:"zbs_volume_id_lte,omitempty"`

	// zbs volume id not
	ZbsVolumeIDNot *string `json:"zbs_volume_id_not,omitempty"`

	// zbs volume id not contains
	ZbsVolumeIDNotContains *string `json:"zbs_volume_id_not_contains,omitempty"`

	// zbs volume id not ends with
	ZbsVolumeIDNotEndsWith *string `json:"zbs_volume_id_not_ends_with,omitempty"`

	// zbs volume id not in
	ZbsVolumeIDNotIn []string `json:"zbs_volume_id_not_in,omitempty"`

	// zbs volume id not starts with
	ZbsVolumeIDNotStartsWith *string `json:"zbs_volume_id_not_starts_with,omitempty"`

	// zbs volume id starts with
	ZbsVolumeIDStartsWith *string `json:"zbs_volume_id_starts_with,omitempty"`
}

// Validate validates this iscsi lun where input
func (m *IscsiLunWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLunWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *IscsiLunWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *IscsiLunWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *IscsiLunWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
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

func (m *IscsiLunWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
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

// ContextValidate validate this iscsi lun where input based on the context it is used
func (m *IscsiLunWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLunWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiLunWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiLunWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiLunWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IscsiLunWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IscsiLunWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiLunWhereInput) UnmarshalBinary(b []byte) error {
	var res IscsiLunWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
