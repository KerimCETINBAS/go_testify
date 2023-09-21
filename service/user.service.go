package service

import "github.com/kerimcetinbas/testify_tut/types"

type IUserService interface {
	GET(int) types.User
}

type userService struct{}

func UserService() IUserService {
	return &userService{}
}

func (s *userService) GET(id int) types.User {

	return types.User{}
}
