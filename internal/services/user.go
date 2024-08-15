package services

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(id uint) (*db.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *db.User) error {
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}

func (s *UserService) FindByEmail(email string) (*db.User, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *UserService) ListUsers(limit, offset int) ([]*db.User, error) {
	return s.userRepo.ListUsers(limit, offset)
}

func (s *UserService) SearchUsers(query string) ([]*db.User, error) {
	return s.userRepo.SearchUsers(query)
}

