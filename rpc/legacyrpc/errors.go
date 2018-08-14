// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2016-2018 The Legenddigital developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package legacyrpc

import (
	"fmt"

	"github.com/Legenddigital/lddld/lddljson"
	"github.com/Legenddigital/lddlwallet/errors"
)

func convertError(err error) *lddljson.RPCError {
	if err, ok := err.(*lddljson.RPCError); ok {
		return err
	}

	code := lddljson.ErrRPCWallet
	if err, ok := err.(*errors.Error); ok {
		switch err.Kind {
		case errors.Bug:
			code = lddljson.ErrRPCInternal.Code
		case errors.Encoding:
			code = lddljson.ErrRPCInvalidParameter
		case errors.Locked:
			code = lddljson.ErrRPCWalletUnlockNeeded
		case errors.Passphrase:
			code = lddljson.ErrRPCWalletPassphraseIncorrect
		case errors.NoPeers:
			code = lddljson.ErrRPCClientNotConnected
		case errors.InsufficientBalance:
			code = lddljson.ErrRPCWalletInsufficientFunds
		}
	}
	return &lddljson.RPCError{
		Code:    code,
		Message: err.Error(),
	}
}

func rpcError(code lddljson.RPCErrorCode, err error) *lddljson.RPCError {
	return &lddljson.RPCError{
		Code:    code,
		Message: err.Error(),
	}
}

func rpcErrorf(code lddljson.RPCErrorCode, format string, args ...interface{}) *lddljson.RPCError {
	return &lddljson.RPCError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// Errors variables that are defined once here to avoid duplication.
var (
	errUnloadedWallet = &lddljson.RPCError{
		Code:    lddljson.ErrRPCWallet,
		Message: "request requires a wallet but wallet has not loaded yet",
	}

	errRPCClientNotConnected = &lddljson.RPCError{
		Code:    lddljson.ErrRPCClientNotConnected,
		Message: "disconnected from consensus RPC",
	}

	errNoNetwork = &lddljson.RPCError{
		Code:    lddljson.ErrRPCClientNotConnected,
		Message: "disconnected from network",
	}

	errAccountNotFound = &lddljson.RPCError{
		Code:    lddljson.ErrRPCWalletInvalidAccountName,
		Message: "account not found",
	}

	errAddressNotInWallet = &lddljson.RPCError{
		Code:    lddljson.ErrRPCWallet,
		Message: "address not found in wallet",
	}

	errNotImportedAccount = &lddljson.RPCError{
		Code:    lddljson.ErrRPCWallet,
		Message: "imported addresses must belong to the imported account",
	}

	errNeedPositiveAmount = &lddljson.RPCError{
		Code:    lddljson.ErrRPCInvalidParameter,
		Message: "amount must be positive",
	}

	errWalletUnlockNeeded = &lddljson.RPCError{
		Code:    lddljson.ErrRPCWalletUnlockNeeded,
		Message: "enter the wallet passphrase with walletpassphrase first",
	}

	errReservedAccountName = &lddljson.RPCError{
		Code:    lddljson.ErrRPCInvalidParameter,
		Message: "account name is reserved by RPC server",
	}
)
