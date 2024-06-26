package api

import (
	"encoding/json"
	"regexp"

	"github.com/xtls/xray-core/infra/conf"
)

const (
	UserNotModified = "users not modified"
	NodeNotModified = "node not modified"
	RuleNotModified = "rules not modified"
)

// Config API config
type Config struct {
	APIHost             string  `mapstructure:"ApiHost"`
	NodeID              int     `mapstructure:"NodeID"`
	Key                 string  `mapstructure:"ApiKey"`
	NodeType            string  `mapstructure:"NodeType"`
	EnableVless         bool    `mapstructure:"EnableVless"`
	VlessFlow           string  `mapstructure:"VlessFlow"`
	Timeout             int     `mapstructure:"Timeout"`
	SpeedLimit          float64 `mapstructure:"SpeedLimit"`
	DeviceLimit         int     `mapstructure:"DeviceLimit"`
	RuleListPath        string  `mapstructure:"RuleListPath"`
}

// NodeStatus Node status
type NodeStatus struct {
	CPU    float64
	Mem    float64
	Disk   float64
	Uptime uint64
}

type NodeInfo struct {
	NodeType          string // Must be V2ray, Trojan, Shadowsocks and Shadowsocks2022
	NodeID            int
	Port              uint32
	SpeedLimit        uint64 // Bps
	AlterID           uint16
	TransportProtocol string
	FakeType          string
	Host              string
	Path              string
	EnableTLS         bool
	EnableVless       bool
	VlessFlow         string
	CipherMethod      string
	ServerKey         string
	ServiceName       string
	Header            json.RawMessage
	NameServerConfig  []*conf.NameServerConfig
	EnableREALITY     bool
	REALITYConfig     *REALITYConfig
	RouteDNS          *RouteDns
}

type UserInfo struct {
	UID         int
	Email       string
	UUID        string
	Passwd      string
	Port        uint32
	AlterID     uint16
	Method      string
	SpeedLimit  uint64 // Bps
	DeviceLimit int
}

type OnlineUser struct {
	UID int
	IP  string
}

type UserTraffic struct {
	UID      int
	Email    string
	Upload   int64
	Download int64
}

type ClientInfo struct {
	APIHost  string
	NodeID   int
	Key      string
	NodeType string
}

type DetectRule struct {
	ID      int
	Pattern *regexp.Regexp
}

type DetectResult struct {
	UID    int
	RuleID int
}

type REALITYConfig struct {
	Dest             string
	ProxyProtocolVer uint64
	ServerNames      []string
	PrivateKey       string
	MinClientVer     string
	MaxClientVer     string
	MaxTimeDiff      uint64
	ShortIds         []string
}

type RouteDns struct {
	Allow    RouteDnsConfig     `json:"allow"`
	Block    RouteDnsConfig     `json:"block"`
	AllowIPs RouteDnsConfig     `json:"allow-ip"`
	Socks5   Socks5PanelOptions `json:"socks"`
	Spoof    []string           `json:"spoof"`
}

type Socks5PanelOptions struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	LocalAddr string `json:"local-addr"`

	// When the resolver is configured with a name, not an IP, e.g. one.one.one.one:53
	// this setting will resolve that name locally rather than on the SOCKS proxy. The
	// name will be resolved either on the local system, or via the bootstrap-resolver
	// if one is setup.
	ResolveLocal  bool                   `json:"resolve-local"`
	Socks5Address string                 `json:"address"`
}

type RouteDnsConfig struct {
	Mode     string             `json:"type"`
	Domains  []string           `json:"domains"`
	Hosts    *conf.HostsWrapper `json:"hosts"`
	AllowIPs conf.StringList    `json:"ips"`
}
