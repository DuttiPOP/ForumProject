package repository

import (
	"ForumProject/model/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (repository *UserRepository) Create(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (email, username, password) VALUES ($1, $2, $3)`, userTable)
	row := repository.db.QueryRow(query, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repository *UserRepository) Get(id int) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, userTable)
	err := repository.db.Get(&user, query, id)
	return user, err
}

func (repository *UserRepository) Delete(id int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, userTable)
	_, err := repository.db.Exec(query, id)
	return err
}
