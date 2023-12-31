// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type AddEnvironmentToProjectRequest struct {
	// projectEnvironmentSchema
	ProjectEnvironmentSchema shared.ProjectEnvironmentSchema `request:"mediaType=application/json"`
	ProjectID                string                          `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *AddEnvironmentToProjectRequest) GetProjectEnvironmentSchema() shared.ProjectEnvironmentSchema {
	if o == nil {
		return shared.ProjectEnvironmentSchema{}
	}
	return o.ProjectEnvironmentSchema
}

func (o *AddEnvironmentToProjectRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// AddEnvironmentToProjectProjectsResponseResponseBody - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type AddEnvironmentToProjectProjectsResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *AddEnvironmentToProjectProjectsResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddEnvironmentToProjectProjectsResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *AddEnvironmentToProjectProjectsResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// AddEnvironmentToProjectProjectsResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type AddEnvironmentToProjectProjectsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *AddEnvironmentToProjectProjectsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddEnvironmentToProjectProjectsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *AddEnvironmentToProjectProjectsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// AddEnvironmentToProjectResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type AddEnvironmentToProjectResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *AddEnvironmentToProjectResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AddEnvironmentToProjectResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *AddEnvironmentToProjectResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type AddEnvironmentToProjectResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *AddEnvironmentToProjectResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *AddEnvironmentToProjectProjectsResponseBody
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	FourHundredAndNineApplicationJSONObject *AddEnvironmentToProjectProjectsResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddEnvironmentToProjectResponse) GetFourHundredAndOneApplicationJSONObject() *AddEnvironmentToProjectResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *AddEnvironmentToProjectResponse) GetFourHundredAndThreeApplicationJSONObject() *AddEnvironmentToProjectProjectsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *AddEnvironmentToProjectResponse) GetFourHundredAndNineApplicationJSONObject() *AddEnvironmentToProjectProjectsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndNineApplicationJSONObject
}

func (o *AddEnvironmentToProjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddEnvironmentToProjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddEnvironmentToProjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
