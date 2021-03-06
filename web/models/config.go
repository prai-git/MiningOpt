package models

type (
	WebType int

	WebConfig struct {
		EtcdServers []string
		HttpPort    uint16
		RpcxPort    uint16
		StartType   WebType
		MyIpAddr    string
	}
)

const (
	WebType_Manager WebType = iota
	WebType_Worker
)
