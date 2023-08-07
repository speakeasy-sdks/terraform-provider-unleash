// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// FeedbackResponseSchema - User feedback information about a particular feedback item.
type FeedbackResponseSchema struct {
	// The name of the feedback session
	FeedbackID *string `json:"feedbackId,omitempty"`
	// When this feedback was given
	Given *time.Time `json:"given,omitempty"`
	// `true` if the user has asked never to see this feedback questionnaire again.
	NeverShow *bool `json:"neverShow,omitempty"`
	// The ID of the user that gave the feedback.
	UserID *int64 `json:"userId,omitempty"`
}