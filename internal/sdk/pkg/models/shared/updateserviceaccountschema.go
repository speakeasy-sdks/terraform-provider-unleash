// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"terraform/internal/sdk/pkg/utils"
)

// UpdateServiceAccountSchema - Describes the properties required to update a service account
type UpdateServiceAccountSchema struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The name of the service account
	Name *string `json:"name,omitempty"`
	// The id of the root role for the service account
	RootRole *int64 `json:"rootRole,omitempty"`
}

func (u UpdateServiceAccountSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateServiceAccountSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateServiceAccountSchema) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *UpdateServiceAccountSchema) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateServiceAccountSchema) GetRootRole() *int64 {
	if o == nil {
		return nil
	}
	return o.RootRole
}
