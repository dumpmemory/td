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
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// PhoneGetGroupParticipantsRequest represents TL type `phone.getGroupParticipants#c558d8ab`.
//
// See https://core.telegram.org/method/phone.getGroupParticipants for reference.
type PhoneGetGroupParticipantsRequest struct {
	// Call field of PhoneGetGroupParticipantsRequest.
	Call InputGroupCall
	// IDs field of PhoneGetGroupParticipantsRequest.
	IDs []InputPeerClass
	// Sources field of PhoneGetGroupParticipantsRequest.
	Sources []int
	// Offset field of PhoneGetGroupParticipantsRequest.
	Offset string
	// Limit field of PhoneGetGroupParticipantsRequest.
	Limit int
}

// PhoneGetGroupParticipantsRequestTypeID is TL type id of PhoneGetGroupParticipantsRequest.
const PhoneGetGroupParticipantsRequestTypeID = 0xc558d8ab

func (g *PhoneGetGroupParticipantsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Call.Zero()) {
		return false
	}
	if !(g.IDs == nil) {
		return false
	}
	if !(g.Sources == nil) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGetGroupParticipantsRequest) String() string {
	if g == nil {
		return "PhoneGetGroupParticipantsRequest(nil)"
	}
	type Alias PhoneGetGroupParticipantsRequest
	return fmt.Sprintf("PhoneGetGroupParticipantsRequest%+v", Alias(*g))
}

// FillFrom fills PhoneGetGroupParticipantsRequest from given interface.
func (g *PhoneGetGroupParticipantsRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
	GetIDs() (value []InputPeerClass)
	GetSources() (value []int)
	GetOffset() (value string)
	GetLimit() (value int)
}) {
	g.Call = from.GetCall()
	g.IDs = from.GetIDs()
	g.Sources = from.GetSources()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneGetGroupParticipantsRequest) TypeID() uint32 {
	return PhoneGetGroupParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneGetGroupParticipantsRequest) TypeName() string {
	return "phone.getGroupParticipants"
}

// TypeInfo returns info about TL type.
func (g *PhoneGetGroupParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.getGroupParticipants",
		ID:   PhoneGetGroupParticipantsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "IDs",
			SchemaName: "ids",
		},
		{
			Name:       "Sources",
			SchemaName: "sources",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *PhoneGetGroupParticipantsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupParticipants#c558d8ab as nil")
	}
	b.PutID(PhoneGetGroupParticipantsRequestTypeID)
	if err := g.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.getGroupParticipants#c558d8ab: field call: %w", err)
	}
	b.PutVectorHeader(len(g.IDs))
	for idx, v := range g.IDs {
		if v == nil {
			return fmt.Errorf("unable to encode phone.getGroupParticipants#c558d8ab: field ids element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.getGroupParticipants#c558d8ab: field ids element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(g.Sources))
	for _, v := range g.Sources {
		b.PutInt(v)
	}
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// GetCall returns value of Call field.
func (g *PhoneGetGroupParticipantsRequest) GetCall() (value InputGroupCall) {
	return g.Call
}

// GetIDs returns value of IDs field.
func (g *PhoneGetGroupParticipantsRequest) GetIDs() (value []InputPeerClass) {
	return g.IDs
}

// MapIDs returns field IDs wrapped in InputPeerClassArray helper.
func (g *PhoneGetGroupParticipantsRequest) MapIDs() (value InputPeerClassArray) {
	return InputPeerClassArray(g.IDs)
}

// GetSources returns value of Sources field.
func (g *PhoneGetGroupParticipantsRequest) GetSources() (value []int) {
	return g.Sources
}

// GetOffset returns value of Offset field.
func (g *PhoneGetGroupParticipantsRequest) GetOffset() (value string) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *PhoneGetGroupParticipantsRequest) GetLimit() (value int) {
	return g.Limit
}

// Decode implements bin.Decoder.
func (g *PhoneGetGroupParticipantsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupParticipants#c558d8ab to nil")
	}
	if err := b.ConsumeID(PhoneGetGroupParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.getGroupParticipants#c558d8ab: %w", err)
	}
	{
		if err := g.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c558d8ab: field call: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c558d8ab: field ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.getGroupParticipants#c558d8ab: field ids: %w", err)
			}
			g.IDs = append(g.IDs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c558d8ab: field sources: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode phone.getGroupParticipants#c558d8ab: field sources: %w", err)
			}
			g.Sources = append(g.Sources, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c558d8ab: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c558d8ab: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneGetGroupParticipantsRequest.
var (
	_ bin.Encoder = &PhoneGetGroupParticipantsRequest{}
	_ bin.Decoder = &PhoneGetGroupParticipantsRequest{}
)

// PhoneGetGroupParticipants invokes method phone.getGroupParticipants#c558d8ab returning error if any.
//
// See https://core.telegram.org/method/phone.getGroupParticipants for reference.
func (c *Client) PhoneGetGroupParticipants(ctx context.Context, request *PhoneGetGroupParticipantsRequest) (*PhoneGroupParticipants, error) {
	var result PhoneGroupParticipants

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
