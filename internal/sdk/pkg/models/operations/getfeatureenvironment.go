// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetFeatureEnvironmentRequest struct {
	Environment string `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *GetFeatureEnvironmentRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *GetFeatureEnvironmentRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *GetFeatureEnvironmentRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// GetFeatureEnvironmentFeaturesResponseResponseBody - The requested resource was not found.
type GetFeatureEnvironmentFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureEnvironmentFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureEnvironmentFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureEnvironmentFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeatureEnvironmentFeaturesResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetFeatureEnvironmentFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureEnvironmentFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureEnvironmentFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureEnvironmentFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeatureEnvironmentResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetFeatureEnvironmentResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureEnvironmentResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureEnvironmentResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureEnvironmentResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetFeatureEnvironmentResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetFeatureEnvironmentResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *GetFeatureEnvironmentFeaturesResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *GetFeatureEnvironmentFeaturesResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureEnvironmentSchema
	FeatureEnvironmentSchema *shared.FeatureEnvironmentSchema
}

func (o *GetFeatureEnvironmentResponse) GetFourHundredAndOneApplicationJSONObject() *GetFeatureEnvironmentResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetFeatureEnvironmentResponse) GetFourHundredAndThreeApplicationJSONObject() *GetFeatureEnvironmentFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetFeatureEnvironmentResponse) GetFourHundredAndFourApplicationJSONObject() *GetFeatureEnvironmentFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *GetFeatureEnvironmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFeatureEnvironmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFeatureEnvironmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFeatureEnvironmentResponse) GetFeatureEnvironmentSchema() *shared.FeatureEnvironmentSchema {
	if o == nil {
		return nil
	}
	return o.FeatureEnvironmentSchema
}
