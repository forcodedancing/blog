package keeper

import (
	"blog/x/admin/types"
)

var _ types.QueryServer = Keeper{}
