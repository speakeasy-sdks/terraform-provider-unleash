// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateAPITokenRequest struct {
	Token string `pathParam:"style=simple,explode=false,name=token"`
	// updateApiTokenSchema
	UpdateAPITokenSchema shared.UpdateAPITokenSchema `request:"mediaType=application/json"`
}

func (o *UpdateAPITokenRequest) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *UpdateAPITokenRequest) GetUpdateAPITokenSchema() shared.UpdateAPITokenSchema {
	if o == nil {
		return shared.UpdateAPITokenSchema{}
	}
	return o.UpdateAPITokenSchema
}

// UpdateAPITokenAPITokensResponseResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type UpdateAPITokenAPITokensResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateAPITokenAPITokensResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateAPITokenAPITokensResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateAPITokenAPITokensResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateAPITokenAPITokensResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type UpdateAPITokenAPITokensResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateAPITokenAPITokensResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateAPITokenAPITokensResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateAPITokenAPITokensResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateAPITokenResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type UpdateAPITokenResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateAPITokenResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateAPITokenResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateAPITokenResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateAPITokenResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *UpdateAPITokenResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *UpdateAPITokenAPITokensResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *UpdateAPITokenAPITokensResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateAPITokenResponse) GetFourHundredAndOneApplicationJSONObject() *UpdateAPITokenResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *UpdateAPITokenResponse) GetFourHundredAndThreeApplicationJSONObject() *UpdateAPITokenAPITokensResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *UpdateAPITokenResponse) GetFourHundredAndFifteenApplicationJSONObject() *UpdateAPITokenAPITokensResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *UpdateAPITokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAPITokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAPITokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
