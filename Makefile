all: bin/main bin/test

clear:
	rm -f build/* bin/*

test: bin/test

main: bin/main

CFLAG= -g -Wall -Wextra -Wpedantic -Wconversion

INCLUDES= src/*.h

COMMON_DEPS= $(INCLUDES) Makefile

build/%.o: src/%.c $(COMMON_DEPS)
	$(CC) $(CFLAGS) -c $< -o $@

bin/test: build/takeNumberTest.o build/takeTheNumber.o $(COMMON_DEPS)
	$(CC) -o bin/testTakeNumber build/takeNumberTest.o build/takeTheNumber.o


bin/main: build/countTheNumber.o build/takeTheNumber.o $(COMMON_DEPS)
	$(CC) -o bin/countTheNumber build/countTheNumber.o build/takeTheNumber.o
