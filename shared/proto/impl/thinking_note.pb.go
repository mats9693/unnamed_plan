// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: thinking_note.proto

package rpc_impl

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

type ThinkingNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ThinkingNote) Reset() {
	*x = ThinkingNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote) ProtoMessage() {}

func (x *ThinkingNote) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote.ProtoReflect.Descriptor instead.
func (*ThinkingNote) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0}
}

type ThinkingNote_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId      string `protobuf:"bytes,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	WriteBy     string `protobuf:"bytes,2,opt,name=write_by,json=writeBy,proto3" json:"write_by,omitempty"`
	Topic       string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Content     string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	IsPublic    bool   `protobuf:"varint,5,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	UpdateTime  int64  `protobuf:"varint,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	CreatedTime int64  `protobuf:"varint,7,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
}

func (x *ThinkingNote_Data) Reset() {
	*x = ThinkingNote_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_Data) ProtoMessage() {}

func (x *ThinkingNote_Data) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_Data.ProtoReflect.Descriptor instead.
func (*ThinkingNote_Data) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ThinkingNote_Data) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *ThinkingNote_Data) GetWriteBy() string {
	if x != nil {
		return x.WriteBy
	}
	return ""
}

func (x *ThinkingNote_Data) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ThinkingNote_Data) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ThinkingNote_Data) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *ThinkingNote_Data) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *ThinkingNote_Data) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

type ThinkingNote_ListByWriterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorId string      `protobuf:"bytes,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	Page       *Pagination `protobuf:"bytes,15,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ThinkingNote_ListByWriterReq) Reset() {
	*x = ThinkingNote_ListByWriterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_ListByWriterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_ListByWriterReq) ProtoMessage() {}

func (x *ThinkingNote_ListByWriterReq) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_ListByWriterReq.ProtoReflect.Descriptor instead.
func (*ThinkingNote_ListByWriterReq) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ThinkingNote_ListByWriterReq) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

func (x *ThinkingNote_ListByWriterReq) GetPage() *Pagination {
	if x != nil {
		return x.Page
	}
	return nil
}

type ThinkingNote_ListByWriterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Notes []*ThinkingNote_Data `protobuf:"bytes,2,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *ThinkingNote_ListByWriterRes) Reset() {
	*x = ThinkingNote_ListByWriterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_ListByWriterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_ListByWriterRes) ProtoMessage() {}

func (x *ThinkingNote_ListByWriterRes) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_ListByWriterRes.ProtoReflect.Descriptor instead.
func (*ThinkingNote_ListByWriterRes) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ThinkingNote_ListByWriterRes) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ThinkingNote_ListByWriterRes) GetNotes() []*ThinkingNote_Data {
	if x != nil {
		return x.Notes
	}
	return nil
}

type ThinkingNote_ListPublicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorId string      `protobuf:"bytes,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	Page       *Pagination `protobuf:"bytes,15,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ThinkingNote_ListPublicReq) Reset() {
	*x = ThinkingNote_ListPublicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_ListPublicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_ListPublicReq) ProtoMessage() {}

func (x *ThinkingNote_ListPublicReq) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_ListPublicReq.ProtoReflect.Descriptor instead.
func (*ThinkingNote_ListPublicReq) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 3}
}

func (x *ThinkingNote_ListPublicReq) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

func (x *ThinkingNote_ListPublicReq) GetPage() *Pagination {
	if x != nil {
		return x.Page
	}
	return nil
}

type ThinkingNote_ListPublicRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Notes []*ThinkingNote_Data `protobuf:"bytes,2,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *ThinkingNote_ListPublicRes) Reset() {
	*x = ThinkingNote_ListPublicRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_ListPublicRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_ListPublicRes) ProtoMessage() {}

func (x *ThinkingNote_ListPublicRes) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_ListPublicRes.ProtoReflect.Descriptor instead.
func (*ThinkingNote_ListPublicRes) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 4}
}

func (x *ThinkingNote_ListPublicRes) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ThinkingNote_ListPublicRes) GetNotes() []*ThinkingNote_Data {
	if x != nil {
		return x.Notes
	}
	return nil
}

type ThinkingNote_CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorId string `protobuf:"bytes,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	Topic      string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Content    string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	IsPublic   bool   `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *ThinkingNote_CreateReq) Reset() {
	*x = ThinkingNote_CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_CreateReq) ProtoMessage() {}

func (x *ThinkingNote_CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_CreateReq.ProtoReflect.Descriptor instead.
func (*ThinkingNote_CreateReq) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 5}
}

func (x *ThinkingNote_CreateReq) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

