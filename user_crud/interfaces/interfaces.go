package interfaces

import "userCrud/structures"


type UserRepository interface {
	GetUserByID(id int) (structures.User,error)
	GetAllUsers() ([]structures.User ,error)
	AddUser(u structures.User) (int,error)
	DelUser(id int) (int,error)
}
