GO_BUILD := go build
GO_VENDER := 0
LDFLAGS := -s -w

DATE := $(shell date +%Y%m%d)
PLATFORM_ARCH := $(shell uname -m)
PLATFORM_OS := $(shell uname -o)

PKG_NAME := app
APP_NAME := app

ifeq ($(GO_VENDER), 1)
GO_BUILD = $(GO_BUILD) -mod=vendor
endif

all: currplatform

currplatform:
	go install github.com/99designs/gqlgen@v0.13.0
	@echo "building app for ${PLATFORM_OS}-${PLATFORM_ARCH}"
	go mod tidy
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
