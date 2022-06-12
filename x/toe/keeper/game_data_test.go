package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/avikj/toe/testutil/keeper"
	"github.com/avikj/toe/testutil/nullify"
	"github.com/avikj/toe/x/toe/keeper"
	"github.com/avikj/toe/x/toe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGameData(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.GameData {
	items := make([]types.GameData, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetGameData(ctx, items[i])
	}
	return items
}

func TestGameDataGet(t *testing.T) {
	keeper, ctx := keepertest.ToeKeeper(t)
	items := createNGameData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGameData(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGameDataRemove(t *testing.T) {
	keeper, ctx := keepertest.ToeKeeper(t)
	items := createNGameData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGameData(ctx,
			item.Index,
		)
		_, found := keeper.GetGameData(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestGameDataGetAll(t *testing.T) {
	keeper, ctx := keepertest.ToeKeeper(t)
	items := createNGameData(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGameData(ctx)),
	)
}
