package penumbra

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	commitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
	ibctmtypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	tmclient "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	abci "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// QueryTx takes a transaction hash and returns the transaction
// note we probably need to change the return type here
func (cc *PenumbraProvider) QueryTx(ctx context.Context, hashHex string) (*ctypes.ResultTx, error) {
	return nil, nil
}

// QueryTxs returns an array of transactions given a tag
// note we probably need to change the return type here
func (cc *PenumbraProvider) QueryTxs(ctx context.Context, page, limit int, events []string) ([]*ctypes.ResultTx, error) {
	return nil, nil
}

// QueryBalance returns the amount of coins in the relayer account
func (cc *PenumbraProvider) QueryBalance(ctx context.Context, keyName string) (sdk.Coins, error) {
	return nil, nil
}

// QueryBalanceWithAddress returns the amount of coins in the relayer account with address as input
// TODO add pagination support
func (cc *PenumbraProvider) QueryBalanceWithAddress(ctx context.Context, address string) (sdk.Coins, error) {
	return nil, nil
}

// QueryUnbondingPeriod returns the unbonding period of the chain
func (cc *PenumbraProvider) QueryUnbondingPeriod(ctx context.Context) (time.Duration, error) {
	return time.Hour * 2, nil
}

// QueryTendermintProof performs an ABCI query with the given key and returns
// the value of the query, the proto encoded merkle proof, and the height of
// the Tendermint block containing the state root. The desired tendermint height
// to perform the query should be set in the client context. The query will be
// performed at one below this height (at the IAVL version) in order to obtain
// the correct merkle proof. Proof queries at height less than or equal to 2 are
// not supported. Queries with a client context height of 0 will perform a query
// at the lastest state available.
// Issue: https://github.com/cosmos/cosmos-sdk/issues/6567
func (cc *PenumbraProvider) QueryTendermintProof(ctx context.Context, height int64, key []byte) ([]byte, []byte, clienttypes.Height, error) {
	// ABCI queries at heights 1, 2 or less than or equal to 0 are not supported.
	// Base app does not support queries for height less than or equal to 1.
	// Therefore, a query at height 2 would be equivalent to a query at height 3.
	// A height of 0 will query with the lastest state.
	if height != 0 && height <= 2 {
		return nil, nil, clienttypes.Height{}, fmt.Errorf("proof queries at height <= 2 are not supported")
	}

	// Use the IAVL height if a valid tendermint height is passed in.
	// A height of 0 will query with the latest state.
	if height != 0 {
		height--
	}

	req := abci.RequestQuery{
		Path:   "state/key",
		Height: height,
		Data:   key,
		Prove:  true,
	}

	res, err := cc.QueryABCI(ctx, req)
	if err != nil {
		return nil, nil, clienttypes.Height{}, err
	}

	merkleProof, err := commitmenttypes.ConvertProofs(res.ProofOps)
	if err != nil {
		return nil, nil, clienttypes.Height{}, err
	}

	cdc := codec.NewProtoCodec(cc.Codec.InterfaceRegistry)

	proofBz, err := cdc.Marshal(&merkleProof)
	if err != nil {
		return nil, nil, clienttypes.Height{}, err
	}

	revision := clienttypes.ParseChainID(cc.PCfg.ChainID)
	return res.Value, proofBz, clienttypes.NewHeight(revision, uint64(res.Height)+1), nil
}

// QueryClientStateResponse retrieves the latest consensus state for a client in state at a given height
func (cc *PenumbraProvider) QueryClientStateResponse(ctx context.Context, height int64, srcClientId string) (*clienttypes.QueryClientStateResponse, error) {
	key := []byte(fmt.Sprintf("penumbra-ibc-commitment/clients/%s/clientState", srcClientId))

	value, proofBz, proofHeight, err := cc.QueryTendermintProof(ctx, height, key)
	if err != nil {
		return nil, err
	}

	// check if client exists
	if len(value) == 0 {
		return nil, sdkerrors.Wrap(clienttypes.ErrClientNotFound, srcClientId)
	}

	cdc := codec.NewProtoCodec(cc.Codec.InterfaceRegistry)

	clientState, err := clienttypes.UnmarshalClientState(cdc, value)
	if err != nil {
		return nil, err
	}

	anyClientState, err := clienttypes.PackClientState(clientState)
	if err != nil {
		return nil, err
	}

	return &clienttypes.QueryClientStateResponse{
		ClientState: anyClientState,
		Proof:       proofBz,
		ProofHeight: proofHeight,
	}, nil
}

