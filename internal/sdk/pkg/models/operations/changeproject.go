// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type ChangeProjectRequest struct {
	// changeProjectSchema
	ChangeProjectSchema shared.ChangeProjectSchema `request:"mediaType=application/json"`
	FeatureName         string                     `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID           string                     `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *ChangeProjectRequest) GetChangeProjectSchema() shared.ChangeProjectSchema {
	if o == nil {
		return shared.ChangeProjectSchema{}
	}
	return o.ChangeProjectSchema
}

func (o *ChangeProjectRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *ChangeProjectRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// ChangeProjectFeaturesResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type ChangeProjectFeaturesResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ChangeProjectFeaturesResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeProjectFeaturesResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ChangeProjectFeaturesResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ChangeProjectFeaturesResponse404ResponseBody - The requested resource was not found.
type ChangeProjectFeaturesResponse404ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ChangeProjectFeaturesResponse404ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeProjectFeaturesResponse404ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ChangeProjectFeaturesResponse404ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ChangeProjectFeaturesResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ChangeProjectFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ChangeProjectFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeProjectFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ChangeProjectFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ChangeProjectFeaturesResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ChangeProjectFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ChangeProjectFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeProjectFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ChangeProjectFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ChangeProjectResponseBody - The request data does not match what we expect.
type ChangeProjectResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ChangeProjectResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ChangeProjectResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ChangeProjectResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ChangeProjectResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *ChangeProjectResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *ChangeProjectFeaturesResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *ChangeProjectFeaturesResponseResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *ChangeProjectFeaturesResponse404ResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *ChangeProjectFeaturesResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ChangeProjectResponse) GetFourHundredApplicationJSONObject() *ChangeProjectResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *ChangeProjectResponse) GetFourHundredAndOneApplicationJSONObject() *ChangeProjectFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *ChangeProjectResponse) GetFourHundredAndThreeApplicationJSONObject() *ChangeProjectFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ChangeProjectResponse) GetFourHundredAndFourApplicationJSONObject() *ChangeProjectFeaturesResponse404ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *ChangeProjectResponse) GetFourHundredAndFifteenApplicationJSONObject() *ChangeProjectFeaturesResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *ChangeProjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ChangeProjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ChangeProjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
