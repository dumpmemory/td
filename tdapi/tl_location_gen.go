// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// Location represents TL type `location#e5925f73`.
type Location struct {
	// Latitude of the location in degrees; as defined by the sender
	Latitude float64
	// Longitude of the location, in degrees; as defined by the sender
	Longitude float64
	// The estimated horizontal accuracy of the location, in meters; as defined by the sender
	// 0 if unknown
	HorizontalAccuracy float64
}

// LocationTypeID is TL type id of Location.
const LocationTypeID = 0xe5925f73

// Ensuring interfaces in compile-time for Location.
var (
	_ bin.Encoder     = &Location{}
	_ bin.Decoder     = &Location{}
	_ bin.BareEncoder = &Location{}
	_ bin.BareDecoder = &Location{}
)

func (l *Location) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Latitude == 0) {
		return false
	}
	if !(l.Longitude == 0) {
		return false
	}
	if !(l.HorizontalAccuracy == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *Location) String() string {
	if l == nil {
		return "Location(nil)"
	}
	type Alias Location
	return fmt.Sprintf("Location%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Location) TypeID() uint32 {
	return LocationTypeID
}

// TypeName returns name of type in TL schema.
func (*Location) TypeName() string {
	return "location"
}

// TypeInfo returns info about TL type.
func (l *Location) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "location",
		ID:   LocationTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Latitude",
			SchemaName: "latitude",
		},
		{
			Name:       "Longitude",
			SchemaName: "longitude",
		},
		{
			Name:       "HorizontalAccuracy",
			SchemaName: "horizontal_accuracy",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *Location) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode location#e5925f73 as nil")
	}
	b.PutID(LocationTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *Location) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode location#e5925f73 as nil")
	}
	b.PutDouble(l.Latitude)
	b.PutDouble(l.Longitude)
	b.PutDouble(l.HorizontalAccuracy)
	return nil
}

// Decode implements bin.Decoder.
func (l *Location) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode location#e5925f73 to nil")
	}
	if err := b.ConsumeID(LocationTypeID); err != nil {
		return fmt.Errorf("unable to decode location#e5925f73: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *Location) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode location#e5925f73 to nil")
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode location#e5925f73: field latitude: %w", err)
		}
		l.Latitude = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode location#e5925f73: field longitude: %w", err)
		}
		l.Longitude = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode location#e5925f73: field horizontal_accuracy: %w", err)
		}
		l.HorizontalAccuracy = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *Location) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode location#e5925f73 as nil")
	}
	b.ObjStart()
	b.PutID("location")
	b.FieldStart("latitude")
	b.PutDouble(l.Latitude)
	b.FieldStart("longitude")
	b.PutDouble(l.Longitude)
	b.FieldStart("horizontal_accuracy")
	b.PutDouble(l.HorizontalAccuracy)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *Location) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode location#e5925f73 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("location"); err != nil {
				return fmt.Errorf("unable to decode location#e5925f73: %w", err)
			}
		case "latitude":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode location#e5925f73: field latitude: %w", err)
			}
			l.Latitude = value
		case "longitude":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode location#e5925f73: field longitude: %w", err)
			}
			l.Longitude = value
		case "horizontal_accuracy":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode location#e5925f73: field horizontal_accuracy: %w", err)
			}
			l.HorizontalAccuracy = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLatitude returns value of Latitude field.
func (l *Location) GetLatitude() (value float64) {
	return l.Latitude
}

// GetLongitude returns value of Longitude field.
func (l *Location) GetLongitude() (value float64) {
	return l.Longitude
}

// GetHorizontalAccuracy returns value of HorizontalAccuracy field.
func (l *Location) GetHorizontalAccuracy() (value float64) {
	return l.HorizontalAccuracy
}