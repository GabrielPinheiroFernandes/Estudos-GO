package controllers

import (
	"APIControlID/interfaces"
	"APIControlID/structs"
	"APIControlID/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	crudApi interfaces.CrudApi
}

// Inicializa o Controller com a interface CrudApi
func NewController(c interfaces.CrudApi) *Controller {
	return &Controller{
		crudApi: c,
	}
}

func (c *Controller) Inicialize() {
	r := gin.Default()

	// Rota de teste
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// Rota de teste para URLBuilder
	r.POST("/urlteste", func(ctx *gin.Context) {
		url, err := utils.UrlBuilder("/paozinquente", map[string]interface{}{
			"id":        12,
			"profissao": "roberto fotografias",
		})
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Falha ao gerar URL!", "erro": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"URL": url})
	})

	// Criando grupo de rotas para usuário
	grpRoutesUser := r.Group("/user")
	{
		// Adicionar usuário
		grpRoutesUser.POST("/add", func(ctx *gin.Context) {
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				ctx.JSON(400, gin.H{"message": "Falha ao ler o corpo da requisição"})
				return
			}

			// Recriar o corpo da requisição para permitir o binding
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			var user structs.User
			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(400, gin.H{"message": "JSON inválido"})
				return
			}

			// Chama a função AddUser na API
			data, err := c.crudApi.AddUser(user)
			if err != nil {
				ctx.JSON(500, gin.H{"message": "Falha ao adicionar usuário", "error": err.Error()})
				return
			}

			// Converte a resposta da API para JSON
			var jsonResponse map[string]interface{}
			if err := json.Unmarshal(data, &jsonResponse); err != nil {
				ctx.JSON(500, gin.H{"message": "Falha ao converter resposta", "details": err.Error()})
				return
			}

			ctx.JSON(200, gin.H{"message": "Usuário adicionado com sucesso", "data": jsonResponse})
		})

		// Criando grupo de rotas para imagens do usuário
		grpRoutesImageUser := grpRoutesUser.Group("/image")
		{
			// Adicionar imagem ao usuário
			grpRoutesImageUser.POST("/add/:id", func(ctx *gin.Context) {
				id := ctx.Param("id")

				bodyBytes, err := io.ReadAll(ctx.Request.Body)
				if err != nil {
					ctx.JSON(400, gin.H{"message": "Falha ao ler o corpo da requisição"})
					return
				}

				idInt, err := strconv.Atoi(id)
				if err != nil {
					ctx.JSON(400, gin.H{"message": "ID inválido"})
					return
				}

				err = c.crudApi.AddImageUser(idInt, bodyBytes)
				if err != nil {
					ctx.JSON(500, gin.H{"message": "Falha ao adicionar imagem ao usuário"})
					return
				}

				ctx.JSON(200, gin.H{"message": "Imagem adicionada com sucesso", "id": id})
			})
		}

		grpRoutesUser.PUT("/edit", func(ctx *gin.Context) {
			ctx.JSON(200,gin.H{})
		})

		// Deletar usuário
		grpRoutesUser.DELETE("/del/:id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				ctx.JSON(400, gin.H{"message": "ID inválido"})
				return
			}

			err = c.crudApi.DelUser(id)
			if err != nil {
				ctx.JSON(500, gin.H{"message": "Falha ao deletar usuário"})
				return
			}

			ctx.JSON(200, gin.H{"message": "Usuário deletado com sucesso", "id": id})
		})
		r.POST("api/v1/notifications", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{})
		})
		controlIDNotifications := r.Group("/api/notifications")
		{
			controlIDNotifications.POST("/dao", func(ctx *gin.Context) {
				printRequest(ctx, "/dao")
			})

			controlIDNotifications.POST("/door", func(ctx *gin.Context) {
				printRequest(ctx, "/door")
			})

			controlIDNotifications.POST("/secbox", func(ctx *gin.Context) {
				printRequest(ctx, "/secbox")
			})
		}
	}

	// Inicia o servidor na porta 8080
	r.Run(":8080")

}

// Função para imprimir tudo que chega na requisição
func printRequest(ctx *gin.Context, r string) {
	// Pegando o corpo da requisição
	bodyBytes, _ := io.ReadAll(ctx.Request.Body)

	// Recriando o corpo da requisição para evitar perda de dados
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Printando os dados no console
	fmt.Println("===================================" + r + "==========================================")
	println("Headers:", ctx.Request.Header)
	println("Query Params:", ctx.Request.URL.Query())
	println("Body:", string(bodyBytes))

	// Respondendo a requisição com os mesmos dados recebidos
	ctx.JSON(200, gin.H{
		"headers": ctx.Request.Header,
		"query":   ctx.Request.URL.Query(),
		"body":    string(bodyBytes),
	})
}
