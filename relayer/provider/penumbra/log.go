package penumbra

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/relayer/v2/relayer/provider"
)

// LogFailedTx takes the transaction and the messages to create it and logs the appropriate data
func (cc *PenumbraProvider) LogFailedTx(res *provider.RelayerTxResponse, err error, msgs []provider.RelayerMessage) {
}

// LogSuccessTx take the transaction and the messages to create it and logs the appropriate data
func (cc *PenumbraProvider) LogSuccessTx(res *sdk.TxResponse, msgs []provider.RelayerMessage) {
}
