#ifndef EXCECAO_H
#define EXCECAO_H
#include <exception>

using namespace std;

/*Definição dos possíveis erros que podem acontecer*/

class erroCurlNaoReconhecido : public exception{
    public:
        const char *what() const throw(){
            
            const char* msgErro = "\n Erro -- O pacote Curl não foi reconhecido e provavelmente não está corretamente instalado na máquina. \n \n";

            return msgErro;
        }
};

class erroProtocoloNaoSuportado: public exception{
    public:
        const char *what() const throw(){
            
            const char* msgErro = "\n Erro -- A URL informada utilizada um protocolo não suportado pelo sistema. \n \n";

            return msgErro;
        }
};

class erroUrlFormatoInvalido: public exception{
    public:
        const char *what() const throw(){
            
            const char* msgErro = "\n Erro -- A URL informada está em um formato inválido. \n \n";

            return msgErro;
        }
};

class erroProblemaConexao : public exception{
    public:
        const char *what() const throw(){
            
            const char* msgErro = "\n Erro -- Não foi possível se conectar, cheque sua conexão com a internet. \n \n";

            return msgErro;
        }
};

class erroGenericoRequisicao : public exception{
    public:
        const char *what() const throw(){
            
            const char* msgErro = "\n Erro ao executar a requisição. \n \n";

            return msgErro;
        }
};

class erroAtividadeoNaoExiste : public exception{
    public:
        const char *what() const throw(){
            
            const char* msgErro = "\n Erro -- Não existe nenhuma atividade cadastrada com esse número. \n \n";

            return msgErro;
        }
};

#endif