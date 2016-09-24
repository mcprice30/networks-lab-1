// Lab 1 - Group 11
// 
// types.h contains the various types used by the client in server.
// In particular, it encapsulates types for both request and response packets.
//
#ifndef SHARED_TYPES_H
#define SHARED_TYPES_H

#include <stdio.h>
#include <stdbool.h>
#include <byteswap.h>
#include <stdint.h>
#include <string.h>

#define REQUEST_TML 8
#define RESPONSE_TML 7
#define NUM_OPERANDS 2

// Convert values to network byte order.
#define HANDLE_SHORT(X) isLittleEndian() ? __bswap_16(X) : X 
#define HANDLE_INT(X) isLittleEndian() ? __bswap_32(X) : X
#define HANDLE_OPERAND_T(X) HANDLE_SHORT(X)

// Encapsulate operand types.
typedef uint16_t operand_t;

// calcrequest represents the data sent in a request packet.
struct calcrequest
{
  unsigned char TML;
  unsigned char RequestID;
  unsigned char OpCode;
  unsigned char NumOperands;
  operand_t Operand1; 
  operand_t Operand2;
} __attribute__((__packed__));

// calcresponse holds the data sent in a response packet.
struct calcresponse
{
  unsigned char TML;
  unsigned char RequestID;
  unsigned char ErrorCode;
  int32_t Result;
} __attribute__((__packed__));

typedef struct calcrequest calcrequest_t;
typedef struct calcresponse calcresponse_t;

// isLittleEndian returns true iff the given system is little endian.
bool isLittleEndian();

// calcrequestFromBytes takes a byte array and the length of the byte array,
// and produces a request packet.
calcrequest_t calcrequestFromBytes(char* bytes, int len);

// readCalculatorRequest, given the # of the current request, will prompt the
// client's user for input, and then build a request packet.
calcrequest_t readCalculatorRequest(unsigned char *reqNum);

// calcresponseToBytes takes a response packet and converts it to bytes.
// It will write the result into bytes, which will have a length of numBytes.
// Returns 0 iff the conversion was successful.
int calcresponseToBytes(calcresponse_t response, char* bytes, int numBytes);

// calcresponseFromBytes will take a byte array of size len and convert it to
// a response packet.
calcresponse_t calcresponseFromBytes(char* bytes, int len);

#endif
