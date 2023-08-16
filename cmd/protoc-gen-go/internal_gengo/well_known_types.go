// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal_gengo

import (
	"strings"

	"github.com/CoinFlowwExchange/protobuf/compiler/protogen"
	"github.com/CoinFlowwExchange/protobuf/internal/genid"
)

// Specialized support for well-known types are hard-coded into the generator
// as opposed to being injected in adjacent .go sources in the generated package
// in order to support specialized build systems like Bazel that always generate
// dynamically from the source .proto files.

func genPackageKnownComment(f *fileInfo) protogen.Comments {
	switch f.Desc.Path() {
	case genid.File_google_protobuf_any_proto:
		return ` Package anypb contains generated types for ` + genid.File_google_protobuf_any_proto + `.

 The Any message is a dynamic representation of any other message value.
 It is functionally a tuple of the full name of the remote message type and
 the serialized bytes of the remote message value.


 Constructing an Any

 An Any message containing another message value is constructed using New:

	any, err := anypb.New(m)
	if err != nil {
		... // handle error
	}
	... // make use of any


 Unmarshaling an Any

 With a populated Any message, the underlying message can be serialized into
 a remote concrete message value in a few ways.

 If the exact concrete type is known, then a new (or pre-existing) instance
 of that message can be passed to the UnmarshalTo method:

	m := new(foopb.MyMessage)
	if err := any.UnmarshalTo(m); err != nil {
		... // handle error
	}
	... // make use of m

 If the exact concrete type is not known, then the UnmarshalNew method can be
 used to unmarshal the contents into a new instance of the remote message type:

	m, err := any.UnmarshalNew()
	if err != nil {
		... // handle error
	}
	... // make use of m

 UnmarshalNew uses the global type registry to resolve the message type and
 construct a new instance of that message to unmarshal into. In order for a
 message type to appear in the global registry, the Go type representing that
 protobuf message type must be linked into the Go binary. For messages
 generated by protoc-gen-go, this is achieved through an import of the
 generated Go package representing a .proto file.

 A common pattern with UnmarshalNew is to use a type switch with the resulting
 proto.Message value:

	switch m := m.(type) {
	case *foopb.MyMessage:
		... // make use of m as a *foopb.MyMessage
	case *barpb.OtherMessage:
		... // make use of m as a *barpb.OtherMessage
	case *bazpb.SomeMessage:
		... // make use of m as a *bazpb.SomeMessage
	}

 This pattern ensures that the generated packages containing the message types
 listed in the case clauses are linked into the Go binary and therefore also
 registered in the global registry.


 Type checking an Any

 In order to type check whether an Any message represents some other message,
 then use the MessageIs method:

	if any.MessageIs((*foopb.MyMessage)(nil)) {
		... // make use of any, knowing that it contains a foopb.MyMessage
	}

 The MessageIs method can also be used with an allocated instance of the target
 message type if the intention is to unmarshal into it if the type matches:

	m := new(foopb.MyMessage)
	if any.MessageIs(m) {
		if err := any.UnmarshalTo(m); err != nil {
			... // handle error
		}
		... // make use of m
	}

`
	case genid.File_google_protobuf_timestamp_proto:
		return ` Package timestamppb contains generated types for ` + genid.File_google_protobuf_timestamp_proto + `.

 The Timestamp message represents a timestamp,
 an instant in time since the Unix epoch (January 1st, 1970).


 Conversion to a Go Time

 The AsTime method can be used to convert a Timestamp message to a
 standard Go time.Time value in UTC:

	t := ts.AsTime()
	... // make use of t as a time.Time

 Converting to a time.Time is a common operation so that the extensive
 set of time-based operations provided by the time package can be leveraged.
 See https://golang.org/pkg/time for more information.

 The AsTime method performs the conversion on a best-effort basis. Timestamps
 with denormal values (e.g., nanoseconds beyond 0 and 99999999, inclusive)
 are normalized during the conversion to a time.Time. To manually check for
 invalid Timestamps per the documented limitations in timestamp.proto,
 additionally call the CheckValid method:

	if err := ts.CheckValid(); err != nil {
		... // handle error
	}


 Conversion from a Go Time

 The timestamppb.New function can be used to construct a Timestamp message
 from a standard Go time.Time value:

	ts := timestamppb.New(t)
	... // make use of ts as a *timestamppb.Timestamp

 In order to construct a Timestamp representing the current time, use Now:

	ts := timestamppb.Now()
	... // make use of ts as a *timestamppb.Timestamp

`
	case genid.File_google_protobuf_duration_proto:
		return ` Package durationpb contains generated types for ` + genid.File_google_protobuf_duration_proto + `.

 The Duration message represents a signed span of time.


 Conversion to a Go Duration

 The AsDuration method can be used to convert a Duration message to a
 standard Go time.Duration value:

	d := dur.AsDuration()
	... // make use of d as a time.Duration

 Converting to a time.Duration is a common operation so that the extensive
 set of time-based operations provided by the time package can be leveraged.
 See https://golang.org/pkg/time for more information.

 The AsDuration method performs the conversion on a best-effort basis.
 Durations with denormal values (e.g., nanoseconds beyond -99999999 and
 +99999999, inclusive; or seconds and nanoseconds with opposite signs)
 are normalized during the conversion to a time.Duration. To manually check for
 invalid Duration per the documented limitations in duration.proto,
 additionally call the CheckValid method:

	if err := dur.CheckValid(); err != nil {
		... // handle error
	}

 Note that the documented limitations in duration.proto does not protect a
 Duration from overflowing the representable range of a time.Duration in Go.
 The AsDuration method uses saturation arithmetic such that an overflow clamps
 the resulting value to the closest representable value (e.g., math.MaxInt64
 for positive overflow and math.MinInt64 for negative overflow).


 Conversion from a Go Duration

 The durationpb.New function can be used to construct a Duration message
 from a standard Go time.Duration value:

	dur := durationpb.New(d)
	... // make use of d as a *durationpb.Duration

`
	case genid.File_google_protobuf_struct_proto:
		return ` Package structpb contains generated types for ` + genid.File_google_protobuf_struct_proto + `.

 The messages (i.e., Value, Struct, and ListValue) defined in struct.proto are
 used to represent arbitrary JSON. The Value message represents a JSON value,
 the Struct message represents a JSON object, and the ListValue message
 represents a JSON array. See https://json.org for more information.

 The Value, Struct, and ListValue types have generated MarshalJSON and
 UnmarshalJSON methods such that they serialize JSON equivalent to what the
 messages themselves represent. Use of these types with the
 "github.com/CoinFlowwExchange/protobuf/encoding/protojson" package
 ensures that they will be serialized as their JSON equivalent.

 # Conversion to and from a Go interface

 The standard Go "encoding/json" package has functionality to serialize
 arbitrary types to a large degree. The Value.AsInterface, Struct.AsMap, and
 ListValue.AsSlice methods can convert the protobuf message representation into
 a form represented by interface{}, map[string]interface{}, and []interface{}.
 This form can be used with other packages that operate on such data structures
 and also directly with the standard json package.

 In order to convert the interface{}, map[string]interface{}, and []interface{}
 forms back as Value, Struct, and ListValue messages, use the NewStruct,
 NewList, and NewValue constructor functions.

 # Example usage

 Consider the following example JSON object:

	{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
			"streetAddress": "21 2nd Street",
			"city": "New York",
			"state": "NY",
			"postalCode": "10021-3100"
		},
		"phoneNumbers": [
			{
				"type": "home",
				"number": "212 555-1234"
			},
			{
				"type": "office",
				"number": "646 555-4567"
			}
		],
		"children": [],
		"spouse": null
	}

 To construct a Value message representing the above JSON object:

	m, err := structpb.NewValue(map[string]interface{}{
		"firstName": "John",
		"lastName":  "Smith",
		"isAlive":   true,
		"age":       27,
		"address": map[string]interface{}{
			"streetAddress": "21 2nd Street",
			"city":          "New York",
			"state":         "NY",
			"postalCode":    "10021-3100",
		},
		"phoneNumbers": []interface{}{
			map[string]interface{}{
				"type":   "home",
				"number": "212 555-1234",
			},
			map[string]interface{}{
				"type":   "office",
				"number": "646 555-4567",
			},
		},
		"children": []interface{}{},
		"spouse":   nil,
	})
	if err != nil {
		... // handle error
	}
	... // make use of m as a *structpb.Value
`
	case genid.File_google_protobuf_field_mask_proto:
		return ` Package fieldmaskpb contains generated types for ` + genid.File_google_protobuf_field_mask_proto + `.

 The FieldMask message represents a set of symbolic field paths.
 The paths are specific to some target message type,
 which is not stored within the FieldMask message itself.


 Constructing a FieldMask

 The New function is used construct a FieldMask:

	var messageType *descriptorpb.DescriptorProto
	fm, err := fieldmaskpb.New(messageType, "field.name", "field.number")
	if err != nil {
		... // handle error
	}
	... // make use of fm

 The "field.name" and "field.number" paths are valid paths according to the
 google.protobuf.DescriptorProto message. Use of a path that does not correlate
 to valid fields reachable from DescriptorProto would result in an error.

 Once a FieldMask message has been constructed,
 the Append method can be used to insert additional paths to the path set:

	var messageType *descriptorpb.DescriptorProto
	if err := fm.Append(messageType, "options"); err != nil {
		... // handle error
	}


 Type checking a FieldMask

 In order to verify that a FieldMask represents a set of fields that are
 reachable from some target message type, use the IsValid method:

	var messageType *descriptorpb.DescriptorProto
	if fm.IsValid(messageType) {
		... // make use of fm
	}

 IsValid needs to be passed the target message type as an input since the
 FieldMask message itself does not store the message type that the set of paths
 are for.
`
	default:
		return ""
	}
}

