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

// MessagesSavedGifsNotModified represents TL type `messages.savedGifsNotModified#e8025ca2`.
// No new saved gifs were found
//
// See https://core.telegram.org/constructor/messages.savedGifsNotModified for reference.
type MessagesSavedGifsNotModified struct {
}

// MessagesSavedGifsNotModifiedTypeID is TL type id of MessagesSavedGifsNotModified.
const MessagesSavedGifsNotModifiedTypeID = 0xe8025ca2

// construct implements constructor of MessagesSavedGifsClass.
func (s MessagesSavedGifsNotModified) construct() MessagesSavedGifsClass { return &s }

// Ensuring interfaces in compile-time for MessagesSavedGifsNotModified.
var (
	_ bin.Encoder     = &MessagesSavedGifsNotModified{}
	_ bin.Decoder     = &MessagesSavedGifsNotModified{}
	_ bin.BareEncoder = &MessagesSavedGifsNotModified{}
	_ bin.BareDecoder = &MessagesSavedGifsNotModified{}

	_ MessagesSavedGifsClass = &MessagesSavedGifsNotModified{}
)

func (s *MessagesSavedGifsNotModified) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSavedGifsNotModified) String() string {
	if s == nil {
		return "MessagesSavedGifsNotModified(nil)"
	}
	type Alias MessagesSavedGifsNotModified
	return fmt.Sprintf("MessagesSavedGifsNotModified%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSavedGifsNotModified) TypeID() uint32 {
	return MessagesSavedGifsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSavedGifsNotModified) TypeName() string {
	return "messages.savedGifsNotModified"
}

// TypeInfo returns info about TL type.
func (s *MessagesSavedGifsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.savedGifsNotModified",
		ID:   MessagesSavedGifsNotModifiedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSavedGifsNotModified) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedGifsNotModified#e8025ca2 as nil")
	}
	b.PutID(MessagesSavedGifsNotModifiedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSavedGifsNotModified) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedGifsNotModified#e8025ca2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSavedGifsNotModified) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedGifsNotModified#e8025ca2 to nil")
	}
	if err := b.ConsumeID(MessagesSavedGifsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.savedGifsNotModified#e8025ca2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSavedGifsNotModified) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedGifsNotModified#e8025ca2 to nil")
	}
	return nil
}

// MessagesSavedGifs represents TL type `messages.savedGifs#84a02a0d`.
// Saved gifs
//
// See https://core.telegram.org/constructor/messages.savedGifs for reference.
type MessagesSavedGifs struct {
	// Hash used for caching, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
	// List of saved gifs
	Gifs []DocumentClass
}

// MessagesSavedGifsTypeID is TL type id of MessagesSavedGifs.
const MessagesSavedGifsTypeID = 0x84a02a0d

// construct implements constructor of MessagesSavedGifsClass.
func (s MessagesSavedGifs) construct() MessagesSavedGifsClass { return &s }

// Ensuring interfaces in compile-time for MessagesSavedGifs.
var (
	_ bin.Encoder     = &MessagesSavedGifs{}
	_ bin.Decoder     = &MessagesSavedGifs{}
	_ bin.BareEncoder = &MessagesSavedGifs{}
	_ bin.BareDecoder = &MessagesSavedGifs{}

	_ MessagesSavedGifsClass = &MessagesSavedGifs{}
)

func (s *MessagesSavedGifs) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Hash == 0) {
		return false
	}
	if !(s.Gifs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSavedGifs) String() string {
	if s == nil {
		return "MessagesSavedGifs(nil)"
	}
	type Alias MessagesSavedGifs
	return fmt.Sprintf("MessagesSavedGifs%+v", Alias(*s))
}

// FillFrom fills MessagesSavedGifs from given interface.
func (s *MessagesSavedGifs) FillFrom(from interface {
	GetHash() (value int64)
	GetGifs() (value []DocumentClass)
}) {
	s.Hash = from.GetHash()
	s.Gifs = from.GetGifs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSavedGifs) TypeID() uint32 {
	return MessagesSavedGifsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSavedGifs) TypeName() string {
	return "messages.savedGifs"
}

