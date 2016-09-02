#ifndef SHARED_IO_H
#define SHARED_IO_H


#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#include "shared/c/types.h"

void readSanitized(const char *promptString, void *readInto, int numBytes);

#endif
