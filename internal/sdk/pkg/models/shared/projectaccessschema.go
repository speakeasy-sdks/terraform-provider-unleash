// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ProjectAccessSchema struct {
	Groups []GroupSchema               `json:"groups"`
	Roles  []RoleSchema                `json:"roles"`
	Users  []UserWithProjectRoleSchema `json:"users"`
}
