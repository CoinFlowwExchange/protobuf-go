// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Benchmark messages for proto2.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datasets/google_message1/proto2/benchmark_message1_proto2.proto

package proto2

import (
	protoreflect "github.com/CoinFlowwExchange/protobuf/reflect/protoreflect"
	protoimpl "github.com/CoinFlowwExchange/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type GoogleMessage1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1   *string                   `protobuf:"bytes,1,req,name=field1" json:"field1,omitempty"`
	Field9   *string                   `protobuf:"bytes,9,opt,name=field9" json:"field9,omitempty"`
	Field18  *string                   `protobuf:"bytes,18,opt,name=field18" json:"field18,omitempty"`
	Field80  *bool                     `protobuf:"varint,80,opt,name=field80,def=0" json:"field80,omitempty"`
	Field81  *bool                     `protobuf:"varint,81,opt,name=field81,def=1" json:"field81,omitempty"`
	Field2   *int32                    `protobuf:"varint,2,req,name=field2" json:"field2,omitempty"`
	Field3   *int32                    `protobuf:"varint,3,req,name=field3" json:"field3,omitempty"`
	Field280 *int32                    `protobuf:"varint,280,opt,name=field280" json:"field280,omitempty"`
	Field6   *int32                    `protobuf:"varint,6,opt,name=field6,def=0" json:"field6,omitempty"`
	Field22  *int64                    `protobuf:"varint,22,opt,name=field22" json:"field22,omitempty"`
	Field4   *string                   `protobuf:"bytes,4,opt,name=field4" json:"field4,omitempty"`
	Field5   []uint64                  `protobuf:"fixed64,5,rep,name=field5" json:"field5,omitempty"`
	Field59  *bool                     `protobuf:"varint,59,opt,name=field59,def=0" json:"field59,omitempty"`
	Field7   *string                   `protobuf:"bytes,7,opt,name=field7" json:"field7,omitempty"`
	Field16  *int32                    `protobuf:"varint,16,opt,name=field16" json:"field16,omitempty"`
	Field130 *int32                    `protobuf:"varint,130,opt,name=field130,def=0" json:"field130,omitempty"`
	Field12  *bool                     `protobuf:"varint,12,opt,name=field12,def=1" json:"field12,omitempty"`
	Field17  *bool                     `protobuf:"varint,17,opt,name=field17,def=1" json:"field17,omitempty"`
	Field13  *bool                     `protobuf:"varint,13,opt,name=field13,def=1" json:"field13,omitempty"`
	Field14  *bool                     `protobuf:"varint,14,opt,name=field14,def=1" json:"field14,omitempty"`
	Field104 *int32                    `protobuf:"varint,104,opt,name=field104,def=0" json:"field104,omitempty"`
	Field100 *int32                    `protobuf:"varint,100,opt,name=field100,def=0" json:"field100,omitempty"`
	Field101 *int32                    `protobuf:"varint,101,opt,name=field101,def=0" json:"field101,omitempty"`
	Field102 *string                   `protobuf:"bytes,102,opt,name=field102" json:"field102,omitempty"`
	Field103 *string                   `protobuf:"bytes,103,opt,name=field103" json:"field103,omitempty"`
	Field29  *int32                    `protobuf:"varint,29,opt,name=field29,def=0" json:"field29,omitempty"`
	Field30  *bool                     `protobuf:"varint,30,opt,name=field30,def=0" json:"field30,omitempty"`
	Field60  *int32                    `protobuf:"varint,60,opt,name=field60,def=-1" json:"field60,omitempty"`
	Field271 *int32                    `protobuf:"varint,271,opt,name=field271,def=-1" json:"field271,omitempty"`
	Field272 *int32                    `protobuf:"varint,272,opt,name=field272,def=-1" json:"field272,omitempty"`
	Field150 *int32                    `protobuf:"varint,150,opt,name=field150" json:"field150,omitempty"`
	Field23  *int32                    `protobuf:"varint,23,opt,name=field23,def=0" json:"field23,omitempty"`
	Field24  *bool                     `protobuf:"varint,24,opt,name=field24,def=0" json:"field24,omitempty"`
	Field25  *int32                    `protobuf:"varint,25,opt,name=field25,def=0" json:"field25,omitempty"`
	Field15  *GoogleMessage1SubMessage `protobuf:"bytes,15,opt,name=field15" json:"field15,omitempty"`
	Field78  *bool                     `protobuf:"varint,78,opt,name=field78" json:"field78,omitempty"`
	Field67  *int32                    `protobuf:"varint,67,opt,name=field67,def=0" json:"field67,omitempty"`
	Field68  *int32                    `protobuf:"varint,68,opt,name=field68" json:"field68,omitempty"`
	Field128 *int32                    `protobuf:"varint,128,opt,name=field128,def=0" json:"field128,omitempty"`
	Field129 *string                   `protobuf:"bytes,129,opt,name=field129,def=xxxxxxxxxxxxxxxxxxxxx" json:"field129,omitempty"`
	Field131 *int32                    `protobuf:"varint,131,opt,name=field131,def=0" json:"field131,omitempty"`
}

