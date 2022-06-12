package serializer_test

import (
	"testing"

	"github.com/namesjc/pcbook/pb"
	"github.com/namesjc/pcbook/sample"
	"github.com/namesjc/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	lattop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, lattop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, lattop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)

}
