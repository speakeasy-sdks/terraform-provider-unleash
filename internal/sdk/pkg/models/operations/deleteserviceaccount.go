// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteServiceAccountRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteServiceAccountRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeleteServiceAccount404ApplicationJSON - The requested resource was not found.
type DeleteServiceAccount404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteServiceAccount404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteServiceAccount404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteServiceAccount404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteServiceAccount403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeleteServiceAccount403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteServiceAccount403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteServiceAccount403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteServiceAccount403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteServiceAccount401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeleteServiceAccount401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteServiceAccount401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteServiceAccount401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteServiceAccount401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeleteServiceAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	DeleteServiceAccount401ApplicationJSONObject *DeleteServiceAccount401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	DeleteServiceAccount403ApplicationJSONObject *DeleteServiceAccount403ApplicationJSON
	// The requested resource was not found.
	DeleteServiceAccount404ApplicationJSONObject *DeleteServiceAccount404ApplicationJSON
}

func (o *DeleteServiceAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteServiceAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteServiceAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteServiceAccountResponse) GetDeleteServiceAccount401ApplicationJSONObject() *DeleteServiceAccount401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteServiceAccount401ApplicationJSONObject
}

func (o *DeleteServiceAccountResponse) GetDeleteServiceAccount403ApplicationJSONObject() *DeleteServiceAccount403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteServiceAccount403ApplicationJSONObject
}

func (o *DeleteServiceAccountResponse) GetDeleteServiceAccount404ApplicationJSONObject() *DeleteServiceAccount404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeleteServiceAccount404ApplicationJSONObject
}
