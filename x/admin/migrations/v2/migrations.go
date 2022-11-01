package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

func UpdateParams(ctx sdk.Context, paramStore *paramtypes.Subspace) error {
	//set new param values
	//paramStore.Set(ctx, types.KeyValue, "newvalue")
	return nil
}

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	if err := migrateValuesWithPrefix(store, cdc); err != nil {
		return err
	}
	//other logics
	return nil
}

func migrateValuesWithPrefix(store sdk.KVStore, cdc codec.BinaryCodec) error {
	oldStoreIter := store.Iterator(nil, nil)

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		oldKey := oldStoreIter.Key()
		//oldVal := store.Get(oldKey)

		// we can transform data here if wanted
		//newKey, newVal := migrateValue(cdc, oldKey, oldVal)
		//store.Set(newKey, newVal)
		store.Delete(oldKey) // Delete old key, value
	}

	return nil
}

func migrateValue(cdc codec.BinaryCodec, oldKey []byte, oldVal []byte) (newKey []byte, newVal []byte) {
	newKey = oldKey
	newVal = oldVal
	return
}
