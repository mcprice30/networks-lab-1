/*
** talker.c -- a datagram "client" demo
*/

#include "shared/c/io.h"

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
