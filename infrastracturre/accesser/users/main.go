package users

import (
	"fmt"

	"github.com/naoking99/boty/domain/user"
	"github.com/naoking99/boty/utils/mysql"
)

// Accesser is
type Accesser struct{}

// GetAll is
func (a Accesser) GetAll() *[]*user.User {
	db := mysql.GetDB()
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
func (a Accesser) Save(u *user.User) {
	db := mysql.GetDB()

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
