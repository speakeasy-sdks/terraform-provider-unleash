// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// GetAddons401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetAddons401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetAddons401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetAddons401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetAddons401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetAddonsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// addonsSchema
	AddonsSchema *shared.AddonsSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetAddons401ApplicationJSONObject *GetAddons401ApplicationJSON
}

func (o *GetAddonsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAddonsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAddonsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAddonsResponse) GetAddonsSchema() *shared.AddonsSchema {
	if o == nil {
		return nil
	}
	return o.AddonsSchema
}

func (o *GetAddonsResponse) GetGetAddons401ApplicationJSONObject() *GetAddons401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetAddons401ApplicationJSONObject
}
