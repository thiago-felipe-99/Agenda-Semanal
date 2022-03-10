#include "agenda.h"
#include "requisicao.h"
#include <iostream>
#include <stdio.h>
#include <fstream>
#include <string>
#include <string.h>
#include <cstring>
#include <algorithm>

using namespace std;

/*Definição das funções a serem utilizada */
void Agenda::adicionarAtividade(){
    string nomeAtividade;
    string diaAtividade;
    string horaInicio;
    string horaFim;
    string body;
    Requisicao requisicao = Requisicao();
 
    cout << "\n-----Nova atividade-----\n"<<endl;
    
    /* Recebe os valores que vamos mandar na requisição */
    cout << "Digite o nome da atividade: ";
    cin.ignore();
    getline(cin, nomeAtividade);
     
    cout << "Digite o dia da semana da atividade: ";
    getline(cin, diaAtividade);

    cout << "Digite o horario de inicio da atividade: ";
    getline(cin, horaInicio);

    cout << "Digite o horario do fim da atividade: ";
    getline(cin, horaFim);  
   
   /* Coloca os valores em um corpo*/
    body = "{ \"nome\": \""+nomeAtividade+"\", \"dia\": \""+diaAtividade+"\", \"início\": "+horaInicio+", \"fim\": "+horaFim+"}";
    /* Envia para a requisição */
    requisicao.post("https://agenda-semanal.herokuapp.com/atividade",body);
        
}

void Agenda::alterarAtividade(){
    string nomeAtividade;
    string diaAtividade;
    string horaInicio;
    string horaFim;
    string idAtividade;
    string body;
    Requisicao requisicao = Requisicao();
    

    cout << "\n|-------Alterando Atividade------\n"<<endl;
    
    /* Recebe a id da atividade a ser alterada*/
    cout << "Digite o id da atividade que quer alterar: ";
    cin.ignore();
    getline(cin, idAtividade);     
    /*Recebe os novos dados */
    cout << "Digite o novo nome da atividade: ";
    getline(cin, nomeAtividade);
     
    cout << "Digite o novo  dia da semana da atividade: ";
    getline(cin, diaAtividade);

    cout << "Digite o novo  horario de inicio da atividade: ";
    getline(cin, horaInicio);

    cout << "Digite o novo horario do fim da atividade: ";
    getline(cin, horaFim);  
    
    /* Coloca os dados em um corpo*/
    body = "{ \"nome\": \""+nomeAtividade+"\", \"dia\": \""+diaAtividade+"\", \"início\": "+horaInicio+", \"fim\": "+horaFim+"}";
    /* Envia para a requisição*/
    requisicao.put("https://agenda-semanal.herokuapp.com/atividade/"+idAtividade,body);
    
}



void Agenda::mostrarAtividades(string questao){

    string atividades;
    Json::Value root;
    Json::Reader reader;
    int tamanho;
    Json::Value guardar;
    bool parsingSuccessful = reader.parse(questao.c_str(), root);

    if (!parsingSuccessful){
        throw erroAtividadeoNaoExiste();
    }

    //o guardar é um json array- com os objetos de cada atividade.
    guardar = root.get("atividades","");
    
    int i;
    i = 0;
   
    tamanho = guardar.size();
 
    /* Passa por todos os objetos do array "atividades" e os printa na tela*/
    cout<<"--------------Atividades-------------------"<<endl;
    while (i < tamanho){
        
        cout<<"Nome da Atividade: "<<guardar[i]["nome"]<<endl;
        cout<<"Dia da semana: "<<guardar[i]["dia"]<<endl;
        cout<<"Horário de início: "<<guardar[i]["início"]<<"h"<<endl;
        cout<<"Horário de término: "<<guardar[i]["fim"]<<"h"<<endl;
        cout<<"----------------------------------------------"<<endl;
        i++;
    }

}
 
