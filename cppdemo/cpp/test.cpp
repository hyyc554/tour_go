#include "test.h"
void Test::call() {
    printf("call from c++ language\n");
}
cwrap.cpp
#include "cwrap.h"
#include "test.h"
void call() {
    Test ctx;
    ctx.call();
}