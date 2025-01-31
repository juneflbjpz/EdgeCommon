// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: service_http_firewall_rule_set.proto

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

// 根据配置创建或修改规则集
type CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetConfigJSON []byte `protobuf:"bytes,1,opt,name=firewallRuleSetConfigJSON,proto3" json:"firewallRuleSetConfigJSON,omitempty"`
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) Reset() {
	*x = CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) ProtoMessage() {}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) GetFirewallRuleSetConfigJSON() []byte {
	if x != nil {
		return x.FirewallRuleSetConfigJSON
	}
	return nil
}

type CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetId int64 `protobuf:"varint,1,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) Reset() {
	*x = CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) ProtoMessage() {}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse.ProtoReflect.Descriptor instead.
func (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

// 设置开启状态
type UpdateHTTPFirewallRuleSetIsOnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetId int64 `protobuf:"varint,1,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`
	IsOn              bool  `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) Reset() {
	*x = UpdateHTTPFirewallRuleSetIsOnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHTTPFirewallRuleSetIsOnRequest) ProtoMessage() {}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHTTPFirewallRuleSetIsOnRequest.ProtoReflect.Descriptor instead.
func (*UpdateHTTPFirewallRuleSetIsOnRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

func (x *UpdateHTTPFirewallRuleSetIsOnRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 查找规则集配置
type FindEnabledHTTPFirewallRuleSetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetId int64 `protobuf:"varint,1,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`
}

func (x *FindEnabledHTTPFirewallRuleSetConfigRequest) Reset() {
	*x = FindEnabledHTTPFirewallRuleSetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFirewallRuleSetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFirewallRuleSetConfigRequest) ProtoMessage() {}

func (x *FindEnabledHTTPFirewallRuleSetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFirewallRuleSetConfigRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFirewallRuleSetConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledHTTPFirewallRuleSetConfigRequest) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

type FindEnabledHTTPFirewallRuleSetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetJSON []byte `protobuf:"bytes,1,opt,name=firewallRuleSetJSON,proto3" json:"firewallRuleSetJSON,omitempty"`
}

func (x *FindEnabledHTTPFirewallRuleSetConfigResponse) Reset() {
	*x = FindEnabledHTTPFirewallRuleSetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFirewallRuleSetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFirewallRuleSetConfigResponse) ProtoMessage() {}

func (x *FindEnabledHTTPFirewallRuleSetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFirewallRuleSetConfigResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFirewallRuleSetConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledHTTPFirewallRuleSetConfigResponse) GetFirewallRuleSetJSON() []byte {
	if x != nil {
		return x.FirewallRuleSetJSON
	}
	return nil
}

// 查找规则集
type FindEnabledHTTPFirewallRuleSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSetId int64 `protobuf:"varint,1,opt,name=firewallRuleSetId,proto3" json:"firewallRuleSetId,omitempty"`
}

func (x *FindEnabledHTTPFirewallRuleSetRequest) Reset() {
	*x = FindEnabledHTTPFirewallRuleSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFirewallRuleSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFirewallRuleSetRequest) ProtoMessage() {}

func (x *FindEnabledHTTPFirewallRuleSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFirewallRuleSetRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFirewallRuleSetRequest) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{5}
}

func (x *FindEnabledHTTPFirewallRuleSetRequest) GetFirewallRuleSetId() int64 {
	if x != nil {
		return x.FirewallRuleSetId
	}
	return 0
}

type FindEnabledHTTPFirewallRuleSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirewallRuleSet *HTTPFirewallRuleSet `protobuf:"bytes,1,opt,name=firewallRuleSet,proto3" json:"firewallRuleSet,omitempty"`
}

func (x *FindEnabledHTTPFirewallRuleSetResponse) Reset() {
	*x = FindEnabledHTTPFirewallRuleSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_http_firewall_rule_set_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledHTTPFirewallRuleSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledHTTPFirewallRuleSetResponse) ProtoMessage() {}

func (x *FindEnabledHTTPFirewallRuleSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_http_firewall_rule_set_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledHTTPFirewallRuleSetResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledHTTPFirewallRuleSetResponse) Descriptor() ([]byte, []int) {
	return file_service_http_firewall_rule_set_proto_rawDescGZIP(), []int{6}
}

func (x *FindEnabledHTTPFirewallRuleSetResponse) GetFirewallRuleSet() *HTTPFirewallRuleSet {
	if x != nil {
		return x.FirewallRuleSet
	}
	return nil
}

