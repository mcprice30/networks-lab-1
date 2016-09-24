// Lab 1 - Group 11
// 
// types.c contains the implementation for the signatures in types.h
//
#include "shared/c/types.h"
#include "shared/c/io.h"

bool isLittleEndian() {
  short a = 1;
  return (bool) ((char*) &a)[0];
}

calcrequest_t readCalculatorRequest(unsigned char *reqNum) {

  unsigned char opCode;
  operand_t operand1, operand2;

  // Prompt the user for input.
  readSanitized("Enter opcode: ", &opCode, sizeof(char));
  readSanitized("Enter operand 1: ", &operand1, sizeof(operand_t));
  readSanitized("Enter operand 2: ", &operand2, sizeof(operand_t));

  // Build the request.
  calcrequest_t request;

  request.TML = REQUEST_TML;
  request.RequestID = (*reqNum)++;
  request.OpCode = opCode;
  request.NumOperands = NUM_OPERANDS;
  request.Operand1 = HANDLE_OPERAND_T(operand1);
  request.Operand2 = HANDLE_OPERAND_T(operand2);

  return request;
}

calcrequest_t calcrequestFromBytes(char* bytes, int len) {

  calcrequest_t request;

  if (len != REQUEST_TML || bytes[0] != REQUEST_TML) {
    return request; // return an empty request.
  }

  request.TML = bytes[0];
  request.RequestID = bytes[1];
  request.OpCode = bytes[2];
  request.NumOperands = bytes[3];
  request.Operand1 = HANDLE_OPERAND_T(*((operand_t*) &bytes[4]));
  request.Operand2 = HANDLE_OPERAND_T(*((operand_t*) &bytes[6]));

  return request;
}

int calcresponseToBytes(calcresponse_t response, char* bytes, int numBytes) {

  calcresponse_t copy; 
  if (numBytes != RESPONSE_TML) {
    return -1;
  }

  copy.TML = response.TML;
  copy.RequestID = response.RequestID;
  copy.ErrorCode= response.ErrorCode;
  copy.Result = HANDLE_INT(response.Result);

  memcpy(bytes, &copy, numBytes);
  return 0;
}

calcresponse_t calcresponseFromBytes(char* bytes, int len) {

  calcresponse_t response;

  if (len != RESPONSE_TML || bytes[0] != RESPONSE_TML) {
    return response; // return empty response.
  }

  // Set response values.
  response.TML = bytes[0];
  response.RequestID = bytes[1];
  response.ErrorCode = bytes[2];
  char* resultOffset = &(bytes[3]);
  response.Result = HANDLE_INT(*((int32_t*)resultOffset));

  return response;
}
