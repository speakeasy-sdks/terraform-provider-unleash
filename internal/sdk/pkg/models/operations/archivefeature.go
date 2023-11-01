// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ArchiveFeatureRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *ArchiveFeatureRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *ArchiveFeatureRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// ArchiveFeature404ApplicationJSON - The requested resource was not found.
type ArchiveFeature404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ArchiveFeature404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ArchiveFeature404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ArchiveFeature404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ArchiveFeature403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ArchiveFeature403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ArchiveFeature403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ArchiveFeature403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ArchiveFeature403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ArchiveFeature401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ArchiveFeature401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ArchiveFeature401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ArchiveFeature401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ArchiveFeature401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ArchiveFeatureResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	ArchiveFeature401ApplicationJSONObject *ArchiveFeature401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	ArchiveFeature403ApplicationJSONObject *ArchiveFeature403ApplicationJSON
	// The requested resource was not found.
	ArchiveFeature404ApplicationJSONObject *ArchiveFeature404ApplicationJSON
}

func (o *ArchiveFeatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ArchiveFeatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ArchiveFeatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ArchiveFeatureResponse) GetArchiveFeature401ApplicationJSONObject() *ArchiveFeature401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ArchiveFeature401ApplicationJSONObject
}

func (o *ArchiveFeatureResponse) GetArchiveFeature403ApplicationJSONObject() *ArchiveFeature403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ArchiveFeature403ApplicationJSONObject
}

func (o *ArchiveFeatureResponse) GetArchiveFeature404ApplicationJSONObject() *ArchiveFeature404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ArchiveFeature404ApplicationJSONObject
}
