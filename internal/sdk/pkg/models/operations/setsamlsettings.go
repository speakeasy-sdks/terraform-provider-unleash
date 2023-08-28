// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// SetSamlSettings415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type SetSamlSettings415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// SetSamlSettings403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type SetSamlSettings403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// SetSamlSettings401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type SetSamlSettings401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// SetSamlSettings400ApplicationJSON - The request data does not match what we expect.
type SetSamlSettings400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type SetSamlSettingsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// samlSettingsSchema
	SamlSettingsSchema *shared.SamlSettingsSchema
	// The request data does not match what we expect.
	SetSamlSettings400ApplicationJSONObject *SetSamlSettings400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	SetSamlSettings401ApplicationJSONObject *SetSamlSettings401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	SetSamlSettings403ApplicationJSONObject *SetSamlSettings403ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	SetSamlSettings415ApplicationJSONObject *SetSamlSettings415ApplicationJSON
}
