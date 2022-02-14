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

// ContentLibraryImageWhereInput content library image where input
// Example: {"AND":"ContentLibraryImageWhereInput[]","NOT":"ContentLibraryImageWhereInput[]","OR":"ContentLibraryImageWhereInput[]","clusters_every":"ClusterWhereInput","clusters_none":"ClusterWhereInput","clusters_some":"ClusterWhereInput","createdAt":"string","createdAt_gt":"string","createdAt_gte":"string","createdAt_in":["string"],"createdAt_lt":"string","createdAt_lte":"string","createdAt_not":"string","createdAt_not_in":["string"],"description":"string","description_contains":"string","description_ends_with":"string","description_gt":"string","description_gte":"string","description_in":["string"],"description_lt":"string","description_lte":"string","description_not":"string","description_not_contains":"string","description_not_ends_with":"string","description_not_in":["string"],"description_not_starts_with":"string","description_starts_with":"string","elf_images_every":"ElfImageWhereInput","elf_images_none":"ElfImageWhereInput","elf_images_some":"ElfImageWhereInput","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","path":"string","path_contains":"string","path_ends_with":"string","path_gt":"string","path_gte":"string","path_in":["string"],"path_lt":"string","path_lte":"string","path_not":"string","path_not_contains":"string","path_not_ends_with":"string","path_not_in":["string"],"path_not_starts_with":"string","path_starts_with":"string","size":0,"size_gt":0,"size_gte":0,"size_in":[0],"size_lt":0,"size_lte":0,"size_not":0,"size_not_in":[0],"vm_disks_every":"VmDiskWhereInput","vm_disks_none":"VmDiskWhereInput","vm_disks_some":"VmDiskWhereInput","vm_snapshots_every":"VmSnapshotWhereInput","vm_snapshots_none":"VmSnapshotWhereInput","vm_snapshots_some":"VmSnapshotWhereInput","vm_templates_every":"VmTemplateWhereInput","vm_templates_none":"VmTemplateWhereInput","vm_templates_some":"VmTemplateWhereInput"}
//
// swagger:model ContentLibraryImageWhereInput
type ContentLibraryImageWhereInput struct {

	// a n d
	AND []*ContentLibraryImageWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*ContentLibraryImageWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*ContentLibraryImageWhereInput `json:"OR,omitempty"`

	// clusters every
	ClustersEvery struct {
		ClusterWhereInput
	} `json:"clusters_every,omitempty"`

	// clusters none
	ClustersNone struct {
		ClusterWhereInput
	} `json:"clusters_none,omitempty"`

	// clusters some
	ClustersSome struct {
		ClusterWhereInput
	} `json:"clusters_some,omitempty"`

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

	// description
	Description *string `json:"description,omitempty"`

	// description contains
	DescriptionContains *string `json:"description_contains,omitempty"`

	// description ends with
	DescriptionEndsWith *string `json:"description_ends_with,omitempty"`

	// description gt
	DescriptionGt *string `json:"description_gt,omitempty"`

	// description gte
	DescriptionGte *string `json:"description_gte,omitempty"`

	// description in
	DescriptionIn []string `json:"description_in,omitempty"`

	// description lt
	DescriptionLt *string `json:"description_lt,omitempty"`

	// description lte
	DescriptionLte *string `json:"description_lte,omitempty"`

	// description not
	DescriptionNot *string `json:"description_not,omitempty"`

	// description not contains
	DescriptionNotContains *string `json:"description_not_contains,omitempty"`

	// description not ends with
	DescriptionNotEndsWith *string `json:"description_not_ends_with,omitempty"`

	// description not in
	DescriptionNotIn []string `json:"description_not_in,omitempty"`

	// description not starts with
	DescriptionNotStartsWith *string `json:"description_not_starts_with,omitempty"`

	// description starts with
	DescriptionStartsWith *string `json:"description_starts_with,omitempty"`

	// elf images every
	ElfImagesEvery struct {
		ElfImageWhereInput
	} `json:"elf_images_every,omitempty"`

	// elf images none
	ElfImagesNone struct {
		ElfImageWhereInput
	} `json:"elf_images_none,omitempty"`

	// elf images some
	ElfImagesSome struct {
		ElfImageWhereInput
	} `json:"elf_images_some,omitempty"`

	// entity async status
	EntityAsyncStatus struct {
		EntityAsyncStatus
	} `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot struct {
		EntityAsyncStatus
	} `json:"entityAsyncStatus_not,omitempty"`

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

	// labels every
	LabelsEvery struct {
		LabelWhereInput
	} `json:"labels_every,omitempty"`

	// labels none
	LabelsNone struct {
		LabelWhereInput
	} `json:"labels_none,omitempty"`

	// labels some
	LabelsSome struct {
		LabelWhereInput
	} `json:"labels_some,omitempty"`

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

	// path
	Path *string `json:"path,omitempty"`

	// path contains
	PathContains *string `json:"path_contains,omitempty"`

	// path ends with
	PathEndsWith *string `json:"path_ends_with,omitempty"`

	// path gt
	PathGt *string `json:"path_gt,omitempty"`

	// path gte
	PathGte *string `json:"path_gte,omitempty"`

	// path in
	PathIn []string `json:"path_in,omitempty"`

	// path lt
	PathLt *string `json:"path_lt,omitempty"`

	// path lte
	PathLte *string `json:"path_lte,omitempty"`

	// path not
	PathNot *string `json:"path_not,omitempty"`

	// path not contains
	PathNotContains *string `json:"path_not_contains,omitempty"`

	// path not ends with
	PathNotEndsWith *string `json:"path_not_ends_with,omitempty"`

	// path not in
	PathNotIn []string `json:"path_not_in,omitempty"`

	// path not starts with
	PathNotStartsWith *string `json:"path_not_starts_with,omitempty"`

	// path starts with
	PathStartsWith *string `json:"path_starts_with,omitempty"`

	// size
	Size *float64 `json:"size,omitempty"`

	// size gt
	SizeGt *float64 `json:"size_gt,omitempty"`

	// size gte
	SizeGte *float64 `json:"size_gte,omitempty"`

	// size in
	SizeIn []float64 `json:"size_in,omitempty"`

	// size lt
	SizeLt *float64 `json:"size_lt,omitempty"`

	// size lte
	SizeLte *float64 `json:"size_lte,omitempty"`

	// size not
	SizeNot *float64 `json:"size_not,omitempty"`

	// size not in
	SizeNotIn []float64 `json:"size_not_in,omitempty"`

	// vm disks every
	VMDisksEvery struct {
		VMDiskWhereInput
	} `json:"vm_disks_every,omitempty"`

	// vm disks none
	VMDisksNone struct {
		VMDiskWhereInput
	} `json:"vm_disks_none,omitempty"`

	// vm disks some
	VMDisksSome struct {
		VMDiskWhereInput
	} `json:"vm_disks_some,omitempty"`

	// vm snapshots every
	VMSnapshotsEvery struct {
		VMSnapshotWhereInput
	} `json:"vm_snapshots_every,omitempty"`

	// vm snapshots none
	VMSnapshotsNone struct {
		VMSnapshotWhereInput
	} `json:"vm_snapshots_none,omitempty"`

	// vm snapshots some
	VMSnapshotsSome struct {
		VMSnapshotWhereInput
	} `json:"vm_snapshots_some,omitempty"`

	// vm templates every
	VMTemplatesEvery struct {
		VMTemplateWhereInput
	} `json:"vm_templates_every,omitempty"`

	// vm templates none
	VMTemplatesNone struct {
		VMTemplateWhereInput
	} `json:"vm_templates_none,omitempty"`

	// vm templates some
	VMTemplatesSome struct {
		VMTemplateWhereInput
	} `json:"vm_templates_some,omitempty"`
}

// Validate validates this content library image where input
func (m *ContentLibraryImageWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateClustersEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClustersSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfImagesEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfImagesNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElfImagesSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisksEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisksNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisksSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshotsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshotsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMSnapshotsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMTemplatesEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMTemplatesNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMTemplatesSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryImageWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *ContentLibraryImageWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *ContentLibraryImageWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *ContentLibraryImageWhereInput) validateClustersEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersEvery) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateClustersNone(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersNone) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateClustersSome(formats strfmt.Registry) error {
	if swag.IsZero(m.ClustersSome) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateElfImagesEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.ElfImagesEvery) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateElfImagesNone(formats strfmt.Registry) error {
	if swag.IsZero(m.ElfImagesNone) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateElfImagesSome(formats strfmt.Registry) error {
	if swag.IsZero(m.ElfImagesSome) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateLabelsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsEvery) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateLabelsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsNone) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateLabelsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsSome) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMDisksEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.VMDisksEvery) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMDisksNone(formats strfmt.Registry) error {
	if swag.IsZero(m.VMDisksNone) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMDisksSome(formats strfmt.Registry) error {
	if swag.IsZero(m.VMDisksSome) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMSnapshotsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshotsEvery) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMSnapshotsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshotsNone) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMSnapshotsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.VMSnapshotsSome) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMTemplatesEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.VMTemplatesEvery) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMTemplatesNone(formats strfmt.Registry) error {
	if swag.IsZero(m.VMTemplatesNone) { // not required
		return nil
	}

	return nil
}

func (m *ContentLibraryImageWhereInput) validateVMTemplatesSome(formats strfmt.Registry) error {
	if swag.IsZero(m.VMTemplatesSome) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this content library image where input based on the context it is used
func (m *ContentLibraryImageWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateClustersEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClustersSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElfImagesEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElfImagesNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateElfImagesSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisksEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisksNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisksSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshotsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshotsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMSnapshotsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMTemplatesEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMTemplatesNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMTemplatesSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryImageWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryImageWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContentLibraryImageWhereInput) contextValidateClustersEvery(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateClustersNone(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateClustersSome(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateElfImagesEvery(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateElfImagesNone(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateElfImagesSome(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateLabelsEvery(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateLabelsNone(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateLabelsSome(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMDisksEvery(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMDisksNone(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMDisksSome(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMSnapshotsEvery(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMSnapshotsNone(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMSnapshotsSome(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMTemplatesEvery(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMTemplatesNone(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ContentLibraryImageWhereInput) contextValidateVMTemplatesSome(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *ContentLibraryImageWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentLibraryImageWhereInput) UnmarshalBinary(b []byte) error {
	var res ContentLibraryImageWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