// TypeInfo returns info about TL type.
func (s *MessagesSavedGifs) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.savedGifs",
		ID:   MessagesSavedGifsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Gifs",
			SchemaName: "gifs",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSavedGifs) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedGifs#84a02a0d as nil")
	}
	b.PutID(MessagesSavedGifsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSavedGifs) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.savedGifs#84a02a0d as nil")
	}
	b.PutLong(s.Hash)
	b.PutVectorHeader(len(s.Gifs))
	for idx, v := range s.Gifs {
		if v == nil {
			return fmt.Errorf("unable to encode messages.savedGifs#84a02a0d: field gifs element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.savedGifs#84a02a0d: field gifs element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSavedGifs) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedGifs#84a02a0d to nil")
	}
	if err := b.ConsumeID(MessagesSavedGifsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.savedGifs#84a02a0d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSavedGifs) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.savedGifs#84a02a0d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedGifs#84a02a0d: field hash: %w", err)
		}
		s.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.savedGifs#84a02a0d: field gifs: %w", err)
		}

		if headerLen > 0 {
			s.Gifs = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.savedGifs#84a02a0d: field gifs: %w", err)
			}
			s.Gifs = append(s.Gifs, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (s *MessagesSavedGifs) GetHash() (value int64) {
	if s == nil {
		return
	}
	return s.Hash
}

// GetGifs returns value of Gifs field.
func (s *MessagesSavedGifs) GetGifs() (value []DocumentClass) {
	if s == nil {
		return
	}
	return s.Gifs
}

// MapGifs returns field Gifs wrapped in DocumentClassArray helper.
func (s *MessagesSavedGifs) MapGifs() (value DocumentClassArray) {
	return DocumentClassArray(s.Gifs)
}

// MessagesSavedGifsClassName is schema name of MessagesSavedGifsClass.
const MessagesSavedGifsClassName = "messages.SavedGifs"

// MessagesSavedGifsClass represents messages.SavedGifs generic type.
//
// See https://core.telegram.org/type/messages.SavedGifs for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagesSavedGifs(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesSavedGifsNotModified: // messages.savedGifsNotModified#e8025ca2
//	case *tg.MessagesSavedGifs: // messages.savedGifs#84a02a0d
//	default: panic(v)
//	}
type MessagesSavedGifsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesSavedGifsClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsModified tries to map MessagesSavedGifsClass to MessagesSavedGifs.
	AsModified() (*MessagesSavedGifs, bool)
}

// AsModified tries to map MessagesSavedGifsNotModified to MessagesSavedGifs.
func (s *MessagesSavedGifsNotModified) AsModified() (*MessagesSavedGifs, bool) {
	return nil, false
}

// AsModified tries to map MessagesSavedGifs to MessagesSavedGifs.
func (s *MessagesSavedGifs) AsModified() (*MessagesSavedGifs, bool) {
	return s, true
}

// DecodeMessagesSavedGifs implements binary de-serialization for MessagesSavedGifsClass.
func DecodeMessagesSavedGifs(buf *bin.Buffer) (MessagesSavedGifsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesSavedGifsNotModifiedTypeID:
		// Decoding messages.savedGifsNotModified#e8025ca2.
		v := MessagesSavedGifsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSavedGifsClass: %w", err)
		}
		return &v, nil
	case MessagesSavedGifsTypeID:
		// Decoding messages.savedGifs#84a02a0d.
		v := MessagesSavedGifs{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesSavedGifsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesSavedGifsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesSavedGifs boxes the MessagesSavedGifsClass providing a helper.
type MessagesSavedGifsBox struct {
	SavedGifs MessagesSavedGifsClass
}

// Decode implements bin.Decoder for MessagesSavedGifsBox.
func (b *MessagesSavedGifsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesSavedGifsBox to nil")
	}
	v, err := DecodeMessagesSavedGifs(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SavedGifs = v
	return nil
}

// Encode implements bin.Encode for MessagesSavedGifsBox.
func (b *MessagesSavedGifsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SavedGifs == nil {
		return fmt.Errorf("unable to encode MessagesSavedGifsClass as nil")
	}
	return b.SavedGifs.Encode(buf)
}
