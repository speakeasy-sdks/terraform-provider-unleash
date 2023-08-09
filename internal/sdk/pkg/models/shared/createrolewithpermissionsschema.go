// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreateRoleWithPermissionsSchemaType string

const (
	CreateRoleWithPermissionsSchemaTypeRootCustom CreateRoleWithPermissionsSchemaType = "root-custom"
	CreateRoleWithPermissionsSchemaTypeCustom     CreateRoleWithPermissionsSchemaType = "custom"
)

func (e CreateRoleWithPermissionsSchemaType) ToPointer() *CreateRoleWithPermissionsSchemaType {
	return &e
}

func (e *CreateRoleWithPermissionsSchemaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "root-custom":
		fallthrough
	case "custom":
		*e = CreateRoleWithPermissionsSchemaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRoleWithPermissionsSchemaType: %v", v)
	}
}

// CreateRoleWithPermissionsSchema - createRoleWithPermissionsSchema
type CreateRoleWithPermissionsSchema struct {
	Description *string                                           `json:"description,omitempty"`
	Name        string                                            `json:"name"`
	Permissions []CreateRoleWithPermissionsSchemaPermissionsInner `json:"permissions,omitempty"`
	Type        *CreateRoleWithPermissionsSchemaType              `json:"type,omitempty"`
}