package keeper

import (
	"context"

	"github.com/avikj/toe/x/toe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NextGameId(c context.Context, req *types.QueryGetNextGameIdRequest) (*types.QueryGetNextGameIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNextGameId(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNextGameIdResponse{NextGameId: val}, nil
}
