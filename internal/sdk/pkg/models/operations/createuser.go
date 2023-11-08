// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// CreateUserUsersResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreateUserUsersResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateUserUsersResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateUserUsersResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateUserUsersResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateUserUsersResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreateUserUsersResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateUserUsersResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateUserUsersResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateUserUsersResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateUserResponseBody - The request data does not match what we expect.
type CreateUserResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateUserResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateUserResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateUserResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreateUserResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *CreateUserResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *CreateUserUsersResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *CreateUserUsersResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The resource was successfully created.
	CreateUserResponseSchema *shared.CreateUserResponseSchema
}

func (o *CreateUserResponse) GetFourHundredApplicationJSONObject() *CreateUserResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *CreateUserResponse) GetFourHundredAndOneApplicationJSONObject() *CreateUserUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *CreateUserResponse) GetFourHundredAndThreeApplicationJSONObject() *CreateUserUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *CreateUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUserResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUserResponse) GetCreateUserResponseSchema() *shared.CreateUserResponseSchema {
	if o == nil {
		return nil
	}
	return o.CreateUserResponseSchema
}
