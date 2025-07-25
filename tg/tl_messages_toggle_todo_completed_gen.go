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

// MessagesToggleTodoCompletedRequest represents TL type `messages.toggleTodoCompleted#d3e03124`.
//
// See https://core.telegram.org/method/messages.toggleTodoCompleted for reference.
type MessagesToggleTodoCompletedRequest struct {
	// Peer field of MessagesToggleTodoCompletedRequest.
	Peer InputPeerClass
	// MsgID field of MessagesToggleTodoCompletedRequest.
	MsgID int
	// Completed field of MessagesToggleTodoCompletedRequest.
	Completed []int
	// Incompleted field of MessagesToggleTodoCompletedRequest.
	Incompleted []int
}

// MessagesToggleTodoCompletedRequestTypeID is TL type id of MessagesToggleTodoCompletedRequest.
const MessagesToggleTodoCompletedRequestTypeID = 0xd3e03124

// Ensuring interfaces in compile-time for MessagesToggleTodoCompletedRequest.
var (
	_ bin.Encoder     = &MessagesToggleTodoCompletedRequest{}
	_ bin.Decoder     = &MessagesToggleTodoCompletedRequest{}
	_ bin.BareEncoder = &MessagesToggleTodoCompletedRequest{}
	_ bin.BareDecoder = &MessagesToggleTodoCompletedRequest{}
)

func (t *MessagesToggleTodoCompletedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Peer == nil) {
		return false
	}
	if !(t.MsgID == 0) {
		return false
	}
	if !(t.Completed == nil) {
		return false
	}
	if !(t.Incompleted == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *MessagesToggleTodoCompletedRequest) String() string {
	if t == nil {
		return "MessagesToggleTodoCompletedRequest(nil)"
	}
	type Alias MessagesToggleTodoCompletedRequest
	return fmt.Sprintf("MessagesToggleTodoCompletedRequest%+v", Alias(*t))
}

// FillFrom fills MessagesToggleTodoCompletedRequest from given interface.
func (t *MessagesToggleTodoCompletedRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetCompleted() (value []int)
	GetIncompleted() (value []int)
}) {
	t.Peer = from.GetPeer()
	t.MsgID = from.GetMsgID()
	t.Completed = from.GetCompleted()
	t.Incompleted = from.GetIncompleted()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesToggleTodoCompletedRequest) TypeID() uint32 {
	return MessagesToggleTodoCompletedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesToggleTodoCompletedRequest) TypeName() string {
	return "messages.toggleTodoCompleted"
}

// TypeInfo returns info about TL type.
func (t *MessagesToggleTodoCompletedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.toggleTodoCompleted",
		ID:   MessagesToggleTodoCompletedRequestTypeID,
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
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Completed",
			SchemaName: "completed",
		},
		{
			Name:       "Incompleted",
			SchemaName: "incompleted",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *MessagesToggleTodoCompletedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleTodoCompleted#d3e03124 as nil")
	}
	b.PutID(MessagesToggleTodoCompletedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *MessagesToggleTodoCompletedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode messages.toggleTodoCompleted#d3e03124 as nil")
	}
	if t.Peer == nil {
		return fmt.Errorf("unable to encode messages.toggleTodoCompleted#d3e03124: field peer is nil")
	}
	if err := t.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.toggleTodoCompleted#d3e03124: field peer: %w", err)
	}
	b.PutInt(t.MsgID)
	b.PutVectorHeader(len(t.Completed))
	for _, v := range t.Completed {
		b.PutInt(v)
	}
	b.PutVectorHeader(len(t.Incompleted))
	for _, v := range t.Incompleted {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *MessagesToggleTodoCompletedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleTodoCompleted#d3e03124 to nil")
	}
	if err := b.ConsumeID(MessagesToggleTodoCompletedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.toggleTodoCompleted#d3e03124: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *MessagesToggleTodoCompletedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode messages.toggleTodoCompleted#d3e03124 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleTodoCompleted#d3e03124: field peer: %w", err)
		}
		t.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleTodoCompleted#d3e03124: field msg_id: %w", err)
		}
		t.MsgID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleTodoCompleted#d3e03124: field completed: %w", err)
		}

		if headerLen > 0 {
			t.Completed = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.toggleTodoCompleted#d3e03124: field completed: %w", err)
			}
			t.Completed = append(t.Completed, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.toggleTodoCompleted#d3e03124: field incompleted: %w", err)
		}

		if headerLen > 0 {
			t.Incompleted = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.toggleTodoCompleted#d3e03124: field incompleted: %w", err)
			}
			t.Incompleted = append(t.Incompleted, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (t *MessagesToggleTodoCompletedRequest) GetPeer() (value InputPeerClass) {
	if t == nil {
		return
	}
	return t.Peer
}

// GetMsgID returns value of MsgID field.
func (t *MessagesToggleTodoCompletedRequest) GetMsgID() (value int) {
	if t == nil {
		return
	}
	return t.MsgID
}

// GetCompleted returns value of Completed field.
func (t *MessagesToggleTodoCompletedRequest) GetCompleted() (value []int) {
	if t == nil {
		return
	}
	return t.Completed
}

// GetIncompleted returns value of Incompleted field.
func (t *MessagesToggleTodoCompletedRequest) GetIncompleted() (value []int) {
	if t == nil {
		return
	}
	return t.Incompleted
}

// MessagesToggleTodoCompleted invokes method messages.toggleTodoCompleted#d3e03124 returning error if any.
//
// See https://core.telegram.org/method/messages.toggleTodoCompleted for reference.
// Can be used by bots.
func (c *Client) MessagesToggleTodoCompleted(ctx context.Context, request *MessagesToggleTodoCompletedRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
