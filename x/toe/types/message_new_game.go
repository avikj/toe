package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewGame = "new_game"

var _ sdk.Msg = &MsgNewGame{}

func NewMsgNewGame(creator string) *MsgNewGame {
	return &MsgNewGame{
		Creator: creator,
	}
}

func (msg *MsgNewGame) Route() string {
	return RouterKey
}

func (msg *MsgNewGame) Type() string {
	return TypeMsgNewGame
}

func (msg *MsgNewGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
