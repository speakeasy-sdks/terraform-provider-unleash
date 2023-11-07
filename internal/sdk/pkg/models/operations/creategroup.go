// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// CreateGroupUsersResponse409ResponseBody - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type CreateGroupUsersResponse409ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateGroupUsersResponse409ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateGroupUsersResponse409ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateGroupUsersResponse409ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateGroupUsersResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreateGroupUsersResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateGroupUsersResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateGroupUsersResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateGroupUsersResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateGroupUsersResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreateGroupUsersResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateGroupUsersResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateGroupUsersResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateGroupUsersResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateGroupResponseBody - The request data does not match what we expect.
type CreateGroupResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateGroupResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateGroupResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateGroupResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreateGroupResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *CreateGroupResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *CreateGroupUsersResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *CreateGroupUsersResponseResponseBody
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	FourHundredAndNineApplicationJSONObject *CreateGroupUsersResponse409ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// groupSchema
	GroupSchema *shared.GroupSchema
}

func (o *CreateGroupResponse) GetFourHundredApplicationJSONObject() *CreateGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *CreateGroupResponse) GetFourHundredAndOneApplicationJSONObject() *CreateGroupUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *CreateGroupResponse) GetFourHundredAndThreeApplicationJSONObject() *CreateGroupUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *CreateGroupResponse) GetFourHundredAndNineApplicationJSONObject() *CreateGroupUsersResponse409ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndNineApplicationJSONObject
}

func (o *CreateGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateGroupResponse) GetGroupSchema() *shared.GroupSchema {
	if o == nil {
		return nil
	}
	return o.GroupSchema
}
