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

// AccountGetPasswordSettingsRequest represents TL type `account.getPasswordSettings#9cd4eaf9`.
// Get private info associated to the password info (recovery email, telegram passport¹ info & so on)
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.getPasswordSettings for reference.
type AccountGetPasswordSettingsRequest struct {
	// The password (see SRP¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Password InputCheckPasswordSRPClass
}

// AccountGetPasswordSettingsRequestTypeID is TL type id of AccountGetPasswordSettingsRequest.
const AccountGetPasswordSettingsRequestTypeID = 0x9cd4eaf9

func (g *AccountGetPasswordSettingsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Password == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetPasswordSettingsRequest) String() string {
	if g == nil {
		return "AccountGetPasswordSettingsRequest(nil)"
	}
	type Alias AccountGetPasswordSettingsRequest
	return fmt.Sprintf("AccountGetPasswordSettingsRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetPasswordSettingsRequest from given interface.
func (g *AccountGetPasswordSettingsRequest) FillFrom(from interface {
	GetPassword() (value InputCheckPasswordSRPClass)
}) {
	g.Password = from.GetPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetPasswordSettingsRequest) TypeID() uint32 {
	return AccountGetPasswordSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetPasswordSettingsRequest) TypeName() string {
	return "account.getPasswordSettings"
}

// TypeInfo returns info about TL type.
func (g *AccountGetPasswordSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getPasswordSettings",
		ID:   AccountGetPasswordSettingsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetPasswordSettingsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPasswordSettings#9cd4eaf9 as nil")
	}
	b.PutID(AccountGetPasswordSettingsRequestTypeID)
	if g.Password == nil {
		return fmt.Errorf("unable to encode account.getPasswordSettings#9cd4eaf9: field password is nil")
	}
	if err := g.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getPasswordSettings#9cd4eaf9: field password: %w", err)
	}
	return nil
}

// GetPassword returns value of Password field.
func (g *AccountGetPasswordSettingsRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	return g.Password
}

// GetPasswordAsNotEmpty returns mapped value of Password field.
func (g *AccountGetPasswordSettingsRequest) GetPasswordAsNotEmpty() (*InputCheckPasswordSRP, bool) {
	return g.Password.AsNotEmpty()
}

// Decode implements bin.Decoder.
func (g *AccountGetPasswordSettingsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPasswordSettings#9cd4eaf9 to nil")
	}
	if err := b.ConsumeID(AccountGetPasswordSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getPasswordSettings#9cd4eaf9: %w", err)
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getPasswordSettings#9cd4eaf9: field password: %w", err)
		}
		g.Password = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetPasswordSettingsRequest.
var (
	_ bin.Encoder = &AccountGetPasswordSettingsRequest{}
	_ bin.Decoder = &AccountGetPasswordSettingsRequest{}
)

// AccountGetPasswordSettings invokes method account.getPasswordSettings#9cd4eaf9 returning error if any.
// Get private info associated to the password info (recovery email, telegram passport¹ info & so on)
//
// Links:
//  1) https://core.telegram.org/passport
//
// Possible errors:
//  400 PASSWORD_HASH_INVALID: The provided password hash is invalid
//
// See https://core.telegram.org/method/account.getPasswordSettings for reference.
func (c *Client) AccountGetPasswordSettings(ctx context.Context, password InputCheckPasswordSRPClass) (*AccountPasswordSettings, error) {
	var result AccountPasswordSettings

	request := &AccountGetPasswordSettingsRequest{
		Password: password,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
