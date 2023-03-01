package helpers

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"time"

	"testing"

	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/elastic-agent-shipper-client/pkg/proto/messages"
	"github.com/stretchr/testify/require"
	"go.elastic.co/fastjson"
)

var marshalResult = []byte{}

func BenchmarkCustomMarshal(b *testing.B) {
	testMapInput := mapstr.M{
		"StrTest":  "teststr",
		"Uint1":    32,
		"Uint2":    556,
		"Float1":   23.0,
		"Float2":   25.343564,
		"TestNil":  nil,
		"TestBool": false,
		"TestList": []interface{}{"strval", 5, false},
		"TestMap":  map[string]string{"testkey": "val"},
		"TestMapStr": mapstr.M{"key1": 5, "key2": "strval", "keymap": mapstr.M{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}},
	}

	testMessage, err := NewValue(testMapInput)
	if err != nil {
		b.Logf("error creating value from struct: %s", err)
		b.FailNow()
	}
	b.ResetTimer()
	b.Run("marshal custom protobuf message.Value type", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsonEventData, err := json.Marshal(testMessage)
			if err != nil {
				b.Logf("error marshaling data: %s", err)
				b.FailNow()
			}
			marshalResult = jsonEventData
		}
	})
	b.Run("standard struct using stdlib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			jsonEventData, err := json.Marshal(testMapInput)
			if err != nil {
				b.Logf("error marshaling data: %s", err)
				b.FailNow()
			}
			marshalResult = jsonEventData
		}
	})
}

func TestAsInterface(t *testing.T) {
	cases := []struct {
		name string
		in   *messages.Value
		exp  interface{}
	}{
		{
			name: "test map conversion",
			in: NewStructValue(&messages.Struct{Data: map[string]*messages.Value{
				"field1": NewStructValue(&messages.Struct{Data: map[string]*messages.Value{
					"value":     NewBoolValue(false),
					"value-str": NewStringValue("test"),
				}}),
			}}),
			exp: map[string]interface{}{
				"field1": map[string]interface{}{
					"value":     false,
					"value-str": "test",
				},
			},
		},
		{
			name: "test list",
			in: NewListValue(&messages.ListValue{Values: []*messages.Value{
				NewFloat32Value(45.3),
				NewStringValue("testStr"),
				NewInt64Value(4590),
			}}),
			exp: []interface{}{
				float32(45.3),
				"testStr",
				int64(4590),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := AsInterface(c.in)
			require.Equal(t, c.exp, res)
		})
	}
}

func TestJSONMarshal(t *testing.T) {
	ts := time.Now().UTC() // the timestamppb used by protobuf will remove the location, so set UTC to make DeepEqual happy
	testMapInput := mapstr.M{
		"@timestamp":    ts,
		"StrTest":       "test",
		"StrTestEscape": `"test_with_quotes"`,
		"Uint1":         uint32(32),
		"Uint2":         uint64(556),
		"Float1":        float32(23.0),
		"Float2":        25.343564,
		"TestNil":       nil,
		"TestBool":      false,
		"TestList":      []interface{}{"strval", 5, false},
		"TestMap":       map[string]string{"testkey": "val"},
		"testMapNested": map[string]interface{}{
			"testlevel1": map[string]interface{}{
				"testLevel2": map[string]interface{}{
					"testlevel3": map[string]int{
						"val":       4,
						"otherval":  10,
						"otherval2": 3425543,
					},
				},
			},
		},
		"TestMapStr": mapstr.M{"key1": 5, "key2": "strval", "keymap": mapstr.M{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}},
	}

	testMessage, err := NewValue(testMapInput)
	require.NoError(t, err)

	jsonWriter := &fastjson.Writer{}
	err = fastjson.Marshal(jsonWriter, testMessage)
	t.Logf("GOT: %s", string(jsonWriter.Bytes()))
	require.NoError(t, err)

	stdJSON, err := json.Marshal(testMapInput)
	require.NoError(t, err)
	t.Logf("EXPECTED: %s", string(stdJSON))

	// JSON string outputs aren't guaranteed to be deterministic, so unmarshal back to map so we can compare, test for JSON errors
	unmarshaledEvent := mapstr.M{}
	err = json.Unmarshal(jsonWriter.Bytes(), &unmarshaledEvent)
	require.NoError(t, err)
	t.Logf("got     : %s", unmarshaledEvent.StringToPrint())

	unmarshaledJSON := mapstr.M{}
	err = json.Unmarshal(stdJSON, &unmarshaledJSON)
	require.NoError(t, err)
	t.Logf("expected: %s", unmarshaledJSON.StringToPrint())

	if !reflect.DeepEqual(unmarshaledJSON, unmarshaledEvent) {
		t.Fatalf("events are different")
	}
}

var result *messages.Value

