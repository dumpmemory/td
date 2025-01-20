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

// MessagesToggleNoForwardsRequest represents TL type `messages.toggleNoForwards#b11eafa2`.
// Enable or disable content protection¹ on a channel or chat
//
// Links:
//  1. https://telegram.org/blog/protected-content-delete-by-date-and-more
//
// See https://core.telegram.org/method/messages.toggleNoForwards for reference.
type MessagesToggleNoForwardsRequest struct {
	// The chat or channel
	Peer InputPeerClass
	// Enable or disable content protection
	Enabled bool
}

// MessagesToggleNoForwardsRequestTypeID is TL type id of MessagesToggleNoForwardsRequest.
const MessagesToggleNoForwardsRequestTypeID = 0xb11eafa2

// Ensuring interfaces in compile-time for MessagesToggleNoForwardsRequest.
var (
	_ bin.Encoder     = &MessagesToggleNoForwardsRequest{}
	_ bin.Decoder     = &MessagesToggleNoForwardsRequest{}
	_ bin.BareEncoder = &MessagesToggleNoForwardsRequest{}
	_ bin.BareDecoder = &MessagesToggleNoForwardsRequest{}
)

func (t *MessagesToggleNoForwardsRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Peer == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesToggleNoForwardsRequest) String() string {
	if t == nil {
		return "MessagesToggleNoForwardsRequest(nil)"
	}
	type Alias MessagesToggleNoForwardsRequest
	return fmt.Sprintf("MessagesToggleNoForwardsRequest%+v", Alias(*t))
}

// FillFrom fills MessagesToggleNoForwardsRequest from given interface.
func (t *MessagesToggleNoForwardsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetEnabled() (value bool)
}) {
	t.Peer = from.GetPeer()
	t.Enabled = from.GetEnabled()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesToggleNoForwardsRequest) TypeID() uint32 {
	return MessagesToggleNoForwardsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesToggleNoForwardsRequest) TypeName() string {
	return "messages.toggleNoForwards"
}

// TypeInfo returns info about TL type.
func (t *MessagesToggleNoForwardsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.toggleNoForwards",
		ID:   MessagesToggleNoForwardsRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *MessagesToggleNoForwardsRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleNoForwards#b11eafa2 as nil")
	}
	b.PutID(MessagesToggleNoForwardsRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesToggleNoForwardsRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleNoForwards#b11eafa2 as nil")
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode messages.toggleNoForwards#b11eafa2: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleNoForwards#b11eafa2: field peer: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesToggleNoForwardsRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleNoForwards#b11eafa2 to nil")
	}
	if err := b.ConsumeID(MessagesToggleNoForwardsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.toggleNoForwards#b11eafa2: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesToggleNoForwardsRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleNoForwards#b11eafa2 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleNoForwards#b11eafa2: field peer: %w", err)
		}
		t.Peer = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleNoForwards#b11eafa2: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (t *MessagesToggleNoForwardsRequest) GetPeer() (value InputPeerClass) {
	if t == nil {
		return
	}
	return t.Peer
}

// GetEnabled returns value of Enabled field.
func (t *MessagesToggleNoForwardsRequest) GetEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.Enabled
}

// MessagesToggleNoForwards invokes method messages.toggleNoForwards#b11eafa2 returning error if any.
// Enable or disable content protection¹ on a channel or chat
//
// Links:
//  1. https://telegram.org/blog/protected-content-delete-by-date-and-more
//
// Possible errors:
//
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_NOT_MODIFIED: No changes were made to chat information because the new information you passed is identical to the current information.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.toggleNoForwards for reference.
func (c *Client) MessagesToggleNoForwards(ctx context.Context, request *MessagesToggleNoForwardsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
