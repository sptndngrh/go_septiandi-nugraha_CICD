package repositories

import (
	"praktikum_23/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func NewMockUserRepo() *MockUserRepository {
	return &MockUserRepository{}
}

func (m *MockUserRepository) Create(user models.User) error {
	ret := m.Called(user)
	return ret.Error(0)
}

func (m *MockUserRepository) Find() ([]models.User, error) {
	ret := m.Called()
	return ret.Get(0).([]models.User), ret.Error(1)
}
