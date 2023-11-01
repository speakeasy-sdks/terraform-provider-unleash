// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// ValidateUserPassword415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type ValidateUserPassword415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateUserPassword415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateUserPassword415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateUserPassword415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ValidateUserPassword401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ValidateUserPassword401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateUserPassword401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateUserPassword401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateUserPassword401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ValidateUserPassword400ApplicationJSON - The request data does not match what we expect.
type ValidateUserPassword400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateUserPassword400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateUserPassword400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateUserPassword400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ValidateUserPasswordResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	ValidateUserPassword400ApplicationJSONObject *ValidateUserPassword400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	ValidateUserPassword401ApplicationJSONObject *ValidateUserPassword401ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	ValidateUserPassword415ApplicationJSONObject *ValidateUserPassword415ApplicationJSON
}

func (o *ValidateUserPasswordResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidateUserPasswordResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidateUserPasswordResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidateUserPasswordResponse) GetValidateUserPassword400ApplicationJSONObject() *ValidateUserPassword400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ValidateUserPassword400ApplicationJSONObject
}

func (o *ValidateUserPasswordResponse) GetValidateUserPassword401ApplicationJSONObject() *ValidateUserPassword401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ValidateUserPassword401ApplicationJSONObject
}

func (o *ValidateUserPasswordResponse) GetValidateUserPassword415ApplicationJSONObject() *ValidateUserPassword415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ValidateUserPassword415ApplicationJSONObject
}
