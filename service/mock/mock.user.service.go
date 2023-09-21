package mock

import (
	"github.com/kerimcetinbas/testify_tut/types"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GET(id int) types.User {
	args := m.Called(id)
	return args.Get(0).(types.User)
}
