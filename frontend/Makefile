CC = g++
LD = g++

CFLAGS = -Wall
LFLAGS = -Wall

MAINOBJS = agenda.o menu.o requisicao.o main.o 

MAIN = main

.cpp.o:
	$(CC) $(CFLAGS) -c $<

all: $(MAIN)

main: $(MAINOBJS)
	$(LD) $(LFLAGS) -o $@ $(MAINOBJS) -lm -lcurl -ljsoncpp

clean:
	rm -f *.o $(MAIN)