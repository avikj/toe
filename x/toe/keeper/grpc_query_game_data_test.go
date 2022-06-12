package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/avikj/toe/testutil/keeper"
	"github.com/avikj/toe/testutil/nullify"
	"github.com/avikj/toe/x/toe/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestGameDataQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ToeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGameData(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetGameDataRequest
		response *types.QueryGetGameDataResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGameDataRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetGameDataResponse{GameData: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGameDataRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetGameDataResponse{GameData: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGameDataRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.GameData(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestGameDataQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ToeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNGameData(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGameDataRequest {
		return &types.QueryAllGameDataRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GameDataAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GameData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GameData),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.GameDataAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GameData), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GameData),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.GameDataAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.GameData),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.GameDataAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
