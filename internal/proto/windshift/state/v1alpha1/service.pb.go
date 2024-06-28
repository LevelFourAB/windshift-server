// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: windshift/state/v1alpha1/service.proto

package statev1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// EnsureStoreRequest creates or updates a state store.
type EnsureStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the store to create or update. Store names are case-sensitive
	// and should only contain the following characters:
	//
	// - `a` to `z`, `A` to `Z` and `0` to `9` are allowed.
	// - `_` and `-` are allowed for separating words.
	Store string `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
}

func (x *EnsureStoreRequest) Reset() {
	*x = EnsureStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnsureStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnsureStoreRequest) ProtoMessage() {}

func (x *EnsureStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnsureStoreRequest.ProtoReflect.Descriptor instead.
func (*EnsureStoreRequest) Descriptor() ([]byte, []int) {
	return file_windshift_state_v1alpha1_service_proto_rawDescGZIP(), []int{0}
}

func (x *EnsureStoreRequest) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

type EnsureStoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnsureStoreResponse) Reset() {
	*x = EnsureStoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnsureStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnsureStoreResponse) ProtoMessage() {}

func (x *EnsureStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnsureStoreResponse.ProtoReflect.Descriptor instead.
func (*EnsureStoreResponse) Descriptor() ([]byte, []int) {
	return file_windshift_state_v1alpha1_service_proto_rawDescGZIP(), []int{1}
}

// GetRequest is the message sent to retrieve a value from a store.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Store to retrieve the value from.
	Store string `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
	// Key to retrieve the value for.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_windshift_state_v1alpha1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// GetResponse is the message returned when retrieving a value from a store.
type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Timestamp of the last update to the key.
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=last_updated,json=lastUpdated,proto3,oneof" json:"last_updated,omitempty"`
	// *
	// The revision of the key.
	Revision uint64 `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
	// *
	// The value of the key.
	Value *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_windshift_state_v1alpha1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *GetResponse) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *GetResponse) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Store to set the value in.
	Store string `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
	// Key to set the value for. Keys are case-sensitive and should only contain
	// the following characters:
	//
	//   - `a` to `z`, `A` to `Z` and `0` to `9` are allowed.
	//   - `_` and `-` are allowed for separating words, but the use of camelCase
	//     is recommended.
	//   - `.` is allowed and used a hierarchy separator.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Value to set.
	Value *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// If set the operation will only succeed if the key does not already
	// exist in the store.
	CreateOnly *bool `protobuf:"varint,4,opt,name=create_only,json=createOnly,proto3,oneof" json:"create_only,omitempty"`
	// If set the operation will only succeed if the current revision of the
	// key matches the given revision.
	LastRevision *uint64 `protobuf:"varint,5,opt,name=last_revision,json=lastRevision,proto3,oneof" json:"last_revision,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_windshift_state_v1alpha1_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetRequest) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SetRequest) GetCreateOnly() bool {
	if x != nil && x.CreateOnly != nil {
		return *x.CreateOnly
	}
	return false
}

func (x *SetRequest) GetLastRevision() uint64 {
	if x != nil && x.LastRevision != nil {
		return *x.LastRevision
	}
	return 0
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The revision of the key.
	Revision uint64 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_windshift_state_v1alpha1_service_proto_rawDescGZIP(), []int{5}
}

func (x *SetResponse) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Store to delete the key from.
	Store string `protobuf:"bytes,1,opt,name=store,proto3" json:"store,omitempty"`
	// Key to delete.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// If set the operation will only succeed if the current revision of the
	// key matches the given revision.
	LastRevision *uint64 `protobuf:"varint,3,opt,name=last_revision,json=lastRevision,proto3,oneof" json:"last_revision,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_windshift_state_v1alpha1_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetStore() string {
	if x != nil {
		return x.Store
	}
	return ""
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DeleteRequest) GetLastRevision() uint64 {
	if x != nil && x.LastRevision != nil {
		return *x.LastRevision
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_windshift_state_v1alpha1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_windshift_state_v1alpha1_service_proto_rawDescGZIP(), []int{7}
}

var File_windshift_state_v1alpha1_service_proto protoreflect.FileDescriptor

var file_windshift_state_v1alpha1_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68,
	0x69, 0x66, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a,
	0x0a, 0x12, 0x45, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x6e,
	0x73, 0x75, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xb9, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x29, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x0d,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xff, 0x02, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x0b, 0x45,
	0x6e, 0x73, 0x75, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x2e, 0x77, 0x69, 0x6e,
	0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x77, 0x69, 0x6e, 0x64, 0x73,
	0x68, 0x69, 0x66, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x24,
	0x2e, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x03, 0x53,
	0x65, 0x74, 0x12, 0x24, 0x2e, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x69, 0x6e, 0x64, 0x73,
	0x68, 0x69, 0x66, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x77, 0x69, 0x6e, 0x64,
	0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8d, 0x02, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x66,
	0x6f, 0x75, 0x72, 0x61, 0x62, 0x2f, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x57,
	0x53, 0x58, 0xaa, 0x02, 0x18, 0x57, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x18,
	0x57, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x24, 0x57, 0x69, 0x6e, 0x64, 0x73,
	0x68, 0x69, 0x66, 0x74, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1a, 0x57, 0x69, 0x6e, 0x64, 0x73, 0x68, 0x69, 0x66, 0x74, 0x3a, 0x3a, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_windshift_state_v1alpha1_service_proto_rawDescOnce sync.Once
	file_windshift_state_v1alpha1_service_proto_rawDescData = file_windshift_state_v1alpha1_service_proto_rawDesc
)

