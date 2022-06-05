// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_http_cache_task_key.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 校验缓存Key
type ValidateHTTPCacheTaskKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *ValidateHTTPCacheTaskKeysRequest) Reset() {
	*x = ValidateHTTPCacheTaskKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_cache_task_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateHTTPCacheTaskKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateHTTPCacheTaskKeysRequest) ProtoMessage() {}

func (x *ValidateHTTPCacheTaskKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_cache_task_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateHTTPCacheTaskKeysRequest.ProtoReflect.Descriptor instead.
func (*ValidateHTTPCacheTaskKeysRequest) Descriptor() ([]byte, []int) {
	return file_service_http_cache_task_key_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateHTTPCacheTaskKeysRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type ValidateHTTPCacheTaskKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailKeys []*ValidateHTTPCacheTaskKeysResponse_FailKey `protobuf:"bytes,1,rep,name=failKeys,proto3" json:"failKeys,omitempty"`
}

func (x *ValidateHTTPCacheTaskKeysResponse) Reset() {
	*x = ValidateHTTPCacheTaskKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_cache_task_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateHTTPCacheTaskKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateHTTPCacheTaskKeysResponse) ProtoMessage() {}

func (x *ValidateHTTPCacheTaskKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_cache_task_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateHTTPCacheTaskKeysResponse.ProtoReflect.Descriptor instead.
func (*ValidateHTTPCacheTaskKeysResponse) Descriptor() ([]byte, []int) {
	return file_service_http_cache_task_key_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateHTTPCacheTaskKeysResponse) GetFailKeys() []*ValidateHTTPCacheTaskKeysResponse_FailKey {
	if x != nil {
		return x.FailKeys
	}
	return nil
}

// 查找需要执行的Key
type FindDoingHTTPCacheTaskKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindDoingHTTPCacheTaskKeysRequest) Reset() {
	*x = FindDoingHTTPCacheTaskKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_cache_task_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDoingHTTPCacheTaskKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDoingHTTPCacheTaskKeysRequest) ProtoMessage() {}

func (x *FindDoingHTTPCacheTaskKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_cache_task_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDoingHTTPCacheTaskKeysRequest.ProtoReflect.Descriptor instead.
func (*FindDoingHTTPCacheTaskKeysRequest) Descriptor() ([]byte, []int) {
	return file_service_http_cache_task_key_proto_rawDescGZIP(), []int{2}
}

func (x *FindDoingHTTPCacheTaskKeysRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindDoingHTTPCacheTaskKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpCacheTaskKeys []*HTTPCacheTaskKey `protobuf:"bytes,1,rep,name=httpCacheTaskKeys,proto3" json:"httpCacheTaskKeys,omitempty"`
}

func (x *FindDoingHTTPCacheTaskKeysResponse) Reset() {
	*x = FindDoingHTTPCacheTaskKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_cache_task_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindDoingHTTPCacheTaskKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindDoingHTTPCacheTaskKeysResponse) ProtoMessage() {}

func (x *FindDoingHTTPCacheTaskKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_cache_task_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindDoingHTTPCacheTaskKeysResponse.ProtoReflect.Descriptor instead.
func (*FindDoingHTTPCacheTaskKeysResponse) Descriptor() ([]byte, []int) {
	return file_service_http_cache_task_key_proto_rawDescGZIP(), []int{3}
}

func (x *FindDoingHTTPCacheTaskKeysResponse) GetHttpCacheTaskKeys() []*HTTPCacheTaskKey {
	if x != nil {
		return x.HttpCacheTaskKeys
	}
	return nil
}

// 更新一组Key状态
type UpdateHTTPCacheTaskKeysStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyResults []*UpdateHTTPCacheTaskKeysStatusRequest_KeyResult `protobuf:"bytes,1,rep,name=keyResults,proto3" json:"keyResults,omitempty"`
}

func (x *UpdateHTTPCacheTaskKeysStatusRequest) Reset() {
	*x = UpdateHTTPCacheTaskKeysStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_cache_task_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPCacheTaskKeysStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPCacheTaskKeysStatusRequest) ProtoMessage() {}

func (x *UpdateHTTPCacheTaskKeysStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_cache_task_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPCacheTaskKeysStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPCacheTaskKeysStatusRequest) Descriptor() ([]byte, []int) {
	return file_service_http_cache_task_key_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateHTTPCacheTaskKeysStatusRequest) GetKeyResults() []*UpdateHTTPCacheTaskKeysStatusRequest_KeyResult {
	if x != nil {
		return x.KeyResults
	}
	return nil
}

type ValidateHTTPCacheTaskKeysResponse_FailKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ReasonCode string `protobuf:"bytes,2,opt,name=reasonCode,proto3" json:"reasonCode,omitempty"`
}

func (x *ValidateHTTPCacheTaskKeysResponse_FailKey) Reset() {
	*x = ValidateHTTPCacheTaskKeysResponse_FailKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_cache_task_key_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateHTTPCacheTaskKeysResponse_FailKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateHTTPCacheTaskKeysResponse_FailKey) ProtoMessage() {}

func (x *ValidateHTTPCacheTaskKeysResponse_FailKey) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_cache_task_key_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateHTTPCacheTaskKeysResponse_FailKey.ProtoReflect.Descriptor instead.
func (*ValidateHTTPCacheTaskKeysResponse_FailKey) Descriptor() ([]byte, []int) {
	return file_service_http_cache_task_key_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ValidateHTTPCacheTaskKeysResponse_FailKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ValidateHTTPCacheTaskKeysResponse_FailKey) GetReasonCode() string {
	if x != nil {
		return x.ReasonCode
	}
	return ""
}

type UpdateHTTPCacheTaskKeysStatusRequest_KeyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeClusterId int64  `protobuf:"varint,2,opt,name=nodeClusterId,proto3" json:"nodeClusterId,omitempty"` // 特意设置的冗余数据
	Error         string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateHTTPCacheTaskKeysStatusRequest_KeyResult) Reset() {
	*x = UpdateHTTPCacheTaskKeysStatusRequest_KeyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_cache_task_key_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPCacheTaskKeysStatusRequest_KeyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPCacheTaskKeysStatusRequest_KeyResult) ProtoMessage() {}

