package database

import "userCrud/structures"

var(
	UserTable []structures.User
)

func init() {
	UserTable = []structures.User{}
}
