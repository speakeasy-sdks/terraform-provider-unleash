// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type AddPublicSignupTokenUserRequest struct {
	// createInvitedUserSchema
	CreateInvitedUserSchema shared.CreateInvitedUserSchema `request:"mediaType=application/json"`
	Token                   string                         `pathParam:"style=simple,explode=false,name=token"`
}

// AddPublicSignupTokenUser409ApplicationJSON - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type AddPublicSignupTokenUser409ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// AddPublicSignupTokenUser400ApplicationJSON - The request data does not match what we expect.
type AddPublicSignupTokenUser400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type AddPublicSignupTokenUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	AddPublicSignupTokenUser400ApplicationJSONObject *AddPublicSignupTokenUser400ApplicationJSON
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	AddPublicSignupTokenUser409ApplicationJSONObject *AddPublicSignupTokenUser409ApplicationJSON
	// userSchema
	UserSchema *shared.UserSchema
}
