// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type ExportStrategiesParameterType string

const (
	ExportStrategiesParameterTypeBoolean ExportStrategiesParameterType = "boolean"
	ExportStrategiesParameterTypeStr     ExportStrategiesParameterType = "str"
	ExportStrategiesParameterTypeNumber  ExportStrategiesParameterType = "number"
)

type ExportStrategiesParameter struct {
	Boolean *bool
	Str     *string
	Number  *float64

	Type ExportStrategiesParameterType
}

func CreateExportStrategiesParameterBoolean(boolean bool) ExportStrategiesParameter {
	typ := ExportStrategiesParameterTypeBoolean

	return ExportStrategiesParameter{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateExportStrategiesParameterStr(str string) ExportStrategiesParameter {
	typ := ExportStrategiesParameterTypeStr

	return ExportStrategiesParameter{
		Str:  &str,
		Type: typ,
	}
}

func CreateExportStrategiesParameterNumber(number float64) ExportStrategiesParameter {
	typ := ExportStrategiesParameterTypeNumber

	return ExportStrategiesParameter{
		Number: &number,
		Type:   typ,
	}
}

func (u *ExportStrategiesParameter) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	boolean := new(bool)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&boolean); err == nil {
		u.Boolean = boolean
		u.Type = ExportStrategiesParameterTypeBoolean
		return nil
	}

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = ExportStrategiesParameterTypeStr
		return nil
	}

	number := new(float64)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&number); err == nil {
		u.Number = number
		u.Type = ExportStrategiesParameterTypeNumber
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ExportStrategiesParameter) MarshalJSON() ([]byte, error) {
	if u.Boolean != nil {
		return json.Marshal(u.Boolean)
	}

	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.Number != nil {
		return json.Marshal(u.Number)
	}

	return nil, nil
}