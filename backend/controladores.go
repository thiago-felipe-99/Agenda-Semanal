package main

import (
	"github.com/gin-gonic/gin"
)

type controlador struct{}

func (controlador *controlador) pegarID(c *gin.Context) {
}

func (controlador *controlador) pegarBodyTarefa(c *gin.Context) {
}

func (controlador *controlador) adicionarTarefa(c *gin.Context) {
}

func (controlador *controlador) atualizarTarefa(c *gin.Context) {
}

func (controlador *controlador) pegarTarefa(c *gin.Context) {
}

func (controlador *controlador) pegarTarefasPorDia(c *gin.Context) {
}

func (controlador *controlador) pegarTarefas(c *gin.Context) {
}

func (controlador *controlador) deletarTarefa(c *gin.Context) {
}

func rotasTarefas(roteamento *gin.RouterGroup, controlador *controlador) {
	roteamento.POST("", controlador.pegarBodyTarefa, controlador.adicionarTarefa)
	roteamento.PUT("/:id", controlador.pegarID, controlador.pegarBodyTarefa, controlador.atualizarTarefa)
	roteamento.GET("/:id", controlador.pegarID, controlador.pegarTarefa)
	roteamento.GET("/dia/:dia", controlador.pegarTarefasPorDia)
	roteamento.GET("", controlador.pegarTarefas)
	roteamento.DELETE("/:id", controlador.pegarID, controlador.deletarTarefa)
}

func rotas(url string, dados *Dados) {
	roteamento := gin.Default()

	rotasTarefas(roteamento.Group("/atividade"), &controlador{})

	if err := roteamento.Run(url); err != nil {
		panic(err)
	}
}
