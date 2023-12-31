// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SplashResponseSchema - Data related to a user having seen a splash screen.
type SplashResponseSchema struct {
	// Indicates whether the user has seen the splash screen or not.
	Seen bool `json:"seen"`
	// The ID of the splash screen that was shown.
	SplashID string `json:"splashId"`
	// The ID of the user that was shown the splash screen.
	UserID int64 `json:"userId"`
}

func (o *SplashResponseSchema) GetSeen() bool {
	if o == nil {
		return false
	}
	return o.Seen
}

func (o *SplashResponseSchema) GetSplashID() string {
	if o == nil {
		return ""
	}
	return o.SplashID
}

func (o *SplashResponseSchema) GetUserID() int64 {
	if o == nil {
		return 0
	}
	return o.UserID
}
