// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type UpdateFeatureTypeLifetimeRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// updateFeatureTypeLifetimeSchema
	UpdateFeatureTypeLifetimeSchema shared.UpdateFeatureTypeLifetimeSchema `request:"mediaType=application/json"`
}

func (o *UpdateFeatureTypeLifetimeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetimeRequest) GetUpdateFeatureTypeLifetimeSchema() shared.UpdateFeatureTypeLifetimeSchema {
	if o == nil {
		return shared.UpdateFeatureTypeLifetimeSchema{}
	}
	return o.UpdateFeatureTypeLifetimeSchema
}

// UpdateFeatureTypeLifetimeUnstableResponse415ResponseBody - The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
type UpdateFeatureTypeLifetimeUnstableResponse415ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse415ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse415ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse415ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetimeUnstableResponse409ResponseBody - The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
type UpdateFeatureTypeLifetimeUnstableResponse409ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse409ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse409ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse409ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetimeUnstableResponse404ResponseBody - The requested resource was not found.
type UpdateFeatureTypeLifetimeUnstableResponse404ResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse404ResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse404ResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetimeUnstableResponse404ResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetimeUnstableResponseResponseBody - The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
type UpdateFeatureTypeLifetimeUnstableResponseResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetimeUnstableResponseResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetimeUnstableResponseResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetimeUnstableResponseResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetimeUnstableResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type UpdateFeatureTypeLifetimeUnstableResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetimeUnstableResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetimeUnstableResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetimeUnstableResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

// UpdateFeatureTypeLifetimeResponseBody - The request data does not match what we expect.
type UpdateFeatureTypeLifetimeResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *UpdateFeatureTypeLifetimeResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UpdateFeatureTypeLifetimeResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *UpdateFeatureTypeLifetimeResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type UpdateFeatureTypeLifetimeResponse struct {
	// The request data does not match what we expect.
	FourHundredApplicationJSONObject *UpdateFeatureTypeLifetimeResponseBody
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	FourHundredAndOneApplicationJSONObject *UpdateFeatureTypeLifetimeUnstableResponseBody
	// The provided user credentials are valid, but the user does not have the necessary permissions to perform this operation
	FourHundredAndThreeApplicationJSONObject *UpdateFeatureTypeLifetimeUnstableResponseResponseBody
	// The requested resource was not found.
	FourHundredAndFourApplicationJSONObject *UpdateFeatureTypeLifetimeUnstableResponse404ResponseBody
	// The provided resource can not be created or updated because it would conflict with the current state of the resource or with an already existing resource, respectively.
	FourHundredAndNineApplicationJSONObject *UpdateFeatureTypeLifetimeUnstableResponse409ResponseBody
	// The operation does not support request payloads of the provided type. Please ensure that you're using one of the listed payload types and that you have specified the right content type in the "content-type" header.
	FourHundredAndFifteenApplicationJSONObject *UpdateFeatureTypeLifetimeUnstableResponse415ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// featureTypeSchema
	FeatureTypeSchema *shared.FeatureTypeSchema
}

func (o *UpdateFeatureTypeLifetimeResponse) GetFourHundredApplicationJSONObject() *UpdateFeatureTypeLifetimeResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetFourHundredAndOneApplicationJSONObject() *UpdateFeatureTypeLifetimeUnstableResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetFourHundredAndThreeApplicationJSONObject() *UpdateFeatureTypeLifetimeUnstableResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndThreeApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetFourHundredAndFourApplicationJSONObject() *UpdateFeatureTypeLifetimeUnstableResponse404ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFourApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetFourHundredAndNineApplicationJSONObject() *UpdateFeatureTypeLifetimeUnstableResponse409ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndNineApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetFourHundredAndFifteenApplicationJSONObject() *UpdateFeatureTypeLifetimeUnstableResponse415ResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndFifteenApplicationJSONObject
}

func (o *UpdateFeatureTypeLifetimeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFeatureTypeLifetimeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFeatureTypeLifetimeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateFeatureTypeLifetimeResponse) GetFeatureTypeSchema() *shared.FeatureTypeSchema {
	if o == nil {
		return nil
	}
	return o.FeatureTypeSchema
}
