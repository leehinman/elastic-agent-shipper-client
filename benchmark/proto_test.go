// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package benchmark

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/elastic-agent-shipper-client/pkg/helpers"
	"github.com/elastic/elastic-agent-shipper-client/pkg/proto/messages"
	"github.com/elastic/go-structform/cborl"
	"github.com/elastic/go-structform/gotype"
	fxamacker "github.com/fxamacker/cbor/v2"
	goccy "github.com/goccy/go-json"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ShallowEvent struct {
	Timestamp  string     `json:"timestamp"`
	Source     Source     `json:"source"`
	DataStream DataStream `json:"data_stream"`
	Metadata   []byte     `json:"metadata"`
	Fields     []byte     `json:"fields"`
}

type DeepEvent struct {
	Timestamp  string     `json:"timestamp"`
	Source     Source     `json:"source"`
	DataStream DataStream `json:"data_stream"`
	Metadata   mapstr.M   `json:"metadata"`
	Fields     mapstr.M   `json:"fields"`
}

type Source struct {
	InputId  string `json:"input_id"`
	StreamId string `json:"stream_id"`
}
type DataStream struct {
	Type      string `json:"type"`
	Dataset   string `json:"dataset"`
	Namespace string `json:"namespace"`
}

func readNdjson(filename string) ([][]byte, error) {
	out := [][]byte{}
	fp := filepath.Join("testdata", filename)
	fileHandle, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %s: %w", fp, err)
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tmp := scanner.Bytes()
		tmp2 := make([]byte, len(tmp))
		copy(tmp2, tmp)
		out = append(out, tmp2)
	}
	return out, nil
}

func bytesToMessagesEvents(input [][]byte) ([]*messages.Event, error) {
	events := []*messages.Event{}
	for _, raw := range input {
		m := mapstr.M{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return events, fmt.Errorf("could not unmarshal: %w: %v", err, string(raw))
		}
		event := messages.Event{}
		tsString, err := m.GetValue("@timestamp")
		if err != nil {
			return events, fmt.Errorf("could not get @timestamp: %w", err)
		}
		err = m.Delete("@timestamp")
		if err != nil {
			return events, fmt.Errorf("could not remove @timestamp")
		}
		ts, err := time.Parse(time.RFC3339, tsString.(string))
		if err != nil {
			return events, fmt.Errorf("could not parse timestamp: %w", err)
		}
		metaMpstr, err := m.GetValue("@metadata")
		if err != nil {
			return events, fmt.Errorf("could not get @metadata: %w", err)
		}
		meta, err := helpers.NewValue(metaMpstr)
		if err != nil {
			return nil, fmt.Errorf("failed to conver metadata to protobuf: %w", err)
		}
		err = m.Delete("@metadata")
		if err != nil {
			return nil, fmt.Errorf("failed to delete @metadata: %w", err)
		}
		fields, err := helpers.NewValue(m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert fields to protobuf: %w", err)
		}

		event.Timestamp = timestamppb.New(ts)
		event.Metadata = meta.GetStructValue()
		event.Fields = fields.GetStructValue()
		event.Source = &messages.Source{
			InputId:  "inputID",
			StreamId: "streamID",
		}
		event.DataStream = &messages.DataStream{
			Type:      "log",
			Namespace: "default",
			Dataset:   "generic",
		}
		events = append(events, &event)
	}
	return events, nil
}

