// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// CreateRole400ApplicationJSON - The request data does not match what we expect.
type CreateRole400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type CreateRoleResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The request data does not match what we expect.
	CreateRole400ApplicationJSONObject *CreateRole400ApplicationJSON
	// roleWithVersionSchema
	RoleWithVersionSchema *shared.RoleWithVersionSchema
}
