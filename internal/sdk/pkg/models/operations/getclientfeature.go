// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetClientFeatureRequest struct {
	FeatureName string `pathParam:"style=simple,explode=false,name=featureName"`
}

type GetClientFeatureResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// clientFeatureSchema
	ClientFeatureSchema *shared.ClientFeatureSchema
}
