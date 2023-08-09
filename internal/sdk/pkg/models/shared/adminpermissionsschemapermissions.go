// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AdminPermissionsSchemaPermissions - Returns permissions available at all three levels (root|project|environment)
type AdminPermissionsSchemaPermissions struct {
	// A list of environments with available permissions per environment
	Environments []AdminPermissionsSchemaPermissionsEnvironmentsInner `json:"environments"`
	// Permissions available at the project level
	Project []AdminPermissionSchema `json:"project"`
	// Permissions available at the root level, i.e. not connected to any specific project or environment
	Root []AdminPermissionSchema `json:"root,omitempty"`
}