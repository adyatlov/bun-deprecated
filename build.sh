#/bin/bash

mkdir -p bin
env GOOS=linux   GOARCH=amd64 go build -o bin/bun_linux_amd64   github.com/adyatlov/bun/bun
env GOOS=darwin  GOARCH=amd64 go build -o bin/bun_darwin_amd64  github.com/adyatlov/bun/bun
env GOOS=windows GOARCH=amd64 go build -o bin/bun_windows_amd64 github.com/adyatlov/bun/bun
