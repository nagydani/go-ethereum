// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package bind

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

<<<<<<< HEAD
// ErrNoCode is returned by call and transact operations for which the requested
// recipient contract to operate on does not exist in the state db or does not
// have any code associated with it (i.e. suicided).
var ErrNoCode = errors.New("no contract code at given address")
=======
var (
	// ErrNoCode is returned by call and transact operations for which the requested
	// recipient contract to operate on does not exist in the state db or does not
	// have any code associated with it (i.e. suicided).
	ErrNoCode = errors.New("no contract code at given address")

	// This error is raised when attempting to perform a pending state action
	// on a backend that doesn't implement PendingContractCaller.
	ErrNoPendingState = errors.New("backend does not support pending state")

	// This error is returned by WaitDeployed if contract creation leaves an
	// empty contract behind.
	ErrNoCodeAfterDeploy = errors.New("no contract code after deployment")
)
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23

// ContractCaller defines the methods needed to allow operating with contract on a read
// only basis.
type ContractCaller interface {
<<<<<<< HEAD
	// HasCode checks if the contract at the given address has any code associated
	// with it or not. This is needed to differentiate between contract internal
	// errors and the local chain being out of sync.
	HasCode(contract common.Address, pending bool) (bool, error)

	// ContractCall executes an Ethereum contract call with the specified data as
	// the input. The pending flag requests execution against the pending block, not
	// the stable head of the chain.
	ContractCall(contract common.Address, data []byte, pending bool) ([]byte, error)
=======
	// CodeAt returns the code of the given account. This is needed to differentiate
	// between contract internal errors and the local chain being out of sync.
	CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error)
	// ContractCall executes an Ethereum contract call with the specified data as the
	// input.
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}

// DeployBackend wraps the operations needed by WaitMined and WaitDeployed.
type DeployBackend interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
}

// PendingContractCaller defines methods to perform contract calls on the pending state.
// Call will try to discover this interface when access to the pending state is requested.
// If the backend does not support the pending state, Call returns ErrNoPendingState.
type PendingContractCaller interface {
	// PendingCodeAt returns the code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error)
	// PendingCallContract executes an Ethereum contract call against the pending state.
	PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error)
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
}

// ContractTransactor defines the methods needed to allow operating with contract
// on a write only basis. Beside the transacting method, the remainder are helpers
// used when the user does not provide some needed values, but rather leaves it up
// to the transactor to decide.
type ContractTransactor interface {
	// PendingCodeAt returns the code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	// PendingNonceAt retrieves the current pending nonce associated with an account.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
<<<<<<< HEAD
	SuggestGasPrice() (*big.Int, error)

	// HasCode checks if the contract at the given address has any code associated
	// with it or not. This is needed to differentiate between contract internal
	// errors and the local chain being out of sync.
	HasCode(contract common.Address, pending bool) (bool, error)

	// EstimateGasLimit tries to estimate the gas needed to execute a specific
=======
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// EstimateGas tries to estimate the gas needed to execute a specific
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
	// transaction based on the current pending state of the backend blockchain.
	// There is no guarantee that this is the true gas limit requirement as other
	// transactions may be added or removed by miners, but it should provide a basis
	// for setting a reasonable default.
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (usedGas *big.Int, err error)
	// SendTransaction injects the transaction into the pending pool for execution.
	SendTransaction(ctx context.Context, tx *types.Transaction) error
}

<<<<<<< HEAD
// ContractBackend defines the methods needed to allow operating with contract
// on a read-write basis.
//
// This interface is essentially the union of ContractCaller and ContractTransactor
// but due to a bug in the Go compiler (https://github.com/golang/go/issues/6977),
// we cannot simply list it as the two interfaces. The other solution is to add a
// third interface containing the common methods, but that convolutes the user API
// as it introduces yet another parameter to require for initialization.
=======
// ContractBackend defines the methods needed to work with contracts on a read-write basis.
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
type ContractBackend interface {
	// HasCode checks if the contract at the given address has any code associated
	// with it or not. This is needed to differentiate between contract internal
	// errors and the local chain being out of sync.
	HasCode(contract common.Address, pending bool) (bool, error)

	// ContractCall executes an Ethereum contract call with the specified data as
	// the input. The pending flag requests execution against the pending block, not
	// the stable head of the chain.
	ContractCall(contract common.Address, data []byte, pending bool) ([]byte, error)

	// PendingAccountNonce retrieves the current pending nonce associated with an
	// account.
	PendingAccountNonce(account common.Address) (uint64, error)

	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice() (*big.Int, error)

	// EstimateGasLimit tries to estimate the gas needed to execute a specific
	// transaction based on the current pending state of the backend blockchain.
	// There is no guarantee that this is the true gas limit requirement as other
	// transactions may be added or removed by miners, but it should provide a basis
	// for setting a reasonable default.
	EstimateGasLimit(sender common.Address, contract *common.Address, value *big.Int, data []byte) (*big.Int, error)

	// SendTransaction injects the transaction into the pending pool for execution.
	SendTransaction(tx *types.Transaction) error
}
