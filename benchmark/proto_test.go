package benchmark

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/elastic-agent-shipper-client/pkg/helpers"
	"github.com/elastic/elastic-agent-shipper-client/pkg/proto/messages"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapStrToEvent(m mapstr.M) (*messages.Event, error) {
	tsString, err := m.GetValue("@timestamp")
	if err != nil {
		return nil, fmt.Errorf("could not get @timestamp: %w", err)
	}
	err = m.Delete("@timestamp")
	if err != nil {
		return nil, fmt.Errorf("could not remove @timestamp")
	}
	ts, err := time.Parse(time.RFC3339, tsString.(string))
	if err != nil {
		return nil, fmt.Errorf("could not parse timestamp: %w", err)
	}

	metaMpstr, err := m.GetValue("@metadata")
	if err != nil {
		return nil, fmt.Errorf("could not get @metadata: %w", err)
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

	return &messages.Event{
		Timestamp: timestamppb.New(ts),
		Metadata:  meta.GetStructValue(),
		Fields:    fields.GetStructValue(),
		Source: &messages.Source{
			InputId:  "inputID",
			StreamId: "streamID",
		},
		DataStream: &messages.DataStream{
			Type:      "log",
			Namespace: "default",
			Dataset:   "generic",
		},
	}, nil
}

func ndjsonToEvents(filename string) ([]*messages.Event, error) {
	events := []*messages.Event{}
	fp := filepath.Join("testdata", filename)
	fileHandle, err := os.Open(fp)
	if err != nil {
		return events, fmt.Errorf("could not open file: %s: %w", fp, err)
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := mapstr.M{}
		if err := json.Unmarshal(scanner.Bytes(), &m); err != nil {
			return events, fmt.Errorf("could not unmarshal: %w", err)
		}
		e, err := mapStrToEvent(m)
		if err != nil {
			return events, fmt.Errorf("could not convert mapstr to event: %w", err)
		}
		events = append(events, e)
	}
	return events, nil
}

func ndjsonToMapStr(filename string) ([]*mapstr.M, error) {
	events := []*mapstr.M{}
	fp := filepath.Join("testdata", filename)
	fileHandle, err := os.Open(fp)
	if err != nil {
		return events, fmt.Errorf("could not open file: %s: %w", fp, err)
	}
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := mapstr.M{}
		if err := json.Unmarshal(scanner.Bytes(), &m); err != nil {
			return events, fmt.Errorf("could not unmarshal: %w", err)
		}
		events = append(events, &m)
	}
	return events, nil
}

func roundtripEvent(m *messages.Event) {
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

func roundtripMapStr(m *mapstr.M) {
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

func BenchmarkEventSingleField(b *testing.B) {
	events, err := ndjsonToEvents("message.ndjson")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range events {
			roundtripEvent(e)
		}
	}
}

func BenchmarkEventMultiField(b *testing.B) {
	events, err := ndjsonToEvents("small_json.ndjson")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range events {
			roundtripEvent(e)
		}
	}
}

func BenchmarkMapStrSingleField(b *testing.B) {
	events, err := ndjsonToMapStr("message.ndjson")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range events {
			roundtripMapStr(e)
		}
	}
}

func BenchmarkMapStrMultiField(b *testing.B) {
	events, err := ndjsonToMapStr("small_json.ndjson")
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range events {
			roundtripMapStr(e)
		}
	}
}
