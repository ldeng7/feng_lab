import ctypes

funcs = {
	"xlangc": "cFun",
	"xlanggo": "goFun",
    "xlangrust": "rustFun",
}

for lib_name in funcs:
    fun_name = funcs[lib_name]
    lib = ctypes.CDLL("lib" + lib_name + ".so")
    f = lib.__getattr__(fun_name)

    buf = ctypes.create_string_buffer(b"python calls xxxx")
    l = f(buf, 13)
    print(str(buf.value[:l], "utf-8"))
