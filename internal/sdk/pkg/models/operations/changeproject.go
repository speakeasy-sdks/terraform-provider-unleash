// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ChangeProjectRequest struct {
	// changeProjectSchema
	ChangeProjectSchema shared.ChangeProjectSchema `request:"mediaType=application/json"`
	FeatureName         string                     `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID           string                     `pathParam:"style=simple,explode=false,name=projectId"`
}

// ChangeProject415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type ChangeProject415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ChangeProject404ApplicationJSON - The requested resource was not found.
type ChangeProject404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ChangeProject403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ChangeProject403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ChangeProject401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ChangeProject401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ChangeProject400ApplicationJSON - The request data does not match what we expect.
type ChangeProject400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type ChangeProjectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The request data does not match what we expect.
	ChangeProject400ApplicationJSONObject *ChangeProject400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	ChangeProject401ApplicationJSONObject *ChangeProject401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	ChangeProject403ApplicationJSONObject *ChangeProject403ApplicationJSON
	// The requested resource was not found.
	ChangeProject404ApplicationJSONObject *ChangeProject404ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	ChangeProject415ApplicationJSONObject *ChangeProject415ApplicationJSON
}
