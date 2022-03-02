package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErroConfigurarBD = &erroPadrão{ // nolint:revive
		Mensagem: "Erro ao configurar o banco de dados",
		Código:   "MAIN-[1]",
	}
	ErroAoValidarID = &erroPadrão{
		Mensagem: "Foi passado um ID inválido",
		Código:   "MAIN-[2]",
	}
)

type id = uuid.UUID

// Atividade representa a entidade atividade da aplicação.
type Atividade struct {
	ID     id            `bson:"_id" json:"id" validate:"required"`
	Nome   string        `bson:"nome" json:"nome" validate:"required"`
	Dia    string        `bson:"dia" json:"dia" validate:"required"`
	Início time.Duration `bson:"início" json:"início" validate:"required"`
	Fim    time.Duration `bson:"fim" json:"fim" validate:"required"`
}

// ParseID retorna um ID a partir de uma string válida.
func ParseID(parse string) (id, *Erro) {
	id, err := uuid.Parse(parse)
	if err != nil {
		return id, erroNovo(ErroAoValidarID, nil, err)
	}

	return id, nil
}

// VariáveisDeAmbiente representa as váriveis de ambiente que a aplicação
// precisa.
type VariáveisDeAmbiente struct {
	MongoDBURI     string
	NomeDB         string
	NomeCollection string
	Port           string
	Host           string
}

// PegandoVariáveisDeAmbiente retorna as variáveis do ambiente.
func PegandoVariáveisDeAmbiente() (variáveis VariáveisDeAmbiente) {
	variáveis.MongoDBURI = os.Getenv("MONGO_DB_URI")
	if variáveis.MongoDBURI == "" {
		variáveis.MongoDBURI = "mongodb://root:root@localhost:2002"
		log.Println("A variável de ambiente MONGO_DB_URI não foi inicializada")
	}
	variáveis.NomeDB = os.Getenv("NOME_DB")
	if variáveis.NomeDB == "" {
		variáveis.NomeDB = "atividade"
		log.Println("A variável de ambiente NOME_DB não foi inicializada")
	}
	variáveis.NomeCollection = os.Getenv("NOME_COLLECTION")
	if variáveis.NomeCollection == "" {
		variáveis.NomeCollection = "atividade"
		log.Println("A variável de ambiente NOME_COLLECTION não foi inicializada")
	}
	variáveis.Port = os.Getenv("PORT")
	if variáveis.Port == "" {
		variáveis.Port = "2001"
		log.Println("A variável de ambiente PORT não foi inicializada")
	}
	variáveis.Host = os.Getenv("HOST_HTTP")
	if variáveis.Host == "" {
		log.Println("A variável de ambiente HOST_HTTP não foi inicializada")
	}

	return variáveis
}

func main() {
	log.Println("Iniciando servidor")
	ambiente := PegandoVariáveisDeAmbiente()
	ctx := context.Background()

	mongoBD, err := mongo.Connect(ctx, options.Client().ApplyURI(ambiente.MongoDBURI))
	if err != nil {
		panic(erroNovo(ErroConfigurarBD, nil, err))
	}

	const maxTimeout = 3 * time.Second

	dados := &Dados{
		Timeout:    maxTimeout,
		Log:        NovoLog(os.Stdout, NívelDebug),
		Collection: mongoBD.Database(ambiente.NomeDB).Collection(ambiente.NomeCollection),
	}

	ctx2, cancel := context.WithTimeout(ctx, maxTimeout)
	defer cancel()

	err = mongoBD.Ping(ctx2, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Banco de dados conectado")

	rotas(fmt.Sprintf("%s:%s", ambiente.Host, ambiente.Port), dados)
}
