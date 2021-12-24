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

// ChannelsDeleteParticipantHistoryRequest represents TL type `channels.deleteParticipantHistory#367544db`.
//
// See https://core.telegram.org/method/channels.deleteParticipantHistory for reference.
type ChannelsDeleteParticipantHistoryRequest struct {
	// Channel field of ChannelsDeleteParticipantHistoryRequest.
	Channel InputChannelClass
	// Participant field of ChannelsDeleteParticipantHistoryRequest.
	Participant InputPeerClass
}

// ChannelsDeleteParticipantHistoryRequestTypeID is TL type id of ChannelsDeleteParticipantHistoryRequest.
const ChannelsDeleteParticipantHistoryRequestTypeID = 0x367544db

// Ensuring interfaces in compile-time for ChannelsDeleteParticipantHistoryRequest.
var (
	_ bin.Encoder     = &ChannelsDeleteParticipantHistoryRequest{}
	_ bin.Decoder     = &ChannelsDeleteParticipantHistoryRequest{}
	_ bin.BareEncoder = &ChannelsDeleteParticipantHistoryRequest{}
	_ bin.BareDecoder = &ChannelsDeleteParticipantHistoryRequest{}
)

func (d *ChannelsDeleteParticipantHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}
	if !(d.Participant == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteParticipantHistoryRequest) String() string {
	if d == nil {
		return "ChannelsDeleteParticipantHistoryRequest(nil)"
	}
	type Alias ChannelsDeleteParticipantHistoryRequest
	return fmt.Sprintf("ChannelsDeleteParticipantHistoryRequest%+v", Alias(*d))
}

// FillFrom fills ChannelsDeleteParticipantHistoryRequest from given interface.
func (d *ChannelsDeleteParticipantHistoryRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetParticipant() (value InputPeerClass)
}) {
	d.Channel = from.GetChannel()
	d.Participant = from.GetParticipant()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsDeleteParticipantHistoryRequest) TypeID() uint32 {
	return ChannelsDeleteParticipantHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsDeleteParticipantHistoryRequest) TypeName() string {
	return "channels.deleteParticipantHistory"
}

// TypeInfo returns info about TL type.
func (d *ChannelsDeleteParticipantHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.deleteParticipantHistory",
		ID:   ChannelsDeleteParticipantHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Participant",
			SchemaName: "participant",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteParticipantHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteParticipantHistory#367544db as nil")
	}
	b.PutID(ChannelsDeleteParticipantHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ChannelsDeleteParticipantHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteParticipantHistory#367544db as nil")
	}
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteParticipantHistory#367544db: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteParticipantHistory#367544db: field channel: %w", err)
	}
	if d.Participant == nil {
		return fmt.Errorf("unable to encode channels.deleteParticipantHistory#367544db: field participant is nil")
	}
	if err := d.Participant.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteParticipantHistory#367544db: field participant: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteParticipantHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteParticipantHistory#367544db to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteParticipantHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteParticipantHistory#367544db: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ChannelsDeleteParticipantHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteParticipantHistory#367544db to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteParticipantHistory#367544db: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteParticipantHistory#367544db: field participant: %w", err)
		}
		d.Participant = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteParticipantHistoryRequest) GetChannel() (value InputChannelClass) {
	if d == nil {
		return
	}
	return d.Channel
}

// GetParticipant returns value of Participant field.
func (d *ChannelsDeleteParticipantHistoryRequest) GetParticipant() (value InputPeerClass) {
	if d == nil {
		return
	}
	return d.Participant
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (d *ChannelsDeleteParticipantHistoryRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return d.Channel.AsNotEmpty()
}

// ChannelsDeleteParticipantHistory invokes method channels.deleteParticipantHistory#367544db returning error if any.
//
// See https://core.telegram.org/method/channels.deleteParticipantHistory for reference.
func (c *Client) ChannelsDeleteParticipantHistory(ctx context.Context, request *ChannelsDeleteParticipantHistoryRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}