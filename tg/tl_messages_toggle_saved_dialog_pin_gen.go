// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessagesToggleSavedDialogPinRequest represents TL type `messages.toggleSavedDialogPin#ac81bbde`.
//
// See https://core.telegram.org/method/messages.toggleSavedDialogPin for reference.
type MessagesToggleSavedDialogPinRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Pinned field of MessagesToggleSavedDialogPinRequest.
	Pinned bool
	// Peer field of MessagesToggleSavedDialogPinRequest.
	Peer InputDialogPeerClass
}

// MessagesToggleSavedDialogPinRequestTypeID is TL type id of MessagesToggleSavedDialogPinRequest.
const MessagesToggleSavedDialogPinRequestTypeID = 0xac81bbde

// Ensuring interfaces in compile-time for MessagesToggleSavedDialogPinRequest.
var (
	_ bin.Encoder     = &MessagesToggleSavedDialogPinRequest{}
	_ bin.Decoder     = &MessagesToggleSavedDialogPinRequest{}
	_ bin.BareEncoder = &MessagesToggleSavedDialogPinRequest{}
	_ bin.BareDecoder = &MessagesToggleSavedDialogPinRequest{}
)

func (t *MessagesToggleSavedDialogPinRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Pinned == false) {
		return false
	}
	if !(t.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesToggleSavedDialogPinRequest) String() string {
	if t == nil {
		return "MessagesToggleSavedDialogPinRequest(nil)"
	}
	type Alias MessagesToggleSavedDialogPinRequest
	return fmt.Sprintf("MessagesToggleSavedDialogPinRequest%+v", Alias(*t))
}

// FillFrom fills MessagesToggleSavedDialogPinRequest from given interface.
func (t *MessagesToggleSavedDialogPinRequest) FillFrom(from interface {
	GetPinned() (value bool)
	GetPeer() (value InputDialogPeerClass)
}) {
	t.Pinned = from.GetPinned()
	t.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesToggleSavedDialogPinRequest) TypeID() uint32 {
	return MessagesToggleSavedDialogPinRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesToggleSavedDialogPinRequest) TypeName() string {
	return "messages.toggleSavedDialogPin"
}

// TypeInfo returns info about TL type.
func (t *MessagesToggleSavedDialogPinRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.toggleSavedDialogPin",
		ID:   MessagesToggleSavedDialogPinRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pinned",
			SchemaName: "pinned",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *MessagesToggleSavedDialogPinRequest) SetFlags() {
	if !(t.Pinned == false) {
		t.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (t *MessagesToggleSavedDialogPinRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleSavedDialogPin#ac81bbde as nil")
	}
	b.PutID(MessagesToggleSavedDialogPinRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesToggleSavedDialogPinRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleSavedDialogPin#ac81bbde as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleSavedDialogPin#ac81bbde: field flags: %w", err)
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode messages.toggleSavedDialogPin#ac81bbde: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleSavedDialogPin#ac81bbde: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesToggleSavedDialogPinRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleSavedDialogPin#ac81bbde to nil")
	}
	if err := b.ConsumeID(MessagesToggleSavedDialogPinRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.toggleSavedDialogPin#ac81bbde: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesToggleSavedDialogPinRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleSavedDialogPin#ac81bbde to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.toggleSavedDialogPin#ac81bbde: field flags: %w", err)
		}
	}
	t.Pinned = t.Flags.Has(0)
	{
		value, err := DecodeInputDialogPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleSavedDialogPin#ac81bbde: field peer: %w", err)
		}
		t.Peer = value
	}
	return nil
}

// SetPinned sets value of Pinned conditional field.
func (t *MessagesToggleSavedDialogPinRequest) SetPinned(value bool) {
	if value {
		t.Flags.Set(0)
		t.Pinned = true
	} else {
		t.Flags.Unset(0)
		t.Pinned = false
	}
}

// GetPinned returns value of Pinned conditional field.
func (t *MessagesToggleSavedDialogPinRequest) GetPinned() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (t *MessagesToggleSavedDialogPinRequest) GetPeer() (value InputDialogPeerClass) {
	if t == nil {
		return
	}
	return t.Peer
}

// MessagesToggleSavedDialogPin invokes method messages.toggleSavedDialogPin#ac81bbde returning error if any.
//
// See https://core.telegram.org/method/messages.toggleSavedDialogPin for reference.
// Can be used by bots.
func (c *Client) MessagesToggleSavedDialogPin(ctx context.Context, request *MessagesToggleSavedDialogPinRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}