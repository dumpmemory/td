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

// MessagesGetScheduledMessagesRequest represents TL type `messages.getScheduledMessages#bdbb0464`.
// Get scheduled messages
//
// See https://core.telegram.org/method/messages.getScheduledMessages for reference.
type MessagesGetScheduledMessagesRequest struct {
	// Peer
	Peer InputPeerClass
	// IDs of scheduled messages
	ID []int
}

// MessagesGetScheduledMessagesRequestTypeID is TL type id of MessagesGetScheduledMessagesRequest.
const MessagesGetScheduledMessagesRequestTypeID = 0xbdbb0464

func (g *MessagesGetScheduledMessagesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetScheduledMessagesRequest) String() string {
	if g == nil {
		return "MessagesGetScheduledMessagesRequest(nil)"
	}
	type Alias MessagesGetScheduledMessagesRequest
	return fmt.Sprintf("MessagesGetScheduledMessagesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetScheduledMessagesRequest from given interface.
func (g *MessagesGetScheduledMessagesRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetScheduledMessagesRequest) TypeID() uint32 {
	return MessagesGetScheduledMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetScheduledMessagesRequest) TypeName() string {
	return "messages.getScheduledMessages"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetScheduledMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getScheduledMessages",
		ID:   MessagesGetScheduledMessagesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetScheduledMessagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getScheduledMessages#bdbb0464 as nil")
	}
	b.PutID(MessagesGetScheduledMessagesRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getScheduledMessages#bdbb0464: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getScheduledMessages#bdbb0464: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetScheduledMessagesRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetID returns value of ID field.
func (g *MessagesGetScheduledMessagesRequest) GetID() (value []int) {
	return g.ID
}

// Decode implements bin.Decoder.
func (g *MessagesGetScheduledMessagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getScheduledMessages#bdbb0464 to nil")
	}
	if err := b.ConsumeID(MessagesGetScheduledMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getScheduledMessages#bdbb0464: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledMessages#bdbb0464: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledMessages#bdbb0464: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getScheduledMessages#bdbb0464: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetScheduledMessagesRequest.
var (
	_ bin.Encoder = &MessagesGetScheduledMessagesRequest{}
	_ bin.Decoder = &MessagesGetScheduledMessagesRequest{}
)

// MessagesGetScheduledMessages invokes method messages.getScheduledMessages#bdbb0464 returning error if any.
// Get scheduled messages
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getScheduledMessages for reference.
func (c *Client) MessagesGetScheduledMessages(ctx context.Context, request *MessagesGetScheduledMessagesRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
