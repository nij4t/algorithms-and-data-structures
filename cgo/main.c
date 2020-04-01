#include <stdio.h>
#include <string.h>

#include "liblogger.h"

int main(int argsc, char **argv) {
    char *msg = "Something";
    Log(msg, 0x04);
    return 0;
}