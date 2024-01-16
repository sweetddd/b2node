package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/testutil/bridge/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateCaller_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateCaller
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateCaller{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateCaller{
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

func TestMsgUpdateCaller_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateCaller
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateCaller{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateCaller{
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

func TestMsgDeleteCaller_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteCaller
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteCaller{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteCaller{
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
