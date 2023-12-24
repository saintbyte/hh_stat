#!/bin/bash
mkdir ./bin
GOOS=windows GOARCH=386 go build -o bin/386.exe  main.go
GOOS=windows GOARCH=amd64 go build -o bin/64.exe  main.go
go build -o bin/main  main.go
