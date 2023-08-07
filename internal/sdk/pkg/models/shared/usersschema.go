// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UsersSchema - Users and root roles
type UsersSchema struct {
	// A list of [root roles](https://docs.getunleash.io/reference/rbac#standard-roles) in the Unleash instance.
	RootRoles []RoleSchema `json:"rootRoles,omitempty"`
	// A list of users in the Unleash instance.
	Users []UserSchema `json:"users"`
}