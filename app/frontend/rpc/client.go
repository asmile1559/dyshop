package rpc

import (
	pbbackenduser "github.com/dyshop/pb/backend/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	UserClient pbbackenduser.UserServiceClient
)

func a() {
	cc, _ := grpc.NewClient(":8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	pbbackenduser.NewUserServiceClient(cc)
}
