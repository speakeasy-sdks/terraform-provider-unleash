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

// DeprecateStrategy404ApplicationJSON - The requested resource was not found.
type DeprecateStrategy404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeprecateStrategy404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeprecateStrategy404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeprecateStrategy404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeprecateStrategy403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeprecateStrategy403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeprecateStrategy403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeprecateStrategy403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeprecateStrategy403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeprecateStrategy401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeprecateStrategy401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeprecateStrategy401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeprecateStrategy401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeprecateStrategy401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeprecateStrategyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	DeprecateStrategy401ApplicationJSONObject *DeprecateStrategy401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	DeprecateStrategy403ApplicationJSONObject *DeprecateStrategy403ApplicationJSON
	// The requested resource was not found.
	DeprecateStrategy404ApplicationJSONObject *DeprecateStrategy404ApplicationJSON
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

func (o *DeprecateStrategyResponse) GetDeprecateStrategy401ApplicationJSONObject() *DeprecateStrategy401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeprecateStrategy401ApplicationJSONObject
}

func (o *DeprecateStrategyResponse) GetDeprecateStrategy403ApplicationJSONObject() *DeprecateStrategy403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeprecateStrategy403ApplicationJSONObject
}

func (o *DeprecateStrategyResponse) GetDeprecateStrategy404ApplicationJSONObject() *DeprecateStrategy404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.DeprecateStrategy404ApplicationJSONObject
}
