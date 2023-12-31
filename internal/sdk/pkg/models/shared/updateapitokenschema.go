// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

// UpdateAPITokenSchema - An object with fields to updated for a given API token.
type UpdateAPITokenSchema struct {
	// The new time when this token should expire.
	ExpiresAt time.Time `json:"expiresAt"`
}

func (u UpdateAPITokenSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateAPITokenSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateAPITokenSchema) GetExpiresAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ExpiresAt
}
