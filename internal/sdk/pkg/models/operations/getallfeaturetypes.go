// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// GetAllFeatureTypesResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetAllFeatureTypesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetAllFeatureTypesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetAllFeatureTypesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetAllFeatureTypesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetAllFeatureTypesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureTypesSchema
	FeatureTypesSchema *shared.FeatureTypesSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Object *GetAllFeatureTypesResponseBody
}

func (o *GetAllFeatureTypesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllFeatureTypesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllFeatureTypesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllFeatureTypesResponse) GetFeatureTypesSchema() *shared.FeatureTypesSchema {
	if o == nil {
		return nil
	}
	return o.FeatureTypesSchema
}

func (o *GetAllFeatureTypesResponse) GetObject() *GetAllFeatureTypesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
