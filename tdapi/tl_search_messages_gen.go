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

// SearchMessagesRequest represents TL type `searchMessages#31f9c3c9`.
type SearchMessagesRequest struct {
	// Chat list in which to search messages; pass null to search in all chats regardless of
	// their chat list. Only Main and Archive chat lists are supported
	ChatList ChatListClass
	// Pass true to search only for messages in channels
	OnlyInChannels bool
	// Query to search for
	Query string
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of messages to be returned; up to 100. For optimal performance, the
	// number of returned messages is chosen by TDLib and can be smaller than the specified
	// limit
	Limit int32
	// Additional filter for messages to search; pass null to search for all messages.
	// Filters searchMessagesFilterMention, searchMessagesFilterUnreadMention,
	// searchMessagesFilterUnreadReaction, searchMessagesFilterFailedToSend, and
	// searchMessagesFilterPinned are unsupported in this function
	Filter SearchMessagesFilterClass
	// If not 0, the minimum date of the messages to return
	MinDate int32
	// If not 0, the maximum date of the messages to return
	MaxDate int32
}

// SearchMessagesRequestTypeID is TL type id of SearchMessagesRequest.
const SearchMessagesRequestTypeID = 0x31f9c3c9

// Ensuring interfaces in compile-time for SearchMessagesRequest.
var (
	_ bin.Encoder     = &SearchMessagesRequest{}
	_ bin.Decoder     = &SearchMessagesRequest{}
	_ bin.BareEncoder = &SearchMessagesRequest{}
	_ bin.BareDecoder = &SearchMessagesRequest{}
)

func (s *SearchMessagesRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatList == nil) {
		return false
	}
	if !(s.OnlyInChannels == false) {
		return false
	}
	if !(s.Query == "") {
		return false
	}
	if !(s.Offset == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}
	if !(s.Filter == nil) {
		return false
	}
	if !(s.MinDate == 0) {
		return false
	}
	if !(s.MaxDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchMessagesRequest) String() string {
	if s == nil {
		return "SearchMessagesRequest(nil)"
	}
	type Alias SearchMessagesRequest
	return fmt.Sprintf("SearchMessagesRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchMessagesRequest) TypeID() uint32 {
	return SearchMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchMessagesRequest) TypeName() string {
	return "searchMessages"
}

// TypeInfo returns info about TL type.
func (s *SearchMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchMessages",
		ID:   SearchMessagesRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatList",
			SchemaName: "chat_list",
		},
		{
			Name:       "OnlyInChannels",
			SchemaName: "only_in_channels",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "MinDate",
			SchemaName: "min_date",
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchMessagesRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchMessages#31f9c3c9 as nil")
	}
	b.PutID(SearchMessagesRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchMessages#31f9c3c9 as nil")
	}
	if s.ChatList == nil {
		return fmt.Errorf("unable to encode searchMessages#31f9c3c9: field chat_list is nil")
	}
	if err := s.ChatList.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchMessages#31f9c3c9: field chat_list: %w", err)
	}
	b.PutBool(s.OnlyInChannels)
	b.PutString(s.Query)
	b.PutString(s.Offset)
	b.PutInt32(s.Limit)
	if s.Filter == nil {
		return fmt.Errorf("unable to encode searchMessages#31f9c3c9: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchMessages#31f9c3c9: field filter: %w", err)
	}
	b.PutInt32(s.MinDate)
	b.PutInt32(s.MaxDate)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchMessagesRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchMessages#31f9c3c9 to nil")
	}
	if err := b.ConsumeID(SearchMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchMessages#31f9c3c9: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchMessages#31f9c3c9 to nil")
	}
	{
		value, err := DecodeChatList(b)
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field chat_list: %w", err)
		}
		s.ChatList = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field only_in_channels: %w", err)
		}
		s.OnlyInChannels = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field query: %w", err)
		}
		s.Query = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field offset: %w", err)
		}
		s.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field limit: %w", err)
		}
		s.Limit = value
	}
	{
		value, err := DecodeSearchMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field filter: %w", err)
		}
		s.Filter = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field min_date: %w", err)
		}
		s.MinDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field max_date: %w", err)
		}
		s.MaxDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchMessages#31f9c3c9 as nil")
	}
	b.ObjStart()
	b.PutID("searchMessages")
	b.Comma()
	b.FieldStart("chat_list")
	if s.ChatList == nil {
		return fmt.Errorf("unable to encode searchMessages#31f9c3c9: field chat_list is nil")
	}
	if err := s.ChatList.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchMessages#31f9c3c9: field chat_list: %w", err)
	}
	b.Comma()
	b.FieldStart("only_in_channels")
	b.PutBool(s.OnlyInChannels)
	b.Comma()
	b.FieldStart("query")
	b.PutString(s.Query)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(s.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.Comma()
	b.FieldStart("filter")
	if s.Filter == nil {
		return fmt.Errorf("unable to encode searchMessages#31f9c3c9: field filter is nil")
	}
	if err := s.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchMessages#31f9c3c9: field filter: %w", err)
	}
	b.Comma()
	b.FieldStart("min_date")
	b.PutInt32(s.MinDate)
	b.Comma()
	b.FieldStart("max_date")
	b.PutInt32(s.MaxDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchMessages#31f9c3c9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchMessages"); err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: %w", err)
			}
		case "chat_list":
			value, err := DecodeTDLibJSONChatList(b)
			if err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field chat_list: %w", err)
			}
			s.ChatList = value
		case "only_in_channels":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field only_in_channels: %w", err)
			}
			s.OnlyInChannels = value
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field query: %w", err)
			}
			s.Query = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field offset: %w", err)
			}
			s.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field limit: %w", err)
			}
			s.Limit = value
		case "filter":
			value, err := DecodeTDLibJSONSearchMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field filter: %w", err)
			}
			s.Filter = value
		case "min_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field min_date: %w", err)
			}
			s.MinDate = value
		case "max_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchMessages#31f9c3c9: field max_date: %w", err)
			}
			s.MaxDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatList returns value of ChatList field.
