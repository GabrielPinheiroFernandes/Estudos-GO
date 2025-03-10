package crudapi

import (
	"APIControlID/structs"
	"APIControlID/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Definição da struct ControlIdCrudApi
type ControlIdCrudApi struct {
	token string
}

// Inicializa a instância de ControlIdCrudApi e retorna o token
func NewControlIdCrudApi() (*ControlIdCrudApi, error) {
	token, err := getToken()
	if err != nil {
		return nil, err
	}

	return &ControlIdCrudApi{
		token: token,
	}, nil
}

func (c *ControlIdCrudApi) AddUser(u structs.User) ([]byte, error) {
	// Construir URL da API para criação de usuário
	url, _ := utils.UrlBuilder("/create_objects.fcgi", map[string]interface{}{"session": c.token})
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

	// **Segunda requisição para adicionar o usuário a um grupo**
	groupURL := os.Getenv("API_URL") + "/create_objects.fcgi?session=" + c.token
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
		// Caso falhe, apagar o usuário criado
		c.DelUser(userID) // Chama a função para deletar o usuário

		return nil, fmt.Errorf("erro ao associar usuário ao grupo: %v", err)
	}
	defer groupResp.Body.Close()

	// Ler resposta da API
	groupResponseBody, err := io.ReadAll(groupResp.Body)
	if err != nil {
		// Caso falhe, apagar o usuário criado
		c.DelUser(userID)

		return nil, fmt.Errorf("erro ao ler resposta da associação de grupo: %v", err)
	}

	// Verificar resposta da API para a associação ao grupo
	if groupResp.StatusCode != http.StatusOK {
		// Caso falhe, apagar o usuário criado
		c.DelUser(userID)

		return nil, fmt.Errorf("erro ao associar usuário ao grupo, status: %d, corpo: %s", groupResp.StatusCode, string(groupResponseBody))
	}

	// Se a associação ao grupo for bem-sucedida, continuar com a criação do cartão
	cardval := convertCard(u.Card_value)

	// URL da API para criar o cartão
	cardUrl, _ := utils.UrlBuilder("/create_objects.fcgi", map[string]interface{}{"session": c.token})

	// Criar o corpo da requisição para adicionar o cartão
	icardVal, err := strconv.Atoi(cardval)
	if err != nil {
		// Caso falhe, apagar o usuário criado
		c.DelUser(userID)

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
		c.DelUser(userID)

		return nil, fmt.Errorf("erro ao marshalling JSON do cartão: %v", err)
	}

	// Enviar a requisição HTTP para criar o cartão
	resp, err = http.Post(cardUrl, "application/json", bytes.NewBuffer(cardBodyJSON))
	if err != nil {
		// Caso falhe, apagar o usuário criado
		c.DelUser(userID)

		return nil, fmt.Errorf("erro ao enviar requisição para cartão: %v", err)
	}

	// Verificar se a resposta HTTP é bem-sucedida (status 2xx)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Caso falhe, apagar o usuário criado
		c.DelUser(userID)

		return nil, fmt.Errorf("erro na resposta HTTP do cartão, código: %d", resp.StatusCode)
	}

	// Retornar resposta JSON
	return responseBody, nil
}

func (c *ControlIdCrudApi) AddImageUser(id int, img []byte) error {
	// Criar a URL com o ID do usuário
	url := os.Getenv("API_URL") + "/user_set_image.fcgi?session=" + c.token + "&user_id=" + strconv.Itoa(id)
	fmt.Println(url)
	// Criar um buffer para os dados da imagem
	imgBuffer := bytes.NewBuffer(img)

	// Fazer a requisição POST
	resp, err := http.Post(url, "application/octet-stream", imgBuffer)
	if err != nil {
		return fmt.Errorf("erro ao enviar a imagem: %v", err)
	}
	defer resp.Body.Close()

	// Ler resposta da API
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao ler resposta da API: %v", err)
	}

	// Verificar se a resposta foi bem-sucedida
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao definir imagem do usuário, status: %d, resposta: %s", resp.StatusCode, string(responseBody))
	}

	fmt.Println("Imagem enviada com sucesso! Resposta da API:", string(responseBody))
	return nil
}

func (c ControlIdCrudApi) DelUser(userID int) error {
	// ===========================
	// 1ª Requisição: Deletar Imagem do Usuário
	// ===========================
	urlDelImage := os.Getenv("API_URL") + "/user_destroy_image.fcgi?session=" + c.token

	// JSON para deletar imagem
	bodyDelImage := map[string]interface{}{
		"user_ids": []int{userID},
	}

	// Converte JSON para []byte
	jsonBody, err := json.Marshal(bodyDelImage)
	if err != nil {
		return fmt.Errorf("erro ao serializar JSON da imagem: %v", err)
	}

	// Enviar requisição HTTP POST
	resp, err := http.Post(urlDelImage, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("erro ao enviar requisição para deletar imagem: %v", err)
	}
	defer resp.Body.Close()

	// Verificar resposta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro ao deletar imagem, status: %d", resp.StatusCode)
	}

	// ===========================
	// 2ª Requisição: Deletar Usuário
	// ===========================
	urlDelUser := os.Getenv("API_URL") + "/destroy_objects.fcgi?session=" + c.token

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
	jsonBody, err = json.Marshal(bodyDelUser)
	if err != nil {
		return fmt.Errorf("erro ao serializar JSON do usuário: %v", err)
	}

	// Enviar requisição HTTP POST
	resp, err = http.Post(urlDelUser, "application/json", bytes.NewBuffer(jsonBody))
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

func (c ControlIdCrudApi) EditUser(u structs.User) (structs.User, error) {

	// Construir a URL com o token de sessão
	url, err := utils.UrlBuilder("modify_objects.fcgi?", map[string]interface{}{"session": c.token})
	if err != nil {
		return structs.User{}, err
	}

	// Criar o corpo da requisição
	body := map[string]interface{}{
		"object": "users",
		"values": map[string]interface{}{
			"begin_time":  12,
			"end_time":    12,
			"Name":        u.Name,
			"Pass":        u.Pass,
			"Card_value":  u.Card_value,
		},
		"where": map[string]interface{}{
			"users": map[string]interface{}{
				"id": u.Id, // Agora usa o ID do usuário passado na função
			},
		},
	}

	// Converter o body para JSON
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return structs.User{}, err
	}

	// Fazer a requisição HTTP POST
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyJson))
	if err != nil {
		return structs.User{}, err
	}
	defer resp.Body.Close()

	// Ler a resposta da API
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return structs.User{}, err
	}

	// Verificar se a API retornou sucesso (pode variar de acordo com a API)
	if resp.StatusCode != 200 {
		return structs.User{}, fmt.Errorf("falha na edição do usuário: %s", string(respBody))
	}

	// Converter a resposta para a struct User (caso a API retorne os dados atualizados)
	var updatedUser structs.User
	if err := json.Unmarshal(respBody, &updatedUser); err != nil {
		return structs.User{}, fmt.Errorf("falha ao processar resposta da API: %v", err)
	}

	return updatedUser, nil
}
