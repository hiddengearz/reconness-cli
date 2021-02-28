package internal

import "github.com/hiddengearz/reconness-cli/sdk/client"

var (
	//CoreClientCfg *cclient.TransportConfig
	//CoreClient    *cclient.CerberusAPI

	//UserClientCfg *cclient.TransportConfig
	//UserClient    *cclient.CerberusAPI
	ClientCfg   *client.TransportConfig
	Client      *client.ReconnessAPI
	Debug       bool
	AgentMode   bool
	OfflineMode bool
	SilentMode  bool
	Pipe        bool = false
	EncKey      string
	Encryption  bool = false
	Hostname    string
)
