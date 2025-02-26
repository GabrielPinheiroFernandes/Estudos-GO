package sqlite

import (
	"database/sql"
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
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer close(db) // Garante que a conexão será fechada ao final da execução

	// Executa a consulta
	rows, err := db.Query(sqlQuery, params...)
	if err != nil {
		return nil, err
	}

	// Retorna as linhas da consulta
	return rows, nil
}
