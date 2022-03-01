package main

import (
	"io"
	"log"
)

var ErroNívelInválido = &erroPadrão{
	Mensagem: "O nível escolhido é inválido",
	Código:   "LOGS-[1]",
}

const (
	NívelPanic uint = iota
	NívelErro
	NívelAviso
	NívelInfo
	NívelDebug
)

type Log struct {
	outPanic      *log.Logger
	outErro       *log.Logger
	outAviso      *log.Logger
	outInformação *log.Logger
	outDebug      *log.Logger
	Nível         uint
}

func (log *Log) Panic(imprimir ...interface{}) {
	log.outPanic.Println(imprimir...)
	panic(imprimir)
}

func (log *Log) Erro(imprimir ...interface{}) {
	if log.Nível < NívelErro {
		return
	}

	log.outErro.Println(imprimir...)
}

func (log *Log) Aviso(imprimir ...interface{}) {
	if log.Nível < NívelAviso {
		return
	}

	log.outAviso.Println(imprimir...)
}

func (log *Log) Informação(imprimir ...interface{}) {
	if log.Nível < NívelInfo {
		return
	}

	log.outInformação.Println(imprimir...)
}

func (log *Log) Debug(imprimir ...interface{}) {
	if log.Nível < NívelDebug {
		return
	}

	log.outDebug.Println(imprimir...)
}

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
