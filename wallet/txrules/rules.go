// Copyright (c) 2016 The btcsuite developers
// Copyright (c) 2016 The Legenddigital developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txrules

import (
	"github.com/Legenddigital/lddld/lddlutil"
	"github.com/Legenddigital/lddld/txscript"
	"github.com/Legenddigital/lddld/wire"
	"github.com/Legenddigital/lddlwallet/errors"
	h "github.com/Legenddigital/lddlwallet/internal/helpers"
)

// DefaultRelayFeePerKb is the default minimum relay fee policy for a mempool.
const DefaultRelayFeePerKb lddlutil.Amount = 1e5

// IsDustAmount determines whether a transaction output value and script length would
// cause the output to be considered dust.  Transactions with dust outputs are
// not standard and are rejected by mempools with default policies.
func IsDustAmount(amount lddlutil.Amount, scriptSize int, relayFeePerKb lddlutil.Amount) bool {
	// Calculate the total (estimated) cost to the network.  This is
	// calculated using the serialize size of the output plus the serial
	// size of a transaction input which redeems it.  The output is assumed
	// to be compressed P2PKH as this is the most common script type.  Use
	// the average size of a compressed P2PKH redeem input (165) rather than
	// the largest possible (txsizes.RedeemP2PKHInputSize).
	totalSize := 8 + 2 + wire.VarIntSerializeSize(uint64(scriptSize)) +
		scriptSize + 165

	// Dust is defined as an output value where the total cost to the network
	// (output size + input size) is greater than 1/3 of the relay fee.
	return int64(amount)*1000/(3*int64(totalSize)) < int64(relayFeePerKb)
}

// IsDustOutput determines whether a transaction output is considered dust.
// Transactions with dust outputs are not standard and are rejected by mempools
// with default policies.
func IsDustOutput(output *wire.TxOut, relayFeePerKb lddlutil.Amount) bool {
	// Unspendable outputs which solely carry data are not checked for dust.
	if txscript.GetScriptClass(output.Version, output.PkScript) == txscript.NullDataTy {
		return false
	}

	// All other unspendable outputs are considered dust.
	if txscript.IsUnspendable(output.Value, output.PkScript) {
		return true
	}

	return IsDustAmount(lddlutil.Amount(output.Value), len(output.PkScript),
		relayFeePerKb)
}

// CheckOutput performs simple consensus and policy tests on a transaction
// output.  Returns with errors.Invalid if output violates consensus rules, and
// errors.Policy if the output violates a non-consensus policy.
func CheckOutput(output *wire.TxOut, relayFeePerKb lddlutil.Amount) error {
	if output.Value < 0 {
		return errors.E(errors.Invalid, "transaction output amount is negative")
	}
	if output.Value > lddlutil.MaxAmount {
		return errors.E(errors.Invalid, "transaction output amount exceeds maximum value")
	}
	if IsDustOutput(output, relayFeePerKb) {
		return errors.E(errors.Policy, "transaction output is dust")
	}
	return nil
}

// FeeForSerializeSize calculates the required fee for a transaction of some
// arbitrary size given a mempool's relay fee policy.
func FeeForSerializeSize(relayFeePerKb lddlutil.Amount, txSerializeSize int) lddlutil.Amount {
	fee := relayFeePerKb * lddlutil.Amount(txSerializeSize) / 1000

	if fee == 0 && relayFeePerKb > 0 {
		fee = relayFeePerKb
	}

	if fee < 0 || fee > lddlutil.MaxAmount {
		fee = lddlutil.MaxAmount
	}

	return fee
}

// PaysHighFees checks whether the signed transaction pays insanely high fees.
// Transactons are defined to have a high fee if they have pay a fee rate that
// is 1000 time higher than the default fee.
func PaysHighFees(totalInput lddlutil.Amount, tx *wire.MsgTx) bool {
	fee := totalInput - h.SumOutputValues(tx.TxOut)
	if fee <= 0 {
		// Impossible to determine
		return false
	}

	maxFee := FeeForSerializeSize(1000*DefaultRelayFeePerKb, tx.SerializeSize())
	return fee > maxFee
}
