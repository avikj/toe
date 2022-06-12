package toe

import (
	"github.com/avikj/toe/x/toe/keeper"
	"github.com/avikj/toe/x/toe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.NextGameId != nil {
		k.SetNextGameId(ctx, *genState.NextGameId)
	}
	// Set all the gameData
	for _, elem := range genState.GameDataList {
		k.SetGameData(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all nextGameId
	nextGameId, found := k.GetNextGameId(ctx)
	if found {
		genesis.NextGameId = &nextGameId
	}
	genesis.GameDataList = k.GetAllGameData(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
