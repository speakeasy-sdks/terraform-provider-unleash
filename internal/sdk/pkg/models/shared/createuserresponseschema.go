// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// CreateUserResponseSchema - An Unleash user after creation
type CreateUserResponseSchema struct {
	// A user is either an actual User or a Service Account
	AccountType *string `json:"accountType,omitempty"`
	// The user was created at this time
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Email of the user
	Email *string `json:"email,omitempty"`
	// Is the welcome email sent to the user or not
	EmailSent *bool `json:"emailSent,omitempty"`
	// The user id
	ID int64 `json:"id"`
	// URL used for the userprofile image
	ImageURL *string `json:"imageUrl,omitempty"`
	// If the user is actively inviting other users, this is the link that can be shared with other users
	InviteLink *string `json:"inviteLink,omitempty"`
	// (Deprecated): Used internally to know which operations the user should be allowed to perform
	//
	// @deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	IsAPI *bool `json:"isAPI,omitempty"`
	// How many unsuccessful attempts at logging in has the user made
	LoginAttempts *int64 `json:"loginAttempts,omitempty"`
	// Name of the user
	Name *string `json:"name,omitempty"`
	// Deprecated
	Permissions []string `json:"permissions,omitempty"`
	// The role to assign to the user. Can be either the role's ID or its unique name.
	RootRole *CreateUserSchemaRootRole `json:"rootRole,omitempty"`
	// The last time this user logged in
	SeenAt *time.Time `json:"seenAt,omitempty"`
	// A unique username for the user
	Username *string `json:"username,omitempty"`
}