func BenchmarkCustomUnmarshal(b *testing.B) {
	testStructType := struct {
		A int
		B string
	}{
		A: 5,
		B: "test",
	}
	testList := mapstr.M{
		"field1": mapstr.M{
			"value":          3,
			"value-str":      "test",
			"value-list":     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			"value-list-str": []string{"value1", "value2", "value3", "valu4", "value5", "value6"},
			"value-struct":   testStructType,
			"value-bytes":    []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
			"value-time":     time.Now(),
			"value-map": map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
			"value-mapstr-nested": mapstr.M{
				"value1": 34.5,
				"value2": 4.5,
				"othermap": mapstr.M{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"other": mapstr.M{
						"key1": 1,
						"key2": 455.6,
						"key3": 987435.434,
					},
				},
			},
		},
	}
	var r *messages.Value
	var err error
	// benchmark performance for the map, which will usually be the most complex type to marshal
	// there's a handful of different possible ways to handle the NewValue conversion, so it's helpful
	// to have a benchmark in case we decide to adjust this in the future
	for i := 0; i < b.N; i++ {
		r, err = NewValue(testList)
		if err != nil {
			b.Logf("error: %s", err)
			b.FailNow()
		}
		result = r
	}
}

func TestStructValue(t *testing.T) {
	testStructType := struct {
		A int
		B string
	}{
		A: 5,
		B: "test",
	}
	ts := time.Now()
	cases := []struct {
		name string
		in   interface{}
		exp  *messages.Value
	}{
		{
			name: "string conversion",
			in:   "test-string",
			exp:  &messages.Value{Kind: &messages.Value_StringValue{StringValue: "test-string"}},
		},
		{
			name: "int64 conversion",
			in:   int64(32),
			exp:  &messages.Value{Kind: &messages.Value_Int64Value{Int64Value: 32}},
		},
		{
			name: "int32 conversion",
			in:   int32(32),
			exp:  &messages.Value{Kind: &messages.Value_Int32Value{Int32Value: 32}},
		},
		{
			name: "uint64 conversion",
			in:   uint64(32),
			exp:  &messages.Value{Kind: &messages.Value_Uint64Value{Uint64Value: 32}},
		},
		{
			name: "uint conversion",
			in:   uint(32),
			exp:  &messages.Value{Kind: &messages.Value_Uint64Value{Uint64Value: 32}},
		},
		{
			name: "float32 conversion",
			in:   float32(32.5),
			exp:  &messages.Value{Kind: &messages.Value_Float32Value{Float32Value: 32.5}},
		},
		{
			name: "float64 conversion",
			in:   float64(32.5),
			exp:  &messages.Value{Kind: &messages.Value_Float64Value{Float64Value: 32.5}},
		},
		{
			name: "int conversion",
			in:   32,
			exp:  &messages.Value{Kind: &messages.Value_Int64Value{Int64Value: 32}},
		},
		{
			name: "nil value",
			in:   nil,
			exp:  &messages.Value{Kind: &messages.Value_NullValue{NullValue: messages.NullValue_NULL_VALUE}},
		},
		{
			name: "test map conversion",
			in: mapstr.M{
				"field1": map[string]interface{}{
					"value":      false,
					"@timestamp": ts,
				},
			},
			exp: NewStructValue(&messages.Struct{Data: map[string]*messages.Value{
				"field1": NewStructValue(&messages.Struct{Data: map[string]*messages.Value{
					"value":      NewBoolValue(false),
					"@timestamp": NewTimestampValue(ts),
				}}),
			}}),
		},
		{
			name: "test struct conversion",
			in:   testStructType,
			exp: NewStructValue(&messages.Struct{Data: map[string]*messages.Value{
				"A": NewInt64Value(5),
				"B": NewStringValue("test"),
			}}),
		},
		{
			name: "list conversion of string type",
			in:   []string{"value1", "value2"},
			exp: NewListValue(&messages.ListValue{Values: []*messages.Value{
				NewStringValue("value1"),
				NewStringValue("value2"),
			}}),
		},
		{
			name: "test list of type int",
			in:   []uint32{45, 56, 7343, 3242, 5673},
			exp: NewListValue(&messages.ListValue{Values: []*messages.Value{
				NewUint32Value(45),
				NewUint32Value(56),
				NewUint32Value(7343),
				NewUint32Value(3242),
				NewUint32Value(5673),
			}}),
		},
		{
			name: "list conversion of interface type",
			in:   []interface{}{"value1", 3},
			exp: NewListValue(&messages.ListValue{Values: []*messages.Value{
				NewStringValue("value1"),
				NewInt64Value(3),
			}}),
		},
		{
			name: "proper handling of byte arrays",
			in:   []byte{0xFF, 0xFF},
			exp:  NewStringValue(base64.StdEncoding.EncodeToString([]byte{0xFF, 0xFF})),
		},
		{
			name: "proper handling of timestamps",
			in:   ts,
			exp:  NewTimestampValue(ts),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res, err := NewValue(c.in)
			require.NoError(t, err)
			require.Equal(t, c.exp, res)
		})
	}
}
