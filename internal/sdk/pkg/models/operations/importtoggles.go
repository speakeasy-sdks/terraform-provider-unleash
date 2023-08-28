// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// ImportToggles404ApplicationJSON - The requested resource was not found.
type ImportToggles404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type ImportTogglesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The requested resource was not found.
	ImportToggles404ApplicationJSONObject *ImportToggles404ApplicationJSON
}
