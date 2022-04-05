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

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
	// QueryRow(), Scan(): https://golang.shop/post/go-databasesql-04-retrieving-ja/
	// Query(): Retrieve multiple search results(rows)
	// QueryRow(): Retrieve a single row
	// Exec(): if you do not want to retrieve search results(CREATE, INSERT, UPDATE, DELETE etc.) https://sourjp.github.io/posts/go-db/
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreateAt,
	)
	return user, err
}

// Update name and email
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// Delete a user
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
