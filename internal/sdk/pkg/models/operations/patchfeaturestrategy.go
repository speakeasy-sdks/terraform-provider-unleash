// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type PatchFeatureStrategyRequest struct {
	// patchesSchema
	RequestBody []shared.PatchSchema `request:"mediaType=application/json"`
	Environment string               `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string               `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string               `pathParam:"style=simple,explode=false,name=projectId"`
	StrategyID  string               `pathParam:"style=simple,explode=false,name=strategyId"`
}

func (o *PatchFeatureStrategyRequest) GetRequestBody() []shared.PatchSchema {
	if o == nil {
		return []shared.PatchSchema{}
	}
	return o.RequestBody
}

func (o *PatchFeatureStrategyRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *PatchFeatureStrategyRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *PatchFeatureStrategyRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *PatchFeatureStrategyRequest) GetStrategyID() string {
	if o == nil {
		return ""
	}
	return o.StrategyID
}

// PatchFeatureStrategy415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type PatchFeatureStrategy415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureStrategy415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureStrategy415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureStrategy415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureStrategy404ApplicationJSON - The requested resource was not found.
type PatchFeatureStrategy404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureStrategy404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureStrategy404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureStrategy404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureStrategy403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type PatchFeatureStrategy403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureStrategy403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureStrategy403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureStrategy403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureStrategy401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type PatchFeatureStrategy401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureStrategy401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureStrategy401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureStrategy401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// PatchFeatureStrategy400ApplicationJSON - The request data does not match what we expect.
type PatchFeatureStrategy400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *PatchFeatureStrategy400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PatchFeatureStrategy400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *PatchFeatureStrategy400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type PatchFeatureStrategyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureStrategySchema
	FeatureStrategySchema *shared.FeatureStrategySchema
	// The request data does not match what we expect.
	PatchFeatureStrategy400ApplicationJSONObject *PatchFeatureStrategy400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	PatchFeatureStrategy401ApplicationJSONObject *PatchFeatureStrategy401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	PatchFeatureStrategy403ApplicationJSONObject *PatchFeatureStrategy403ApplicationJSON
	// The requested resource was not found.
	PatchFeatureStrategy404ApplicationJSONObject *PatchFeatureStrategy404ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	PatchFeatureStrategy415ApplicationJSONObject *PatchFeatureStrategy415ApplicationJSON
}

func (o *PatchFeatureStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchFeatureStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchFeatureStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchFeatureStrategyResponse) GetFeatureStrategySchema() *shared.FeatureStrategySchema {
	if o == nil {
		return nil
	}
	return o.FeatureStrategySchema
}

func (o *PatchFeatureStrategyResponse) GetPatchFeatureStrategy400ApplicationJSONObject() *PatchFeatureStrategy400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureStrategy400ApplicationJSONObject
}

func (o *PatchFeatureStrategyResponse) GetPatchFeatureStrategy401ApplicationJSONObject() *PatchFeatureStrategy401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureStrategy401ApplicationJSONObject
}

func (o *PatchFeatureStrategyResponse) GetPatchFeatureStrategy403ApplicationJSONObject() *PatchFeatureStrategy403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureStrategy403ApplicationJSONObject
}

func (o *PatchFeatureStrategyResponse) GetPatchFeatureStrategy404ApplicationJSONObject() *PatchFeatureStrategy404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureStrategy404ApplicationJSONObject
}

func (o *PatchFeatureStrategyResponse) GetPatchFeatureStrategy415ApplicationJSONObject() *PatchFeatureStrategy415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PatchFeatureStrategy415ApplicationJSONObject
}
