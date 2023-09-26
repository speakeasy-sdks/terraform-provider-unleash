// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetTagRequest struct {
	Type  string `pathParam:"style=simple,explode=false,name=type"`
	Value string `pathParam:"style=simple,explode=false,name=value"`
}

// GetTag404ApplicationJSON - The requested resource was not found.
type GetTag404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetTag403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetTag403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetTag401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetTag401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type GetTagResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetTag401ApplicationJSONObject *GetTag401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetTag403ApplicationJSONObject *GetTag403ApplicationJSON
	// The requested resource was not found.
	GetTag404ApplicationJSONObject *GetTag404ApplicationJSON
	// tagWithVersionSchema
	TagWithVersionSchema *shared.TagWithVersionSchema
}
