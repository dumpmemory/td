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

// PhoneLeaveGroupCallRequest represents TL type `phone.leaveGroupCall#500377f9`.
//
// See https://core.telegram.org/method/phone.leaveGroupCall for reference.
type PhoneLeaveGroupCallRequest struct {
	// Call field of PhoneLeaveGroupCallRequest.
	Call InputGroupCall
	// Source field of PhoneLeaveGroupCallRequest.
	Source int
}

// PhoneLeaveGroupCallRequestTypeID is TL type id of PhoneLeaveGroupCallRequest.
const PhoneLeaveGroupCallRequestTypeID = 0x500377f9

func (l *PhoneLeaveGroupCallRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Call.Zero()) {
		return false
	}
	if !(l.Source == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *PhoneLeaveGroupCallRequest) String() string {
	if l == nil {
		return "PhoneLeaveGroupCallRequest(nil)"
	}
	type Alias PhoneLeaveGroupCallRequest
	return fmt.Sprintf("PhoneLeaveGroupCallRequest%+v", Alias(*l))
}

// FillFrom fills PhoneLeaveGroupCallRequest from given interface.
func (l *PhoneLeaveGroupCallRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
	GetSource() (value int)
}) {
	l.Call = from.GetCall()
	l.Source = from.GetSource()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneLeaveGroupCallRequest) TypeID() uint32 {
	return PhoneLeaveGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneLeaveGroupCallRequest) TypeName() string {
	return "phone.leaveGroupCall"
}

// TypeInfo returns info about TL type.
func (l *PhoneLeaveGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.leaveGroupCall",
		ID:   PhoneLeaveGroupCallRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Source",
			SchemaName: "source",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *PhoneLeaveGroupCallRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode phone.leaveGroupCall#500377f9 as nil")
	}
	b.PutID(PhoneLeaveGroupCallRequestTypeID)
	if err := l.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.leaveGroupCall#500377f9: field call: %w", err)
	}
	b.PutInt(l.Source)
	return nil
}

// GetCall returns value of Call field.
func (l *PhoneLeaveGroupCallRequest) GetCall() (value InputGroupCall) {
	return l.Call
}

// GetSource returns value of Source field.
func (l *PhoneLeaveGroupCallRequest) GetSource() (value int) {
	return l.Source
}

// Decode implements bin.Decoder.
func (l *PhoneLeaveGroupCallRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode phone.leaveGroupCall#500377f9 to nil")
	}
	if err := b.ConsumeID(PhoneLeaveGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.leaveGroupCall#500377f9: %w", err)
	}
	{
		if err := l.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.leaveGroupCall#500377f9: field call: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.leaveGroupCall#500377f9: field source: %w", err)
		}
		l.Source = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneLeaveGroupCallRequest.
var (
	_ bin.Encoder = &PhoneLeaveGroupCallRequest{}
	_ bin.Decoder = &PhoneLeaveGroupCallRequest{}
)

// PhoneLeaveGroupCall invokes method phone.leaveGroupCall#500377f9 returning error if any.
//
// See https://core.telegram.org/method/phone.leaveGroupCall for reference.
func (c *Client) PhoneLeaveGroupCall(ctx context.Context, request *PhoneLeaveGroupCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
