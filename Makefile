# Directory Names
OBJ=obj
BUILD=build
SHARED=shared
SRC=src
BIN=bin

# Compilation flags
CC=gcc
CFLAGS=-Wall -Isrc/

# Go setup flags
GO_DOWNLOAD_SRC="https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz"
GO_TARBALL=godownload.tar.gz
GROOT=$(shell pwd)/$(BIN)/go
export GOROOT:=$(GROOT)

# Go compilation flags.
SETPATH=GOPATH=$(shell pwd)
GOCMD=$(BIN)/go/bin/go
GO=$(SETPATH) $(GOCMD) build
GOFLAGS=

# =================================
#			Main commands
# =================================
all: tcp udp
	@echo "Done building all components"

setup: setup-dirs setup-golang
	@echo "Done with setup"

clean: 
	@rm -f $(OBJ)/*.o $(OBJ)/$(SHARED)/*.o $(BUILD)/*; \
	echo "Done with cleanup"

tcp: tcp_server tcp_client tcp_util
	@echo "Done building TCP components"

udp: udp_server udp_client udp_util
	@echo "Done building UDP components"

util: tcp_util udp_util
	@echo "Done building util components"

# =================================
#			Setup project and install go
# =================================
setup-dirs:
	@mkdir -p $(BUILD) $(OBJ)/$(SHARED) $(BIN)

setup-golang:
	@echo "Downloading Go..."
	@curl $(GO_DOWNLOAD_SRC) > $(BIN)/$(GO_TARBALL) && \
	echo "Extracting Go..." && \
	tar -xzf $(BIN)/$(GO_TARBALL) -C $(BIN) && \
	echo "Go installed successfully!"

# =================================
#			Build shared C files.
# =================================
shared: src/$(SHARED)/*
	@$(CC) $(CFLAGS) -c src/shared/c/io.c -o $(OBJ)/$(SHARED)/io.o; \
	$(CC) $(CFLAGS) -c src/shared/c/types.c -o $(OBJ)/$(SHARED)/types.o; \
	echo "Done building shared files"

# =================================
#			Build individual components.
# =================================
tcp_server_base: src/tcp/server/* 
	@$(CC) $(CFLAGS) -c src/tcp/server/*.c -o $(OBJ)/tcp_server.o

tcp_server: shared tcp_server_base 
	@$(CC) $(CFLAGS) $(OBJ)/tcp_server.o $(OBJ)/$(SHARED)/*o -o $(BUILD)/TCP_server

tcp_client: src/tcp/client/*
	@$(GO) $(GOFLAGS) -o $(BUILD)/TCP_client tcp/client

tcp_util: src/util/TCPServerDisplay.c
	@$(CC) $(CFLAGS) src/util/TCPServerDisplay.c -o $(BUILD)/TCP_diagnose

udp_server: src/udp/server/*
	@$(GO) $(GOFLAGS) -o $(BUILD)/UDP_server udp/server

udp_client_base: src/udp/client/* 
	@$(CC) $(CFLAGS) -c src/udp/client/*.c -o $(OBJ)/udp_client.o

udp_client: shared udp_client_base
	@$(CC) $(CFLAGS) $(OBJ)/udp_client.o $(OBJ)/$(SHARED)/*o -o $(BUILD)/UDP_client

udp_util: src/util/UDPServerDisplay.c
	@$(CC) $(CFLAGS) src/util/UDPServerDisplay.c -o $(BUILD)/UDP_diagnose
