package models

type UfwRules struct {
	Status string `json:"Status"`
	Rules  []struct {
		Index           int         `json:"Index"`
		Action          string      `json:"Action"`
		NetworkProtocol string      `json:"Network_protocol"`
		ToPorts         []int       `json:"To_ports"`
		ToTransport     string      `json:"To_transport"`
		FromIP          string      `json:"From_ip"`
		ToPortRanges    interface{} `json:"To_port_ranges"`
	} `json:"Rules"`
}
