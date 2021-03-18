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

// ChatAdminWithInvites represents TL type `chatAdminWithInvites#dfd2330f`.
//
// See https://core.telegram.org/constructor/chatAdminWithInvites for reference.
type ChatAdminWithInvites struct {
	// AdminID field of ChatAdminWithInvites.
	AdminID int
	// InvitesCount field of ChatAdminWithInvites.
	InvitesCount int
	// RevokedInvitesCount field of ChatAdminWithInvites.
	RevokedInvitesCount int
}

// ChatAdminWithInvitesTypeID is TL type id of ChatAdminWithInvites.
const ChatAdminWithInvitesTypeID = 0xdfd2330f

func (c *ChatAdminWithInvites) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.AdminID == 0) {
		return false
	}
	if !(c.InvitesCount == 0) {
		return false
	}
	if !(c.RevokedInvitesCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAdminWithInvites) String() string {
	if c == nil {
		return "ChatAdminWithInvites(nil)"
	}
	type Alias ChatAdminWithInvites
	return fmt.Sprintf("ChatAdminWithInvites%+v", Alias(*c))
}

// FillFrom fills ChatAdminWithInvites from given interface.
func (c *ChatAdminWithInvites) FillFrom(from interface {
	GetAdminID() (value int)
	GetInvitesCount() (value int)
	GetRevokedInvitesCount() (value int)
}) {
	c.AdminID = from.GetAdminID()
	c.InvitesCount = from.GetInvitesCount()
	c.RevokedInvitesCount = from.GetRevokedInvitesCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatAdminWithInvites) TypeID() uint32 {
	return ChatAdminWithInvitesTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatAdminWithInvites) TypeName() string {
	return "chatAdminWithInvites"
}

// TypeInfo returns info about TL type.
func (c *ChatAdminWithInvites) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatAdminWithInvites",
		ID:   ChatAdminWithInvitesTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AdminID",
			SchemaName: "admin_id",
		},
		{
			Name:       "InvitesCount",
			SchemaName: "invites_count",
		},
		{
			Name:       "RevokedInvitesCount",
			SchemaName: "revoked_invites_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatAdminWithInvites) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdminWithInvites#dfd2330f as nil")
	}
	b.PutID(ChatAdminWithInvitesTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatAdminWithInvites) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdminWithInvites#dfd2330f as nil")
	}
	b.PutInt(c.AdminID)
	b.PutInt(c.InvitesCount)
	b.PutInt(c.RevokedInvitesCount)
	return nil
}

// GetAdminID returns value of AdminID field.
func (c *ChatAdminWithInvites) GetAdminID() (value int) {
	return c.AdminID
}

// GetInvitesCount returns value of InvitesCount field.
func (c *ChatAdminWithInvites) GetInvitesCount() (value int) {
	return c.InvitesCount
}

// GetRevokedInvitesCount returns value of RevokedInvitesCount field.
func (c *ChatAdminWithInvites) GetRevokedInvitesCount() (value int) {
	return c.RevokedInvitesCount
}

// Decode implements bin.Decoder.
func (c *ChatAdminWithInvites) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdminWithInvites#dfd2330f to nil")
	}
	if err := b.ConsumeID(ChatAdminWithInvitesTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAdminWithInvites#dfd2330f: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatAdminWithInvites) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdminWithInvites#dfd2330f to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdminWithInvites#dfd2330f: field admin_id: %w", err)
		}
		c.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdminWithInvites#dfd2330f: field invites_count: %w", err)
		}
		c.InvitesCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatAdminWithInvites#dfd2330f: field revoked_invites_count: %w", err)
		}
		c.RevokedInvitesCount = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChatAdminWithInvites.
var (
	_ bin.Encoder     = &ChatAdminWithInvites{}
	_ bin.Decoder     = &ChatAdminWithInvites{}
	_ bin.BareEncoder = &ChatAdminWithInvites{}
	_ bin.BareDecoder = &ChatAdminWithInvites{}
)