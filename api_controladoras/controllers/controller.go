package controllers

import (
	"APIControlID/interfaces"
	"APIControlID/structs"
	"APIControlID/utils"
	"bytes"
	"encoding/json"
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

		grpRoutesUser.PUT("/edit/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			// Convertendo ID para inteiro
			iId, err := strconv.Atoi(id)
			if err != nil {
				ctx.JSON(400, gin.H{"message": "Forneça um ID válido!", "erro": err.Error()})
				return
			}

			// Lendo o JSON do corpo da requisição
			var user structs.User
			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(400, gin.H{"message": "JSON inválido", "error": err.Error()})
				return
			}

			// Chamando a função EditUser
			data, err := c.crudApi.EditUser(iId, user)
			if err != nil {
				ctx.JSON(500, gin.H{"message": "Falha ao editar usuário!", "erro": err.Error()})
				return
			}

			// Retornando a struct User como JSON
			ctx.JSON(200, gin.H{
				"id":   iId,
				"user": string(data), // Aqui o `data` é uma struct User que será automaticamente convertida para JSON
			})
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
				utils.PrintRequest(ctx, "/dao")
			})

			controlIDNotifications.POST("/door", func(ctx *gin.Context) {
				utils.PrintRequest(ctx, "/door")
			})

			controlIDNotifications.POST("/secbox", func(ctx *gin.Context) {
				utils.PrintRequest(ctx, "/secbox")
			})
		}
	}

	// Inicia o servidor na porta 8080
	r.Run(":8080")

}
