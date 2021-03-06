package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/avikj/toe/testutil/keeper"
	"github.com/avikj/toe/testutil/nullify"
	"github.com/avikj/toe/x/toe/types"
)

func TestNextGameIdQuery(t *testing.T) {
	keeper, ctx := keepertest.ToeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNextGameId(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNextGameIdRequest
		response *types.QueryGetNextGameIdResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNextGameIdRequest{},
			response: &types.QueryGetNextGameIdResponse{NextGameId: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NextGameId(wctx, tc.request)
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
