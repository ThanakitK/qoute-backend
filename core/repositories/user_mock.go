package repositories

import (
	"backend/core/models"

	"github.com/stretchr/testify/mock"
)

type userRepoMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *userRepoMock {
	return &userRepoMock{}
}

func (m *userRepoMock) GetUser(email string) (result models.UserModel, err error) {
	args := m.Called(email)
	return args.Get(0).(models.UserModel), args.Error(1)
}

func (m *userRepoMock) CreateUser(user models.CreateUserModel) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *userRepoMock) UpdateUser(id string, user models.UpdateUserModel) (result models.UserModel, err error) {
	args := m.Called(id, user)
	return args.Get(0).(models.UserModel), args.Error(1)
}
