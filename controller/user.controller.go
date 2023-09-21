package controller

import (
	"github.com/kerimcetinbas/testify_tut/service"
	"github.com/kerimcetinbas/testify_tut/types"
)

type IUserController interface {
	GET(int) types.User
}
type userController struct {
	userService service.IUserService
}

func UserController(userService service.IUserService) IUserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GET(id int) types.User {
	return c.userService.GET(id)
}
