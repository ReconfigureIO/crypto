# variable definitions
NAME := crypto
VERSION := $(shell git describe --tags --always --dirty)
GOVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDDATE := $(shell date -u +"%B %d, %Y")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>")
PKG_RELEASE ?= 1
PROJECT_URL := "https://github.com/ReconfigureIO/$(NAME)"

.PHONY: test vendor

test:
	go test github.com/ReconfigureIO/crypto

vendor: examples/md5/vendor/github.com/ReconfigureIO/$(NAME)/md5

examples/md5/vendor/github.com/ReconfigureIO/$(NAME)/md5: md5
	mkdir -p examples/md5/vendor/github.com/ReconfigureIO/$(NAME)/
	cp md5 examples/md5/vendor/github.com/ReconfigureIO/$(NAME)/md5
