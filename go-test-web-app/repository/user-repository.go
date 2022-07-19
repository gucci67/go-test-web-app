package repository

import "go-test-web-app/entity"

type UserRepository struct{}

func (e *UserRepository) Select(userId int) (entity.User, error) {
	db := GetDB()
	var user entity.User

	query := `
		SELECT
			*
		FROM
			users
		WHERE
			user_id = ?`

	db.Get(user, query, userId)

	if err := db.Get(&user, query, userId); err != nil {
		return user, err
	}

	return user, nil
}

func (e *UserRepository) SelectAll() ([]entity.User, error) {
	db := GetDB()
	var users []entity.User

	query := `
		SELECT
			*
		FROM
			users`

	if err := db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (e *UserRepository) Insert(user entity.User) error {
	db := GetDB()

	query := `
		INSERT INTO
			users (name, address)
		VALUES
			(:name, :address)`

	if _, err := db.NamedExec(query, user); err != nil {
		return err
	}

	return nil
}

func (e *UserRepository) Update(user entity.User) error {
	db := GetDB()

	query := `
		UPDATE
			users
		SET
			name = :name
			, address = :address`

	if _, err := db.NamedExec(query, user); err != nil {
		return err
	}

	return nil
}
