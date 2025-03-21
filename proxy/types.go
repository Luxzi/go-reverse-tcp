package proxy

type Service struct {
	Name    string
	Mapping TcpMapping
}

type TcpMapping struct {
	InternalPort    int
	ExternalPort    int
	InternalAddress string
	ExternalAddress string
}
