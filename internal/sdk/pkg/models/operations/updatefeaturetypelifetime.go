// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateFeatureTypeLifetimeRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// updateFeatureTypeLifetimeSchema
	UpdateFeatureTypeLifetimeSchema shared.UpdateFeatureTypeLifetimeSchema `request:"mediaType=application/json"`
}

func (o *UpdateFeatureTypeLifetimeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetimeRequest) GetUpdateFeatureTypeLifetimeSchema() shared.UpdateFeatureTypeLifetimeSchema {
	if o == nil {
		return shared.UpdateFeatureTypeLifetimeSchema{}
	}
	return o.UpdateFeatureTypeLifetimeSchema
}

// UpdateFeatureTypeLifetime415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type UpdateFeatureTypeLifetime415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetime415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetime415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetime415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetime409ApplicationJSON - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type UpdateFeatureTypeLifetime409ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetime409ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetime409ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetime409ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetime404ApplicationJSON - The requested resource was not found.
type UpdateFeatureTypeLifetime404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetime404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetime404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetime404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetime403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type UpdateFeatureTypeLifetime403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetime403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetime403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetime403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetime401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type UpdateFeatureTypeLifetime401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetime401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetime401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetime401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetime400ApplicationJSON - The request data does not match what we expect.
type UpdateFeatureTypeLifetime400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetime400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetime400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetime400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateFeatureTypeLifetimeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureTypeSchema
	FeatureTypeSchema *shared.FeatureTypeSchema
	// The request data does not match what we expect.
	UpdateFeatureTypeLifetime400ApplicationJSONObject *UpdateFeatureTypeLifetime400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	UpdateFeatureTypeLifetime401ApplicationJSONObject *UpdateFeatureTypeLifetime401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	UpdateFeatureTypeLifetime403ApplicationJSONObject *UpdateFeatureTypeLifetime403ApplicationJSON
	// The requested resource was not found.
	UpdateFeatureTypeLifetime404ApplicationJSONObject *UpdateFeatureTypeLifetime404ApplicationJSON
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	UpdateFeatureTypeLifetime409ApplicationJSONObject *UpdateFeatureTypeLifetime409ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	UpdateFeatureTypeLifetime415ApplicationJSONObject *UpdateFeatureTypeLifetime415ApplicationJSON
}

func (o *UpdateFeatureTypeLifetimeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFeatureTypeLifetimeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFeatureTypeLifetimeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateFeatureTypeLifetimeResponse) GetFeatureTypeSchema() *shared.FeatureTypeSchema {
	if o == nil {
		return nil
	}
	return o.FeatureTypeSchema
}

func (o *UpdateFeatureTypeLifetimeResponse) GetUpdateFeatureTypeLifetime400ApplicationJSONObject() *UpdateFeatureTypeLifetime400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateFeatureTypeLifetime400ApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetUpdateFeatureTypeLifetime401ApplicationJSONObject() *UpdateFeatureTypeLifetime401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateFeatureTypeLifetime401ApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetUpdateFeatureTypeLifetime403ApplicationJSONObject() *UpdateFeatureTypeLifetime403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateFeatureTypeLifetime403ApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetUpdateFeatureTypeLifetime404ApplicationJSONObject() *UpdateFeatureTypeLifetime404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateFeatureTypeLifetime404ApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetUpdateFeatureTypeLifetime409ApplicationJSONObject() *UpdateFeatureTypeLifetime409ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateFeatureTypeLifetime409ApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetUpdateFeatureTypeLifetime415ApplicationJSONObject() *UpdateFeatureTypeLifetime415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateFeatureTypeLifetime415ApplicationJSONObject
}
