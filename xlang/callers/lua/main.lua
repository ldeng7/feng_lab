local ffi = require "ffi"

local funcs = {
	xlangc = "cFun",
	xlanggo = "goFun",
}

for lib_name, fun_name in pairs(funcs) do
    ffi.cdef("int " .. fun_name .. "(void* s, int i);")
    local lib = ffi.load(lib_name)
    local f = lib[fun_name]

    local buf = ffi.new("char[14]", "lua calls xxxx")
    local l = f(buf, 10)
    print(string.sub(ffi.string(buf), 1, l))
end
