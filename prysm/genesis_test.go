package prysm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "prysm/testutil/keeper"
	"prysm/testutil/nullify"
	"prysm/x/prysm"
	"prysm/x/prysm/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PrysmKeeper(t)
	prysm.InitGenesis(ctx, *k, genesisState)
	got := prysm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
