// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UsersGroupsBaseSchema struct {
	Groups []GroupSchema `json:"groups,omitempty"`
	Users  []UserSchema  `json:"users,omitempty"`
}

func (o *UsersGroupsBaseSchema) GetGroups() []GroupSchema {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *UsersGroupsBaseSchema) GetUsers() []UserSchema {
	if o == nil {
		return nil
	}
	return o.Users
}
