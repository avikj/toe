package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPlaceMarker = "place_marker"

var _ sdk.Msg = &MsgPlaceMarker{}

func NewMsgPlaceMarker(creator string, gameId uint64, pos uint64) *MsgPlaceMarker {
	return &MsgPlaceMarker{
		Creator: creator,
		GameId:  gameId,
		Pos:     pos,
	}
}

func (msg *MsgPlaceMarker) Route() string {
	return RouterKey
}

func (msg *MsgPlaceMarker) Type() string {
	return TypeMsgPlaceMarker
}

func (msg *MsgPlaceMarker) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPlaceMarker) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPlaceMarker) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
