#include "goc.h"

int bridge_int_func(initFunc f)
{
    return f();
}

int forty_two()
{
    return 42;
}