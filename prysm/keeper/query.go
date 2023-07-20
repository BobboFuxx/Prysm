package keeper

import (
	"prysm/x/prysm/types"
)

var _ types.QueryServer = Keeper{}
