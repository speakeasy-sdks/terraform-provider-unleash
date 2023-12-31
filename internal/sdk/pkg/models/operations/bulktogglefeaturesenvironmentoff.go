// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type BulkToggleFeaturesEnvironmentOffRequest struct {
	// bulkToggleFeaturesSchema
	BulkToggleFeaturesSchema shared.BulkToggleFeaturesSchema `request:"mediaType=application/json"`
	Environment              string                          `pathParam:"style=simple,explode=false,name=environment"`
	ProjectID                string                          `pathParam:"style=simple,explode=false,name=projectId"`
}

func (o *BulkToggleFeaturesEnvironmentOffRequest) GetBulkToggleFeaturesSchema() shared.BulkToggleFeaturesSchema {
	if o == nil {
		return shared.BulkToggleFeaturesSchema{}
	}
	return o.BulkToggleFeaturesSchema
}

func (o *BulkToggleFeaturesEnvironmentOffRequest) GetEnvironment() string {
	if o == nil {
		return ""
	}
	return o.Environment
}

func (o *BulkToggleFeaturesEnvironmentOffRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// BulkToggleFeaturesEnvironmentOffFeaturesResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type BulkToggleFeaturesEnvironmentOffFeaturesResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkToggleFeaturesEnvironmentOffFeaturesResponse413ResponseBody - The request body is larger than what we accept. By default we only accept bodies of 100kB or less
type BulkToggleFeaturesEnvironmentOffFeaturesResponse413ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse413ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse413ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse413ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkToggleFeaturesEnvironmentOffFeaturesResponse404ResponseBody - The requested resource was not found.
type BulkToggleFeaturesEnvironmentOffFeaturesResponse404ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse404ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse404ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponse404ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkToggleFeaturesEnvironmentOffFeaturesResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type BulkToggleFeaturesEnvironmentOffFeaturesResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkToggleFeaturesEnvironmentOffFeaturesResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type BulkToggleFeaturesEnvironmentOffFeaturesResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkToggleFeaturesEnvironmentOffFeaturesResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// BulkToggleFeaturesEnvironmentOffResponseBody - The request data does not match what we expect.
type BulkToggleFeaturesEnvironmentOffResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *BulkToggleFeaturesEnvironmentOffResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BulkToggleFeaturesEnvironmentOffResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *BulkToggleFeaturesEnvironmentOffResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type BulkToggleFeaturesEnvironmentOffResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *BulkToggleFeaturesEnvironmentOffResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *BulkToggleFeaturesEnvironmentOffFeaturesResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *BulkToggleFeaturesEnvironmentOffFeaturesResponseResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *BulkToggleFeaturesEnvironmentOffFeaturesResponse404ResponseBody
	// The request body is larger than what we accept. By default we only accept bodies of 100kB or less
	FourHundredAndThirteenApplicationJSONObject *BulkToggleFeaturesEnvironmentOffFeaturesResponse413ResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *BulkToggleFeaturesEnvironmentOffFeaturesResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetFourHundredApplicationJSONObject() *BulkToggleFeaturesEnvironmentOffResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetFourHundredAndOneApplicationJSONObject() *BulkToggleFeaturesEnvironmentOffFeaturesResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetFourHundredAndThreeApplicationJSONObject() *BulkToggleFeaturesEnvironmentOffFeaturesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetFourHundredAndFourApplicationJSONObject() *BulkToggleFeaturesEnvironmentOffFeaturesResponse404ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetFourHundredAndThirteenApplicationJSONObject() *BulkToggleFeaturesEnvironmentOffFeaturesResponse413ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThirteenApplicationJSONObject
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetFourHundredAndFifteenApplicationJSONObject() *BulkToggleFeaturesEnvironmentOffFeaturesResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BulkToggleFeaturesEnvironmentOffResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
