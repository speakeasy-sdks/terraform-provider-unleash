// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateUserSchema - The payload must contain at least one of the name and email properties, though which one is up to you. For the user to be able to log in to the system, the user must have an email.
type CreateUserSchema struct {
	// The user's email address. Must be provided if username is not provided.
	Email *string `json:"email,omitempty"`
	// The user's name (not the user's username).
	Name *string `json:"name,omitempty"`
	// Password for the user
	Password *string `json:"password,omitempty"`
	// The role to assign to the user. Can be either the role's ID or its unique name.
	RootRole CreateUserSchemaRootRole `json:"rootRole"`
	// Whether to send a welcome email with a login link to the user or not. Defaults to `true`.
	SendEmail *bool `json:"sendEmail,omitempty"`
	// The user's username. Must be provided if email is not provided.
	Username *string `json:"username,omitempty"`
}
