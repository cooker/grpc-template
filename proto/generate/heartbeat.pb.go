// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: heartbeat.proto

package proto

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

type HeartBeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//创建时间
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MsgId     string                 `protobuf:"bytes,2,opt,name=msgId,proto3" json:"msgId,omitempty"`
	FromBy    string                 `protobuf:"bytes,3,opt,name=fromBy,proto3" json:"fromBy,omitempty"`
}

func (x *HeartBeatRequest) Reset() {
	*x = HeartBeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heartbeat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBeatRequest) ProtoMessage() {}

func (x *HeartBeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heartbeat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBeatRequest.ProtoReflect.Descriptor instead.
func (*HeartBeatRequest) Descriptor() ([]byte, []int) {
	return file_heartbeat_proto_rawDescGZIP(), []int{0}
}

func (x *HeartBeatRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *HeartBeatRequest) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *HeartBeatRequest) GetFromBy() string {
	if x != nil {
		return x.FromBy
	}
	return ""
}

var File_heartbeat_proto protoreflect.FileDescriptor

var file_heartbeat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x72, 0x6f, 0x6d, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72,
	0x6f, 0x6d, 0x42, 0x79, 0x32, 0x4b, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heartbeat_proto_rawDescOnce sync.Once
	file_heartbeat_proto_rawDescData = file_heartbeat_proto_rawDesc
)

func file_heartbeat_proto_rawDescGZIP() []byte {
	file_heartbeat_proto_rawDescOnce.Do(func() {
		file_heartbeat_proto_rawDescData = protoimpl.X.CompressGZIP(file_heartbeat_proto_rawDescData)
	})
	return file_heartbeat_proto_rawDescData
}

var file_heartbeat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_heartbeat_proto_goTypes = []interface{}{
	(*HeartBeatRequest)(nil),      // 0: proto.HeartBeatRequest
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*MessageResponse)(nil),       // 2: proto.MessageResponse
}
var file_heartbeat_proto_depIdxs = []int32{
	1, // 0: proto.HeartBeatRequest.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: proto.HeartBeatService.Send:input_type -> proto.HeartBeatRequest
	2, // 2: proto.HeartBeatService.Send:output_type -> proto.MessageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_heartbeat_proto_init() }
func file_heartbeat_proto_init() {
	if File_heartbeat_proto != nil {
		return
	}
	file_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_heartbeat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBeatRequest); i {
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
			RawDescriptor: file_heartbeat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heartbeat_proto_goTypes,
		DependencyIndexes: file_heartbeat_proto_depIdxs,
		MessageInfos:      file_heartbeat_proto_msgTypes,
	}.Build()
	File_heartbeat_proto = out.File
	file_heartbeat_proto_rawDesc = nil
	file_heartbeat_proto_goTypes = nil
	file_heartbeat_proto_depIdxs = nil
}
