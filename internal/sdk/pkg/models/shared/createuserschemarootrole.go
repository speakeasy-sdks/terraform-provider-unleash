// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// CreateUserSchemaRootRole2 - The role to assign to the user. Can be either the role's ID or its unique name.
type CreateUserSchemaRootRole2 string

const (
	CreateUserSchemaRootRole2Admin  CreateUserSchemaRootRole2 = "Admin"
	CreateUserSchemaRootRole2Editor CreateUserSchemaRootRole2 = "Editor"
	CreateUserSchemaRootRole2Viewer CreateUserSchemaRootRole2 = "Viewer"
	CreateUserSchemaRootRole2Owner  CreateUserSchemaRootRole2 = "Owner"
	CreateUserSchemaRootRole2Member CreateUserSchemaRootRole2 = "Member"
)

func (e CreateUserSchemaRootRole2) ToPointer() *CreateUserSchemaRootRole2 {
	return &e
}

func (e *CreateUserSchemaRootRole2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Admin":
		fallthrough
	case "Editor":
		fallthrough
	case "Viewer":
		fallthrough
	case "Owner":
		fallthrough
	case "Member":
		*e = CreateUserSchemaRootRole2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateUserSchemaRootRole2: %v", v)
	}
}

type CreateUserSchemaRootRoleType string

const (
	CreateUserSchemaRootRoleTypeInteger                   CreateUserSchemaRootRoleType = "integer"
	CreateUserSchemaRootRoleTypeCreateUserSchemaRootRole2 CreateUserSchemaRootRoleType = "createUserSchema_rootRole_2"
)

type CreateUserSchemaRootRole struct {
	Integer                   *int64
	CreateUserSchemaRootRole2 *CreateUserSchemaRootRole2

	Type CreateUserSchemaRootRoleType
}

func CreateCreateUserSchemaRootRoleInteger(integer int64) CreateUserSchemaRootRole {
	typ := CreateUserSchemaRootRoleTypeInteger

	return CreateUserSchemaRootRole{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateCreateUserSchemaRootRoleCreateUserSchemaRootRole2(createUserSchemaRootRole2 CreateUserSchemaRootRole2) CreateUserSchemaRootRole {
	typ := CreateUserSchemaRootRoleTypeCreateUserSchemaRootRole2

	return CreateUserSchemaRootRole{
		CreateUserSchemaRootRole2: &createUserSchemaRootRole2,
		Type:                      typ,
	}
}

func (u *CreateUserSchemaRootRole) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	integer := new(int64)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&integer); err == nil {
		u.Integer = integer
		u.Type = CreateUserSchemaRootRoleTypeInteger
		return nil
	}

	createUserSchemaRootRole2 := new(CreateUserSchemaRootRole2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createUserSchemaRootRole2); err == nil {
		u.CreateUserSchemaRootRole2 = createUserSchemaRootRole2
		u.Type = CreateUserSchemaRootRoleTypeCreateUserSchemaRootRole2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateUserSchemaRootRole) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return json.Marshal(u.Integer)
	}

	if u.CreateUserSchemaRootRole2 != nil {
		return json.Marshal(u.CreateUserSchemaRootRole2)
	}

	return nil, nil
}