// QueryClientState retrieves the latest consensus state for a client in state at a given height
// and unpacks it to exported client state interface
func (cc *PenumbraProvider) QueryClientState(ctx context.Context, height int64, clientid string) (ibcexported.ClientState, error) {
	clientStateRes, err := cc.QueryClientStateResponse(ctx, height, clientid)
	if err != nil {
		return nil, err
	}

	clientStateExported, err := clienttypes.UnpackClientState(clientStateRes.ClientState)
	if err != nil {
		return nil, err
	}

	return clientStateExported, nil
}

// QueryClientConsensusState retrieves the latest consensus state for a client in state at a given height
func (cc *PenumbraProvider) QueryClientConsensusState(ctx context.Context, chainHeight int64, clientid string, clientHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	return &clienttypes.QueryConsensusStateResponse{}, nil
}

//DefaultUpgradePath is the default IBC upgrade path set for an on-chain light client

func (cc *PenumbraProvider) NewClientState(dstUpdateHeader ibcexported.Header, dstTrustingPeriod, dstUbdPeriod time.Duration, allowUpdateAfterExpiry, allowUpdateAfterMisbehaviour bool) (ibcexported.ClientState, error) {
	return &tmclient.ClientState{}, nil
}

// QueryUpgradeProof performs an abci query with the given key and returns the proto encoded merkle proof
// for the query and the height at which the proof will succeed on a tendermint verifier.
func (cc *PenumbraProvider) QueryUpgradeProof(ctx context.Context, key []byte, height uint64) ([]byte, clienttypes.Height, error) {
	return []byte(""), clienttypes.Height{}, nil
}

// QueryUpgradedClient returns upgraded client info
func (cc *PenumbraProvider) QueryUpgradedClient(ctx context.Context, height int64) (*clienttypes.QueryClientStateResponse, error) {
	return &clienttypes.QueryClientStateResponse{}, nil
}

// QueryUpgradedConsState returns upgraded consensus state and height of client
func (cc *PenumbraProvider) QueryUpgradedConsState(ctx context.Context, height int64) (*clienttypes.QueryConsensusStateResponse, error) {
	return &clienttypes.QueryConsensusStateResponse{}, nil
}

// QueryConsensusState returns a consensus state for a given chain to be used as a
// client in another chain, fetches latest height when passed 0 as arg
func (cc *PenumbraProvider) QueryConsensusState(ctx context.Context, height int64) (ibcexported.ConsensusState, int64, error) {
	return &ibctmtypes.ConsensusState{}, 0, nil
}

// QueryClients queries all the clients!
// TODO add pagination support
func (cc *PenumbraProvider) QueryClients(ctx context.Context) (clienttypes.IdentifiedClientStates, error) {
	return nil, nil
}

// QueryConnection returns the remote end of a given connection
func (cc *PenumbraProvider) QueryConnection(ctx context.Context, height int64, connectionid string) (*conntypes.QueryConnectionResponse, error) {
	return nil, nil
}

// QueryConnections gets any connections on a chain
// TODO add pagination support
func (cc *PenumbraProvider) QueryConnections(ctx context.Context) (conns []*conntypes.IdentifiedConnection, err error) {
	return nil, nil
}

// QueryConnectionsUsingClient gets any connections that exist between chain and counterparty
// TODO add pagination support
func (cc *PenumbraProvider) QueryConnectionsUsingClient(ctx context.Context, height int64, clientid string) (*conntypes.QueryConnectionsResponse, error) {
	return nil, nil
}

// GenerateConnHandshakeProof generates all the proofs needed to prove the existence of the
// connection state on this chain. A counterparty should use these generated proofs.
func (cc *PenumbraProvider) GenerateConnHandshakeProof(ctx context.Context, height int64, clientId, connId string) (clientState ibcexported.ClientState, clientStateProof []byte, consensusProof []byte, connectionProof []byte, connectionProofHeight ibcexported.Height, err error) {
	return nil, nil, nil, nil, clienttypes.Height{}, nil
}

// QueryChannel returns the channel associated with a channelID
func (cc *PenumbraProvider) QueryChannel(ctx context.Context, height int64, channelid, portid string) (chanRes *chantypes.QueryChannelResponse, err error) {
	return nil, nil
}

func (cc *PenumbraProvider) queryChannelABCI(ctx context.Context, height int64, portID, channelID string) (*chantypes.QueryChannelResponse, error) {
	return nil, nil
}

// QueryChannelClient returns the client state of the client supporting a given channel
func (cc *PenumbraProvider) QueryChannelClient(ctx context.Context, height int64, channelid, portid string) (*clienttypes.IdentifiedClientState, error) {
	return nil, nil
}

