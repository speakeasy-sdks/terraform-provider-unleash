// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// BulkMetricsEdgeResponseResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type BulkMetricsEdgeResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkMetricsEdgeResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkMetricsEdgeResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkMetricsEdgeResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkMetricsEdgeResponseBody - The request body is larger than what we accept. By default we only accept bodies of 100kB or less
type BulkMetricsEdgeResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkMetricsEdgeResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkMetricsEdgeResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkMetricsEdgeResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkMetricsResponseBody - The request data does not match what we expect.
type BulkMetricsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkMetricsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkMetricsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkMetricsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type BulkMetricsResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *BulkMetricsResponseBody
	// The request body is larger than what we accept. By default we only accept bodies of 100kB or less
	FourHundredAndThirteenApplicationJSONObject *BulkMetricsEdgeResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *BulkMetricsEdgeResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *BulkMetricsResponse) GetFourHundredApplicationJSONObject() *BulkMetricsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *BulkMetricsResponse) GetFourHundredAndThirteenApplicationJSONObject() *BulkMetricsEdgeResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThirteenApplicationJSONObject
}

func (o *BulkMetricsResponse) GetFourHundredAndFifteenApplicationJSONObject() *BulkMetricsEdgeResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
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