func (x *UpdateHTTPCacheTaskKeysStatusRequest_KeyResult) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_cache_task_key_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPCacheTaskKeysStatusRequest_KeyResult.ProtoReflect.Descriptor instead.
func (*UpdateHTTPCacheTaskKeysStatusRequest_KeyResult) Descriptor() ([]byte, []int) {
	return file_service_http_cache_task_key_proto_rawDescGZIP(), []int{4, 0}
}

func (x *UpdateHTTPCacheTaskKeysStatusRequest_KeyResult) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateHTTPCacheTaskKeysStatusRequest_KeyResult) GetNodeClusterId() int64 {
	if x != nil {
		return x.NodeClusterId
	}
	return 0
}

func (x *UpdateHTTPCacheTaskKeysStatusRequest_KeyResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_service_http_cache_task_key_proto protoreflect.FileDescriptor

var file_service_http_cache_task_key_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x20, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x21, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c,
	0x4b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x62, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x4b,
	0x65, 0x79, 0x73, 0x1a, 0x3b, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x4b, 0x65, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x37, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x48, 0x54, 0x54,
	0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x68, 0x0a, 0x22, 0x46, 0x69, 0x6e,
	0x64, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x4b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x2e,
	0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79,
	0x52, 0x11, 0x68, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b,
	0x65, 0x79, 0x73, 0x22, 0xd3, 0x01, 0x0a, 0x24, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54,
	0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x52, 0x0a, 0x0a,
	0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x1a, 0x57, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xcb, 0x02, 0x0a, 0x17, 0x48, 0x54,
	0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x19, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65,
	0x79, 0x73, 0x12, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x6b, 0x0a, 0x1a, 0x66, 0x69, 0x6e, 0x64, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x48, 0x54, 0x54, 0x50,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x25, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x48, 0x54, 0x54, 0x50,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x6f,
	0x69, 0x6e, 0x67, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x1d,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_cache_task_key_proto_rawDescOnce sync.Once
	file_service_http_cache_task_key_proto_rawDescData = file_service_http_cache_task_key_proto_rawDesc
)

func file_service_http_cache_task_key_proto_rawDescGZIP() []byte {
	file_service_http_cache_task_key_proto_rawDescOnce.Do(func() {
		file_service_http_cache_task_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_cache_task_key_proto_rawDescData)
	})
	return file_service_http_cache_task_key_proto_rawDescData
}

