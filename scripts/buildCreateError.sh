#!/bin/bash
rm -f bin/*
mkdir -p bin/
go build -o bin/createError ./cmd/createError
echo "Run ./bin/createError"
