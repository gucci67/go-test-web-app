package service

import (
	"go-test-web-app/entity"
	"go-test-web-app/repository"
)

var userRepository = repository.UserRepository{}

type UserService struct{}

// ユーザを取得します。
func (s *UserService) GetUser(userId int) (entity.User, error) {
	return userRepository.Select(userId)
}

// ユーザ一覧を取得します。
func (s *UserService) GetUserList() ([]entity.User, error) {
	return userRepository.SelectAll()
}

// ユーザを登録します。
func (s *UserService) CreateUser(name, address string) error {
	user := entity.User{
		Name:    name,
		Address: address,
	}
	return userRepository.Insert(user)
}

// ユーザを更新します。
func (s *UserService) UpdateUser(userId int, name, address string) error {
	user := entity.User{
		UserId:  userId,
		Name:    name,
		Address: address,
	}
	return userRepository.Update(user)
}
