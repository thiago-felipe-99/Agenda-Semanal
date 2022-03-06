/*Trabalho Final
Nome: Catarina Oliveira Dowsley Fernandes
DRE: 11802011
Disciplina: Linguagens de Programação - EEL670
Professor: Miguel Elias
Universidade Federal do Rio de Janeiro - 2021.2*/


#include <iostream>
#include <string>
using namespace std;
#include <jsoncpp/json/json.h>
#include <jsoncpp/json/reader.h>
#include <jsoncpp/json/writer.h>
#include <jsoncpp/json/value.h>
#include "menu.h"
#include "requisicao.h"
#include "excecao.h"
#include "agenda.h"

int main() {
    Menu menu;
    unsigned escolha_menu;
	string nomeBusca;

/////////////////// 
Requisicao requisicao = Requisicao();
Agenda agenda = Agenda();
unsigned opcao;
string resposta;
string url;
string atividadeId;
string atividadeDia;
string mensagem;
string questao;
string dia;
string idAtividade;


cout << "\n|---------------------AGENDA-------------------------|\n";
	/*Menu*/
	do {
        
        menu.mostrarMenuPrincipal();
        escolha_menu = menu.receberOpcaoMenu();

		switch (escolha_menu) {

			/*Caso 1 - Adicionar uma nova atividade*/
			case 1:  
				agenda.adicionarAtividade();
				break;

			/*Caso 2 - Deletar uma atividade*/
			case 2:
				cout << "\nEscolha a id da atividade que quer deletar: ";
			 	
					 do {
					 	cin>>opcao;
					 	if (opcao == 0){
							 cout << "Erro - você deve digitar um número maior que 0";
					 	}
					 }while(opcao == 0);
				try{
				requisicao.deletar("https://agenda-semanal.herokuapp.com/atividade/"+to_string(opcao));
				}
				catch(const erroAtividadeoNaoExiste &erro){
					cerr << erro.what();
				}
				break;
			/*Caso 5- - Ver uma atividade a partir do ID*/
			case 5:     
                cout << "\nEscolha a id da atividade que quer ver: ";			    
				cin.ignore();
				getline(cin, idAtividade);
				questao = requisicao.get("https://agenda-semanal.herokuapp.com/atividade/"+idAtividade);
				agenda.mostrarAtividadesId(questao);
            	break;
            
			/*Caso 4 - Ver todas as atividades de um dia escolhido*/
			case 4:   
			   
				cout << "\nDigite o dia da semana desejado: ";
				cin.ignore();
				getline(cin, dia);
				
				questao = requisicao.get("https://agenda-semanal.herokuapp.com/atividade/dia/"+dia);
				agenda.mostrarAtividadesDia(questao);
            	break;
            
			/*Caso 3 - Mostrar todas as atividades*/
            case 3: 

    			questao = requisicao.get("https://agenda-semanal.herokuapp.com/atividade");
				agenda.mostrarAtividades(questao);
                 break;

            /*Caso 6 - Alterar uma atividade*/
            case 6:
				agenda.alterarAtividade();
				break;

			/*Caso 7 - Salvar todas as atividades em um arquivo CSV*/
            case 7:
				questao = requisicao.get("https://agenda-semanal.herokuapp.com/atividade");
    
				agenda.salvarAtividades(questao);
				break;

            /*Sair do Programa*/
			case 8:
				menu.finalizarPrograma();
                escolha_menu = 9;
				break;

            default:
                escolha_menu = 0;
                cout << "Opção Inválida - Tente Novamente" << endl;
            
	    }

    } while (escolha_menu != 9);

return 0;
}
