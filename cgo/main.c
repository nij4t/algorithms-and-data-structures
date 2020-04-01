#include <stdio.h>
#include <string.h>

#include "liblogger.h"

int main(int argsc, char **argv) {
    Log(argv[1], 0x04);
    return 0;
}