// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetFeaturesRequest struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *GetFeaturesRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// GetFeaturesFeaturesResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetFeaturesFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeaturesFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeaturesFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeaturesFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeaturesFeaturesResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetFeaturesFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeaturesFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeaturesFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeaturesFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeaturesResponseBody - The request data does not match what we expect.
type GetFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetFeaturesResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *GetFeaturesResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetFeaturesFeaturesResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *GetFeaturesFeaturesResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featuresSchema
	FeaturesSchema *shared.FeaturesSchema
}

func (o *GetFeaturesResponse) GetFourHundredApplicationJSONObject() *GetFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *GetFeaturesResponse) GetFourHundredAndOneApplicationJSONObject() *GetFeaturesFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetFeaturesResponse) GetFourHundredAndThreeApplicationJSONObject() *GetFeaturesFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetFeaturesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFeaturesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFeaturesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFeaturesResponse) GetFeaturesSchema() *shared.FeaturesSchema {
	if o == nil {
		return nil
	}
	return o.FeaturesSchema
}
