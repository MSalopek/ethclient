package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateStorage = "create_storage"
)

var _ sdk.Msg = &MsgCreateStorage{}

func NewMsgCreateStorage(creator, address, storage, block string) *MsgCreateStorage {
	return &MsgCreateStorage{
		Creator: creator,
		Address: address,
		Storage: storage,
		Block:   block,
	}
}

func (msg *MsgCreateStorage) Route() string {
	return RouterKey
}

func (msg *MsgCreateStorage) Type() string {
	return TypeMsgCreateStorage
}

func (msg *MsgCreateStorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
