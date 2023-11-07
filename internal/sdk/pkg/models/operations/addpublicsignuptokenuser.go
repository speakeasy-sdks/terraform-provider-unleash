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

func (o *AddPublicSignupTokenUserRequest) GetCreateInvitedUserSchema() shared.CreateInvitedUserSchema {
	if o == nil {
		return shared.CreateInvitedUserSchema{}
	}
	return o.CreateInvitedUserSchema
}

func (o *AddPublicSignupTokenUserRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

// AddPublicSignupTokenUserPublicSignupTokensResponseBody - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type AddPublicSignupTokenUserPublicSignupTokensResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *AddPublicSignupTokenUserPublicSignupTokensResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddPublicSignupTokenUserPublicSignupTokensResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *AddPublicSignupTokenUserPublicSignupTokensResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// AddPublicSignupTokenUserResponseBody - The request data does not match what we expect.
type AddPublicSignupTokenUserResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *AddPublicSignupTokenUserResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddPublicSignupTokenUserResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *AddPublicSignupTokenUserResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type AddPublicSignupTokenUserResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *AddPublicSignupTokenUserResponseBody
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	FourHundredAndNineApplicationJSONObject *AddPublicSignupTokenUserPublicSignupTokensResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// userSchema
	UserSchema *shared.UserSchema
}

func (o *AddPublicSignupTokenUserResponse) GetFourHundredApplicationJSONObject() *AddPublicSignupTokenUserResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *AddPublicSignupTokenUserResponse) GetFourHundredAndNineApplicationJSONObject() *AddPublicSignupTokenUserPublicSignupTokensResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndNineApplicationJSONObject
}

func (o *AddPublicSignupTokenUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddPublicSignupTokenUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddPublicSignupTokenUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddPublicSignupTokenUserResponse) GetUserSchema() *shared.UserSchema {
	if o == nil {
		return nil
	}
	return o.UserSchema
}