func file_windshift_state_v1alpha1_service_proto_rawDescGZIP() []byte {
	file_windshift_state_v1alpha1_service_proto_rawDescOnce.Do(func() {
		file_windshift_state_v1alpha1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_windshift_state_v1alpha1_service_proto_rawDescData)
	})
	return file_windshift_state_v1alpha1_service_proto_rawDescData
}

var file_windshift_state_v1alpha1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_windshift_state_v1alpha1_service_proto_goTypes = []interface{}{
	(*EnsureStoreRequest)(nil),    // 0: windshift.state.v1alpha1.EnsureStoreRequest
	(*EnsureStoreResponse)(nil),   // 1: windshift.state.v1alpha1.EnsureStoreResponse
	(*GetRequest)(nil),            // 2: windshift.state.v1alpha1.GetRequest
	(*GetResponse)(nil),           // 3: windshift.state.v1alpha1.GetResponse
	(*SetRequest)(nil),            // 4: windshift.state.v1alpha1.SetRequest
	(*SetResponse)(nil),           // 5: windshift.state.v1alpha1.SetResponse
	(*DeleteRequest)(nil),         // 6: windshift.state.v1alpha1.DeleteRequest
	(*DeleteResponse)(nil),        // 7: windshift.state.v1alpha1.DeleteResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 9: google.protobuf.Any
}
var file_windshift_state_v1alpha1_service_proto_depIdxs = []int32{
	8, // 0: windshift.state.v1alpha1.GetResponse.last_updated:type_name -> google.protobuf.Timestamp
	9, // 1: windshift.state.v1alpha1.GetResponse.value:type_name -> google.protobuf.Any
	9, // 2: windshift.state.v1alpha1.SetRequest.value:type_name -> google.protobuf.Any
	0, // 3: windshift.state.v1alpha1.StateService.EnsureStore:input_type -> windshift.state.v1alpha1.EnsureStoreRequest
	2, // 4: windshift.state.v1alpha1.StateService.Get:input_type -> windshift.state.v1alpha1.GetRequest
	4, // 5: windshift.state.v1alpha1.StateService.Set:input_type -> windshift.state.v1alpha1.SetRequest
	6, // 6: windshift.state.v1alpha1.StateService.Delete:input_type -> windshift.state.v1alpha1.DeleteRequest
	1, // 7: windshift.state.v1alpha1.StateService.EnsureStore:output_type -> windshift.state.v1alpha1.EnsureStoreResponse
	3, // 8: windshift.state.v1alpha1.StateService.Get:output_type -> windshift.state.v1alpha1.GetResponse
	5, // 9: windshift.state.v1alpha1.StateService.Set:output_type -> windshift.state.v1alpha1.SetResponse
	7, // 10: windshift.state.v1alpha1.StateService.Delete:output_type -> windshift.state.v1alpha1.DeleteResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_windshift_state_v1alpha1_service_proto_init() }
func file_windshift_state_v1alpha1_service_proto_init() {
	if File_windshift_state_v1alpha1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_windshift_state_v1alpha1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnsureStoreRequest); i {
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
		file_windshift_state_v1alpha1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnsureStoreResponse); i {
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
		file_windshift_state_v1alpha1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_windshift_state_v1alpha1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_windshift_state_v1alpha1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_windshift_state_v1alpha1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResponse); i {
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
		file_windshift_state_v1alpha1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_windshift_state_v1alpha1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
	file_windshift_state_v1alpha1_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_windshift_state_v1alpha1_service_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_windshift_state_v1alpha1_service_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_windshift_state_v1alpha1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_windshift_state_v1alpha1_service_proto_goTypes,
		DependencyIndexes: file_windshift_state_v1alpha1_service_proto_depIdxs,
		MessageInfos:      file_windshift_state_v1alpha1_service_proto_msgTypes,
	}.Build()
	File_windshift_state_v1alpha1_service_proto = out.File
	file_windshift_state_v1alpha1_service_proto_rawDesc = nil
	file_windshift_state_v1alpha1_service_proto_goTypes = nil
	file_windshift_state_v1alpha1_service_proto_depIdxs = nil
}
