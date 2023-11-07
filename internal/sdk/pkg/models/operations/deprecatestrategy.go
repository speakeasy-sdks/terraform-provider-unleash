// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeprecateStrategyRequest struct {
	StrategyName string `pathParam:"style=simple,explode=false,name=strategyName"`
}

func (o *DeprecateStrategyRequest) GetStrategyName() string {
	if o == nil {
		return ""
	}
	return o.StrategyName
}

// DeprecateStrategyStrategiesResponseResponseBody - The requested resource was not found.
type DeprecateStrategyStrategiesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeprecateStrategyStrategiesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeprecateStrategyStrategiesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeprecateStrategyStrategiesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeprecateStrategyStrategiesResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeprecateStrategyStrategiesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeprecateStrategyStrategiesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeprecateStrategyStrategiesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeprecateStrategyStrategiesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeprecateStrategyResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeprecateStrategyResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeprecateStrategyResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeprecateStrategyResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeprecateStrategyResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeprecateStrategyResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *DeprecateStrategyResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *DeprecateStrategyStrategiesResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *DeprecateStrategyStrategiesResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeprecateStrategyResponse) GetFourHundredAndOneApplicationJSONObject() *DeprecateStrategyResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *DeprecateStrategyResponse) GetFourHundredAndThreeApplicationJSONObject() *DeprecateStrategyStrategiesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *DeprecateStrategyResponse) GetFourHundredAndFourApplicationJSONObject() *DeprecateStrategyStrategiesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *DeprecateStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeprecateStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeprecateStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
