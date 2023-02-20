package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Device is a unique, end-user device.
// Devices belong to one User and may be onboarded into an organization
type Device struct {
	Base
	UserID                   string         `json:"user_id"`
	OrganizationID           uuid.UUID      `json:"organization_id"`
	PublicKey                string         `json:"public_key" gorm:"uniqueIndex"`
	LocalIP                  string         `json:"local_ip"`
	AllowedIPs               pq.StringArray `json:"allowed_ips" gorm:"type:text[]"`
	TunnelIP                 string         `json:"tunnel_ip"`
	ChildPrefix              pq.StringArray `json:"child_prefix" gorm:"type:text[]"`
	Relay                    bool           `json:"relay"`
	OrganizationPrefix       string         `json:"organization_prefix"`
	ReflexiveIPv4            string         `json:"reflexive_ip4"`
	EndpointLocalAddressIPv4 string         `json:"endpoint_local_address_ip4"`
	SymmetricNat             bool           `json:"symmetric_nat"`
	Hostname                 string         `json:"hostname"`
}

// AddDevice is the information needed to add a new Device.
type AddDevice struct {
	UserID                   string         `json:"user_id" example:"694aa002-5d19-495e-980b-3d8fd508ea10"`
	OrganizationID           uuid.UUID      `json:"organization_id" example:"694aa002-5d19-495e-980b-3d8fd508ea10"`
	PublicKey                string         `json:"public_key"`
	LocalIP                  string         `json:"local_ip" example:"10.1.1.1"`
	TunnelIP                 string         `json:"tunnel_ip" example:"1.2.3.4"`
	ChildPrefix              pq.StringArray `json:"child_prefix" example:"172.16.42.0/24"`
	Relay                    bool           `json:"relay"`
	ReflexiveIPv4            string         `json:"reflexive_ip4"`
	EndpointLocalAddressIPv4 string         `json:"endpoint_local_address_ip4" example:"1.2.3.4"`
	SymmetricNat             bool           `json:"symmetric_nat"`
	Hostname                 string         `json:"hostname" example:"myhost"`
}

// UpdateDevice is the information needed to update a Device.
type UpdateDevice struct {
	OrganizationID           uuid.UUID `json:"organization_id" example:"694aa002-5d19-495e-980b-3d8fd508ea10"`
	LocalIP                  string    `json:"local_ip" example:"10.1.1.1"`
	ChildPrefix              []string  `json:"child_prefix" example:"172.16.42.0/24"`
	ReflexiveIPv4            string    `json:"reflexive_ip4"`
	EndpointLocalAddressIPv4 string    `json:"endpoint_local_address_ip4" example:"1.2.3.4"`
	SymmetricNat             bool      `json:"symmetric_nat"`
	Hostname                 string    `json:"hostname" example:"myhost"`
}