func (x *ThinkingNote_CreateReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ThinkingNote_CreateReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ThinkingNote_CreateReq) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type ThinkingNote_CreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ThinkingNote_CreateRes) Reset() {
	*x = ThinkingNote_CreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_CreateRes) ProtoMessage() {}

func (x *ThinkingNote_CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_CreateRes.ProtoReflect.Descriptor instead.
func (*ThinkingNote_CreateRes) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 6}
}

type ThinkingNote_ModifyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorId string `protobuf:"bytes,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	NoteId     string `protobuf:"bytes,2,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Topic      string `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Content    string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	IsPublic   bool   `protobuf:"varint,6,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *ThinkingNote_ModifyReq) Reset() {
	*x = ThinkingNote_ModifyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_ModifyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_ModifyReq) ProtoMessage() {}

func (x *ThinkingNote_ModifyReq) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_ModifyReq.ProtoReflect.Descriptor instead.
func (*ThinkingNote_ModifyReq) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 7}
}

func (x *ThinkingNote_ModifyReq) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

func (x *ThinkingNote_ModifyReq) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

func (x *ThinkingNote_ModifyReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ThinkingNote_ModifyReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ThinkingNote_ModifyReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ThinkingNote_ModifyReq) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type ThinkingNote_ModifyRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ThinkingNote_ModifyRes) Reset() {
	*x = ThinkingNote_ModifyRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_ModifyRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_ModifyRes) ProtoMessage() {}

func (x *ThinkingNote_ModifyRes) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_ModifyRes.ProtoReflect.Descriptor instead.
func (*ThinkingNote_ModifyRes) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 8}
}

