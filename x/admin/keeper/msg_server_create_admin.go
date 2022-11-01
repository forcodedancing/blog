package keeper

import (
	"context"

	"blog/x/admin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAdmin(goCtx context.Context, msg *types.MsgCreateAdmin) (*types.MsgCreateAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Admin{
		Creator: msg.Creator,
		Title:   msg.Title,
		Gender:  msg.Gender,
	}
	id := k.AppendAdmin(ctx, post)

	post.Id = id
	err := ctx.EventManager().EmitTypedEvent(&post)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreateAdminResponse{Id: id}, nil
}
