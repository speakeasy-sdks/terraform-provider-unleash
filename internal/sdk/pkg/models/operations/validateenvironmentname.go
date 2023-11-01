// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// ValidateEnvironmentName401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ValidateEnvironmentName401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateEnvironmentName401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateEnvironmentName401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateEnvironmentName401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ValidateEnvironmentName400ApplicationJSON - The request data does not match what we expect.
type ValidateEnvironmentName400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateEnvironmentName400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateEnvironmentName400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateEnvironmentName400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ValidateEnvironmentNameResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	ValidateEnvironmentName400ApplicationJSONObject *ValidateEnvironmentName400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	ValidateEnvironmentName401ApplicationJSONObject *ValidateEnvironmentName401ApplicationJSON
}

func (o *ValidateEnvironmentNameResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidateEnvironmentNameResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidateEnvironmentNameResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidateEnvironmentNameResponse) GetValidateEnvironmentName400ApplicationJSONObject() *ValidateEnvironmentName400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ValidateEnvironmentName400ApplicationJSONObject
}

func (o *ValidateEnvironmentNameResponse) GetValidateEnvironmentName401ApplicationJSONObject() *ValidateEnvironmentName401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ValidateEnvironmentName401ApplicationJSONObject
}
