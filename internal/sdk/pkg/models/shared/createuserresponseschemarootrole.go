// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// CreateUserResponseSchemaRootRole2 - Which [root role](https://docs.getunleash.io/reference/rbac#standard-roles) this user is assigned. Usually a numeric role ID, but can be a string when returning newly created user with an explicit string role.
type CreateUserResponseSchemaRootRole2 string

const (
	CreateUserResponseSchemaRootRole2Admin  CreateUserResponseSchemaRootRole2 = "Admin"
	CreateUserResponseSchemaRootRole2Editor CreateUserResponseSchemaRootRole2 = "Editor"
	CreateUserResponseSchemaRootRole2Viewer CreateUserResponseSchemaRootRole2 = "Viewer"
	CreateUserResponseSchemaRootRole2Owner  CreateUserResponseSchemaRootRole2 = "Owner"
	CreateUserResponseSchemaRootRole2Member CreateUserResponseSchemaRootRole2 = "Member"
)

func (e CreateUserResponseSchemaRootRole2) ToPointer() *CreateUserResponseSchemaRootRole2 {
	return &e
}

func (e *CreateUserResponseSchemaRootRole2) UnmarshalJSON(data []byte) error {
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
		*e = CreateUserResponseSchemaRootRole2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateUserResponseSchemaRootRole2: %v", v)
	}
}

type CreateUserResponseSchemaRootRoleType string

const (
	CreateUserResponseSchemaRootRoleTypeInteger                           CreateUserResponseSchemaRootRoleType = "integer"
	CreateUserResponseSchemaRootRoleTypeCreateUserResponseSchemaRootRole2 CreateUserResponseSchemaRootRoleType = "createUserResponseSchema_rootRole_2"
)

type CreateUserResponseSchemaRootRole struct {
	Integer                           *int64
	CreateUserResponseSchemaRootRole2 *CreateUserResponseSchemaRootRole2

	Type CreateUserResponseSchemaRootRoleType
}

func CreateCreateUserResponseSchemaRootRoleInteger(integer int64) CreateUserResponseSchemaRootRole {
	typ := CreateUserResponseSchemaRootRoleTypeInteger

	return CreateUserResponseSchemaRootRole{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateCreateUserResponseSchemaRootRoleCreateUserResponseSchemaRootRole2(createUserResponseSchemaRootRole2 CreateUserResponseSchemaRootRole2) CreateUserResponseSchemaRootRole {
	typ := CreateUserResponseSchemaRootRoleTypeCreateUserResponseSchemaRootRole2

	return CreateUserResponseSchemaRootRole{
		CreateUserResponseSchemaRootRole2: &createUserResponseSchemaRootRole2,
		Type:                              typ,
	}
}

func (u *CreateUserResponseSchemaRootRole) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	integer := new(int64)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&integer); err == nil {
		u.Integer = integer
		u.Type = CreateUserResponseSchemaRootRoleTypeInteger
		return nil
	}

	createUserResponseSchemaRootRole2 := new(CreateUserResponseSchemaRootRole2)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&createUserResponseSchemaRootRole2); err == nil {
		u.CreateUserResponseSchemaRootRole2 = createUserResponseSchemaRootRole2
		u.Type = CreateUserResponseSchemaRootRoleTypeCreateUserResponseSchemaRootRole2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateUserResponseSchemaRootRole) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return json.Marshal(u.Integer)
	}

	if u.CreateUserResponseSchemaRootRole2 != nil {
		return json.Marshal(u.CreateUserResponseSchemaRootRole2)
	}

	return nil, nil
}