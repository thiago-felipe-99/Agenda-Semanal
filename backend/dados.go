package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErroSalvarAtividadeBD = &erroPadrão{
	Mensagem: "Erro ao salvar atividade no banco de dados",
	Código:   "DADOS-[1]",
}

var ErroAtualizarAtividadeBD = &erroPadrão{
	Mensagem: "Erro ao atualizar atividade no banco de dados",
	Código:   "DADOS-[2]",
}

var ErroAtividadeNãoEncontradaBD = &erroPadrão{
	Mensagem: "Ativiade não encontrada no banco de dados",
	Código:   "DADOS-[3]",
}

var ErroPegarAtividadeBD = &erroPadrão{
	Mensagem: "Erro ao pegar atividade no banco de dados",
	Código:   "DADOS-[4]",
}

var ErroDeletarAtividade = &erroPadrão{
	Mensagem: "Erro ao deletat atividade no banco de dados",
	Código:   "DADOS-[5]",
}

type Dados struct {
	Timeout    time.Duration
	Collection *mongo.Collection
	Log        *Log
}

func (dados *Dados) SalvarAtividade(ctx context.Context, atividade *Atividade) *Erro {
	dados.Log.Informação("Salvando atividade no banco de dados a atividade com ID:", atividade.ID)

	ctx, cancel := context.WithTimeout(ctx, dados.Timeout)
	defer cancel()

	_, err := dados.Collection.InsertOne(ctx, atividade)
	if err != nil {
		return erroNovo(ErroSalvarAtividadeBD, nil, err)
	}

	return nil
}

func (dados *Dados) AtualizarAtividade(ctx context.Context, _id id, atividade *Atividade) *Erro {
	dados.Log.Informação("Atualizando atividade no banco de dados a atividade com ID:", atividade.ID)

	ctx, cancel := context.WithTimeout(ctx, dados.Timeout)
	defer cancel()

	query := bson.D{{Key: "$set", Value: atividade}}

	_, err := dados.Collection.UpdateByID(ctx, _id, query)
	if err != nil {
		return erroNovo(ErroAtualizarAtividadeBD, nil, err)
	}

	return nil
}

func (dados *Dados) PegarAtividade(ctx context.Context, _id id) (*Atividade, *Erro) {
	dados.Log.Informação("Pegando atividade no banco de dados com ID:", _id)

	ctx, cancel := context.WithTimeout(ctx, dados.Timeout)
	defer cancel()

	var atividade Atividade

	err := dados.Collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&atividade)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, erroNovo(ErroAtividadeNãoEncontradaBD, nil, err)
		}

		return nil, erroNovo(ErroPegarAtividadeBD, nil, err)
	}

	return &atividade, nil
}

func (dados *Dados) Deletar(ctx context.Context, _id id) *Erro {
	dados.Log.Informação("Deletando atividade no banco de dados com ID:", _id)

	ctx, cancel := context.WithTimeout(ctx, dados.Timeout)
	defer cancel()

	_, err := dados.Collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return erroNovo(ErroDeletarAtividade, nil, err)
	}

	return nil
}
