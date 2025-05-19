package repository

import (
	"golang-tutorial/contract"
	"golang-tutorial/entity"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func implUserRepository(db *gorm.DB) contract.UserRepository {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) GetUser(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepo) CreateUser(user *entity.User) error {
	return r.db.Table("users").Create(user).Error
}

func (r *UserRepo) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepo) CheckEmail(email string) (bool, error) {
	var exists bool
	err := r.db.Raw("SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists).Error
	return exists, err
}
