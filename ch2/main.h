#ifndef MAIN_H
#define MAIN_H

#include <stdint.h>

extern char* runHandler(uint64_t, uint64_t);
char* runHandler(uint64_t f, uint64_t r) {
	((void (*)(uint64_t)) f) (r);
}

#endif
