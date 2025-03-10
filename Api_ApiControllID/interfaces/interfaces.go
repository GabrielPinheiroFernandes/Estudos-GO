package interfaces

import "APIControlID/structs"

type CrudApi interface {
	AddUser(structs.User) ([]byte, error)
	DelUser(int) error
	AddImageUser(int, []byte) error
	EditUser(structs.User) (structs.User, error)
}
