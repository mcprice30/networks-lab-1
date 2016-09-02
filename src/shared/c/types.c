#include <stdint.h>

#include "shared/c/types.h"
#include "shared/c/io.h"

bool isLittleEndian() {
  short a = 1;
  return (bool) ((char*) &a)[0];
}

calcrequest_t readCalculatorRequest(unsigned char *reqNum) {

  unsigned char opCode;
  operand_t operand1, operand2;

  readSanitized("Enter opcode: ", &opCode, sizeof(char));
  readSanitized("Enter operand 1: ", &operand1, sizeof(operand_t));
  readSanitized("Enter operand 2: ", &operand2, sizeof(operand_t));

  calcrequest_t request;

  request.TML = 8;
  request.RequestID = (*reqNum)++;
  request.OpCode = opCode;
  request.NumOperands = 2;
  request.Operand1 = HANDLE_OPERAND_T(operand1);
  request.Operand2 = HANDLE_OPERAND_T(operand2);

  return request;
}


calcresponse_t calcresponseFromBytes(char* bytes, int len) {

  calcresponse_t response;

  if (len != 7) {
    return response;
  }

  response.TML = bytes[0];
  response.RequestID = bytes[1];
  response.ErrorCode = bytes[2];
  char* resultOffset = &(bytes[3]);
  response.Result = HANDLE_INT(*((int32_t*)resultOffset));

  return response;
}
