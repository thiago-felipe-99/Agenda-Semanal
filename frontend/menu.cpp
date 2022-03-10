using namespace std;
#include "menu.h"
#include "agenda.h"
#include "requisicao.h"



/*Função que mostra o menu para o usuário*/
void Menu::mostrarMenuPrincipal() {

    cout << "\n|----------------------------------------------------|"
         << "\n|                   MENU PRINCIPAL                   |"
         << "\n|----------------------------------------------------|"
         << "\n| 1 - Adicionar Atividade                            |"
         << "\n| 2 - Deletar Atividade                              |"
         << "\n| 3 - Ver todas as atividades                        |"
         << "\n| 4 - Ver atividades de um dia específico            |"
         << "\n| 5 - Ver atividade a partir de do ID                |"
         << "\n| 6 - Alterar Atividade                              |"
         << "\n| 7 - Salvar atividades em um arquivo csv            |"
         << "\n| 8 - Sair                                           |"
         << "\n|----------------------------------------------------|\n"
      
            ;   
};  


/*Função que recebe a opção do usuário*/
unsigned Menu::receberOpcaoMenu(){
    char entrada_usuario[256];
    unsigned escolhaMenu;
    
    cout << "\nDigite a opção desejada:";
    cin  >> entrada_usuario;
    cout<<"\n"<<endl;

    escolhaMenu = atoi(entrada_usuario);
    
    return escolhaMenu;
};
unsigned Menu::receberOpcaoMenuDia(){
    unsigned dia;

    cout << "\n Digite o dia da semana desejado:";
    cin  >> dia;
    
    return dia;
};
/*Função que finaliza o programa*/
void Menu::finalizarPrograma(){
    cout << "\n|----------------FIM DO PROGRAMA-----------------|\n\n";
};