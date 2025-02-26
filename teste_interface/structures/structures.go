package structures

import (
	"fmt"
	"interface/interfaces"
)



type ControleLG struct{}

type ControleSamsung struct{}

type TV struct{}


type Pessoa struct {
	Nome string
	Idade int
	Status bool
}
func (p Pessoa) String() string{
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos de idade.",p.Nome,p.Idade)
}


type PessoaFisica struct {
	Pessoa
	Cpf string
	Sobrenome string
}
func (pf PessoaFisica) Doc() string{
	return fmt.Sprintf("Meu CPF é %s",pf.Cpf)
}


type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	Cnpj string
}
func (pj PessoaJuridica) Doc() string{
	return fmt.Sprintf("Meu CNPJ é %s",pj.Cnpj)
}



func (t TV) MudarDeCanal(c interfaces.Controle) {
	c.ProximoCanal()
}

func (c ControleLG) ProximoCanal() {
	fmt.Println("Passando para o próximo canal na TV da LG")
}

func (c ControleLG) CanalAnterior() {
	fmt.Println("Passando para o canal anterior na TV da LG")
}

func (c ControleSamsung) ProximoCanal() {
	fmt.Println("Passando para o próximo canal na TV da Samsung")
}

func (c ControleSamsung) CanalAnterior() {
	fmt.Println("Passando para o canal anterior na TV da Samsung")
}


