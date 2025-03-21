package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func dispatchServiceProxy(address string, port int) {
	proxy, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Fatal(err)
	}
	defer proxy.Close()

	collectConnections(proxy)
}

func collectConnections(proxy net.Listener) {
	for {
		connection, err := proxy.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())
	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)
	defer connection.Close()
	for {
		_, err := connection.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			println("END OF FILE")
			break
		}
		packet = append(packet, tmp...)

		fmt.Println(packet)
	}
	num, _ := connection.Write(packet)
	fmt.Printf("Wrote back %d bytes, the payload is %s\n", num, string(packet))
	connection.Close()
}
