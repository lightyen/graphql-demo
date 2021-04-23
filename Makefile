GO_BUILD := go build
GO_VENDER := 1
LDFLAGS := -s -w

DATE := $(shell date +%Y%m%d)
PLATFORM_ARCH := $(shell uname -m)
PLATFORM_OS := $(shell uname -o)

PKG_NAME := app
APP_NAME := app

ifeq ($(GO_VENDER), 1)
GO_BUILD = go build -mod=vendor
endif

all: currplatform

currplatform:
	go mod tidy
	@echo "building app for ${PLATFORM_OS}-${PLATFORM_ARCH}"
	gqlgen
	$(GO_BUILD) -ldflags="${LDFLAGS}" -o $(APP_NAME) $(PKG_NAME)

clean:
	@rm app

check:
ifeq ($(GO_VENDER), 1)
	GOFLAGS=-mod=vendor golangci-lint run --enable golint
else
	golangci-lint run --enable golint
endif

init:
	go install github.com/99designs/gqlgen@v0.13.0
	go get -d github.com/99designs/gqlgen@v0.13.0
ifeq ($(GO_VENDER), 1)
	go mod vendor
endif
	