// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BLSRegistryCoordinatorWithIndices

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// IBLSRegistryCoordinatorWithIndicesOperatorKickParam is an auto generated low-level Go binding around an user-defined struct.
type IBLSRegistryCoordinatorWithIndicesOperatorKickParam struct {
	Operator          common.Address
	Pubkey            BN254G1Point
	OperatorIdsToSwap [][32]byte
}

// IBLSRegistryCoordinatorWithIndicesOperatorSetParam is an auto generated low-level Go binding around an user-defined struct.
type IBLSRegistryCoordinatorWithIndicesOperatorSetParam struct {
	MaxOperatorCount        uint32
	KickBIPsOfOperatorStake uint16
	KickBIPsOfTotalStake    uint16
}

// IRegistryCoordinatorOperator is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorOperator struct {
	OperatorId [32]byte
	Status     uint8
}

// IRegistryCoordinatorQuorumBitmapUpdate is an auto generated low-level Go binding around an user-defined struct.
type IRegistryCoordinatorQuorumBitmapUpdate struct {
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
	QuorumBitmap          *big.Int
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// BLSRegistryCoordinatorWithIndicesMetaData contains all meta data concerning the BLSRegistryCoordinatorWithIndices contract.
var BLSRegistryCoordinatorWithIndicesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISlasher\",\"name\":\"_slasher\",\"type\":\"address\"},{\"internalType\":\"contractIServiceManager\",\"name\":\"_serviceManager\",\"type\":\"address\"},{\"internalType\":\"contractIStakeRegistry\",\"name\":\"_stakeRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIBLSPubkeyRegistry\",\"name\":\"_blsPubkeyRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIIndexRegistry\",\"name\":\"_indexRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"OperatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"OperatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIBLSRegistryCoordinatorWithIndices.OperatorSetParam\",\"name\":\"operatorSetParams\",\"type\":\"tuple\"}],\"name\":\"OperatorSetParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"OperatorSocketUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"log_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"log_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"log_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"log_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"log_named_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"val\",\"type\":\"uint256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256[]\",\"name\":\"val\",\"type\":\"int256[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"val\",\"type\":\"address[]\"}],\"name\":\"log_named_array\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"val\",\"type\":\"bytes\"}],\"name\":\"log_named_bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"}],\"name\":\"log_named_bytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"name\":\"log_named_decimal_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"name\":\"log_named_int\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"log_named_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"log_named_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"log_string\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"log_uint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"logs\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"IS_TEST\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsPubkeyRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSPubkeyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"deregistrationData\",\"type\":\"bytes\"}],\"name\":\"deregisterOperatorWithCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"operatorIdsToSwap\",\"type\":\"bytes32[]\"}],\"name\":\"deregisterOperatorWithCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"excludedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"excludeSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"excludedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"getCurrentQuorumBitmapByOperatorId\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"enumIRegistryCoordinator.OperatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIRegistryCoordinator.Operator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getOperatorSetParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIBLSRegistryCoordinatorWithIndices.OperatorSetParam\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getQuorumBitmapByOperatorIdAtBlockNumberByIndex\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"operatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getQuorumBitmapIndicesByOperatorIdsAtBlockNumber\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getQuorumBitmapUpdateByOperatorIdByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"quorumBitmap\",\"type\":\"uint192\"}],\"internalType\":\"structIRegistryCoordinator.QuorumBitmapUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"}],\"name\":\"getQuorumBitmapUpdateByOperatorIdLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"indexRegistry\",\"outputs\":[{\"internalType\":\"contractIIndexRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIBLSRegistryCoordinatorWithIndices.OperatorSetParam[]\",\"name\":\"_operatorSetParams\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numRegistries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"registrationData\",\"type\":\"bytes\"}],\"name\":\"registerOperatorWithCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkey\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"operatorIdsToSwap\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIBLSRegistryCoordinatorWithIndices.OperatorKickParam[]\",\"name\":\"operatorKickParams\",\"type\":\"tuple[]\"}],\"name\":\"registerOperatorWithCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkey\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"registerOperatorWithCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registries\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceManager\",\"outputs\":[{\"internalType\":\"contractIServiceManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxOperatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfOperatorStake\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"kickBIPsOfTotalStake\",\"type\":\"uint16\"}],\"internalType\":\"structIBLSRegistryCoordinatorWithIndices.OperatorSetParam\",\"name\":\"operatorSetParam\",\"type\":\"tuple\"}],\"name\":\"setOperatorSetParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"contractISlasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifactSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetArtifacts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"targetedArtifacts_\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedContracts_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSelectors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targetedSenders_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"socket\",\"type\":\"string\"}],\"name\":\"updateSocket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BLSRegistryCoordinatorWithIndicesABI is the input ABI used to generate the binding from.
// Deprecated: Use BLSRegistryCoordinatorWithIndicesMetaData.ABI instead.
var BLSRegistryCoordinatorWithIndicesABI = BLSRegistryCoordinatorWithIndicesMetaData.ABI

