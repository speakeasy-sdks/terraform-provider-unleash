// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"terraform/internal/sdk/pkg/utils"
	"time"
)

// GroupSchema - A detailed information about a user group
type GroupSchema struct {
	// When was this group created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A user who created this group
	CreatedBy *string `json:"createdBy,omitempty"`
	// A custom description of the group
	Description *string `json:"description,omitempty"`
	// The group id
	ID *int64 `json:"id,omitempty"`
	// A list of SSO groups that should map to this Unleash group
	MappingsSSO []string `json:"mappingsSSO,omitempty"`
	// The name of the group
	Name string `json:"name"`
	// A list of projects where this group is used
	Projects []string `json:"projects,omitempty"`
	// A role id that is used as the root role for all users in this group. This can be either the id of the Viewer, Editor or Admin role.
	RootRole *float64 `json:"rootRole,omitempty"`
	// A list of users belonging to this group
	Users []GroupUserModelSchema `json:"users,omitempty"`
}

func (g GroupSchema) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GroupSchema) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GroupSchema) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GroupSchema) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *GroupSchema) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GroupSchema) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GroupSchema) GetMappingsSSO() []string {
	if o == nil {
		return nil
	}
	return o.MappingsSSO
}

func (o *GroupSchema) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GroupSchema) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *GroupSchema) GetRootRole() *float64 {
	if o == nil {
		return nil
	}
	return o.RootRole
}

func (o *GroupSchema) GetUsers() []GroupUserModelSchema {
	if o == nil {
		return nil
	}
	return o.Users
}
