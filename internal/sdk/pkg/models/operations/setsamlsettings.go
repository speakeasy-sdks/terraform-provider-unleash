// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

// SetSamlSettingsAuthResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type SetSamlSettingsAuthResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSamlSettingsAuthResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSamlSettingsAuthResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSamlSettingsAuthResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSamlSettingsAuthResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type SetSamlSettingsAuthResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSamlSettingsAuthResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSamlSettingsAuthResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSamlSettingsAuthResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSamlSettingsAuthResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type SetSamlSettingsAuthResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSamlSettingsAuthResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSamlSettingsAuthResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSamlSettingsAuthResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// SetSamlSettingsResponseBody - The request data does not match what we expect.
type SetSamlSettingsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *SetSamlSettingsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SetSamlSettingsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *SetSamlSettingsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type SetSamlSettingsResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *SetSamlSettingsResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *SetSamlSettingsAuthResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *SetSamlSettingsAuthResponseResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *SetSamlSettingsAuthResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// samlSettingsSchema
	SamlSettingsSchema *shared.SamlSettingsSchema
}

func (o *SetSamlSettingsResponse) GetFourHundredApplicationJSONObject() *SetSamlSettingsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *SetSamlSettingsResponse) GetFourHundredAndOneApplicationJSONObject() *SetSamlSettingsAuthResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *SetSamlSettingsResponse) GetFourHundredAndThreeApplicationJSONObject() *SetSamlSettingsAuthResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *SetSamlSettingsResponse) GetFourHundredAndFifteenApplicationJSONObject() *SetSamlSettingsAuthResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *SetSamlSettingsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SetSamlSettingsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SetSamlSettingsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *SetSamlSettingsResponse) GetSamlSettingsSchema() *shared.SamlSettingsSchema {
	if o == nil {
		return nil
	}
	return o.SamlSettingsSchema
}
