// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PatchSchemaOp string

const (
	PatchSchemaOpAdd     PatchSchemaOp = "add"
	PatchSchemaOpRemove  PatchSchemaOp = "remove"
	PatchSchemaOpReplace PatchSchemaOp = "replace"
	PatchSchemaOpCopy    PatchSchemaOp = "copy"
	PatchSchemaOpMove    PatchSchemaOp = "move"
)

func (e PatchSchemaOp) ToPointer() *PatchSchemaOp {
	return &e
}

func (e *PatchSchemaOp) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "add":
		fallthrough
	case "remove":
		fallthrough
	case "replace":
		fallthrough
	case "copy":
		fallthrough
	case "move":
		*e = PatchSchemaOp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PatchSchemaOp: %v", v)
	}
}

type PatchSchema struct {
	From  *string       `json:"from,omitempty"`
	Op    PatchSchemaOp `json:"op"`
	Path  string        `json:"path"`
	Value interface{}   `json:"value,omitempty"`
}
