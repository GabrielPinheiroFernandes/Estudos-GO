package repository

import (
	"errors"
	"fmt"
	"userCrud/database"
	"userCrud/structures"
)

var ()

type LocalUserRepository struct{}

func (lu LocalUserRepository) AddUser(u structures.User) (int, error) {
	database.UserTable = append(database.UserTable, u)
	return len(database.UserTable), nil
}

func (lu LocalUserRepository) GetAllUsers() ([]structures.User, error) {
	users := database.UserTable
	if len(users) == 0 {
		return users, fmt.Errorf("não existem usuarios na tabela")
	}
	return users, nil
}

func (lu LocalUserRepository) GetUserByID(id int) (structures.User, error) {
	// Obtém a tabela de usuários
	users := database.UserTable

	// Itera pela lista de usuários
	for _, user := range users {
		if user.Id == id {
			// fmt.Println(user)

			return user, nil // Retorna o usuário encontrado e nenhum erro
		}
	}

	// Retorna erro caso o ID não seja encontrado
	return structures.User{}, errors.New("usuário não encontrado")
}

func (lu LocalUserRepository) DelUser(id int) (int, error) {
	users := database.UserTable

	// Encontra o índice do usuário pelo ID
	for i, user := range users {
		if user.Id == id {
			// Remove o usuário do slice
			database.UserTable = append(users[:i], users[i+1:]...)
			return id, nil // Usuário deletado com sucesso
		}
	}

	// Se o usuário não for encontrado, retorna erro

	return id, errors.New("usuario não encontrado")

}
