#!/bin/bash

#rm -f libsmth.{o,a} test
g++ -L./lib/ -Wall -c -o libsmth.o smth.cpp && ar rcs libsmth.a libsmth.o ../proj/CapPicture.o
#go build --ldflags '-extldflags "-static"' -o test main.go && rm -f libsmth.{o,a}
go build --ldflags '-extldflags "-static"' -o test main.go
