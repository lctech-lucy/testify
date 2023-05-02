package assert

import "google.golang.org/protobuf/reflect/protoreflect"

type ProtoMessage interface {
	Reset()
	ProtoReflect() protoreflect.Message
}
