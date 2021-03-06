package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/avikj/toe/testutil/network"
	"github.com/avikj/toe/testutil/nullify"
	"github.com/avikj/toe/x/toe/client/cli"
	"github.com/avikj/toe/x/toe/types"
)

func networkWithNextGameIdObjects(t *testing.T) (*network.Network, types.NextGameId) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	nextGameId := &types.NextGameId{}
	nullify.Fill(&nextGameId)
	state.NextGameId = nextGameId
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.NextGameId
}

func TestShowNextGameId(t *testing.T) {
	net, obj := networkWithNextGameIdObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.NextGameId
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowNextGameId(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetNextGameIdResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.NextGameId)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.NextGameId),
				)
			}
		})
	}
}
