package toe

import (
	"math/rand"

	"github.com/avikj/toe/testutil/sample"
	toesimulation "github.com/avikj/toe/x/toe/simulation"
	"github.com/avikj/toe/x/toe/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = toesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgNewGame = "op_weight_msg_new_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgNewGame int = 100

	opWeightMsgJoinGame = "op_weight_msg_join_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgJoinGame int = 100

	opWeightMsgPlaceMarker = "op_weight_msg_place_marker"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPlaceMarker int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	toeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&toeGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgNewGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgNewGame, &weightMsgNewGame, nil,
		func(_ *rand.Rand) {
			weightMsgNewGame = defaultWeightMsgNewGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgNewGame,
		toesimulation.SimulateMsgNewGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgJoinGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgJoinGame, &weightMsgJoinGame, nil,
		func(_ *rand.Rand) {
			weightMsgJoinGame = defaultWeightMsgJoinGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgJoinGame,
		toesimulation.SimulateMsgJoinGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPlaceMarker int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPlaceMarker, &weightMsgPlaceMarker, nil,
		func(_ *rand.Rand) {
			weightMsgPlaceMarker = defaultWeightMsgPlaceMarker
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPlaceMarker,
		toesimulation.SimulateMsgPlaceMarker(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
