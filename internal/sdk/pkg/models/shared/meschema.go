// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MeSchema - Detailed user information
type MeSchema struct {
	// User feedback information
	Feedback []FeedbackResponseSchema `json:"feedback"`
	// User permissions for projects and environments
	Permissions []PermissionSchema `json:"permissions"`
	// Splash screen configuration
	Splash map[string]bool `json:"splash"`
	// An Unleash user
	User UserSchema `json:"user"`
}

func (o *MeSchema) GetFeedback() []FeedbackResponseSchema {
	if o == nil {
		return []FeedbackResponseSchema{}
	}
	return o.Feedback
}

func (o *MeSchema) GetPermissions() []PermissionSchema {
	if o == nil {
		return []PermissionSchema{}
	}
	return o.Permissions
}

func (o *MeSchema) GetSplash() map[string]bool {
	if o == nil {
		return map[string]bool{}
	}
	return o.Splash
}

func (o *MeSchema) GetUser() UserSchema {
	if o == nil {
		return UserSchema{}
	}
	return o.User
}
