package utils

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

// Função para imprimir tudo que chega na requisição
func PrintRequest(ctx *gin.Context, r string) {
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
