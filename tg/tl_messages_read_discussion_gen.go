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

// MessagesReadDiscussionRequest represents TL type `messages.readDiscussion#f731a9f4`.
// Mark a thread¹ as read
//
// Links:
//  1. https://core.telegram.org/api/threads
//
// See https://core.telegram.org/method/messages.readDiscussion for reference.
type MessagesReadDiscussionRequest struct {
	// Group ID
	Peer InputPeerClass
	// ID of message that started the thread
	MsgID int
	// ID up to which thread messages were read
	ReadMaxID int
}

// MessagesReadDiscussionRequestTypeID is TL type id of MessagesReadDiscussionRequest.
const MessagesReadDiscussionRequestTypeID = 0xf731a9f4

// Ensuring interfaces in compile-time for MessagesReadDiscussionRequest.
var (
	_ bin.Encoder     = &MessagesReadDiscussionRequest{}
	_ bin.Decoder     = &MessagesReadDiscussionRequest{}
	_ bin.BareEncoder = &MessagesReadDiscussionRequest{}
	_ bin.BareDecoder = &MessagesReadDiscussionRequest{}
)

func (r *MessagesReadDiscussionRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.MsgID == 0) {
		return false
	}
	if !(r.ReadMaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReadDiscussionRequest) String() string {
	if r == nil {
		return "MessagesReadDiscussionRequest(nil)"
	}
	type Alias MessagesReadDiscussionRequest
	return fmt.Sprintf("MessagesReadDiscussionRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReadDiscussionRequest from given interface.
func (r *MessagesReadDiscussionRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetReadMaxID() (value int)
}) {
	r.Peer = from.GetPeer()
	r.MsgID = from.GetMsgID()
	r.ReadMaxID = from.GetReadMaxID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReadDiscussionRequest) TypeID() uint32 {
	return MessagesReadDiscussionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReadDiscussionRequest) TypeName() string {
	return "messages.readDiscussion"
}

// TypeInfo returns info about TL type.
func (r *MessagesReadDiscussionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.readDiscussion",
		ID:   MessagesReadDiscussionRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "ReadMaxID",
			SchemaName: "read_max_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReadDiscussionRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readDiscussion#f731a9f4 as nil")
	}
	b.PutID(MessagesReadDiscussionRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReadDiscussionRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readDiscussion#f731a9f4 as nil")
	}
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.readDiscussion#f731a9f4: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.readDiscussion#f731a9f4: field peer: %w", err)
	}
	b.PutInt(r.MsgID)
	b.PutInt(r.ReadMaxID)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReadDiscussionRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readDiscussion#f731a9f4 to nil")
	}
	if err := b.ConsumeID(MessagesReadDiscussionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readDiscussion#f731a9f4: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReadDiscussionRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readDiscussion#f731a9f4 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.readDiscussion#f731a9f4: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readDiscussion#f731a9f4: field msg_id: %w", err)
		}
		r.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readDiscussion#f731a9f4: field read_max_id: %w", err)
		}
		r.ReadMaxID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (r *MessagesReadDiscussionRequest) GetPeer() (value InputPeerClass) {
	if r == nil {
		return
	}
	return r.Peer
}

// GetMsgID returns value of MsgID field.
func (r *MessagesReadDiscussionRequest) GetMsgID() (value int) {
	if r == nil {
		return
	}
	return r.MsgID
}

// GetReadMaxID returns value of ReadMaxID field.
func (r *MessagesReadDiscussionRequest) GetReadMaxID() (value int) {
	if r == nil {
		return
	}
	return r.ReadMaxID
}

// MessagesReadDiscussion invokes method messages.readDiscussion#f731a9f4 returning error if any.
// Mark a thread¹ as read
//
// Links:
//  1. https://core.telegram.org/api/threads
//
// Possible errors:
//
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.readDiscussion for reference.
func (c *Client) MessagesReadDiscussion(ctx context.Context, request *MessagesReadDiscussionRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
