package keeper

import (
	"context"

	"blog/x/blog/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var post = types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}
	id := k.AppendPost(ctx, post)

	post.Id = id
	err := ctx.EventManager().EmitTypedEvent(&post)
	if err != nil {
		return nil, err
	}

	return &types.MsgCreatePostResponse{Id: id}, nil
}
