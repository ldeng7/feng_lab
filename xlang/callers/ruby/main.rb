require 'ffi'

module GoLib
  extend FFI::Library
  ffi_lib ['go_lib', 'xlanggo']
  attach_function :goFun, [:int, :pointer], :int
end

def call_go()
  buf = FFI::MemoryPointer.new :char, 5
  puts GoLib.goFun 1, buf
  puts buf.read_string
end

call_go
