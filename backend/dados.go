package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErroSalvarAtividadeBD = &erroPadrão{ // nolint:revive
		Mensagem: "Erro ao salvar atividade no banco de dados",
		Código:   "DADOS-[1]",
	}
	ErroAtualizarAtividadeBD = &erroPadrão{
		Mensagem: "Erro ao atualizar atividade no banco de dados",
		Código:   "DADOS-[2]",
	}
	ErroAtividadeNãoEncontradaBD = &erroPadrão{
		Mensagem: "Ativiade não encontrada no banco de dados",
		Código:   "DADOS-[3]",
	}
	ErroPegarAtividadeBD = &erroPadrão{
		Mensagem: "Erro ao pegar atividade no banco de dados",
		Código:   "DADOS-[4]",
	}
	ErroDeletarAtividadeBD = &erroPadrão{
		Mensagem: "Erro ao deletar atividade no banco de dados",
		Código:   "DADOS-[5]",
	}
	ErroPegarAtividadeDiaDB = &erroPadrão{
		Mensagem: "Erro ao pegar a atividade por dia",
		Código:   "DADOS-[5]",
	}
)

// Dados representa um banco de dados na aplicação.
type Dados struct {
	Timeout    time.Duration
	Collection *mongo.Collection
	Log        *Log
}

// SalvarAtividade escreve uma atividade no banco de dados.
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

// AtualizarAtividade é o método que altera uma atividade já existente no banco
// de dados.
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

// PegarAtividade retorna uma atividade salva no banco de dados.
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

// PegarAtividadeDia retorna todas as atividades de um dia do banco de dados.
func (dados *Dados) PegarAtividadeDia(ctx context.Context, dia string) ([]*Atividade, *Erro) {
	dados.Log.Informação("Pegando atividades no banco de dados do:", dia)

	ctx, cancel := context.WithTimeout(ctx, dados.Timeout)
	defer cancel()

	cursor, err := dados.Collection.Find(ctx, bson.M{"dia": dia})
	if err != nil {
		return nil, erroNovo(ErroPegarAtividadeDiaDB, nil, err)
	}

	atividades := []*Atividade{}

	err = cursor.All(ctx, &atividades)
	if err != nil {
		return nil, erroNovo(ErroPegarAtividadeDiaDB, nil, err)
	}

	return atividades, nil
}

// PegarAtividades retorna todas as atividades do banco de dados.
func (dados *Dados) PegarAtividades(ctx context.Context) ([]*Atividade, *Erro) {
	dados.Log.Informação("Pegando todas as atividades do banco de dados")

	ctx, cancel := context.WithTimeout(ctx, dados.Timeout)
	defer cancel()

	cursor, err := dados.Collection.Find(ctx, nil)
	if err != nil {
		return nil, erroNovo(ErroPegarAtividadeDiaDB, nil, err)
	}

	atividades := []*Atividade{}

	err = cursor.All(ctx, &atividades)
	if err != nil {
		return nil, erroNovo(ErroPegarAtividadeDiaDB, nil, err)
	}

	return atividades, nil
}

// Deletar remove uma atividade do banco de dados.
func (dados *Dados) Deletar(ctx context.Context, _id id) *Erro {
	dados.Log.Informação("Deletando atividade no banco de dados com ID:", _id)

	ctx, cancel := context.WithTimeout(ctx, dados.Timeout)
	defer cancel()

	_, err := dados.Collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return erroNovo(ErroDeletarAtividadeBD, nil, err)
	}

	return nil
}