func (s *SearchMessagesRequest) GetChatList() (value ChatListClass) {
	if s == nil {
		return
	}
	return s.ChatList
}

// GetOnlyInChannels returns value of OnlyInChannels field.
func (s *SearchMessagesRequest) GetOnlyInChannels() (value bool) {
	if s == nil {
		return
	}
	return s.OnlyInChannels
}

// GetQuery returns value of Query field.
func (s *SearchMessagesRequest) GetQuery() (value string) {
	if s == nil {
		return
	}
	return s.Query
}

// GetOffset returns value of Offset field.
func (s *SearchMessagesRequest) GetOffset() (value string) {
	if s == nil {
		return
	}
	return s.Offset
}

// GetLimit returns value of Limit field.
func (s *SearchMessagesRequest) GetLimit() (value int32) {
	if s == nil {
		return
	}
	return s.Limit
}

// GetFilter returns value of Filter field.
func (s *SearchMessagesRequest) GetFilter() (value SearchMessagesFilterClass) {
	if s == nil {
		return
	}
	return s.Filter
}

// GetMinDate returns value of MinDate field.
func (s *SearchMessagesRequest) GetMinDate() (value int32) {
	if s == nil {
		return
	}
	return s.MinDate
}

// GetMaxDate returns value of MaxDate field.
func (s *SearchMessagesRequest) GetMaxDate() (value int32) {
	if s == nil {
		return
	}
	return s.MaxDate
}

// SearchMessages invokes method searchMessages#31f9c3c9 returning error if any.
func (c *Client) SearchMessages(ctx context.Context, request *SearchMessagesRequest) (*FoundMessages, error) {
	var result FoundMessages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
