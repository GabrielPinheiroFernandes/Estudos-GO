package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func connect() (*sql.DB, error) {
	// Abre a conexão com o banco de dados SQLite
	db, err := sql.Open("sqlite3", "/home/gabriel/Desktop/estudo/user_crud/database/sqlite/sqlite.db")
	if err != nil {
		return nil, err
	}

	// Verifica se a conexão foi bem-sucedida
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Retorna a conexão bem-sucedida
	return db, nil
}

func close(db *sql.DB) error {
	// Fecha a conexão com o banco de dados
	err := db.Close()
	if err != nil {
		return err
	}
	// Retorna nil caso a conexão seja fechada com sucesso
	return nil
}

func ExecQuerySqlite(sqlQuery string, params []interface{}) (*sql.Rows, error) {
	// Estabelece a conexão com o banco de dados
	db, err := connect() // A função connect() deve retornar uma *sql.DB
	if err != nil {
		fmt.Println("Houve algum erro de conexão:", err)
		return nil, err
	}
	defer db.Close() // Garante que a conexão será fechada ao final da execução

	// Executa a consulta
	rows, err := db.Query(sqlQuery, params...)
	if err != nil {
		return nil, err
	}
	// defer rows.Close() // Fecha o conjunto de linhas quando a função terminar

	return rows, nil
}
