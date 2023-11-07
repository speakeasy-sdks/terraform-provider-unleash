// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateFeatureStrategyRequest struct {
	Environment string `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
	StrategyID  string `pathParam:"style=simple,explode=false,name=strategyId"`
	// updateFeatureStrategySchema
	UpdateFeatureStrategySchema shared.UpdateFeatureStrategySchema `request:"mediaType=application/json"`
}

func (o *UpdateFeatureStrategyRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *UpdateFeatureStrategyRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *UpdateFeatureStrategyRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *UpdateFeatureStrategyRequest) GetStrategyID() string {
	if o == nil {
		return ""
	}
	return o.StrategyID
}

func (o *UpdateFeatureStrategyRequest) GetUpdateFeatureStrategySchema() shared.UpdateFeatureStrategySchema {
	if o == nil {
		return shared.UpdateFeatureStrategySchema{}
	}
	return o.UpdateFeatureStrategySchema
}

// UpdateFeatureStrategyFeaturesResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type UpdateFeatureStrategyFeaturesResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategyFeaturesResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategyFeaturesResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategyFeaturesResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureStrategyFeaturesResponse404ResponseBody - The requested resource was not found.
type UpdateFeatureStrategyFeaturesResponse404ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategyFeaturesResponse404ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategyFeaturesResponse404ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategyFeaturesResponse404ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureStrategyFeaturesResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type UpdateFeatureStrategyFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategyFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategyFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategyFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureStrategyFeaturesResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type UpdateFeatureStrategyFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategyFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategyFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategyFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureStrategyResponseBody - The request data does not match what we expect.
type UpdateFeatureStrategyResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategyResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategyResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategyResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateFeatureStrategyResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *UpdateFeatureStrategyResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *UpdateFeatureStrategyFeaturesResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *UpdateFeatureStrategyFeaturesResponseResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *UpdateFeatureStrategyFeaturesResponse404ResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *UpdateFeatureStrategyFeaturesResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureStrategySchema
	FeatureStrategySchema *shared.FeatureStrategySchema
}

func (o *UpdateFeatureStrategyResponse) GetFourHundredApplicationJSONObject() *UpdateFeatureStrategyResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *UpdateFeatureStrategyResponse) GetFourHundredAndOneApplicationJSONObject() *UpdateFeatureStrategyFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *UpdateFeatureStrategyResponse) GetFourHundredAndThreeApplicationJSONObject() *UpdateFeatureStrategyFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *UpdateFeatureStrategyResponse) GetFourHundredAndFourApplicationJSONObject() *UpdateFeatureStrategyFeaturesResponse404ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *UpdateFeatureStrategyResponse) GetFourHundredAndFifteenApplicationJSONObject() *UpdateFeatureStrategyFeaturesResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *UpdateFeatureStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFeatureStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFeatureStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateFeatureStrategyResponse) GetFeatureStrategySchema() *shared.FeatureStrategySchema {
	if o == nil {
		return nil
	}
	return o.FeatureStrategySchema
}
