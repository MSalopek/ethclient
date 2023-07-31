package params

import (
	"strings"

	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
)

var (
	// EthereumRpcURL defines the configuration key for Ethereum RPC URL.
	//nolint: gosec
	EthereumRpcURL = "ethereum-rpc-url"

	// custom application configuration TOML template.
	customTemplate = `
###############################################################################
###                        Custom ETH Configuration                         ###
###############################################################################
# "ethereum-rpc-url defines a URL of an ethereum RPC node to
# it is used to query eth state and proofs
#
ethereum-rpc-url = "{{ .EthereumRpcURL }}"
`
)

// CustomConfigTemplate defines a custom application configuration TOML
func CustomConfigTemplate() string {
	config := serverconfig.DefaultConfigTemplate
	lines := strings.Split(config, "\n")
	// add config at the second line of the file
	lines[2] += customTemplate
	return strings.Join(lines, "\n")
}

// CustomAppConfig defines Gaia's custom application configuration.
type CustomAppConfig struct {
	serverconfig.Config

	// BypassMinFeeMsgTypes defines custom message types the operator may set that
	// will bypass minimum fee checks during CheckTx.
	ethereumRpcURL []string `mapstructure:"ethereum-rpc-url"`
}
