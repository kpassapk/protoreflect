package main

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"tmp/protoreflect/json"
)

func main() {
	dest := json.ComplexType{}
	src := json.InputJson

	// In JSON, the string values get converted to right type!
	protojson.Unmarshal(src, &dest)
	fmt.Println(dest.Id)
	fmt.Println(dest.Name)
	fmt.Println(dest.Birthday.AsTime().Year())
	//Output:
	//412
	//one
	//1972
}
