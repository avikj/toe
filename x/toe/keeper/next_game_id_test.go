package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/avikj/toe/testutil/keeper"
	"github.com/avikj/toe/testutil/nullify"
	"github.com/avikj/toe/x/toe/keeper"
	"github.com/avikj/toe/x/toe/types"
)

func createTestNextGameId(keeper *keeper.Keeper, ctx sdk.Context) types.NextGameId {
	item := types.NextGameId{}
	keeper.SetNextGameId(ctx, item)
	return item
}

func TestNextGameIdGet(t *testing.T) {
	keeper, ctx := keepertest.ToeKeeper(t)
	item := createTestNextGameId(keeper, ctx)
	rst, found := keeper.GetNextGameId(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestNextGameIdRemove(t *testing.T) {
	keeper, ctx := keepertest.ToeKeeper(t)
	createTestNextGameId(keeper, ctx)
	keeper.RemoveNextGameId(ctx)
	_, found := keeper.GetNextGameId(ctx)
	require.False(t, found)
}
