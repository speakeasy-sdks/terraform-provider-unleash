// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// CreatePat404ApplicationJSON - The requested resource was not found.
type CreatePat404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// CreatePat403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreatePat403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// CreatePat401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreatePat401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type CreatePatResponse struct {
	ContentType string
	Headers     map[string][]string
	StatusCode  int
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	CreatePat401ApplicationJSONObject *CreatePat401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	CreatePat403ApplicationJSONObject *CreatePat403ApplicationJSON
	// The requested resource was not found.
	CreatePat404ApplicationJSONObject *CreatePat404ApplicationJSON
	// The resource was successfully created.
	PatSchema *shared.PatSchema
}
