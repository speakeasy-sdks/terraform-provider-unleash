// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type CreateFeatureRequest struct {
	// createFeatureSchema
	CreateFeatureSchema shared.CreateFeatureSchema `request:"mediaType=application/json"`
	ProjectID           string                     `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *CreateFeatureRequest) GetCreateFeatureSchema() shared.CreateFeatureSchema {
	if o == nil {
		return shared.CreateFeatureSchema{}
	}
	return o.CreateFeatureSchema
}

func (o *CreateFeatureRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// CreateFeature415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type CreateFeature415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateFeature415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateFeature415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateFeature415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateFeature404ApplicationJSON - The requested resource was not found.
type CreateFeature404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateFeature404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateFeature404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateFeature404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateFeature403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type CreateFeature403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateFeature403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateFeature403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateFeature403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// CreateFeature401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type CreateFeature401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *CreateFeature401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateFeature401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateFeature401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreateFeatureResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	CreateFeature401ApplicationJSONObject *CreateFeature401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	CreateFeature403ApplicationJSONObject *CreateFeature403ApplicationJSON
	// The requested resource was not found.
	CreateFeature404ApplicationJSONObject *CreateFeature404ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	CreateFeature415ApplicationJSONObject *CreateFeature415ApplicationJSON
	// featureSchema
	FeatureSchema *shared.FeatureSchema
}

func (o *CreateFeatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateFeatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateFeatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateFeatureResponse) GetCreateFeature401ApplicationJSONObject() *CreateFeature401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateFeature401ApplicationJSONObject
}

func (o *CreateFeatureResponse) GetCreateFeature403ApplicationJSONObject() *CreateFeature403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateFeature403ApplicationJSONObject
}

func (o *CreateFeatureResponse) GetCreateFeature404ApplicationJSONObject() *CreateFeature404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateFeature404ApplicationJSONObject
}

func (o *CreateFeatureResponse) GetCreateFeature415ApplicationJSONObject() *CreateFeature415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateFeature415ApplicationJSONObject
}

func (o *CreateFeatureResponse) GetFeatureSchema() *shared.FeatureSchema {
	if o == nil {
		return nil
	}
	return o.FeatureSchema
}
