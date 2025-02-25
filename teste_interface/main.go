package main

import (
	"fmt"
	"interface/interfaces"
	"interface/structures"
)

func Show(d interfaces.Document) {
	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	p := structures.Pessoa{
		Nome:   "Roberto",
		Idade:  22,
		Status: true,
	}

	pf := structures.PessoaFisica{
		Pessoa:    p,
		Cpf:       "000.000.000-00",
		Sobrenome: "Nunes",
	}

	pj := structures.PessoaJuridica{
		Pessoa:      p,
		RazaoSocial: "Roberto Nunes LTDA",
		Cnpj:        "00.000.000/0000-00",
	}

	tv := structures.TV{}

	lg := structures.ControleLG{}
	s := structures.ControleSamsung{}

	tv.MudarDeCanal(lg)
	tv.MudarDeCanal(s)
	Show(pj)
	Show(pf)
}
