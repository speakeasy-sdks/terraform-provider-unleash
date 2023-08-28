// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// ValidateTagType415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type ValidateTagType415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ValidateTagType409ApplicationJSON - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type ValidateTagType409ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ValidateTagType403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ValidateTagType403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ValidateTagType401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ValidateTagType401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ValidateTagType400ApplicationJSON - The request data does not match what we expect.
type ValidateTagType400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type ValidateTagTypeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// validateTagTypeSchema
	ValidateTagTypeSchema *shared.ValidateTagTypeSchema
	// The request data does not match what we expect.
	ValidateTagType400ApplicationJSONObject *ValidateTagType400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	ValidateTagType401ApplicationJSONObject *ValidateTagType401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	ValidateTagType403ApplicationJSONObject *ValidateTagType403ApplicationJSON
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	ValidateTagType409ApplicationJSONObject *ValidateTagType409ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	ValidateTagType415ApplicationJSONObject *ValidateTagType415ApplicationJSON
}
