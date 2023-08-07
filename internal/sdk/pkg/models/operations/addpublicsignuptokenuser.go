// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unleash/internal/sdk/pkg/models/shared"
)

type AddPublicSignupTokenUserRequest struct {
	// createInvitedUserSchema
	CreateInvitedUserSchema shared.CreateInvitedUserSchema `request:"mediaType=application/json"`
	Token                   string                         `pathParam:"style=simple,explode=false,name=token"`
}

type AddPublicSignupTokenUserResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	AddPublicSignupTokenUser409Response *shared.AddPublicSignupTokenUser409Response
	// userSchema
	UserSchema *shared.UserSchema
	// The request data does not match what we expect.
	ValidatePublicSignupToken400Response *shared.ValidatePublicSignupToken400Response
}