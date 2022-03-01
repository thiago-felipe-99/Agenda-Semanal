package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	ErroIDNãoExisteNoContexto = &erroPadrão{
		Mensagem: "O ID não existe no contexto",
		Código:   "CONTROLADORES-[1]",
	}
	ErroConverterIDDoContexto = &erroPadrão{
		Mensagem: "Erro ao converte o ID do contexto",
		Código:   "CONTROLADORES-[2]",
	}
)

type mensagemJSON struct {
	Mensagem  string
	Erro      []string
	Atividade *Atividade
}

type controlador struct {
	Log *Log
}

func (controlador *controlador) enviarErro(ginC *gin.Context, erro *Erro) {
	var (
		código   int
		mensagem string
	)

	switch erro.Código {
	case ErroAoValidarID.Código:
		código = http.StatusBadRequest
		mensagem = erro.Mensagem
	default:
		código = http.StatusInternalServerError
		mensagem = "Ocoreu um erro inesperado"

		controlador.Log.Erro(erro.Traçado())
	}

	ginC.JSON(código, gin.H{"erro": mensagem})
	ginC.Abort()
}

func (controlador *controlador) pegarID(ginC *gin.Context) {
	_id, erro := ParseID(ginC.Params.ByName("id"))
	if erro != nil {
		controlador.enviarErro(ginC, erro)

		return
	}

	ginC.Set("id", &_id)
	ginC.Next()
}

func (controlador *controlador) pegarIDContexto(ginC *gin.Context) (*id, *Erro) {
	IDGet, existe := ginC.Get("id")
	if !existe {
		return nil, erroNovo(ErroIDNãoExisteNoContexto, nil, nil)
	}

	id, okay := IDGet.(*id)
	if !okay {
		return nil, erroNovo(ErroConverterIDDoContexto, nil, nil)
	}

	return id, nil
}

func (controlador *controlador) pegarBodyTarefa(ginC *gin.Context) {
}

func (controlador *controlador) adicionarTarefa(ginC *gin.Context) {
}

func (controlador *controlador) atualizarTarefa(ginC *gin.Context) {
	_id, erro := controlador.pegarIDContexto(ginC)
	if erro != nil {
		controlador.enviarErro(ginC, erro)

		return
	}

	mensagem := fmt.Sprintf("Tarefa com ID %s atualizada com sucesso", _id)

	ginC.JSON(http.StatusOK, mensagemJSON{Mensagem: mensagem, Erro: nil, Atividade: nil})
}

func (controlador *controlador) pegarTarefa(ginC *gin.Context) {
	_id, erro := controlador.pegarIDContexto(ginC)
	if erro != nil {
		controlador.enviarErro(ginC, erro)

		return
	}

	mensagem := fmt.Sprintf("Tarefa com ID %s pega com sucesso", _id)

	ginC.JSON(http.StatusOK, mensagemJSON{Mensagem: mensagem, Erro: nil, Atividade: nil})
}

func (controlador *controlador) pegarTarefasPorDia(ginC *gin.Context) {
}

func (controlador *controlador) pegarTarefas(ginC *gin.Context) {
}

func (controlador *controlador) deletarTarefa(ginC *gin.Context) {
	_id, erro := controlador.pegarIDContexto(ginC)
	if erro != nil {
		controlador.enviarErro(ginC, erro)

		return
	}

	mensagem := fmt.Sprintf("Tarefa com ID %s deletada com sucesso", _id)

	ginC.JSON(http.StatusOK, mensagemJSON{Mensagem: mensagem, Erro: nil, Atividade: nil})
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

	rotasTarefas(roteamento.Group("/atividade"), &controlador{
		Log: NovoLog(os.Stdout, NívelDebug),
	})

	if err := roteamento.Run(url); err != nil {
		panic(err)
	}
}
