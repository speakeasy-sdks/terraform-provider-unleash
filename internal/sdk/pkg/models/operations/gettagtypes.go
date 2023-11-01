// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// GetTagTypes403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetTagTypes403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetTagTypes403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetTagTypes403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetTagTypes403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetTagTypes401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetTagTypes401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetTagTypes401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetTagTypes401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetTagTypes401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetTagTypesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetTagTypes401ApplicationJSONObject *GetTagTypes401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetTagTypes403ApplicationJSONObject *GetTagTypes403ApplicationJSON
	// tagTypesSchema
	TagTypesSchema *shared.TagTypesSchema
}

func (o *GetTagTypesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTagTypesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTagTypesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTagTypesResponse) GetGetTagTypes401ApplicationJSONObject() *GetTagTypes401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTagTypes401ApplicationJSONObject
}

func (o *GetTagTypesResponse) GetGetTagTypes403ApplicationJSONObject() *GetTagTypes403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTagTypes403ApplicationJSONObject
}

func (o *GetTagTypesResponse) GetTagTypesSchema() *shared.TagTypesSchema {
	if o == nil {
		return nil
	}
	return o.TagTypesSchema
}
