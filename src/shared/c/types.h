#ifndef SHARED_TYPES_H
#define SHARED_TYPES_H

#include <stdio.h>
#include <stdbool.h>
#include <byteswap.h>

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

bool isLittleEndian();

calcrequest_t readCalculatorRequest(unsigned char *reqNum);

#endif