func rtMessagesEvent(m *messages.Event) {
	b, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	new := messages.Event{}
	err = proto.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func bytesToMapStr(input [][]byte) ([]*mapstr.M, error) {
	events := []*mapstr.M{}
	for _, raw := range input {
		m := mapstr.M{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return events, fmt.Errorf("could not unmarshal: %w: %v", err, string(raw))
		}
		events = append(events, &m)
	}
	return events, nil
}

func rtMapStrStdJSON(m *mapstr.M) {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	new := mapstr.M{}
	err = json.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func bytesToMessagesShallowEvents(input [][]byte) ([]*messages.ShallowEvent, error) {
	events := []*messages.ShallowEvent{}
	for _, raw := range input {
		m := mapstr.M{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return events, fmt.Errorf("could not unmarshal: %w: %v", err, string(raw))
		}
		event := messages.ShallowEvent{}
		tsString, err := m.GetValue("@timestamp")
		if err != nil {
			return events, fmt.Errorf("could not get @timestamp: %w", err)
		}
		err = m.Delete("@timestamp")
		if err != nil {
			return events, fmt.Errorf("could not remove @timestamp")
		}
		ts, err := time.Parse(time.RFC3339, tsString.(string))
		if err != nil {
			return events, fmt.Errorf("could not parse timestamp: %w", err)
		}
		metaMpstr, err := m.GetValue("@metadata")
		if err != nil {
			return events, fmt.Errorf("could not get @metadata: %w", err)
		}
		err = m.Delete("@metadata")
		metaBytes, err := json.Marshal(metaMpstr)
		if err != nil {
			return events, fmt.Errorf("could not marshal meta map string: %w", err)
		}
		fieldsBytes, err := json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("could not marshal fields map string: %w", err)
		}
		event.Timestamp = timestamppb.New(ts)
		event.Metadata = metaBytes
		event.Fields = fieldsBytes
		event.Source = &messages.Source{
			InputId:  "inputID",
			StreamId: "streamID",
		}
		event.DataStream = &messages.DataStream{
			Type:      "log",
			Namespace: "default",
			Dataset:   "generic",
		}
		events = append(events, &event)
	}
	return events, nil
}

func rtMessagesShallowEvent(m *messages.ShallowEvent) {
	b, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	new := messages.ShallowEvent{}
	err = proto.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func rtMessagesShallowEventFull(m *messages.ShallowEvent) {
	b, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	new := messages.ShallowEvent{}
	err = proto.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
	fields := mapstr.M{}
	err = json.Unmarshal([]byte(new.Fields), &fields)
	if err != nil {
		panic(err)
	}
	b, err = json.Marshal(fields)
	if err != nil {
		panic(err)
	}
	meta := mapstr.M{}
	err = json.Unmarshal([]byte(new.Metadata), &meta)
	if err != nil {
		panic(err)
	}
	b, err = json.Marshal(meta)
	if err != nil {
		panic(err)
	}
}

func rtMessagesShallowEventFullGoJSON(m *messages.ShallowEvent) {
	b, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	new := messages.ShallowEvent{}
	err = proto.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
	fields := mapstr.M{}
	err = goccy.Unmarshal([]byte(new.Fields), &fields)
	if err != nil {
		panic(err)
	}
	b, err = goccy.Marshal(fields)
	if err != nil {
		panic(err)
	}
	meta := mapstr.M{}
	err = goccy.Unmarshal([]byte(new.Metadata), &meta)
	if err != nil {
		panic(err)
	}
	b, err = goccy.Marshal(meta)
	if err != nil {
		panic(err)
	}
}

func rtMapStrGoJSON(m *mapstr.M) {
	b, err := goccy.Marshal(m)
	if err != nil {
		panic(err)
	}
	new := mapstr.M{}
	err = goccy.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func bytesToShallowEvents(input [][]byte) ([]*ShallowEvent, error) {
	events := []*ShallowEvent{}
	for _, raw := range input {
		m := mapstr.M{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return events, fmt.Errorf("could not unmarshal: %w: %v", err, string(raw))
		}
		event := ShallowEvent{}
		tsString, err := m.GetValue("@timestamp")
		if err != nil {
			return events, fmt.Errorf("could not get @timestamp: %w", err)
		}
		err = m.Delete("@timestamp")
		if err != nil {
			return events, fmt.Errorf("could not remove @timestamp")
		}
		metaMpstr, err := m.GetValue("@metadata")
		if err != nil {
			return events, fmt.Errorf("could not get @metadata: %w", err)
		}
		err = m.Delete("@metadata")
		metaBytes, err := json.Marshal(metaMpstr)
		if err != nil {
			return events, fmt.Errorf("could not marshal meta map string: %w", err)
		}
		fieldsBytes, err := json.Marshal(m)
		if err != nil {
			return nil, fmt.Errorf("could not marshal fields map string: %w", err)
		}
		event.Timestamp = tsString.(string)
		event.Metadata = metaBytes
		event.Fields = fieldsBytes
		event.Source = Source{
			InputId:  "inputID",
			StreamId: "streamID",
		}
		event.DataStream = DataStream{
			Type:      "log",
			Namespace: "default",
			Dataset:   "generic",
		}
		events = append(events, &event)
	}
	return events, nil
}

func rtShallowEventStdJSON(e *ShallowEvent) {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	new := ShallowEvent{}
	err = json.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func rtShallowEventGoJSON(e *ShallowEvent) {
	b, err := goccy.Marshal(e)
	if err != nil {
		panic(err)
	}
	new := ShallowEvent{}
	err = goccy.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func bytesToDeepEvents(input [][]byte) ([]*DeepEvent, error) {
	events := []*DeepEvent{}
	for _, raw := range input {
		m := mapstr.M{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return events, fmt.Errorf("could not unmarshal: %w: %v", err, string(raw))
		}
		event := DeepEvent{}
		tsString, err := m.GetValue("@timestamp")
		if err != nil {
			return events, fmt.Errorf("could not get @timestamp: %w", err)
		}
		err = m.Delete("@timestamp")
		if err != nil {
			return events, fmt.Errorf("could not remove @timestamp")
		}
		metaMpstr, err := m.GetValue("@metadata")
		if err != nil {
			return events, fmt.Errorf("could not get @metadata: %w", err)
		}
		err = m.Delete("@metadata")

		event.Timestamp = tsString.(string)
		event.Metadata = metaMpstr.(map[string]interface{})
		event.Fields = m
		event.Source = Source{
			InputId:  "inputID",
			StreamId: "streamID",
		}
		event.DataStream = DataStream{
			Type:      "log",
			Namespace: "default",
			Dataset:   "generic",
		}
		events = append(events, &event)
	}
	return events, nil
}

func rtDeepEventStdJSON(e *DeepEvent) {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	new := DeepEvent{}
	err = json.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func rtDeepEventStructFormCBORL(e *DeepEvent) {
	var buf bytes.Buffer
	visitor := cborl.NewVisitor(&buf)
	folder, _ := gotype.NewIterator(visitor)
	folder.Fold(*e)

	unfolder, _ := gotype.NewUnfolder(nil)
	parser := cborl.NewParser(unfolder)
	new := DeepEvent{}
	err := unfolder.SetTarget(&new)
	if err != nil {
		panic(err)
	}
	parser.Parse(buf.Bytes())
}

func rtDeepEventGoJSON(e *DeepEvent) {
	b, err := goccy.Marshal(e)
	if err != nil {
		panic(err)
	}
	new := DeepEvent{}
	err = goccy.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func rtDeepEventCBORL(e *DeepEvent) {
	b, err := fxamacker.Marshal(e)
	if err != nil {
		panic(err)
	}
	new := DeepEvent{}
	err = fxamacker.Unmarshal(b, &new)
	if err != nil {
		panic(err)
	}
}

func BenchmarkMarshalUnmarshal(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
	}{
		{"Single", "message.ndjson"},
		{"Multi", "small_json.ndjson"},
	}
	for _, bm := range benchmarks {
		rawBytes, err := readNdjson(bm.filename)
		if err != nil {
			panic(err)
		}
		b.Run(bm.name+"OriginalProtobuf", func(b *testing.B) {
			events, err := bytesToMessagesEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtMessagesEvent(e)
				}
			}
		})
		b.Run(bm.name+"MapStrStdJSON", func(b *testing.B) {
			events, err := bytesToMapStr(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtMapStrStdJSON(e)
				}
			}
		})
		b.Run(bm.name+"ShallowProtobuf", func(b *testing.B) {
			events, err := bytesToMessagesShallowEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtMessagesShallowEvent(e)
				}
			}
		})
		b.Run(bm.name+"ShallowProtobufFull", func(b *testing.B) {
			events, err := bytesToMessagesShallowEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtMessagesShallowEventFull(e)
				}
			}
		})
		b.Run(bm.name+"ShallowProtobufFullGoJSON", func(b *testing.B) {
			events, err := bytesToMessagesShallowEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtMessagesShallowEventFullGoJSON(e)
				}
			}
		})
		b.Run(bm.name+"MapStrGoJSON", func(b *testing.B) {
			events, err := bytesToMapStr(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtMapStrGoJSON(e)
				}
			}
		})
		b.Run(bm.name+"ShallowEventStdJSON", func(b *testing.B) {
			events, err := bytesToShallowEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtShallowEventStdJSON(e)
				}
			}
		})
		b.Run(bm.name+"ShallowEventGoJSON", func(b *testing.B) {
			events, err := bytesToShallowEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtShallowEventGoJSON(e)
				}
			}
		})
		b.Run(bm.name+"DeepEventStdJSON", func(b *testing.B) {
			events, err := bytesToDeepEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtDeepEventStdJSON(e)
				}
			}
		})
		b.Run(bm.name+"DeepEventSructFormCBORL", func(b *testing.B) {
			events, err := bytesToDeepEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtDeepEventStructFormCBORL(e)
				}
			}
		})
		b.Run(bm.name+"DeepEventGoJSON", func(b *testing.B) {
			events, err := bytesToDeepEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtDeepEventGoJSON(e)
				}
			}
		})
		b.Run(bm.name+"DeepEventCBORL", func(b *testing.B) {
			events, err := bytesToDeepEvents(rawBytes)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, e := range events {
					rtDeepEventCBORL(e)
				}
			}
		})
	}
}
