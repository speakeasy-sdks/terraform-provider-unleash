// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type OverwriteEnvironmentFeatureVariantsRequest struct {
	// variantsSchema
	RequestBody []shared.VariantSchema `request:"mediaType=application/json"`
	Environment string                 `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string                 `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string                 `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *OverwriteEnvironmentFeatureVariantsRequest) GetRequestBody() []shared.VariantSchema {
	if o == nil {
		return []shared.VariantSchema{}
	}
	return o.RequestBody
}

func (o *OverwriteEnvironmentFeatureVariantsRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *OverwriteEnvironmentFeatureVariantsRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *OverwriteEnvironmentFeatureVariantsRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// OverwriteEnvironmentFeatureVariants403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type OverwriteEnvironmentFeatureVariants403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *OverwriteEnvironmentFeatureVariants403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OverwriteEnvironmentFeatureVariants403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *OverwriteEnvironmentFeatureVariants403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// OverwriteEnvironmentFeatureVariants401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type OverwriteEnvironmentFeatureVariants401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *OverwriteEnvironmentFeatureVariants401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OverwriteEnvironmentFeatureVariants401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *OverwriteEnvironmentFeatureVariants401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// OverwriteEnvironmentFeatureVariants400ApplicationJSON - The request data does not match what we expect.
type OverwriteEnvironmentFeatureVariants400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *OverwriteEnvironmentFeatureVariants400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *OverwriteEnvironmentFeatureVariants400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *OverwriteEnvironmentFeatureVariants400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type OverwriteEnvironmentFeatureVariantsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureVariantsSchema
	FeatureVariantsSchema *shared.FeatureVariantsSchema
	// The request data does not match what we expect.
	OverwriteEnvironmentFeatureVariants400ApplicationJSONObject *OverwriteEnvironmentFeatureVariants400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	OverwriteEnvironmentFeatureVariants401ApplicationJSONObject *OverwriteEnvironmentFeatureVariants401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	OverwriteEnvironmentFeatureVariants403ApplicationJSONObject *OverwriteEnvironmentFeatureVariants403ApplicationJSON
}

func (o *OverwriteEnvironmentFeatureVariantsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *OverwriteEnvironmentFeatureVariantsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *OverwriteEnvironmentFeatureVariantsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *OverwriteEnvironmentFeatureVariantsResponse) GetFeatureVariantsSchema() *shared.FeatureVariantsSchema {
	if o == nil {
		return nil
	}
	return o.FeatureVariantsSchema
}

func (o *OverwriteEnvironmentFeatureVariantsResponse) GetOverwriteEnvironmentFeatureVariants400ApplicationJSONObject() *OverwriteEnvironmentFeatureVariants400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.OverwriteEnvironmentFeatureVariants400ApplicationJSONObject
}

func (o *OverwriteEnvironmentFeatureVariantsResponse) GetOverwriteEnvironmentFeatureVariants401ApplicationJSONObject() *OverwriteEnvironmentFeatureVariants401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.OverwriteEnvironmentFeatureVariants401ApplicationJSONObject
}

func (o *OverwriteEnvironmentFeatureVariantsResponse) GetOverwriteEnvironmentFeatureVariants403ApplicationJSONObject() *OverwriteEnvironmentFeatureVariants403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.OverwriteEnvironmentFeatureVariants403ApplicationJSONObject
}
