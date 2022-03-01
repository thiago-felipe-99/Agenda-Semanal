package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_tradução "github.com/go-playground/validator/v10/translations/en"
	es_tradução "github.com/go-playground/validator/v10/translations/es"
	pt_tradução "github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/google/uuid"
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
	ErroDiaInvalido = &erroPadrão{
		Mensagem: "Foi passado um dia inválido",
		Código:   "CONTROLADORES-[3]",
	}
	ErroRequisiçãoSemBody = &erroPadrão{
		Mensagem: "Não foi passado body na requisição",
		Código:   "CONTROLADORES-[4]",
	}
	ErroDecodificarJSON = &erroPadrão{
		Mensagem: "Foi passado um JSON inválido",
		Código:   "CONTROLADORES-[5]",
	}
	ErroValidarBody = &erroPadrão{
		Mensagem: "Não foi possível validar o body",
		Código:   "CONTROLADORES-[6]",
	}
	ErroFimMenorInício = &erroPadrão{
		Mensagem: "A data final é menor que a data inicial",
		Código:   "CONTROLADORES-[7]",
	}
	ErroTempoInválido = &erroPadrão{
		Mensagem: "O tempo não pode exceder 24h",
		Código:   "CONTROLADORES-[8]",
	}
	ErroAtividadeNãoExisteNoContexto = &erroPadrão{
		Mensagem: "A atividade não existe no contexto",
		Código:   "CONTROLADORES-[9]",
	}
	ErroConverterAtividadeDoContexto = &erroPadrão{
		Mensagem: "Erro ao converte a atividade do contexto",
		Código:   "CONTROLADORES-[10]",
	}
)

type mensagemJSON struct {
	Mensagem  string
	Erros     []string
	Atividade []Atividade
}

type controlador struct {
	Log       *Log
	validator *validator.Validate
}

var uni *ut.UniversalTranslator //nolint: gochecknoglobals

func pegarTradutor(c *gin.Context) *ut.Translator {
	trans, existe := uni.GetTranslator(c.Request.Header.Get("Accept-Language"))
	if !existe {
		trans, _ = uni.GetTranslator("pt_BR")
	}

	return &trans
}

