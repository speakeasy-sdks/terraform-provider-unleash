// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetProjectEnvironmentsRequest struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *GetProjectEnvironmentsRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// GetProjectEnvironmentsEnvironmentsResponseResponseBody - The requested resource was not found.
type GetProjectEnvironmentsEnvironmentsResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetProjectEnvironmentsEnvironmentsResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetProjectEnvironmentsEnvironmentsResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetProjectEnvironmentsEnvironmentsResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetProjectEnvironmentsEnvironmentsResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetProjectEnvironmentsEnvironmentsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetProjectEnvironmentsEnvironmentsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetProjectEnvironmentsEnvironmentsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetProjectEnvironmentsEnvironmentsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetProjectEnvironmentsResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetProjectEnvironmentsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetProjectEnvironmentsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetProjectEnvironmentsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetProjectEnvironmentsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetProjectEnvironmentsResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetProjectEnvironmentsResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *GetProjectEnvironmentsEnvironmentsResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *GetProjectEnvironmentsEnvironmentsResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// environmentsProjectSchema
	EnvironmentsProjectSchema *shared.EnvironmentsProjectSchema
}

func (o *GetProjectEnvironmentsResponse) GetFourHundredAndOneApplicationJSONObject() *GetProjectEnvironmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetProjectEnvironmentsResponse) GetFourHundredAndThreeApplicationJSONObject() *GetProjectEnvironmentsEnvironmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetProjectEnvironmentsResponse) GetFourHundredAndFourApplicationJSONObject() *GetProjectEnvironmentsEnvironmentsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *GetProjectEnvironmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProjectEnvironmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProjectEnvironmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetProjectEnvironmentsResponse) GetEnvironmentsProjectSchema() *shared.EnvironmentsProjectSchema {
	if o == nil {
		return nil
	}
	return o.EnvironmentsProjectSchema
}
