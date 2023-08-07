// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ProjectsSchema - An overview of all the projects in the Unleash instance
type ProjectsSchema struct {
	// A list of projects in the Unleash instance
	Projects []ProjectSchema `json:"projects"`
	Version  int64           `json:"version"`
}