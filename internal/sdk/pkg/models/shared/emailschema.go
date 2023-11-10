// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EmailSchema - Represents the email of a user. Used to send email communication (reset password, welcome mail etc)
type EmailSchema struct {
	// The email address
	Email string `json:"email"`
}

func (o *EmailSchema) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}
