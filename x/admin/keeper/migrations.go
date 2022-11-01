package keeper

import (
	v2 "blog/x/admin/migrations/v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) MigrateV1toV2(ctx sdk.Context) error {
	// if we want to upgrade parameters
	//if err := v1.UpdateParams(ctx, &m.keeper.paramstore); err != nil {
	//	return err
	//}

	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
