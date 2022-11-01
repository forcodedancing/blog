package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateAdmin = "create_admin"

var _ sdk.Msg = &MsgCreateAdmin{}

func NewMsgCreateAdmin(creator string, title string, gender string) *MsgCreateAdmin {
	return &MsgCreateAdmin{
		Creator: creator,
		Title:   title,
		Gender:  gender,
	}
}

func (msg *MsgCreateAdmin) Route() string {
	return RouterKey
}

func (msg *MsgCreateAdmin) Type() string {
	return TypeMsgCreateAdmin
}

func (msg *MsgCreateAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
