package grpc

import (
	"net"

	userpb "github.com/IAMBURNINGMAN/proto/proto/user"
	"github.com/IAMBURNINGMAN/user-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc user.UserService) error {
	// 1. listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	// 2. grpc server
	grpcSrv := grpc.NewServer()

	// 3. register handler
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))

	// 4. serve
	return grpcSrv.Serve(lis)
}
