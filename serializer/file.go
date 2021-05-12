package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func WriteProtobufToJsonFile(message proto.Message, filename string) error {
	data, err := ProtobufToJson(message)
	if err != nil {
		return fmt.Errorf("can't convert protobuf to json files: %v", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("can't write json message to file: %v", err)
	}
	return nil
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("can't marshal protobuf message: %v", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("can't write binary data to file")
	}
	return nil
}

func ReadProtobufToBinaryFile(filename string, message proto.Message) error {
	date, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("can't convert binary file to protobuf")
	}

	err = proto.Unmarshal(date, message)
	if err != nil {
		return fmt.Errorf("can't unmarshal data to protobuf")
	}
	return nil
}

func ProtobufToJson(message proto.Message) (string, error) {
	marshaller := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
		AnyResolver:  nil,
	}
	return marshaller.MarshalToString(message)
}
