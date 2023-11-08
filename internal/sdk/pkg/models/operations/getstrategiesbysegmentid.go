// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetStrategiesBySegmentIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetStrategiesBySegmentIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetStrategiesBySegmentIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// adminStrategiesSchema
	AdminStrategiesSchema *shared.AdminStrategiesSchema
}

func (o *GetStrategiesBySegmentIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetStrategiesBySegmentIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetStrategiesBySegmentIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetStrategiesBySegmentIDResponse) GetAdminStrategiesSchema() *shared.AdminStrategiesSchema {
	if o == nil {
		return nil
	}
	return o.AdminStrategiesSchema
}
