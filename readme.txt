Trabalho Final

Autores:
-------------------------------------------
Nome: Catarina Oliveira Dowsley Fernandes |
DRE: 11802011				  |
------------------------------------------|
Nome: Thiago Felipe Cruz e Souza	  |
DRE:119035356				  |
-------------------------------------------

Disciplina: Linguagens de Programação - EEL670
Professor: Miguel Elias
Universidade Federal do Rio de Janeiro - 2021.2

----Instruções sobre o Programa-----

Esse programa gerencia um banco de dados de atividades cadastradas pelo próprio usuário.

Instruções sobre como utilizar o Menu.
Basta selecionar uma das seguintes opções:
1. Adicionar Atividade
2. Deletar Atividade 
3. Ver todas as atividades                
4. Ver atividades de um dia específico         
5. Ver atividade a partir de do ID      
6. Alterar Atividade                          
7. Salvar atividades em um arquivo csv     
8. Sair         


Para compilar basta digitar no terminal:
make

Para executar:
./main

Para compliar novamente após alguma alteração no código é necessário digitar:
make clean

---Informações sobre o Programa-----

Esse trabalho foi feito usando o Visual Studio Code para a edição dos arquivos e foi compilado em Linux (Ubuntu- 18.04).

O arquivo atividades.csv é gerado após selecionada a opção 7. Salvar atividades em um arquivo csv. 


Foi feito o deploy do backend no Heroku. Nossa url base é:
https://agenda-semanal.herokuapp.com

Nossos endpoints:

POST   /atividade   - Usado para criar uma nova atividade      
PUT    /atividade/:id    - Usado para alterar uma atividade 
GET    /atividade/:id    - Usado pra encontrar uma atividade já existente pelo ID 
GET    /atividade/dia/:dia - Usado pra encontrar as atividades de um dia da semana específico
GET    /atividade     - Usado para encontrar todas as atividades cadastradas    
DELETE /atividade/:id - Usado para deletar uma atividade a partir do seu id 
