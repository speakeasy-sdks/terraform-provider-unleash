// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetEnvironmentFeatureVariantsRequest struct {
	Environment string `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *GetEnvironmentFeatureVariantsRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *GetEnvironmentFeatureVariantsRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *GetEnvironmentFeatureVariantsRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// GetEnvironmentFeatureVariantsFeaturesResponseResponseBody - The requested resource was not found.
type GetEnvironmentFeatureVariantsFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentFeatureVariantsFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentFeatureVariantsFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentFeatureVariantsFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetEnvironmentFeatureVariantsFeaturesResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetEnvironmentFeatureVariantsFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentFeatureVariantsFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentFeatureVariantsFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentFeatureVariantsFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetEnvironmentFeatureVariantsResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetEnvironmentFeatureVariantsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentFeatureVariantsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentFeatureVariantsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentFeatureVariantsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetEnvironmentFeatureVariantsResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetEnvironmentFeatureVariantsResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *GetEnvironmentFeatureVariantsFeaturesResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *GetEnvironmentFeatureVariantsFeaturesResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
}

func (o *GetEnvironmentFeatureVariantsResponse) GetFourHundredAndOneApplicationJSONObject() *GetEnvironmentFeatureVariantsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetEnvironmentFeatureVariantsResponse) GetFourHundredAndThreeApplicationJSONObject() *GetEnvironmentFeatureVariantsFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetEnvironmentFeatureVariantsResponse) GetFourHundredAndFourApplicationJSONObject() *GetEnvironmentFeatureVariantsFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *GetEnvironmentFeatureVariantsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEnvironmentFeatureVariantsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEnvironmentFeatureVariantsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEnvironmentFeatureVariantsResponse) GetFeatureVariantsSchema() *shared.FeatureVariantsSchema {
	if o == nil {
		return nil
	}
	return o.FeatureVariantsSchema
}
