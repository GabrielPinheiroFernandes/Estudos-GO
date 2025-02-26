package controllers

import (
	"fmt"
	"userCrud/interfaces"
	"userCrud/structures"
)

type Controller struct {
	UserRepository interfaces.UserRepository
}

func NewController(repo interfaces.UserRepository) *Controller {
	return &Controller{
		UserRepository: repo,
	}
}

func (c *Controller) Run() {
	var (
		user structures.User
	)
	//=======================================================
	//============LOOP DE ADICIONAR USUARIOS=================
	//=======================================================
	ids := []int{}
	for i := 0; i < 10; i++ {
		user = structures.User{
			Id:       i + 1,
			Name:     "Gabriel Pinheiro",
			Username: "GabGamer",
			Pass:     "123",
			Image:    "linkAws.com.br",
		}
		data, err := c.UserRepository.AddUser(user)
		if err != nil {
			fmt.Println("Erro ao adicionar usuario:", err)
		}
		ids = append(ids, data)
	}
	fmt.Printf("Sucesso adicionando os usuarios %d \n\n", ids)

	//=======================================================
	//============Recuperar todos os usuarios================
	//=======================================================

	data, err := c.UserRepository.GetAllUsers()
	if err != nil {
		fmt.Println("Ocorreu algum erro ao tentar obter todos os usuários:", err)
		// return
	}

	for _, user := range data {
		fmt.Println(user) // Aqui, você imprime cada usuário
	}
	println("")

	//=======================================================
	//============Recuperar um unico usuairo pelo id=========
	//=======================================================

	var returnUserById structures.User
	returnUserById, err = c.UserRepository.GetUserByID(7)
	if err != nil {
		fmt.Println("Ocorreu algum erro ao tentar obter o usuario:", err)
		// return
	}
	fmt.Printf("Usuario: %v \n\n", returnUserById)

	//=======================================================
	//============Deletar usuario pelo id====================
	//=======================================================

	deluser, err := c.UserRepository.DelUser(7)
	if err != nil {
		fmt.Printf("usuario %d inexistente \n", deluser)
		// return
	}
	fmt.Printf("usuario %d deletado com sucesso! \n\n", deluser)

	//=======================================================
	//============Deletar usuario inexistente pelo id========
	//=======================================================

	deluser, err = c.UserRepository.DelUser(7)
	if err != nil {
		fmt.Printf("usuario %d inexistente \n\n", deluser)
		// return
	}
	fmt.Printf("usuario %d deletado com sucesso! \n\n", deluser)

	//=======================================================
	//===Recuperar todos os usuarios com usuario deletado====
	//=======================================================

	data, err = c.UserRepository.GetAllUsers()
	if err != nil {
		fmt.Println("Ocorreu algum erro ao tentar obter todos os usuários:", err)
		// return
	}

	for _, user := range data {
		fmt.Println(user) // Aqui, você imprime cada usuário
	}
	println("")

	//=======================================================
	//====tentando recuperar usuario inexistente pelo id ====
	//=======================================================
	returnUserById, err = c.UserRepository.GetUserByID(7)
	if err != nil {
		fmt.Println("Ocorreu algum erro ao tentar obter o usuario:", err,"\n ")
		// return
	} else {
		fmt.Println(returnUserById)
	}
	println("")

}
