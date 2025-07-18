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

// StarRefProgram represents TL type `starRefProgram#dd0c66f2`.
// Indo about an affiliate program offered by a bot¹
//
// Links:
//  1. https://core.telegram.org/api/bots/referrals
//
// See https://core.telegram.org/constructor/starRefProgram for reference.
type StarRefProgram struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// ID of the bot that offers the program
	BotID int64
	// An affiliate gets a commission of starRefProgram¹.commission_permille‰ Telegram
	// Stars² for every mini app transaction made by users they refer
	//
	// Links:
	//  1) https://core.telegram.org/constructor/starRefProgram
	//  2) https://core.telegram.org/api/stars
	CommissionPermille int
	// An affiliate gets a commission for every mini app transaction made by users they refer
	// for duration_months months after a referral link is imported, starting the bot for
	// the first time
	//
	// Use SetDurationMonths and GetDurationMonths helpers.
	DurationMonths int
	// Point in time (Unix timestamp) when the affiliate program will be closed (optional, if
	// not set the affiliate program isn't scheduled to be closed)
	//
	// Use SetEndDate and GetEndDate helpers.
	EndDate int
	// The amount of daily revenue per user in Telegram Stars of the bot that created the
	// affiliate program. To obtain the approximated revenue per referred user, multiply this
	// value by commission_permille and divide by 1000.
	//
	// Use SetDailyRevenuePerUser and GetDailyRevenuePerUser helpers.
	DailyRevenuePerUser StarsAmountClass
}

// StarRefProgramTypeID is TL type id of StarRefProgram.
const StarRefProgramTypeID = 0xdd0c66f2

// Ensuring interfaces in compile-time for StarRefProgram.
var (
	_ bin.Encoder     = &StarRefProgram{}
	_ bin.Decoder     = &StarRefProgram{}
	_ bin.BareEncoder = &StarRefProgram{}
	_ bin.BareDecoder = &StarRefProgram{}
)

func (s *StarRefProgram) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.BotID == 0) {
		return false
	}
	if !(s.CommissionPermille == 0) {
		return false
	}
	if !(s.DurationMonths == 0) {
		return false
	}
	if !(s.EndDate == 0) {
		return false
	}
	if !(s.DailyRevenuePerUser == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StarRefProgram) String() string {
	if s == nil {
		return "StarRefProgram(nil)"
	}
	type Alias StarRefProgram
	return fmt.Sprintf("StarRefProgram%+v", Alias(*s))
}

