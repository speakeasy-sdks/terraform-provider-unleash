// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type ListTagsRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
}

// ListTags404ApplicationJSON - The requested resource was not found.
type ListTags404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ListTags403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ListTags403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

// ListTags401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ListTags401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

type ListTagsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	ListTags401ApplicationJSONObject *ListTags401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	ListTags403ApplicationJSONObject *ListTags403ApplicationJSON
	// The requested resource was not found.
	ListTags404ApplicationJSONObject *ListTags404ApplicationJSON
	// tagsSchema
	TagsSchema *shared.TagsSchema
}
