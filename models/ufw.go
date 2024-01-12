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

type AddUfwRule struct {
	Source struct {
		Type   string        `json:"type"`
		IP     string        `json:"ip"`
		Range  []interface{} `json:"range"`
		Subnet struct {
			IP     string `json:"ip"`
			Prefix string `json:"prefix"`
		} `json:"subnet"`
	} `json:"source"`
	Port struct {
		Type   string        `json:"type"`
		Number string        `json:"number"`
		Range  []interface{} `json:"range"`
	} `json:"port"`
	Protocol string `json:"protocol"`
	Action   string `json:"action"`
}

type DeleteUfwRule struct {
	Index []int `json:"index"`
}
