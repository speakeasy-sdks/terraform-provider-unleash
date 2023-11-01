// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetEnvironmentFeatureVariantsRequest struct {
	Environment string `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *GetEnvironmentFeatureVariantsRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *GetEnvironmentFeatureVariantsRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *GetEnvironmentFeatureVariantsRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// GetEnvironmentFeatureVariants404ApplicationJSON - The requested resource was not found.
type GetEnvironmentFeatureVariants404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentFeatureVariants404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentFeatureVariants404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentFeatureVariants404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetEnvironmentFeatureVariants403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetEnvironmentFeatureVariants403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentFeatureVariants403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentFeatureVariants403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentFeatureVariants403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetEnvironmentFeatureVariants401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetEnvironmentFeatureVariants401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEnvironmentFeatureVariants401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEnvironmentFeatureVariants401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEnvironmentFeatureVariants401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetEnvironmentFeatureVariantsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetEnvironmentFeatureVariants401ApplicationJSONObject *GetEnvironmentFeatureVariants401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetEnvironmentFeatureVariants403ApplicationJSONObject *GetEnvironmentFeatureVariants403ApplicationJSON
	// The requested resource was not found.
	GetEnvironmentFeatureVariants404ApplicationJSONObject *GetEnvironmentFeatureVariants404ApplicationJSON
}

func (o *GetEnvironmentFeatureVariantsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEnvironmentFeatureVariantsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEnvironmentFeatureVariantsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEnvironmentFeatureVariantsResponse) GetFeatureVariantsSchema() *shared.FeatureVariantsSchema {
	if o == nil {
		return nil
	}
	return o.FeatureVariantsSchema
}

func (o *GetEnvironmentFeatureVariantsResponse) GetGetEnvironmentFeatureVariants401ApplicationJSONObject() *GetEnvironmentFeatureVariants401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetEnvironmentFeatureVariants401ApplicationJSONObject
}

func (o *GetEnvironmentFeatureVariantsResponse) GetGetEnvironmentFeatureVariants403ApplicationJSONObject() *GetEnvironmentFeatureVariants403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetEnvironmentFeatureVariants403ApplicationJSONObject
}

func (o *GetEnvironmentFeatureVariantsResponse) GetGetEnvironmentFeatureVariants404ApplicationJSONObject() *GetEnvironmentFeatureVariants404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetEnvironmentFeatureVariants404ApplicationJSONObject
}