// BLSRegistryCoordinatorWithIndices is an auto generated Go binding around an Ethereum contract.
type BLSRegistryCoordinatorWithIndices struct {
	BLSRegistryCoordinatorWithIndicesCaller     // Read-only binding to the contract
	BLSRegistryCoordinatorWithIndicesTransactor // Write-only binding to the contract
	BLSRegistryCoordinatorWithIndicesFilterer   // Log filterer for contract events
}

// BLSRegistryCoordinatorWithIndicesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BLSRegistryCoordinatorWithIndicesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSRegistryCoordinatorWithIndicesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BLSRegistryCoordinatorWithIndicesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSRegistryCoordinatorWithIndicesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BLSRegistryCoordinatorWithIndicesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BLSRegistryCoordinatorWithIndicesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BLSRegistryCoordinatorWithIndicesSession struct {
	Contract     *BLSRegistryCoordinatorWithIndices // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BLSRegistryCoordinatorWithIndicesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BLSRegistryCoordinatorWithIndicesCallerSession struct {
	Contract *BLSRegistryCoordinatorWithIndicesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// BLSRegistryCoordinatorWithIndicesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BLSRegistryCoordinatorWithIndicesTransactorSession struct {
	Contract     *BLSRegistryCoordinatorWithIndicesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// BLSRegistryCoordinatorWithIndicesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BLSRegistryCoordinatorWithIndicesRaw struct {
	Contract *BLSRegistryCoordinatorWithIndices // Generic contract binding to access the raw methods on
}

// BLSRegistryCoordinatorWithIndicesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BLSRegistryCoordinatorWithIndicesCallerRaw struct {
	Contract *BLSRegistryCoordinatorWithIndicesCaller // Generic read-only contract binding to access the raw methods on
}

// BLSRegistryCoordinatorWithIndicesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BLSRegistryCoordinatorWithIndicesTransactorRaw struct {
	Contract *BLSRegistryCoordinatorWithIndicesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBLSRegistryCoordinatorWithIndices creates a new instance of BLSRegistryCoordinatorWithIndices, bound to a specific deployed contract.
