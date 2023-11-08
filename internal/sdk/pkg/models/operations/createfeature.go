// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateFeatureRequest struct {
	// createFeatureSchema
	CreateFeatureSchema shared.CreateFeatureSchema `request:"mediaType=application/json"`
	ProjectID           string                     `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *CreateFeatureRequest) GetCreateFeatureSchema() shared.CreateFeatureSchema {
	if o == nil {
		return shared.CreateFeatureSchema{}
	}
	return o.CreateFeatureSchema
}

func (o *CreateFeatureRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// CreateFeatureFeaturesResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type CreateFeatureFeaturesResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateFeatureFeaturesResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateFeatureFeaturesResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateFeatureFeaturesResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateFeatureFeaturesResponseResponseBody - The requested resource was not found.
type CreateFeatureFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateFeatureFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateFeatureFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateFeatureFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateFeatureFeaturesResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreateFeatureFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateFeatureFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateFeatureFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateFeatureFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateFeatureResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreateFeatureResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateFeatureResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateFeatureResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateFeatureResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreateFeatureResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *CreateFeatureResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *CreateFeatureFeaturesResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *CreateFeatureFeaturesResponseResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *CreateFeatureFeaturesResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureSchema
	FeatureSchema *shared.FeatureSchema
}

func (o *CreateFeatureResponse) GetFourHundredAndOneApplicationJSONObject() *CreateFeatureResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *CreateFeatureResponse) GetFourHundredAndThreeApplicationJSONObject() *CreateFeatureFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *CreateFeatureResponse) GetFourHundredAndFourApplicationJSONObject() *CreateFeatureFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *CreateFeatureResponse) GetFourHundredAndFifteenApplicationJSONObject() *CreateFeatureFeaturesResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *CreateFeatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateFeatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateFeatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateFeatureResponse) GetFeatureSchema() *shared.FeatureSchema {
	if o == nil {
		return nil
	}
	return o.FeatureSchema
}
