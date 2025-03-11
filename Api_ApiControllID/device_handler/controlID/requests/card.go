package requests

import (
	"APIControlID/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AddUserCard(userID int, token string, cv string) ([]byte, error) {
	// Se a associação ao grupo for bem-sucedida, continuar com a criação do cartão
	cardval := utils.ConvertCard(cv)

	// URL da API para criar o cartão
	cardUrl, _ := utils.BuildUrl("/create_objects.fcgi", map[string]interface{}{"session": token})

	// Criar o corpo da requisição para adicionar o cartão
	icardVal, err := strconv.Atoi(cardval)
	if err != nil {
		// Caso falhe, apagar o usuário criado
		return nil, fmt.Errorf("erro ao converter card para int: %v", err)
	}
	cardBody := map[string]interface{}{
		"object": "cards",
		"values": []interface{}{
			map[string]interface{}{
				"value":   icardVal,
				"user_id": userID, // Certifique-se de que userID está definido corretamente
			},
		},
	}

	// Converter o corpo para JSON
	cardBodyJSON, err := json.Marshal(cardBody)
	if err != nil {
		// Caso falhe, apagar o usuário criado
		return nil, fmt.Errorf("erro ao marshalling JSON do cartão: %v", err)
	}

	// Enviar a requisição HTTP para criar o cartão
	resp, err := http.Post(cardUrl, "application/json", bytes.NewBuffer(cardBodyJSON))
	if err != nil {
		// Caso falhe, apagar o usuário criado
		return nil, fmt.Errorf("erro ao enviar requisição para cartão: %v", err)
	}

	// Verificar se a resposta HTTP é bem-sucedida (status 2xx)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Caso falhe, apagar o usuário criado
		return nil, fmt.Errorf("erro na resposta HTTP do cartão, código: %d", resp.StatusCode)
	}

	return cardBodyJSON, nil
}
