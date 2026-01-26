package grpc

import (
	"context"

	userpb "github.com/IAMBURNINGMAN/proto/proto/user"
	"github.com/IAMBURNINGMAN/user-service/internal/user"
)

type Handler struct {
	svc user.UserService
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(
	ctx context.Context,
	req *userpb.CreateUserRequest,
) (*userpb.CreateUserResponse, error) {

	// 1. protobuf → domain
	u := user.UserStruct{
		Email: req.Email,
	}

	// 2. вызов бизнес-логики
	createdUser, err := h.svc.CreateUser(u)
	if err != nil {
		return nil, err
	}

	// 3. domain → protobuf
	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(createdUser.ID),
			Email: createdUser.Email,
		},
	}, nil
}
func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}
	pbUsers := make([]*userpb.User, 0, len(users))
	for _, u := range users {
		pbUsers = append(pbUsers, &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		})
	}
	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	u, err := h.svc.GetUserById(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    uint32(u.ID),
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	u := user.UserStruct{
		Email: req.Email,
	}
	updatedUser, err := h.svc.UpdateUser(uint(req.Id), u)
	if err != nil {
		return nil, err
	}
	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(updatedUser.ID),
			Email: updatedUser.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := h.svc.DeleteUser(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{Success: true}, nil
}
