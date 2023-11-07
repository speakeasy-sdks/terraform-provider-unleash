// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetEventsRequest struct {
	// The name of the project whose events you want to retrieve
	Project *string `queryParam:"style=form,explode=true,name=project"`
}

func (o *GetEventsRequest) GetProject() *string {
	if o == nil {
		return nil
	}
	return o.Project
}

// GetEventsResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetEventsResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEventsResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEventsResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEventsResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// eventsSchema
	EventsSchema *shared.EventsSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Object *GetEventsResponseBody
}

func (o *GetEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEventsResponse) GetEventsSchema() *shared.EventsSchema {
	if o == nil {
		return nil
	}
	return o.EventsSchema
}

func (o *GetEventsResponse) GetObject() *GetEventsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
