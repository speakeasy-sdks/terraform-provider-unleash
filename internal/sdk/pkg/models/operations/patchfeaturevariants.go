// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type PatchFeatureVariantsRequest struct {
	// patchesSchema
	RequestBody []shared.PatchSchema `request:"mediaType=application/json"`
	FeatureName string               `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string               `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *PatchFeatureVariantsRequest) GetRequestBody() []shared.PatchSchema {
	if o == nil {
		return []shared.PatchSchema{}
	}
	return o.RequestBody
}

func (o *PatchFeatureVariantsRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *PatchFeatureVariantsRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// PatchFeatureVariantsFeaturesResponse404ResponseBody - The requested resource was not found.
type PatchFeatureVariantsFeaturesResponse404ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureVariantsFeaturesResponse404ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureVariantsFeaturesResponse404ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureVariantsFeaturesResponse404ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureVariantsFeaturesResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type PatchFeatureVariantsFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureVariantsFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureVariantsFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureVariantsFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureVariantsFeaturesResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type PatchFeatureVariantsFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureVariantsFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureVariantsFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureVariantsFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureVariantsResponseBody - The request data does not match what we expect.
type PatchFeatureVariantsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureVariantsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureVariantsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureVariantsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type PatchFeatureVariantsResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *PatchFeatureVariantsResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *PatchFeatureVariantsFeaturesResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *PatchFeatureVariantsFeaturesResponseResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *PatchFeatureVariantsFeaturesResponse404ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
}

func (o *PatchFeatureVariantsResponse) GetFourHundredApplicationJSONObject() *PatchFeatureVariantsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *PatchFeatureVariantsResponse) GetFourHundredAndOneApplicationJSONObject() *PatchFeatureVariantsFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *PatchFeatureVariantsResponse) GetFourHundredAndThreeApplicationJSONObject() *PatchFeatureVariantsFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *PatchFeatureVariantsResponse) GetFourHundredAndFourApplicationJSONObject() *PatchFeatureVariantsFeaturesResponse404ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *PatchFeatureVariantsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchFeatureVariantsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchFeatureVariantsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchFeatureVariantsResponse) GetFeatureVariantsSchema() *shared.FeatureVariantsSchema {
	if o == nil {
		return nil
	}
	return o.FeatureVariantsSchema
}
