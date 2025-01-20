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

// ChannelsToggleSlowModeRequest represents TL type `channels.toggleSlowMode#edd49ef0`.
// Toggle supergroup slow mode: if enabled, users will only be able to send one message
// every seconds seconds
//
// See https://core.telegram.org/method/channels.toggleSlowMode for reference.
type ChannelsToggleSlowModeRequest struct {
	// The supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// Users will only be able to send one message every seconds seconds, 0 to disable the
	// limitation
	Seconds int
}

// ChannelsToggleSlowModeRequestTypeID is TL type id of ChannelsToggleSlowModeRequest.
const ChannelsToggleSlowModeRequestTypeID = 0xedd49ef0

// Ensuring interfaces in compile-time for ChannelsToggleSlowModeRequest.
var (
	_ bin.Encoder     = &ChannelsToggleSlowModeRequest{}
	_ bin.Decoder     = &ChannelsToggleSlowModeRequest{}
	_ bin.BareEncoder = &ChannelsToggleSlowModeRequest{}
	_ bin.BareDecoder = &ChannelsToggleSlowModeRequest{}
)

func (t *ChannelsToggleSlowModeRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Seconds == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsToggleSlowModeRequest) String() string {
	if t == nil {
		return "ChannelsToggleSlowModeRequest(nil)"
	}
	type Alias ChannelsToggleSlowModeRequest
	return fmt.Sprintf("ChannelsToggleSlowModeRequest%+v", Alias(*t))
}

// FillFrom fills ChannelsToggleSlowModeRequest from given interface.
func (t *ChannelsToggleSlowModeRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetSeconds() (value int)
}) {
	t.Channel = from.GetChannel()
	t.Seconds = from.GetSeconds()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsToggleSlowModeRequest) TypeID() uint32 {
	return ChannelsToggleSlowModeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsToggleSlowModeRequest) TypeName() string {
	return "channels.toggleSlowMode"
}

// TypeInfo returns info about TL type.
func (t *ChannelsToggleSlowModeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.toggleSlowMode",
		ID:   ChannelsToggleSlowModeRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Seconds",
			SchemaName: "seconds",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ChannelsToggleSlowModeRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleSlowMode#edd49ef0 as nil")
	}
	b.PutID(ChannelsToggleSlowModeRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ChannelsToggleSlowModeRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleSlowMode#edd49ef0 as nil")
	}
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleSlowMode#edd49ef0: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleSlowMode#edd49ef0: field channel: %w", err)
	}
	b.PutInt(t.Seconds)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleSlowModeRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleSlowMode#edd49ef0 to nil")
	}
	if err := b.ConsumeID(ChannelsToggleSlowModeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ChannelsToggleSlowModeRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleSlowMode#edd49ef0 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: field seconds: %w", err)
		}
		t.Seconds = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsToggleSlowModeRequest) GetChannel() (value InputChannelClass) {
	if t == nil {
		return
	}
	return t.Channel
}

// GetSeconds returns value of Seconds field.
func (t *ChannelsToggleSlowModeRequest) GetSeconds() (value int) {
	if t == nil {
		return
	}
	return t.Seconds
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (t *ChannelsToggleSlowModeRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return t.Channel.AsNotEmpty()
}

// ChannelsToggleSlowMode invokes method channels.toggleSlowMode#edd49ef0 returning error if any.
// Toggle supergroup slow mode: if enabled, users will only be able to send one message
// every seconds seconds
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	400 CHAT_NOT_MODIFIED: No changes were made to chat information because the new information you passed is identical to the current information.
//	400 SECONDS_INVALID: Invalid duration provided.
//
// See https://core.telegram.org/method/channels.toggleSlowMode for reference.
func (c *Client) ChannelsToggleSlowMode(ctx context.Context, request *ChannelsToggleSlowModeRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