func NewBLSRegistryCoordinatorWithIndices(address common.Address, backend bind.ContractBackend) (*BLSRegistryCoordinatorWithIndices, error) {
	contract, err := bindBLSRegistryCoordinatorWithIndices(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndices{BLSRegistryCoordinatorWithIndicesCaller: BLSRegistryCoordinatorWithIndicesCaller{contract: contract}, BLSRegistryCoordinatorWithIndicesTransactor: BLSRegistryCoordinatorWithIndicesTransactor{contract: contract}, BLSRegistryCoordinatorWithIndicesFilterer: BLSRegistryCoordinatorWithIndicesFilterer{contract: contract}}, nil
}

// NewBLSRegistryCoordinatorWithIndicesCaller creates a new read-only instance of BLSRegistryCoordinatorWithIndices, bound to a specific deployed contract.
func NewBLSRegistryCoordinatorWithIndicesCaller(address common.Address, caller bind.ContractCaller) (*BLSRegistryCoordinatorWithIndicesCaller, error) {
	contract, err := bindBLSRegistryCoordinatorWithIndices(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesCaller{contract: contract}, nil
}

// NewBLSRegistryCoordinatorWithIndicesTransactor creates a new write-only instance of BLSRegistryCoordinatorWithIndices, bound to a specific deployed contract.
func NewBLSRegistryCoordinatorWithIndicesTransactor(address common.Address, transactor bind.ContractTransactor) (*BLSRegistryCoordinatorWithIndicesTransactor, error) {
	contract, err := bindBLSRegistryCoordinatorWithIndices(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesTransactor{contract: contract}, nil
}

// NewBLSRegistryCoordinatorWithIndicesFilterer creates a new log filterer instance of BLSRegistryCoordinatorWithIndices, bound to a specific deployed contract.
func NewBLSRegistryCoordinatorWithIndicesFilterer(address common.Address, filterer bind.ContractFilterer) (*BLSRegistryCoordinatorWithIndicesFilterer, error) {
	contract, err := bindBLSRegistryCoordinatorWithIndices(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesFilterer{contract: contract}, nil
}

// bindBLSRegistryCoordinatorWithIndices binds a generic wrapper to an already deployed contract.
func bindBLSRegistryCoordinatorWithIndices(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BLSRegistryCoordinatorWithIndicesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSRegistryCoordinatorWithIndices.Contract.BLSRegistryCoordinatorWithIndicesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.BLSRegistryCoordinatorWithIndicesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.BLSRegistryCoordinatorWithIndicesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BLSRegistryCoordinatorWithIndices.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) ISTEST() (bool, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ISTEST(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) ISTEST() (bool, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ISTEST(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) BlsPubkeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "blsPubkeyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) BlsPubkeyRegistry() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.BlsPubkeyRegistry(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// BlsPubkeyRegistry is a free data retrieval call binding the contract method 0x3561deb1.
//
// Solidity: function blsPubkeyRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) BlsPubkeyRegistry() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.BlsPubkeyRegistry(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) ExcludeArtifacts() ([]string, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ExcludeArtifacts(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) ExcludeArtifacts() ([]string, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ExcludeArtifacts(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) ExcludeContracts() ([]common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ExcludeContracts(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ExcludeContracts(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) ExcludeSenders() ([]common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ExcludeSenders(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ExcludeSenders(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// GetCurrentQuorumBitmapByOperatorId is a free data retrieval call binding the contract method 0x3431af25.
//
// Solidity: function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) view returns(uint192)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) GetCurrentQuorumBitmapByOperatorId(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "getCurrentQuorumBitmapByOperatorId", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentQuorumBitmapByOperatorId is a free data retrieval call binding the contract method 0x3431af25.
//
// Solidity: function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) view returns(uint192)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) GetCurrentQuorumBitmapByOperatorId(operatorId [32]byte) (*big.Int, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetCurrentQuorumBitmapByOperatorId(&_BLSRegistryCoordinatorWithIndices.CallOpts, operatorId)
}

// GetCurrentQuorumBitmapByOperatorId is a free data retrieval call binding the contract method 0x3431af25.
//
// Solidity: function getCurrentQuorumBitmapByOperatorId(bytes32 operatorId) view returns(uint192)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) GetCurrentQuorumBitmapByOperatorId(operatorId [32]byte) (*big.Int, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetCurrentQuorumBitmapByOperatorId(&_BLSRegistryCoordinatorWithIndices.CallOpts, operatorId)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) GetOperator(opts *bind.CallOpts, operator common.Address) (IRegistryCoordinatorOperator, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "getOperator", operator)

	if err != nil {
		return *new(IRegistryCoordinatorOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorOperator)).(*IRegistryCoordinatorOperator)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperator, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetOperator(&_BLSRegistryCoordinatorWithIndices.CallOpts, operator)
}

// GetOperator is a free data retrieval call binding the contract method 0x5865c60c.
//
// Solidity: function getOperator(address operator) view returns((bytes32,uint8))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) GetOperator(operator common.Address) (IRegistryCoordinatorOperator, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetOperator(&_BLSRegistryCoordinatorWithIndices.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetOperatorId(&_BLSRegistryCoordinatorWithIndices.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetOperatorId(&_BLSRegistryCoordinatorWithIndices.CallOpts, operator)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) GetOperatorSetParams(opts *bind.CallOpts, quorumNumber uint8) (IBLSRegistryCoordinatorWithIndicesOperatorSetParam, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "getOperatorSetParams", quorumNumber)

	if err != nil {
		return *new(IBLSRegistryCoordinatorWithIndicesOperatorSetParam), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSRegistryCoordinatorWithIndicesOperatorSetParam)).(*IBLSRegistryCoordinatorWithIndicesOperatorSetParam)

	return out0, err

}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) GetOperatorSetParams(quorumNumber uint8) (IBLSRegistryCoordinatorWithIndicesOperatorSetParam, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetOperatorSetParams(&_BLSRegistryCoordinatorWithIndices.CallOpts, quorumNumber)
}

// GetOperatorSetParams is a free data retrieval call binding the contract method 0xe65797ad.
//
// Solidity: function getOperatorSetParams(uint8 quorumNumber) view returns((uint32,uint16,uint16))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) GetOperatorSetParams(quorumNumber uint8) (IBLSRegistryCoordinatorWithIndicesOperatorSetParam, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetOperatorSetParams(&_BLSRegistryCoordinatorWithIndices.CallOpts, quorumNumber)
}

// GetQuorumBitmapByOperatorIdAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x3064620d.
//
// Solidity: function getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) GetQuorumBitmapByOperatorIdAtBlockNumberByIndex(opts *bind.CallOpts, operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "getQuorumBitmapByOperatorIdAtBlockNumberByIndex", operatorId, blockNumber, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapByOperatorIdAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x3064620d.
//
// Solidity: function getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) GetQuorumBitmapByOperatorIdAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetQuorumBitmapByOperatorIdAtBlockNumberByIndex(&_BLSRegistryCoordinatorWithIndices.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapByOperatorIdAtBlockNumberByIndex is a free data retrieval call binding the contract method 0x3064620d.
//
// Solidity: function getQuorumBitmapByOperatorIdAtBlockNumberByIndex(bytes32 operatorId, uint32 blockNumber, uint256 index) view returns(uint192)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) GetQuorumBitmapByOperatorIdAtBlockNumberByIndex(operatorId [32]byte, blockNumber uint32, index *big.Int) (*big.Int, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetQuorumBitmapByOperatorIdAtBlockNumberByIndex(&_BLSRegistryCoordinatorWithIndices.CallOpts, operatorId, blockNumber, index)
}

// GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber is a free data retrieval call binding the contract method 0x85020d49.
//
// Solidity: function getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber(opts *bind.CallOpts, blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "getQuorumBitmapIndicesByOperatorIdsAtBlockNumber", blockNumber, operatorIds)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber is a free data retrieval call binding the contract method 0x85020d49.
//
// Solidity: function getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber(&_BLSRegistryCoordinatorWithIndices.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber is a free data retrieval call binding the contract method 0x85020d49.
//
// Solidity: function getQuorumBitmapIndicesByOperatorIdsAtBlockNumber(uint32 blockNumber, bytes32[] operatorIds) view returns(uint32[])
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber(blockNumber uint32, operatorIds [][32]byte) ([]uint32, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetQuorumBitmapIndicesByOperatorIdsAtBlockNumber(&_BLSRegistryCoordinatorWithIndices.CallOpts, blockNumber, operatorIds)
}

// GetQuorumBitmapUpdateByOperatorIdByIndex is a free data retrieval call binding the contract method 0x0159f1ce.
//
// Solidity: function getQuorumBitmapUpdateByOperatorIdByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) GetQuorumBitmapUpdateByOperatorIdByIndex(opts *bind.CallOpts, operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "getQuorumBitmapUpdateByOperatorIdByIndex", operatorId, index)

	if err != nil {
		return *new(IRegistryCoordinatorQuorumBitmapUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryCoordinatorQuorumBitmapUpdate)).(*IRegistryCoordinatorQuorumBitmapUpdate)

	return out0, err

}

// GetQuorumBitmapUpdateByOperatorIdByIndex is a free data retrieval call binding the contract method 0x0159f1ce.
//
// Solidity: function getQuorumBitmapUpdateByOperatorIdByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) GetQuorumBitmapUpdateByOperatorIdByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetQuorumBitmapUpdateByOperatorIdByIndex(&_BLSRegistryCoordinatorWithIndices.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByOperatorIdByIndex is a free data retrieval call binding the contract method 0x0159f1ce.
//
// Solidity: function getQuorumBitmapUpdateByOperatorIdByIndex(bytes32 operatorId, uint256 index) view returns((uint32,uint32,uint192))
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) GetQuorumBitmapUpdateByOperatorIdByIndex(operatorId [32]byte, index *big.Int) (IRegistryCoordinatorQuorumBitmapUpdate, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetQuorumBitmapUpdateByOperatorIdByIndex(&_BLSRegistryCoordinatorWithIndices.CallOpts, operatorId, index)
}

// GetQuorumBitmapUpdateByOperatorIdLength is a free data retrieval call binding the contract method 0x055a62b6.
//
// Solidity: function getQuorumBitmapUpdateByOperatorIdLength(bytes32 operatorId) view returns(uint256)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) GetQuorumBitmapUpdateByOperatorIdLength(opts *bind.CallOpts, operatorId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "getQuorumBitmapUpdateByOperatorIdLength", operatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumBitmapUpdateByOperatorIdLength is a free data retrieval call binding the contract method 0x055a62b6.
//
// Solidity: function getQuorumBitmapUpdateByOperatorIdLength(bytes32 operatorId) view returns(uint256)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) GetQuorumBitmapUpdateByOperatorIdLength(operatorId [32]byte) (*big.Int, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetQuorumBitmapUpdateByOperatorIdLength(&_BLSRegistryCoordinatorWithIndices.CallOpts, operatorId)
}

// GetQuorumBitmapUpdateByOperatorIdLength is a free data retrieval call binding the contract method 0x055a62b6.
//
// Solidity: function getQuorumBitmapUpdateByOperatorIdLength(bytes32 operatorId) view returns(uint256)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) GetQuorumBitmapUpdateByOperatorIdLength(operatorId [32]byte) (*big.Int, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.GetQuorumBitmapUpdateByOperatorIdLength(&_BLSRegistryCoordinatorWithIndices.CallOpts, operatorId)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) IndexRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "indexRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) IndexRegistry() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.IndexRegistry(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// IndexRegistry is a free data retrieval call binding the contract method 0x9e9923c2.
//
// Solidity: function indexRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) IndexRegistry() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.IndexRegistry(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) NumRegistries(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "numRegistries")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) NumRegistries() (*big.Int, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.NumRegistries(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// NumRegistries is a free data retrieval call binding the contract method 0xd72d8dd6.
//
// Solidity: function numRegistries() view returns(uint256)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) NumRegistries() (*big.Int, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.NumRegistries(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) Registries(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "registries", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.Registries(&_BLSRegistryCoordinatorWithIndices.CallOpts, arg0)
}

