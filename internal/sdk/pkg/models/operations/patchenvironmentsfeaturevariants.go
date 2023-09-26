// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type PatchEnvironmentsFeatureVariantsRequest struct {
	// patchesSchema
	RequestBody []shared.PatchSchema `request:"mediaType=application/json"`
	Environment string               `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string               `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string               `pathParam:"style=simple,explode=false,name=projectId"`
}

// PatchEnvironmentsFeatureVariants404ApplicationJSON - The requested resource was not found.
type PatchEnvironmentsFeatureVariants404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// PatchEnvironmentsFeatureVariants403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type PatchEnvironmentsFeatureVariants403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// PatchEnvironmentsFeatureVariants401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type PatchEnvironmentsFeatureVariants401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// PatchEnvironmentsFeatureVariants400ApplicationJSON - The request data does not match what we expect.
type PatchEnvironmentsFeatureVariants400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type PatchEnvironmentsFeatureVariantsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
	// The request data does not match what we expect.
	PatchEnvironmentsFeatureVariants400ApplicationJSONObject *PatchEnvironmentsFeatureVariants400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	PatchEnvironmentsFeatureVariants401ApplicationJSONObject *PatchEnvironmentsFeatureVariants401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	PatchEnvironmentsFeatureVariants403ApplicationJSONObject *PatchEnvironmentsFeatureVariants403ApplicationJSON
	// The requested resource was not found.
	PatchEnvironmentsFeatureVariants404ApplicationJSONObject *PatchEnvironmentsFeatureVariants404ApplicationJSON
}
