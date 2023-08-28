// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeprecateStrategyRequest struct {
	StrategyName string `pathParam:"style=simple,explode=false,name=strategyName"`
}

// DeprecateStrategy404ApplicationJSON - The requested resource was not found.
type DeprecateStrategy404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// DeprecateStrategy403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeprecateStrategy403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// DeprecateStrategy401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeprecateStrategy401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type DeprecateStrategyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	DeprecateStrategy401ApplicationJSONObject *DeprecateStrategy401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	DeprecateStrategy403ApplicationJSONObject *DeprecateStrategy403ApplicationJSON
	// The requested resource was not found.
	DeprecateStrategy404ApplicationJSONObject *DeprecateStrategy404ApplicationJSON
}
