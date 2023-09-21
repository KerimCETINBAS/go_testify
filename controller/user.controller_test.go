package controller

import (
	"testing"

	"github.com/kerimcetinbas/testify_tut/service/mock"
	"github.com/kerimcetinbas/testify_tut/types"
	"github.com/stretchr/testify/suite"
)

type UserControllerTestSuite struct {
	suite.Suite
}

func (s *UserControllerTestSuite) TestGET() {
	mockUserService := mock.MockUserService{}
	uc := UserController(&mockUserService)

	s.Implements((*IUserController)(nil), uc)

	mockUserService.On("GET", 1).Return(types.User{})
	user := uc.GET(1)
	s.Equal(types.User{}, user)

}

func TestUserServiceTestSuite(t *testing.T) {

	suite.Run(t, new(UserControllerTestSuite))
}
