// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveFavoriteProjectRequest struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
}

// RemoveFavoriteProject404ApplicationJSON - The requested resource was not found.
type RemoveFavoriteProject404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// RemoveFavoriteProject401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type RemoveFavoriteProject401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type RemoveFavoriteProjectResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	RemoveFavoriteProject401ApplicationJSONObject *RemoveFavoriteProject401ApplicationJSON
	// The requested resource was not found.
	RemoveFavoriteProject404ApplicationJSONObject *RemoveFavoriteProject404ApplicationJSON
}