// QueryConnectionChannels queries the channels associated with a connection
func (cc *PenumbraProvider) QueryConnectionChannels(ctx context.Context, height int64, connectionid string) ([]*chantypes.IdentifiedChannel, error) {
	return nil, nil
}

// QueryChannels returns all the channels that are registered on a chain
// TODO add pagination support
func (cc *PenumbraProvider) QueryChannels(ctx context.Context) ([]*chantypes.IdentifiedChannel, error) {
	return nil, nil
}

// QueryPacketCommitments returns an array of packet commitments
// TODO add pagination support
func (cc *PenumbraProvider) QueryPacketCommitments(ctx context.Context, height uint64, channelid, portid string) (commitments *chantypes.QueryPacketCommitmentsResponse, err error) {
	return nil, nil
}

// QueryPacketAcknowledgements returns an array of packet acks
// TODO add pagination support
func (cc *PenumbraProvider) QueryPacketAcknowledgements(ctx context.Context, height uint64, channelid, portid string) (acknowledgements []*chantypes.PacketState, err error) {
	return nil, nil
}

// QueryUnreceivedPackets returns a list of unrelayed packet commitments
func (cc *PenumbraProvider) QueryUnreceivedPackets(ctx context.Context, height uint64, channelid, portid string, seqs []uint64) ([]uint64, error) {
	return nil, nil
}

// QueryUnreceivedAcknowledgements returns a list of unrelayed packet acks
func (cc *PenumbraProvider) QueryUnreceivedAcknowledgements(ctx context.Context, height uint64, channelid, portid string, seqs []uint64) ([]uint64, error) {
	return nil, nil
}

// QueryNextSeqRecv returns the next seqRecv for a configured channel
func (cc *PenumbraProvider) QueryNextSeqRecv(ctx context.Context, height int64, channelid, portid string) (recvRes *chantypes.QueryNextSequenceReceiveResponse, err error) {
	return nil, nil
}

// QueryPacketCommitment returns the packet commitment proof at a given height
func (cc *PenumbraProvider) QueryPacketCommitment(ctx context.Context, height int64, channelid, portid string, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	return nil, nil
}

// QueryPacketAcknowledgement returns the packet ack proof at a given height
func (cc *PenumbraProvider) QueryPacketAcknowledgement(ctx context.Context, height int64, channelid, portid string, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	return nil, nil
}

// QueryPacketReceipt returns the packet receipt proof at a given height
func (cc *PenumbraProvider) QueryPacketReceipt(ctx context.Context, height int64, channelid, portid string, seq uint64) (recRes *chantypes.QueryPacketReceiptResponse, err error) {
	return nil, nil
}

func (cc *PenumbraProvider) QueryLatestHeight(ctx context.Context) (int64, error) {
	stat, err := cc.RPCClient.Status(ctx)
	if err != nil {
		return -1, err
	} else if stat.SyncInfo.CatchingUp {
		return -1, fmt.Errorf("node at %s running chain %s not caught up", cc.PCfg.RPCAddr, cc.PCfg.ChainID)
	}
	return stat.SyncInfo.LatestBlockHeight, nil
}

// QueryHeaderAtHeight returns the header at a given height
func (cc *PenumbraProvider) QueryHeaderAtHeight(ctx context.Context, height int64) (ibcexported.Header, error) {
	var (
		page    = 1
		perPage = 100000
	)
	if height <= 0 {
		return nil, fmt.Errorf("must pass in valid height, %d not valid", height)
	}

	res, err := cc.RPCClient.Commit(ctx, &height)
	if err != nil {
		return nil, err
	}

	val, err := cc.RPCClient.Validators(ctx, &height, &page, &perPage)
	if err != nil {
		return nil, err
	}

	protoVal, err := tmtypes.NewValidatorSet(val.Validators).ToProto()
	if err != nil {
		return nil, err
	}

	return &tmclient.Header{
		// NOTE: This is not a SignedHeader
		// We are missing a light.Commit type here
		SignedHeader: res.SignedHeader.ToProto(),
		ValidatorSet: protoVal,
	}, nil
}

// QueryDenomTrace takes a denom from IBC and queries the information about it
func (cc *PenumbraProvider) QueryDenomTrace(ctx context.Context, denom string) (*transfertypes.DenomTrace, error) {
	return nil, nil
}

// QueryDenomTraces returns all the denom traces from a given chain
// TODO add pagination support
func (cc *PenumbraProvider) QueryDenomTraces(ctx context.Context, offset, limit uint64, height int64) ([]transfertypes.DenomTrace, error) {
	return nil, nil
}

func (cc *PenumbraProvider) QueryStakingParams(ctx context.Context) (*stakingtypes.Params, error) {
	return nil, nil
}
