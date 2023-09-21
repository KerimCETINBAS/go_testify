package service

import (
	"testing"

	"github.com/kerimcetinbas/testify_tut/types"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
}

func (s *UserServiceTestSuite) TestGET() {
	us := UserService()
	user := us.GET(1)

	s.Implements((*IUserService)(nil), us)
	s.Equal(types.User{}, user)

}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