var file_service_http_cache_task_key_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_http_cache_task_key_proto_goTypes = []interface{}{
	(*ValidateHTTPCacheTaskKeysRequest)(nil),               // 0: pb.ValidateHTTPCacheTaskKeysRequest
	(*ValidateHTTPCacheTaskKeysResponse)(nil),              // 1: pb.ValidateHTTPCacheTaskKeysResponse
	(*FindDoingHTTPCacheTaskKeysRequest)(nil),              // 2: pb.FindDoingHTTPCacheTaskKeysRequest
	(*FindDoingHTTPCacheTaskKeysResponse)(nil),             // 3: pb.FindDoingHTTPCacheTaskKeysResponse
	(*UpdateHTTPCacheTaskKeysStatusRequest)(nil),           // 4: pb.UpdateHTTPCacheTaskKeysStatusRequest
	(*ValidateHTTPCacheTaskKeysResponse_FailKey)(nil),      // 5: pb.ValidateHTTPCacheTaskKeysResponse.FailKey
	(*UpdateHTTPCacheTaskKeysStatusRequest_KeyResult)(nil), // 6: pb.UpdateHTTPCacheTaskKeysStatusRequest.KeyResult
	(*HTTPCacheTaskKey)(nil),                               // 7: pb.HTTPCacheTaskKey
	(*RPCSuccess)(nil),                                     // 8: pb.RPCSuccess
}
var file_service_http_cache_task_key_proto_depIdxs = []int32{
	5, // 0: pb.ValidateHTTPCacheTaskKeysResponse.failKeys:type_name -> pb.ValidateHTTPCacheTaskKeysResponse.FailKey
	7, // 1: pb.FindDoingHTTPCacheTaskKeysResponse.httpCacheTaskKeys:type_name -> pb.HTTPCacheTaskKey
	6, // 2: pb.UpdateHTTPCacheTaskKeysStatusRequest.keyResults:type_name -> pb.UpdateHTTPCacheTaskKeysStatusRequest.KeyResult
	0, // 3: pb.HTTPCacheTaskKeyService.validateHTTPCacheTaskKeys:input_type -> pb.ValidateHTTPCacheTaskKeysRequest
	2, // 4: pb.HTTPCacheTaskKeyService.findDoingHTTPCacheTaskKeys:input_type -> pb.FindDoingHTTPCacheTaskKeysRequest
	4, // 5: pb.HTTPCacheTaskKeyService.updateHTTPCacheTaskKeysStatus:input_type -> pb.UpdateHTTPCacheTaskKeysStatusRequest
	1, // 6: pb.HTTPCacheTaskKeyService.validateHTTPCacheTaskKeys:output_type -> pb.ValidateHTTPCacheTaskKeysResponse
	3, // 7: pb.HTTPCacheTaskKeyService.findDoingHTTPCacheTaskKeys:output_type -> pb.FindDoingHTTPCacheTaskKeysResponse
	8, // 8: pb.HTTPCacheTaskKeyService.updateHTTPCacheTaskKeysStatus:output_type -> pb.RPCSuccess
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_http_cache_task_key_proto_init() }
func file_service_http_cache_task_key_proto_init() {
	if File_service_http_cache_task_key_proto != nil {
		return
	}
	file_models_model_http_cache_task_key_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_cache_task_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateHTTPCacheTaskKeysRequest); i {
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
		file_service_http_cache_task_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateHTTPCacheTaskKeysResponse); i {
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
		file_service_http_cache_task_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDoingHTTPCacheTaskKeysRequest); i {
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
		file_service_http_cache_task_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindDoingHTTPCacheTaskKeysResponse); i {
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
		file_service_http_cache_task_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPCacheTaskKeysStatusRequest); i {
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
		file_service_http_cache_task_key_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateHTTPCacheTaskKeysResponse_FailKey); i {
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
		file_service_http_cache_task_key_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPCacheTaskKeysStatusRequest_KeyResult); i {
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
			RawDescriptor: file_service_http_cache_task_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_cache_task_key_proto_goTypes,
		DependencyIndexes: file_service_http_cache_task_key_proto_depIdxs,
		MessageInfos:      file_service_http_cache_task_key_proto_msgTypes,
	}.Build()
	File_service_http_cache_task_key_proto = out.File
	file_service_http_cache_task_key_proto_rawDesc = nil
	file_service_http_cache_task_key_proto_goTypes = nil
	file_service_http_cache_task_key_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPCacheTaskKeyServiceClient is the client API for HTTPCacheTaskKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPCacheTaskKeyServiceClient interface {
	// 校验缓存Key
	ValidateHTTPCacheTaskKeys(ctx context.Context, in *ValidateHTTPCacheTaskKeysRequest, opts ...grpc.CallOption) (*ValidateHTTPCacheTaskKeysResponse, error)
	// 查找需要执行的Key
	FindDoingHTTPCacheTaskKeys(ctx context.Context, in *FindDoingHTTPCacheTaskKeysRequest, opts ...grpc.CallOption) (*FindDoingHTTPCacheTaskKeysResponse, error)
	// 更新一组Key状态
	UpdateHTTPCacheTaskKeysStatus(ctx context.Context, in *UpdateHTTPCacheTaskKeysStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
}

type hTTPCacheTaskKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPCacheTaskKeyServiceClient(cc grpc.ClientConnInterface) HTTPCacheTaskKeyServiceClient {
	return &hTTPCacheTaskKeyServiceClient{cc}
}

func (c *hTTPCacheTaskKeyServiceClient) ValidateHTTPCacheTaskKeys(ctx context.Context, in *ValidateHTTPCacheTaskKeysRequest, opts ...grpc.CallOption) (*ValidateHTTPCacheTaskKeysResponse, error) {
	out := new(ValidateHTTPCacheTaskKeysResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPCacheTaskKeyService/validateHTTPCacheTaskKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCacheTaskKeyServiceClient) FindDoingHTTPCacheTaskKeys(ctx context.Context, in *FindDoingHTTPCacheTaskKeysRequest, opts ...grpc.CallOption) (*FindDoingHTTPCacheTaskKeysResponse, error) {
	out := new(FindDoingHTTPCacheTaskKeysResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPCacheTaskKeyService/findDoingHTTPCacheTaskKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPCacheTaskKeyServiceClient) UpdateHTTPCacheTaskKeysStatus(ctx context.Context, in *UpdateHTTPCacheTaskKeysStatusRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPCacheTaskKeyService/updateHTTPCacheTaskKeysStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPCacheTaskKeyServiceServer is the server API for HTTPCacheTaskKeyService service.
type HTTPCacheTaskKeyServiceServer interface {
	// 校验缓存Key
	ValidateHTTPCacheTaskKeys(context.Context, *ValidateHTTPCacheTaskKeysRequest) (*ValidateHTTPCacheTaskKeysResponse, error)
	// 查找需要执行的Key
	FindDoingHTTPCacheTaskKeys(context.Context, *FindDoingHTTPCacheTaskKeysRequest) (*FindDoingHTTPCacheTaskKeysResponse, error)
	// 更新一组Key状态
	UpdateHTTPCacheTaskKeysStatus(context.Context, *UpdateHTTPCacheTaskKeysStatusRequest) (*RPCSuccess, error)
}

// UnimplementedHTTPCacheTaskKeyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPCacheTaskKeyServiceServer struct {
}

func (*UnimplementedHTTPCacheTaskKeyServiceServer) ValidateHTTPCacheTaskKeys(context.Context, *ValidateHTTPCacheTaskKeysRequest) (*ValidateHTTPCacheTaskKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateHTTPCacheTaskKeys not implemented")
}
func (*UnimplementedHTTPCacheTaskKeyServiceServer) FindDoingHTTPCacheTaskKeys(context.Context, *FindDoingHTTPCacheTaskKeysRequest) (*FindDoingHTTPCacheTaskKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDoingHTTPCacheTaskKeys not implemented")
}
func (*UnimplementedHTTPCacheTaskKeyServiceServer) UpdateHTTPCacheTaskKeysStatus(context.Context, *UpdateHTTPCacheTaskKeysStatusRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPCacheTaskKeysStatus not implemented")
}

func RegisterHTTPCacheTaskKeyServiceServer(s *grpc.Server, srv HTTPCacheTaskKeyServiceServer) {
	s.RegisterService(&_HTTPCacheTaskKeyService_serviceDesc, srv)
}

func _HTTPCacheTaskKeyService_ValidateHTTPCacheTaskKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateHTTPCacheTaskKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCacheTaskKeyServiceServer).ValidateHTTPCacheTaskKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCacheTaskKeyService/ValidateHTTPCacheTaskKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCacheTaskKeyServiceServer).ValidateHTTPCacheTaskKeys(ctx, req.(*ValidateHTTPCacheTaskKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCacheTaskKeyService_FindDoingHTTPCacheTaskKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDoingHTTPCacheTaskKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCacheTaskKeyServiceServer).FindDoingHTTPCacheTaskKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCacheTaskKeyService/FindDoingHTTPCacheTaskKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCacheTaskKeyServiceServer).FindDoingHTTPCacheTaskKeys(ctx, req.(*FindDoingHTTPCacheTaskKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPCacheTaskKeyService_UpdateHTTPCacheTaskKeysStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPCacheTaskKeysStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPCacheTaskKeyServiceServer).UpdateHTTPCacheTaskKeysStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPCacheTaskKeyService/UpdateHTTPCacheTaskKeysStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPCacheTaskKeyServiceServer).UpdateHTTPCacheTaskKeysStatus(ctx, req.(*UpdateHTTPCacheTaskKeysStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTPCacheTaskKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPCacheTaskKeyService",
	HandlerType: (*HTTPCacheTaskKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "validateHTTPCacheTaskKeys",
			Handler:    _HTTPCacheTaskKeyService_ValidateHTTPCacheTaskKeys_Handler,
		},
		{
			MethodName: "findDoingHTTPCacheTaskKeys",
			Handler:    _HTTPCacheTaskKeyService_FindDoingHTTPCacheTaskKeys_Handler,
		},
		{
			MethodName: "updateHTTPCacheTaskKeysStatus",
			Handler:    _HTTPCacheTaskKeyService_UpdateHTTPCacheTaskKeysStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_cache_task_key.proto",
}
