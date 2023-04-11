/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsAddDevice struct for ModelsAddDevice
type ModelsAddDevice struct {
	ChildPrefix             []string         `json:"child_prefix,omitempty"`
	Discovery               bool             `json:"discovery,omitempty"`
	EndpointLocalAddressIp4 string           `json:"endpoint_local_address_ip4,omitempty"`
	Endpoints               []ModelsEndpoint `json:"endpoints,omitempty"`
	Hostname                string           `json:"hostname,omitempty"`
	OrganizationId          string           `json:"organization_id,omitempty"`
	Os                      string           `json:"os,omitempty"`
	PublicKey               string           `json:"public_key,omitempty"`
	Relay                   bool             `json:"relay,omitempty"`
	SymmetricNat            bool             `json:"symmetric_nat,omitempty"`
	TunnelIp                string           `json:"tunnel_ip,omitempty"`
	TunnelIpV6              string           `json:"tunnel_ip_v6,omitempty"`
	UserId                  string           `json:"user_id,omitempty"`
}
