local ffi = require "ffi"
local go_lib = ffi.load("xlanggo")
ffi.cdef[[
    int goFun(int i, void* s);
]]

local call_go = function()
    local buf = ffi.new("char[5]", {})
    print(go_lib.goFun(1, buf))
    print(ffi.string(buf))
end

call_go()
