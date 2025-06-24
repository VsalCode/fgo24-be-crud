package models

import (
	"context"
	"crud-backend/utils"
	"strconv"
	"github.com/jackc/pgx/v5"
)
type User struct {
	Id         int    `form:"id" db:"id" json:"id"`
	Username   string `form:"username" db:"username" json:"username"`  
	Email      string `form:"email" db:"email" json:"email"`
	Phone      string  `form:"phone" db:"phone" json:"phone"`        
	Password   string `form:"password" db:"password" json:"password"`         
}

func FindAllUsers() ([]User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return []User{}, err
	}
	defer conn.Close() 

	query := `SELECT id, username, email, phone, password FROM users`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func GetUserDetail(param string) ([]User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return []User{}, err
	}
	defer conn.Close() 

	query := `SELECT id, username, email, phone, password FROM users WHERE id = $1`

	id, _ := strconv.Atoi(param)

	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func HandleCreateUser(user User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `INSERT INTO users (username, email, phone, password) VALUES ($1, $2, $3, md5($4))
	`

	_, err = conn.Exec(context.Background(), query, user.Username, user.Email, user.Phone, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func HandleDeleteUser(id string) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `DELETE FROM users WHERE id = $1`

	idInt, _ := strconv.Atoi(id)

	_, err = conn.Exec(context.Background(), query, idInt)
	if err != nil {
		return err
	}

	return nil
}