// Copyright (c) 2015 The btcsuite developers
// Copyright (c) 2015-2017 The Legenddigital developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

//+build !generate

package rpchelp

import "github.com/Legenddigital/lddld/lddljson"

// Common return types.
var (
	returnsBool        = []interface{}{(*bool)(nil)}
	returnsNumber      = []interface{}{(*float64)(nil)}
	returnsString      = []interface{}{(*string)(nil)}
	returnsStringArray = []interface{}{(*[]string)(nil)}
	returnsLTRArray    = []interface{}{(*[]lddljson.ListTransactionsResult)(nil)}
)

// Methods contains all methods and result types that help is generated for,
// for every locale.
var Methods = []struct {
	Method      string
	ResultTypes []interface{}
}{
	{"accountaddressindex", []interface{}{(*int)(nil)}},
	{"accountsyncaddressindex", nil},
	{"addmultisigaddress", returnsString},
	{"consolidate", returnsString},
	{"createmultisig", []interface{}{(*lddljson.CreateMultiSigResult)(nil)}},
	{"dumpprivkey", returnsString},
	{"getaccount", returnsString},
	{"getaccountaddress", returnsString},
	{"getaddressesbyaccount", returnsStringArray},
	{"getbalance", []interface{}{(*lddljson.GetBalanceResult)(nil)}},
	{"getbestblockhash", returnsString},
	{"getblockcount", returnsNumber},
	{"getinfo", []interface{}{(*lddljson.InfoWalletResult)(nil)}},
	{"getmasterpubkey", []interface{}{(*string)(nil)}},
	{"getmultisigoutinfo", []interface{}{(*lddljson.GetMultisigOutInfoResult)(nil)}},
	{"getnewaddress", returnsString},
	{"getrawchangeaddress", returnsString},
	{"getreceivedbyaccount", returnsNumber},
	{"getreceivedbyaddress", returnsNumber},
	{"gettickets", []interface{}{(*lddljson.GetTicketsResult)(nil)}},
	{"gettransaction", []interface{}{(*lddljson.GetTransactionResult)(nil)}},
	{"getvotechoices", []interface{}{(*lddljson.GetVoteChoicesResult)(nil)}},
	{"help", append(returnsString, returnsString[0])},
	{"importprivkey", nil},
	{"importscript", nil},
	{"keypoolrefill", nil},
	{"listaccounts", []interface{}{(*map[string]float64)(nil)}},
	{"listlockunspent", []interface{}{(*[]lddljson.TransactionInput)(nil)}},
	{"listreceivedbyaccount", []interface{}{(*[]lddljson.ListReceivedByAccountResult)(nil)}},
	{"listreceivedbyaddress", []interface{}{(*[]lddljson.ListReceivedByAddressResult)(nil)}},
	{"listsinceblock", []interface{}{(*lddljson.ListSinceBlockResult)(nil)}},
	{"listtransactions", returnsLTRArray},
	{"listunspent", []interface{}{(*lddljson.ListUnspentResult)(nil)}},
	{"lockunspent", returnsBool},
	{"redeemmultisigout", []interface{}{(*lddljson.RedeemMultiSigOutResult)(nil)}},
	{"redeemmultisigouts", []interface{}{(*lddljson.RedeemMultiSigOutResult)(nil)}},
	{"rescanwallet", nil},
	{"revoketickets", nil},
	{"sendfrom", returnsString},
	{"sendmany", returnsString},
	{"sendtoaddress", returnsString},
	{"sendtomultisig", returnsString},
	{"settxfee", returnsBool},
	{"setvotechoice", nil},
	{"signmessage", returnsString},
	{"signrawtransaction", []interface{}{(*lddljson.SignRawTransactionResult)(nil)}},
	{"signrawtransactions", []interface{}{(*lddljson.SignRawTransactionsResult)(nil)}},
	{"startautobuyer", nil},
	{"stopautobuyer", nil},
	{"sweepaccount", []interface{}{(*lddljson.SweepAccountResult)(nil)}},
	{"validateaddress", []interface{}{(*lddljson.ValidateAddressWalletResult)(nil)}},
	{"verifymessage", returnsBool},
	{"version", []interface{}{(*map[string]lddljson.VersionResult)(nil)}},
	{"walletlock", nil},
	{"walletpassphrase", nil},
	{"walletpassphrasechange", nil},
	{"createnewaccount", nil},
	{"exportwatchingwallet", returnsString},
	{"getbestblock", []interface{}{(*lddljson.GetBestBlockResult)(nil)}},
	{"getunconfirmedbalance", returnsNumber},
	{"listaddresstransactions", returnsLTRArray},
	{"listalltransactions", returnsLTRArray},
	{"renameaccount", nil},
	{"walletislocked", returnsBool},
	{"walletinfo", []interface{}{(*lddljson.WalletInfoResult)(nil)}},

	// TODO Alphabetize
	{"purchaseticket", returnsString},
	{"generatevote", []interface{}{(*lddljson.GenerateVoteResult)(nil)}},
	{"getstakeinfo", []interface{}{(*lddljson.GetStakeInfoResult)(nil)}},
	{"getticketfee", returnsNumber},
	{"setticketfee", returnsBool},
	{"getwalletfee", returnsNumber},
	{"addticket", nil},
	{"listscripts", []interface{}{(*lddljson.ListScriptsResult)(nil)}},
	{"stakepooluserinfo", []interface{}{(*lddljson.StakePoolUserInfoResult)(nil)}},
	{"ticketsforaddress", returnsBool},
}

// HelpDescs contains the locale-specific help strings along with the locale.
var HelpDescs = []struct {
	Locale   string // Actual locale, e.g. en_US
	GoLocale string // Locale used in Go names, e.g. EnUS
	Descs    map[string]string
}{
	{"en_US", "EnUS", helpDescsEnUS}, // helpdescs_en_US.go
}