// Default values for GoogleMessage1 fields.
const (
	Default_GoogleMessage1_Field80  = bool(false)
	Default_GoogleMessage1_Field81  = bool(true)
	Default_GoogleMessage1_Field6   = int32(0)
	Default_GoogleMessage1_Field59  = bool(false)
	Default_GoogleMessage1_Field130 = int32(0)
	Default_GoogleMessage1_Field12  = bool(true)
	Default_GoogleMessage1_Field17  = bool(true)
	Default_GoogleMessage1_Field13  = bool(true)
	Default_GoogleMessage1_Field14  = bool(true)
	Default_GoogleMessage1_Field104 = int32(0)
	Default_GoogleMessage1_Field100 = int32(0)
	Default_GoogleMessage1_Field101 = int32(0)
	Default_GoogleMessage1_Field29  = int32(0)
	Default_GoogleMessage1_Field30  = bool(false)
	Default_GoogleMessage1_Field60  = int32(-1)
	Default_GoogleMessage1_Field271 = int32(-1)
	Default_GoogleMessage1_Field272 = int32(-1)
	Default_GoogleMessage1_Field23  = int32(0)
	Default_GoogleMessage1_Field24  = bool(false)
	Default_GoogleMessage1_Field25  = int32(0)
	Default_GoogleMessage1_Field67  = int32(0)
	Default_GoogleMessage1_Field128 = int32(0)
	Default_GoogleMessage1_Field129 = string("xxxxxxxxxxxxxxxxxxxxx")
	Default_GoogleMessage1_Field131 = int32(0)
)

func (x *GoogleMessage1) Reset() {
	*x = GoogleMessage1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleMessage1) ProtoMessage() {}

