package keeper

import (
	"context"
	"strconv"
	"github.com/avikj/toe/x/toe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NewGame(goCtx context.Context, msg *types.MsgNewGame) (*types.MsgNewGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	nextGameIdStruct, _ := k.GetNextGameId(ctx); // get nextGameId value from store

	// create new gameData entry {creator: message.Creator, playerX: “”, playerO: “”, boardState: “_________”, id: nextGameIdValue}
	newGameData := types.GameData{
		Index: strconv.FormatUint(nextGameIdStruct.Value, 10),
		GameId: nextGameIdStruct.Value,
		Creator: msg.Creator,
		PlayerX: "",
		PlayerO: "",
		BoardState: "_________",
	}

	nextGameIdStruct.Value++;
	k.SetNextGameId(ctx, nextGameIdStruct)// increment nextGameId value in store
	// place new gameData entry in store
	k.SetGameData(ctx, newGameData)

	//emit event?
	return &types.MsgNewGameResponse{GameId: nextGameIdStruct.Value-1}, nil
}
 
