// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dynamicpb_test

import (
	"testing"

	"github.com/CoinFlowwExchange/protobuf/proto"
	"github.com/CoinFlowwExchange/protobuf/reflect/protoreflect"
	"github.com/CoinFlowwExchange/protobuf/reflect/protoregistry"
	"github.com/CoinFlowwExchange/protobuf/testing/prototest"
	"github.com/CoinFlowwExchange/protobuf/types/dynamicpb"

	testpb "github.com/CoinFlowwExchange/protobuf/internal/testprotos/test"
	test3pb "github.com/CoinFlowwExchange/protobuf/internal/testprotos/test3"
)

func TestConformance(t *testing.T) {
	for _, message := range []proto.Message{
		(*testpb.TestAllTypes)(nil),
		(*test3pb.TestAllTypes)(nil),
		(*testpb.TestAllExtensions)(nil),
	} {
		mt := dynamicpb.NewMessageType(message.ProtoReflect().Descriptor())
		prototest.Message{}.Test(t, mt)
	}
}

func TestDynamicExtensions(t *testing.T) {
	for _, message := range []proto.Message{
		(*testpb.TestAllExtensions)(nil),
	} {
		mt := dynamicpb.NewMessageType(message.ProtoReflect().Descriptor())
		prototest.Message{
			Resolver: extResolver{},
		}.Test(t, mt)
	}
}

func TestDynamicEnums(t *testing.T) {
	for _, enum := range []protoreflect.Enum{
		testpb.TestAllTypes_FOO,
		test3pb.TestAllTypes_FOO,
	} {
		et := dynamicpb.NewEnumType(enum.Descriptor())
		prototest.Enum{}.Test(t, et)
	}
}

type extResolver struct{}

func (extResolver) FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error) {
	xt, err := protoregistry.GlobalTypes.FindExtensionByName(field)
	if err != nil {
		return nil, err
	}
	return dynamicpb.NewExtensionType(xt.TypeDescriptor().Descriptor()), nil
}

func (extResolver) FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error) {
	xt, err := protoregistry.GlobalTypes.FindExtensionByNumber(message, field)
	if err != nil {
		return nil, err
	}
	return dynamicpb.NewExtensionType(xt.TypeDescriptor().Descriptor()), nil
}

func (extResolver) RangeExtensionsByMessage(message protoreflect.FullName, f func(protoreflect.ExtensionType) bool) {
	protoregistry.GlobalTypes.RangeExtensionsByMessage(message, func(xt protoreflect.ExtensionType) bool {
		return f(dynamicpb.NewExtensionType(xt.TypeDescriptor().Descriptor()))
	})
}
