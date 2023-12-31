// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteApplicationRequest struct {
	AppName string `pathParam:"style=simple,explode=false,name=appName"`
}

func (o *DeleteApplicationRequest) GetAppName() string {
	if o == nil {
		return ""
	}
	return o.AppName
}

// DeleteApplicationMetricsResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type DeleteApplicationMetricsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteApplicationMetricsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteApplicationMetricsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteApplicationMetricsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// DeleteApplicationResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type DeleteApplicationResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *DeleteApplicationResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DeleteApplicationResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *DeleteApplicationResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type DeleteApplicationResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *DeleteApplicationResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *DeleteApplicationMetricsResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteApplicationResponse) GetFourHundredAndOneApplicationJSONObject() *DeleteApplicationResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *DeleteApplicationResponse) GetFourHundredAndThreeApplicationJSONObject() *DeleteApplicationMetricsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *DeleteApplicationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteApplicationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteApplicationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
