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

// GetEvents401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetEvents401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetEvents401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetEvents401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetEvents401ApplicationJSON) GetName() *string {
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
	GetEvents401ApplicationJSONObject *GetEvents401ApplicationJSON
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

func (o *GetEventsResponse) GetGetEvents401ApplicationJSONObject() *GetEvents401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetEvents401ApplicationJSONObject
}
