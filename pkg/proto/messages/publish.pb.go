// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: messages/publish.proto

package messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. If present, this request will only be accepted if the uuid
	// matches the current uuid of the shipper process. The uuid identifies
	// the current shipper process, and is updated when the shipper restarts.
	// Its current value is returned in every shipper API reply.
	// A uuid in a PublishRequest is used for enforcing at-least-once delivery
	// guarantees: inputs may include the known shipper uuid with their request,
	// ensuring it will be rejected if the shipper restarts. In this case the
	// input should rewind to the last known-good position in its data sequence.
	// Note that this issue only arises during error states, since Agent only
	// restarts the shipper when its process is terminated or nonresponsive.
	Uuid   string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_messages_publish_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PublishRequest) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// Event is a translation of beat.Event into protobuf.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Creation timestamp of the event.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Source of the generated event.
	Source *Source `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	// Data stream for the event.
	DataStream *DataStream `protobuf:"bytes,4,opt,name=data_stream,json=dataStream,proto3" json:"data_stream,omitempty"`
	// Metadata JSON object (map[string]google.protobuf.Value)
	Metadata *Struct `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Field JSON object (map[string]google.protobuf.Value)
	Fields *Struct `protobuf:"bytes,6,opt,name=fields,proto3" json:"fields,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_messages_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_messages_publish_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Event) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Event) GetDataStream() *DataStream {
	if x != nil {
		return x.DataStream
	}
	return nil
}

func (x *Event) GetMetadata() *Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Event) GetFields() *Struct {
	if x != nil {
		return x.Fields
	}
	return nil
}

// Source information required for proper event tracking, processing and routing
type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Input ID in the agent policy.
	InputId string `protobuf:"bytes,1,opt,name=input_id,json=inputId,proto3" json:"input_id,omitempty"`
	// Stream ID in the agent policy (Optional, some inputs don't use streams).
	// Not to be confused with data streams in Elasticsearch.
	StreamId string `protobuf:"bytes,2,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_publish_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_messages_publish_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_messages_publish_proto_rawDescGZIP(), []int{2}
}

func (x *Source) GetInputId() string {
	if x != nil {
		return x.InputId
	}
	return ""
}

func (x *Source) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

// Elastic data stream
// See https://www.elastic.co/blog/an-introduction-to-the-elastic-data-stream-naming-scheme
type DataStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Generic type describing the data
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Describes the data ingested and its structure
	Dataset string `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// User-configurable arbitrary grouping
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *DataStream) Reset() {
	*x = DataStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_publish_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStream) ProtoMessage() {}

func (x *DataStream) ProtoReflect() protoreflect.Message {
	mi := &file_messages_publish_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStream.ProtoReflect.Descriptor instead.
func (*DataStream) Descriptor() ([]byte, []int) {
	return file_messages_publish_proto_rawDescGZIP(), []int{3}
}

func (x *DataStream) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DataStream) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *DataStream) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type PublishReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uuid of the shipper process, generated on startup. Clients can use this
	// to detect when the shipper restarts.
	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// The number of events accepted by the shipper, in the same order as the
	// PublishRequest. If uuid in the reply differs from uuid in the request,
	// then accepted_count is always zero.
	AcceptedCount int32 `protobuf:"varint,2,opt,name=accepted_count,json=acceptedCount,proto3" json:"accepted_count,omitempty"`
	// The final internal index for the events that were accepted. Inputs that
	// want to guarantee event persistence can do it with this field: when the
	// persisted_index of a PublishReply or PersistedIndexReply is >= this value,
	// the events from this publish request have been persisted and the input can
	// safely advance. See the API README for details.
	AcceptedIndex int64 `protobuf:"varint,3,opt,name=accepted_index,json=acceptedIndex,proto3" json:"accepted_index,omitempty"`
	// The highest sequential index that has been persisted. (See the API
	// README for details on what "persisted" entails.)
	PersistedIndex int64 `protobuf:"varint,4,opt,name=persisted_index,json=persistedIndex,proto3" json:"persisted_index,omitempty"`
}

func (x *PublishReply) Reset() {
	*x = PublishReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_publish_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishReply) ProtoMessage() {}

func (x *PublishReply) ProtoReflect() protoreflect.Message {
	mi := &file_messages_publish_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishReply.ProtoReflect.Descriptor instead.
func (*PublishReply) Descriptor() ([]byte, []int) {
	return file_messages_publish_proto_rawDescGZIP(), []int{4}
}

func (x *PublishReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PublishReply) GetAcceptedCount() int32 {
	if x != nil {
		return x.AcceptedCount
	}
	return 0
}

func (x *PublishReply) GetAcceptedIndex() int64 {
	if x != nil {
		return x.AcceptedIndex
	}
	return 0
}

func (x *PublishReply) GetPersistedIndex() int64 {
	if x != nil {
		return x.PersistedIndex
	}
	return 0
}

var File_messages_publish_proto protoreflect.FileDescriptor

var file_messages_publish_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xde, 0x02, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x41, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x45, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6c, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x40, 0x0a, 0x06,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x58,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x0c, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_messages_publish_proto_rawDescOnce sync.Once
	file_messages_publish_proto_rawDescData = file_messages_publish_proto_rawDesc
)

func file_messages_publish_proto_rawDescGZIP() []byte {
	file_messages_publish_proto_rawDescOnce.Do(func() {
		file_messages_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_publish_proto_rawDescData)
	})
	return file_messages_publish_proto_rawDescData
}

var file_messages_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_messages_publish_proto_goTypes = []interface{}{
	(*PublishRequest)(nil),        // 0: elastic.agent.shipper.v1.messages.PublishRequest
	(*Event)(nil),                 // 1: elastic.agent.shipper.v1.messages.Event
	(*Source)(nil),                // 2: elastic.agent.shipper.v1.messages.Source
	(*DataStream)(nil),            // 3: elastic.agent.shipper.v1.messages.DataStream
	(*PublishReply)(nil),          // 4: elastic.agent.shipper.v1.messages.PublishReply
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Struct)(nil),                // 6: elastic.agent.shipper.v1.messages.Struct
}
var file_messages_publish_proto_depIdxs = []int32{
	1, // 0: elastic.agent.shipper.v1.messages.PublishRequest.events:type_name -> elastic.agent.shipper.v1.messages.Event
	5, // 1: elastic.agent.shipper.v1.messages.Event.timestamp:type_name -> google.protobuf.Timestamp
	2, // 2: elastic.agent.shipper.v1.messages.Event.source:type_name -> elastic.agent.shipper.v1.messages.Source
	3, // 3: elastic.agent.shipper.v1.messages.Event.data_stream:type_name -> elastic.agent.shipper.v1.messages.DataStream
	6, // 4: elastic.agent.shipper.v1.messages.Event.metadata:type_name -> elastic.agent.shipper.v1.messages.Struct
	6, // 5: elastic.agent.shipper.v1.messages.Event.fields:type_name -> elastic.agent.shipper.v1.messages.Struct
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_messages_publish_proto_init() }
func file_messages_publish_proto_init() {
	if File_messages_publish_proto != nil {
		return
	}
	file_messages_struct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_publish_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_publish_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStream); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_publish_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_publish_proto_goTypes,
		DependencyIndexes: file_messages_publish_proto_depIdxs,
		MessageInfos:      file_messages_publish_proto_msgTypes,
	}.Build()
	File_messages_publish_proto = out.File
	file_messages_publish_proto_rawDesc = nil
	file_messages_publish_proto_goTypes = nil
	file_messages_publish_proto_depIdxs = nil
}
