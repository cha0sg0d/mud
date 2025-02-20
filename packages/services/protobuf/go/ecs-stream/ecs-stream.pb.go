// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.3
// source: proto/ecs-stream.proto

package ecs_stream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ECSEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType   string `protobuf:"bytes,1,opt,name=eventType,proto3" json:"eventType,omitempty"`
	ComponentId string `protobuf:"bytes,2,opt,name=componentId,proto3" json:"componentId,omitempty"`
	EntityId    string `protobuf:"bytes,3,opt,name=entityId,proto3" json:"entityId,omitempty"`
	Value       []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Tx          string `protobuf:"bytes,5,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *ECSEvent) Reset() {
	*x = ECSEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECSEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECSEvent) ProtoMessage() {}

func (x *ECSEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECSEvent.ProtoReflect.Descriptor instead.
func (*ECSEvent) Descriptor() ([]byte, []int) {
	return file_proto_ecs_stream_proto_rawDescGZIP(), []int{0}
}

func (x *ECSEvent) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *ECSEvent) GetComponentId() string {
	if x != nil {
		return x.ComponentId
	}
	return ""
}

func (x *ECSEvent) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *ECSEvent) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ECSEvent) GetTx() string {
	if x != nil {
		return x.Tx
	}
	return ""
}

// Request to subscribe to an ECSStream. The required parameter is 'worldAddress', while others
// are opt-in based on which data the client is interested in receiving.
type ECSStreamBlockBundleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldAddress          string `protobuf:"bytes,1,opt,name=worldAddress,proto3" json:"worldAddress,omitempty"`
	BlockNumber           bool   `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockHash             bool   `protobuf:"varint,3,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockTimestamp        bool   `protobuf:"varint,4,opt,name=blockTimestamp,proto3" json:"blockTimestamp,omitempty"`
	TransactionsConfirmed bool   `protobuf:"varint,5,opt,name=transactionsConfirmed,proto3" json:"transactionsConfirmed,omitempty"`
	EcsEvents             bool   `protobuf:"varint,6,opt,name=ecsEvents,proto3" json:"ecsEvents,omitempty"`
}

func (x *ECSStreamBlockBundleRequest) Reset() {
	*x = ECSStreamBlockBundleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECSStreamBlockBundleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECSStreamBlockBundleRequest) ProtoMessage() {}

func (x *ECSStreamBlockBundleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECSStreamBlockBundleRequest.ProtoReflect.Descriptor instead.
func (*ECSStreamBlockBundleRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecs_stream_proto_rawDescGZIP(), []int{1}
}

func (x *ECSStreamBlockBundleRequest) GetWorldAddress() string {
	if x != nil {
		return x.WorldAddress
	}
	return ""
}

func (x *ECSStreamBlockBundleRequest) GetBlockNumber() bool {
	if x != nil {
		return x.BlockNumber
	}
	return false
}

func (x *ECSStreamBlockBundleRequest) GetBlockHash() bool {
	if x != nil {
		return x.BlockHash
	}
	return false
}

func (x *ECSStreamBlockBundleRequest) GetBlockTimestamp() bool {
	if x != nil {
		return x.BlockTimestamp
	}
	return false
}

func (x *ECSStreamBlockBundleRequest) GetTransactionsConfirmed() bool {
	if x != nil {
		return x.TransactionsConfirmed
	}
	return false
}

func (x *ECSStreamBlockBundleRequest) GetEcsEvents() bool {
	if x != nil {
		return x.EcsEvents
	}
	return false
}

// ECSStream response. The fields are populated based on the request which must have been sent when
// starting the subscription.
type ECSStreamBlockBundleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber           uint32      `protobuf:"varint,1,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockHash             string      `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockTimestamp        uint32      `protobuf:"varint,3,opt,name=blockTimestamp,proto3" json:"blockTimestamp,omitempty"`
	TransactionsConfirmed []string    `protobuf:"bytes,4,rep,name=transactionsConfirmed,proto3" json:"transactionsConfirmed,omitempty"`
	EcsEvents             []*ECSEvent `protobuf:"bytes,5,rep,name=ecsEvents,proto3" json:"ecsEvents,omitempty"`
}

