// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package messages

import (
	"encoding/json"
	"fmt"
)

// MarshalJSON implements the JSON interface for the value type
func (val *Value) MarshalJSON() ([]byte, error) {
	switch typ := val.GetKind().(type) {
	case *Value_NullValue:
		return []byte("null"), nil
	case *Value_NumberValue:
		return json.Marshal(typ.NumberValue)
	case *Value_StringValue:
		return json.Marshal(typ.StringValue)
	case *Value_BoolValue:
		if typ.BoolValue {
			return []byte("true"), nil
		}
		return []byte("false"), nil
	case *Value_StructValue:
		data, err := typ.StructValue.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("error marshaling within value: %w", err)
		}
		return data, nil
	case *Value_ListValue:
		data, err := typ.ListValue.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("error marshaling within value: %w", err)
		}
		return data, nil
	case *Value_TimestampValue:
		data, err := typ.TimestampValue.AsTime().MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("error marshaling timestamp within value: %w", err)
		}
		return data, nil
	default:
		return nil, fmt.Errorf("Unknown type %T in event", typ)
	}
}

// MarshalJSON implements the JSON interface for the struct type
func (sv *Struct) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(sv.GetData())
	if err != nil {
		return nil, fmt.Errorf("error marshaling struct type: %w", err)
	}
	return data, nil
}

// MarshalJSON implements the JSON interface for the list Value type
func (lv *ListValue) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(lv.GetValues())
	if err != nil {
		return nil, fmt.Errorf("error marshaling list type: %w", err)
	}
	return data, nil

}
