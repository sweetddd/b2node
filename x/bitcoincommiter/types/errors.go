package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/bitcoincommiter module sentinel errors
var (
	ErrSample            = errorsmod.Register(ModuleName, 1100, "sample error")
	GetBitcoinRPCErr     = errorsmod.Register(ModuleName, 2001, "init bitcoin rpc fail.")
	ListUnspentErr       = errorsmod.Register(ModuleName, 2002, "list unspent fail.")
	DecodeUnspentHashErr = errorsmod.Register(ModuleName, 2003, "decode unspent hash fail.")
	CreatInscribeErr     = errorsmod.Register(ModuleName, 2004, "create inscription fail.")
	BackupRecoveryKeyErr = errorsmod.Register(ModuleName, 2005, "backup recovery key fail.")
	InscribeErr          = errorsmod.Register(ModuleName, 2006, "inscribe fail.")
)
