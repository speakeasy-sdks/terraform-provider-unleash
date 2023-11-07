// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// GetValidTokensEdgeResponseResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type GetValidTokensEdgeResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetValidTokensEdgeResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetValidTokensEdgeResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetValidTokensEdgeResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetValidTokensEdgeResponseBody - The request body is larger than what we accept. By default we only accept bodies of 100kB or less
type GetValidTokensEdgeResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetValidTokensEdgeResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetValidTokensEdgeResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetValidTokensEdgeResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetValidTokensResponseBody - The request data does not match what we expect.
type GetValidTokensResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetValidTokensResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetValidTokensResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetValidTokensResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetValidTokensResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *GetValidTokensResponseBody
	// The request body is larger than what we accept. By default we only accept bodies of 100kB or less
	FourHundredAndThirteenApplicationJSONObject *GetValidTokensEdgeResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *GetValidTokensEdgeResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// validatedEdgeTokensSchema
	ValidatedEdgeTokensSchema *shared.ValidatedEdgeTokensSchema
}

func (o *GetValidTokensResponse) GetFourHundredApplicationJSONObject() *GetValidTokensResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *GetValidTokensResponse) GetFourHundredAndThirteenApplicationJSONObject() *GetValidTokensEdgeResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThirteenApplicationJSONObject
}

func (o *GetValidTokensResponse) GetFourHundredAndFifteenApplicationJSONObject() *GetValidTokensEdgeResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *GetValidTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetValidTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetValidTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetValidTokensResponse) GetValidatedEdgeTokensSchema() *shared.ValidatedEdgeTokensSchema {
	if o == nil {
		return nil
	}
	return o.ValidatedEdgeTokensSchema
}
