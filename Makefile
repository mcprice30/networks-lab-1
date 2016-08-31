all: tcp udp
	echo "Done building all components"

setup:
	mkdir -p build

clean: setup
	rm build/*

tcp: tcp_server tcp_client tcp_util
	echo "Done building TCP components"

udp: udp_server udp_client udp_util
	echo "Done building UDP components"

util: tcp_util udp_util
	echo "Done building util components"

tcp_server: TCP-server.c setup
	gcc TCP-server.c -o build/TCP_server

tcp_client: TCP-client.c setup
	gcc TCP-client.c -o build/TCP_client

tcp_util: TCPServerDisplay.c setup
	gcc TCPServerDisplay.c -o build/TCP_diagnose

udp_server: UDP-server.c setup
	gcc UDP-server.c -o build/UDP_server

udp_client: UDP-client.c setup
	gcc UDP-client.c -o build/UDP_client

udp_util: UDPServerDisplay.c setup
	gcc UDPServerDisplay.c -o build/UDP_diagnose
