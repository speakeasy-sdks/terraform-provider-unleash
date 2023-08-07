// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// FeedbackCreateSchema - User feedback information to be created.
type FeedbackCreateSchema struct {
	// The name of the feedback session
	FeedbackID string `json:"feedbackId"`
	// `true` if the user has asked never to see this feedback questionnaire again. Defaults to `false`.
	NeverShow *bool `json:"neverShow,omitempty"`
}