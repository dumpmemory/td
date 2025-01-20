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

// SMSJobsUpdateSettingsRequest represents TL type `smsjobs.updateSettings#93fa0bf`.
// Update SMS job settings (official clients only).
//
// See https://core.telegram.org/method/smsjobs.updateSettings for reference.
type SMSJobsUpdateSettingsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Allow international numbers?
	AllowInternational bool
}

// SMSJobsUpdateSettingsRequestTypeID is TL type id of SMSJobsUpdateSettingsRequest.
const SMSJobsUpdateSettingsRequestTypeID = 0x93fa0bf

// Ensuring interfaces in compile-time for SMSJobsUpdateSettingsRequest.
var (
	_ bin.Encoder     = &SMSJobsUpdateSettingsRequest{}
	_ bin.Decoder     = &SMSJobsUpdateSettingsRequest{}
	_ bin.BareEncoder = &SMSJobsUpdateSettingsRequest{}
	_ bin.BareDecoder = &SMSJobsUpdateSettingsRequest{}
)

func (u *SMSJobsUpdateSettingsRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.AllowInternational == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *SMSJobsUpdateSettingsRequest) String() string {
	if u == nil {
		return "SMSJobsUpdateSettingsRequest(nil)"
	}
	type Alias SMSJobsUpdateSettingsRequest
	return fmt.Sprintf("SMSJobsUpdateSettingsRequest%+v", Alias(*u))
}

// FillFrom fills SMSJobsUpdateSettingsRequest from given interface.
func (u *SMSJobsUpdateSettingsRequest) FillFrom(from interface {
	GetAllowInternational() (value bool)
}) {
	u.AllowInternational = from.GetAllowInternational()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SMSJobsUpdateSettingsRequest) TypeID() uint32 {
	return SMSJobsUpdateSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SMSJobsUpdateSettingsRequest) TypeName() string {
	return "smsjobs.updateSettings"
}

// TypeInfo returns info about TL type.
func (u *SMSJobsUpdateSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "smsjobs.updateSettings",
		ID:   SMSJobsUpdateSettingsRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AllowInternational",
			SchemaName: "allow_international",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *SMSJobsUpdateSettingsRequest) SetFlags() {
	if !(u.AllowInternational == false) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *SMSJobsUpdateSettingsRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode smsjobs.updateSettings#93fa0bf as nil")
	}
	b.PutID(SMSJobsUpdateSettingsRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *SMSJobsUpdateSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode smsjobs.updateSettings#93fa0bf as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode smsjobs.updateSettings#93fa0bf: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *SMSJobsUpdateSettingsRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode smsjobs.updateSettings#93fa0bf to nil")
	}
	if err := b.ConsumeID(SMSJobsUpdateSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode smsjobs.updateSettings#93fa0bf: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *SMSJobsUpdateSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode smsjobs.updateSettings#93fa0bf to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode smsjobs.updateSettings#93fa0bf: field flags: %w", err)
		}
	}
	u.AllowInternational = u.Flags.Has(0)
	return nil
}

// SetAllowInternational sets value of AllowInternational conditional field.
func (u *SMSJobsUpdateSettingsRequest) SetAllowInternational(value bool) {
	if value {
		u.Flags.Set(0)
		u.AllowInternational = true
	} else {
		u.Flags.Unset(0)
		u.AllowInternational = false
	}
}

// GetAllowInternational returns value of AllowInternational conditional field.
func (u *SMSJobsUpdateSettingsRequest) GetAllowInternational() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// SMSJobsUpdateSettings invokes method smsjobs.updateSettings#93fa0bf returning error if any.
// Update SMS job settings (official clients only).
//
// Possible errors:
//
//	400 NOT_JOINED: The current user hasn't joined the Peer-to-Peer Login Program.
//
// See https://core.telegram.org/method/smsjobs.updateSettings for reference.
func (c *Client) SMSJobsUpdateSettings(ctx context.Context, request *SMSJobsUpdateSettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
