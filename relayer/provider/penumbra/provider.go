package penumbra

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	chantypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/cosmos/relayer/v2/relayer/provider"
	lens "github.com/strangelove-ventures/lens/client"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_ provider.ChainProvider = &PenumbraProvider{}
	_ provider.KeyProvider   = &PenumbraProvider{}
	_ provider.QueryProvider = &PenumbraProvider{}
)

type PenumbraMessage struct {
	// Add message abstraction here
}

func NewPenumbraMessage(msg sdk.Msg) provider.RelayerMessage {
	return PenumbraMessage{}
}

func PenumbraMsg(rm provider.RelayerMessage) sdk.Msg {
	return nil
}

func PenumbraMsgs(rm ...provider.RelayerMessage) []sdk.Msg {
	sdkMsgs := make([]sdk.Msg, 0)
	return sdkMsgs
}

func (cm PenumbraMessage) Type() string {
	return ""
}

func (cm PenumbraMessage) MsgBytes() ([]byte, error) {
	return []byte(""), nil
}

// MarshalLogObject is used to encode cm to a zap logger with the zap.Object field type.
func (cm PenumbraMessage) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return nil
}

type PenumbraProviderConfig struct {
	// Provider Config goes here
}

func (pc PenumbraProviderConfig) Validate() error {
	return nil
}

// NewProvider validates the PenumbraProviderConfig, instantiates a ChainClient and then instantiates a PenumbraProvider
func (pc PenumbraProviderConfig) NewProvider(log *zap.Logger, homepath string, debug bool) (provider.ChainProvider, error) {
	return &PenumbraProvider{}, nil
}

func (pc PenumbraProvider) Init() error {
	return nil
}

// ChainClientConfig builds a ChainClientConfig struct from a PenumbraProviderConfig, this is used
// to instantiate an instance of ChainClient from lens which is how we build the PenumbraProvider
func ChainClientConfig(pcfg *PenumbraProviderConfig) *lens.ChainClientConfig {
	return &lens.ChainClientConfig{}
}

type PenumbraProvider struct {
	//
}

func (cc *PenumbraProvider) ProviderConfig() provider.ProviderConfig {
	return nil
}

func (cc *PenumbraProvider) ChainId() string {
	return ""
}

func (cc *PenumbraProvider) Type() string {
	return "penumbra"
}

func (cc *PenumbraProvider) Key() string {
	return ""
}

func (cc *PenumbraProvider) Timeout() string {
	return ""
}

// Address returns the chains configured address as a string
func (cc *PenumbraProvider) Address() (string, error) {
	return "", nil
}

func (cc *PenumbraProvider) TrustingPeriod(ctx context.Context) (time.Duration, error) {
	return 0, nil
}

