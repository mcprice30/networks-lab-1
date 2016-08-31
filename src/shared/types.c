#include "shared/types.h"
#include "shared/io.h"

bool isLittleEndian() {
  short a = 1;
  return (int) ((char*) &a)[0];
}


calcrequest_t readCalculatorRequest(unsigned char *reqNum) {

  unsigned char opCode;
  unsigned short operand1, operand2;

  readSanitized("Enter opcode: ", &opCode, sizeof(char));
  readSanitized("Enter operand 1: ", &operand1, sizeof(short));
  readSanitized("Enter operand 2: ", &operand2, sizeof(short));

  calcrequest_t request;

  request.TML = 8;
  request.RequestID = (*reqNum)++;
  request.OpCode = opCode;
  request.NumOperands = 2;
  request.Operand1 = HANDLE_SHORT(operand1);
  request.Operand2 = HANDLE_SHORT(operand2);

  return request;
}