void Agenda::mostrarAtividadesDia(string questao){
    
    Json::FastWriter fastwriter;
    Json::Value root;
    Json::Reader reader;
    Json::Value guardar;
    int tamanho;
    string nome;

    bool parsingSuccessful = reader.parse(questao.c_str(), root);

    if (!parsingSuccessful){
        throw erroAtividadeoNaoExiste();
    }

    //O guardar é um json array- com os objetos de cada atividade do dia escolhido.
     guardar = root.get("atividades","");
    
    int i;
    i = 0;
   
    tamanho = guardar.size();
   
    /*Caso não há nenhuma atividade para o dia escolhido*/
    nome = fastwriter.write(guardar[0]["nome"]);
        if (nome == "null\n"){
            i = tamanho+1;
            cout<<"Nenhuma atividade registrada para esse dia"<<endl;
        }
    /* Passa por todos os objetos do array "atividades" daquele dia e os printa na tela*/
    while (i < tamanho){
        
     
        cout<<"Nome da Atividade: "<<guardar[i]["nome"]<<endl;
        cout<<"Horário de início: "<<guardar[i]["início"]<<"h"<<endl;
        cout<<"Horário de término: "<<guardar[i]["fim"]<<"h"<<endl;
        cout<<"----------------------------------------------"<<endl;
        i++;
    }


}
void Agenda::mostrarAtividadesId(string questao){
    
    Json::Value root;
    Json::Reader reader;
    Json::Value guardar;
    Json::FastWriter fastwriter;
    string nome;
    int i = 0;
    bool parsingSuccessful = reader.parse(questao.c_str(), root);

    if (!parsingSuccessful){
        throw erroAtividadeoNaoExiste();
    }

     guardar = root.get("atividades","");

        nome = fastwriter.write(guardar[0]["nome"]);
        if (nome == "null\n"){
            i = 1;
            cout<<"Nenhuma atividade registrada com esse ID"<<endl;
        }
    
    //o guardar é um json array- com os objetos de cada atividade. 

        if (i == 0){
        cout<<"\nAtividade:\n"<<endl;
        /* Passa pelo único objeto que é recebido em "atividades"*/
        cout<<"Nome da Atividade: "<<guardar[0]["nome"]<<endl;
        cout<<"Dia da semana: "<<guardar[0]["dia"]<<endl;
        cout<<"Horário de início: "<<guardar[0]["início"]<<"h"<<endl;
        cout<<"Horário de término: "<<guardar[0]["fim"]<<"h"<<endl;
        cout<<"----------------------------------------------"<<endl;
        }
}

   

void Agenda::salvarAtividades(string questao ){
   

    Json::Value root;
    Json::Reader reader;
    Json::Value guardar;
    int tamanho;
    Json::Value nome1;
    Json::FastWriter fastwriter;
    string  dia ;
    string nome;
    string inicio;
    string fim  ;
    string body;
   

   bool parsingSuccessful = reader.parse(questao.c_str(), root);
   

    if (!parsingSuccessful)
        throw erroAtividadeoNaoExiste();
    
    guardar = root.get("atividades","");


    fstream fout;
  
    // opens an existing csv file or creates a new file.
    fout.open("atividades.csv", ios::out | ios::app);
        
    int i;
    i = 0;

    fout<<"nome;dia;inicio;fim\n";
  
    tamanho = guardar.size();
   
    cout<<"--------------Atividades Salvas-------------"<<endl;
    while (i < tamanho){
        
        /*Guarda os dados dos objetos em uma string*/
        nome = fastwriter.write(guardar[i]["nome"]);
        dia = fastwriter.write(guardar[i]["dia"]);
        inicio = fastwriter.write(guardar[i]["início"]);
        fim = fastwriter.write(guardar[i]["fim"]);

        //todos eles tem um /n no final , logo preciso tirá-los
        nome.erase(std::remove(nome.begin(), nome.end(), '\n'), nome.end());
        dia.erase(std::remove(dia.begin(), dia.end(), '\n'), dia.end());
        inicio.erase(std::remove(inicio.begin(), inicio.end(), '\n'), inicio.end());
        fim.erase(std::remove(fim.begin(), fim.end(), '\n'), fim.end());
       
        /*Escreve os dados no rquivo csv, os dados são separados por ; o que significa que vão pertencer a colunas diferntes*/
        fout << nome<< ";"
             << dia << ";"
             << inicio << ";"
             << fim <<";"
             << "\n";
        
        i++;
    } 


}

