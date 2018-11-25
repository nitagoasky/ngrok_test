package conn

import (
	"net"
	"ngrok/log"
	"fmt"
	// "io"
	"ngrok/util"
	// "io/ioutil"
	// "bytes"
)

const (
	NotAuthorized = `HTTP/1.0 401 Not Authorized
WWW-Authenticate: Basic realm="ngrok"
Content-Length: 23

Authorization required
`

	NotFound = `HTTP/1.0 404 Not Found
Content-Length: %d

Tunnel %s not found
`

	BadRequest = `HTTP/1.0 400 Bad Request
Content-Length: 12

Bad Request
`
	success = `HTTP/1.1 200 OK
Content-Length: 8

gaotian
`
)

func Listen(addr string) {


	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch :=make(chan int)
	fmt.Println(listener)
	log.Debug("begin");
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Error("Failed to accept new TCP connection of : %v", err)
			}
			
			conn.Write(util.GetData())

		}
		<-ch
		fmt.Println("end!")
	}()
	ch <- 1
}