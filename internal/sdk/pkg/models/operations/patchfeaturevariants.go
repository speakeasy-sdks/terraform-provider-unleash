// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type PatchFeatureVariantsRequest struct {
	// patchesSchema
	RequestBody []shared.PatchSchema `request:"mediaType=application/json"`
	FeatureName string               `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string               `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *PatchFeatureVariantsRequest) GetRequestBody() []shared.PatchSchema {
	if o == nil {
		return []shared.PatchSchema{}
	}
	return o.RequestBody
}

func (o *PatchFeatureVariantsRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *PatchFeatureVariantsRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// PatchFeatureVariants404ApplicationJSON - The requested resource was not found.
type PatchFeatureVariants404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureVariants404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureVariants404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureVariants404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureVariants403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type PatchFeatureVariants403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureVariants403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureVariants403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureVariants403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureVariants401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type PatchFeatureVariants401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureVariants401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureVariants401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureVariants401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureVariants400ApplicationJSON - The request data does not match what we expect.
type PatchFeatureVariants400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureVariants400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureVariants400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureVariants400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type PatchFeatureVariantsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
	// The request data does not match what we expect.
	PatchFeatureVariants400ApplicationJSONObject *PatchFeatureVariants400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	PatchFeatureVariants401ApplicationJSONObject *PatchFeatureVariants401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	PatchFeatureVariants403ApplicationJSONObject *PatchFeatureVariants403ApplicationJSON
	// The requested resource was not found.
	PatchFeatureVariants404ApplicationJSONObject *PatchFeatureVariants404ApplicationJSON
}

func (o *PatchFeatureVariantsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchFeatureVariantsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchFeatureVariantsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchFeatureVariantsResponse) GetFeatureVariantsSchema() *shared.FeatureVariantsSchema {
	if o == nil {
		return nil
	}
	return o.FeatureVariantsSchema
}

func (o *PatchFeatureVariantsResponse) GetPatchFeatureVariants400ApplicationJSONObject() *PatchFeatureVariants400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureVariants400ApplicationJSONObject
}

func (o *PatchFeatureVariantsResponse) GetPatchFeatureVariants401ApplicationJSONObject() *PatchFeatureVariants401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureVariants401ApplicationJSONObject
}

func (o *PatchFeatureVariantsResponse) GetPatchFeatureVariants403ApplicationJSONObject() *PatchFeatureVariants403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureVariants403ApplicationJSONObject
}

func (o *PatchFeatureVariantsResponse) GetPatchFeatureVariants404ApplicationJSONObject() *PatchFeatureVariants404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureVariants404ApplicationJSONObject
}
