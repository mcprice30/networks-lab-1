// Lab 1 - Group 11
// 
// io.h contains helper functionality for reading user input.
//
#ifndef SHARED_IO_H
#define SHARED_IO_H


#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#include "shared/c/types.h"

// readSanitized takes a prompt string, a pointer to read into, and the number
// of bytes allocated to that pointer, and will repeatedly prompt the user
// with the given prompt string until a valid value is provided, which will
// be read into the given pointer.
void readSanitized(const char *promptString, void *readInto, int numBytes);

#endif
