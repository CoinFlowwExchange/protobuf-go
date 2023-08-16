// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-protos. DO NOT EDIT.

package genid

import (
	protoreflect "github.com/CoinFlowwExchange/protobuf/reflect/protoreflect"
)

const File_google_protobuf_api_proto = "google/protobuf/api.proto"

// Names for google.protobuf.Api.
const (
	Api_message_name     protoreflect.Name     = "Api"
	Api_message_fullname protoreflect.FullName = "google.protobuf.Api"
)

// Field names for google.protobuf.Api.
const (
	Api_Name_field_name          protoreflect.Name = "name"
	Api_Methods_field_name       protoreflect.Name = "methods"
	Api_Options_field_name       protoreflect.Name = "options"
	Api_Version_field_name       protoreflect.Name = "version"
	Api_SourceContext_field_name protoreflect.Name = "source_context"
	Api_Mixins_field_name        protoreflect.Name = "mixins"
	Api_Syntax_field_name        protoreflect.Name = "syntax"

	Api_Name_field_fullname          protoreflect.FullName = "google.protobuf.Api.name"
	Api_Methods_field_fullname       protoreflect.FullName = "google.protobuf.Api.methods"
	Api_Options_field_fullname       protoreflect.FullName = "google.protobuf.Api.options"
	Api_Version_field_fullname       protoreflect.FullName = "google.protobuf.Api.version"
	Api_SourceContext_field_fullname protoreflect.FullName = "google.protobuf.Api.source_context"
	Api_Mixins_field_fullname        protoreflect.FullName = "google.protobuf.Api.mixins"
	Api_Syntax_field_fullname        protoreflect.FullName = "google.protobuf.Api.syntax"
)

// Field numbers for google.protobuf.Api.
const (
	Api_Name_field_number          protoreflect.FieldNumber = 1
	Api_Methods_field_number       protoreflect.FieldNumber = 2
	Api_Options_field_number       protoreflect.FieldNumber = 3
	Api_Version_field_number       protoreflect.FieldNumber = 4
	Api_SourceContext_field_number protoreflect.FieldNumber = 5
	Api_Mixins_field_number        protoreflect.FieldNumber = 6
	Api_Syntax_field_number        protoreflect.FieldNumber = 7
)

// Names for google.protobuf.Method.
const (
	Method_message_name     protoreflect.Name     = "Method"
	Method_message_fullname protoreflect.FullName = "google.protobuf.Method"
)

// Field names for google.protobuf.Method.
const (
	Method_Name_field_name              protoreflect.Name = "name"
	Method_RequestTypeUrl_field_name    protoreflect.Name = "request_type_url"
	Method_RequestStreaming_field_name  protoreflect.Name = "request_streaming"
	Method_ResponseTypeUrl_field_name   protoreflect.Name = "response_type_url"
	Method_ResponseStreaming_field_name protoreflect.Name = "response_streaming"
	Method_Options_field_name           protoreflect.Name = "options"
	Method_Syntax_field_name            protoreflect.Name = "syntax"

	Method_Name_field_fullname              protoreflect.FullName = "google.protobuf.Method.name"
	Method_RequestTypeUrl_field_fullname    protoreflect.FullName = "google.protobuf.Method.request_type_url"
	Method_RequestStreaming_field_fullname  protoreflect.FullName = "google.protobuf.Method.request_streaming"
	Method_ResponseTypeUrl_field_fullname   protoreflect.FullName = "google.protobuf.Method.response_type_url"
	Method_ResponseStreaming_field_fullname protoreflect.FullName = "google.protobuf.Method.response_streaming"
	Method_Options_field_fullname           protoreflect.FullName = "google.protobuf.Method.options"
	Method_Syntax_field_fullname            protoreflect.FullName = "google.protobuf.Method.syntax"
)

// Field numbers for google.protobuf.Method.
const (
	Method_Name_field_number              protoreflect.FieldNumber = 1
	Method_RequestTypeUrl_field_number    protoreflect.FieldNumber = 2
	Method_RequestStreaming_field_number  protoreflect.FieldNumber = 3
	Method_ResponseTypeUrl_field_number   protoreflect.FieldNumber = 4
	Method_ResponseStreaming_field_number protoreflect.FieldNumber = 5
	Method_Options_field_number           protoreflect.FieldNumber = 6
	Method_Syntax_field_number            protoreflect.FieldNumber = 7
)

// Names for google.protobuf.Mixin.
const (
	Mixin_message_name     protoreflect.Name     = "Mixin"
	Mixin_message_fullname protoreflect.FullName = "google.protobuf.Mixin"
)

// Field names for google.protobuf.Mixin.
const (
	Mixin_Name_field_name protoreflect.Name = "name"
	Mixin_Root_field_name protoreflect.Name = "root"

	Mixin_Name_field_fullname protoreflect.FullName = "google.protobuf.Mixin.name"
	Mixin_Root_field_fullname protoreflect.FullName = "google.protobuf.Mixin.root"
)

// Field numbers for google.protobuf.Mixin.
const (
	Mixin_Name_field_number protoreflect.FieldNumber = 1
	Mixin_Root_field_number protoreflect.FieldNumber = 2
)
