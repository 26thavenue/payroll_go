package repository

import (
	"github.com/26thavenue/payroll_go/internal/db"
	"github.com/26thavenue/payroll_go/pkg"
	"gorm.io/gorm"
	"github.com/26thavenue/payroll_go/internal/utils"
)

type UserRepository interface {
	CreateUser(user *db.User) error
    GetUserByID(id uint) (*db.User, error)
    UpdateUser(user *db.User) error
    DeleteUser(id uint) error
	FindByEmail(email string) (*db.User, error)
    Authenticate(username, password string) (*db.User, error)
    ChangePassword(userID uint, oldPassword, newPassword string) error
	ListUsers(limit, offset int) ([]*db.User, error)
    SearchUsers(query string) ([]*db.User, error)
}


type GormUserRepository struct {
	db *gorm.DB
}


func (r *GormUserRepository) FindByEmail(email string) (*db.User, error) {
    var user db.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, pkg.ErrUserNotFound
        }
        return nil, err
    }
    return &user, nil
}

func (r *GormUserRepository) CreateUser(user *db.User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepository) GetUserByID(id uint) (*db.User, error) {
	var user db.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pkg.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) UpdateUser(user *db.User) error {
	return r.db.Save(user).Error
}

func (r *GormUserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&db.User{}, id).Error
}

func (r *GormUserRepository) Authenticate(username, password string) (*db.User, error) {
	var user db.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pkg.ErrInvalidCredentials
		}
		return nil, err
	}


	err = utils.ComparePasswords(user.Password, password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) ListUsers(limit, offset int) ([]*db.User, error) {
	var users []*db.User
	err := r.db.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *GormUserRepository) SearchUsers(query string) ([]*db.User, error) {
	var users []*db.User
	err := r.db.Where("fullname LIKE ? OR email LIKE ?", "%"+query+"%", "%"+query+"%").Find(&users).Error
	return users, err
}
