package crudapi

import (
	"APIControlID/device_handler/controlID/requests"
	"APIControlID/structs"
	"encoding/json"
	"fmt"
)

// Definição da struct ControlIdCrudApi
type ControlIdCrudApi struct {
	token string
}

// Inicializa a instância de ControlIdCrudApi e retorna o token
func NewControlIdCrudApi() (*ControlIdCrudApi, error) {
	token, err := requests.GetToken()
	if err != nil {
		return nil, err
	}

	return &ControlIdCrudApi{
		token: token,
	}, nil
}

func (c *ControlIdCrudApi) AddUser(u structs.User) ([]byte, error) {
	responseBody, err := requests.AddUser(u, c.token)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(responseBody, &jsonResponse); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
	}

	// Buscar o ID do usuário
	ids, ok := jsonResponse["ids"].([]interface{})
	if !ok || len(ids) == 0 {
		return nil, fmt.Errorf("erro: ID do usuário não encontrado na resposta: %s", string(responseBody))
	}
	userIDFloat, ok := ids[0].(float64) // O JSON pode interpretar números como float64
	if !ok {
		return nil, fmt.Errorf("erro ao converter user_id para int64")
	}
	userID := int(userIDFloat)

	_, err = requests.AddUserGroup(userID, c.token)
	if err != nil {
		c.DelUser(userID)
		return nil, fmt.Errorf("erro ao Concluir requisição:  %v", err)
	}

	_, err = requests.AddUserCard(userID, c.token, u.Card_value)
	if err != nil {
		c.DelUser(userID)
		return nil, fmt.Errorf("erro ao Concluir requisição:  %v", err)
	}
	// Retornar resposta JSON
	return responseBody, nil
}

func (c *ControlIdCrudApi) AddImageUser(id int, img []byte) error {
	// Criar a URL com o ID do usuário

	responseBody, err := requests.AddImage(img, c.token, id)
	if err != nil {
		return fmt.Errorf("erro ao Fazer requisição: %v", err)
	}
	fmt.Println("Imagem enviada com sucesso! Resposta da API:", string(responseBody))
	return nil
}

func (c ControlIdCrudApi) DelUser(userID int) error {

	err := requests.DelUser(userID, c.token)
	if err != nil {
		return fmt.Errorf("erro ao enviar requisição para deletar imagem do usuário: %v", err)
	}
	err = requests.DelUser(userID, c.token)
	if err != nil {
		return fmt.Errorf("erro ao enviar requisição para deletar imagem do usuário: %v", err)
	}
	return nil
}

func (c ControlIdCrudApi) EditUser(id int, u structs.User) ([]byte, error) {

	response, err := requests.EditUser(id, c.token, u)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição para editar o usuário: %v", err)
	}

	return response, nil
}
