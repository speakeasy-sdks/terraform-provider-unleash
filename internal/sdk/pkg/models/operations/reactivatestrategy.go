// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ReactivateStrategyRequest struct {
	StrategyName string `pathParam:"style=simple,explode=false,name=strategyName"`
}

func (o *ReactivateStrategyRequest) GetStrategyName() string {
	if o == nil {
		return ""
	}
	return o.StrategyName
}

// ReactivateStrategy404ApplicationJSON - The requested resource was not found.
type ReactivateStrategy404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ReactivateStrategy404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ReactivateStrategy404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ReactivateStrategy404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ReactivateStrategy403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ReactivateStrategy403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ReactivateStrategy403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ReactivateStrategy403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ReactivateStrategy403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ReactivateStrategy401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ReactivateStrategy401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ReactivateStrategy401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ReactivateStrategy401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ReactivateStrategy401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ReactivateStrategyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	ReactivateStrategy401ApplicationJSONObject *ReactivateStrategy401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	ReactivateStrategy403ApplicationJSONObject *ReactivateStrategy403ApplicationJSON
	// The requested resource was not found.
	ReactivateStrategy404ApplicationJSONObject *ReactivateStrategy404ApplicationJSON
}

func (o *ReactivateStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReactivateStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReactivateStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReactivateStrategyResponse) GetReactivateStrategy401ApplicationJSONObject() *ReactivateStrategy401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ReactivateStrategy401ApplicationJSONObject
}

func (o *ReactivateStrategyResponse) GetReactivateStrategy403ApplicationJSONObject() *ReactivateStrategy403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ReactivateStrategy403ApplicationJSONObject
}

func (o *ReactivateStrategyResponse) GetReactivateStrategy404ApplicationJSONObject() *ReactivateStrategy404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ReactivateStrategy404ApplicationJSONObject
}
