// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ReviveFeatureRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
}

func (o *ReviveFeatureRequest) GetFeatureName() string {
	if o == nil {
		return ""
	}
	return o.FeatureName
}

// ReviveFeatureArchiveResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type ReviveFeatureArchiveResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ReviveFeatureArchiveResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ReviveFeatureArchiveResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ReviveFeatureArchiveResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ReviveFeatureArchiveResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type ReviveFeatureArchiveResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ReviveFeatureArchiveResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ReviveFeatureArchiveResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ReviveFeatureArchiveResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// ReviveFeatureResponseBody - The request data does not match what we expect.
type ReviveFeatureResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *ReviveFeatureResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ReviveFeatureResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ReviveFeatureResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type ReviveFeatureResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *ReviveFeatureResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *ReviveFeatureArchiveResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *ReviveFeatureArchiveResponseResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReviveFeatureResponse) GetFourHundredApplicationJSONObject() *ReviveFeatureResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *ReviveFeatureResponse) GetFourHundredAndOneApplicationJSONObject() *ReviveFeatureArchiveResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *ReviveFeatureResponse) GetFourHundredAndThreeApplicationJSONObject() *ReviveFeatureArchiveResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *ReviveFeatureResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReviveFeatureResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReviveFeatureResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
