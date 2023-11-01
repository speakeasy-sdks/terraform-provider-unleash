// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveEnvironmentFromProjectRequest struct {
	Environment string `pathParam:"style=simple,explode=false,name=environment"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *RemoveEnvironmentFromProjectRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *RemoveEnvironmentFromProjectRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// RemoveEnvironmentFromProject403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type RemoveEnvironmentFromProject403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *RemoveEnvironmentFromProject403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RemoveEnvironmentFromProject403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RemoveEnvironmentFromProject403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// RemoveEnvironmentFromProject401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type RemoveEnvironmentFromProject401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *RemoveEnvironmentFromProject401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RemoveEnvironmentFromProject401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RemoveEnvironmentFromProject401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// RemoveEnvironmentFromProject400ApplicationJSON - The request data does not match what we expect.
type RemoveEnvironmentFromProject400ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *RemoveEnvironmentFromProject400ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RemoveEnvironmentFromProject400ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RemoveEnvironmentFromProject400ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type RemoveEnvironmentFromProjectResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The request data does not match what we expect.
	RemoveEnvironmentFromProject400ApplicationJSONObject *RemoveEnvironmentFromProject400ApplicationJSON
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	RemoveEnvironmentFromProject401ApplicationJSONObject *RemoveEnvironmentFromProject401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	RemoveEnvironmentFromProject403ApplicationJSONObject *RemoveEnvironmentFromProject403ApplicationJSON
}

func (o *RemoveEnvironmentFromProjectResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveEnvironmentFromProjectResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveEnvironmentFromProjectResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveEnvironmentFromProjectResponse) GetRemoveEnvironmentFromProject400ApplicationJSONObject() *RemoveEnvironmentFromProject400ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.RemoveEnvironmentFromProject400ApplicationJSONObject
}

func (o *RemoveEnvironmentFromProjectResponse) GetRemoveEnvironmentFromProject401ApplicationJSONObject() *RemoveEnvironmentFromProject401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.RemoveEnvironmentFromProject401ApplicationJSONObject
}

func (o *RemoveEnvironmentFromProjectResponse) GetRemoveEnvironmentFromProject403ApplicationJSONObject() *RemoveEnvironmentFromProject403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.RemoveEnvironmentFromProject403ApplicationJSONObject
}