type ThinkingNote_DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperatorId string `protobuf:"bytes,1,opt,name=operator_id,json=operatorId,proto3" json:"operator_id,omitempty"`
	Password   string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	NoteId     string `protobuf:"bytes,3,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (x *ThinkingNote_DeleteReq) Reset() {
	*x = ThinkingNote_DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_DeleteReq) ProtoMessage() {}

func (x *ThinkingNote_DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_DeleteReq.ProtoReflect.Descriptor instead.
func (*ThinkingNote_DeleteReq) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 9}
}

func (x *ThinkingNote_DeleteReq) GetOperatorId() string {
	if x != nil {
		return x.OperatorId
	}
	return ""
}

func (x *ThinkingNote_DeleteReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ThinkingNote_DeleteReq) GetNoteId() string {
	if x != nil {
		return x.NoteId
	}
	return ""
}

type ThinkingNote_DeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ThinkingNote_DeleteRes) Reset() {
	*x = ThinkingNote_DeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thinking_note_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThinkingNote_DeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThinkingNote_DeleteRes) ProtoMessage() {}

func (x *ThinkingNote_DeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_thinking_note_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThinkingNote_DeleteRes.ProtoReflect.Descriptor instead.
func (*ThinkingNote_DeleteRes) Descriptor() ([]byte, []int) {
	return file_thinking_note_proto_rawDescGZIP(), []int{0, 10}
}

var File_thinking_note_proto protoreflect.FileDescriptor

var file_thinking_note_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e,
	0x6f, 0x74, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x86, 0x08, 0x0a, 0x0c, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x74, 0x65, 0x1a, 0xcb, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x1a, 0x5a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x1a, 0x5e, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e,
	0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x1a, 0x58, 0x0a, 0x0d,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a,
	0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x1a, 0x5c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x35, 0x0a,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74,
	0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x1a, 0x79, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x1a,
	0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x1a, 0xae, 0x01, 0x0a,
	0x09, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x1a, 0x0b, 0x0a,
	0x09, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x1a, 0x61, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x0b, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x32, 0xdb, 0x03, 0x0a, 0x0d, 0x49,
	0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x66, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x74,
	0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b,
	0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x60, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x12, 0x28, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74,
	0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x74,
	0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x52, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x24, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e,
	0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x06,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x24, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f,
	0x74, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x74,
	0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x73, 0x12, 0x54, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x74,
	0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x24, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74,
	0x65, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x42, 0x0e, 0x48, 0x03, 0x5a, 0x0a, 0x2e, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thinking_note_proto_rawDescOnce sync.Once
	file_thinking_note_proto_rawDescData = file_thinking_note_proto_rawDesc
)

func file_thinking_note_proto_rawDescGZIP() []byte {
	file_thinking_note_proto_rawDescOnce.Do(func() {
		file_thinking_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_thinking_note_proto_rawDescData)
	})
	return file_thinking_note_proto_rawDescData
}

var file_thinking_note_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_thinking_note_proto_goTypes = []interface{}{
	(*ThinkingNote)(nil),                 // 0: thinkingNote.ThinkingNote
	(*ThinkingNote_Data)(nil),            // 1: thinkingNote.ThinkingNote.Data
	(*ThinkingNote_ListByWriterReq)(nil), // 2: thinkingNote.ThinkingNote.ListByWriterReq
	(*ThinkingNote_ListByWriterRes)(nil), // 3: thinkingNote.ThinkingNote.ListByWriterRes
	(*ThinkingNote_ListPublicReq)(nil),   // 4: thinkingNote.ThinkingNote.ListPublicReq
	(*ThinkingNote_ListPublicRes)(nil),   // 5: thinkingNote.ThinkingNote.ListPublicRes
	(*ThinkingNote_CreateReq)(nil),       // 6: thinkingNote.ThinkingNote.CreateReq
	(*ThinkingNote_CreateRes)(nil),       // 7: thinkingNote.ThinkingNote.CreateRes
	(*ThinkingNote_ModifyReq)(nil),       // 8: thinkingNote.ThinkingNote.ModifyReq
	(*ThinkingNote_ModifyRes)(nil),       // 9: thinkingNote.ThinkingNote.ModifyRes
	(*ThinkingNote_DeleteReq)(nil),       // 10: thinkingNote.ThinkingNote.DeleteReq
	(*ThinkingNote_DeleteRes)(nil),       // 11: thinkingNote.ThinkingNote.DeleteRes
	(*Pagination)(nil),                   // 12: common.Pagination
}
var file_thinking_note_proto_depIdxs = []int32{
	12, // 0: thinkingNote.ThinkingNote.ListByWriterReq.page:type_name -> common.Pagination
	1,  // 1: thinkingNote.ThinkingNote.ListByWriterRes.notes:type_name -> thinkingNote.ThinkingNote.Data
	12, // 2: thinkingNote.ThinkingNote.ListPublicReq.page:type_name -> common.Pagination
	1,  // 3: thinkingNote.ThinkingNote.ListPublicRes.notes:type_name -> thinkingNote.ThinkingNote.Data
	2,  // 4: thinkingNote.IThinkingNote.ListByWriter:input_type -> thinkingNote.ThinkingNote.ListByWriterReq
	4,  // 5: thinkingNote.IThinkingNote.ListPublic:input_type -> thinkingNote.ThinkingNote.ListPublicReq
	6,  // 6: thinkingNote.IThinkingNote.Create:input_type -> thinkingNote.ThinkingNote.CreateReq
	8,  // 7: thinkingNote.IThinkingNote.Modify:input_type -> thinkingNote.ThinkingNote.ModifyReq
	10, // 8: thinkingNote.IThinkingNote.Delete:input_type -> thinkingNote.ThinkingNote.DeleteReq
	3,  // 9: thinkingNote.IThinkingNote.ListByWriter:output_type -> thinkingNote.ThinkingNote.ListByWriterRes
	5,  // 10: thinkingNote.IThinkingNote.ListPublic:output_type -> thinkingNote.ThinkingNote.ListPublicRes
	7,  // 11: thinkingNote.IThinkingNote.Create:output_type -> thinkingNote.ThinkingNote.CreateRes
	9,  // 12: thinkingNote.IThinkingNote.Modify:output_type -> thinkingNote.ThinkingNote.ModifyRes
	11, // 13: thinkingNote.IThinkingNote.Delete:output_type -> thinkingNote.ThinkingNote.DeleteRes
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_thinking_note_proto_init() }
func file_thinking_note_proto_init() {
	if File_thinking_note_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_thinking_note_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote); i {
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
		file_thinking_note_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_Data); i {
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
		file_thinking_note_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_ListByWriterReq); i {
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
		file_thinking_note_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_ListByWriterRes); i {
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
		file_thinking_note_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_ListPublicReq); i {
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
		file_thinking_note_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_ListPublicRes); i {
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
		file_thinking_note_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_CreateReq); i {
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
		file_thinking_note_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_CreateRes); i {
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
		file_thinking_note_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_ModifyReq); i {
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
		file_thinking_note_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_ModifyRes); i {
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
		file_thinking_note_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_DeleteReq); i {
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
		file_thinking_note_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThinkingNote_DeleteRes); i {
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
			RawDescriptor: file_thinking_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thinking_note_proto_goTypes,
		DependencyIndexes: file_thinking_note_proto_depIdxs,
		MessageInfos:      file_thinking_note_proto_msgTypes,
	}.Build()
	File_thinking_note_proto = out.File
	file_thinking_note_proto_rawDesc = nil
	file_thinking_note_proto_goTypes = nil
	file_thinking_note_proto_depIdxs = nil
}