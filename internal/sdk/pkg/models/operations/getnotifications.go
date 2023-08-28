// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"terraform/internal/sdk/pkg/models/shared"
)

type GetNotificationsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// notificationsSchema
	NotificationsSchema []shared.NotificationsSchema
}
