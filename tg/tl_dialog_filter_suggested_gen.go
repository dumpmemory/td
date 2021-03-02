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
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// DialogFilterSuggested represents TL type `dialogFilterSuggested#77744d4a`.
// Suggested folders¹
//
// Links:
//  1) https://core.telegram.org/api/folders
//
// See https://core.telegram.org/constructor/dialogFilterSuggested for reference.
type DialogFilterSuggested struct {
	// Folder info¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Filter DialogFilter
	// Folder¹ description
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Description string
}

// DialogFilterSuggestedTypeID is TL type id of DialogFilterSuggested.
const DialogFilterSuggestedTypeID = 0x77744d4a

func (d *DialogFilterSuggested) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Filter.Zero()) {
		return false
	}
	if !(d.Description == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DialogFilterSuggested) String() string {
	if d == nil {
		return "DialogFilterSuggested(nil)"
	}
	type Alias DialogFilterSuggested
	return fmt.Sprintf("DialogFilterSuggested%+v", Alias(*d))
}

// FillFrom fills DialogFilterSuggested from given interface.
func (d *DialogFilterSuggested) FillFrom(from interface {
	GetFilter() (value DialogFilter)
	GetDescription() (value string)
}) {
	d.Filter = from.GetFilter()
	d.Description = from.GetDescription()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DialogFilterSuggested) TypeID() uint32 {
	return DialogFilterSuggestedTypeID
}

// TypeName returns name of type in TL schema.
func (*DialogFilterSuggested) TypeName() string {
	return "dialogFilterSuggested"
}

// TypeInfo returns info about TL type.
func (d *DialogFilterSuggested) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dialogFilterSuggested",
		ID:   DialogFilterSuggestedTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DialogFilterSuggested) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilterSuggested#77744d4a as nil")
	}
	b.PutID(DialogFilterSuggestedTypeID)
	if err := d.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogFilterSuggested#77744d4a: field filter: %w", err)
	}
	b.PutString(d.Description)
	return nil
}

// GetFilter returns value of Filter field.
func (d *DialogFilterSuggested) GetFilter() (value DialogFilter) {
	return d.Filter
}

// GetDescription returns value of Description field.
func (d *DialogFilterSuggested) GetDescription() (value string) {
	return d.Description
}

// Decode implements bin.Decoder.
func (d *DialogFilterSuggested) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilterSuggested#77744d4a to nil")
	}
	if err := b.ConsumeID(DialogFilterSuggestedTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogFilterSuggested#77744d4a: %w", err)
	}
	{
		if err := d.Filter.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialogFilterSuggested#77744d4a: field filter: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterSuggested#77744d4a: field description: %w", err)
		}
		d.Description = value
	}
	return nil
}

// Ensuring interfaces in compile-time for DialogFilterSuggested.
var (
	_ bin.Encoder = &DialogFilterSuggested{}
	_ bin.Decoder = &DialogFilterSuggested{}
)
