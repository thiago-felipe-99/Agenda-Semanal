package main

type Erro struct {
	Mensagem    string
	ErroInicial *Erro
	ErroExterno error
	Código      string
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
