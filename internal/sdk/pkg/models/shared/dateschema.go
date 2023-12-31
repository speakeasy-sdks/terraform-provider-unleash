// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/pkg/utils"
	"time"
)

type DateSchemaType string

const (
	DateSchemaTypeDateTime DateSchemaType = "date-time"
	DateSchemaTypeNumber   DateSchemaType = "number"
)

type DateSchema struct {
	DateTime *time.Time
	Number   *float64

	Type DateSchemaType
}

func CreateDateSchemaDateTime(dateTime time.Time) DateSchema {
	typ := DateSchemaTypeDateTime

	return DateSchema{
		DateTime: &dateTime,
		Type:     typ,
	}
}

func CreateDateSchemaNumber(number float64) DateSchema {
	typ := DateSchemaTypeNumber

	return DateSchema{
		Number: &number,
		Type:   typ,
	}
}

func (u *DateSchema) UnmarshalJSON(data []byte) error {

	dateTime := new(time.Time)
	if err := utils.UnmarshalJSON(data, &dateTime, "", true, true); err == nil {
		u.DateTime = dateTime
		u.Type = DateSchemaTypeDateTime
		return nil
	}

	number := new(float64)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = number
		u.Type = DateSchemaTypeNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DateSchema) MarshalJSON() ([]byte, error) {
	if u.DateTime != nil {
		return utils.MarshalJSON(u.DateTime, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
