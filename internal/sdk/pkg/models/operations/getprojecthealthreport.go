// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetProjectHealthReportRequest struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
}

// GetProjectHealthReport404ApplicationJSON - The requested resource was not found.
type GetProjectHealthReport404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetProjectHealthReport403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetProjectHealthReport403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// GetProjectHealthReport401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetProjectHealthReport401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type GetProjectHealthReportResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetProjectHealthReport401ApplicationJSONObject *GetProjectHealthReport401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetProjectHealthReport403ApplicationJSONObject *GetProjectHealthReport403ApplicationJSON
	// The requested resource was not found.
	GetProjectHealthReport404ApplicationJSONObject *GetProjectHealthReport404ApplicationJSON
	// healthReportSchema
	HealthReportSchema *shared.HealthReportSchema
}
