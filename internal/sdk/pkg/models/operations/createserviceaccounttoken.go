// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CreateServiceAccountTokenRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// patSchema
	PatSchema shared.PatSchema `request:"mediaType=application/json"`
}

func (o *CreateServiceAccountTokenRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateServiceAccountTokenRequest) GetPatSchema() shared.PatSchema {
	if o == nil {
		return shared.PatSchema{}
	}
	return o.PatSchema
}

// CreateServiceAccountToken415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type CreateServiceAccountToken415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateServiceAccountToken415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateServiceAccountToken415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateServiceAccountToken415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateServiceAccountToken409ApplicationJSON - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type CreateServiceAccountToken409ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateServiceAccountToken409ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateServiceAccountToken409ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateServiceAccountToken409ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateServiceAccountToken404ApplicationJSON - The requested resource was not found.
type CreateServiceAccountToken404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateServiceAccountToken404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateServiceAccountToken404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateServiceAccountToken404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateServiceAccountToken403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreateServiceAccountToken403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateServiceAccountToken403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateServiceAccountToken403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateServiceAccountToken403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateServiceAccountToken401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreateServiceAccountToken401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateServiceAccountToken401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateServiceAccountToken401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateServiceAccountToken401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreateServiceAccountTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	CreateServiceAccountToken401ApplicationJSONObject *CreateServiceAccountToken401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	CreateServiceAccountToken403ApplicationJSONObject *CreateServiceAccountToken403ApplicationJSON
	// The requested resource was not found.
	CreateServiceAccountToken404ApplicationJSONObject *CreateServiceAccountToken404ApplicationJSON
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	CreateServiceAccountToken409ApplicationJSONObject *CreateServiceAccountToken409ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	CreateServiceAccountToken415ApplicationJSONObject *CreateServiceAccountToken415ApplicationJSON
	// The resource was successfully created.
	PatSchema *shared.PatSchema
}

func (o *CreateServiceAccountTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateServiceAccountTokenResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateServiceAccountTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateServiceAccountTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateServiceAccountTokenResponse) GetCreateServiceAccountToken401ApplicationJSONObject() *CreateServiceAccountToken401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateServiceAccountToken401ApplicationJSONObject
}

func (o *CreateServiceAccountTokenResponse) GetCreateServiceAccountToken403ApplicationJSONObject() *CreateServiceAccountToken403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateServiceAccountToken403ApplicationJSONObject
}

func (o *CreateServiceAccountTokenResponse) GetCreateServiceAccountToken404ApplicationJSONObject() *CreateServiceAccountToken404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateServiceAccountToken404ApplicationJSONObject
}

func (o *CreateServiceAccountTokenResponse) GetCreateServiceAccountToken409ApplicationJSONObject() *CreateServiceAccountToken409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateServiceAccountToken409ApplicationJSONObject
}

func (o *CreateServiceAccountTokenResponse) GetCreateServiceAccountToken415ApplicationJSONObject() *CreateServiceAccountToken415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateServiceAccountToken415ApplicationJSONObject
}

func (o *CreateServiceAccountTokenResponse) GetPatSchema() *shared.PatSchema {
	if o == nil {
		return nil
	}
	return o.PatSchema
}
