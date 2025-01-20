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

// HelpGetPeerColorsRequest represents TL type `help.getPeerColors#da80f42f`.
// Get the set of accent color palettes »¹ that can be used for message accents.
//
// Links:
//  1. https://core.telegram.org/api/colors
//
// See https://core.telegram.org/method/help.getPeerColors for reference.
type HelpGetPeerColorsRequest struct {
	// Hash used for caching, for more info click here¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// HelpGetPeerColorsRequestTypeID is TL type id of HelpGetPeerColorsRequest.
const HelpGetPeerColorsRequestTypeID = 0xda80f42f

// Ensuring interfaces in compile-time for HelpGetPeerColorsRequest.
var (
	_ bin.Encoder     = &HelpGetPeerColorsRequest{}
	_ bin.Decoder     = &HelpGetPeerColorsRequest{}
	_ bin.BareEncoder = &HelpGetPeerColorsRequest{}
	_ bin.BareDecoder = &HelpGetPeerColorsRequest{}
)

func (g *HelpGetPeerColorsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetPeerColorsRequest) String() string {
	if g == nil {
		return "HelpGetPeerColorsRequest(nil)"
	}
	type Alias HelpGetPeerColorsRequest
	return fmt.Sprintf("HelpGetPeerColorsRequest%+v", Alias(*g))
}

// FillFrom fills HelpGetPeerColorsRequest from given interface.
func (g *HelpGetPeerColorsRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetPeerColorsRequest) TypeID() uint32 {
	return HelpGetPeerColorsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetPeerColorsRequest) TypeName() string {
	return "help.getPeerColors"
}

// TypeInfo returns info about TL type.
func (g *HelpGetPeerColorsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getPeerColors",
		ID:   HelpGetPeerColorsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetPeerColorsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getPeerColors#da80f42f as nil")
	}
	b.PutID(HelpGetPeerColorsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetPeerColorsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getPeerColors#da80f42f as nil")
	}
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetPeerColorsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getPeerColors#da80f42f to nil")
	}
	if err := b.ConsumeID(HelpGetPeerColorsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getPeerColors#da80f42f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetPeerColorsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getPeerColors#da80f42f to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.getPeerColors#da80f42f: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *HelpGetPeerColorsRequest) GetHash() (value int) {
	if g == nil {
		return
	}
	return g.Hash
}

// HelpGetPeerColors invokes method help.getPeerColors#da80f42f returning error if any.
// Get the set of accent color palettes »¹ that can be used for message accents.
//
// Links:
//  1. https://core.telegram.org/api/colors
//
// See https://core.telegram.org/method/help.getPeerColors for reference.
func (c *Client) HelpGetPeerColors(ctx context.Context, hash int) (HelpPeerColorsClass, error) {
	var result HelpPeerColorsBox

	request := &HelpGetPeerColorsRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PeerColors, nil
}
