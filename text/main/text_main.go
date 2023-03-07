package main

import (
	"google.golang.org/protobuf/encoding/prototext"
	"tmp/protoreflect/text"
)

func main() {
	// There is a 'prototext' package with a Marshal function, but it doesn't look to be
	// general purpose enough for us to use it to decode CSV files, for example.

	dest := text.ComplexType{}
	dest.Id = 3
	dest.Name = "Joe"

	str, err := prototext.Marshal(&dest)
	if err != nil {
		panic(err)
	}

	println(string(str))
}
