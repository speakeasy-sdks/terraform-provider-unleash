// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetFeatureVariantsRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

// GetFeatureVariants404ApplicationJSON - The requested resource was not found.
type GetFeatureVariants404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetFeatureVariants403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetFeatureVariants403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetFeatureVariants401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetFeatureVariants401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type GetFeatureVariantsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetFeatureVariants401ApplicationJSONObject *GetFeatureVariants401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetFeatureVariants403ApplicationJSONObject *GetFeatureVariants403ApplicationJSON
	// The requested resource was not found.
	GetFeatureVariants404ApplicationJSONObject *GetFeatureVariants404ApplicationJSON
}