// Registries is a free data retrieval call binding the contract method 0x6347c900.
//
// Solidity: function registries(uint256 ) view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) Registries(arg0 *big.Int) (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.Registries(&_BLSRegistryCoordinatorWithIndices.CallOpts, arg0)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) ServiceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "serviceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) ServiceManager() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ServiceManager(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// ServiceManager is a free data retrieval call binding the contract method 0x3998fdd3.
//
// Solidity: function serviceManager() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) ServiceManager() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.ServiceManager(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) Slasher() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.Slasher(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) Slasher() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.Slasher(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) StakeRegistry() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.StakeRegistry(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) StakeRegistry() (common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.StakeRegistry(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) TargetArtifactSelectors() ([]StdInvariantFuzzSelector, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetArtifactSelectors(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzSelector, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetArtifactSelectors(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) TargetArtifacts() ([]string, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetArtifacts(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) TargetArtifacts() ([]string, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetArtifacts(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) TargetContracts() ([]common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetContracts(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) TargetContracts() ([]common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetContracts(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetSelectors(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetSelectors(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BLSRegistryCoordinatorWithIndices.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) TargetSenders() ([]common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetSenders(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesCallerSession) TargetSenders() ([]common.Address, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.TargetSenders(&_BLSRegistryCoordinatorWithIndices.CallOpts)
}

// DeregisterOperatorWithCoordinator is a paid mutator transaction binding the contract method 0xc81b1ff4.
//
// Solidity: function deregisterOperatorWithCoordinator(bytes quorumNumbers, bytes deregistrationData) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) DeregisterOperatorWithCoordinator(opts *bind.TransactOpts, quorumNumbers []byte, deregistrationData []byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "deregisterOperatorWithCoordinator", quorumNumbers, deregistrationData)
}

// DeregisterOperatorWithCoordinator is a paid mutator transaction binding the contract method 0xc81b1ff4.
//
// Solidity: function deregisterOperatorWithCoordinator(bytes quorumNumbers, bytes deregistrationData) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) DeregisterOperatorWithCoordinator(quorumNumbers []byte, deregistrationData []byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.DeregisterOperatorWithCoordinator(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, deregistrationData)
}

// DeregisterOperatorWithCoordinator is a paid mutator transaction binding the contract method 0xc81b1ff4.
//
// Solidity: function deregisterOperatorWithCoordinator(bytes quorumNumbers, bytes deregistrationData) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) DeregisterOperatorWithCoordinator(quorumNumbers []byte, deregistrationData []byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.DeregisterOperatorWithCoordinator(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, deregistrationData)
}

// DeregisterOperatorWithCoordinator0 is a paid mutator transaction binding the contract method 0xf807d627.
//
// Solidity: function deregisterOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, bytes32[] operatorIdsToSwap) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) DeregisterOperatorWithCoordinator0(opts *bind.TransactOpts, quorumNumbers []byte, pubkey BN254G1Point, operatorIdsToSwap [][32]byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "deregisterOperatorWithCoordinator0", quorumNumbers, pubkey, operatorIdsToSwap)
}

// DeregisterOperatorWithCoordinator0 is a paid mutator transaction binding the contract method 0xf807d627.
//
// Solidity: function deregisterOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, bytes32[] operatorIdsToSwap) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) DeregisterOperatorWithCoordinator0(quorumNumbers []byte, pubkey BN254G1Point, operatorIdsToSwap [][32]byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.DeregisterOperatorWithCoordinator0(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, pubkey, operatorIdsToSwap)
}

// DeregisterOperatorWithCoordinator0 is a paid mutator transaction binding the contract method 0xf807d627.
//
// Solidity: function deregisterOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, bytes32[] operatorIdsToSwap) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) DeregisterOperatorWithCoordinator0(quorumNumbers []byte, pubkey BN254G1Point, operatorIdsToSwap [][32]byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.DeregisterOperatorWithCoordinator0(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, pubkey, operatorIdsToSwap)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) Failed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "failed")
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) Failed() (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.Failed(&_BLSRegistryCoordinatorWithIndices.TransactOpts)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) Failed() (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.Failed(&_BLSRegistryCoordinatorWithIndices.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x74056186.
//
// Solidity: function initialize((uint32,uint16,uint16)[] _operatorSetParams) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) Initialize(opts *bind.TransactOpts, _operatorSetParams []IBLSRegistryCoordinatorWithIndicesOperatorSetParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "initialize", _operatorSetParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x74056186.
//
// Solidity: function initialize((uint32,uint16,uint16)[] _operatorSetParams) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) Initialize(_operatorSetParams []IBLSRegistryCoordinatorWithIndicesOperatorSetParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.Initialize(&_BLSRegistryCoordinatorWithIndices.TransactOpts, _operatorSetParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x74056186.
//
// Solidity: function initialize((uint32,uint16,uint16)[] _operatorSetParams) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) Initialize(_operatorSetParams []IBLSRegistryCoordinatorWithIndicesOperatorSetParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.Initialize(&_BLSRegistryCoordinatorWithIndices.TransactOpts, _operatorSetParams)
}

// RegisterOperatorWithCoordinator is a paid mutator transaction binding the contract method 0x526ea94e.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, bytes registrationData) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) RegisterOperatorWithCoordinator(opts *bind.TransactOpts, quorumNumbers []byte, registrationData []byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "registerOperatorWithCoordinator", quorumNumbers, registrationData)
}

