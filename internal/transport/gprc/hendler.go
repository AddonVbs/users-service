package transportgrpc

import (
	"context"
	"net/http"

	userpb "github.com/AddonVbs/project-protos/proto/user"
	user "github.com/AddonVbs/users-service/internal/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc user.UserService1
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.UserService1) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	user := new(user.User)

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	us, err := h.svc.CreateUser(req.Email, req.Password)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	apiUser := userpb.User{
		Id:       uint32(us.Id),
		Email:    us.Email,
		Password: us.Password, // или уберите, если не нужно возвращать пароль
	}

	return &userpb.CreateUserResponse{User: &apiUser}, nil

}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id := req.Id

	err := h.svc.DeleteUser(int(id))
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {

	us, err := h.svc.GetUserForTasks(int(req.UserID))
	if err != nil {
		return nil, err
	}

	protoUsers := userpb.User{
		Id:
	}

	return &userpb.GetUserResponse{User: &protoUsers}, nil
}
func (h *Handler) ListUsers(ctx context.Context, req *emptypb.Empty) (*userpb.ListUsersResponse, error) {
	us, err := h.svc.GetAllUser()
	if err != nil {
		return nil, err
	}

	var protoUsers []*userpb.User
	for _, u := range us {
		protoUsers = append(protoUsers, &userpb.User{
			Id:       uint32(u.Id),
			UserId: 	
			Email:    u.Email,
			Password: u.Password,
		})
	}

	return &userpb.ListUsersResponse{Users: protoUsers}, nil
}
func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	d := user.User{Id: int(req.Id)}
	if req.Email != "" {
		d.Email = req.Email
	}
	if req.Password != "" {
		d.Password = req.Password
	}

	updated, err := h.svc.UpdataUser(d.Id, d)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "update user: %v", err)
	}

	result := userpb.User{
		Id:       uint32(updated.Id),
		Email:    updated.Email,
		Password: updated.Password,
	}
	// !!! здесь мапим домен -> proto
	return &userpb.UpdateUserResponse{User: &result}, nil
}
