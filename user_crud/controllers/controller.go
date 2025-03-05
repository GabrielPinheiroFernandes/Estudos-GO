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
			Name:     "Gabriel Pinheiro",
			Username: "GabGamer",
			Pass:     "123",
			ImagePath:    "linkAws.com.br",
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

	var (
		returnUserById structures.User
		id_user int
	)
	id_user=7
	returnUserById, err = c.UserRepository.GetUserByID(id_user)
	if err != nil {
		fmt.Println("Ocorreu algum erro ao tentar obter o usuario:", err)
		// return
	}
	
	if returnUserById.Id == 0 {
		fmt.Printf("usuario %d inexistente \n", id_user)
	} else 
	{
		fmt.Printf("Usuario %v: %v \n\n", id_user,returnUserById)
	}

	//=======================================================
	//============Deletar usuario pelo id====================
	//=======================================================

	deluser, err := c.UserRepository.DelUser(7)
	if err != nil {
		fmt.Printf("Algo deu errado ao tentar apagar o usuario!")
		// return
	}
	fmt.Printf("usuario %d apagado no banco \n\n", deluser)

	//=======================================================
	//====tentando recuperar usuario inexistente pelo id ====
	//=======================================================
	usr_id:=299
	returnUserById, err = c.UserRepository.GetUserByID(usr_id)
	if err != nil {
		fmt.Println("Ocorreu algum erro ao tentar obter o usuario:", err,"\n ")
		// return
	}
	if returnUserById.Id  >= 1 {
		fmt.Println(returnUserById)
	} else {
		fmt.Println("Usuario não existente")
	}
	println("")

}
