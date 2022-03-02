#!/usr/bin/env sh

git clone https://github.com/thiago-felipe-99/Agenda-Semanal.git /agenda-semanal
cd /agenda-semanal/backend
go build -o main
./main
