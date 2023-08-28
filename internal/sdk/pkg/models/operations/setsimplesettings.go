// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// SetSimpleSettings415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type SetSimpleSettings415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// SetSimpleSettings403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type SetSimpleSettings403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// SetSimpleSettings401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type SetSimpleSettings401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// SetSimpleSettings400ApplicationJSON - The request data does not match what we expect.
type SetSimpleSettings400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type SetSimpleSettingsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// passwordAuthSchema
	PasswordAuthSchema *shared.PasswordAuthSchema
	// The request data does not match what we expect.
	SetSimpleSettings400ApplicationJSONObject *SetSimpleSettings400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	SetSimpleSettings401ApplicationJSONObject *SetSimpleSettings401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	SetSimpleSettings403ApplicationJSONObject *SetSimpleSettings403ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	SetSimpleSettings415ApplicationJSONObject *SetSimpleSettings415ApplicationJSON
}
