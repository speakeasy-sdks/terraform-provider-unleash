// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateFeedbackRequest struct {
	// feedbackUpdateSchema
	FeedbackUpdateSchema shared.FeedbackUpdateSchema `request:"mediaType=application/json"`
	ID                   string                      `pathParam:"style=simple,explode=false,name=id"`
}

func (o *UpdateFeedbackRequest) GetFeedbackUpdateSchema() shared.FeedbackUpdateSchema {
	if o == nil {
		return shared.FeedbackUpdateSchema{}
	}
	return o.FeedbackUpdateSchema
}

func (o *UpdateFeedbackRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// UpdateFeedbackAdminUIResponseResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type UpdateFeedbackAdminUIResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeedbackAdminUIResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeedbackAdminUIResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeedbackAdminUIResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeedbackAdminUIResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type UpdateFeedbackAdminUIResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeedbackAdminUIResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeedbackAdminUIResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeedbackAdminUIResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeedbackResponseBody - The request data does not match what we expect.
type UpdateFeedbackResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeedbackResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeedbackResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeedbackResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateFeedbackResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *UpdateFeedbackResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *UpdateFeedbackAdminUIResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *UpdateFeedbackAdminUIResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// feedbackResponseSchema
	FeedbackResponseSchema *shared.FeedbackResponseSchema
}

func (o *UpdateFeedbackResponse) GetFourHundredApplicationJSONObject() *UpdateFeedbackResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *UpdateFeedbackResponse) GetFourHundredAndOneApplicationJSONObject() *UpdateFeedbackAdminUIResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *UpdateFeedbackResponse) GetFourHundredAndFifteenApplicationJSONObject() *UpdateFeedbackAdminUIResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *UpdateFeedbackResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFeedbackResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFeedbackResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateFeedbackResponse) GetFeedbackResponseSchema() *shared.FeedbackResponseSchema {
	if o == nil {
		return nil
	}
	return o.FeedbackResponseSchema
}
