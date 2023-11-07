// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetUserRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetUserUsersResponseResponseBody - The requested resource was not found.
type GetUserUsersResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetUserUsersResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetUserUsersResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetUserUsersResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetUserUsersResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetUserUsersResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetUserUsersResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetUserUsersResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetUserUsersResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetUserResponseBody - The request data does not match what we expect.
type GetUserResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetUserResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetUserResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetUserResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetUserResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *GetUserResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetUserUsersResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *GetUserUsersResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// userSchema
	UserSchema *shared.UserSchema
}

func (o *GetUserResponse) GetFourHundredApplicationJSONObject() *GetUserResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *GetUserResponse) GetFourHundredAndOneApplicationJSONObject() *GetUserUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetUserResponse) GetFourHundredAndFourApplicationJSONObject() *GetUserUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *GetUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUserResponse) GetUserSchema() *shared.UserSchema {
	if o == nil {
		return nil
	}
	return o.UserSchema
}