// RegisterOperatorWithCoordinator is a paid mutator transaction binding the contract method 0x526ea94e.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, bytes registrationData) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) RegisterOperatorWithCoordinator(quorumNumbers []byte, registrationData []byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.RegisterOperatorWithCoordinator(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, registrationData)
}

// RegisterOperatorWithCoordinator is a paid mutator transaction binding the contract method 0x526ea94e.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, bytes registrationData) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) RegisterOperatorWithCoordinator(quorumNumbers []byte, registrationData []byte) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.RegisterOperatorWithCoordinator(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, registrationData)
}

// RegisterOperatorWithCoordinator0 is a paid mutator transaction binding the contract method 0x8f84880e.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, string socket, (address,(uint256,uint256),bytes32[])[] operatorKickParams) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) RegisterOperatorWithCoordinator0(opts *bind.TransactOpts, quorumNumbers []byte, pubkey BN254G1Point, socket string, operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "registerOperatorWithCoordinator0", quorumNumbers, pubkey, socket, operatorKickParams)
}

// RegisterOperatorWithCoordinator0 is a paid mutator transaction binding the contract method 0x8f84880e.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, string socket, (address,(uint256,uint256),bytes32[])[] operatorKickParams) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) RegisterOperatorWithCoordinator0(quorumNumbers []byte, pubkey BN254G1Point, socket string, operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.RegisterOperatorWithCoordinator0(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, pubkey, socket, operatorKickParams)
}

// RegisterOperatorWithCoordinator0 is a paid mutator transaction binding the contract method 0x8f84880e.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, string socket, (address,(uint256,uint256),bytes32[])[] operatorKickParams) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) RegisterOperatorWithCoordinator0(quorumNumbers []byte, pubkey BN254G1Point, socket string, operatorKickParams []IBLSRegistryCoordinatorWithIndicesOperatorKickParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.RegisterOperatorWithCoordinator0(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, pubkey, socket, operatorKickParams)
}

// RegisterOperatorWithCoordinator1 is a paid mutator transaction binding the contract method 0xc66ab9ca.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, string socket) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) RegisterOperatorWithCoordinator1(opts *bind.TransactOpts, quorumNumbers []byte, pubkey BN254G1Point, socket string) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "registerOperatorWithCoordinator1", quorumNumbers, pubkey, socket)
}

// RegisterOperatorWithCoordinator1 is a paid mutator transaction binding the contract method 0xc66ab9ca.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, string socket) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) RegisterOperatorWithCoordinator1(quorumNumbers []byte, pubkey BN254G1Point, socket string) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.RegisterOperatorWithCoordinator1(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, pubkey, socket)
}

// RegisterOperatorWithCoordinator1 is a paid mutator transaction binding the contract method 0xc66ab9ca.
//
// Solidity: function registerOperatorWithCoordinator(bytes quorumNumbers, (uint256,uint256) pubkey, string socket) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) RegisterOperatorWithCoordinator1(quorumNumbers []byte, pubkey BN254G1Point, socket string) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.RegisterOperatorWithCoordinator1(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumbers, pubkey, socket)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParam) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) SetOperatorSetParams(opts *bind.TransactOpts, quorumNumber uint8, operatorSetParam IBLSRegistryCoordinatorWithIndicesOperatorSetParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "setOperatorSetParams", quorumNumber, operatorSetParam)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParam) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParam IBLSRegistryCoordinatorWithIndicesOperatorSetParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.SetOperatorSetParams(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumber, operatorSetParam)
}

// SetOperatorSetParams is a paid mutator transaction binding the contract method 0x5b0b829f.
//
// Solidity: function setOperatorSetParams(uint8 quorumNumber, (uint32,uint16,uint16) operatorSetParam) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) SetOperatorSetParams(quorumNumber uint8, operatorSetParam IBLSRegistryCoordinatorWithIndicesOperatorSetParam) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.SetOperatorSetParams(&_BLSRegistryCoordinatorWithIndices.TransactOpts, quorumNumber, operatorSetParam)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactor) UpdateSocket(opts *bind.TransactOpts, socket string) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.contract.Transact(opts, "updateSocket", socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.UpdateSocket(&_BLSRegistryCoordinatorWithIndices.TransactOpts, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x0cf4b767.
//
// Solidity: function updateSocket(string socket) returns()
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesTransactorSession) UpdateSocket(socket string) (*types.Transaction, error) {
	return _BLSRegistryCoordinatorWithIndices.Contract.UpdateSocket(&_BLSRegistryCoordinatorWithIndices.TransactOpts, socket)
}

// BLSRegistryCoordinatorWithIndicesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesInitializedIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesInitialized represents a Initialized event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterInitialized(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesInitializedIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesInitializedIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesInitialized) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesInitialized)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseInitialized(log types.Log) (*BLSRegistryCoordinatorWithIndicesInitialized, error) {
	event := new(BLSRegistryCoordinatorWithIndicesInitialized)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesOperatorDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesOperatorDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesOperatorDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesOperatorDeregistered represents a OperatorDeregistered event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesOperatorDeregistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*BLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesOperatorDeregisteredIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesOperatorDeregistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesOperatorDeregistered)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorDeregistered is a log parse operation binding the contract event 0x396fdcb180cb0fea26928113fb0fd1c3549863f9cd563e6a184f1d578116c8e4.
//
// Solidity: event OperatorDeregistered(address indexed operator, bytes32 indexed operatorId)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseOperatorDeregistered(log types.Log) (*BLSRegistryCoordinatorWithIndicesOperatorDeregistered, error) {
	event := new(BLSRegistryCoordinatorWithIndicesOperatorDeregistered)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesOperatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesOperatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesOperatorRegistered represents a OperatorRegistered event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesOperatorRegistered struct {
	Operator   common.Address
	OperatorId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address, operatorId [][32]byte) (*BLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesOperatorRegisteredIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesOperatorRegistered, operator []common.Address, operatorId [][32]byte) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "OperatorRegistered", operatorRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesOperatorRegistered)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistered is a log parse operation binding the contract event 0xe8e68cef1c3a761ed7be7e8463a375f27f7bc335e51824223cacce636ec5c3fe.
