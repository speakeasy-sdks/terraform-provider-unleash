// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// SetSamlSettings415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type SetSamlSettings415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSamlSettings415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSamlSettings415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSamlSettings415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSamlSettings403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type SetSamlSettings403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSamlSettings403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSamlSettings403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSamlSettings403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSamlSettings401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type SetSamlSettings401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSamlSettings401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSamlSettings401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSamlSettings401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSamlSettings400ApplicationJSON - The request data does not match what we expect.
type SetSamlSettings400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSamlSettings400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSamlSettings400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSamlSettings400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type SetSamlSettingsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// samlSettingsSchema
	SamlSettingsSchema *shared.SamlSettingsSchema
	// The request data does not match what we expect.
	SetSamlSettings400ApplicationJSONObject *SetSamlSettings400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	SetSamlSettings401ApplicationJSONObject *SetSamlSettings401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	SetSamlSettings403ApplicationJSONObject *SetSamlSettings403ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	SetSamlSettings415ApplicationJSONObject *SetSamlSettings415ApplicationJSON
}

func (o *SetSamlSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetSamlSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetSamlSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SetSamlSettingsResponse) GetSamlSettingsSchema() *shared.SamlSettingsSchema {
	if o == nil {
		return nil
	}
	return o.SamlSettingsSchema
}

func (o *SetSamlSettingsResponse) GetSetSamlSettings400ApplicationJSONObject() *SetSamlSettings400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SetSamlSettings400ApplicationJSONObject
}

func (o *SetSamlSettingsResponse) GetSetSamlSettings401ApplicationJSONObject() *SetSamlSettings401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SetSamlSettings401ApplicationJSONObject
}

func (o *SetSamlSettingsResponse) GetSetSamlSettings403ApplicationJSONObject() *SetSamlSettings403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SetSamlSettings403ApplicationJSONObject
}

func (o *SetSamlSettingsResponse) GetSetSamlSettings415ApplicationJSONObject() *SetSamlSettings415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.SetSamlSettings415ApplicationJSONObject
}
