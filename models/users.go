package models

import (
	"context"
	"crud-backend/utils"
	"fmt"
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

	query := `SELECT id, username, email, phone, password FROM users`
	fmt.Println("🔍 Executing query:", query)

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