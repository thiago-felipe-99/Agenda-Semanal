package main

import "fmt"

type Erro struct {
	Mensagem    string
	ErroInicial *Erro
	ErroExterno error
	Código      string
}

func (erro *Erro) Traçado() string {
	mensagem := fmt.Sprintf("Erro Da Aplicação[%s]: %s", erro.Código, erro.Mensagem)

	if erro.ErroExterno != nil {
		mensagem += "\n\t" + ErroExterno(erro.ErroExterno)
	}

	if erro.ErroInicial != nil {
		mensagem += "\n\t" + erro.ErroInicial.Traçado()
	}

	return mensagem
}

func (erro *Erro) Error() string {
	return erro.Traçado()
}

func ErroExterno(erro error) string {
	return fmt.Sprintf("Erro Externo: %s", erro.Error())
}

type erroPadrão struct {
	Mensagem string
	Código   string
}

func erroNovo(padrão *erroPadrão, inicial *Erro, externo error) *Erro {
	return &Erro{
		Mensagem:    padrão.Mensagem,
		Código:      padrão.Código,
		ErroInicial: inicial,
		ErroExterno: externo,
	}
}
