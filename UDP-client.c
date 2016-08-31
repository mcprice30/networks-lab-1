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
#include <byteswap.h>

// #define SERVERPORT "10010"	// the port users will be connecting to

#define HANDLE_SHORT(X) isLittleEndian() ? __bswap_16(X) : X 
#define HANLDE_INT(X) isLittleEndian() ? __bswap_32(X) : X

struct calcrequest
{
  unsigned char TML;
  unsigned char RequestID;
  unsigned char OpCode;
  unsigned char NumOperands;
  unsigned short Operand1; 
  unsigned short Operand2;
} __attribute__((__packed__));

typedef struct calcrequest calcrequest_t;

bool isLittleEndian() {
  short a = 1;
  return (int) ((char*) &a)[0];
}

void readSanitized(const char *promptString, void *readInto, int numBytes) {

  int input;
  bool okSize;
  int valsRead;

  do {
    fputs(promptString, stdout);

    valsRead = scanf("%d", &input);  

    if (valsRead != 1) {
      perror("invalid input");
    }

    okSize = true;

    if (numBytes == 1)
    { 
      if (input > 0xff)
      {
        okSize = false;
        fprintf(stderr, "value must fit in 1 byte");
      } else {
        *(unsigned char*)readInto = (unsigned char)input;
      }
    } else if (numBytes == 2)
    {
      if (input > 0xffff)
      {
        okSize = false;
        fprintf(stderr, "value must fit in 2 bytes");
      } else {
        *(unsigned short*)readInto = (unsigned short)input;
      }
    } else {
      *(unsigned int*)readInto = input;
    }

  } while (valsRead != 1 || !okSize);

}

int main(int argc, char *argv[])
{
	int sockfd;
	struct addrinfo hints, *servinfo, *p;
	int rv;
	int numbytes;

	if (argc != 3) {
		fprintf(stderr,"usage: talker hostname message\n");
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

      
      unsigned char opCode;
      unsigned short operand1, operand2;

      readSanitized("Enter opcode: ", &opCode, sizeof(char));
      printf("%02x\n", opCode);
      readSanitized("Enter operand 1: ", &operand1, sizeof(short));
      printf("%04x\n", operand1);
      readSanitized("Enter operand 2: ", &operand2, sizeof(short));
      printf("%04x\n", operand2);

      calcrequest_t request;

      request.TML = 8;
      request.RequestID = reqNum++; 
      request.OpCode = opCode;
      request.NumOperands = 2;
      request.Operand1 = HANDLE_SHORT(operand1);
      request.Operand2 = HANDLE_SHORT(operand2);

      if ((numbytes = sendto(sockfd, &request, 8, 0,
               p->ai_addr, p->ai_addrlen)) == -1) {
          perror("talker: sendto");
          exit(1);
      }

    }

	freeaddrinfo(servinfo);

	printf("talker: sent %d bytes to %s\n", numbytes, argv[1]);
	close(sockfd);

	return 0;
}
