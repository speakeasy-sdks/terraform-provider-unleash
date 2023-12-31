// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetHealthResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// healthCheckSchema
	HealthCheckSchema *shared.HealthCheckSchema
}

func (o *GetHealthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHealthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHealthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetHealthResponse) GetHealthCheckSchema() *shared.HealthCheckSchema {
	if o == nil {
		return nil
	}
	return o.HealthCheckSchema
}
