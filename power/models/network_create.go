// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkCreate network create
// Example: {"cidr":"192.168.1.0/24","gateway":"192.168.1.1","ipAddressRanges":[{"endingIPAddress":"192.168.1.254","startingIPAddress":"192.168.1.2"}],"mtu":1450,"name":"sample-network","type":"vlan"}
//
// swagger:model NetworkCreate
type NetworkCreate struct {

	// access config
	AccessConfig AccessConfig `json:"accessConfig,omitempty"`

	// Indicates if the network is advertised externally of the workspace to PER and\or peer networks
	// Enum: ["enable","disable"]
	Advertise *string `json:"advertise,omitempty"`

	// Indicates if ARP broadcast is enabled
	// Enum: ["enable","disable"]
	ArpBroadcast *string `json:"arpBroadcast,omitempty"`

	// Network in CIDR notation (192.168.0.0/24)
	Cidr string `json:"cidr,omitempty"`

	// DNS Servers. If not specified, default is 127.0.0.1 for 'vlan' (private network) and 9.9.9.9 for 'pub-vlan' (public network)
	DNSServers []string `json:"dnsServers"`

	// Gateway IP Address
	Gateway string `json:"gateway,omitempty"`

	// IP Address Ranges
	IPAddressRanges []*IPAddressRange `json:"ipAddressRanges"`

	// (deprecated - replaced by mtu) Enable MTU Jumbo Network (for multi-zone locations only)
	Jumbo bool `json:"jumbo,omitempty"`

	// Maximum transmission unit
	// Maximum: 9000
	// Minimum: 1450
	Mtu *int64 `json:"mtu,omitempty"`

	// Network Name
	// Max Length: 128
	// Pattern: ^[a-zA-Z0-9-_][a-zA-Z0-9-_]*$
	Name string `json:"name,omitempty"`

	// Network Peer information
	Peer *NetworkCreatePeer `json:"peer,omitempty"`

	// Type of Network - 'vlan' (private network) 'pub-vlan' (public network) 'dhcp-vlan' (for satellite locations only)
	// Required: true
	// Enum: ["vlan","pub-vlan","dhcp-vlan"]
	Type *string `json:"type"`

	// user tags
	UserTags Tags `json:"userTags,omitempty"`
}

// Validate validates this network create
func (m *NetworkCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAdvertise(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArpBroadcast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddressRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreate) validateAccessConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessConfig) { // not required
		return nil
	}

	if err := m.AccessConfig.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accessConfig")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accessConfig")
		}
		return err
	}

	return nil
}

var networkCreateTypeAdvertisePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enable","disable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkCreateTypeAdvertisePropEnum = append(networkCreateTypeAdvertisePropEnum, v)
	}
}

const (

	// NetworkCreateAdvertiseEnable captures enum value "enable"
	NetworkCreateAdvertiseEnable string = "enable"

	// NetworkCreateAdvertiseDisable captures enum value "disable"
	NetworkCreateAdvertiseDisable string = "disable"
)

// prop value enum
func (m *NetworkCreate) validateAdvertiseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkCreateTypeAdvertisePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkCreate) validateAdvertise(formats strfmt.Registry) error {
	if swag.IsZero(m.Advertise) { // not required
		return nil
	}

	// value enum
	if err := m.validateAdvertiseEnum("advertise", "body", *m.Advertise); err != nil {
		return err
	}

	return nil
}

var networkCreateTypeArpBroadcastPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enable","disable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkCreateTypeArpBroadcastPropEnum = append(networkCreateTypeArpBroadcastPropEnum, v)
	}
}

const (

	// NetworkCreateArpBroadcastEnable captures enum value "enable"
	NetworkCreateArpBroadcastEnable string = "enable"

	// NetworkCreateArpBroadcastDisable captures enum value "disable"
	NetworkCreateArpBroadcastDisable string = "disable"
)

// prop value enum
func (m *NetworkCreate) validateArpBroadcastEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkCreateTypeArpBroadcastPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkCreate) validateArpBroadcast(formats strfmt.Registry) error {
	if swag.IsZero(m.ArpBroadcast) { // not required
		return nil
	}

	// value enum
	if err := m.validateArpBroadcastEnum("arpBroadcast", "body", *m.ArpBroadcast); err != nil {
		return err
	}

	return nil
}

func (m *NetworkCreate) validateIPAddressRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddressRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAddressRanges); i++ {
		if swag.IsZero(m.IPAddressRanges[i]) { // not required
			continue
		}

		if m.IPAddressRanges[i] != nil {
			if err := m.IPAddressRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkCreate) validateMtu(formats strfmt.Registry) error {
	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", *m.Mtu, 1450, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", *m.Mtu, 9000, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkCreate) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 128); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[a-zA-Z0-9-_][a-zA-Z0-9-_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NetworkCreate) validatePeer(formats strfmt.Registry) error {
	if swag.IsZero(m.Peer) { // not required
		return nil
	}

	if m.Peer != nil {
		if err := m.Peer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer")
			}
			return err
		}
	}

	return nil
}

var networkCreateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vlan","pub-vlan","dhcp-vlan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkCreateTypeTypePropEnum = append(networkCreateTypeTypePropEnum, v)
	}
}

const (

	// NetworkCreateTypeVlan captures enum value "vlan"
	NetworkCreateTypeVlan string = "vlan"

	// NetworkCreateTypePubDashVlan captures enum value "pub-vlan"
	NetworkCreateTypePubDashVlan string = "pub-vlan"

	// NetworkCreateTypeDhcpDashVlan captures enum value "dhcp-vlan"
	NetworkCreateTypeDhcpDashVlan string = "dhcp-vlan"
)

// prop value enum
func (m *NetworkCreate) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkCreateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkCreate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *NetworkCreate) validateUserTags(formats strfmt.Registry) error {
	if swag.IsZero(m.UserTags) { // not required
		return nil
	}

	if err := m.UserTags.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userTags")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userTags")
		}
		return err
	}

	return nil
}

// ContextValidate validate this network create based on the context it is used
func (m *NetworkCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPAddressRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkCreate) contextValidateAccessConfig(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AccessConfig) { // not required
		return nil
	}

	if err := m.AccessConfig.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accessConfig")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accessConfig")
		}
		return err
	}

	return nil
}

func (m *NetworkCreate) contextValidateIPAddressRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAddressRanges); i++ {

		if m.IPAddressRanges[i] != nil {

			if swag.IsZero(m.IPAddressRanges[i]) { // not required
				return nil
			}

			if err := m.IPAddressRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkCreate) contextValidatePeer(ctx context.Context, formats strfmt.Registry) error {

	if m.Peer != nil {

		if swag.IsZero(m.Peer) { // not required
			return nil
		}

		if err := m.Peer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("peer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("peer")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkCreate) contextValidateUserTags(ctx context.Context, formats strfmt.Registry) error {

	if err := m.UserTags.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userTags")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("userTags")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCreate) UnmarshalBinary(b []byte) error {
	var res NetworkCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