// CreateClient creates an sdk.Msg to update the client on src with consensus state from dst
func (cc *PenumbraProvider) CreateClient(clientState ibcexported.ClientState, dstHeader ibcexported.Header) (provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) SubmitMisbehavior( /*TBD*/ ) (provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) UpdateClient(srcClientId string, dstHeader ibcexported.Header) (provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ConnectionOpenInit(srcClientId, dstClientId string, dstHeader ibcexported.Header) ([]provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ConnectionOpenTry(ctx context.Context, dstQueryProvider provider.QueryProvider, dstHeader ibcexported.Header, srcClientId, dstClientId, srcConnId, dstConnId string) ([]provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ConnectionOpenAck(ctx context.Context, dstQueryProvider provider.QueryProvider, dstHeader ibcexported.Header, srcClientId, srcConnId, dstClientId, dstConnId string) ([]provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ConnectionOpenConfirm(ctx context.Context, dstQueryProvider provider.QueryProvider, dstHeader ibcexported.Header, dstConnId, srcClientId, srcConnId string) ([]provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ChannelOpenInit(srcClientId, srcConnId, srcPortId, srcVersion, dstPortId string, order chantypes.Order, dstHeader ibcexported.Header) ([]provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ChannelOpenTry(ctx context.Context, dstQueryProvider provider.QueryProvider, dstHeader ibcexported.Header, srcPortId, dstPortId, srcChanId, dstChanId, srcVersion, srcConnectionId, srcClientId string) ([]provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ChannelOpenAck(ctx context.Context, dstQueryProvider provider.QueryProvider, dstHeader ibcexported.Header, srcClientId, srcPortId, srcChanId, dstChanId, dstPortId string) ([]provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ChannelOpenConfirm(ctx context.Context, dstQueryProvider provider.QueryProvider, dstHeader ibcexported.Header, srcClientId, srcPortId, srcChanId, dstPortId, dstChanId string) ([]provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ChannelCloseInit(srcPortId, srcChanId string) (provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) ChannelCloseConfirm(ctx context.Context, dstQueryProvider provider.QueryProvider, dsth int64, dstChanId, dstPortId, srcPortId, srcChanId string) (provider.RelayerMessage, error) {
	return nil, nil
}

// GetIBCUpdateHeader updates the off chain tendermint light client and
// returns an IBC Update Header which can be used to update an on chain
// light client on the destination chain. The source is used to construct
// the header data.
func (cc *PenumbraProvider) GetIBCUpdateHeader(ctx context.Context, srch int64, dst provider.ChainProvider, dstClientId string) (ibcexported.Header, error) {
	return nil, nil
}

func (cc *PenumbraProvider) GetLightSignedHeaderAtHeight(ctx context.Context, h int64) (ibcexported.Header, error) {
	return nil, nil
}

// InjectTrustedFields injects the necessary trusted fields for a header to update a light
// client stored on the destination chain, using the information provided by the source
// chain.
// TrustedHeight is the latest height of the IBC client on dst
// TrustedValidators is the validator set of srcChain at the TrustedHeight
// InjectTrustedFields returns a copy of the header with TrustedFields modified
func (cc *PenumbraProvider) InjectTrustedFields(ctx context.Context, header ibcexported.Header, dst provider.ChainProvider, dstClientId string) (ibcexported.Header, error) {
	return nil, nil
}

// MsgRelayAcknowledgement constructs the MsgAcknowledgement which is to be sent to the sending chain.
// The counterparty represents the receiving chain where the acknowledgement would be stored.
func (cc *PenumbraProvider) MsgRelayAcknowledgement(ctx context.Context, dst provider.ChainProvider, dstChanId, dstPortId, srcChanId, srcPortId string, dsth int64, packet provider.RelayPacket) (provider.RelayerMessage, error) {
	return nil, nil
}

// MsgTransfer creates a new transfer message
func (cc *PenumbraProvider) MsgTransfer(amount sdk.Coin, dstChainId, dstAddr, srcPortId, srcChanId string, timeoutHeight, timeoutTimestamp uint64) (provider.RelayerMessage, error) {
	return nil, nil
}

// MsgRelayTimeout constructs the MsgTimeout which is to be sent to the sending chain.
// The counterparty represents the receiving chain where the receipts would have been
// stored.
func (cc *PenumbraProvider) MsgRelayTimeout(ctx context.Context, dst provider.ChainProvider, dsth int64, packet provider.RelayPacket, dstChanId, dstPortId, srcChanId, srcPortId string) (provider.RelayerMessage, error) {
	return nil, nil
}

// MsgRelayRecvPacket constructs the MsgRecvPacket which is to be sent to the receiving chain.
// The counterparty represents the sending chain where the packet commitment would be stored.
func (cc *PenumbraProvider) MsgRelayRecvPacket(ctx context.Context, dst provider.ChainProvider, dsth int64, packet provider.RelayPacket, dstChanId, dstPortId, srcChanId, srcPortId string) (provider.RelayerMessage, error) {
	return nil, nil
}

// RelayPacketFromSequence relays a packet with a given seq on src and returns recvPacket msgs, timeoutPacketmsgs and error
func (cc *PenumbraProvider) RelayPacketFromSequence(ctx context.Context, src, dst provider.ChainProvider, srch, dsth, seq uint64, dstChanId, dstPortId, srcChanId, srcPortId, srcClientId string) (provider.RelayerMessage, provider.RelayerMessage, error) {
	return nil, nil, fmt.Errorf("should have errored before here")
}

// AcknowledgementFromSequence relays an acknowledgement with a given seq on src, source is the sending chain, destination is the receiving chain
func (cc *PenumbraProvider) AcknowledgementFromSequence(ctx context.Context, dst provider.ChainProvider, dsth, seq uint64, dstChanId, dstPortId, srcChanId, srcPortId string) (provider.RelayerMessage, error) {
	return nil, nil
}

func (cc *PenumbraProvider) MsgUpgradeClient(srcClientId string, consRes *clienttypes.QueryConsensusStateResponse, clientRes *clienttypes.QueryClientStateResponse) (provider.RelayerMessage, error) {
	return nil, nil
}

// AutoUpdateClient update client automatically to prevent expiry
func (cc *PenumbraProvider) AutoUpdateClient(ctx context.Context, dst provider.ChainProvider, thresholdTime time.Duration, srcClientId, dstClientId string) (time.Duration, error) {
	return 0, nil
}

// FindMatchingClient will determine if there exists a client with identical client and consensus states
// to the client which would have been created. Source is the chain that would be adding a client
// which would track the counterparty. Therefore we query source for the existing clients
// and check if any match the counterparty. The counterparty must have a matching consensus state
// to the latest consensus state of a potential match. The provided client state is the client
// state that will be created if there exist no matches.
func (cc *PenumbraProvider) FindMatchingClient(ctx context.Context, counterparty provider.ChainProvider, clientState ibcexported.ClientState) (string, bool) {
	return "", false
}

func (cc *PenumbraProvider) QueryConsensusStateABCI(ctx context.Context, clientID string, height ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	return &clienttypes.QueryConsensusStateResponse{}, nil
}

// WaitForNBlocks blocks until the next block on a given chain
func (cc *PenumbraProvider) WaitForNBlocks(ctx context.Context, n int64) error {
	return nil
}

// SendMessage attempts to sign, encode & send a RelayerMessage
// This is used extensively in the relayer as an extension of the Provider interface
func (cc *PenumbraProvider) SendMessage(ctx context.Context, msg provider.RelayerMessage) (*provider.RelayerTxResponse, bool, error) {
	return cc.SendMessages(ctx, []provider.RelayerMessage{msg})
}

// SendMessages attempts to sign, encode, & send a slice of RelayerMessages
// This is used extensively in the relayer as an extension of the Provider interface
//
// NOTE: An error is returned if there was an issue sending the transaction. A successfully sent, but failed
// transaction will not return an error. If a transaction is successfully sent, the result of the execution
// of that transaction will be logged. A boolean indicating if a transaction was successfully
// sent and executed successfully is returned.
func (cc *PenumbraProvider) SendMessages(ctx context.Context, msgs []provider.RelayerMessage) (*provider.RelayerTxResponse, bool, error) {
	return &provider.RelayerTxResponse{}, true, nil
}
