package requests

import (
	"APIControlID/structs"
	"APIControlID/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AddUser(u structs.User, token string) ([]byte, error) {
	// Construir URL da API para criação de usuário
	url, _ := utils.BuildUrl("/create_objects.fcgi", map[string]interface{}{"session": token})
	fmt.Print(url)
	// Criar corpo da requisição para adicionar usuário
	userBody := map[string]interface{}{
		"object": "users",
		"values": []interface{}{
			map[string]interface{}{
				"name":         u.Name,
				"registration": "",
				"password":     "",
				"salt":         "",
			},
		},
	}

	// Converter corpo para JSON
	userBodyJson, err := json.Marshal(userBody)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter body para JSON: %v", err)
	}

	// Enviar requisição HTTP POST para criar usuário
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(userBodyJson))
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %v", err)
	}
	defer resp.Body.Close()

	// Ler resposta da API
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %v", err)
	}

	// Verificar resposta da API
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao criar usuário, status: %d, corpo: %s", resp.StatusCode, string(responseBody))
	}

	// Extrair o ID do usuário da resposta

	return responseBody, nil
}

func DelUser(userID int, token string) error {
	// ===========================
	// 2ª Requisição: Deletar Usuário
	// ===========================
	urlDelUser, err := utils.BuildUrl("/destroy_objects.fcgi", map[string]interface{}{"session": token})
	if err != nil {
		return fmt.Errorf("erro ao montar a requisiçao: %v", err)
	}
	// JSON para deletar usuário
	bodyDelUser := map[string]interface{}{
		"object": "users",
		"where": map[string]interface{}{
			"users": map[string]interface{}{
				"id": []int{userID},
			},
		},
	}

	// Converte JSON para []byte
	jsonBody, err := json.Marshal(bodyDelUser)
	if err != nil {
		return fmt.Errorf("erro ao serializar JSON do usuário: %v", err)
	}

	// Enviar requisição HTTP POST
	resp, err := http.Post(urlDelUser, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("erro ao enviar requisição para deletar usuário: %v", err)
	}
	defer resp.Body.Close()

	// Verificar resposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao deletar usuário, status: %d", resp.StatusCode)
	}
	return nil

}

func EditUser(userID int, token string, u structs.User) ([]byte, error) {
	// Construir a URL com o token de sessão
	url, err := utils.BuildUrl("/modify_objects.fcgi", map[string]interface{}{"session": token})
	if err != nil {
		return nil, err
	}
	fmt.Println(userID)
	// Criar o corpo da requisição
	body := map[string]interface{}{
		"object": "users",
		"values": map[string]interface{}{
			"begin_time": 1745568008, // Certifique-se de que este é o formato correto esperado pela API
			"end_time":   64818,      // Se for timestamp, deve ser um número UNIX Timestamp
			"name":       u.Name,     // Pode precisar ser "Name" (com maiúscula) dependendo da API
			"password":   u.Pass,     // Pode precisar ser "Pass" dependendo da API
		},
		"where": map[string]interface{}{
			"users": map[string]interface{}{
				"id": userID, // Deve ser um número inteiro válido
			},
		},
	}

	// Converter o body para JSON
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Fazer a requisição HTTP POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Ler a resposta da API
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Verificar se a API retornou sucesso (pode variar de acordo com a API)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("falha na edição do usuário: %s", string(respBody))
	}

	// Converter a resposta para a struct User (caso a API retorne os dados atualizados)
	return respBody, nil
}
