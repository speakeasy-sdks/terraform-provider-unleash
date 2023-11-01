// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// SetSimpleSettings415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type SetSimpleSettings415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSimpleSettings415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSimpleSettings415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSimpleSettings415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSimpleSettings403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type SetSimpleSettings403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSimpleSettings403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSimpleSettings403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSimpleSettings403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSimpleSettings401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type SetSimpleSettings401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSimpleSettings401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSimpleSettings401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSimpleSettings401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSimpleSettings400ApplicationJSON - The request data does not match what we expect.
type SetSimpleSettings400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSimpleSettings400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSimpleSettings400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSimpleSettings400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type SetSimpleSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// passwordAuthSchema
	PasswordAuthSchema *shared.PasswordAuthSchema
	// The request data does not match what we expect.
	SetSimpleSettings400ApplicationJSONObject *SetSimpleSettings400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	SetSimpleSettings401ApplicationJSONObject *SetSimpleSettings401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	SetSimpleSettings403ApplicationJSONObject *SetSimpleSettings403ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	SetSimpleSettings415ApplicationJSONObject *SetSimpleSettings415ApplicationJSON
}

func (o *SetSimpleSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetSimpleSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetSimpleSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SetSimpleSettingsResponse) GetPasswordAuthSchema() *shared.PasswordAuthSchema {
	if o == nil {
		return nil
	}
	return o.PasswordAuthSchema
}

func (o *SetSimpleSettingsResponse) GetSetSimpleSettings400ApplicationJSONObject() *SetSimpleSettings400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SetSimpleSettings400ApplicationJSONObject
}

func (o *SetSimpleSettingsResponse) GetSetSimpleSettings401ApplicationJSONObject() *SetSimpleSettings401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SetSimpleSettings401ApplicationJSONObject
}

func (o *SetSimpleSettingsResponse) GetSetSimpleSettings403ApplicationJSONObject() *SetSimpleSettings403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SetSimpleSettings403ApplicationJSONObject
}

func (o *SetSimpleSettingsResponse) GetSetSimpleSettings415ApplicationJSONObject() *SetSimpleSettings415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SetSimpleSettings415ApplicationJSONObject
}
