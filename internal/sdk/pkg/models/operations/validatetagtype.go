// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// ValidateTagTypeTagsResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type ValidateTagTypeTagsResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateTagTypeTagsResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateTagTypeTagsResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateTagTypeTagsResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ValidateTagTypeTagsResponse409ResponseBody - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type ValidateTagTypeTagsResponse409ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateTagTypeTagsResponse409ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateTagTypeTagsResponse409ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateTagTypeTagsResponse409ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ValidateTagTypeTagsResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ValidateTagTypeTagsResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateTagTypeTagsResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateTagTypeTagsResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateTagTypeTagsResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ValidateTagTypeTagsResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ValidateTagTypeTagsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateTagTypeTagsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateTagTypeTagsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateTagTypeTagsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ValidateTagTypeResponseBody - The request data does not match what we expect.
type ValidateTagTypeResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateTagTypeResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateTagTypeResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateTagTypeResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ValidateTagTypeResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *ValidateTagTypeResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *ValidateTagTypeTagsResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *ValidateTagTypeTagsResponseResponseBody
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	FourHundredAndNineApplicationJSONObject *ValidateTagTypeTagsResponse409ResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *ValidateTagTypeTagsResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// validateTagTypeSchema
	ValidateTagTypeSchema *shared.ValidateTagTypeSchema
}

func (o *ValidateTagTypeResponse) GetFourHundredApplicationJSONObject() *ValidateTagTypeResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *ValidateTagTypeResponse) GetFourHundredAndOneApplicationJSONObject() *ValidateTagTypeTagsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *ValidateTagTypeResponse) GetFourHundredAndThreeApplicationJSONObject() *ValidateTagTypeTagsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ValidateTagTypeResponse) GetFourHundredAndNineApplicationJSONObject() *ValidateTagTypeTagsResponse409ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndNineApplicationJSONObject
}

func (o *ValidateTagTypeResponse) GetFourHundredAndFifteenApplicationJSONObject() *ValidateTagTypeTagsResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *ValidateTagTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidateTagTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidateTagTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidateTagTypeResponse) GetValidateTagTypeSchema() *shared.ValidateTagTypeSchema {
	if o == nil {
		return nil
	}
	return o.ValidateTagTypeSchema
}
