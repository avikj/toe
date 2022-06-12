package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NextGameId:   &NextGameId{Value: uint64(DefaultIndex)},
		GameDataList: []GameData{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in gameData
	gameDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.GameDataList {
		index := string(GameDataKey(elem.Index))
		if _, ok := gameDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for gameData")
		}
		gameDataIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
