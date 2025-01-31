// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_user_ticket_log.proto

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

// 创建日志
type CreateUserTicketLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketId int64  `protobuf:"varint,1,opt,name=userTicketId,proto3" json:"userTicketId,omitempty"`
	Status       string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Comment      string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateUserTicketLogRequest) Reset() {
	*x = CreateUserTicketLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserTicketLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserTicketLogRequest) ProtoMessage() {}

func (x *CreateUserTicketLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserTicketLogRequest.ProtoReflect.Descriptor instead.
func (*CreateUserTicketLogRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_log_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserTicketLogRequest) GetUserTicketId() int64 {
	if x != nil {
		return x.UserTicketId
	}
	return 0
}

func (x *CreateUserTicketLogRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateUserTicketLogRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CreateUserTicketLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketLogId int64 `protobuf:"varint,1,opt,name=userTicketLogId,proto3" json:"userTicketLogId,omitempty"`
}

func (x *CreateUserTicketLogResponse) Reset() {
	*x = CreateUserTicketLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserTicketLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserTicketLogResponse) ProtoMessage() {}

func (x *CreateUserTicketLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserTicketLogResponse.ProtoReflect.Descriptor instead.
func (*CreateUserTicketLogResponse) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_log_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserTicketLogResponse) GetUserTicketLogId() int64 {
	if x != nil {
		return x.UserTicketLogId
	}
	return 0
}

// 删除日志
type DeleteUserTicketLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketLogId int64 `protobuf:"varint,1,opt,name=userTicketLogId,proto3" json:"userTicketLogId,omitempty"`
}

func (x *DeleteUserTicketLogRequest) Reset() {
	*x = DeleteUserTicketLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTicketLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTicketLogRequest) ProtoMessage() {}

func (x *DeleteUserTicketLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTicketLogRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserTicketLogRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_log_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserTicketLogRequest) GetUserTicketLogId() int64 {
	if x != nil {
		return x.UserTicketLogId
	}
	return 0
}

// 查询日志数量
type CountUserTicketLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketId int64 `protobuf:"varint,1,opt,name=userTicketId,proto3" json:"userTicketId,omitempty"`
}

func (x *CountUserTicketLogsRequest) Reset() {
	*x = CountUserTicketLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserTicketLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserTicketLogsRequest) ProtoMessage() {}

func (x *CountUserTicketLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserTicketLogsRequest.ProtoReflect.Descriptor instead.
func (*CountUserTicketLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_log_proto_rawDescGZIP(), []int{3}
}

func (x *CountUserTicketLogsRequest) GetUserTicketId() int64 {
	if x != nil {
		return x.UserTicketId
	}
	return 0
}

// 列出单页日志
type ListUserTicketLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketId int64 `protobuf:"varint,1,opt,name=userTicketId,proto3" json:"userTicketId,omitempty"`
	Offset       int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Size         int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListUserTicketLogsRequest) Reset() {
	*x = ListUserTicketLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserTicketLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTicketLogsRequest) ProtoMessage() {}

func (x *ListUserTicketLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTicketLogsRequest.ProtoReflect.Descriptor instead.
func (*ListUserTicketLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_log_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserTicketLogsRequest) GetUserTicketId() int64 {
	if x != nil {
		return x.UserTicketId
	}
	return 0
}

func (x *ListUserTicketLogsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUserTicketLogsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListUserTicketLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTicketLogs []*UserTicketLog `protobuf:"bytes,1,rep,name=userTicketLogs,proto3" json:"userTicketLogs,omitempty"`
}

func (x *ListUserTicketLogsResponse) Reset() {
	*x = ListUserTicketLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_ticket_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserTicketLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTicketLogsResponse) ProtoMessage() {}

func (x *ListUserTicketLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_ticket_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTicketLogsResponse.ProtoReflect.Descriptor instead.
func (*ListUserTicketLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_user_ticket_log_proto_rawDescGZIP(), []int{5}
}

func (x *ListUserTicketLogsResponse) GetUserTicketLogs() []*UserTicketLog {
	if x != nil {
		return x.UserTicketLogs
	}
	return nil
}

var File_service_user_ticket_log_proto protoreflect.FileDescriptor

var file_service_user_ticket_log_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x72, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x22,
	0x46, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x0f, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x1a, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x19, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x57, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52,
	0x0e, 0x75, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x32,
	0xd7, 0x02, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x13, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x12,
	0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x45, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4b, 0x0a, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1e,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_ticket_log_proto_rawDescOnce sync.Once
	file_service_user_ticket_log_proto_rawDescData = file_service_user_ticket_log_proto_rawDesc
)

func file_service_user_ticket_log_proto_rawDescGZIP() []byte {
	file_service_user_ticket_log_proto_rawDescOnce.Do(func() {
		file_service_user_ticket_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_ticket_log_proto_rawDescData)
	})
	return file_service_user_ticket_log_proto_rawDescData
}

