// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateFeatureRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
	// updateFeatureSchema
	UpdateFeatureSchema shared.UpdateFeatureSchema `request:"mediaType=application/json"`
}

// UpdateFeature415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type UpdateFeature415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// UpdateFeature404ApplicationJSON - The requested resource was not found.
type UpdateFeature404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// UpdateFeature403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type UpdateFeature403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// UpdateFeature401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type UpdateFeature401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type UpdateFeatureResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureSchema
	FeatureSchema *shared.FeatureSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	UpdateFeature401ApplicationJSONObject *UpdateFeature401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	UpdateFeature403ApplicationJSONObject *UpdateFeature403ApplicationJSON
	// The requested resource was not found.
	UpdateFeature404ApplicationJSONObject *UpdateFeature404ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	UpdateFeature415ApplicationJSONObject *UpdateFeature415ApplicationJSON
}
