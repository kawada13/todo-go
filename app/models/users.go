package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

//　ユーザー作成
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(
		cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrpt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// ユーザーの更新

func (u User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

//　ユーザーの削除

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`

	_, err = Db.Exec(cmd, u.ID)

	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// ユーザーを取得
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select * from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}