func genMessageKnownFunctions(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	switch m.Desc.FullName() {
	case genid.Any_message_fullname:
		g.P("// New marshals src into a new Any instance.")
		g.P("func New(src ", protoPackage.Ident("Message"), ") (*Any, error) {")
		g.P("	dst := new(Any)")
		g.P("	if err := dst.MarshalFrom(src); err != nil {")
		g.P("		return nil, err")
		g.P("	}")
		g.P("	return dst, nil")
		g.P("}")
		g.P()

		g.P("// MarshalFrom marshals src into dst as the underlying message")
		g.P("// using the provided marshal options.")
		g.P("//")
		g.P("// If no options are specified, call dst.MarshalFrom instead.")
		g.P("func MarshalFrom(dst *Any, src ", protoPackage.Ident("Message"), ", opts ", protoPackage.Ident("MarshalOptions"), ") error {")
		g.P("	const urlPrefix = \"type.googleapis.com/\"")
		g.P("	if src == nil {")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"invalid nil source message\")")
		g.P("	}")
		g.P("	b, err := opts.Marshal(src)")
		g.P("	if err != nil {")
		g.P("		return err")
		g.P("	}")
		g.P("	dst.TypeUrl = urlPrefix + string(src.ProtoReflect().Descriptor().FullName())")
		g.P("	dst.Value = b")
		g.P("	return nil")
		g.P("}")
		g.P()

		g.P("// UnmarshalTo unmarshals the underlying message from src into dst")
		g.P("// using the provided unmarshal options.")
		g.P("// It reports an error if dst is not of the right message type.")
		g.P("//")
		g.P("// If no options are specified, call src.UnmarshalTo instead.")
		g.P("func UnmarshalTo(src *Any, dst ", protoPackage.Ident("Message"), ", opts ", protoPackage.Ident("UnmarshalOptions"), ") error {")
		g.P("	if src == nil {")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"invalid nil source message\")")
		g.P("	}")
		g.P("	if !src.MessageIs(dst) {")
		g.P("		got := dst.ProtoReflect().Descriptor().FullName()")
		g.P("		want := src.MessageName()")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"mismatched message type: got %q, want %q\", got, want)")
		g.P("	}")
		g.P("	return opts.Unmarshal(src.GetValue(), dst)")
		g.P("}")
		g.P()

		g.P("// UnmarshalNew unmarshals the underlying message from src into dst,")
		g.P("// which is newly created message using a type resolved from the type URL.")
		g.P("// The message type is resolved according to opt.Resolver,")
		g.P("// which should implement protoregistry.MessageTypeResolver.")
		g.P("// It reports an error if the underlying message type could not be resolved.")
		g.P("//")
		g.P("// If no options are specified, call src.UnmarshalNew instead.")
		g.P("func UnmarshalNew(src *Any, opts ", protoPackage.Ident("UnmarshalOptions"), ") (dst ", protoPackage.Ident("Message"), ", err error) {")
		g.P("	if src.GetTypeUrl() == \"\" {")
		g.P("		return nil, ", protoimplPackage.Ident("X"), ".NewError(\"invalid empty type URL\")")
		g.P("	}")
		g.P("	if opts.Resolver == nil {")
		g.P("		opts.Resolver = ", protoregistryPackage.Ident("GlobalTypes"))
		g.P("	}")
		g.P("	r, ok := opts.Resolver.(", protoregistryPackage.Ident("MessageTypeResolver"), ")")
		g.P("	if !ok {")
		g.P("		return nil, ", protoregistryPackage.Ident("NotFound"))
		g.P("	}")
		g.P("	mt, err := r.FindMessageByURL(src.GetTypeUrl())")
		g.P("	if err != nil {")
		g.P("		if err == ", protoregistryPackage.Ident("NotFound"), " {")
		g.P("			return nil, err")
		g.P("		}")
		g.P("		return nil, ", protoimplPackage.Ident("X"), ".NewError(\"could not resolve %q: %v\", src.GetTypeUrl(), err)")
		g.P("	}")
		g.P("	dst = mt.New().Interface()")
		g.P("	return dst, opts.Unmarshal(src.GetValue(), dst)")
		g.P("}")
		g.P()

		g.P("// MessageIs reports whether the underlying message is of the same type as m.")
		g.P("func (x *Any) MessageIs(m ", protoPackage.Ident("Message"), ") bool {")
		g.P("	if m == nil {")
		g.P("		return false")
		g.P("	}")
		g.P("	url := x.GetTypeUrl()")
		g.P("	name := string(m.ProtoReflect().Descriptor().FullName())")
		g.P("	if !", stringsPackage.Ident("HasSuffix"), "(url, name) {")
		g.P("		return false")
		g.P("	}")
		g.P("	return len(url) == len(name) || url[len(url)-len(name)-1] == '/'")
		g.P("}")
		g.P()

		g.P("// MessageName reports the full name of the underlying message,")
		g.P("// returning an empty string if invalid.")
		g.P("func (x *Any) MessageName() ", protoreflectPackage.Ident("FullName"), " {")
		g.P("	url := x.GetTypeUrl()")
		g.P("	name := ", protoreflectPackage.Ident("FullName"), "(url)")
		g.P("	if i := ", stringsPackage.Ident("LastIndexByte"), "(url, '/'); i >= 0 {")
		g.P("		name = name[i+len(\"/\"):]")
		g.P("	}")
		g.P("	if !name.IsValid() {")
		g.P("		return \"\"")
		g.P("	}")
		g.P("	return name")
		g.P("}")
		g.P()

		g.P("// MarshalFrom marshals m into x as the underlying message.")
		g.P("func (x *Any) MarshalFrom(m ", protoPackage.Ident("Message"), ") error {")
		g.P("	return MarshalFrom(x, m, ", protoPackage.Ident("MarshalOptions"), "{})")
		g.P("}")
		g.P()

		g.P("// UnmarshalTo unmarshals the contents of the underlying message of x into m.")
		g.P("// It resets m before performing the unmarshal operation.")
		g.P("// It reports an error if m is not of the right message type.")
		g.P("func (x *Any) UnmarshalTo(m ", protoPackage.Ident("Message"), ") error {")
		g.P("	return UnmarshalTo(x, m, ", protoPackage.Ident("UnmarshalOptions"), "{})")
		g.P("}")
		g.P()

		g.P("// UnmarshalNew unmarshals the contents of the underlying message of x into")
		g.P("// a newly allocated message of the specified type.")
		g.P("// It reports an error if the underlying message type could not be resolved.")
		g.P("func (x *Any) UnmarshalNew() (", protoPackage.Ident("Message"), ", error) {")
		g.P("	return UnmarshalNew(x, ", protoPackage.Ident("UnmarshalOptions"), "{})")
		g.P("}")
		g.P()

	case genid.Timestamp_message_fullname:
		g.P("// Now constructs a new Timestamp from the current time.")
		g.P("func Now() *Timestamp {")
		g.P("	return New(", timePackage.Ident("Now"), "())")
		g.P("}")
		g.P()

		g.P("// New constructs a new Timestamp from the provided time.Time.")
		g.P("func New(t ", timePackage.Ident("Time"), ") *Timestamp {")
		g.P("	return &Timestamp{Seconds: int64(t.Unix()), Nanos: int32(t.Nanosecond())}")
		g.P("}")
		g.P()

		g.P("// AsTime converts x to a time.Time.")
		g.P("func (x *Timestamp) AsTime() ", timePackage.Ident("Time"), " {")
		g.P("	return ", timePackage.Ident("Unix"), "(int64(x.GetSeconds()), int64(x.GetNanos())).UTC()")
		g.P("}")
		g.P()

		g.P("// IsValid reports whether the timestamp is valid.")
		g.P("// It is equivalent to CheckValid == nil.")
		g.P("func (x *Timestamp) IsValid() bool {")
		g.P("	return x.check() == 0")
		g.P("}")
		g.P()

		g.P("// CheckValid returns an error if the timestamp is invalid.")
		g.P("// In particular, it checks whether the value represents a date that is")
		g.P("// in the range of 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive.")
		g.P("// An error is reported for a nil Timestamp.")
		g.P("func (x *Timestamp) CheckValid() error {")
		g.P("	switch x.check() {")
		g.P("	case invalidNil:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"invalid nil Timestamp\")")
		g.P("	case invalidUnderflow:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"timestamp (%v) before 0001-01-01\", x)")
		g.P("	case invalidOverflow:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"timestamp (%v) after 9999-12-31\", x)")
		g.P("	case invalidNanos:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"timestamp (%v) has out-of-range nanos\", x)")
		g.P("	default:")
		g.P("		return nil")
		g.P("	}")
		g.P("}")
		g.P()

		g.P("const (")
		g.P("	_ = iota")
		g.P("	invalidNil")
		g.P("	invalidUnderflow")
		g.P("	invalidOverflow")
		g.P("	invalidNanos")
		g.P(")")
		g.P()

		g.P("func (x *Timestamp) check() uint {")
		g.P("	const minTimestamp = -62135596800  // Seconds between 1970-01-01T00:00:00Z and 0001-01-01T00:00:00Z, inclusive")
		g.P("	const maxTimestamp = +253402300799 // Seconds between 1970-01-01T00:00:00Z and 9999-12-31T23:59:59Z, inclusive")
		g.P("	secs := x.GetSeconds()")
		g.P("	nanos := x.GetNanos()")
		g.P("	switch {")
		g.P("	case x == nil:")
		g.P("		return invalidNil")
		g.P("	case secs < minTimestamp:")
		g.P("		return invalidUnderflow")
		g.P("	case secs > maxTimestamp:")
		g.P("		return invalidOverflow")
		g.P("	case nanos < 0 || nanos >= 1e9:")
		g.P("		return invalidNanos")
		g.P("	default:")
		g.P("		return 0")
		g.P("	}")
		g.P("}")
		g.P()

	case genid.Duration_message_fullname:
		g.P("// New constructs a new Duration from the provided time.Duration.")
		g.P("func New(d ", timePackage.Ident("Duration"), ") *Duration {")
		g.P("	nanos := d.Nanoseconds()")
		g.P("	secs := nanos / 1e9")
		g.P("	nanos -= secs * 1e9")
		g.P("	return &Duration{Seconds: int64(secs), Nanos: int32(nanos)}")
		g.P("}")
		g.P()

		g.P("// AsDuration converts x to a time.Duration,")
		g.P("// returning the closest duration value in the event of overflow.")
		g.P("func (x *Duration) AsDuration() ", timePackage.Ident("Duration"), " {")
		g.P("	secs := x.GetSeconds()")
		g.P("	nanos := x.GetNanos()")
		g.P("	d := ", timePackage.Ident("Duration"), "(secs) * ", timePackage.Ident("Second"))
		g.P("	overflow := d/", timePackage.Ident("Second"), " != ", timePackage.Ident("Duration"), "(secs)")
		g.P("	d += ", timePackage.Ident("Duration"), "(nanos) * ", timePackage.Ident("Nanosecond"))
		g.P("	overflow = overflow || (secs < 0 && nanos < 0 && d > 0)")
		g.P("	overflow = overflow || (secs > 0 && nanos > 0 && d < 0)")
		g.P("	if overflow {")
		g.P("		switch {")
		g.P("		case secs < 0:")
		g.P("			return ", timePackage.Ident("Duration"), "(", mathPackage.Ident("MinInt64"), ")")
		g.P("		case secs > 0:")
		g.P("			return ", timePackage.Ident("Duration"), "(", mathPackage.Ident("MaxInt64"), ")")
		g.P("		}")
		g.P("	}")
		g.P("	return d")
		g.P("}")
		g.P()

		g.P("// IsValid reports whether the duration is valid.")
		g.P("// It is equivalent to CheckValid == nil.")
		g.P("func (x *Duration) IsValid() bool {")
		g.P("	return x.check() == 0")
		g.P("}")
		g.P()

		g.P("// CheckValid returns an error if the duration is invalid.")
		g.P("// In particular, it checks whether the value is within the range of")
		g.P("// -10000 years to +10000 years inclusive.")
		g.P("// An error is reported for a nil Duration.")
		g.P("func (x *Duration) CheckValid() error {")
		g.P("	switch x.check() {")
		g.P("	case invalidNil:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"invalid nil Duration\")")
		g.P("	case invalidUnderflow:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"duration (%v) exceeds -10000 years\", x)")
		g.P("	case invalidOverflow:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"duration (%v) exceeds +10000 years\", x)")
		g.P("	case invalidNanosRange:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"duration (%v) has out-of-range nanos\", x)")
		g.P("	case invalidNanosSign:")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"duration (%v) has seconds and nanos with different signs\", x)")
		g.P("	default:")
		g.P("		return nil")
		g.P("	}")
		g.P("}")
		g.P()

		g.P("const (")
		g.P("	_ = iota")
		g.P("	invalidNil")
		g.P("	invalidUnderflow")
		g.P("	invalidOverflow")
		g.P("	invalidNanosRange")
		g.P("	invalidNanosSign")
		g.P(")")
		g.P()

		g.P("func (x *Duration) check() uint {")
		g.P("	const absDuration = 315576000000 // 10000yr * 365.25day/yr * 24hr/day * 60min/hr * 60sec/min")
		g.P("	secs := x.GetSeconds()")
		g.P("	nanos := x.GetNanos()")
		g.P("	switch {")
		g.P("	case x == nil:")
		g.P("		return invalidNil")
		g.P("	case secs < -absDuration:")
		g.P("		return invalidUnderflow")
		g.P("	case secs > +absDuration:")
		g.P("		return invalidOverflow")
		g.P("	case nanos <= -1e9 || nanos >= +1e9:")
		g.P("		return invalidNanosRange")
		g.P("	case (secs > 0 && nanos < 0) || (secs < 0 && nanos > 0):")
		g.P("		return invalidNanosSign")
		g.P("	default:")
		g.P("		return 0")
		g.P("	}")
		g.P("}")
		g.P()

	case genid.Struct_message_fullname:
		g.P("// NewStruct constructs a Struct from a general-purpose Go map.")
		g.P("// The map keys must be valid UTF-8.")
		g.P("// The map values are converted using NewValue.")
		g.P("func NewStruct(v map[string]interface{}) (*Struct, error) {")
		g.P("	x := &Struct{Fields: make(map[string]*Value, len(v))}")
		g.P("	for k, v := range v {")
		g.P("		if !", utf8Package.Ident("ValidString"), "(k) {")
		g.P("			return nil, ", protoimplPackage.Ident("X"), ".NewError(\"invalid UTF-8 in string: %q\", k)")
		g.P("		}")
		g.P("		var err error")
		g.P("		x.Fields[k], err = NewValue(v)")
		g.P("		if err != nil {")
		g.P("			return nil, err")
		g.P("		}")
		g.P("	}")
		g.P("	return x, nil")
		g.P("}")
		g.P()

		g.P("// AsMap converts x to a general-purpose Go map.")
		g.P("// The map values are converted by calling Value.AsInterface.")
		g.P("func (x *Struct) AsMap() map[string]interface{} {")
		g.P("	f := x.GetFields()")
		g.P("	vs := make(map[string]interface{}, len(f))")
		g.P("	for k, v := range f {")
		g.P("		vs[k] = v.AsInterface()")
		g.P("	}")
		g.P("	return vs")
		g.P("}")
		g.P()

		g.P("func (x *Struct) MarshalJSON() ([]byte, error) {")
		g.P("	return ", protojsonPackage.Ident("Marshal"), "(x)")
		g.P("}")
		g.P()

		g.P("func (x *Struct) UnmarshalJSON(b []byte) error {")
		g.P("	return ", protojsonPackage.Ident("Unmarshal"), "(b, x)")
		g.P("}")
		g.P()

	case genid.ListValue_message_fullname:
		g.P("// NewList constructs a ListValue from a general-purpose Go slice.")
		g.P("// The slice elements are converted using NewValue.")
		g.P("func NewList(v []interface{}) (*ListValue, error) {")
		g.P("	x := &ListValue{Values: make([]*Value, len(v))}")
		g.P("	for i, v := range v {")
		g.P("		var err error")
		g.P("		x.Values[i], err = NewValue(v)")
		g.P("		if err != nil {")
		g.P("			return nil, err")
		g.P("		}")
		g.P("	}")
		g.P("	return x, nil")
		g.P("}")
		g.P()

		g.P("// AsSlice converts x to a general-purpose Go slice.")
		g.P("// The slice elements are converted by calling Value.AsInterface.")
		g.P("func (x *ListValue) AsSlice() []interface{} {")
		g.P("	vals := x.GetValues()")
		g.P("	vs := make([]interface{}, len(vals))")
		g.P("	for i, v := range vals {")
		g.P("		vs[i] = v.AsInterface()")
		g.P("	}")
		g.P("	return vs")
		g.P("}")
		g.P()

		g.P("func (x *ListValue) MarshalJSON() ([]byte, error) {")
		g.P("	return ", protojsonPackage.Ident("Marshal"), "(x)")
		g.P("}")
		g.P()

		g.P("func (x *ListValue) UnmarshalJSON(b []byte) error {")
		g.P("	return ", protojsonPackage.Ident("Unmarshal"), "(b, x)")
		g.P("}")
		g.P()

	case genid.Value_message_fullname:
		g.P("// NewValue constructs a Value from a general-purpose Go interface.")
		g.P("//")
		g.P("//	╔════════════════════════╤════════════════════════════════════════════╗")
		g.P("//	║ Go type                │ Conversion                                 ║")
		g.P("//	╠════════════════════════╪════════════════════════════════════════════╣")
		g.P("//	║ nil                    │ stored as NullValue                        ║")
		g.P("//	║ bool                   │ stored as BoolValue                        ║")
		g.P("//	║ int, int32, int64      │ stored as NumberValue                      ║")
		g.P("//	║ uint, uint32, uint64   │ stored as NumberValue                      ║")
		g.P("//	║ float32, float64       │ stored as NumberValue                      ║")
		g.P("//	║ string                 │ stored as StringValue; must be valid UTF-8 ║")
		g.P("//	║ []byte                 │ stored as StringValue; base64-encoded      ║")
		g.P("//	║ map[string]interface{} │ stored as StructValue                      ║")
		g.P("//	║ []interface{}          │ stored as ListValue                        ║")
		g.P("//	╚════════════════════════╧════════════════════════════════════════════╝")
		g.P("//")
		g.P("// When converting an int64 or uint64 to a NumberValue, numeric precision loss")
		g.P("// is possible since they are stored as a float64.")
		g.P("func NewValue(v interface{}) (*Value, error) {")
		g.P("	switch v := v.(type) {")
		g.P("	case nil:")
		g.P("		return NewNullValue(), nil")
		g.P("	case bool:")
		g.P("		return NewBoolValue(v), nil")
		g.P("	case int:")
		g.P("		return NewNumberValue(float64(v)), nil")
		g.P("	case int32:")
		g.P("		return NewNumberValue(float64(v)), nil")
		g.P("	case int64:")
		g.P("		return NewNumberValue(float64(v)), nil")
		g.P("	case uint:")
		g.P("		return NewNumberValue(float64(v)), nil")
		g.P("	case uint32:")
		g.P("		return NewNumberValue(float64(v)), nil")
		g.P("	case uint64:")
		g.P("		return NewNumberValue(float64(v)), nil")
		g.P("	case float32:")
		g.P("		return NewNumberValue(float64(v)), nil")
		g.P("	case float64:")
		g.P("		return NewNumberValue(float64(v)), nil")
		g.P("	case string:")
		g.P("		if !", utf8Package.Ident("ValidString"), "(v) {")
		g.P("			return nil, ", protoimplPackage.Ident("X"), ".NewError(\"invalid UTF-8 in string: %q\", v)")
		g.P("		}")
		g.P("		return NewStringValue(v), nil")
		g.P("	case []byte:")
		g.P("		s := ", base64Package.Ident("StdEncoding"), ".EncodeToString(v)")
		g.P("		return NewStringValue(s), nil")
		g.P("	case map[string]interface{}:")
		g.P("		v2, err := NewStruct(v)")
		g.P("		if err != nil {")
		g.P("			return nil, err")
		g.P("		}")
		g.P("		return NewStructValue(v2), nil")
		g.P("	case []interface{}:")
		g.P("		v2, err := NewList(v)")
		g.P("		if err != nil {")
		g.P("			return nil, err")
		g.P("		}")
		g.P("		return NewListValue(v2), nil")
		g.P("	default:")
		g.P("		return nil, ", protoimplPackage.Ident("X"), ".NewError(\"invalid type: %T\", v)")
		g.P("	}")
		g.P("}")
		g.P()

		g.P("// NewNullValue constructs a new null Value.")
		g.P("func NewNullValue() *Value {")
		g.P("	return &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}}")
		g.P("}")
		g.P()

		g.P("// NewBoolValue constructs a new boolean Value.")
		g.P("func NewBoolValue(v bool) *Value {")
		g.P("	return &Value{Kind: &Value_BoolValue{BoolValue: v}}")
		g.P("}")
		g.P()

		g.P("// NewNumberValue constructs a new number Value.")
		g.P("func NewNumberValue(v float64) *Value {")
		g.P("	return &Value{Kind: &Value_NumberValue{NumberValue: v}}")
		g.P("}")
		g.P()

		g.P("// NewStringValue constructs a new string Value.")
		g.P("func NewStringValue(v string) *Value {")
		g.P("	return &Value{Kind: &Value_StringValue{StringValue: v}}")
		g.P("}")
		g.P()

		g.P("// NewStructValue constructs a new struct Value.")
		g.P("func NewStructValue(v *Struct) *Value {")
		g.P("	return &Value{Kind: &Value_StructValue{StructValue: v}}")
		g.P("}")
		g.P()

		g.P("// NewListValue constructs a new list Value.")
		g.P("func NewListValue(v *ListValue) *Value {")
		g.P("	return &Value{Kind: &Value_ListValue{ListValue: v}}")
		g.P("}")
		g.P()

		g.P("// AsInterface converts x to a general-purpose Go interface.")
		g.P("//")
		g.P("// Calling Value.MarshalJSON and \"encoding/json\".Marshal on this output produce")
		g.P("// semantically equivalent JSON (assuming no errors occur).")
		g.P("//")
		g.P("// Floating-point values (i.e., \"NaN\", \"Infinity\", and \"-Infinity\") are")
		g.P("// converted as strings to remain compatible with MarshalJSON.")
		g.P("func (x *Value) AsInterface() interface{} {")
		g.P("	switch v := x.GetKind().(type) {")
		g.P("	case *Value_NumberValue:")
		g.P("		if v != nil {")
		g.P("			switch {")
		g.P("			case ", mathPackage.Ident("IsNaN"), "(v.NumberValue):")
		g.P("				return \"NaN\"")
		g.P("			case ", mathPackage.Ident("IsInf"), "(v.NumberValue, +1):")
		g.P("				return \"Infinity\"")
		g.P("			case ", mathPackage.Ident("IsInf"), "(v.NumberValue, -1):")
		g.P("				return \"-Infinity\"")
		g.P("			default:")
		g.P("				return v.NumberValue")
		g.P("			}")
		g.P("		}")
		g.P("	case *Value_StringValue:")
		g.P("		if v != nil {")
		g.P("			return v.StringValue")
		g.P("		}")
		g.P("	case *Value_BoolValue:")
		g.P("		if v != nil {")
		g.P("			return v.BoolValue")
		g.P("		}")
		g.P("	case *Value_StructValue:")
		g.P("		if v != nil {")
		g.P("			return v.StructValue.AsMap()")
		g.P("		}")
		g.P("	case *Value_ListValue:")
		g.P("		if v != nil {")
		g.P("			return v.ListValue.AsSlice()")
		g.P("		}")
		g.P("	}")
		g.P("	return nil")
		g.P("}")
		g.P()

		g.P("func (x *Value) MarshalJSON() ([]byte, error) {")
		g.P("	return ", protojsonPackage.Ident("Marshal"), "(x)")
		g.P("}")
		g.P()

		g.P("func (x *Value) UnmarshalJSON(b []byte) error {")
		g.P("	return ", protojsonPackage.Ident("Unmarshal"), "(b, x)")
		g.P("}")
		g.P()

	case genid.FieldMask_message_fullname:
		g.P("// New constructs a field mask from a list of paths and verifies that")
		g.P("// each one is valid according to the specified message type.")
		g.P("func New(m ", protoPackage.Ident("Message"), ", paths ...string) (*FieldMask, error) {")
		g.P("	x := new(FieldMask)")
		g.P("	return x, x.Append(m, paths...)")
		g.P("}")
		g.P()

		g.P("// Union returns the union of all the paths in the input field masks.")
		g.P("func Union(mx *FieldMask, my *FieldMask, ms ...*FieldMask) *FieldMask {")
		g.P("	var out []string")
		g.P("	out = append(out, mx.GetPaths()...)")
		g.P("	out = append(out, my.GetPaths()...)")
		g.P("	for _, m := range ms {")
		g.P("		out = append(out, m.GetPaths()...)")
		g.P("	}")
		g.P("	return &FieldMask{Paths: normalizePaths(out)}")
		g.P("}")
		g.P()

		g.P("// Intersect returns the intersection of all the paths in the input field masks.")
		g.P("func Intersect(mx *FieldMask, my *FieldMask, ms ...*FieldMask) *FieldMask {")
		g.P("	var ss1, ss2 []string // reused buffers for performance")
		g.P("	intersect := func(out, in []string) []string {")
		g.P("		ss1 = normalizePaths(append(ss1[:0], in...))")
		g.P("		ss2 = normalizePaths(append(ss2[:0], out...))")
		g.P("		out = out[:0]")
		g.P("		for i1, i2 := 0, 0; i1 < len(ss1) && i2 < len(ss2); {")
		g.P("			switch s1, s2 := ss1[i1], ss2[i2]; {")
		g.P("			case hasPathPrefix(s1, s2):")
		g.P("				out = append(out, s1)")
		g.P("				i1++")
		g.P("			case hasPathPrefix(s2, s1):")
		g.P("				out = append(out, s2)")
		g.P("				i2++")
		g.P("			case lessPath(s1, s2):")
		g.P("				i1++")
		g.P("			case lessPath(s2, s1):")
		g.P("				i2++")
		g.P("			}")
		g.P("		}")
		g.P("		return out")
		g.P("	}")
		g.P()
		g.P("	out := Union(mx, my, ms...).GetPaths()")
		g.P("	out = intersect(out, mx.GetPaths())")
		g.P("	out = intersect(out, my.GetPaths())")
		g.P("	for _, m := range ms {")
		g.P("		out = intersect(out, m.GetPaths())")
		g.P("	}")
		g.P("	return &FieldMask{Paths: normalizePaths(out)}")
		g.P("}")
		g.P()

		g.P("// IsValid reports whether all the paths are syntactically valid and")
		g.P("// refer to known fields in the specified message type.")
		g.P("// It reports false for a nil FieldMask.")
		g.P("func (x *FieldMask) IsValid(m ", protoPackage.Ident("Message"), ") bool {")
		g.P("	paths := x.GetPaths()")
		g.P("	return x != nil && numValidPaths(m, paths) == len(paths)")
		g.P("}")
		g.P()

		g.P("// Append appends a list of paths to the mask and verifies that each one")
		g.P("// is valid according to the specified message type.")
		g.P("// An invalid path is not appended and breaks insertion of subsequent paths.")
		g.P("func (x *FieldMask) Append(m ", protoPackage.Ident("Message"), ", paths ...string) error {")
		g.P("	numValid := numValidPaths(m, paths)")
		g.P("	x.Paths = append(x.Paths, paths[:numValid]...)")
		g.P("	paths = paths[numValid:]")
		g.P("	if len(paths) > 0 {")
		g.P("		name := m.ProtoReflect().Descriptor().FullName()")
		g.P("		return ", protoimplPackage.Ident("X"), ".NewError(\"invalid path %q for message %q\", paths[0], name)")
		g.P("	}")
		g.P("	return nil")
		g.P("}")
		g.P()

		g.P("func numValidPaths(m ", protoPackage.Ident("Message"), ", paths []string) int {")
		g.P("	md0 := m.ProtoReflect().Descriptor()")
		g.P("	for i, path := range paths {")
		g.P("		md := md0")
		g.P("		if !rangeFields(path, func(field string) bool {")
		g.P("			// Search the field within the message.")
		g.P("			if md == nil {")
		g.P("				return false // not within a message")
		g.P("			}")
		g.P("			fd := md.Fields().ByName(", protoreflectPackage.Ident("Name"), "(field))")
		g.P("			// The real field name of a group is the message name.")
		g.P("			if fd == nil {")
		g.P("				gd := md.Fields().ByName(", protoreflectPackage.Ident("Name"), "(", stringsPackage.Ident("ToLower"), "(field)))")
		g.P("				if gd != nil && gd.Kind() == ", protoreflectPackage.Ident("GroupKind"), " && string(gd.Message().Name()) == field {")
		g.P("					fd = gd")
		g.P("				}")
		g.P("			} else if fd.Kind() == ", protoreflectPackage.Ident("GroupKind"), " && string(fd.Message().Name()) != field {")
		g.P("				fd = nil")
		g.P("			}")
		g.P("			if fd == nil {")
		g.P("				return false // message has does not have this field")
		g.P("			}")
		g.P()
		g.P("			// Identify the next message to search within.")
		g.P("			md = fd.Message() // may be nil")
		g.P()
		g.P("			// Repeated fields are only allowed at the last position.")
		g.P("			if fd.IsList() || fd.IsMap() {")
		g.P("				md = nil")
		g.P("			}")
		g.P()
		g.P("			return true")
		g.P("		}) {")
		g.P("			return i")
		g.P("		}")
		g.P("	}")
		g.P("	return len(paths)")
		g.P("}")
		g.P()

		g.P("// Normalize converts the mask to its canonical form where all paths are sorted")
		g.P("// and redundant paths are removed.")
		g.P("func (x *FieldMask) Normalize() {")
		g.P("	x.Paths = normalizePaths(x.Paths)")
		g.P("}")
		g.P()
		g.P("func normalizePaths(paths []string) []string {")
		g.P("	", sortPackage.Ident("Slice"), "(paths, func(i, j int) bool {")
		g.P("		return lessPath(paths[i], paths[j])")
		g.P("	})")
		g.P()
		g.P("	// Elide any path that is a prefix match on the previous.")
		g.P("	out := paths[:0]")
		g.P("	for _, path := range paths {")
		g.P("		if len(out) > 0 && hasPathPrefix(path, out[len(out)-1]) {")
		g.P("			continue")
		g.P("		}")
		g.P("		out = append(out, path)")
		g.P("	}")
		g.P("	return out")
		g.P("}")
		g.P()

		g.P("// hasPathPrefix is like strings.HasPrefix, but further checks for either")
		g.P("// an exact matche or that the prefix is delimited by a dot.")
		g.P("func hasPathPrefix(path, prefix string) bool {")
		g.P("	return ", stringsPackage.Ident("HasPrefix"), "(path, prefix) && (len(path) == len(prefix) || path[len(prefix)] == '.')")
		g.P("}")
		g.P()

		g.P("// lessPath is a lexicographical comparison where dot is specially treated")
		g.P("// as the smallest symbol.")
		g.P("func lessPath(x, y string) bool {")
		g.P("	for i := 0; i < len(x) && i < len(y); i++ {")
		g.P("		if x[i] != y[i] {")
		g.P("			return (x[i] - '.') < (y[i] - '.')")
		g.P("		}")
		g.P("	}")
		g.P("	return len(x) < len(y)")
		g.P("}")
		g.P()

		g.P("// rangeFields is like strings.Split(path, \".\"), but avoids allocations by")
		g.P("// iterating over each field in place and calling a iterator function.")
		g.P("func rangeFields(path string, f func(field string) bool) bool {")
		g.P("	for {")
		g.P("		var field string")
		g.P("		if i := ", stringsPackage.Ident("IndexByte"), "(path, '.'); i >= 0 {")
		g.P("			field, path = path[:i], path[i:]")
		g.P("		} else {")
		g.P("			field, path = path, \"\"")
		g.P("		}")
		g.P()
		g.P("		if !f(field) {")
		g.P("			return false")
		g.P("		}")
		g.P()
		g.P("		if len(path) == 0 {")
		g.P("			return true")
		g.P("		}")
		g.P("		path = ", stringsPackage.Ident("TrimPrefix"), "(path, \".\")")
		g.P("	}")
		g.P("}")
		g.P()

	case genid.BoolValue_message_fullname,
		genid.Int32Value_message_fullname,
		genid.Int64Value_message_fullname,
		genid.UInt32Value_message_fullname,
		genid.UInt64Value_message_fullname,
		genid.FloatValue_message_fullname,
		genid.DoubleValue_message_fullname,
		genid.StringValue_message_fullname,
		genid.BytesValue_message_fullname:
		funcName := strings.TrimSuffix(m.GoIdent.GoName, "Value")
		typeName := strings.ToLower(funcName)
		switch typeName {
		case "float":
			typeName = "float32"
		case "double":
			typeName = "float64"
		case "bytes":
			typeName = "[]byte"
		}

		g.P("// ", funcName, " stores v in a new ", m.GoIdent, " and returns a pointer to it.")
		g.P("func ", funcName, "(v ", typeName, ") *", m.GoIdent, " {")
		g.P("	return &", m.GoIdent, "{Value: v}")
		g.P("}")
		g.P()
	}
}
