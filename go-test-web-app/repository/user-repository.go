package repository

import "go-test-web-app/entity"

type UserRepository struct{}

func (e *UserRepository) Select(userId int) (entity.User, error) {
	db := GetDB()
	var user entity.User

	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (e *UserRepository) SelectAll() ([]entity.User, error) {
	db := GetDB()
	var users []entity.User

	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (e *UserRepository) Insert(user entity.User) error {
	db := GetDB()

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (e *UserRepository) Update(user entity.User) error {
	db := GetDB()

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
