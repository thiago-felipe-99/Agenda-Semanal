package main

import (
	"time"

	"github.com/google/uuid"
)

type id = uuid.UUID

type Atividade struct {
	ID     id            `bson:"_id"`
	Nome   string        `bson:"nome"`
	Dia    string        `bson:"dia"`
	Início time.Duration `bson:"início"`
	Fim    time.Duration `bson:"fim"`
}

func main() {
	rotas("127.0.0.1:8080", &Dados{})
}
