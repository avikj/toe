package keeper

import (
	"context"
	"crypto/sha256"
	"strconv"
	"github.com/avikj/toe/x/toe/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) JoinGame(goCtx context.Context, msg *types.MsgJoinGame) (*types.MsgJoinGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// msg.Creator
	//msg.GameId

	gameData, found := k.GetGameData(
		ctx,
		strconv.FormatUint(msg.GameId, 10),
	)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game not found %s", msg.GameId)
	}
	if gameData.PlayerX != "" {
		return nil, sdkerrors.Wrapf(types.ErrGameNotOpen, "game not open %s", msg.GameId)
	}
	if gameData.Creator == msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrCantPlaySelf, "can't play against self %s", msg.GameId)
	}

	if (sha256.Sum256([]byte(gameData.Creator+msg.Creator))[0] & 128) != 0 {
		gameData.PlayerX = gameData.Creator
		gameData.PlayerO = msg.Creator
	} else {
		gameData.PlayerX = msg.Creator
		gameData.PlayerO = gameData.Creator
	}

	k.SetGameData(ctx, gameData)

	return &types.MsgJoinGameResponse{
		PlayerX: gameData.PlayerX, 
		PlayerO: gameData.PlayerO,
	}, nil
}
