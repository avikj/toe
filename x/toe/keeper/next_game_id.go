package keeper

import (
	"github.com/avikj/toe/x/toe/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNextGameId set nextGameId in the store
func (k Keeper) SetNextGameId(ctx sdk.Context, nextGameId types.NextGameId) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextGameIdKey))
	b := k.cdc.MustMarshal(&nextGameId)
	store.Set([]byte{0}, b)
}

// GetNextGameId returns nextGameId
func (k Keeper) GetNextGameId(ctx sdk.Context) (val types.NextGameId, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextGameIdKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNextGameId removes nextGameId from the store
func (k Keeper) RemoveNextGameId(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextGameIdKey))
	store.Delete([]byte{0})
}
