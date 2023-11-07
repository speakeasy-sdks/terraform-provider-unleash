// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeletePatRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeletePatRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// DeletePatPersonalAccessTokensResponseResponseBody - The requested resource was not found.
type DeletePatPersonalAccessTokensResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeletePatPersonalAccessTokensResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeletePatPersonalAccessTokensResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeletePatPersonalAccessTokensResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeletePatPersonalAccessTokensResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeletePatPersonalAccessTokensResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeletePatPersonalAccessTokensResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeletePatPersonalAccessTokensResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeletePatPersonalAccessTokensResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeletePatResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeletePatResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeletePatResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeletePatResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeletePatResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeletePatResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *DeletePatResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *DeletePatPersonalAccessTokensResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *DeletePatPersonalAccessTokensResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeletePatResponse) GetFourHundredAndOneApplicationJSONObject() *DeletePatResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *DeletePatResponse) GetFourHundredAndThreeApplicationJSONObject() *DeletePatPersonalAccessTokensResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *DeletePatResponse) GetFourHundredAndFourApplicationJSONObject() *DeletePatPersonalAccessTokensResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *DeletePatResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePatResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePatResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
