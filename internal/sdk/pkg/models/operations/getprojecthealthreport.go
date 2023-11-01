// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetProjectHealthReportRequest struct {
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *GetProjectHealthReportRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// GetProjectHealthReport404ApplicationJSON - The requested resource was not found.
type GetProjectHealthReport404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetProjectHealthReport404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetProjectHealthReport404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetProjectHealthReport404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetProjectHealthReport403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetProjectHealthReport403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetProjectHealthReport403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetProjectHealthReport403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetProjectHealthReport403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetProjectHealthReport401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetProjectHealthReport401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetProjectHealthReport401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetProjectHealthReport401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetProjectHealthReport401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetProjectHealthReportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetProjectHealthReport401ApplicationJSONObject *GetProjectHealthReport401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetProjectHealthReport403ApplicationJSONObject *GetProjectHealthReport403ApplicationJSON
	// The requested resource was not found.
	GetProjectHealthReport404ApplicationJSONObject *GetProjectHealthReport404ApplicationJSON
	// healthReportSchema
	HealthReportSchema *shared.HealthReportSchema
}

func (o *GetProjectHealthReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetProjectHealthReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetProjectHealthReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetProjectHealthReportResponse) GetGetProjectHealthReport401ApplicationJSONObject() *GetProjectHealthReport401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetProjectHealthReport401ApplicationJSONObject
}

func (o *GetProjectHealthReportResponse) GetGetProjectHealthReport403ApplicationJSONObject() *GetProjectHealthReport403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetProjectHealthReport403ApplicationJSONObject
}

func (o *GetProjectHealthReportResponse) GetGetProjectHealthReport404ApplicationJSONObject() *GetProjectHealthReport404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetProjectHealthReport404ApplicationJSONObject
}

func (o *GetProjectHealthReportResponse) GetHealthReportSchema() *shared.HealthReportSchema {
	if o == nil {
		return nil
	}
	return o.HealthReportSchema
}
