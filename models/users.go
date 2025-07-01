package models

import (
	"context"
	"crud-backend/dto"
	"crud-backend/utils"
	"encoding/json"
	"strconv"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id       int    `form:"id" db:"id" json:"id"`
	Username string `form:"username" db:"username" json:"username"`
	Email    string `form:"email" db:"email" json:"email"`
	Phone    string `form:"phone" db:"phone" json:"phone"`
	Password string `form:"password" db:"password" json:"password"`
}

func FindUserByName(search string) ([]User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return []User{}, err
	}
	defer conn.Close()

	var rows pgx.Rows
	if search == "" {
		rows, err = conn.Query(
			context.Background(),
			`SELECT id, username, email, phone, password FROM users`,
		)
	} else {
		rows, err = conn.Query(
			context.Background(),
			`SELECT id, username, email, phone, password FROM users 
     WHERE username ILIKE $1`,
			"%"+search+"%")
	}
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

func GetUserDetail(param string) (User, error) {
    redisKey := "id:" + param
    result := utils.RedisClient.Exists(context.Background(), redisKey)

    if result.Val() != 0 {
        data := utils.RedisClient.Get(context.Background(), redisKey)
        str := data.Val()
        users := User{}
        if err := json.Unmarshal([]byte(str), &users); err != nil {
            return User{}, err
        }
        return users, nil
    }

    conn, err := utils.DBConnect()
    if err != nil {
        return User{}, err
    }
    defer conn.Close()

    query := `SELECT id, username, email, phone, password FROM users WHERE id = $1`
    id, _ := strconv.Atoi(param)

    rows, err := conn.Query(context.Background(), query, id)
    if err != nil {
        return User{}, err
    }
    defer rows.Close()

    users, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
    if err != nil {
        return User{}, err
    }

    encoded, err := json.Marshal(users)
    if err == nil {
        utils.RedisClient.Set(context.Background(), redisKey, string(encoded), 0)
    }

    return users, nil
}

func HandleCreateUser(user dto.CreateUserRequest) error {
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

func HandleUpdateUser(id string, user dto.UpdateUserRequest) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	querySelect := `SELECT username, email, phone, password FROM users WHERE id = $1`
	row := conn.QueryRow(context.Background(), querySelect, id)

	var currentUsername, currentEmail, currentPhone, currentPassword string
	err = row.Scan(&currentUsername, &currentEmail, &currentPhone, &currentPassword)
	if err != nil {
		return err
	}

	_, err = conn.Exec(
		context.Background(),
		`
		UPDATE users SET 
    username = COALESCE($1, $2), 
    email = COALESCE($3, $4), 
    phone = COALESCE($5, $6), 
    password = COALESCE(md5($7), md5($8))
    WHERE id = $9
		`,
		user.Username, currentUsername,
		user.Email, currentEmail,
		user.Phone, currentPhone,
		user.Password, currentPassword,
		id)
	if err != nil {
		return err
	}

	return nil
}
