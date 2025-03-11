package requests

import (
	"APIControlID/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func AddImage(img []byte, token string, id int) ([]byte, error) {
	sId := strconv.Itoa(id)

	url, err := utils.UrlBuilder("/user_set_image.fcgi", map[string]interface{}{"session": token, "user_id": sId})

	fmt.Println(url)
	// Criar um buffer para os dados da imagem
	imgBuffer := bytes.NewBuffer(img)

	// Fazer a requisição POST
	resp, err := http.Post(url, "application/octet-stream", imgBuffer)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar a imagem: %v", err)
	}
	defer resp.Body.Close()

	// Ler resposta da API
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta da API: %v", err)
	}

	// Verificar se a resposta foi bem-sucedida
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro ao definir imagem do usuário, status: %d, resposta: %s", resp.StatusCode, string(responseBody))
	}
	return responseBody, nil
}

func DelImage(userID int, token string) error {

	urlDelImage, err := utils.UrlBuilder("/user_destroy_image.fcgi", map[string]interface{}{"session": token})
	if err != nil {
		return fmt.Errorf("erro ao montar a requisiçao: %v", err)

	}
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
	return nil

}
