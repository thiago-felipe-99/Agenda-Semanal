package main

import (
	"io"
	"log"
)

var ErroNívelInválido = &erroPadrão{ // nolint:revive
	Mensagem: "O nível escolhido é inválido",
	Código:   "LOGS-[1]",
}

// Possíveis níveis que o log pod mostrar.
const (
	NívelPanic uint = iota
	NívelErro
	NívelAviso
	NívelInfo
	NívelDebug
)

// Log representa um log na aplicação.
type Log struct {
	outPanic      *log.Logger
	outErro       *log.Logger
	outAviso      *log.Logger
	outInformação *log.Logger
	outDebug      *log.Logger
	Nível         uint
}

// Panic imprime o nível panic no log.
func (log *Log) Panic(imprimir ...interface{}) {
	log.outPanic.Println(imprimir...)
	panic(imprimir)
}

// Erro imprime o nível erro no log.
func (log *Log) Erro(imprimir ...interface{}) {
	if log.Nível < NívelErro {
		return
	}

	log.outErro.Println(imprimir...)
}

// Aviso imprime o nível aviso no log.
func (log *Log) Aviso(imprimir ...interface{}) {
	if log.Nível < NívelAviso {
		return
	}

	log.outAviso.Println(imprimir...)
}

// Informação imprime o nível Informação no log.
func (log *Log) Informação(imprimir ...interface{}) {
	if log.Nível < NívelInfo {
		return
	}

	log.outInformação.Println(imprimir...)
}

// Debug imprime o nível Debug no log.
func (log *Log) Debug(imprimir ...interface{}) {
	if log.Nível < NívelDebug {
		return
	}

	log.outDebug.Println(imprimir...)
}

// NovoLog criar um novo log da aplicação.
func NovoLog(out io.Writer, nível uint) *Log {
	if nível < NívelPanic || nível > NívelDebug {
		panic(erroNovo(ErroNívelInválido, nil, nil))
	}

	return &Log{
		outPanic:      log.New(out, "PANIC - ", log.Ldate|log.Ltime|log.Lshortfile),
		outErro:       log.New(out, "ERRO - ", log.Ldate|log.Ltime|log.Lshortfile),
		outAviso:      log.New(out, "AVISO - ", log.Ldate|log.Ltime|log.Lshortfile),
		outInformação: log.New(out, "INFORMAÇÃO - ", log.Ldate|log.Ltime|log.Lshortfile),
		outDebug:      log.New(out, "DEBUG - ", log.Ldate|log.Ltime|log.Lshortfile),
		Nível:         nível,
	}
}
