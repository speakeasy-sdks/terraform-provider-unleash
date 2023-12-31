// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteUserRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteUserUsersResponseResponseBody - The requested resource was not found.
type DeleteUserUsersResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteUserUsersResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteUserUsersResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteUserUsersResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteUserUsersResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeleteUserUsersResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteUserUsersResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteUserUsersResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteUserUsersResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteUserResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeleteUserResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteUserResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteUserResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteUserResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeleteUserResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *DeleteUserResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *DeleteUserUsersResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *DeleteUserUsersResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteUserResponse) GetFourHundredAndOneApplicationJSONObject() *DeleteUserResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *DeleteUserResponse) GetFourHundredAndThreeApplicationJSONObject() *DeleteUserUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *DeleteUserResponse) GetFourHundredAndFourApplicationJSONObject() *DeleteUserUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *DeleteUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