var File_service_http_firewall_rule_set_proto protoreflect.FileDescriptor

var file_service_http_firewall_rule_set_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c,
	0x6c, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x72, 0x0a, 0x32, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x19, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a,
	0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x19, 0x66, 0x69, 0x72, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x63, 0x0a, 0x33, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x24, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x73, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69,
	0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x22, 0x5b, 0x0a, 0x2b, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64,
	0x22, 0x60, 0x0a, 0x2c, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48,
	0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x4a, 0x53,
	0x4f, 0x4e, 0x22, 0x55, 0x0a, 0x25, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x66,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x26, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x32, 0x9d, 0x04, 0x0a, 0x1a, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x2b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72,
	0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x1d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x49, 0x73, 0x4f, 0x6e, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52,
	0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x49, 0x73, 0x4f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x89, 0x01, 0x0a, 0x24, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x77, 0x0a,
	0x1e, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50,
	0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x29, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x62, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x48, 0x54, 0x54, 0x50, 0x46,
	0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_http_firewall_rule_set_proto_rawDescOnce sync.Once
	file_service_http_firewall_rule_set_proto_rawDescData = file_service_http_firewall_rule_set_proto_rawDesc
)

func file_service_http_firewall_rule_set_proto_rawDescGZIP() []byte {
	file_service_http_firewall_rule_set_proto_rawDescOnce.Do(func() {
		file_service_http_firewall_rule_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_http_firewall_rule_set_proto_rawDescData)
	})
	return file_service_http_firewall_rule_set_proto_rawDescData
}

