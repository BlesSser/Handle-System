package listener

import (
	"fmt"
	"net"
	"os"
)

//todo
//Start TCP Service
func StartTCPService() (err error) {
	// listen to incoming tcp connections
	l, err := net.Listen("tcp", "host:port")
	if err != nil {
		return err
	}
	defer l.Close()
	// A common pattern is to start a loop to continously accept connections
	for {
		//accept connections using Listener.Accept()
		c, err := l.Accept()
		if err != nil {
			return err
		}
		//It's common to handle accepted connection on different goroutines
		go HandleRequest(c)
	}
}

//todo
//Start UDP listener
func StartUDPService() (err error) {

	addr, err := net.ResolveUDPAddr("udp", "host:port")
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		HandleRequest(conn)
	}

}
