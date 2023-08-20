###############################################
#
# Makefile
#
###############################################

.DEFAULT_GOAL := run

.PHONY: test

run: deps
	go run .

build: deps
	go build .

clean: lint format
	rm -rf main vendor

test:
	go test -v -count=1 .

lint:
	go vet

format:
	go fmt

deps:
	go mod tidy
	go mod vendor

# ldflags for smallest possible application (~40% smaller)
# `go tool link`
# `-w` disable DWARF generation
# `-s` disable debug symbols
# CGO_ENABLED=0 to disable libc bindings
dist:
	CGO_ENABLED=0 go build --ldflags "-s -w" -o main .

# ldflags for static builds (does not work on macos)
# `-extldflags "-static"` omit dyld dynamic loader
static:
	CGO_ENABLED=0 go build --ldflags '-s -w -linkmode external -extldflags "-static"' .

macintel:
	GOOS=darwin GOARCH=amd64 go build .

mac:
	GOOS=darwin GOARCH=arm64 go build .

linux:
	GOOS=linux GOARCH=amd64 go build .

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