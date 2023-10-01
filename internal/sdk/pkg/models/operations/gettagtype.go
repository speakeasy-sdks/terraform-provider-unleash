// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetTagTypeRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

// GetTagType403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetTagType403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetTagType401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetTagType401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type GetTagTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetTagType401ApplicationJSONObject *GetTagType401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetTagType403ApplicationJSONObject *GetTagType403ApplicationJSON
	// tagTypeSchema
	TagTypeSchema *shared.TagTypeSchema
}
