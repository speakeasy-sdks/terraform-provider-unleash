// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// RegisterClientMetrics400ApplicationJSON - The request data does not match what we expect.
type RegisterClientMetrics400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *RegisterClientMetrics400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RegisterClientMetrics400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RegisterClientMetrics400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type RegisterClientMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	RegisterClientMetrics400ApplicationJSONObject *RegisterClientMetrics400ApplicationJSON
}

func (o *RegisterClientMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RegisterClientMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RegisterClientMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RegisterClientMetricsResponse) GetRegisterClientMetrics400ApplicationJSONObject() *RegisterClientMetrics400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.RegisterClientMetrics400ApplicationJSONObject
}