// FillFrom fills StarRefProgram from given interface.
func (s *StarRefProgram) FillFrom(from interface {
	GetBotID() (value int64)
	GetCommissionPermille() (value int)
	GetDurationMonths() (value int, ok bool)
	GetEndDate() (value int, ok bool)
	GetDailyRevenuePerUser() (value StarsAmountClass, ok bool)
}) {
	s.BotID = from.GetBotID()
	s.CommissionPermille = from.GetCommissionPermille()
	if val, ok := from.GetDurationMonths(); ok {
		s.DurationMonths = val
	}

	if val, ok := from.GetEndDate(); ok {
		s.EndDate = val
	}

	if val, ok := from.GetDailyRevenuePerUser(); ok {
		s.DailyRevenuePerUser = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StarRefProgram) TypeID() uint32 {
	return StarRefProgramTypeID
}

// TypeName returns name of type in TL schema.
func (*StarRefProgram) TypeName() string {
	return "starRefProgram"
}

// TypeInfo returns info about TL type.
func (s *StarRefProgram) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "starRefProgram",
		ID:   StarRefProgramTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotID",
			SchemaName: "bot_id",
		},
		{
			Name:       "CommissionPermille",
			SchemaName: "commission_permille",
		},
		{
			Name:       "DurationMonths",
			SchemaName: "duration_months",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "EndDate",
			SchemaName: "end_date",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "DailyRevenuePerUser",
			SchemaName: "daily_revenue_per_user",
			Null:       !s.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StarRefProgram) SetFlags() {
	if !(s.DurationMonths == 0) {
		s.Flags.Set(0)
	}
	if !(s.EndDate == 0) {
		s.Flags.Set(1)
	}
	if !(s.DailyRevenuePerUser == nil) {
		s.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (s *StarRefProgram) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starRefProgram#dd0c66f2 as nil")
	}
	b.PutID(StarRefProgramTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StarRefProgram) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode starRefProgram#dd0c66f2 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode starRefProgram#dd0c66f2: field flags: %w", err)
	}
	b.PutLong(s.BotID)
	b.PutInt(s.CommissionPermille)
	if s.Flags.Has(0) {
		b.PutInt(s.DurationMonths)
	}
	if s.Flags.Has(1) {
		b.PutInt(s.EndDate)
	}
	if s.Flags.Has(2) {
		if s.DailyRevenuePerUser == nil {
			return fmt.Errorf("unable to encode starRefProgram#dd0c66f2: field daily_revenue_per_user is nil")
		}
		if err := s.DailyRevenuePerUser.Encode(b); err != nil {
			return fmt.Errorf("unable to encode starRefProgram#dd0c66f2: field daily_revenue_per_user: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StarRefProgram) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starRefProgram#dd0c66f2 to nil")
	}
	if err := b.ConsumeID(StarRefProgramTypeID); err != nil {
		return fmt.Errorf("unable to decode starRefProgram#dd0c66f2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StarRefProgram) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode starRefProgram#dd0c66f2 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode starRefProgram#dd0c66f2: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode starRefProgram#dd0c66f2: field bot_id: %w", err)
		}
		s.BotID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starRefProgram#dd0c66f2: field commission_permille: %w", err)
		}
		s.CommissionPermille = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starRefProgram#dd0c66f2: field duration_months: %w", err)
		}
		s.DurationMonths = value
	}
	if s.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode starRefProgram#dd0c66f2: field end_date: %w", err)
		}
		s.EndDate = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeStarsAmount(b)
		if err != nil {
			return fmt.Errorf("unable to decode starRefProgram#dd0c66f2: field daily_revenue_per_user: %w", err)
		}
		s.DailyRevenuePerUser = value
	}
	return nil
}

// GetBotID returns value of BotID field.
func (s *StarRefProgram) GetBotID() (value int64) {
	if s == nil {
		return
	}
	return s.BotID
}

// GetCommissionPermille returns value of CommissionPermille field.
func (s *StarRefProgram) GetCommissionPermille() (value int) {
	if s == nil {
		return
	}
	return s.CommissionPermille
}

// SetDurationMonths sets value of DurationMonths conditional field.
func (s *StarRefProgram) SetDurationMonths(value int) {
	s.Flags.Set(0)
	s.DurationMonths = value
}

// GetDurationMonths returns value of DurationMonths conditional field and
// boolean which is true if field was set.
func (s *StarRefProgram) GetDurationMonths() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.DurationMonths, true
}

// SetEndDate sets value of EndDate conditional field.
func (s *StarRefProgram) SetEndDate(value int) {
	s.Flags.Set(1)
	s.EndDate = value
}

// GetEndDate returns value of EndDate conditional field and
// boolean which is true if field was set.
func (s *StarRefProgram) GetEndDate() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.EndDate, true
}

// SetDailyRevenuePerUser sets value of DailyRevenuePerUser conditional field.
func (s *StarRefProgram) SetDailyRevenuePerUser(value StarsAmountClass) {
	s.Flags.Set(2)
	s.DailyRevenuePerUser = value
}

// GetDailyRevenuePerUser returns value of DailyRevenuePerUser conditional field and
// boolean which is true if field was set.
func (s *StarRefProgram) GetDailyRevenuePerUser() (value StarsAmountClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.DailyRevenuePerUser, true
}
