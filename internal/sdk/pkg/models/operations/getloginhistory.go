// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

// GetLoginHistoryUnstableResponseBody - The requested resource was not found.
type GetLoginHistoryUnstableResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetLoginHistoryUnstableResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetLoginHistoryUnstableResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetLoginHistoryUnstableResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetLoginHistoryResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetLoginHistoryResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetLoginHistoryResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetLoginHistoryResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetLoginHistoryResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetLoginHistoryResponse struct {
	// loginHistorySchema
	TwoHundredTextCsvRes *string
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *GetLoginHistoryResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *GetLoginHistoryUnstableResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// loginHistorySchema
	LoginHistorySchema *shared.LoginHistorySchema
}

func (o *GetLoginHistoryResponse) GetTwoHundredTextCsvRes() *string {
	if o == nil {
		return nil
	}
	return o.TwoHundredTextCsvRes
}

func (o *GetLoginHistoryResponse) GetFourHundredAndOneApplicationJSONObject() *GetLoginHistoryResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *GetLoginHistoryResponse) GetFourHundredAndFourApplicationJSONObject() *GetLoginHistoryUnstableResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *GetLoginHistoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLoginHistoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLoginHistoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLoginHistoryResponse) GetLoginHistorySchema() *shared.LoginHistorySchema {
	if o == nil {
		return nil
	}
	return o.LoginHistorySchema
}
