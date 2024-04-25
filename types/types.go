package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/planetmint/faucet/config"
)

type IndexParams struct {
	RecvAddr string
	Tx       sdk.TxResponse
	Config   *config.Config
}

type SendParams struct {
	Address []string `mapstructure:"address"`
}