func (x *GoogleMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleMessage1.ProtoReflect.Descriptor instead.
func (*GoogleMessage1) Descriptor() ([]byte, []int) {
	return file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *GoogleMessage1) GetField1() string {
	if x != nil && x.Field1 != nil {
		return *x.Field1
	}
	return ""
}

func (x *GoogleMessage1) GetField9() string {
	if x != nil && x.Field9 != nil {
		return *x.Field9
	}
	return ""
}

func (x *GoogleMessage1) GetField18() string {
	if x != nil && x.Field18 != nil {
		return *x.Field18
	}
	return ""
}

func (x *GoogleMessage1) GetField80() bool {
	if x != nil && x.Field80 != nil {
		return *x.Field80
	}
	return Default_GoogleMessage1_Field80
}

func (x *GoogleMessage1) GetField81() bool {
	if x != nil && x.Field81 != nil {
		return *x.Field81
	}
	return Default_GoogleMessage1_Field81
}

func (x *GoogleMessage1) GetField2() int32 {
	if x != nil && x.Field2 != nil {
		return *x.Field2
	}
	return 0
}

func (x *GoogleMessage1) GetField3() int32 {
	if x != nil && x.Field3 != nil {
		return *x.Field3
	}
	return 0
}

func (x *GoogleMessage1) GetField280() int32 {
	if x != nil && x.Field280 != nil {
		return *x.Field280
	}
	return 0
}

func (x *GoogleMessage1) GetField6() int32 {
	if x != nil && x.Field6 != nil {
		return *x.Field6
	}
	return Default_GoogleMessage1_Field6
}

func (x *GoogleMessage1) GetField22() int64 {
	if x != nil && x.Field22 != nil {
		return *x.Field22
	}
	return 0
}

func (x *GoogleMessage1) GetField4() string {
	if x != nil && x.Field4 != nil {
		return *x.Field4
	}
	return ""
}

func (x *GoogleMessage1) GetField5() []uint64 {
	if x != nil {
		return x.Field5
	}
	return nil
}

func (x *GoogleMessage1) GetField59() bool {
	if x != nil && x.Field59 != nil {
		return *x.Field59
	}
	return Default_GoogleMessage1_Field59
}

func (x *GoogleMessage1) GetField7() string {
	if x != nil && x.Field7 != nil {
		return *x.Field7
	}
	return ""
}

func (x *GoogleMessage1) GetField16() int32 {
	if x != nil && x.Field16 != nil {
		return *x.Field16
	}
	return 0
}

func (x *GoogleMessage1) GetField130() int32 {
	if x != nil && x.Field130 != nil {
		return *x.Field130
	}
	return Default_GoogleMessage1_Field130
}

func (x *GoogleMessage1) GetField12() bool {
	if x != nil && x.Field12 != nil {
		return *x.Field12
	}
	return Default_GoogleMessage1_Field12
}

func (x *GoogleMessage1) GetField17() bool {
	if x != nil && x.Field17 != nil {
		return *x.Field17
	}
	return Default_GoogleMessage1_Field17
}

func (x *GoogleMessage1) GetField13() bool {
	if x != nil && x.Field13 != nil {
		return *x.Field13
	}
	return Default_GoogleMessage1_Field13
}

func (x *GoogleMessage1) GetField14() bool {
	if x != nil && x.Field14 != nil {
		return *x.Field14
	}
	return Default_GoogleMessage1_Field14
}

func (x *GoogleMessage1) GetField104() int32 {
	if x != nil && x.Field104 != nil {
		return *x.Field104
	}
	return Default_GoogleMessage1_Field104
}

func (x *GoogleMessage1) GetField100() int32 {
	if x != nil && x.Field100 != nil {
		return *x.Field100
	}
	return Default_GoogleMessage1_Field100
}

func (x *GoogleMessage1) GetField101() int32 {
	if x != nil && x.Field101 != nil {
		return *x.Field101
	}
	return Default_GoogleMessage1_Field101
}

func (x *GoogleMessage1) GetField102() string {
	if x != nil && x.Field102 != nil {
		return *x.Field102
	}
	return ""
}

func (x *GoogleMessage1) GetField103() string {
	if x != nil && x.Field103 != nil {
		return *x.Field103
	}
	return ""
}

func (x *GoogleMessage1) GetField29() int32 {
	if x != nil && x.Field29 != nil {
		return *x.Field29
	}
	return Default_GoogleMessage1_Field29
}

func (x *GoogleMessage1) GetField30() bool {
	if x != nil && x.Field30 != nil {
		return *x.Field30
	}
	return Default_GoogleMessage1_Field30
}

func (x *GoogleMessage1) GetField60() int32 {
	if x != nil && x.Field60 != nil {
		return *x.Field60
	}
	return Default_GoogleMessage1_Field60
}

func (x *GoogleMessage1) GetField271() int32 {
	if x != nil && x.Field271 != nil {
		return *x.Field271
	}
	return Default_GoogleMessage1_Field271
}

func (x *GoogleMessage1) GetField272() int32 {
	if x != nil && x.Field272 != nil {
		return *x.Field272
	}
	return Default_GoogleMessage1_Field272
}

func (x *GoogleMessage1) GetField150() int32 {
	if x != nil && x.Field150 != nil {
		return *x.Field150
	}
	return 0
}

func (x *GoogleMessage1) GetField23() int32 {
	if x != nil && x.Field23 != nil {
		return *x.Field23
	}
	return Default_GoogleMessage1_Field23
}

func (x *GoogleMessage1) GetField24() bool {
	if x != nil && x.Field24 != nil {
		return *x.Field24
	}
	return Default_GoogleMessage1_Field24
}

func (x *GoogleMessage1) GetField25() int32 {
	if x != nil && x.Field25 != nil {
		return *x.Field25
	}
	return Default_GoogleMessage1_Field25
}

func (x *GoogleMessage1) GetField15() *GoogleMessage1SubMessage {
	if x != nil {
		return x.Field15
	}
	return nil
}

func (x *GoogleMessage1) GetField78() bool {
	if x != nil && x.Field78 != nil {
		return *x.Field78
	}
	return false
}

func (x *GoogleMessage1) GetField67() int32 {
	if x != nil && x.Field67 != nil {
		return *x.Field67
	}
	return Default_GoogleMessage1_Field67
}

func (x *GoogleMessage1) GetField68() int32 {
	if x != nil && x.Field68 != nil {
		return *x.Field68
	}
	return 0
}

func (x *GoogleMessage1) GetField128() int32 {
	if x != nil && x.Field128 != nil {
		return *x.Field128
	}
	return Default_GoogleMessage1_Field128
}

func (x *GoogleMessage1) GetField129() string {
	if x != nil && x.Field129 != nil {
		return *x.Field129
	}
	return Default_GoogleMessage1_Field129
}

func (x *GoogleMessage1) GetField131() int32 {
	if x != nil && x.Field131 != nil {
		return *x.Field131
	}
	return Default_GoogleMessage1_Field131
}

type GoogleMessage1SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1   *int32  `protobuf:"varint,1,opt,name=field1,def=0" json:"field1,omitempty"`
	Field2   *int32  `protobuf:"varint,2,opt,name=field2,def=0" json:"field2,omitempty"`
	Field3   *int32  `protobuf:"varint,3,opt,name=field3,def=0" json:"field3,omitempty"`
	Field15  *string `protobuf:"bytes,15,opt,name=field15" json:"field15,omitempty"`
	Field12  *bool   `protobuf:"varint,12,opt,name=field12,def=1" json:"field12,omitempty"`
	Field13  *int64  `protobuf:"varint,13,opt,name=field13" json:"field13,omitempty"`
	Field14  *int64  `protobuf:"varint,14,opt,name=field14" json:"field14,omitempty"`
	Field16  *int32  `protobuf:"varint,16,opt,name=field16" json:"field16,omitempty"`
	Field19  *int32  `protobuf:"varint,19,opt,name=field19,def=2" json:"field19,omitempty"`
	Field20  *bool   `protobuf:"varint,20,opt,name=field20,def=1" json:"field20,omitempty"`
	Field28  *bool   `protobuf:"varint,28,opt,name=field28,def=1" json:"field28,omitempty"`
	Field21  *uint64 `protobuf:"fixed64,21,opt,name=field21" json:"field21,omitempty"`
	Field22  *int32  `protobuf:"varint,22,opt,name=field22" json:"field22,omitempty"`
	Field23  *bool   `protobuf:"varint,23,opt,name=field23,def=0" json:"field23,omitempty"`
	Field206 *bool   `protobuf:"varint,206,opt,name=field206,def=0" json:"field206,omitempty"`
	Field203 *uint32 `protobuf:"fixed32,203,opt,name=field203" json:"field203,omitempty"`
	Field204 *int32  `protobuf:"varint,204,opt,name=field204" json:"field204,omitempty"`
	Field205 *string `protobuf:"bytes,205,opt,name=field205" json:"field205,omitempty"`
	Field207 *uint64 `protobuf:"varint,207,opt,name=field207" json:"field207,omitempty"`
	Field300 *uint64 `protobuf:"varint,300,opt,name=field300" json:"field300,omitempty"`
}

