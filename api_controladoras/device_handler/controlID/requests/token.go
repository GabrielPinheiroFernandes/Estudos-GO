package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetToken() (string, error) {

	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("erro ao carregar variáveis de ambiente: %v", err)
	}

	// Definir URL da API de login
	url := os.Getenv("API_URL") + "/login.fcgi"

	// Definir corpo da requisição com credenciais
	body := map[string]interface{}{
		"login":    os.Getenv("USER_LOGIN"),
		"password": os.Getenv("USER_PASS"),
	}

	// Converter o corpo para JSON
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("erro ao converter o body: %v", err)
	}

	// Fazer a requisição POST para login
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyJson))
	if err != nil {
		return "", fmt.Errorf("erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	// Ler resposta
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erro ao ler resposta: %v", err)
	}

	// Decodificar resposta JSON
	var responseData map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseData); err != nil {
		return "", fmt.Errorf("erro ao decodificar resposta JSON: %v\nResposta: %s", err, string(responseBody))
	}

	// Extrair o token da resposta
	token, ok := responseData["session"].(string)
	if !ok || token == "" {
		return "", fmt.Errorf("token não encontrado ou inválido na resposta: %s", string(responseBody))
	}

	// Retornar a instância com o token
	return token, nil
}
