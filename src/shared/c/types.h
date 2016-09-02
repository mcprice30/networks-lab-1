#ifndef SHARED_TYPES_H
#define SHARED_TYPES_H

#include <stdio.h>
#include <stdbool.h>
#include <byteswap.h>
#include <stdint.h>

#define HANDLE_SHORT(X) isLittleEndian() ? __bswap_16(X) : X 
#define HANDLE_INT(X) isLittleEndian() ? __bswap_32(X) : X
#define HANDLE_OPERAND_T(X) HANDLE_SHORT(X)

typedef int16_t operand_t;

struct calcrequest
{
  unsigned char TML;
  unsigned char RequestID;
  unsigned char OpCode;
  unsigned char NumOperands;
  operand_t Operand1; 
  operand_t Operand2;
} __attribute__((__packed__));


struct calcresponse
{
  unsigned char TML;
  unsigned char RequestID;
  unsigned char ErrorCode;
  int32_t Result;
} __attribute__((__packed__));

typedef struct calcrequest calcrequest_t;
typedef struct calcresponse calcresponse_t;

bool isLittleEndian();

calcrequest_t readCalculatorRequest(unsigned char *reqNum);

calcresponse_t calcresponseFromBytes(char* bytes, int len);

#endif
