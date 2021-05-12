package serializer_test

import (
	"testing"

	pb "github.com/pkbhowmick/just-grpc/proto"

	"github.com/golang/protobuf/proto"

	"github.com/pkbhowmick/just-grpc/sample"
	"github.com/pkbhowmick/just-grpc/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	jsonFile := "../example/laptop.json"
	binFile := "../example/laptop.bin"
	laptop := sample.NewLaptop()
	err := serializer.WriteProtobufToJsonFile(laptop, jsonFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToBinaryFile(laptop, binFile)
	require.NoError(t, err)

	laptopCopy := &pb.Laptop{}

	err = serializer.ReadProtobufToBinaryFile(binFile, laptopCopy)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop, laptopCopy))
}
