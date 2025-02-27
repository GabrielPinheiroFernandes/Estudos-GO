package database

import (
	"fmt"
	"log"
	"userCrud/database/sqlite"
	"userCrud/structures"
)

var (
	UserTable []structures.User
)

func init() {
	UserTable = []structures.User{}
}

func sqliteInit() {
	sql := "SELECT * FROM Users"
	params := []interface{}{}

	// Executa a query e obtém os resultados
	rows, err := sqlite.ExecQuerySqlite(sql, params)
	if err != nil {
		fmt.Println("Erro ao executar query", err)
		return
	}
	defer rows.Close() // Garante que as linhas serão fechadas após o uso

	// Itera sobre os resultados da query
	for rows.Next() {

		var id int
		var username string

		// Scaneia os dados das colunas para as variáveis
		err := rows.Scan(&id, &username)
		if err != nil {
			log.Fatal(err)
		}

		// Imprime os dados obtidos
		fmt.Printf("ID: %d, Username: %s\n", id, username)
	}

	// Verifica se ocorreu algum erro ao iterar sobre as linhas
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
