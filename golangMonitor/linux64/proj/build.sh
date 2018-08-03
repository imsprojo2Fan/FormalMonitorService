#!/bin/bash
rm libCapPicture.so && rm CapPicture.o
gcc -c -fPIC -o CapPicture.o CapPicture.c
gcc -shared -o libCapPicture.so CapPicture.o #&& go run main.go
