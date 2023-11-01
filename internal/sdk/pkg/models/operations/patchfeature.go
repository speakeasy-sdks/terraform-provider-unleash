// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type PatchFeatureRequest struct {
	// patchesSchema
	RequestBody []shared.PatchSchema `request:"mediaType=application/json"`
	FeatureName string               `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string               `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *PatchFeatureRequest) GetRequestBody() []shared.PatchSchema {
	if o == nil {
		return []shared.PatchSchema{}
	}
	return o.RequestBody
}

func (o *PatchFeatureRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *PatchFeatureRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// PatchFeature415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type PatchFeature415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeature415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeature415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeature415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeature404ApplicationJSON - The requested resource was not found.
type PatchFeature404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeature404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeature404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeature404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeature403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type PatchFeature403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeature403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeature403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeature403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeature401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type PatchFeature401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeature401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeature401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeature401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type PatchFeatureResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureSchema
	FeatureSchema *shared.FeatureSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	PatchFeature401ApplicationJSONObject *PatchFeature401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	PatchFeature403ApplicationJSONObject *PatchFeature403ApplicationJSON
	// The requested resource was not found.
	PatchFeature404ApplicationJSONObject *PatchFeature404ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	PatchFeature415ApplicationJSONObject *PatchFeature415ApplicationJSON
}

func (o *PatchFeatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchFeatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchFeatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchFeatureResponse) GetFeatureSchema() *shared.FeatureSchema {
	if o == nil {
		return nil
	}
	return o.FeatureSchema
}

func (o *PatchFeatureResponse) GetPatchFeature401ApplicationJSONObject() *PatchFeature401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeature401ApplicationJSONObject
}

func (o *PatchFeatureResponse) GetPatchFeature403ApplicationJSONObject() *PatchFeature403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeature403ApplicationJSONObject
}

func (o *PatchFeatureResponse) GetPatchFeature404ApplicationJSONObject() *PatchFeature404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeature404ApplicationJSONObject
}

func (o *PatchFeatureResponse) GetPatchFeature415ApplicationJSONObject() *PatchFeature415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeature415ApplicationJSONObject
}
