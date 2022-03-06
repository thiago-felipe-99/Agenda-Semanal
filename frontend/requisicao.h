#ifndef REQUISICAO_H
#define REQUISICAO_H
#include <iostream>
#include <string>
#include <curl/curl.h>
#include "excecao.h"
using namespace std;

/*Declaração da classe e das funções que usamos para as requisições*/
class Requisicao{
    public:
        Requisicao(){};
        string get(string url);
        string post(string url, string body);
        string deletar(string url);
        string put(string url, string body);

    
    private:
        string curlRequisicaoGet(string url);
        string curlRequisicaoPost(string url, string body);
        string curlRequisicaoDeletar(string url);
        string curlRequisicaoPut(string url, string body);
        static size_t WriteCallBack(char *contents, size_t size, size_t nmemb, char*buffer_in);

};

#endif