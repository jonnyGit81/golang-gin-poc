#!/usr/bin/env bash

Set -xe

#install package and depedencies for delploying into AWS Beanstalk
go get github.com/gin-gionic/gin

go get github.com/go-playground/validator

# build command, this is will generate a binary file named as application in the bin folder
go build -o bin/application server.go