var file_service_http_firewall_rule_set_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_http_firewall_rule_set_proto_goTypes = []interface{}{
	(*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest)(nil),  // 0: pb.CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest
	(*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse)(nil), // 1: pb.CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse
	(*UpdateHTTPFirewallRuleSetIsOnRequest)(nil),                // 2: pb.UpdateHTTPFirewallRuleSetIsOnRequest
	(*FindEnabledHTTPFirewallRuleSetConfigRequest)(nil),         // 3: pb.FindEnabledHTTPFirewallRuleSetConfigRequest
	(*FindEnabledHTTPFirewallRuleSetConfigResponse)(nil),        // 4: pb.FindEnabledHTTPFirewallRuleSetConfigResponse
	(*FindEnabledHTTPFirewallRuleSetRequest)(nil),               // 5: pb.FindEnabledHTTPFirewallRuleSetRequest
	(*FindEnabledHTTPFirewallRuleSetResponse)(nil),              // 6: pb.FindEnabledHTTPFirewallRuleSetResponse
	(*HTTPFirewallRuleSet)(nil),                                 // 7: pb.HTTPFirewallRuleSet
	(*RPCSuccess)(nil),                                          // 8: pb.RPCSuccess
}
var file_service_http_firewall_rule_set_proto_depIdxs = []int32{
	7, // 0: pb.FindEnabledHTTPFirewallRuleSetResponse.firewallRuleSet:type_name -> pb.HTTPFirewallRuleSet
	0, // 1: pb.HTTPFirewallRuleSetService.createOrUpdateHTTPFirewallRuleSetFromConfig:input_type -> pb.CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest
	2, // 2: pb.HTTPFirewallRuleSetService.updateHTTPFirewallRuleSetIsOn:input_type -> pb.UpdateHTTPFirewallRuleSetIsOnRequest
	3, // 3: pb.HTTPFirewallRuleSetService.findEnabledHTTPFirewallRuleSetConfig:input_type -> pb.FindEnabledHTTPFirewallRuleSetConfigRequest
	5, // 4: pb.HTTPFirewallRuleSetService.findEnabledHTTPFirewallRuleSet:input_type -> pb.FindEnabledHTTPFirewallRuleSetRequest
	1, // 5: pb.HTTPFirewallRuleSetService.createOrUpdateHTTPFirewallRuleSetFromConfig:output_type -> pb.CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse
	8, // 6: pb.HTTPFirewallRuleSetService.updateHTTPFirewallRuleSetIsOn:output_type -> pb.RPCSuccess
	4, // 7: pb.HTTPFirewallRuleSetService.findEnabledHTTPFirewallRuleSetConfig:output_type -> pb.FindEnabledHTTPFirewallRuleSetConfigResponse
	6, // 8: pb.HTTPFirewallRuleSetService.findEnabledHTTPFirewallRuleSet:output_type -> pb.FindEnabledHTTPFirewallRuleSetResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_http_firewall_rule_set_proto_init() }
func file_service_http_firewall_rule_set_proto_init() {
	if File_service_http_firewall_rule_set_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_http_firewall_rule_set_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_http_firewall_rule_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHTTPFirewallRuleSetIsOnRequest); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFirewallRuleSetConfigRequest); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFirewallRuleSetConfigResponse); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFirewallRuleSetRequest); i {
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
		file_service_http_firewall_rule_set_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledHTTPFirewallRuleSetResponse); i {
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
			RawDescriptor: file_service_http_firewall_rule_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_http_firewall_rule_set_proto_goTypes,
		DependencyIndexes: file_service_http_firewall_rule_set_proto_depIdxs,
		MessageInfos:      file_service_http_firewall_rule_set_proto_msgTypes,
	}.Build()
	File_service_http_firewall_rule_set_proto = out.File
	file_service_http_firewall_rule_set_proto_rawDesc = nil
	file_service_http_firewall_rule_set_proto_goTypes = nil
	file_service_http_firewall_rule_set_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HTTPFirewallRuleSetServiceClient is the client API for HTTPFirewallRuleSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPFirewallRuleSetServiceClient interface {
	// 根据配置创建或修改规则集
	CreateOrUpdateHTTPFirewallRuleSetFromConfig(ctx context.Context, in *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest, opts ...grpc.CallOption) (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse, error)
	// 设置开启状态
	UpdateHTTPFirewallRuleSetIsOn(ctx context.Context, in *UpdateHTTPFirewallRuleSetIsOnRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找规则集配置
	FindEnabledHTTPFirewallRuleSetConfig(ctx context.Context, in *FindEnabledHTTPFirewallRuleSetConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFirewallRuleSetConfigResponse, error)
	// 查找规则集信息
	FindEnabledHTTPFirewallRuleSet(ctx context.Context, in *FindEnabledHTTPFirewallRuleSetRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFirewallRuleSetResponse, error)
}

type hTTPFirewallRuleSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPFirewallRuleSetServiceClient(cc grpc.ClientConnInterface) HTTPFirewallRuleSetServiceClient {
	return &hTTPFirewallRuleSetServiceClient{cc}
}

func (c *hTTPFirewallRuleSetServiceClient) CreateOrUpdateHTTPFirewallRuleSetFromConfig(ctx context.Context, in *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest, opts ...grpc.CallOption) (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse, error) {
	out := new(CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFirewallRuleSetService/createOrUpdateHTTPFirewallRuleSetFromConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFirewallRuleSetServiceClient) UpdateHTTPFirewallRuleSetIsOn(ctx context.Context, in *UpdateHTTPFirewallRuleSetIsOnRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.HTTPFirewallRuleSetService/updateHTTPFirewallRuleSetIsOn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFirewallRuleSetServiceClient) FindEnabledHTTPFirewallRuleSetConfig(ctx context.Context, in *FindEnabledHTTPFirewallRuleSetConfigRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFirewallRuleSetConfigResponse, error) {
	out := new(FindEnabledHTTPFirewallRuleSetConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFirewallRuleSetService/findEnabledHTTPFirewallRuleSetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFirewallRuleSetServiceClient) FindEnabledHTTPFirewallRuleSet(ctx context.Context, in *FindEnabledHTTPFirewallRuleSetRequest, opts ...grpc.CallOption) (*FindEnabledHTTPFirewallRuleSetResponse, error) {
	out := new(FindEnabledHTTPFirewallRuleSetResponse)
	err := c.cc.Invoke(ctx, "/pb.HTTPFirewallRuleSetService/findEnabledHTTPFirewallRuleSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPFirewallRuleSetServiceServer is the server API for HTTPFirewallRuleSetService service.
type HTTPFirewallRuleSetServiceServer interface {
	// 根据配置创建或修改规则集
	CreateOrUpdateHTTPFirewallRuleSetFromConfig(context.Context, *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse, error)
	// 设置开启状态
	UpdateHTTPFirewallRuleSetIsOn(context.Context, *UpdateHTTPFirewallRuleSetIsOnRequest) (*RPCSuccess, error)
	// 查找规则集配置
	FindEnabledHTTPFirewallRuleSetConfig(context.Context, *FindEnabledHTTPFirewallRuleSetConfigRequest) (*FindEnabledHTTPFirewallRuleSetConfigResponse, error)
	// 查找规则集信息
	FindEnabledHTTPFirewallRuleSet(context.Context, *FindEnabledHTTPFirewallRuleSetRequest) (*FindEnabledHTTPFirewallRuleSetResponse, error)
}

// UnimplementedHTTPFirewallRuleSetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPFirewallRuleSetServiceServer struct {
}

func (*UnimplementedHTTPFirewallRuleSetServiceServer) CreateOrUpdateHTTPFirewallRuleSetFromConfig(context.Context, *CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest) (*CreateOrUpdateHTTPFirewallRuleSetFromConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateHTTPFirewallRuleSetFromConfig not implemented")
}
func (*UnimplementedHTTPFirewallRuleSetServiceServer) UpdateHTTPFirewallRuleSetIsOn(context.Context, *UpdateHTTPFirewallRuleSetIsOnRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHTTPFirewallRuleSetIsOn not implemented")
}
func (*UnimplementedHTTPFirewallRuleSetServiceServer) FindEnabledHTTPFirewallRuleSetConfig(context.Context, *FindEnabledHTTPFirewallRuleSetConfigRequest) (*FindEnabledHTTPFirewallRuleSetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPFirewallRuleSetConfig not implemented")
}
func (*UnimplementedHTTPFirewallRuleSetServiceServer) FindEnabledHTTPFirewallRuleSet(context.Context, *FindEnabledHTTPFirewallRuleSetRequest) (*FindEnabledHTTPFirewallRuleSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledHTTPFirewallRuleSet not implemented")
}

func RegisterHTTPFirewallRuleSetServiceServer(s *grpc.Server, srv HTTPFirewallRuleSetServiceServer) {
	s.RegisterService(&_HTTPFirewallRuleSetService_serviceDesc, srv)
}

func _HTTPFirewallRuleSetService_CreateOrUpdateHTTPFirewallRuleSetFromConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFirewallRuleSetServiceServer).CreateOrUpdateHTTPFirewallRuleSetFromConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFirewallRuleSetService/CreateOrUpdateHTTPFirewallRuleSetFromConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFirewallRuleSetServiceServer).CreateOrUpdateHTTPFirewallRuleSetFromConfig(ctx, req.(*CreateOrUpdateHTTPFirewallRuleSetFromConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFirewallRuleSetService_UpdateHTTPFirewallRuleSetIsOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHTTPFirewallRuleSetIsOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFirewallRuleSetServiceServer).UpdateHTTPFirewallRuleSetIsOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFirewallRuleSetService/UpdateHTTPFirewallRuleSetIsOn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFirewallRuleSetServiceServer).UpdateHTTPFirewallRuleSetIsOn(ctx, req.(*UpdateHTTPFirewallRuleSetIsOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFirewallRuleSetService_FindEnabledHTTPFirewallRuleSetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPFirewallRuleSetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFirewallRuleSetServiceServer).FindEnabledHTTPFirewallRuleSetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFirewallRuleSetService/FindEnabledHTTPFirewallRuleSetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFirewallRuleSetServiceServer).FindEnabledHTTPFirewallRuleSetConfig(ctx, req.(*FindEnabledHTTPFirewallRuleSetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFirewallRuleSetService_FindEnabledHTTPFirewallRuleSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledHTTPFirewallRuleSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFirewallRuleSetServiceServer).FindEnabledHTTPFirewallRuleSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HTTPFirewallRuleSetService/FindEnabledHTTPFirewallRuleSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFirewallRuleSetServiceServer).FindEnabledHTTPFirewallRuleSet(ctx, req.(*FindEnabledHTTPFirewallRuleSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTPFirewallRuleSetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HTTPFirewallRuleSetService",
	HandlerType: (*HTTPFirewallRuleSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createOrUpdateHTTPFirewallRuleSetFromConfig",
			Handler:    _HTTPFirewallRuleSetService_CreateOrUpdateHTTPFirewallRuleSetFromConfig_Handler,
		},
		{
			MethodName: "updateHTTPFirewallRuleSetIsOn",
			Handler:    _HTTPFirewallRuleSetService_UpdateHTTPFirewallRuleSetIsOn_Handler,
		},
		{
			MethodName: "findEnabledHTTPFirewallRuleSetConfig",
			Handler:    _HTTPFirewallRuleSetService_FindEnabledHTTPFirewallRuleSetConfig_Handler,
		},
		{
			MethodName: "findEnabledHTTPFirewallRuleSet",
			Handler:    _HTTPFirewallRuleSetService_FindEnabledHTTPFirewallRuleSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_http_firewall_rule_set.proto",
}
