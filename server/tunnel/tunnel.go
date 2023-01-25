package tunnel

type Tunnel interface {
	Start()
	Hostname() string
	Protocol() string
	PublicServerPort() uint16
	PrivateServerPort() uint16
}
