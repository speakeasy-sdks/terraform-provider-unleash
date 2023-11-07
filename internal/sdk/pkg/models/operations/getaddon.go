// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetAddonRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAddonRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetAddonResponseBody - Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
type GetAddonResponseBody struct {
	// The ID of the error instance
	ID *string `json:"id,omitempty"`
	// A description of what went wrong.
	Message *string `json:"message,omitempty"`
	// The name of the error kind
	Name *string `json:"name,omitempty"`
}

func (o *GetAddonResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetAddonResponseBody) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetAddonResponseBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type GetAddonResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// addonSchema
	AddonSchema *shared.AddonSchema
	// Authorization information is missing or invalid. Provide a valid API token as the `authorization` header, e.g. `authorization:*.*.my-admin-token`.
	Object *GetAddonResponseBody
}

func (o *GetAddonResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAddonResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAddonResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAddonResponse) GetAddonSchema() *shared.AddonSchema {
	if o == nil {
		return nil
	}
	return o.AddonSchema
}

func (o *GetAddonResponse) GetObject() *GetAddonResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
