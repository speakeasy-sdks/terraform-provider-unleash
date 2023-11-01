// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// BulkMetrics415ApplicationJSON - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type BulkMetrics415ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkMetrics415ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkMetrics415ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkMetrics415ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkMetrics413ApplicationJSON - The request body is larger than what we accept. By default we only accept bodies of 100kB or less
type BulkMetrics413ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkMetrics413ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkMetrics413ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkMetrics413ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkMetrics400ApplicationJSON - The request data does not match what we expect.
type BulkMetrics400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkMetrics400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkMetrics400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkMetrics400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type BulkMetricsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	BulkMetrics400ApplicationJSONObject *BulkMetrics400ApplicationJSON
	// The request body is larger than what we accept. By default we only accept bodies of 100kB or less
	BulkMetrics413ApplicationJSONObject *BulkMetrics413ApplicationJSON
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	BulkMetrics415ApplicationJSONObject *BulkMetrics415ApplicationJSON
}

func (o *BulkMetricsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BulkMetricsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BulkMetricsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *BulkMetricsResponse) GetBulkMetrics400ApplicationJSONObject() *BulkMetrics400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.BulkMetrics400ApplicationJSONObject
}

func (o *BulkMetricsResponse) GetBulkMetrics413ApplicationJSONObject() *BulkMetrics413ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.BulkMetrics413ApplicationJSONObject
}

func (o *BulkMetricsResponse) GetBulkMetrics415ApplicationJSONObject() *BulkMetrics415ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.BulkMetrics415ApplicationJSONObject
}
