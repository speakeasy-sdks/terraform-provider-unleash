// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// CreateTagTypeTagsResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type CreateTagTypeTagsResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateTagTypeTagsResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateTagTypeTagsResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateTagTypeTagsResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateTagTypeTagsResponse409ResponseBody - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type CreateTagTypeTagsResponse409ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateTagTypeTagsResponse409ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateTagTypeTagsResponse409ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateTagTypeTagsResponse409ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateTagTypeTagsResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreateTagTypeTagsResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateTagTypeTagsResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateTagTypeTagsResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateTagTypeTagsResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateTagTypeTagsResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreateTagTypeTagsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateTagTypeTagsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateTagTypeTagsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateTagTypeTagsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateTagTypeResponseBody - The request data does not match what we expect.
type CreateTagTypeResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateTagTypeResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateTagTypeResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateTagTypeResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreateTagTypeResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *CreateTagTypeResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *CreateTagTypeTagsResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *CreateTagTypeTagsResponseResponseBody
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	FourHundredAndNineApplicationJSONObject *CreateTagTypeTagsResponse409ResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *CreateTagTypeTagsResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The resource was successfully created.
	TagTypeSchema *shared.TagTypeSchema
}

func (o *CreateTagTypeResponse) GetFourHundredApplicationJSONObject() *CreateTagTypeResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *CreateTagTypeResponse) GetFourHundredAndOneApplicationJSONObject() *CreateTagTypeTagsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *CreateTagTypeResponse) GetFourHundredAndThreeApplicationJSONObject() *CreateTagTypeTagsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *CreateTagTypeResponse) GetFourHundredAndNineApplicationJSONObject() *CreateTagTypeTagsResponse409ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndNineApplicationJSONObject
}

func (o *CreateTagTypeResponse) GetFourHundredAndFifteenApplicationJSONObject() *CreateTagTypeTagsResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *CreateTagTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTagTypeResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateTagTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTagTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTagTypeResponse) GetTagTypeSchema() *shared.TagTypeSchema {
	if o == nil {
		return nil
	}
	return o.TagTypeSchema
}
