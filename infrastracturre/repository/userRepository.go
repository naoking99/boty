package repository

import (
	"github.com/naoking99/boty/domain"
	"github.com/naoking99/boty/utils/mysql"
)

// GetUsers is
func GetUsers() *[]*domain.User {
	db := mysql.GetDB()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	users := []*domain.User{}

	for rows.Next() {
		var id int
		var email string

		err := rows.Scan(&id, &email)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, domain.CreateUser(id, "testname", email))
	}

	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	return &users
}
