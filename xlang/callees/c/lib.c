int cFun(void *s, int i) {
    *(char*)(s + i) = 'c';
    return i + 1;
}
