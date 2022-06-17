package types

type Other struct {
	TxIndex    string `json:"tx_index"`
	RPCAddress string `json:"rpc_address"`
}

type ProtocolVersion struct {
	P2P   string `json:"p2p"`
	Block string `json:"block"`
	App   string `json:"app"`
}

type ApplicationVersion struct {
	Name             string   `json:"name"`
	ServerName       string   `json:"server_name"`
	Version          string   `json:"version"`
	Commit           string   `json:"commit"`
	BuildTags        string   `json:"build_tags"`
	Go               string   `json:"go"`
	BuildDeps        []string `json:"build_deps"`
	CosmosSdkVersion string   `json:"cosmos_sdk_version"`
}

type NodeInfo struct {
	ProtocolVersion ProtocolVersion `json:"protocol_version"`
	ID              string          `json:"-"`
	ListenAddr      string          `json:"-"`
	Network         string          `json:"network"`
	Version         string          `json:"version"`
	Channels        string          `json:"channels"`
	Moniker         string          `json:"moniker"`
	Other           Other           `json:"other"`
}

type Result struct {
	NodeInfo           NodeInfo           `json:"node_info"`
	ApplicationVersion ApplicationVersion `json:"application_version"`
}
