package keeper

import (
	"context"

	"github.com/avikj/toe/x/toe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlaceMarker(goCtx context.Context, msg *types.MsgPlaceMarker) (*types.MsgPlaceMarkerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgPlaceMarkerResponse{}, nil
}
