// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ChangeProjectSchema - Data required to move a feature toggle to a project.
type ChangeProjectSchema struct {
	// The project to move the feature toggle to.
	NewProjectID string `json:"newProjectId"`
}
