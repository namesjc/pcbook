package service_test

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/namesjc/pcbook/pb"
	"github.com/namesjc/pcbook/sample"
	"github.com/namesjc/pcbook/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopServer, serverAddress := startTestLaptopServer((t))
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	exceptedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	fmt.Println(req)

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, exceptedID, res.Id)

	// check that the laptop is saved to the store
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check that the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)
}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	// reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	// go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	// defer conn.Close()
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	require.Equal(t, laptop1, laptop2)
}
