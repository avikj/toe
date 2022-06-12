package keeper

import (
	"github.com/avikj/toe/x/toe/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetGameData set a specific gameData in the store from its index
func (k Keeper) SetGameData(ctx sdk.Context, gameData types.GameData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameDataKeyPrefix))
	b := k.cdc.MustMarshal(&gameData)
	store.Set(types.GameDataKey(
		gameData.Index,
	), b)
}

// GetGameData returns a gameData from its index
func (k Keeper) GetGameData(
	ctx sdk.Context,
	index string,

) (val types.GameData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameDataKeyPrefix))

	b := store.Get(types.GameDataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGameData removes a gameData from the store
func (k Keeper) RemoveGameData(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameDataKeyPrefix))
	store.Delete(types.GameDataKey(
		index,
	))
}

// GetAllGameData returns all gameData
func (k Keeper) GetAllGameData(ctx sdk.Context) (list []types.GameData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GameDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.GameData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
