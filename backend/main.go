package main

import (
	"context"
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

func main() {
	uri := "mongodb://root:root@localhost:2001"
	ctx := context.Background()
	nomeDB := "atividades"
	collectionNome := "atividades"

	mongoBD, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(erroNovo(ErroConfigurarBD, nil, err))
	}

	const maxTimeout = 3 * time.Second

	dados := &Dados{
		Timeout:    maxTimeout,
		Log:        NovoLog(os.Stdout, NívelDebug),
		Collection: mongoBD.Database(nomeDB).Collection(collectionNome),
	}

	err = mongoBD.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	rotas("localhost:2000", dados)
}
