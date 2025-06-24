package models

import (
	"context"
	"crud-backend/utils"
	"strconv"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Id         int    `db:"id" json:"id"`
	Username   string `db:"username" json:"username"`  
	Email      string `db:"email" json:"email"`
	Phone      string  `db:"phone" json:"phone"`        
	Password   string `db:"password" json:"password"`         
}

func FindAllUsers() ([]User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return []User{}, err
	}
	defer conn.Close() 

	query := `SELECT id, username, email FROM users`

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