var file_service_user_ticket_log_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_user_ticket_log_proto_goTypes = []interface{}{
	(*CreateUserTicketLogRequest)(nil),  // 0: pb.CreateUserTicketLogRequest
	(*CreateUserTicketLogResponse)(nil), // 1: pb.CreateUserTicketLogResponse
	(*DeleteUserTicketLogRequest)(nil),  // 2: pb.DeleteUserTicketLogRequest
	(*CountUserTicketLogsRequest)(nil),  // 3: pb.CountUserTicketLogsRequest
	(*ListUserTicketLogsRequest)(nil),   // 4: pb.ListUserTicketLogsRequest
	(*ListUserTicketLogsResponse)(nil),  // 5: pb.ListUserTicketLogsResponse
	(*UserTicketLog)(nil),               // 6: pb.UserTicketLog
	(*RPCSuccess)(nil),                  // 7: pb.RPCSuccess
	(*RPCCountResponse)(nil),            // 8: pb.RPCCountResponse
}
var file_service_user_ticket_log_proto_depIdxs = []int32{
	6, // 0: pb.ListUserTicketLogsResponse.userTicketLogs:type_name -> pb.UserTicketLog
	0, // 1: pb.UserTicketLogService.createUserTicketLog:input_type -> pb.CreateUserTicketLogRequest
	2, // 2: pb.UserTicketLogService.deleteUserTicketLog:input_type -> pb.DeleteUserTicketLogRequest
	3, // 3: pb.UserTicketLogService.countUserTicketLogs:input_type -> pb.CountUserTicketLogsRequest
	4, // 4: pb.UserTicketLogService.listUserTicketLogs:input_type -> pb.ListUserTicketLogsRequest
	1, // 5: pb.UserTicketLogService.createUserTicketLog:output_type -> pb.CreateUserTicketLogResponse
	7, // 6: pb.UserTicketLogService.deleteUserTicketLog:output_type -> pb.RPCSuccess
	8, // 7: pb.UserTicketLogService.countUserTicketLogs:output_type -> pb.RPCCountResponse
	5, // 8: pb.UserTicketLogService.listUserTicketLogs:output_type -> pb.ListUserTicketLogsResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_user_ticket_log_proto_init() }
func file_service_user_ticket_log_proto_init() {
	if File_service_user_ticket_log_proto != nil {
		return
	}
	file_models_model_user_ticket_log_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_ticket_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserTicketLogRequest); i {
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
		file_service_user_ticket_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserTicketLogResponse); i {
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
		file_service_user_ticket_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTicketLogRequest); i {
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
		file_service_user_ticket_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserTicketLogsRequest); i {
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
		file_service_user_ticket_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserTicketLogsRequest); i {
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
		file_service_user_ticket_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserTicketLogsResponse); i {
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
			RawDescriptor: file_service_user_ticket_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_ticket_log_proto_goTypes,
		DependencyIndexes: file_service_user_ticket_log_proto_depIdxs,
		MessageInfos:      file_service_user_ticket_log_proto_msgTypes,
	}.Build()
	File_service_user_ticket_log_proto = out.File
	file_service_user_ticket_log_proto_rawDesc = nil
	file_service_user_ticket_log_proto_goTypes = nil
	file_service_user_ticket_log_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserTicketLogServiceClient is the client API for UserTicketLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserTicketLogServiceClient interface {
	// 创建日志
	CreateUserTicketLog(ctx context.Context, in *CreateUserTicketLogRequest, opts ...grpc.CallOption) (*CreateUserTicketLogResponse, error)
	// 删除日志
	DeleteUserTicketLog(ctx context.Context, in *DeleteUserTicketLogRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查询日志数量
	CountUserTicketLogs(ctx context.Context, in *CountUserTicketLogsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页日志
	ListUserTicketLogs(ctx context.Context, in *ListUserTicketLogsRequest, opts ...grpc.CallOption) (*ListUserTicketLogsResponse, error)
}

type userTicketLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserTicketLogServiceClient(cc grpc.ClientConnInterface) UserTicketLogServiceClient {
	return &userTicketLogServiceClient{cc}
}

func (c *userTicketLogServiceClient) CreateUserTicketLog(ctx context.Context, in *CreateUserTicketLogRequest, opts ...grpc.CallOption) (*CreateUserTicketLogResponse, error) {
	out := new(CreateUserTicketLogResponse)
	err := c.cc.Invoke(ctx, "/pb.UserTicketLogService/createUserTicketLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTicketLogServiceClient) DeleteUserTicketLog(ctx context.Context, in *DeleteUserTicketLogRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.UserTicketLogService/deleteUserTicketLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTicketLogServiceClient) CountUserTicketLogs(ctx context.Context, in *CountUserTicketLogsRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, "/pb.UserTicketLogService/countUserTicketLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userTicketLogServiceClient) ListUserTicketLogs(ctx context.Context, in *ListUserTicketLogsRequest, opts ...grpc.CallOption) (*ListUserTicketLogsResponse, error) {
	out := new(ListUserTicketLogsResponse)
	err := c.cc.Invoke(ctx, "/pb.UserTicketLogService/listUserTicketLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserTicketLogServiceServer is the server API for UserTicketLogService service.
type UserTicketLogServiceServer interface {
	// 创建日志
	CreateUserTicketLog(context.Context, *CreateUserTicketLogRequest) (*CreateUserTicketLogResponse, error)
	// 删除日志
	DeleteUserTicketLog(context.Context, *DeleteUserTicketLogRequest) (*RPCSuccess, error)
	// 查询日志数量
	CountUserTicketLogs(context.Context, *CountUserTicketLogsRequest) (*RPCCountResponse, error)
	// 列出单页日志
	ListUserTicketLogs(context.Context, *ListUserTicketLogsRequest) (*ListUserTicketLogsResponse, error)
}

// UnimplementedUserTicketLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserTicketLogServiceServer struct {
}

func (*UnimplementedUserTicketLogServiceServer) CreateUserTicketLog(context.Context, *CreateUserTicketLogRequest) (*CreateUserTicketLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserTicketLog not implemented")
}
func (*UnimplementedUserTicketLogServiceServer) DeleteUserTicketLog(context.Context, *DeleteUserTicketLogRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserTicketLog not implemented")
}
func (*UnimplementedUserTicketLogServiceServer) CountUserTicketLogs(context.Context, *CountUserTicketLogsRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUserTicketLogs not implemented")
}
func (*UnimplementedUserTicketLogServiceServer) ListUserTicketLogs(context.Context, *ListUserTicketLogsRequest) (*ListUserTicketLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserTicketLogs not implemented")
}

func RegisterUserTicketLogServiceServer(s *grpc.Server, srv UserTicketLogServiceServer) {
	s.RegisterService(&_UserTicketLogService_serviceDesc, srv)
}

func _UserTicketLogService_CreateUserTicketLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserTicketLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTicketLogServiceServer).CreateUserTicketLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserTicketLogService/CreateUserTicketLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTicketLogServiceServer).CreateUserTicketLog(ctx, req.(*CreateUserTicketLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTicketLogService_DeleteUserTicketLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserTicketLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTicketLogServiceServer).DeleteUserTicketLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserTicketLogService/DeleteUserTicketLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTicketLogServiceServer).DeleteUserTicketLog(ctx, req.(*DeleteUserTicketLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTicketLogService_CountUserTicketLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountUserTicketLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTicketLogServiceServer).CountUserTicketLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserTicketLogService/CountUserTicketLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTicketLogServiceServer).CountUserTicketLogs(ctx, req.(*CountUserTicketLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserTicketLogService_ListUserTicketLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserTicketLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserTicketLogServiceServer).ListUserTicketLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserTicketLogService/ListUserTicketLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserTicketLogServiceServer).ListUserTicketLogs(ctx, req.(*ListUserTicketLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserTicketLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserTicketLogService",
	HandlerType: (*UserTicketLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUserTicketLog",
			Handler:    _UserTicketLogService_CreateUserTicketLog_Handler,
		},
		{
			MethodName: "deleteUserTicketLog",
			Handler:    _UserTicketLogService_DeleteUserTicketLog_Handler,
		},
		{
			MethodName: "countUserTicketLogs",
			Handler:    _UserTicketLogService_CountUserTicketLogs_Handler,
		},
		{
			MethodName: "listUserTicketLogs",
			Handler:    _UserTicketLogService_ListUserTicketLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_user_ticket_log.proto",
}
