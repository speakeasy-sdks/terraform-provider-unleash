// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateEnvironmentRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// updateEnvironmentSchema
	UpdateEnvironmentSchema shared.UpdateEnvironmentSchema `request:"mediaType=application/json"`
}

func (o *UpdateEnvironmentRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateEnvironmentRequest) GetUpdateEnvironmentSchema() shared.UpdateEnvironmentSchema {
	if o == nil {
		return shared.UpdateEnvironmentSchema{}
	}
	return o.UpdateEnvironmentSchema
}

// UpdateEnvironmentEnvironmentsResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type UpdateEnvironmentEnvironmentsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateEnvironmentEnvironmentsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateEnvironmentEnvironmentsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateEnvironmentEnvironmentsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateEnvironmentResponseBody - The request data does not match what we expect.
type UpdateEnvironmentResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateEnvironmentResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateEnvironmentResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateEnvironmentResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateEnvironmentResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *UpdateEnvironmentResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *UpdateEnvironmentEnvironmentsResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// environmentSchema
	EnvironmentSchema *shared.EnvironmentSchema
}

func (o *UpdateEnvironmentResponse) GetFourHundredApplicationJSONObject() *UpdateEnvironmentResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *UpdateEnvironmentResponse) GetFourHundredAndOneApplicationJSONObject() *UpdateEnvironmentEnvironmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *UpdateEnvironmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateEnvironmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateEnvironmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateEnvironmentResponse) GetEnvironmentSchema() *shared.EnvironmentSchema {
	if o == nil {
		return nil
	}
	return o.EnvironmentSchema
}
