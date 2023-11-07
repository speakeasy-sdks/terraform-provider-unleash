// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetFeatureStrategyRequest struct {
	Environment string `pathParam:"style=simple,explode=false,name=environment"`
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
	StrategyID  string `pathParam:"style=simple,explode=false,name=strategyId"`
}

func (o *GetFeatureStrategyRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *GetFeatureStrategyRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *GetFeatureStrategyRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *GetFeatureStrategyRequest) GetStrategyID() string {
	if o == nil {
		return ""
	}
	return o.StrategyID
}

// GetFeatureStrategyFeaturesResponseResponseBody - The requested resource was not found.
type GetFeatureStrategyFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureStrategyFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureStrategyFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureStrategyFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeatureStrategyFeaturesResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetFeatureStrategyFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureStrategyFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureStrategyFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureStrategyFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeatureStrategyResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetFeatureStrategyResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureStrategyResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureStrategyResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureStrategyResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetFeatureStrategyResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetFeatureStrategyResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *GetFeatureStrategyFeaturesResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *GetFeatureStrategyFeaturesResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureStrategySchema
	FeatureStrategySchema *shared.FeatureStrategySchema
}

func (o *GetFeatureStrategyResponse) GetFourHundredAndOneApplicationJSONObject() *GetFeatureStrategyResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetFeatureStrategyResponse) GetFourHundredAndThreeApplicationJSONObject() *GetFeatureStrategyFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetFeatureStrategyResponse) GetFourHundredAndFourApplicationJSONObject() *GetFeatureStrategyFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *GetFeatureStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFeatureStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFeatureStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFeatureStrategyResponse) GetFeatureStrategySchema() *shared.FeatureStrategySchema {
	if o == nil {
		return nil
	}
	return o.FeatureStrategySchema
}