// Default values for GoogleMessage1SubMessage fields.
const (
	Default_GoogleMessage1SubMessage_Field1   = int32(0)
	Default_GoogleMessage1SubMessage_Field2   = int32(0)
	Default_GoogleMessage1SubMessage_Field3   = int32(0)
	Default_GoogleMessage1SubMessage_Field12  = bool(true)
	Default_GoogleMessage1SubMessage_Field19  = int32(2)
	Default_GoogleMessage1SubMessage_Field20  = bool(true)
	Default_GoogleMessage1SubMessage_Field28  = bool(true)
	Default_GoogleMessage1SubMessage_Field23  = bool(false)
	Default_GoogleMessage1SubMessage_Field206 = bool(false)
)

func (x *GoogleMessage1SubMessage) Reset() {
	*x = GoogleMessage1SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoogleMessage1SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoogleMessage1SubMessage) ProtoMessage() {}

func (x *GoogleMessage1SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoogleMessage1SubMessage.ProtoReflect.Descriptor instead.
func (*GoogleMessage1SubMessage) Descriptor() ([]byte, []int) {
	return file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescGZIP(), []int{1}
}

func (x *GoogleMessage1SubMessage) GetField1() int32 {
	if x != nil && x.Field1 != nil {
		return *x.Field1
	}
	return Default_GoogleMessage1SubMessage_Field1
}

func (x *GoogleMessage1SubMessage) GetField2() int32 {
	if x != nil && x.Field2 != nil {
		return *x.Field2
	}
	return Default_GoogleMessage1SubMessage_Field2
}

func (x *GoogleMessage1SubMessage) GetField3() int32 {
	if x != nil && x.Field3 != nil {
		return *x.Field3
	}
	return Default_GoogleMessage1SubMessage_Field3
}

func (x *GoogleMessage1SubMessage) GetField15() string {
	if x != nil && x.Field15 != nil {
		return *x.Field15
	}
	return ""
}

func (x *GoogleMessage1SubMessage) GetField12() bool {
	if x != nil && x.Field12 != nil {
		return *x.Field12
	}
	return Default_GoogleMessage1SubMessage_Field12
}

func (x *GoogleMessage1SubMessage) GetField13() int64 {
	if x != nil && x.Field13 != nil {
		return *x.Field13
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField14() int64 {
	if x != nil && x.Field14 != nil {
		return *x.Field14
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField16() int32 {
	if x != nil && x.Field16 != nil {
		return *x.Field16
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField19() int32 {
	if x != nil && x.Field19 != nil {
		return *x.Field19
	}
	return Default_GoogleMessage1SubMessage_Field19
}

func (x *GoogleMessage1SubMessage) GetField20() bool {
	if x != nil && x.Field20 != nil {
		return *x.Field20
	}
	return Default_GoogleMessage1SubMessage_Field20
}

func (x *GoogleMessage1SubMessage) GetField28() bool {
	if x != nil && x.Field28 != nil {
		return *x.Field28
	}
	return Default_GoogleMessage1SubMessage_Field28
}

func (x *GoogleMessage1SubMessage) GetField21() uint64 {
	if x != nil && x.Field21 != nil {
		return *x.Field21
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField22() int32 {
	if x != nil && x.Field22 != nil {
		return *x.Field22
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField23() bool {
	if x != nil && x.Field23 != nil {
		return *x.Field23
	}
	return Default_GoogleMessage1SubMessage_Field23
}

func (x *GoogleMessage1SubMessage) GetField206() bool {
	if x != nil && x.Field206 != nil {
		return *x.Field206
	}
	return Default_GoogleMessage1SubMessage_Field206
}

func (x *GoogleMessage1SubMessage) GetField203() uint32 {
	if x != nil && x.Field203 != nil {
		return *x.Field203
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField204() int32 {
	if x != nil && x.Field204 != nil {
		return *x.Field204
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField205() string {
	if x != nil && x.Field205 != nil {
		return *x.Field205
	}
	return ""
}

func (x *GoogleMessage1SubMessage) GetField207() uint64 {
	if x != nil && x.Field207 != nil {
		return *x.Field207
	}
	return 0
}

func (x *GoogleMessage1SubMessage) GetField300() uint64 {
	if x != nil && x.Field300 != nil {
		return *x.Field300
	}
	return 0
}

var File_datasets_google_message1_proto2_benchmark_message1_proto2_proto protoreflect.FileDescriptor

var file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x22, 0xf7, 0x09, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x38, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x38, 0x12, 0x1f, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x30, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x38, 0x30, 0x12, 0x1e, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x31, 0x18, 0x51, 0x20,
	0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x38, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x33, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x33, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x38, 0x30, 0x18, 0x98,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x38, 0x30, 0x12,
	0x19, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x01, 0x30, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x32, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x34, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x34, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x18, 0x05, 0x20, 0x03, 0x28, 0x06, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x35, 0x12, 0x1f, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x35, 0x39, 0x18,
	0x3b, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x35, 0x39, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x36, 0x12, 0x1e, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x33, 0x30, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x30, 0x12, 0x1e, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x12, 0x1e, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x37, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x37, 0x12, 0x1e, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x33, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x12, 0x1e, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x12, 0x1d, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x30, 0x34, 0x18, 0x68, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x30, 0x34, 0x12, 0x1d, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x30, 0x30, 0x18, 0x64, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x30, 0x30, 0x12, 0x1d, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x31, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x30, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x32,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x32,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x33, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x39, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x39, 0x12, 0x1f, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x33, 0x30, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73,
	0x65, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x12, 0x1c, 0x0a, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x36, 0x30, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x2d, 0x31, 0x52,
	0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x30, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0x37, 0x31, 0x18, 0x8f, 0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x2d, 0x31, 0x52,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x31, 0x12, 0x1f, 0x0a, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x37, 0x32, 0x18, 0x90, 0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x2d, 0x31,
	0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x37, 0x32, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x35, 0x30, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x30, 0x12, 0x1b, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x33, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x32, 0x33, 0x12, 0x1f, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x34, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x34, 0x12, 0x1b, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x35,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x35, 0x12, 0x45, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x37, 0x38, 0x18, 0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x37, 0x38, 0x12, 0x1b, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x37, 0x18, 0x43,
	0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x37,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x38, 0x18, 0x44, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x36, 0x38, 0x12, 0x1e, 0x0a, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x32, 0x38, 0x18, 0x80, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30,
	0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x38, 0x12, 0x32, 0x0a, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x32, 0x39, 0x18, 0x81, 0x01, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x15, 0x78,
	0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78, 0x78,
	0x78, 0x78, 0x78, 0x78, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x39, 0x12, 0x1e,
	0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x31, 0x18, 0x83, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33, 0x31, 0x22, 0xda,
	0x04, 0x0a, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x31, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x19, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x30, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x12, 0x19, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x3a, 0x01, 0x30, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x35, 0x12, 0x1e, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x32, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x33, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x33,
	0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x36, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x36, 0x12, 0x1b, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x39, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01, 0x32, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31,
	0x39, 0x12, 0x1e, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x30, 0x12, 0x1e, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x38, 0x18, 0x1c, 0x20, 0x01,
	0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x38, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x32, 0x12, 0x1f, 0x0a, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x33,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x07, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x32, 0x33, 0x12, 0x22, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x30, 0x36, 0x18, 0xce, 0x01, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65,
	0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x36, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x30, 0x33, 0x18, 0xcb, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x08, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x33, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x32, 0x30, 0x34, 0x18, 0xcc, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0x30, 0x34, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x35,
	0x18, 0xcd, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30,
	0x35, 0x12, 0x1b, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x37, 0x18, 0xcf, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x30, 0x37, 0x12, 0x1b,
	0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x30, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x30, 0x30, 0x42, 0x25, 0x0a, 0x1e, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x48, 0x01, 0xf8,
	0x01, 0x01,
}

var (
	file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescOnce sync.Once
	file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescData = file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDesc
)

func file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescGZIP() []byte {
	file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescOnce.Do(func() {
		file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescData)
	})
	return file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDescData
}

var file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_goTypes = []interface{}{
	(*GoogleMessage1)(nil),           // 0: benchmarks.proto2.GoogleMessage1
	(*GoogleMessage1SubMessage)(nil), // 1: benchmarks.proto2.GoogleMessage1SubMessage
}
var file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_depIdxs = []int32{
	1, // 0: benchmarks.proto2.GoogleMessage1.field15:type_name -> benchmarks.proto2.GoogleMessage1SubMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_init() }
func file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_init() {
	if File_datasets_google_message1_proto2_benchmark_message1_proto2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleMessage1); i {
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
		file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoogleMessage1SubMessage); i {
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
			RawDescriptor: file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_goTypes,
		DependencyIndexes: file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_depIdxs,
		MessageInfos:      file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_msgTypes,
	}.Build()
	File_datasets_google_message1_proto2_benchmark_message1_proto2_proto = out.File
	file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_rawDesc = nil
	file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_goTypes = nil
	file_datasets_google_message1_proto2_benchmark_message1_proto2_proto_depIdxs = nil
}
