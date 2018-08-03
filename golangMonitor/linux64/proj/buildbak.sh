#!/bin/bash

#rm -f libsmth.{o,a} test
#g++ -L=../lib/ -Wall -c -o ../lib/libsmth.o smth.cpp -rpath=./:./HCNetSDKCom:../lib -lhcnetsdk
#go build --ldflags '-extldflags "-static"' -o test main.go && rm -f libsmth.{o,a}
#g++ smth.cpp -L=../lib/ -Wall -c -o ../lib/libsmth.o && ar rcs ../lib/libsmth.a ../lib/libsmth.o 
#g++ smth.cpp -L=../lib/ -Wall -c -o libsmth.o && ar rcs libsmth.a libsmth.o 
go build --ldflags '-extldflags "-static"' -o test main.go
