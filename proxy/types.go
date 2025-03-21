package main

type Service struct {
	Mapping TcpMapping
}

type TcpMapping struct {
	InternalPort    int
	ExternalPort    int
	InternalAddress string
	ExternalAddress string
}
