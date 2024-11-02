// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: worker/worker.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MethodName string       `protobuf:"bytes,1,opt,name=MethodName,proto3" json:"MethodName,omitempty"`
	MethodArgs []*anypb.Any `protobuf:"bytes,2,rep,name=MethodArgs,proto3" json:"MethodArgs,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	mi := &file_worker_worker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_worker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_worker_worker_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteRequest) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *ExecuteRequest) GetMethodArgs() []*anypb.Any {
	if x != nil {
		return x.MethodArgs
	}
	return nil
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32      `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string     `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Result  *anypb.Any `protobuf:"bytes,3,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	mi := &file_worker_worker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worker_worker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_worker_worker_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ExecuteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ExecuteResponse) GetResult() *anypb.Any {
	if x != nil {
		return x.Result
	}
	return nil
}

type SimpleString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SimpleString) Reset() {
	*x = SimpleString{}
	mi := &file_worker_worker_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleString) ProtoMessage() {}

func (x *SimpleString) ProtoReflect() protoreflect.Message {
	mi := &file_worker_worker_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleString.ProtoReflect.Descriptor instead.
func (*SimpleString) Descriptor() ([]byte, []int) {
	return file_worker_worker_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SimpleInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SimpleInt) Reset() {
	*x = SimpleInt{}
	mi := &file_worker_worker_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleInt) ProtoMessage() {}

func (x *SimpleInt) ProtoReflect() protoreflect.Message {
	mi := &file_worker_worker_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleInt.ProtoReflect.Descriptor instead.
func (*SimpleInt) Descriptor() ([]byte, []int) {
	return file_worker_worker_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleInt) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SimpleFloat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SimpleFloat) Reset() {
	*x = SimpleFloat{}
	mi := &file_worker_worker_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleFloat) ProtoMessage() {}

func (x *SimpleFloat) ProtoReflect() protoreflect.Message {
	mi := &file_worker_worker_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleFloat.ProtoReflect.Descriptor instead.
func (*SimpleFloat) Descriptor() ([]byte, []int) {
	return file_worker_worker_proto_rawDescGZIP(), []int{4}
}

func (x *SimpleFloat) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_worker_worker_proto protoreflect.FileDescriptor

var file_worker_worker_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x62,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x0e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41,
	0x72, 0x67, 0x73, 0x22, 0x71, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x24, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x23, 0x0a, 0x0b, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0x4c, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x42,
	0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x62,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_worker_proto_rawDescOnce sync.Once
	file_worker_worker_proto_rawDescData = file_worker_worker_proto_rawDesc
)

func file_worker_worker_proto_rawDescGZIP() []byte {
	file_worker_worker_proto_rawDescOnce.Do(func() {
		file_worker_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_worker_proto_rawDescData)
	})
	return file_worker_worker_proto_rawDescData
}

var file_worker_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_worker_worker_proto_goTypes = []any{
	(*ExecuteRequest)(nil),  // 0: worker.pb.ExecuteRequest
	(*ExecuteResponse)(nil), // 1: worker.pb.ExecuteResponse
	(*SimpleString)(nil),    // 2: worker.pb.SimpleString
	(*SimpleInt)(nil),       // 3: worker.pb.SimpleInt
	(*SimpleFloat)(nil),     // 4: worker.pb.SimpleFloat
	(*anypb.Any)(nil),       // 5: google.protobuf.Any
}
var file_worker_worker_proto_depIdxs = []int32{
	5, // 0: worker.pb.ExecuteRequest.MethodArgs:type_name -> google.protobuf.Any
	5, // 1: worker.pb.ExecuteResponse.Result:type_name -> google.protobuf.Any
	0, // 2: worker.pb.Worker.Execute:input_type -> worker.pb.ExecuteRequest
	1, // 3: worker.pb.Worker.Execute:output_type -> worker.pb.ExecuteResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_worker_worker_proto_init() }
func file_worker_worker_proto_init() {
	if File_worker_worker_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_worker_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_worker_proto_goTypes,
		DependencyIndexes: file_worker_worker_proto_depIdxs,
		MessageInfos:      file_worker_worker_proto_msgTypes,
	}.Build()
	File_worker_worker_proto = out.File
	file_worker_worker_proto_rawDesc = nil
	file_worker_worker_proto_goTypes = nil
	file_worker_worker_proto_depIdxs = nil
}