func (x *ECSStreamBlockBundleReply) Reset() {
	*x = ECSStreamBlockBundleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecs_stream_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECSStreamBlockBundleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECSStreamBlockBundleReply) ProtoMessage() {}

func (x *ECSStreamBlockBundleReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecs_stream_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECSStreamBlockBundleReply.ProtoReflect.Descriptor instead.
func (*ECSStreamBlockBundleReply) Descriptor() ([]byte, []int) {
	return file_proto_ecs_stream_proto_rawDescGZIP(), []int{2}
}

func (x *ECSStreamBlockBundleReply) GetBlockNumber() uint32 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *ECSStreamBlockBundleReply) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *ECSStreamBlockBundleReply) GetBlockTimestamp() uint32 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

func (x *ECSStreamBlockBundleReply) GetTransactionsConfirmed() []string {
	if x != nil {
		return x.TransactionsConfirmed
	}
	return nil
}

func (x *ECSStreamBlockBundleReply) GetEcsEvents() []*ECSEvent {
	if x != nil {
		return x.EcsEvents
	}
	return nil
}

var File_proto_ecs_stream_proto protoreflect.FileDescriptor

var file_proto_ecs_stream_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x73, 0x2d, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x63, 0x73, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x45, 0x43, 0x53, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x78, 0x22, 0xfd, 0x01, 0x0a, 0x1b, 0x45, 0x43, 0x53, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x34,
	0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xec, 0x01, 0x0a, 0x19, 0x45, 0x43, 0x53, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x34, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x31,
	0x0a, 0x09, 0x65, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x65, 0x63, 0x73, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x43,
	0x53, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x63, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x32, 0x7f, 0x0a, 0x10, 0x45, 0x43, 0x53, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x2e, 0x65, 0x63, 0x73, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x43, 0x53,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x63, 0x73, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x43, 0x53, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x18, 0x5a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x6f, 0x2f, 0x65, 0x63, 0x73, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ecs_stream_proto_rawDescOnce sync.Once
	file_proto_ecs_stream_proto_rawDescData = file_proto_ecs_stream_proto_rawDesc
)

func file_proto_ecs_stream_proto_rawDescGZIP() []byte {
	file_proto_ecs_stream_proto_rawDescOnce.Do(func() {
		file_proto_ecs_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ecs_stream_proto_rawDescData)
	})
	return file_proto_ecs_stream_proto_rawDescData
}

var file_proto_ecs_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_ecs_stream_proto_goTypes = []interface{}{
	(*ECSEvent)(nil),                    // 0: ecsstream.ECSEvent
	(*ECSStreamBlockBundleRequest)(nil), // 1: ecsstream.ECSStreamBlockBundleRequest
	(*ECSStreamBlockBundleReply)(nil),   // 2: ecsstream.ECSStreamBlockBundleReply
}
var file_proto_ecs_stream_proto_depIdxs = []int32{
	0, // 0: ecsstream.ECSStreamBlockBundleReply.ecsEvents:type_name -> ecsstream.ECSEvent
	1, // 1: ecsstream.ECSStreamService.SubscribeToStreamLatest:input_type -> ecsstream.ECSStreamBlockBundleRequest
	2, // 2: ecsstream.ECSStreamService.SubscribeToStreamLatest:output_type -> ecsstream.ECSStreamBlockBundleReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_ecs_stream_proto_init() }
func file_proto_ecs_stream_proto_init() {
	if File_proto_ecs_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ecs_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECSEvent); i {
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
		file_proto_ecs_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECSStreamBlockBundleRequest); i {
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
		file_proto_ecs_stream_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECSStreamBlockBundleReply); i {
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
			RawDescriptor: file_proto_ecs_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ecs_stream_proto_goTypes,
		DependencyIndexes: file_proto_ecs_stream_proto_depIdxs,
		MessageInfos:      file_proto_ecs_stream_proto_msgTypes,
	}.Build()
	File_proto_ecs_stream_proto = out.File
	file_proto_ecs_stream_proto_rawDesc = nil
	file_proto_ecs_stream_proto_goTypes = nil
	file_proto_ecs_stream_proto_depIdxs = nil
}