func (controlador *controlador) enviarErro(ginC *gin.Context, erro *Erro) {
	var (
		código   int
		mensagem string
	)

	switch erro.Código {
	case ErroAoValidarID.Código, ErroDiaInvalido.Código,
		ErroRequisiçãoSemBody.Código, ErroDecodificarJSON.Código,
		ErroFimMenorInício.Código, ErroTempoInválido.Código:
		{
			código = http.StatusBadRequest
			mensagem = erro.Mensagem
		}
	default:
		código = http.StatusInternalServerError
		mensagem = "Ocoreu um erro inesperado"

		controlador.Log.Erro(erro.Traçado())
	}

	ginC.JSON(código, mensagemJSON{
		Mensagem:  mensagem,
		Erros:     []string{mensagem},
		Atividade: nil,
	})
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

//nolint: funlen
func (controlador *controlador) pegarBodyTarefa(ginC *gin.Context) {
	decodificador := json.NewDecoder(ginC.Request.Body)
	atividadeJSON := struct {
		Nome   string        `json:"nome" validate:"required"`
		Dia    string        `json:"dia" validate:"required"`
		Início time.Duration `json:"início" validate:"required"`
		Fim    time.Duration `json:"fim" validate:"required"`
	}{}

	err := decodificador.Decode(&atividadeJSON)
	if err != nil {
		if errors.Is(err, io.EOF) {
			controlador.enviarErro(ginC, erroNovo(ErroRequisiçãoSemBody, nil, err))

			return
		}

		controlador.enviarErro(ginC, erroNovo(ErroDecodificarJSON, nil, err))

		return
	}

	err = controlador.validator.Struct(atividadeJSON)
	if err != nil {
		if erros, ok := err.(validator.ValidationErrors); ok { //nolint: errorlint
			mensagens := []string{}

			tradutor := pegarTradutor(ginC)
			if tradutor == nil {
				controlador.enviarErro(ginC, erroNovo(ErroValidarBody, nil, nil))

				return
			}

			for _, erro := range erros.Translate(*tradutor) {
				mensagens = append(mensagens, erro)
			}

			ginC.JSON(http.StatusBadRequest, mensagemJSON{
				Mensagem:  "Foi passado valores inválidos no body",
				Erros:     mensagens,
				Atividade: nil,
			})
			ginC.Abort()

			return
		}

		controlador.enviarErro(ginC, erroNovo(ErroValidarBody, nil, nil))

		return
	}

	if !diaValido(atividadeJSON.Dia) {
		controlador.enviarErro(ginC, erroNovo(ErroDiaInvalido, nil, nil))

		return
	}

	if atividadeJSON.Fim < atividadeJSON.Início {
		controlador.enviarErro(ginC, erroNovo(ErroFimMenorInício, nil, nil))

		return
	}

	const dia = 24 * time.Hour

	if atividadeJSON.Fim >= dia {
		controlador.enviarErro(ginC, erroNovo(ErroTempoInválido, nil, nil))

		return
	}

	atividade := &Atividade{ //nolint: exhaustivestruct
		Nome:   atividadeJSON.Nome,
		Dia:    atividadeJSON.Dia,
		Início: atividadeJSON.Início,
		Fim:    atividadeJSON.Fim,
	}

	ginC.Set("atividade", atividade)
	ginC.Next()
}

func (controlador *controlador) pegarAtividadeContexto(c *gin.Context) (*Atividade, *Erro) {
	pessoaGet, existe := c.Get("atividade")
	if !existe {
		return nil, erroNovo(ErroAtividadeNãoExisteNoContexto, nil, nil)
	}

	pessoa, okay := pessoaGet.(*Atividade)
	if !okay {
		return nil, erroNovo(ErroConverterAtividadeDoContexto, nil, nil)
	}

	return pessoa, nil
}

func (controlador *controlador) adicionarAtividade(ginC *gin.Context) {
	atividade, erro := controlador.pegarAtividadeContexto(ginC)
	if erro != nil {
		controlador.enviarErro(ginC, erro)

		return
	}

	_id := uuid.New()
	atividade.ID = _id

	mensagem := fmt.Sprintf("Tarefa com ID %s adicionada com sucesso", _id)

	ginC.JSON(http.StatusCreated, mensagemJSON{
		Mensagem:  mensagem,
		Erros:     nil,
		Atividade: []Atividade{*atividade},
	})
}

func (controlador *controlador) atualizarAtividade(ginC *gin.Context) {
	_id, erro := controlador.pegarIDContexto(ginC)
	if erro != nil {
		controlador.enviarErro(ginC, erro)

		return
	}

	atividade, erro := controlador.pegarAtividadeContexto(ginC)
	if erro != nil {
		controlador.enviarErro(ginC, erro)

		return
	}

	atividade.ID = *_id

	mensagem := fmt.Sprintf("Tarefa com ID %s atualizada com sucesso", _id)

	ginC.JSON(http.StatusOK, mensagemJSON{
		Mensagem:  mensagem,
		Erros:     nil,
		Atividade: []Atividade{*atividade},
	})
}

func (controlador *controlador) pegarTarefa(ginC *gin.Context) {
	_id, erro := controlador.pegarIDContexto(ginC)
	if erro != nil {
		controlador.enviarErro(ginC, erro)

		return
	}

	mensagem := fmt.Sprintf("Tarefa com ID %s econtrada com sucesso", _id)

	ginC.JSON(http.StatusOK, mensagemJSON{
		Mensagem:  mensagem,
		Erros:     nil,
		Atividade: nil,
	})
}

func diaValido(dia string) bool {
	dias := []string{"domindo", "segunda", "terça", "quarta", "quinta", "sexta", "sábado"}

	for _, valido := range dias {
		if valido == dia {
			return true
		}
	}

	return false
}

func (controlador *controlador) pegarTarefasPorDia(ginC *gin.Context) {
	dia := ginC.Params.ByName("dia")
	if !diaValido(dia) {
		controlador.enviarErro(ginC, erroNovo(ErroDiaInvalido, nil, nil))

		return
	}

	mensagem := fmt.Sprintf("Tarefas do dia %s econtradas com sucesso", dia)

	ginC.JSON(http.StatusOK, mensagemJSON{
		Mensagem:  mensagem,
		Erros:     nil,
		Atividade: nil,
	})
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

	ginC.JSON(http.StatusOK, mensagemJSON{
		Mensagem:  mensagem,
		Erros:     nil,
		Atividade: nil,
	})
}

func rotasTarefas(roteamento *gin.RouterGroup, controlador *controlador) {
	roteamento.POST("", controlador.pegarBodyTarefa, controlador.adicionarAtividade)
	roteamento.PUT("/:id", controlador.pegarID, controlador.pegarBodyTarefa, controlador.atualizarAtividade)
	roteamento.GET("/:id", controlador.pegarID, controlador.pegarTarefa)
	roteamento.GET("/dia/:dia", controlador.pegarTarefasPorDia)
	roteamento.GET("", controlador.pegarTarefas)
	roteamento.DELETE("/:id", controlador.pegarID, controlador.deletarTarefa)
}

func rotas(url string, _ *Dados) {
	roteamento := gin.Default()

	validate := validator.New()

	uni = ut.New(pt_BR.New(), en.New(), es.New())

	trans, _ := uni.GetTranslator("pt_BR")

	err := pt_tradução.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	trans, _ = uni.GetTranslator("en")

	err = en_tradução.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	trans, _ = uni.GetTranslator("es")

	err = es_tradução.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	rotasTarefas(roteamento.Group("/atividade"), &controlador{
		Log:       NovoLog(os.Stdout, NívelDebug),
		validator: validate,
	})

	if err := roteamento.Run(url); err != nil {
		panic(err)
	}
}
