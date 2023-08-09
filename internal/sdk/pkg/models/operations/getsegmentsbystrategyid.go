// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetSegmentsByStrategyIDRequest struct {
	StrategyID string `pathParam:"style=simple,explode=false,name=strategyId"`
}

type GetSegmentsByStrategyIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// segmentsSchema
	SegmentsSchema *shared.SegmentsSchema
}
