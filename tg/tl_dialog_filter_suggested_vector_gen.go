// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// DialogFilterSuggestedVector is a box for Vector<DialogFilterSuggested>
type DialogFilterSuggestedVector struct {
	// Elements of Vector<DialogFilterSuggested>
	Elems []DialogFilterSuggested
}

// Encode implements bin.Encoder.
func (vec *DialogFilterSuggestedVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<DialogFilterSuggested> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<DialogFilterSuggested>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *DialogFilterSuggestedVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<DialogFilterSuggested> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<DialogFilterSuggested>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value DialogFilterSuggested
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<DialogFilterSuggested>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for DialogFilterSuggestedVector.
var (
	_ bin.Encoder = &DialogFilterSuggestedVector{}
	_ bin.Decoder = &DialogFilterSuggestedVector{}
)