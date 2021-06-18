// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateDataVolume create data volume
// swagger:model CreateDataVolume
type CreateDataVolume struct {

	// PVM Instance (ID or Name) to base volume affinity policy against; required if requesting affinity and affinityVolume is not provided
	AffinityPVMInstance *string `json:"affinityPVMInstance,omitempty"`

	// Affinity policy for data volume being created; ignored if volumePool provided; for policy affinity requires one of affinityPVMInstance or affinityVolume to be specified; for policy anti-affinity requires one of antiAffinityPVMInstances or antiAffinityVolumes to be specified
	// Enum: [affinity anti-affinity]
	AffinityPolicy *string `json:"affinityPolicy,omitempty"`

	// Volume (ID or Name) to base volume affinity policy against; required if requesting affinity and affinityPVMInstance is not provided
	AffinityVolume *string `json:"affinityVolume,omitempty"`

	// List of pvmInstances to base volume anti-affinity policy against; required if requesting anti-affinity and antiAffinityVolumes is not provided
	AntiAffinityPVMInstances []string `json:"antiAffinityPVMInstances"`

	// List of volumes to base volume anti-affinity policy against; required if requesting anti-affinity and antiAffinityPVMInstances is not provided
	AntiAffinityVolumes []string `json:"antiAffinityVolumes"`

	// Type of Disk, required if affinityPolicy and volumePool not provided, otherwise ignored
	DiskType string `json:"diskType,omitempty"`

	// Volume Name
	// Required: true
	Name *string `json:"name"`

	// Indicates if the volume is shareable between VMs
	Shareable *bool `json:"shareable,omitempty"`

	// Volume Size (GB)
	// Required: true
	Size *float64 `json:"size"`

	// Volume pool where the volume will be created; if provided then diskType and affinityPolicy values will be ignored
	VolumePool string `json:"volumePool,omitempty"`
}

// Validate validates this create data volume
func (m *CreateDataVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAffinityPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createDataVolumeTypeAffinityPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["affinity","anti-affinity"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDataVolumeTypeAffinityPolicyPropEnum = append(createDataVolumeTypeAffinityPolicyPropEnum, v)
	}
}

const (

	// CreateDataVolumeAffinityPolicyAffinity captures enum value "affinity"
	CreateDataVolumeAffinityPolicyAffinity string = "affinity"

	// CreateDataVolumeAffinityPolicyAntiAffinity captures enum value "anti-affinity"
	CreateDataVolumeAffinityPolicyAntiAffinity string = "anti-affinity"
)

// prop value enum
func (m *CreateDataVolume) validateAffinityPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createDataVolumeTypeAffinityPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateDataVolume) validateAffinityPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.AffinityPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateAffinityPolicyEnum("affinityPolicy", "body", *m.AffinityPolicy); err != nil {
		return err
	}

	return nil
}

func (m *CreateDataVolume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateDataVolume) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateDataVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDataVolume) UnmarshalBinary(b []byte) error {
	var res CreateDataVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
