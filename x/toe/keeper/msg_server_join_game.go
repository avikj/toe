package keeper

import (
	"context"

	"github.com/avikj/toe/x/toe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) JoinGame(goCtx context.Context, msg *types.MsgJoinGame) (*types.MsgJoinGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgJoinGameResponse{}, nil
}
