#include <stdio.h>
//#include "../../src/CapPicture.h"
#include "CapPicture.h"

extern "C" int testX() {
    printf("Hello world from C++\n");
     Demo_Capture();
    fflush (stdout);
    return 42;
}
