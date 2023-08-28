// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type AddTagRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	// tagSchema
	TagSchema shared.TagSchema `request:"mediaType=application/json"`
}

// AddTag404ApplicationJSON - The requested resource was not found.
type AddTag404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// AddTag403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type AddTag403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// AddTag401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type AddTag401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// AddTag400ApplicationJSON - The request data does not match what we expect.
type AddTag400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type AddTagResponse struct {
	ContentType string
	Headers     map[string][]string
	StatusCode  int
	RawResponse *http.Response
	// The request data does not match what we expect.
	AddTag400ApplicationJSONObject *AddTag400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	AddTag401ApplicationJSONObject *AddTag401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	AddTag403ApplicationJSONObject *AddTag403ApplicationJSON
	// The requested resource was not found.
	AddTag404ApplicationJSONObject *AddTag404ApplicationJSON
	// The resource was successfully created.
	TagSchema *shared.TagSchema
}
