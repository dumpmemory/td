// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// PaymentsValidateRequestedInfoRequest represents TL type `payments.validateRequestedInfo#770a8e74`.
// Submit requested order information for validation
//
// See https://core.telegram.org/method/payments.validateRequestedInfo for reference.
type PaymentsValidateRequestedInfoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Save order information to re-use it for future orders
	Save bool
	// Message ID of payment form
	MsgID int
	// Requested order information
	Info PaymentRequestedInfo
}

// PaymentsValidateRequestedInfoRequestTypeID is TL type id of PaymentsValidateRequestedInfoRequest.
const PaymentsValidateRequestedInfoRequestTypeID = 0x770a8e74

func (v *PaymentsValidateRequestedInfoRequest) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Flags.Zero()) {
		return false
	}
	if !(v.Save == false) {
		return false
	}
	if !(v.MsgID == 0) {
		return false
	}
	if !(v.Info.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *PaymentsValidateRequestedInfoRequest) String() string {
	if v == nil {
		return "PaymentsValidateRequestedInfoRequest(nil)"
	}
	type Alias PaymentsValidateRequestedInfoRequest
	return fmt.Sprintf("PaymentsValidateRequestedInfoRequest%+v", Alias(*v))
}

// FillFrom fills PaymentsValidateRequestedInfoRequest from given interface.
func (v *PaymentsValidateRequestedInfoRequest) FillFrom(from interface {
	GetSave() (value bool)
	GetMsgID() (value int)
	GetInfo() (value PaymentRequestedInfo)
}) {
	v.Save = from.GetSave()
	v.MsgID = from.GetMsgID()
	v.Info = from.GetInfo()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsValidateRequestedInfoRequest) TypeID() uint32 {
	return PaymentsValidateRequestedInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsValidateRequestedInfoRequest) TypeName() string {
	return "payments.validateRequestedInfo"
}

// TypeInfo returns info about TL type.
func (v *PaymentsValidateRequestedInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.validateRequestedInfo",
		ID:   PaymentsValidateRequestedInfoRequestTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "Save",
			SchemaName: "save",
			Null:       !v.Flags.Has(0),
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Info",
			SchemaName: "info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *PaymentsValidateRequestedInfoRequest) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode payments.validateRequestedInfo#770a8e74 as nil")
	}
	b.PutID(PaymentsValidateRequestedInfoRequestTypeID)
	if !(v.Save == false) {
		v.Flags.Set(0)
	}
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validateRequestedInfo#770a8e74: field flags: %w", err)
	}
	b.PutInt(v.MsgID)
	if err := v.Info.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.validateRequestedInfo#770a8e74: field info: %w", err)
	}
	return nil
}

// SetSave sets value of Save conditional field.
func (v *PaymentsValidateRequestedInfoRequest) SetSave(value bool) {
	if value {
		v.Flags.Set(0)
		v.Save = true
	} else {
		v.Flags.Unset(0)
		v.Save = false
	}
}

// GetSave returns value of Save conditional field.
func (v *PaymentsValidateRequestedInfoRequest) GetSave() (value bool) {
	return v.Flags.Has(0)
}

// GetMsgID returns value of MsgID field.
func (v *PaymentsValidateRequestedInfoRequest) GetMsgID() (value int) {
	return v.MsgID
}

// GetInfo returns value of Info field.
func (v *PaymentsValidateRequestedInfoRequest) GetInfo() (value PaymentRequestedInfo) {
	return v.Info
}

// Decode implements bin.Decoder.
func (v *PaymentsValidateRequestedInfoRequest) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode payments.validateRequestedInfo#770a8e74 to nil")
	}
	if err := b.ConsumeID(PaymentsValidateRequestedInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.validateRequestedInfo#770a8e74: %w", err)
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#770a8e74: field flags: %w", err)
		}
	}
	v.Save = v.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#770a8e74: field msg_id: %w", err)
		}
		v.MsgID = value
	}
	{
		if err := v.Info.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.validateRequestedInfo#770a8e74: field info: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsValidateRequestedInfoRequest.
var (
	_ bin.Encoder = &PaymentsValidateRequestedInfoRequest{}
	_ bin.Decoder = &PaymentsValidateRequestedInfoRequest{}
)

// PaymentsValidateRequestedInfo invokes method payments.validateRequestedInfo#770a8e74 returning error if any.
// Submit requested order information for validation
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//
// See https://core.telegram.org/method/payments.validateRequestedInfo for reference.
func (c *Client) PaymentsValidateRequestedInfo(ctx context.Context, request *PaymentsValidateRequestedInfoRequest) (*PaymentsValidatedRequestedInfo, error) {
	var result PaymentsValidatedRequestedInfo

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
