#ifndef SHARED_IO_H
#define SHARED_IO_H


#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
/*
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include "shared/types.h" */

void readSanitized(const char *promptString, void *readInto, int numBytes);

#endif
