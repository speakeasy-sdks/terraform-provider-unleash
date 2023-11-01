// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetFeatureUsageSummaryRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetFeatureUsageSummaryRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// GetFeatureUsageSummary404ApplicationJSON - The requested resource was not found.
type GetFeatureUsageSummary404ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureUsageSummary404ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureUsageSummary404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureUsageSummary404ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeatureUsageSummary403ApplicationJSON - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type GetFeatureUsageSummary403ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureUsageSummary403ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureUsageSummary403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureUsageSummary403ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// GetFeatureUsageSummary401ApplicationJSON - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetFeatureUsageSummary401ApplicationJSON struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetFeatureUsageSummary401ApplicationJSON) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetFeatureUsageSummary401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFeatureUsageSummary401ApplicationJSON) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetFeatureUsageSummaryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureUsageSchema
	FeatureUsageSchema *shared.FeatureUsageSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	GetFeatureUsageSummary401ApplicationJSONObject *GetFeatureUsageSummary401ApplicationJSON
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	GetFeatureUsageSummary403ApplicationJSONObject *GetFeatureUsageSummary403ApplicationJSON
	// The requested resource was not found.
	GetFeatureUsageSummary404ApplicationJSONObject *GetFeatureUsageSummary404ApplicationJSON
}

func (o *GetFeatureUsageSummaryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFeatureUsageSummaryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFeatureUsageSummaryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFeatureUsageSummaryResponse) GetFeatureUsageSchema() *shared.FeatureUsageSchema {
	if o == nil {
		return nil
	}
	return o.FeatureUsageSchema
}

func (o *GetFeatureUsageSummaryResponse) GetGetFeatureUsageSummary401ApplicationJSONObject() *GetFeatureUsageSummary401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetFeatureUsageSummary401ApplicationJSONObject
}

func (o *GetFeatureUsageSummaryResponse) GetGetFeatureUsageSummary403ApplicationJSONObject() *GetFeatureUsageSummary403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetFeatureUsageSummary403ApplicationJSONObject
}

func (o *GetFeatureUsageSummaryResponse) GetGetFeatureUsageSummary404ApplicationJSONObject() *GetFeatureUsageSummary404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetFeatureUsageSummary404ApplicationJSONObject
}
