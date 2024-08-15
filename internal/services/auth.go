package services

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/internal/repository"
	"github.com/26thavenue/payroll_go/pkg"
	"github.com/26thavenue/payroll_go/internal/utils"
		
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Authenticate(username, password string) (*db.User, error) {
	return s.userRepo.Authenticate(username, password)
}

func (s *AuthService) Register(user *db.User) error {
	existingUser, _ := s.userRepo.FindByEmail(user.Email)
	if existingUser != nil {
		return pkg.ErrUsernameTaken
	}
	
	user.Password, _ = utils.HashPassword(user.Password)
	
	return s.userRepo.CreateUser(user)
}