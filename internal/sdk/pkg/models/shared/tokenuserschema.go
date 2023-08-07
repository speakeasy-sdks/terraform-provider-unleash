// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TokenUserSchema - A user identified by a token
type TokenUserSchema struct {
	// A username or email identifying which user created this token
	CreatedBy string `json:"createdBy"`
	// The email of the user
	Email string `json:"email"`
	// The user id
	ID int64 `json:"id"`
	// The name of the user
	Name string `json:"name"`
	// A role holds permissions to allow Unleash to decide what actions a role holder is allowed to perform
	Role RoleSchema `json:"role"`
	// A token uniquely identifying a user
	Token string `json:"token"`
}