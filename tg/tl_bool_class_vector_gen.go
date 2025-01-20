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

// BoolClassVector is a box for Vector<Bool>
type BoolClassVector struct {
	// Elements of Vector<Bool>
	Elems []bool
}

// BoolClassVectorTypeID is TL type id of BoolClassVector.
const BoolClassVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for BoolClassVector.
var (
	_ bin.Encoder     = &BoolClassVector{}
	_ bin.Decoder     = &BoolClassVector{}
	_ bin.BareEncoder = &BoolClassVector{}
	_ bin.BareDecoder = &BoolClassVector{}
)

func (vec *BoolClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *BoolClassVector) String() string {
	if vec == nil {
		return "BoolClassVector(nil)"
	}
	type Alias BoolClassVector
	return fmt.Sprintf("BoolClassVector%+v", Alias(*vec))
}

// FillFrom fills BoolClassVector from given interface.
func (vec *BoolClassVector) FillFrom(from interface {
	GetElems() (value []bool)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BoolClassVector) TypeID() uint32 {
	return BoolClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*BoolClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *BoolClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   BoolClassVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *BoolClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<Bool> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *BoolClassVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<Bool> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for _, v := range vec.Elems {
		b.PutBool(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *BoolClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<Bool> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *BoolClassVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<Bool> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<Bool>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]bool, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode Vector<Bool>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *BoolClassVector) GetElems() (value []bool) {
	if vec == nil {
		return
	}
	return vec.Elems
}
