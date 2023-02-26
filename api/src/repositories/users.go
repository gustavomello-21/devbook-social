package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gustavomello-21/devbook/api/src/models"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *user {
	return &user{db: db}
}

func (u user) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("INSERT INTO users (fullname, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	result, err := statement.Exec(user.Fullname, user.Nick, user.Email, user.Passord)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return uint64(lastId), nil
}

func (u user) Find(NameOrNick string) ([]models.User, error) {
	nameOrNick := fmt.Sprintf("%%%s%%", NameOrNick)

	rows, err := u.db.Query(
		"SELECT id, fullname, nick, email, createdAt FROM users WHERE fullname LIKE ? OR nick LIKE ?",
		nameOrNick,
		nameOrNick,
	)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Fullname, &user.Nick, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (u user) FindById(targetId interface{}) (models.User, error) {
	row, err := u.db.Query("SELECT id, fullname, nick, email, createdAt FROM users WHERE id = ?", targetId)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	for row.Next() {
		if err = row.Scan(&user.ID, &user.Fullname, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, nil
		}
	}

	return user, nil
}

func (u *user) Update(targetId int, user models.User) error {
	statement, err := u.db.Prepare("UPDATE users SET fullname = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(user.Fullname, user.Nick, user.Email, targetId)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (u *user) Delete(targetId int) error {
	statement, err := u.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(targetId)

	return err
}
