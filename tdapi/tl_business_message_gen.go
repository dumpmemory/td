// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// BusinessMessage represents TL type `businessMessage#fa604646`.
type BusinessMessage struct {
	// The message
	Message Message
	// Message that is replied by the message in the same chat; may be null if none
	ReplyToMessage Message
}

// BusinessMessageTypeID is TL type id of BusinessMessage.
const BusinessMessageTypeID = 0xfa604646

// Ensuring interfaces in compile-time for BusinessMessage.
var (
	_ bin.Encoder     = &BusinessMessage{}
	_ bin.Decoder     = &BusinessMessage{}
	_ bin.BareEncoder = &BusinessMessage{}
	_ bin.BareDecoder = &BusinessMessage{}
)

func (b *BusinessMessage) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Message.Zero()) {
		return false
	}
	if !(b.ReplyToMessage.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessMessage) String() string {
	if b == nil {
		return "BusinessMessage(nil)"
	}
	type Alias BusinessMessage
	return fmt.Sprintf("BusinessMessage%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessMessage) TypeID() uint32 {
	return BusinessMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessMessage) TypeName() string {
	return "businessMessage"
}

// TypeInfo returns info about TL type.
func (b *BusinessMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessMessage",
		ID:   BusinessMessageTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "ReplyToMessage",
			SchemaName: "reply_to_message",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessMessage) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessMessage#fa604646 as nil")
	}
	buf.PutID(BusinessMessageTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessMessage) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessMessage#fa604646 as nil")
	}
	if err := b.Message.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessMessage#fa604646: field message: %w", err)
	}
	if err := b.ReplyToMessage.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessMessage#fa604646: field reply_to_message: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessMessage) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessMessage#fa604646 to nil")
	}
	if err := buf.ConsumeID(BusinessMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode businessMessage#fa604646: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessMessage) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessMessage#fa604646 to nil")
	}
	{
		if err := b.Message.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessMessage#fa604646: field message: %w", err)
		}
	}
	{
		if err := b.ReplyToMessage.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessMessage#fa604646: field reply_to_message: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessMessage) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessMessage#fa604646 as nil")
	}
	buf.ObjStart()
	buf.PutID("businessMessage")
	buf.Comma()
	buf.FieldStart("message")
	if err := b.Message.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessMessage#fa604646: field message: %w", err)
	}
	buf.Comma()
	buf.FieldStart("reply_to_message")
	if err := b.ReplyToMessage.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessMessage#fa604646: field reply_to_message: %w", err)
	}
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessMessage) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessMessage#fa604646 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessMessage"); err != nil {
				return fmt.Errorf("unable to decode businessMessage#fa604646: %w", err)
			}
		case "message":
			if err := b.Message.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessMessage#fa604646: field message: %w", err)
			}
		case "reply_to_message":
			if err := b.ReplyToMessage.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessMessage#fa604646: field reply_to_message: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetMessage returns value of Message field.
func (b *BusinessMessage) GetMessage() (value Message) {
	if b == nil {
		return
	}
	return b.Message
}

// GetReplyToMessage returns value of ReplyToMessage field.
func (b *BusinessMessage) GetReplyToMessage() (value Message) {
	if b == nil {
		return
	}
	return b.ReplyToMessage
}