//
// Solidity: event OperatorRegistered(address indexed operator, bytes32 indexed operatorId)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseOperatorRegistered(log types.Log) (*BLSRegistryCoordinatorWithIndicesOperatorRegistered, error) {
	event := new(BLSRegistryCoordinatorWithIndicesOperatorRegistered)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator is returned from FilterOperatorSetParamsUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetParamsUpdated events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated represents a OperatorSetParamsUpdated event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated struct {
	QuorumNumber      uint8
	OperatorSetParams IBLSRegistryCoordinatorWithIndicesOperatorSetParam
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetParamsUpdated is a free log retrieval operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterOperatorSetParamsUpdated(opts *bind.FilterOpts, quorumNumber []uint8) (*BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdatedIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "OperatorSetParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetParamsUpdated is a free log subscription operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchOperatorSetParamsUpdated(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated, quorumNumber []uint8) (event.Subscription, error) {

	var quorumNumberRule []interface{}
	for _, quorumNumberItem := range quorumNumber {
		quorumNumberRule = append(quorumNumberRule, quorumNumberItem)
	}

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "OperatorSetParamsUpdated", quorumNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSetParamsUpdated is a log parse operation binding the contract event 0x3ee6fe8d54610244c3e9d3c066ae4aee997884aa28f10616ae821925401318ac.
//
// Solidity: event OperatorSetParamsUpdated(uint8 indexed quorumNumber, (uint32,uint16,uint16) operatorSetParams)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseOperatorSetParamsUpdated(log types.Log) (*BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated, error) {
	event := new(BLSRegistryCoordinatorWithIndicesOperatorSetParamsUpdated)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "OperatorSetParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator is returned from FilterOperatorSocketUpdate and is used to iterate over the raw logs and unpacked data for OperatorSocketUpdate events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate represents a OperatorSocketUpdate event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate struct {
	Operator common.Address
	Socket   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketUpdate is a free log retrieval operation binding the contract event 0x36538cd2b13583680f88eb5649e2e9c6a4540738f146954b468c403189107747.
//
// Solidity: event OperatorSocketUpdate(address operator, string socket)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterOperatorSocketUpdate(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "OperatorSocketUpdate")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesOperatorSocketUpdateIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "OperatorSocketUpdate", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketUpdate is a free log subscription operation binding the contract event 0x36538cd2b13583680f88eb5649e2e9c6a4540738f146954b468c403189107747.
//
// Solidity: event OperatorSocketUpdate(address operator, string socket)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchOperatorSocketUpdate(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "OperatorSocketUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSocketUpdate is a log parse operation binding the contract event 0x36538cd2b13583680f88eb5649e2e9c6a4540738f146954b468c403189107747.
//
// Solidity: event OperatorSocketUpdate(address operator, string socket)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseOperatorSocketUpdate(log types.Log) (*BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate, error) {
	event := new(BLSRegistryCoordinatorWithIndicesOperatorSocketUpdate)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "OperatorSocketUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLog represents a Log event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLog(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLog) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLog)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLog(log types.Log) (*BLSRegistryCoordinatorWithIndicesLog, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLog)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogAddressIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogAddress represents a LogAddress event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogAddress(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogAddressIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogAddressIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogAddress) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogAddress)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogAddress(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogAddress, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogAddress)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogArrayIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogArray represents a LogArray event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogArray(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogArrayIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogArrayIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogArray) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogArray)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogArray(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogArray, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogArray)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogArray0Iterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogArray0 represents a LogArray0 event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogArray0(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogArray0Iterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogArray0Iterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogArray0) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogArray0)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogArray0(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogArray0, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogArray0)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogArray1Iterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogArray1 represents a LogArray1 event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogArray1(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogArray1Iterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogArray1Iterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogArray1) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogArray1)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogArray1(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogArray1, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogArray1)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogBytesIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogBytes represents a LogBytes event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogBytes(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogBytesIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogBytesIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogBytes) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogBytes)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogBytes(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogBytes, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogBytes)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogBytes32Iterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogBytes32 represents a LogBytes32 event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogBytes32Iterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogBytes32Iterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogBytes32) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogBytes32)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogBytes32(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogBytes32, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogBytes32)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogIntIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogInt represents a LogInt event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogInt(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogIntIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogIntIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogInt) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogInt)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogInt(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogInt, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogInt)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedAddressIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedAddress represents a LogNamedAddress event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedAddressIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedAddressIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedAddress)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_address", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedAddress(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedAddress, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedAddress)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedArrayIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedArray // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedArray)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedArray)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedArray represents a LogNamedArray event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedArrayIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedArrayIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedArray)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_array", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedArray(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedArray, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedArray)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedArray0Iterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedArray0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedArray0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedArray0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedArray0 represents a LogNamedArray0 event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedArray0Iterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedArray0Iterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedArray0)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_array0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedArray0(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedArray0, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedArray0)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedArray1Iterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedArray1 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedArray1)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedArray1)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedArray1 represents a LogNamedArray1 event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedArray1Iterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedArray1Iterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedArray1)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_array1", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedArray1(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedArray1, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedArray1)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedBytesIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedBytes represents a LogNamedBytes event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedBytesIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedBytesIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedBytes)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedBytes(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedBytes, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedBytes)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedBytes32Iterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedBytes32 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedBytes32)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedBytes32)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedBytes32 represents a LogNamedBytes32 event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedBytes32Iterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedBytes32Iterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedBytes32)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedBytes32(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedBytes32, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedBytes32)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedDecimalIntIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedDecimalIntIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedDecimalIntIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedDecimalInt(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedDecimalInt)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedDecimalUintIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedDecimalUintIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedDecimalUintIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedDecimalUint(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedDecimalUint)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedIntIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedInt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedInt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedInt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedInt represents a LogNamedInt event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedIntIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedIntIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedInt)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_int", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedInt(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedInt, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedInt)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedStringIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedString represents a LogNamedString event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedStringIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedStringIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedString) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedString)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedString(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedString, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedString)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedUintIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogNamedUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogNamedUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogNamedUint represents a LogNamedUint event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogNamedUintIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogNamedUintIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogNamedUint)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogNamedUint(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogNamedUint, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogNamedUint)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogStringIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogString represents a LogString event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogString(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogStringIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogStringIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogString) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogString)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_string", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogString(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogString, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogString)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogUintIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogUint represents a LogUint event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogUint(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogUintIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogUintIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogUint) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogUint)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogUint(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogUint, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogUint)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BLSRegistryCoordinatorWithIndicesLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogsIterator struct {
	Event *BLSRegistryCoordinatorWithIndicesLogs // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BLSRegistryCoordinatorWithIndicesLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BLSRegistryCoordinatorWithIndicesLogs)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BLSRegistryCoordinatorWithIndicesLogs)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BLSRegistryCoordinatorWithIndicesLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BLSRegistryCoordinatorWithIndicesLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BLSRegistryCoordinatorWithIndicesLogs represents a Logs event raised by the BLSRegistryCoordinatorWithIndices contract.
type BLSRegistryCoordinatorWithIndicesLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) FilterLogs(opts *bind.FilterOpts) (*BLSRegistryCoordinatorWithIndicesLogsIterator, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &BLSRegistryCoordinatorWithIndicesLogsIterator{contract: _BLSRegistryCoordinatorWithIndices.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *BLSRegistryCoordinatorWithIndicesLogs) (event.Subscription, error) {

	logs, sub, err := _BLSRegistryCoordinatorWithIndices.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BLSRegistryCoordinatorWithIndicesLogs)
				if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "logs", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_BLSRegistryCoordinatorWithIndices *BLSRegistryCoordinatorWithIndicesFilterer) ParseLogs(log types.Log) (*BLSRegistryCoordinatorWithIndicesLogs, error) {
	event := new(BLSRegistryCoordinatorWithIndicesLogs)
	if err := _BLSRegistryCoordinatorWithIndices.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
