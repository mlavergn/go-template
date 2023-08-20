###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := run

.PHONY: test

run: mod
	go run *.go

build: mod
	go build *.go

clean: lint format
	rm -rf main vendor

test:
	go test -v -count=1 ./...

lint:
	go vet

format:
	go fmt

mod:
	go mod tidy
	go mod vendor

# ldflags for smallest possible binary (~40% smaller)
# `go tool link`
# `-w` disable DWARF generation
# `-s` disable debug symbols
dist:
	go build --ldflags "-s -w" -o main *.go

macintel:
	GOOS=darwin GOARCH=amd64 go build *.go

mac:
	GOOS=darwin GOARCH=arm64 go build *.go

linux:
	GOOS=linux GOARCH=amd64 go build *.go

# MT7621 - MIPS 32 + softfloat
linuxmips32:
	GOOS=linux GOARCH=mips32le GOMIPS=softfloat go build *.go

#
# Publishing
#

VERSION := 1.0.0
PROJECT := demo
REPO := repo

github:
	open "https://github.com/${REPO}/${PROJECT}"

release:
	zip -r ${PROJECT}.zip LICENSE README.md Makefile *.go go.mod
	gh release create v${VERSION} ${PROJECT}.zip --target master --notes "${VERSION} - ${PROJECT}"
	open "https://github.com/${REPO}/${PROJECT}/releases"

st:
	open -a SourceTree .