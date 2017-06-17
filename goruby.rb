require 'ffi'
require 'webrick'

module GoModule
  extend FFI::Library
  ffi_lib ['go_server', '/home/ldeng/1/work/go/src/github.com/ldeng7/gonginx/libgonginx.so']
  attach_function :setStr, [:string], :void
  attach_function :startServer, [:int], :void
end

GoModule.startServer 8081

server = WEBrick::HTTPServer.new Port: 8080
server.mount_proc '/' do |req, resp|
  GoModule.setStr req.body
  resp.body = 'pid: ' + Process.pid.to_s + "\n"
end
server.start

