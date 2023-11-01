// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetFeatureVariantsRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *GetFeatureVariantsRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *GetFeatureVariantsRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// GetFeatureVariants404ApplicationJSON - The requested resource was not found.
type GetFeatureVariants404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureVariants404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureVariants404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureVariants404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeatureVariants403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetFeatureVariants403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureVariants403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureVariants403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureVariants403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeatureVariants401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetFeatureVariants401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureVariants401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureVariants401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureVariants401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetFeatureVariantsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetFeatureVariants401ApplicationJSONObject *GetFeatureVariants401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetFeatureVariants403ApplicationJSONObject *GetFeatureVariants403ApplicationJSON
	// The requested resource was not found.
	GetFeatureVariants404ApplicationJSONObject *GetFeatureVariants404ApplicationJSON
}

func (o *GetFeatureVariantsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFeatureVariantsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFeatureVariantsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFeatureVariantsResponse) GetFeatureVariantsSchema() *shared.FeatureVariantsSchema {
	if o == nil {
		return nil
	}
	return o.FeatureVariantsSchema
}

func (o *GetFeatureVariantsResponse) GetGetFeatureVariants401ApplicationJSONObject() *GetFeatureVariants401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetFeatureVariants401ApplicationJSONObject
}

func (o *GetFeatureVariantsResponse) GetGetFeatureVariants403ApplicationJSONObject() *GetFeatureVariants403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetFeatureVariants403ApplicationJSONObject
}

func (o *GetFeatureVariantsResponse) GetGetFeatureVariants404ApplicationJSONObject() *GetFeatureVariants404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetFeatureVariants404ApplicationJSONObject
}
