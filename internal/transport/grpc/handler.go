package grpc

import (
	userpb "github.com/IAMBURNINGMAN/proto"
	"github.com/IAMBURNINGMAN/user-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}
