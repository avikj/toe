package keeper

import (
	"context"
	"strconv"

	"github.com/avikj/toe/x/toe/gamelogic"
	"github.com/avikj/toe/x/toe/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlaceMarker(goCtx context.Context, msg *types.MsgPlaceMarker) (*types.MsgPlaceMarkerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	gameData, found := k.GetGameData(
		ctx,
		strconv.FormatUint(msg.GameId, 10),
	)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "game not found %s", msg.GameId)
	}

	if gameData.PlayerX == "" {
		return nil, sdkerrors.Wrapf(types.ErrWaitingForOpponent, 
			"game not started, waiting for player to join %s", msg.GameId)
	}

	//gameStatus, _ := gamelogic.GetStatus(gameData.BoardState)

	/*if gameStatus == gamelogic.XWon || gameStatus == gamelogic.OWon || gameStatus == gamelogic.Tie {
		return nil, sdkerrors.Wrapf(types.ErrGameOver, 
			"game already over, status=%s", gameStatus.String())
	}*/

	marker := gamelogic.EmptyMarker
	if msg.Creator == gameData.PlayerX {
		marker = gamelogic.XMarker
	} else if msg.Creator == gameData.PlayerO {
		marker = gamelogic.OMarker
	} else {
		return nil, sdkerrors.Wrapf(types.ErrNotInGame, "not participating in this game %s", msg.GameId)
	}

	newBoardState, err := gamelogic.PlaceMarker(gameData.BoardState, marker, uint(msg.Pos))

	if err != nil {
		return nil, err
	}

	gameData.BoardState = newBoardState

	newStatus, _ := gamelogic.GetStatus(newBoardState)
	k.SetGameData(ctx, gameData)
	
	return &types.MsgPlaceMarkerResponse{GameStatus: newStatus.String()}, nil
}
