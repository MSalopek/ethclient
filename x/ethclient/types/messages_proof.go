package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgStoreProof = "store_proof"
)

var _ sdk.Msg = &MsgStoreProof{}

func NewMsgStoreProof(creator, address, storage string, block int64) *MsgStoreProof {
	return &MsgStoreProof{
		Creator: creator,
		Address: address,
		Storage: storage,
		Block:   block,
	}
}

func (msg *MsgStoreProof) Route() string {
	return RouterKey
}

func (msg *MsgStoreProof) Type() string {
	return TypeMsgStoreProof
}

func (msg *MsgStoreProof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStoreProof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStoreProof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Address == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "address cannot be empty")
	}
	if msg.Storage == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "storage cannot be empty")
	}
	// check that storage is a hexadecimal string
	if !strings.HasPrefix(msg.Storage, "0x") {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "storage must be a hexadecimal string ( start with 0x )")
	}

	if msg.Block == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "block cannot be empty")
	}
	return nil
}
