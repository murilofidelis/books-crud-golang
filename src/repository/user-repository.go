package repository

import (
	"api/src/models"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (userRepository User) Save(user models.User) (uint64, error) {

	lastId := 0

	err := userRepository.db.QueryRow(
		"INSERT INTO dev_book.user (name, nick, email, password) VALUES($1, $2, $3, $4) RETURNING id",
		user.Name, user.Nick, user.Email, user.Password).Scan(&lastId)

	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (userRepository User) FindByFilter(nameParam string) ([]models.User, error) {

	rows, err := userRepository.db.Query(`
	 SELECT id, name, nick, email, cretae_at FROM dev_book.user 
	 WHERE LOWER(name) LIKE '%' || $1 || '%' OR LOWER(nick) LIKE '%' || $2 || '%' `,
		nameParam, nameParam)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (userRepository User) FindById(userId uint64) (models.User, error) {
	rows, err := userRepository.db.Query(
		"SELECT id, name, nick, email, cretae_at FROM dev_book.user WHERE id = $1", userId)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
