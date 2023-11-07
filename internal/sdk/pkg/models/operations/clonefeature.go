// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CloneFeatureRequest struct {
	// cloneFeatureSchema
	CloneFeatureSchema shared.CloneFeatureSchema `request:"mediaType=application/json"`
	FeatureName        string                    `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID          string                    `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *CloneFeatureRequest) GetCloneFeatureSchema() shared.CloneFeatureSchema {
	if o == nil {
		return shared.CloneFeatureSchema{}
	}
	return o.CloneFeatureSchema
}

func (o *CloneFeatureRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *CloneFeatureRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// CloneFeatureFeaturesResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type CloneFeatureFeaturesResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CloneFeatureFeaturesResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CloneFeatureFeaturesResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CloneFeatureFeaturesResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CloneFeatureFeaturesResponseResponseBody - The requested resource was not found.
type CloneFeatureFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CloneFeatureFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CloneFeatureFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CloneFeatureFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CloneFeatureFeaturesResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CloneFeatureFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CloneFeatureFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CloneFeatureFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CloneFeatureFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CloneFeatureResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CloneFeatureResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CloneFeatureResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CloneFeatureResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CloneFeatureResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CloneFeatureResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *CloneFeatureResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *CloneFeatureFeaturesResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *CloneFeatureFeaturesResponseResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *CloneFeatureFeaturesResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureSchema
	FeatureSchema *shared.FeatureSchema
}

func (o *CloneFeatureResponse) GetFourHundredAndOneApplicationJSONObject() *CloneFeatureResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *CloneFeatureResponse) GetFourHundredAndThreeApplicationJSONObject() *CloneFeatureFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *CloneFeatureResponse) GetFourHundredAndFourApplicationJSONObject() *CloneFeatureFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *CloneFeatureResponse) GetFourHundredAndFifteenApplicationJSONObject() *CloneFeatureFeaturesResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *CloneFeatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CloneFeatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CloneFeatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CloneFeatureResponse) GetFeatureSchema() *shared.FeatureSchema {
	if o == nil {
		return nil
	}
	return o.FeatureSchema
}
