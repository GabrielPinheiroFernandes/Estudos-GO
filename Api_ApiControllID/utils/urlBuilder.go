package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// UrlBuilder constrói uma URL com base nos parâmetros fornecidos
func UrlBuilder(route string, params ...map[string]interface{}) (string, error) {
	// Se não houver parâmetros, apenas retorne a URL base
	if len(params) == 0 {
		return route, nil
	}
	// Carregar as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("erro ao carregar o .env: %v", err)
	}

	// Caso tenha parâmetros, use o primeiro map passado
	paramMap := params[0]

	// Construa a URL com base nos parâmetros
	queryParams := "?"
	for key, value := range paramMap {
		// Adiciona os parâmetros de forma segura na query string
		queryParams += fmt.Sprintf("%s=%v&", key, value)
	}

	// Remove o último '&' extra
	queryParams = queryParams[:len(queryParams)-1]

	urlBase:=os.Getenv("API_URL")
	// fmt.Println(urlBase + route + queryParams)
	// Retorne a URL com os parâmetros
	return urlBase + route + queryParams, nil
}
