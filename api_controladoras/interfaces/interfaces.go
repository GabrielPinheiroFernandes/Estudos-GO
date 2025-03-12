package interfaces

import "APIControlID/structs"

type CrudApi interface {
	AddUser(structs.User) ([]byte, error)
	DelUser(int) error
	AddImageUser(int, []byte) error
	EditUser(int, structs.User) ([]byte, error)
}
