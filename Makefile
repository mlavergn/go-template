###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := run

.PHONY: test

run: mod
	go run main.go

build: mod
	go build main.go

clean:
	rm -r vendor

test:
	go test -v -count=1 ./...

lint:
	go vet

format:
	go fmt

mod:
	go mod tidy
	go mod vendor

macintel:
	GOOS=darwin GOARCH=amd64 go build main.go

mac:
	GOOS=darwin GOARCH=arm64 go build main.go

linux:
	GOOS=linux GOARCH=amd64 go build main.go

# MT7621 - MIPS 32 + softfloat
linuxmips32:
	GOOS=linux GOARCH=mips32le GOMIPS=softfloat go build main.go

#
# Publishing
#

VERSION := 1.0.0
PROJECT := demo
REPO := mlavergne

github:
	open "https://github.com/${REPO}/${PROJECT}"

release:
	zip -r ${PROJECT}.zip LICENSE README.md Makefile main.go go.mod vendor
	hub release create -m "${VERSION} - ${PROJECT}" -a ${PROJECT}.zip -t master "v${VERSION}"
	open "https://github.com/${REPO}/${PROJECT}/releases"

st:
	open -a SourceTree .