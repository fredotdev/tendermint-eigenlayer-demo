package main

import (
	"log"
	"os"
	"encoding/hex"
	"errors"
	"strconv"
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"


	abcitypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/ed25519"

	blsregistry "github.com/me/example/abi"
)

type TendermintApplication struct {
	ethClient *ethclient.Client
	contract  *blsregistry.BLSRegistryCoordinatorWithIndices
	queue []ed25519.PubKey
}

var _ abcitypes.Application = (*TendermintApplication)(nil)

// Define a constant for transaction types
const (
	TypeVerifySignature = "verify_signature"
	// Define other transaction types here
)

// TxPayload represents a common payload for transactions
type TxPayload struct {
	Type string `json:"type"`
	Data json.RawMessage `json:"data"`
}

// TxVerifySignature represents the data needed to verify a signature
type TxVerifySignature struct {
	From      string `json:"from"`
	Message   string `json:"message"`
	Signature string `json:"signature"`
	PubKey    ed25519.PubKey `json:"pubKey"` // This would be the ed25519 public key as a hex string
}

func (app *TendermintApplication) CallGetOperator(address common.Address) (uint8, error) {    
    operatorStatus, err := app.contract.GetOperator(nil, address)
    if err != nil {
        log.Printf("Failed to call GetOperator: %v", err)
        return operatorStatus.Status, err 
    }

    status := operatorStatus.Status

    return status, nil 
}

func NewContractInstance(client *ethclient.Client, contractABI abi.ABI, contractAddress common.Address) (*bind.BoundContract, error) {
	contractInstance := bind.NewBoundContract(contractAddress, contractABI, client, client, client)
	return contractInstance, nil
}

func NewTendermintApplication(ethereumNodeURL string) *TendermintApplication {
    // Connect to the Ethereum node
    client, err := ethclient.Dial(ethereumNodeURL)
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    // Get the contract address from the environment variable
    contractAddress := os.Getenv("CONTRACT_ADDRESS")
    if contractAddress == "" {
        log.Fatal("CONTRACT_ADDRESS environment variable is not set")
    }
    address := common.HexToAddress(contractAddress)

    // Use the new name 'blsregistry' to access the package's New function
    contractInstance, err := blsregistry.NewBLSRegistryCoordinatorWithIndices(address, client)
    if err != nil {
        log.Fatalf("Failed to instantiate the contract: %v", err)
    }

    return &TendermintApplication{
        ethClient: client,
        contract:  contractInstance,
		queue: []ed25519.PubKey{},
    }
}

func (TendermintApplication) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (TendermintApplication) SetOption(req abcitypes.RequestSetOption) abcitypes.ResponseSetOption {
	return abcitypes.ResponseSetOption{}
}

func (app *TendermintApplication) VerifySignature(address common.Address, message, sigHex string) (bool, error) {
	sig, err := hex.DecodeString(sigHex)
	if err != nil {
		return false, err
	}

	if len(sig) != 65 {
		return false, errors.New("invalid signature length")
	}

	prefix := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)))
	prefixedMsg := append(prefix, []byte(message)...)

	pubKey, err := crypto.SigToPub(crypto.Keccak256(prefixedMsg), sig)
	if err != nil {
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return recoveredAddr == address, nil
}

func (app *TendermintApplication) handleVerifySignatureTx(data json.RawMessage) abcitypes.ResponseDeliverTx {
	var tx TxVerifySignature
	err := json.Unmarshal(data, &tx)
	if err != nil {
		return abcitypes.ResponseDeliverTx{Code: 1, Log: "verify signature tx decode error"}
	}

	address := common.HexToAddress(tx.From)
	valid, err := app.VerifySignature(address, tx.Message, tx.Signature)
	if err != nil || !valid {
		return abcitypes.ResponseDeliverTx{Code: 2, Log: "invalid signature"}
	}

	status, err := app.CallGetOperator(common.HexToAddress(tx.From))
	if err != nil || status == 0 {
		return abcitypes.ResponseDeliverTx{Code: 2, Log: "invalid ethereum address"}
	}

	app.queue = append(app.queue, tx.PubKey)
	if err != nil {
		return abcitypes.ResponseDeliverTx{Code: 3, Log: "error adding to validator queue"}
	}

	return abcitypes.ResponseDeliverTx{Code: 0}
}

func (app *TendermintApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	// First, unmarshal the transaction into the common payload structure
	var txPayload TxPayload
	err := json.Unmarshal(req.Tx, &txPayload)
	if err != nil {
		return abcitypes.ResponseDeliverTx{Code: 1, Log: "tx decode error"}
	}

	// Dispatch to the appropriate handler based on the transaction type
	switch txPayload.Type {
	case TypeVerifySignature:
		return app.handleVerifySignatureTx(txPayload.Data)
	// Add cases for other transaction types here
	default:
		return abcitypes.ResponseDeliverTx{Code: 1, Log: "unknown tx type"}
	}
}

func (TendermintApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	return abcitypes.ResponseCheckTx{Code: 0}
}

func (TendermintApplication) Commit() abcitypes.ResponseCommit {
	return abcitypes.ResponseCommit{}
}

func (TendermintApplication) Query(req abcitypes.RequestQuery) abcitypes.ResponseQuery {
	return abcitypes.ResponseQuery{Code: 0}
}

func (TendermintApplication) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (TendermintApplication) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return abcitypes.ResponseBeginBlock{}
}

func (app *TendermintApplication) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	var validatorUpdates []abcitypes.ValidatorUpdate

	for _, tmPubKey := range app.queue {
		// Convert the ed25519 public key to a byte representation
		pubKeyBytes := tmPubKey.Bytes()

		// Now create the ValidatorUpdate using the Tendermint helper function
		validatorUpdate := abcitypes.UpdateValidator(pubKeyBytes, 10, ed25519.KeyType)

		// Append to the validatorUpdates slice
		validatorUpdates = append(validatorUpdates, validatorUpdate)
	}

	// Clear the queue after processing
	app.queue = nil // Reset the queue

	return abcitypes.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
	}
}

func (TendermintApplication) ListSnapshots(abcitypes.RequestListSnapshots) abcitypes.ResponseListSnapshots {
	return abcitypes.ResponseListSnapshots{}
}

func (TendermintApplication) OfferSnapshot(abcitypes.RequestOfferSnapshot) abcitypes.ResponseOfferSnapshot {
	return abcitypes.ResponseOfferSnapshot{}
}

func (TendermintApplication) LoadSnapshotChunk(abcitypes.RequestLoadSnapshotChunk) abcitypes.ResponseLoadSnapshotChunk {
	return abcitypes.ResponseLoadSnapshotChunk{}
}

func (TendermintApplication) ApplySnapshotChunk(abcitypes.RequestApplySnapshotChunk) abcitypes.ResponseApplySnapshotChunk {
	return abcitypes.ResponseApplySnapshotChunk{}
}
