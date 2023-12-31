// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// GetPatsPersonalAccessTokensResponseResponseBody - The requested resource was not found.
type GetPatsPersonalAccessTokensResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetPatsPersonalAccessTokensResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPatsPersonalAccessTokensResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPatsPersonalAccessTokensResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetPatsPersonalAccessTokensResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetPatsPersonalAccessTokensResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetPatsPersonalAccessTokensResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPatsPersonalAccessTokensResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPatsPersonalAccessTokensResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetPatsResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetPatsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetPatsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPatsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPatsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetPatsResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetPatsResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *GetPatsPersonalAccessTokensResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *GetPatsPersonalAccessTokensResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// patsSchema
	PatsSchema *shared.PatsSchema
}

func (o *GetPatsResponse) GetFourHundredAndOneApplicationJSONObject() *GetPatsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetPatsResponse) GetFourHundredAndThreeApplicationJSONObject() *GetPatsPersonalAccessTokensResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *GetPatsResponse) GetFourHundredAndFourApplicationJSONObject() *GetPatsPersonalAccessTokensResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *GetPatsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPatsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPatsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPatsResponse) GetPatsSchema() *shared.PatsSchema {
	if o == nil {
		return nil
	}
	return o.PatsSchema
}
