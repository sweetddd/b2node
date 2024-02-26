package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/testutil/bridge/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateCallerGroup_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCallerGroup
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCallerGroup{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCallerGroup{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateCallerGroup_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCallerGroup
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCallerGroup{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCallerGroup{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteCallerGroup_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteCallerGroup
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteCallerGroup{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteCallerGroup{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
