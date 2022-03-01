package main

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErroConfigurarBD = &erroPadrão{
	Mensagem: "Erro ao configurar o banco de dados",
	Código:   "MAIN-[1]",
}

type id = uuid.UUID

type Atividade struct {
	ID     id            `bson:"_id"`
	Nome   string        `bson:"nome"`
	Dia    string        `bson:"dia"`
	Início time.Duration `bson:"início"`
	Fim    time.Duration `bson:"fim"`
}

func main() {
	uri := ""
	ctx := context.Background()
	nomeDB := ""
	collectionNome := ""

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

	rotas("127.0.0.1:8080", dados)
}
