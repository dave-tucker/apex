/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsRegKey struct for ModelsRegKey
type ModelsRegKey struct {
	// BearerToken is the bearer token the client should use to authenticate the device registration request.
	BearerToken string `json:"bearer_token,omitempty"`
	Description string `json:"description,omitempty"`
	DeviceId    string `json:"device_id,omitempty"`
	Expiration  string `json:"expiration,omitempty"`
	Id          string `json:"id,omitempty"`
	OwnerId     string `json:"owner_id,omitempty"`
	VpcId       string `json:"vpc_id,omitempty"`
}
