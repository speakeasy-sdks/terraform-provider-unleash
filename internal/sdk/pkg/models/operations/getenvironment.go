// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetEnvironmentRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetEnvironmentRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// GetEnvironmentEnvironmentsResponseResponseBody - The requested resource was not found.
type GetEnvironmentEnvironmentsResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentEnvironmentsResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentEnvironmentsResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentEnvironmentsResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetEnvironmentEnvironmentsResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetEnvironmentEnvironmentsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentEnvironmentsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentEnvironmentsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentEnvironmentsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetEnvironmentResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetEnvironmentResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetEnvironmentResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetEnvironmentResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *GetEnvironmentEnvironmentsResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *GetEnvironmentEnvironmentsResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// environmentSchema
	EnvironmentSchema *shared.EnvironmentSchema
}

func (o *GetEnvironmentResponse) GetFourHundredAndOneApplicationJSONObject() *GetEnvironmentResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetEnvironmentResponse) GetFourHundredAndThreeApplicationJSONObject() *GetEnvironmentEnvironmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetEnvironmentResponse) GetFourHundredAndFourApplicationJSONObject() *GetEnvironmentEnvironmentsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *GetEnvironmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEnvironmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEnvironmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEnvironmentResponse) GetEnvironmentSchema() *shared.EnvironmentSchema {
	if o == nil {
		return nil
	}
	return o.EnvironmentSchema
}
