package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
)

func CreateService(name string, internalAddress string, externalAddress string, internalPort int, externalPort int) {
	service := Service{
		Name: name,
		Mapping: TcpMapping{
			InternalPort:    internalPort,
			ExternalPort:    externalPort,
			InternalAddress: internalAddress,
			ExternalAddress: externalAddress,
		},
	}

	dispatch(service)
}

func dispatch(service Service) {
	proxy, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", service.Mapping.ExternalAddress, service.Mapping.ExternalPort))
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
