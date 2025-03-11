package requests

import (
	"APIControlID/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AddUserGroup(userID int, token string) ([]byte, error) {

	// **Segunda requisição para adicionar o usuário a um grupo**
	groupURL, _ := utils.UrlBuilder("/create_objects.fcgi", map[string]interface{}{"session": token})
	// Criar corpo da requisição para adicionar o usuário ao grupo
	groupBody := map[string]interface{}{
		"object": "user_groups",
		"fields": []interface{}{"user_id", "group_id"},
		"values": []interface{}{
			map[string]interface{}{
				"user_id":  userID,
				"group_id": 1,
			},
		},
	}

	// Converter corpo para JSON
	groupBodyJson, err := json.Marshal(groupBody)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter body para JSON: %v", err)
	}

	// Enviar requisição HTTP POST para adicionar usuário ao grupo
	groupResp, err := http.Post(groupURL, "application/json", bytes.NewBuffer(groupBodyJson))
	if err != nil {
		return nil, fmt.Errorf("erro ao associar usuário ao grupo: %v", err)
	}
	defer groupResp.Body.Close()

	// Ler resposta da API
	groupResponseBody, err := io.ReadAll(groupResp.Body)
	if err != nil {
		// Caso falhe, apagar o usuário criado
		return nil, fmt.Errorf("erro ao ler resposta da associação de grupo: %v", err)
	}

	// Verificar resposta da API para a associação ao grupo
	if groupResp.StatusCode != http.StatusOK {
		// Caso falhe, apagar o usuário criado
		return nil, fmt.Errorf("erro ao associar usuário ao grupo, status: %d, corpo: %s", groupResp.StatusCode, string(groupResponseBody))
	}
	return groupResponseBody, nil

}
