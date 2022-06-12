package keeper

import (
	"context"

	"github.com/avikj/toe/x/toe/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GameDataAll(c context.Context, req *types.QueryAllGameDataRequest) (*types.QueryAllGameDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var gameDatas []types.GameData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	gameDataStore := prefix.NewStore(store, types.KeyPrefix(types.GameDataKeyPrefix))

	pageRes, err := query.Paginate(gameDataStore, req.Pagination, func(key []byte, value []byte) error {
		var gameData types.GameData
		if err := k.cdc.Unmarshal(value, &gameData); err != nil {
			return err
		}

		gameDatas = append(gameDatas, gameData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGameDataResponse{GameData: gameDatas, Pagination: pageRes}, nil
}

func (k Keeper) GameData(c context.Context, req *types.QueryGetGameDataRequest) (*types.QueryGetGameDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetGameData(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGameDataResponse{GameData: val}, nil
}
