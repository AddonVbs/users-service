package transportgrpc

import (
	"context"

	userpb "github.com/AddonVbs/project-protos/proto/user"
	user "github.com/AddonVbs/users-service/internal/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc *user.UserService1
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.UserService1) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

}
func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {

}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {

}
func (h *Handler) ListUsers(ctx context.Context, req *emptypb.Empty) (*userpb.ListUsersResponse, error) {

}
func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {

}
