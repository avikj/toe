package keeper

import (
	"context"

	"github.com/avikj/toe/x/toe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NewGame(goCtx context.Context, msg *types.MsgNewGame) (*types.MsgNewGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgNewGameResponse{}, nil
}
