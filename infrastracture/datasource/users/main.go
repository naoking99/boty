package users

import (
	"fmt"

	"github.com/naoking99/boty/domain/user"
	"github.com/naoking99/boty/infrastracture/config/mysql"
)

// Repository is
type Repository struct{}

// GetAll is
func (r Repository) GetAll() *[]*user.User {
	db := mysql.DBPool()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	users := []*user.User{}

	for rows.Next() {
		var id int
		var email string

		err := rows.Scan(&id, &email)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user.New(id, "testname", email))
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	return &users
}

// Save is
func (r Repository) Save(u *user.User) {
	db := mysql.DBPool()

	stmtInsert, err := db.Prepare("INSERT INTO users(email) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(u.Email())
	if err != nil {
		panic(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(lastInsertID)
}
