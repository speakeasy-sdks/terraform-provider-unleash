// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ChangePasswordSchema - Change password as long as the token is a valid token
type ChangePasswordSchema struct {
	// The new password for the user
	Password string `json:"password"`
	// A reset token used to validate that the user is allowed to change the password.
	Token string `json:"token"`
}

func (o *ChangePasswordSchema) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ChangePasswordSchema) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}
