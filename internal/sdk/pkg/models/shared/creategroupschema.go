// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"terraform/internal/sdk/pkg/utils"
)

// CreateGroupSchemaUsersUser - A minimal user object
type CreateGroupSchemaUsersUser struct {
	// The user id
	ID int64 `json:"id"`
}

func (o *CreateGroupSchemaUsersUser) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// CreateGroupSchemaUsers - A minimal user object
type CreateGroupSchemaUsers struct {
	// A minimal user object
	User CreateGroupSchemaUsersUser `json:"user"`
}

func (o *CreateGroupSchemaUsers) GetUser() CreateGroupSchemaUsersUser {
	if o == nil {
		return CreateGroupSchemaUsersUser{}
	}
	return o.User
}

// CreateGroupSchema - A detailed information about a user group
type CreateGroupSchema struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// A custom description of the group
	Description *string `json:"description,omitempty"`
	// A list of SSO groups that should map to this Unleash group
	MappingsSSO []string `json:"mappingsSSO,omitempty"`
	// The name of the group
	Name string `json:"name"`
	// A role id that is used as the root role for all users in this group. This can be either the id of the Viewer, Editor or Admin role.
	RootRole *float64 `json:"rootRole,omitempty"`
	// A list of users belonging to this group
	Users []CreateGroupSchemaUsers `json:"users,omitempty"`
}

func (c CreateGroupSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateGroupSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateGroupSchema) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *CreateGroupSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateGroupSchema) GetMappingsSSO() []string {
	if o == nil {
		return nil
	}
	return o.MappingsSSO
}

func (o *CreateGroupSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateGroupSchema) GetRootRole() *float64 {
	if o == nil {
		return nil
	}
	return o.RootRole
}

func (o *CreateGroupSchema) GetUsers() []CreateGroupSchemaUsers {
	if o == nil {
		return nil
	}
	return o.Users
}
