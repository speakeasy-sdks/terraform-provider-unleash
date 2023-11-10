// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

type UserWithProjectRoleSchema struct {
	AddedAt  *time.Time `json:"addedAt,omitempty"`
	Email    *string    `json:"email,omitempty"`
	ID       float64    `json:"id"`
	ImageURL *string    `json:"imageUrl,omitempty"`
	IsAPI    bool       `json:"isAPI"`
	Name     *string    `json:"name,omitempty"`
	RoleID   *float64   `json:"roleId,omitempty"`
}

func (u UserWithProjectRoleSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserWithProjectRoleSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UserWithProjectRoleSchema) GetAddedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *UserWithProjectRoleSchema) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *UserWithProjectRoleSchema) GetID() float64 {
	if o == nil {
		return 0.0
	}
	return o.ID
}

func (o *UserWithProjectRoleSchema) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *UserWithProjectRoleSchema) GetIsAPI() bool {
	if o == nil {
		return false
	}
	return o.IsAPI
}

func (o *UserWithProjectRoleSchema) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UserWithProjectRoleSchema) GetRoleID() *float64 {
	if o == nil {
		return nil
	}
	return o.RoleID
}
