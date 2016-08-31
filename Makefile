CC=gcc
PWD=$(pwd)
GPATH=$(GOPATH)
SETPATH=GOPATH=$(GPATH):$(shell pwd)
GO=$(SETPATH) go build
GOFLAGS=
CFLAGS=-Wall -Isrc/
OBJ=obj
BUILD=build
SHARED=shared
SRC=src
SHARED_SOURCES = $(shell find $(SRC)/$(SHARED) -name '*.c')
SHARED_OBJECTS = $(SHARED_SOURCES:%.c=%.o)

all: tcp udp
	@echo "Done building all components"

setup:
	@mkdir -p $(BUILD) $(OBJ)/$(SHARED); \

clean: setup
	@rm -f $(OBJ)/* $(BUILD)/*; \
	echo "Done with cleanup"

tcp: tcp_server tcp_client tcp_util
	@echo "Done building TCP components"

udp: udp_server udp_client udp_util
	@echo "Done building UDP components"

util: tcp_util udp_util
	@echo "Done building util components"


shared: src/$(SHARED)/* setup 
	@$(CC) $(CFLAGS) -c src/shared/c/io.c -o $(OBJ)/$(SHARED)/io.o; \
	$(CC) $(CFLAGS) -c src/shared/c/types.c -o $(OBJ)/$(SHARED)/types.o; \
	echo "Done building shared files"

tcp_server: src/tcp/server/* setup
	@$(CC) $(CFLAGS) src/tcp/server/*.c -o $(BUILD)/TCP_server

tcp_client: src/tcp/client/* setup
	@$(GO) $(GOFLAGS) -o $(BUILD)/TCP_client tcp/client

tcp_util: src/util/TCPServerDisplay.c setup
	@$(CC) $(CFLAGS) src/util/TCPServerDisplay.c -o $(BUILD)/TCP_diagnose

udp_server: src/udp/server/* setup
	@$(CC) $(CFLAGS) src/udp/server/*.c -o $(BUILD)/UDP_server

udp_client_base: src/udp/client/* setup
	@$(CC) $(CFLAGS) -c src/udp/client/*.c -o $(OBJ)/udp_client.o

udp_client: shared udp_client_base setup
	@$(CC) $(CFLAGS) $(OBJ)/udp_client.o $(OBJ)/$(SHARED)/*o -o $(BUILD)/UDP_client

udp_util: src/util/UDPServerDisplay.c setup
	@$(CC) $(CFLAGS) src/util/UDPServerDisplay.c -o $(BUILD)/UDP_diagnose
