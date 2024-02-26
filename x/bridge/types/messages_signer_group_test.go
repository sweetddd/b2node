package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/evmos/ethermint/testutil/bridge/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateSignerGroup_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSignerGroup
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSignerGroup{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSignerGroup{
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

func TestMsgUpdateSignerGroup_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSignerGroup
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateSignerGroup{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateSignerGroup{
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

func TestMsgDeleteSignerGroup_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteSignerGroup
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteSignerGroup{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteSignerGroup{
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
