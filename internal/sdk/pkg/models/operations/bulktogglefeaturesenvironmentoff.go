// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type BulkToggleFeaturesEnvironmentOffRequest struct {
	// bulkToggleFeaturesSchema
	BulkToggleFeaturesSchema shared.BulkToggleFeaturesSchema `request:"mediaType=application/json"`
	Environment              string                          `pathParam:"style=simple,explode=false,name=environment"`
	ProjectID                string                          `pathParam:"style=simple,explode=false,name=projectId"`
}

// BulkToggleFeaturesEnvironmentOff415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type BulkToggleFeaturesEnvironmentOff415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// BulkToggleFeaturesEnvironmentOff413ApplicationJSON - The request body is larger than what we accept. By default we only accept bodies of 100kB or less
type BulkToggleFeaturesEnvironmentOff413ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// BulkToggleFeaturesEnvironmentOff404ApplicationJSON - The requested resource was not found.
type BulkToggleFeaturesEnvironmentOff404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// BulkToggleFeaturesEnvironmentOff403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type BulkToggleFeaturesEnvironmentOff403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// BulkToggleFeaturesEnvironmentOff401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type BulkToggleFeaturesEnvironmentOff401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// BulkToggleFeaturesEnvironmentOff400ApplicationJSON - The request data does not match what we expect.
type BulkToggleFeaturesEnvironmentOff400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type BulkToggleFeaturesEnvironmentOffResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The request data does not match what we expect.
	BulkToggleFeaturesEnvironmentOff400ApplicationJSONObject *BulkToggleFeaturesEnvironmentOff400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	BulkToggleFeaturesEnvironmentOff401ApplicationJSONObject *BulkToggleFeaturesEnvironmentOff401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	BulkToggleFeaturesEnvironmentOff403ApplicationJSONObject *BulkToggleFeaturesEnvironmentOff403ApplicationJSON
	// The requested resource was not found.
	BulkToggleFeaturesEnvironmentOff404ApplicationJSONObject *BulkToggleFeaturesEnvironmentOff404ApplicationJSON
	// The request body is larger than what we accept. By default we only accept bodies of 100kB or less
	BulkToggleFeaturesEnvironmentOff413ApplicationJSONObject *BulkToggleFeaturesEnvironmentOff413ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	BulkToggleFeaturesEnvironmentOff415ApplicationJSONObject *BulkToggleFeaturesEnvironmentOff415ApplicationJSON
}
