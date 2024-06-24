// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: eventpush.proto

package pb

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

type EventPushCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vin           string               `protobuf:"bytes,1,opt,name=vin,proto3" json:"vin,omitempty"`
	State         VVA_CommandState     `protobuf:"varint,2,opt,name=state,json=acpState,proto3,enum=proto.VVA_CommandState" json:"state,omitempty"`
	Condition     VVA_CommandCondition `protobuf:"varint,3,opt,name=condition,json=acpCondition,proto3,enum=proto.VVA_CommandCondition" json:"condition,omitempty"`
	Type          ACP_CommandType      `protobuf:"varint,4,opt,name=type,json=acpCommandType,proto3,enum=proto.ACP_CommandType" json:"type,omitempty"`
	ProcessId     int64                `protobuf:"varint,5,opt,name=process_id,json=pid,proto3" json:"process_id,omitempty"`
	TrackingId    string               `protobuf:"bytes,6,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	CorrelationId string               `protobuf:"bytes,7,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ErrorCodes    []int32              `protobuf:"varint,8,rep,packed,name=error_codes,json=errorCodes,proto3" json:"error_codes,omitempty"`
	Guid          string               `protobuf:"bytes,9,opt,name=guid,proto3" json:"guid,omitempty"`
	TimestampInS  int64                `protobuf:"varint,10,opt,name=timestamp_in_s,json=timestamp,proto3" json:"timestamp_in_s,omitempty"`
}

func (x *EventPushCommand) Reset() {
	*x = EventPushCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventpush_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPushCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPushCommand) ProtoMessage() {}

func (x *EventPushCommand) ProtoReflect() protoreflect.Message {
	mi := &file_eventpush_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPushCommand.ProtoReflect.Descriptor instead.
func (*EventPushCommand) Descriptor() ([]byte, []int) {
	return file_eventpush_proto_rawDescGZIP(), []int{0}
}

func (x *EventPushCommand) GetVin() string {
	if x != nil {
		return x.Vin
	}
	return ""
}

func (x *EventPushCommand) GetState() VVA_CommandState {
	if x != nil {
		return x.State
	}
	return VVA_UNKNOWN_COMMAND_STATE
}

func (x *EventPushCommand) GetCondition() VVA_CommandCondition {
	if x != nil {
		return x.Condition
	}
	return VVA_UNKNWON_COMMAND_CONDITION
}

func (x *EventPushCommand) GetType() ACP_CommandType {
	if x != nil {
		return x.Type
	}
	return ACP_UNKNOWNCOMMANDTYPE
}

func (x *EventPushCommand) GetProcessId() int64 {
	if x != nil {
		return x.ProcessId
	}
	return 0
}

func (x *EventPushCommand) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *EventPushCommand) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *EventPushCommand) GetErrorCodes() []int32 {
	if x != nil {
		return x.ErrorCodes
	}
	return nil
}

func (x *EventPushCommand) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *EventPushCommand) GetTimestampInS() int64 {
	if x != nil {
		return x.TimestampInS
	}
	return 0
}

var File_eventpush_proto protoreflect.FileDescriptor

var file_eventpush_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x63, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x83, 0x03, 0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x73,
	0x68, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x56, 0x41, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x08, 0x61, 0x63, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x56, 0x41, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x63,
	0x70, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x43, 0x50, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0e, 0x61, 0x63, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x75, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x5f, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x21, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x61, 0x69, 0x6d, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x62, 0x63, 0x61, 0x72, 0x6b, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventpush_proto_rawDescOnce sync.Once
	file_eventpush_proto_rawDescData = file_eventpush_proto_rawDesc
)

func file_eventpush_proto_rawDescGZIP() []byte {
	file_eventpush_proto_rawDescOnce.Do(func() {
		file_eventpush_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventpush_proto_rawDescData)
	})
	return file_eventpush_proto_rawDescData
}

var file_eventpush_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eventpush_proto_goTypes = []any{
	(*EventPushCommand)(nil),  // 0: proto.EventPushCommand
	(VVA_CommandState)(0),     // 1: proto.VVA.CommandState
	(VVA_CommandCondition)(0), // 2: proto.VVA.CommandCondition
	(ACP_CommandType)(0),      // 3: proto.ACP.CommandType
}
var file_eventpush_proto_depIdxs = []int32{
	1, // 0: proto.EventPushCommand.state:type_name -> proto.VVA.CommandState
	2, // 1: proto.EventPushCommand.condition:type_name -> proto.VVA.CommandCondition
	3, // 2: proto.EventPushCommand.type:type_name -> proto.ACP.CommandType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eventpush_proto_init() }
func file_eventpush_proto_init() {
	if File_eventpush_proto != nil {
		return
	}
	file_acp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eventpush_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EventPushCommand); i {
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
			RawDescriptor: file_eventpush_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eventpush_proto_goTypes,
		DependencyIndexes: file_eventpush_proto_depIdxs,
		MessageInfos:      file_eventpush_proto_msgTypes,
	}.Build()
	File_eventpush_proto = out.File
	file_eventpush_proto_rawDesc = nil
	file_eventpush_proto_goTypes = nil
	file_eventpush_proto_depIdxs = nil
}
