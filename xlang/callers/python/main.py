import ctypes

go_lib = ctypes.CDLL("libxlanggo.so")

def call_go():
    buf = ctypes.create_string_buffer(5)
    print(go_lib.goFun(1, buf))
    print(str(buf.value, "utf-8"))

if __name__ == "__main__":
    call_go()
