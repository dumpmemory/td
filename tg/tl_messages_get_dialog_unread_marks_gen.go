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

// MessagesGetDialogUnreadMarksRequest represents TL type `messages.getDialogUnreadMarks#21202222`.
// Get dialogs manually marked as unread
//
// See https://core.telegram.org/method/messages.getDialogUnreadMarks for reference.
type MessagesGetDialogUnreadMarksRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// ParentPeer field of MessagesGetDialogUnreadMarksRequest.
	//
	// Use SetParentPeer and GetParentPeer helpers.
	ParentPeer InputPeerClass
}

// MessagesGetDialogUnreadMarksRequestTypeID is TL type id of MessagesGetDialogUnreadMarksRequest.
const MessagesGetDialogUnreadMarksRequestTypeID = 0x21202222

// Ensuring interfaces in compile-time for MessagesGetDialogUnreadMarksRequest.
var (
	_ bin.Encoder     = &MessagesGetDialogUnreadMarksRequest{}
	_ bin.Decoder     = &MessagesGetDialogUnreadMarksRequest{}
	_ bin.BareEncoder = &MessagesGetDialogUnreadMarksRequest{}
	_ bin.BareDecoder = &MessagesGetDialogUnreadMarksRequest{}
)

func (g *MessagesGetDialogUnreadMarksRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ParentPeer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDialogUnreadMarksRequest) String() string {
	if g == nil {
		return "MessagesGetDialogUnreadMarksRequest(nil)"
	}
	type Alias MessagesGetDialogUnreadMarksRequest
	return fmt.Sprintf("MessagesGetDialogUnreadMarksRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetDialogUnreadMarksRequest from given interface.
func (g *MessagesGetDialogUnreadMarksRequest) FillFrom(from interface {
	GetParentPeer() (value InputPeerClass, ok bool)
}) {
	if val, ok := from.GetParentPeer(); ok {
		g.ParentPeer = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDialogUnreadMarksRequest) TypeID() uint32 {
	return MessagesGetDialogUnreadMarksRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDialogUnreadMarksRequest) TypeName() string {
	return "messages.getDialogUnreadMarks"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDialogUnreadMarksRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDialogUnreadMarks",
		ID:   MessagesGetDialogUnreadMarksRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ParentPeer",
			SchemaName: "parent_peer",
			Null:       !g.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *MessagesGetDialogUnreadMarksRequest) SetFlags() {
	if !(g.ParentPeer == nil) {
		g.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (g *MessagesGetDialogUnreadMarksRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogUnreadMarks#21202222 as nil")
	}
	b.PutID(MessagesGetDialogUnreadMarksRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetDialogUnreadMarksRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogUnreadMarks#21202222 as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDialogUnreadMarks#21202222: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		if g.ParentPeer == nil {
			return fmt.Errorf("unable to encode messages.getDialogUnreadMarks#21202222: field parent_peer is nil")
		}
		if err := g.ParentPeer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getDialogUnreadMarks#21202222: field parent_peer: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogUnreadMarksRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogUnreadMarks#21202222 to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogUnreadMarksRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogUnreadMarks#21202222: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetDialogUnreadMarksRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogUnreadMarks#21202222 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getDialogUnreadMarks#21202222: field flags: %w", err)
		}
	}
	if g.Flags.Has(0) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogUnreadMarks#21202222: field parent_peer: %w", err)
		}
		g.ParentPeer = value
	}
	return nil
}

// SetParentPeer sets value of ParentPeer conditional field.
func (g *MessagesGetDialogUnreadMarksRequest) SetParentPeer(value InputPeerClass) {
	g.Flags.Set(0)
	g.ParentPeer = value
}

// GetParentPeer returns value of ParentPeer conditional field and
// boolean which is true if field was set.
func (g *MessagesGetDialogUnreadMarksRequest) GetParentPeer() (value InputPeerClass, ok bool) {
	if g == nil {
		return
	}
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.ParentPeer, true
}

// MessagesGetDialogUnreadMarks invokes method messages.getDialogUnreadMarks#21202222 returning error if any.
// Get dialogs manually marked as unread
//
// See https://core.telegram.org/method/messages.getDialogUnreadMarks for reference.
func (c *Client) MessagesGetDialogUnreadMarks(ctx context.Context, request *MessagesGetDialogUnreadMarksRequest) ([]DialogPeerClass, error) {
	var result DialogPeerClassVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []DialogPeerClass(result.Elems), nil
}
