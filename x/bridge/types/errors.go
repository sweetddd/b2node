package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/bridge module sentinel errors
var (
	ErrIndexExist            = errorsmod.Register(ModuleName, 1100, "index already exist")
	ErrIndexNotExist         = errorsmod.Register(ModuleName, 1101, "index does not exist")
	ErrUnauthorized          = errorsmod.Register(ModuleName, 1102, "only admin can  do this")
	ErrNotCallerGroupMembers = errorsmod.Register(ModuleName, 1103, "only caller group members can do this")
	ErrInvalidStatus         = errorsmod.Register(ModuleName, 1104, "current status is incorrect to do this")
	ErrNotOwner              = errorsmod.Register(ModuleName, 1105, "only owner can do this")
	ErrNotSignerGroupMembers = errorsmod.Register(ModuleName, 1106, "only signer group members can do this")
	ErrAlreadySigned         = errorsmod.Register(ModuleName, 1107, "this sender already signed")
)
