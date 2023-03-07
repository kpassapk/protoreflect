package main

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"tmp/protoreflect/registry"
)

func main() {
	// The 'protoregistry' package helps register types and recover a protoreflect.Descriptor.

	// It is enough for the types to be compiled into the binary for them to be available
	// through the proto registry.
	_ = (*registry.ComplexType)(nil)

	// Print all registered types.
	// name := protoreflect.FullName("tmp.protoreflect.protoregistry.example")
	protoregistry.GlobalTypes.RangeMessages(func(t protoreflect.MessageType) bool {
		fmt.Println(t.Descriptor().FullName())
		return true
	})

	name := protoreflect.FullName("tmp.protoreflect.registry.example.ComplexType")

	desc, err := protoregistry.GlobalTypes.FindMessageByName(name)
	if err != nil {
		panic(err)
	}
	fmt.Println(desc)
	msg := desc.New().Interface()

	err = protojson.Unmarshal(registry.InputJson, msg)
	if err != nil {
		panic(err)
	}

	mt := msg.(*registry.ComplexType)
	fmt.Println(mt.Id)
	fmt.Println(mt.Name)
	fmt.Println(mt.Birthday.AsTime().Year())
}
