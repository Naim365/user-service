package handler

import (
	"user-service/service"
)

type UserHandler struct {
	// pb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// func (h *UserHandler) ReadUser(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
// 	user, err := h.service.GetUser(req.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ReadUserResponse{
// 		Id:    user.ID,
// 		Name:  user.Name,
// 		Email: user.Email,
// 	}, nil
// }


