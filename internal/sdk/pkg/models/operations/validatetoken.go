// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// ValidateTokenAuthResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type ValidateTokenAuthResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateTokenAuthResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateTokenAuthResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateTokenAuthResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ValidateTokenResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ValidateTokenResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ValidateTokenResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ValidateTokenResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ValidateTokenResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ValidateTokenResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *ValidateTokenResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *ValidateTokenAuthResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// tokenUserSchema
	TokenUserSchema *shared.TokenUserSchema
}

func (o *ValidateTokenResponse) GetFourHundredAndOneApplicationJSONObject() *ValidateTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *ValidateTokenResponse) GetFourHundredAndFifteenApplicationJSONObject() *ValidateTokenAuthResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *ValidateTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ValidateTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ValidateTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ValidateTokenResponse) GetTokenUserSchema() *shared.TokenUserSchema {
	if o == nil {
		return nil
	}
	return o.TokenUserSchema
}
