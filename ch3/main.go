package main

import (
	"bytes"
	"flag"
	"net"
	"net/http"
	"strconv"
	"sync"
	"syscall"
)

func makeResp(payload []byte) []byte {
	buf := bytes.NewBufferString("HTTP/1.1 200 OK\r\nContent-Length: ")
	buf.WriteString(strconv.Itoa(len(payload)))
	buf.WriteString("\r\n\r\n")
	buf.Write(payload)
	return buf.Bytes()
}

type HttpServer struct{}

func (s *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method[:3]))
}

func runHttpServer(port int) {
	println("runHttpServer")
	http.ListenAndServe(":"+strconv.Itoa(port), &HttpServer{})
}

func runTcp1Server(port int) {
	println("runTcp1Server")
	l, _ := net.ListenTCP("tcp4", &net.TCPAddr{Port: port})
	rr := make([]byte, 1024)
	chs := [4]chan *net.TCPConn{}
	for i := 0; i < 4; i++ {
		chs[i] = make(chan *net.TCPConn)
	}

	go func() {
		var i uint8
		for {
			conn, _ := l.AcceptTCP()
			chs[i&3] <- conn
			i++
		}
	}()

	for i := 0; i < 4; i++ {
		go func(i int) {
			for conn := range chs[i] {
				go func(conn *net.TCPConn) {
					r := make([]byte, 3)
					for {
						n, _ := conn.Read(r)
						if n == 0 {
							break
						}
						conn.Read(rr)
						conn.Write(makeResp(r))
					}
					conn.Close()
				}(conn)
			}
		}(i)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func runTcp2Server(port int) {
	println("runTcp2Server")
	lfd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	syscall.SetsockoptInt(lfd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	syscall.Bind(lfd, &syscall.SockaddrInet4{Port: port})
	syscall.Listen(lfd, 42)
	rr := make([]byte, 1024)
	efds := [4]int{}
	for i := 0; i < 4; i++ {
		efds[i], _ = syscall.EpollCreate(42)
	}

	go func() {
		var i uint8
		for {
			cfd, _, _ := syscall.Accept(lfd)
			ev := &syscall.EpollEvent{
				Events: syscall.EPOLLIN,
				Fd:     int32(cfd),
			}
			syscall.EpollCtl(efds[i&3], syscall.EPOLL_CTL_ADD, cfd, ev)
			i++
		}
	}()

	for i := 0; i < 4; i++ {
		go func(i int) {
			evs := make([]syscall.EpollEvent, 128)
			for {
				n, _ := syscall.EpollWait(efds[i], evs, 0)
				for j := 0; j < n; j++ {
					cfd := int(evs[j].Fd)
					r := make([]byte, 3)
					nr, _ := syscall.Read(cfd, r)
					if nr == 0 {
						syscall.EpollCtl(efds[i], syscall.EPOLL_CTL_DEL, cfd, nil)
						syscall.Close(cfd)
						continue
					}
					syscall.Read(cfd, rr)
					syscall.Write(cfd, makeResp(r))
				}
			}
		}(i)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func main() {
	typ := flag.Int("t", 0, "")
	port := flag.Int("p", 8080, "")
	flag.Parse()
	switch *typ {
	case 0:
		runHttpServer(*port)
	case 1:
		runTcp1Server(*port)
	case 2:
		runTcp2Server(*port)
	}
}
