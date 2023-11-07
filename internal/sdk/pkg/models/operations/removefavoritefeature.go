// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type RemoveFavoriteFeatureRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
	ProjectID   string `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *RemoveFavoriteFeatureRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

func (o *RemoveFavoriteFeatureRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// RemoveFavoriteFeatureFeaturesResponseBody - The requested resource was not found.
type RemoveFavoriteFeatureFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *RemoveFavoriteFeatureFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RemoveFavoriteFeatureFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RemoveFavoriteFeatureFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// RemoveFavoriteFeatureResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type RemoveFavoriteFeatureResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *RemoveFavoriteFeatureResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RemoveFavoriteFeatureResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *RemoveFavoriteFeatureResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type RemoveFavoriteFeatureResponse struct {
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *RemoveFavoriteFeatureResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *RemoveFavoriteFeatureFeaturesResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemoveFavoriteFeatureResponse) GetFourHundredAndOneApplicationJSONObject() *RemoveFavoriteFeatureResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *RemoveFavoriteFeatureResponse) GetFourHundredAndFourApplicationJSONObject() *RemoveFavoriteFeatureFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *RemoveFavoriteFeatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveFavoriteFeatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveFavoriteFeatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
