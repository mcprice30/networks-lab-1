// Lab 1 - Group 11
// 
// io.c contains the implementation of helper functions defined in io.h
//

#include "shared/c/io.h"

void readSanitized(const char *promptString, void *readInto, int numBytes) {

  int input;
  bool okSize;
  int valsRead;

  do { // Repeat until we get valid input.

    fputs(promptString, stdout); // Print the prompt, read in the user's input.
    valsRead = scanf("%d", &input);  

    // We did not get valid input. Retry.
    if (valsRead != 1) {
      perror("invalid input");
      while (fgetc(stdin) != '\n'){}
      continue;
    }

    // Will be true iff we got a valid value.
    okSize = true;

    if (numBytes == 1) {  // We are trying to read a byte.
      if (input > 0xff) {
        okSize = false;
        fprintf(stderr, "value must fit in 1 byte\n");
      } else {
        *(unsigned char*)readInto = (unsigned char)input;
      }
    } else if (numBytes == 2) { // We are trying to read a 2-byte value.
      if (input > 0xffff) {
        okSize = false;
        fprintf(stderr, "value must fit in 2 bytes\n");
      } else {
        *(operand_t*)readInto = (operand_t)input;
      }
    } else { // We are trying to read a 4-byte value.
      *(unsigned int*)readInto = input;
    }

  } while (valsRead != 1 || !okSize); // See if the input was valid.

}
