package main

import (
	"fmt"
	s "teste/structs"
)

func main() {
	// Chamando a função Gerapessoa do pacote structs
	p1 := s.Gerapessoa("João", "Silva", 30, false)
	p2 := s.Gerapessoa("Gabriel", "Fernandes", 21, true)

	var statusP string
	if p1.Status {
		statusP = "Vivo"
	} else {
		statusP = "Morto"
	}
	fmt.Printf("O nome do usuário é %s %s, tem %d anos e ele está %s.\n", p1.Nome, p1.Sobrenome, p1.Idade, statusP)

	if p2.Status {
		statusP = "Vivo"
	} else {
		statusP = "Morto"
	}
	fmt.Printf("O nome do usuário é %s %s, tem %d anos e ele está %s.\n", p2.Nome, p2.Sobrenome, p2.Idade, statusP)
	mudaNome(&p1)
	fmt.Println(p1.Nome)
}

func mudaNome(p *s.Pessoa) {
	p.Nome = "novo nome"
}
