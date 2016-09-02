/*
** talker.c -- a datagram "client" demo
*/

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include "shared/c/types.h"
#include "shared/c/io.h"


#define MAXBUFLEN 100

int main(int argc, char *argv[])
{
	int sockfd;
	struct addrinfo hints, *servinfo, *p;
	int rv;
	int numbytes;
	char buf[MAXBUFLEN];
	struct sockaddr_storage their_addr;
	socklen_t addr_len;

	if (argc != 3) {
		fprintf(stderr,"usage: %s hostname port\n", argv[0]);
		exit(1);
	}

	memset(&hints, 0, sizeof hints);
	hints.ai_family = AF_UNSPEC;
	hints.ai_socktype = SOCK_DGRAM;

	if ((rv = getaddrinfo(argv[1], argv[2], &hints, &servinfo)) != 0) {
		fprintf(stderr, "getaddrinfo: %s\n", gai_strerror(rv));
		return 1;
	}

	// loop through all the results and make a socket
	for(p = servinfo; p != NULL; p = p->ai_next) {
		if ((sockfd = socket(p->ai_family, p->ai_socktype,
				p->ai_protocol)) == -1) {
			perror("talker: socket");
			continue;
		}

		break;
	}

	if (p == NULL) {
		fprintf(stderr, "talker: failed to create socket\n");
		return 2;
	}

    unsigned char reqNum = 0;

    while(1) {

     
      calcrequest_t request = readCalculatorRequest(&reqNum);

      if ((numbytes = sendto(sockfd, &request, 8, 0,
               p->ai_addr, p->ai_addrlen)) == -1) {
          perror("talker: sendto");
          exit(1);
      }

	  addr_len = sizeof their_addr;

	  if ((numbytes = recvfrom(sockfd, buf, MAXBUFLEN-1 , 0,
		(struct sockaddr *) &their_addr, &addr_len)) == -1) {
		perror("recvfrom");
		exit(1);
	  }
      char tml = buf[0];
      char i;

      for (i = (char)0; i < tml; i++) {
        printf("0x%02x\n", buf[(int)i]);
      }

      calcresponse_t response = calcresponseFromBytes(buf, (int) tml);
      printf("Request #%d: %d\n", response.RequestID, response.Result);


	  //printf("listener: packet contains \"%s\"\n", buf);
    }

	freeaddrinfo(servinfo);

	printf("talker: sent %d bytes to %s\n", numbytes, argv[1]);
	close(sockfd);

	return 0;
}
