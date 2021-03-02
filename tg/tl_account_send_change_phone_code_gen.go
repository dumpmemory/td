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

// AccountSendChangePhoneCodeRequest represents TL type `account.sendChangePhoneCode#82574ae5`.
// Verify a new phone number to associate to the current account
//
// See https://core.telegram.org/method/account.sendChangePhoneCode for reference.
type AccountSendChangePhoneCodeRequest struct {
	// New phone number
	PhoneNumber string
	// Phone code settings
	Settings CodeSettings
}

// AccountSendChangePhoneCodeRequestTypeID is TL type id of AccountSendChangePhoneCodeRequest.
const AccountSendChangePhoneCodeRequestTypeID = 0x82574ae5

func (s *AccountSendChangePhoneCodeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSendChangePhoneCodeRequest) String() string {
	if s == nil {
		return "AccountSendChangePhoneCodeRequest(nil)"
	}
	type Alias AccountSendChangePhoneCodeRequest
	return fmt.Sprintf("AccountSendChangePhoneCodeRequest%+v", Alias(*s))
}

// FillFrom fills AccountSendChangePhoneCodeRequest from given interface.
func (s *AccountSendChangePhoneCodeRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetSettings() (value CodeSettings)
}) {
	s.PhoneNumber = from.GetPhoneNumber()
	s.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSendChangePhoneCodeRequest) TypeID() uint32 {
	return AccountSendChangePhoneCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSendChangePhoneCodeRequest) TypeName() string {
	return "account.sendChangePhoneCode"
}

// TypeInfo returns info about TL type.
func (s *AccountSendChangePhoneCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.sendChangePhoneCode",
		ID:   AccountSendChangePhoneCodeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSendChangePhoneCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendChangePhoneCode#82574ae5 as nil")
	}
	b.PutID(AccountSendChangePhoneCodeRequestTypeID)
	b.PutString(s.PhoneNumber)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.sendChangePhoneCode#82574ae5: field settings: %w", err)
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *AccountSendChangePhoneCodeRequest) GetPhoneNumber() (value string) {
	return s.PhoneNumber
}

// GetSettings returns value of Settings field.
func (s *AccountSendChangePhoneCodeRequest) GetSettings() (value CodeSettings) {
	return s.Settings
}

// Decode implements bin.Decoder.
func (s *AccountSendChangePhoneCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendChangePhoneCode#82574ae5 to nil")
	}
	if err := b.ConsumeID(AccountSendChangePhoneCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendChangePhoneCode#82574ae5: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendChangePhoneCode#82574ae5: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.sendChangePhoneCode#82574ae5: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSendChangePhoneCodeRequest.
var (
	_ bin.Encoder = &AccountSendChangePhoneCodeRequest{}
	_ bin.Decoder = &AccountSendChangePhoneCodeRequest{}
)

// AccountSendChangePhoneCode invokes method account.sendChangePhoneCode#82574ae5 returning error if any.
// Verify a new phone number to associate to the current account
//
// Possible errors:
//  400 PHONE_NUMBER_INVALID: The phone number is invalid
//
// See https://core.telegram.org/method/account.sendChangePhoneCode for reference.
func (c *Client) AccountSendChangePhoneCode(ctx context.Context, request *AccountSendChangePhoneCodeRequest) (*AuthSentCode, error) {
	var result AuthSentCode

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
