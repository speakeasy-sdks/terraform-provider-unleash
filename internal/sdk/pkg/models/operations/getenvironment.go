// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetEnvironmentRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

// GetEnvironment404ApplicationJSON - The requested resource was not found.
type GetEnvironment404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetEnvironment403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetEnvironment403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetEnvironment401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetEnvironment401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type GetEnvironmentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// environmentSchema
	EnvironmentSchema *shared.EnvironmentSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetEnvironment401ApplicationJSONObject *GetEnvironment401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetEnvironment403ApplicationJSONObject *GetEnvironment403ApplicationJSON
	// The requested resource was not found.
	GetEnvironment404ApplicationJSONObject *GetEnvironment404ApplicationJSON
}
