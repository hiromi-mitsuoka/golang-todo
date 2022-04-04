package models

import (
	"log"
	"time"
)

type User struct {
	ID       int
	UUID     string
	Name     string
	Email    string
	Password string
	CreateAt time.Time
}

// Method of type User, return value is an error
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at
	) values (?,?,?,?,?)`

	// Db exists in the models package, so there is no need to import it
	_, err = Db.Exec(
		cmd, createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now(),
	)

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
