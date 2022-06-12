package toe_test

import (
	"testing"

	keepertest "github.com/avikj/toe/testutil/keeper"
	"github.com/avikj/toe/testutil/nullify"
	"github.com/avikj/toe/x/toe"
	"github.com/avikj/toe/x/toe/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NextGameId: &types.NextGameId{
			Value: 67,
		},
		GameDataList: []types.GameData{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ToeKeeper(t)
	toe.InitGenesis(ctx, *k, genesisState)
	got := toe.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.NextGameId, got.NextGameId)
	require.ElementsMatch(t, genesisState.GameDataList, got.GameDataList)
	// this line is used by starport scaffolding # genesis/test/assert
}
