require 'ffi'

$funcs = {
  xlangc: :cFun,
  xlanggo: :goFun,
}

module FLib
  extend FFI::Library
  $funcs.each do |lib_name, fun_name|
    ffi_lib lib_name
    attach_function fun_name, [:pointer, :int], :int
  end
end

$funcs.each do |_, fun_name|
  buf = FFI::MemoryPointer.from_string "ruby calls xxxx"
  l = FLib.send fun_name, buf, 11
  puts buf.read_string[0...l]
end
