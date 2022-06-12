package keeper_test

import (
	"testing"

	testkeeper "github.com/avikj/toe/testutil/keeper"
	"github.com/avikj/toe/x/toe/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ToeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
