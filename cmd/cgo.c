#include <stdio.h>
#include <string.h>

#include "../lib/liblogger.h"

int main(int argsc, char **argv) {
    Log(argv[1], 0x04);
    return 0;
}