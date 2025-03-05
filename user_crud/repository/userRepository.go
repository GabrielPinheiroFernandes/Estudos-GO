package repository

import (
	"errors"
	"fmt"
	"userCrud/database"
	"userCrud/database/sqlite"
	"userCrud/structures"
)

type LocalUserRepository struct{}
type SqliteUserRepository struct{}

//=======================================================
//=======================LOCAL===========================
//=======================================================

func (lu *LocalUserRepository) AddUser(u structures.User) (int, error) {
	database.UserTable = append(database.UserTable, u)
	return len(database.UserTable), nil
}

func (lu *LocalUserRepository) GetAllUsers() ([]structures.User, error) {
	users := database.UserTable
	if len(users) == 0 {
		return users, fmt.Errorf("não existem usuarios na tabela")
	}
	return users, nil
}

func (lu *LocalUserRepository) GetUserByID(id int) (structures.User, error) {
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

func (lu *LocalUserRepository) DelUser(id int) (int, error) {
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

//=======================================================
//=======================SQLITE3=========================
//=======================================================

func (s *SqliteUserRepository) GetUserByID(id int) (structures.User, error) {
	sql := "SELECT * FROM Users WHERE ID=?"
	params := []interface{}{id}
	data, err := sqlite.ExecQuerySqlite(sql, params)
	
	if (err != nil){
		return structures.User{}, err
	}
	

	var usr structures.User
	for data.Next() {

		var id int
		var name, username, pass, imagePath string

		err := data.Scan(&id, &name, &username, &pass, &imagePath)
		if err != nil || id < 0{
			return structures.User{}, err
		}

		unHashPass := pass
		usr = structures.User{
			Id:        id,
			Name:      name,
			Username:  username,
			Pass:      unHashPass,
			ImagePath: imagePath,
		}
	}
	return usr, err
}
func (s *SqliteUserRepository) GetAllUsers() ([]structures.User, error) {
	sql := "SELECT * FROM Users"
	data, err := sqlite.ExecQuerySqlite(sql, []interface{}{})
	if err != nil {
		return []structures.User{}, err
	}
	defer data.Close() // Sempre feche os resultados quando terminar

	slOfUs := []structures.User{}

	// Itera sobre as linhas retornadas
	for data.Next() {
		var id int
		var name, username, pass, imagePath string

		// Escaneia os valores das colunas na linha atual
		err := data.Scan(&id, &name, &username, &pass, &imagePath)
		if err != nil {
			return nil, err // Se ocorrer erro ao escanear a linha, retorne o erro
		}

		// Adiciona o usuário à lista
		slOfUs = append(slOfUs, structures.User{
			Id:        id,
			Name:      name,
			Username:  username,
			Pass:      pass,
			ImagePath: imagePath,
		})
	}

	// Verifica se ocorreu algum erro durante a iteração
	if err := data.Err(); err != nil {
		return nil, err
	}

	return slOfUs, nil
}

func (s *SqliteUserRepository) AddUser(u structures.User) (int, error) {
	// fmt.Print("Chegou aqui!")

	hashedPass := u.Pass //Implementar hash de senha no futuro
	sql := "INSERT INTO Users (name, username, pass, ImagePath) VALUES (?, ?, ?, ?) RETURNING Id"
	params := []interface{}{u.Name, u.Username, hashedPass, u.ImagePath}
	// fmt.Print(params...)

	// Executa a consulta
	data, err := sqlite.ExecQuerySqlite(sql, params)
	if err != nil {
		return 0, err
	}
	defer data.Close() // Fechar os resultados ao final

	var lastId int
	// Itera sobre as linhas retornadas
	if data.Next() {
		// Escaneia o valor do ID retornado pelo INSERT
		err := data.Scan(&lastId)
		if err != nil {
			return 0, err
		}
	}

	// Verifica se ocorreu algum erro durante a iteração das linhas
	if err := data.Err(); err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *SqliteUserRepository) DelUser(id int) (int, error) {
	sql := "DELETE FROM Users WHERE ID=?"
	params := []interface{}{id}

	_, err := sqlite.ExecQuerySqlite(sql, params)

	if err != nil {
		// fmt.Print("Algo deu errado na exclusao!")
		return 0, err
	}

	return id, err

}
