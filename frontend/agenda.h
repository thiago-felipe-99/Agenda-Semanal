#ifndef AGENDA_H
#define AGENDA_H

#include <jsoncpp/json/json.h>
#include <jsoncpp/json/reader.h>
#include <jsoncpp/json/writer.h>
#include <jsoncpp/json/value.h>
#include "requisicao.h"
#include "menu.h"


class Agenda{

    public:
        Agenda(){};

        /*Declaração de todas as classes a serem utilizada*/
        void deletarQuestao();
        void adicionarAtividade();
        void mostrarAtividades(string questao);
        void alterarAtividade();
        void salvarAtividades(string questao);
        void mostrarAtividadesDia(string questao);
        void mostrarAtividadesId(string questao);
    private:
};
#endif