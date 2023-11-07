// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// UpdateFeatureStrategySegmentsStrategiesResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type UpdateFeatureStrategySegmentsStrategiesResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureStrategySegmentsStrategiesResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type UpdateFeatureStrategySegmentsStrategiesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureStrategySegmentsStrategiesResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type UpdateFeatureStrategySegmentsStrategiesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategySegmentsStrategiesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureStrategySegmentsResponseBody - The request data does not match what we expect.
type UpdateFeatureStrategySegmentsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureStrategySegmentsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureStrategySegmentsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureStrategySegmentsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateFeatureStrategySegmentsResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *UpdateFeatureStrategySegmentsResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *UpdateFeatureStrategySegmentsStrategiesResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *UpdateFeatureStrategySegmentsStrategiesResponseResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *UpdateFeatureStrategySegmentsStrategiesResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The resource was successfully created.
	UpdateFeatureStrategySegmentsSchema *shared.UpdateFeatureStrategySegmentsSchema
}

func (o *UpdateFeatureStrategySegmentsResponse) GetFourHundredApplicationJSONObject() *UpdateFeatureStrategySegmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *UpdateFeatureStrategySegmentsResponse) GetFourHundredAndOneApplicationJSONObject() *UpdateFeatureStrategySegmentsStrategiesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *UpdateFeatureStrategySegmentsResponse) GetFourHundredAndThreeApplicationJSONObject() *UpdateFeatureStrategySegmentsStrategiesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *UpdateFeatureStrategySegmentsResponse) GetFourHundredAndFifteenApplicationJSONObject() *UpdateFeatureStrategySegmentsStrategiesResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *UpdateFeatureStrategySegmentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFeatureStrategySegmentsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *UpdateFeatureStrategySegmentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFeatureStrategySegmentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateFeatureStrategySegmentsResponse) GetUpdateFeatureStrategySegmentsSchema() *shared.UpdateFeatureStrategySegmentsSchema {
	if o == nil {
		return nil
	}
	return o.UpdateFeatureStrategySegmentsSchema
}
