#ifndef MENU_H
#define MENU_H
#include <iostream>
#include <string>
using namespace std;


class Menu {
    public:
    /*Declaração das funções */
        void mostrarMenuPrincipal();
        unsigned receberOpcaoMenu();
        unsigned receberOpcaoMenuDia();
        void finalizarPrograma();

};

#endif