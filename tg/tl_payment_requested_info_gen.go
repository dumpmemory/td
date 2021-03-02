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

// PaymentRequestedInfo represents TL type `paymentRequestedInfo#909c3f94`.
// Order info provided by the user
//
// See https://core.telegram.org/constructor/paymentRequestedInfo for reference.
type PaymentRequestedInfo struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// User's full name
	//
	// Use SetName and GetName helpers.
	Name string
	// User's phone number
	//
	// Use SetPhone and GetPhone helpers.
	Phone string
	// User's email address
	//
	// Use SetEmail and GetEmail helpers.
	Email string
	// User's shipping address
	//
	// Use SetShippingAddress and GetShippingAddress helpers.
	ShippingAddress PostAddress
}

// PaymentRequestedInfoTypeID is TL type id of PaymentRequestedInfo.
const PaymentRequestedInfoTypeID = 0x909c3f94

func (p *PaymentRequestedInfo) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Name == "") {
		return false
	}
	if !(p.Phone == "") {
		return false
	}
	if !(p.Email == "") {
		return false
	}
	if !(p.ShippingAddress.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentRequestedInfo) String() string {
	if p == nil {
		return "PaymentRequestedInfo(nil)"
	}
	type Alias PaymentRequestedInfo
	return fmt.Sprintf("PaymentRequestedInfo%+v", Alias(*p))
}

// FillFrom fills PaymentRequestedInfo from given interface.
func (p *PaymentRequestedInfo) FillFrom(from interface {
	GetName() (value string, ok bool)
	GetPhone() (value string, ok bool)
	GetEmail() (value string, ok bool)
	GetShippingAddress() (value PostAddress, ok bool)
}) {
	if val, ok := from.GetName(); ok {
		p.Name = val
	}

	if val, ok := from.GetPhone(); ok {
		p.Phone = val
	}

	if val, ok := from.GetEmail(); ok {
		p.Email = val
	}

	if val, ok := from.GetShippingAddress(); ok {
		p.ShippingAddress = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentRequestedInfo) TypeID() uint32 {
	return PaymentRequestedInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentRequestedInfo) TypeName() string {
	return "paymentRequestedInfo"
}

// TypeInfo returns info about TL type.
func (p *PaymentRequestedInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentRequestedInfo",
		ID:   PaymentRequestedInfoTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "Name",
			SchemaName: "name",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "Phone",
			SchemaName: "phone",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "Email",
			SchemaName: "email",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "ShippingAddress",
			SchemaName: "shipping_address",
			Null:       !p.Flags.Has(3),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentRequestedInfo) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentRequestedInfo#909c3f94 as nil")
	}
	b.PutID(PaymentRequestedInfoTypeID)
	if !(p.Name == "") {
		p.Flags.Set(0)
	}
	if !(p.Phone == "") {
		p.Flags.Set(1)
	}
	if !(p.Email == "") {
		p.Flags.Set(2)
	}
	if !(p.ShippingAddress.Zero()) {
		p.Flags.Set(3)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentRequestedInfo#909c3f94: field flags: %w", err)
	}
	if p.Flags.Has(0) {
		b.PutString(p.Name)
	}
	if p.Flags.Has(1) {
		b.PutString(p.Phone)
	}
	if p.Flags.Has(2) {
		b.PutString(p.Email)
	}
	if p.Flags.Has(3) {
		if err := p.ShippingAddress.Encode(b); err != nil {
			return fmt.Errorf("unable to encode paymentRequestedInfo#909c3f94: field shipping_address: %w", err)
		}
	}
	return nil
}

// SetName sets value of Name conditional field.
func (p *PaymentRequestedInfo) SetName(value string) {
	p.Flags.Set(0)
	p.Name = value
}

// GetName returns value of Name conditional field and
// boolean which is true if field was set.
func (p *PaymentRequestedInfo) GetName() (value string, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Name, true
}

// SetPhone sets value of Phone conditional field.
func (p *PaymentRequestedInfo) SetPhone(value string) {
	p.Flags.Set(1)
	p.Phone = value
}

// GetPhone returns value of Phone conditional field and
// boolean which is true if field was set.
func (p *PaymentRequestedInfo) GetPhone() (value string, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Phone, true
}

// SetEmail sets value of Email conditional field.
func (p *PaymentRequestedInfo) SetEmail(value string) {
	p.Flags.Set(2)
	p.Email = value
}

// GetEmail returns value of Email conditional field and
// boolean which is true if field was set.
func (p *PaymentRequestedInfo) GetEmail() (value string, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.Email, true
}

// SetShippingAddress sets value of ShippingAddress conditional field.
func (p *PaymentRequestedInfo) SetShippingAddress(value PostAddress) {
	p.Flags.Set(3)
	p.ShippingAddress = value
}

// GetShippingAddress returns value of ShippingAddress conditional field and
// boolean which is true if field was set.
func (p *PaymentRequestedInfo) GetShippingAddress() (value PostAddress, ok bool) {
	if !p.Flags.Has(3) {
		return value, false
	}
	return p.ShippingAddress, true
}

// Decode implements bin.Decoder.
func (p *PaymentRequestedInfo) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentRequestedInfo#909c3f94 to nil")
	}
	if err := b.ConsumeID(PaymentRequestedInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentRequestedInfo#909c3f94: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentRequestedInfo#909c3f94: field flags: %w", err)
		}
	}
	if p.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentRequestedInfo#909c3f94: field name: %w", err)
		}
		p.Name = value
	}
	if p.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentRequestedInfo#909c3f94: field phone: %w", err)
		}
		p.Phone = value
	}
	if p.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentRequestedInfo#909c3f94: field email: %w", err)
		}
		p.Email = value
	}
	if p.Flags.Has(3) {
		if err := p.ShippingAddress.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentRequestedInfo#909c3f94: field shipping_address: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentRequestedInfo.
var (
	_ bin.Encoder = &PaymentRequestedInfo{}
	_ bin.Decoder = &PaymentRequestedInfo{}
)
