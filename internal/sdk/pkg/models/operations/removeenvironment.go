// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveEnvironmentRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *RemoveEnvironmentRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// RemoveEnvironmentEnvironmentsResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type RemoveEnvironmentEnvironmentsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *RemoveEnvironmentEnvironmentsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RemoveEnvironmentEnvironmentsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RemoveEnvironmentEnvironmentsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// RemoveEnvironmentResponseBody - The request data does not match what we expect.
type RemoveEnvironmentResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *RemoveEnvironmentResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RemoveEnvironmentResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RemoveEnvironmentResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type RemoveEnvironmentResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *RemoveEnvironmentResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *RemoveEnvironmentEnvironmentsResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemoveEnvironmentResponse) GetFourHundredApplicationJSONObject() *RemoveEnvironmentResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *RemoveEnvironmentResponse) GetFourHundredAndOneApplicationJSONObject() *RemoveEnvironmentEnvironmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *RemoveEnvironmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveEnvironmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveEnvironmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
