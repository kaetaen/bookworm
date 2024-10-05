package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/kaetaen/bookworm/config"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) Save() (int64, error) {
	db, err := sql.Open("mysql", config.Env["MYSQL_URI"])
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		u.Username,
		u.Email,
		u.Password,
	)
	if err != nil {
		return 0, fmt.Errorf("Error when try save User: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error when try save User: %v", err)
	}
	return id, nil

}
