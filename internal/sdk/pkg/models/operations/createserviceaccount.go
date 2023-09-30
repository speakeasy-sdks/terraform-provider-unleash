// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// CreateServiceAccount415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type CreateServiceAccount415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// CreateServiceAccount409ApplicationJSON - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type CreateServiceAccount409ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// CreateServiceAccount403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreateServiceAccount403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// CreateServiceAccount401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreateServiceAccount401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// CreateServiceAccount400ApplicationJSON - The request data does not match what we expect.
type CreateServiceAccount400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type CreateServiceAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	CreateServiceAccount400ApplicationJSONObject *CreateServiceAccount400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	CreateServiceAccount401ApplicationJSONObject *CreateServiceAccount401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	CreateServiceAccount403ApplicationJSONObject *CreateServiceAccount403ApplicationJSON
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	CreateServiceAccount409ApplicationJSONObject *CreateServiceAccount409ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	CreateServiceAccount415ApplicationJSONObject *CreateServiceAccount415ApplicationJSON
	// The resource was successfully created.
	ServiceAccountSchema *shared.ServiceAccountSchema
}
