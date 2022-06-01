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

// AnswerWebAppQueryRequest represents TL type `answerWebAppQuery#a0b49cf1`.
type AnswerWebAppQueryRequest struct {
	// Identifier of the web app query
	WebAppQueryID string
	// The result of the query
	Result InputInlineQueryResultClass
}

// AnswerWebAppQueryRequestTypeID is TL type id of AnswerWebAppQueryRequest.
const AnswerWebAppQueryRequestTypeID = 0xa0b49cf1

// Ensuring interfaces in compile-time for AnswerWebAppQueryRequest.
var (
	_ bin.Encoder     = &AnswerWebAppQueryRequest{}
	_ bin.Decoder     = &AnswerWebAppQueryRequest{}
	_ bin.BareEncoder = &AnswerWebAppQueryRequest{}
	_ bin.BareDecoder = &AnswerWebAppQueryRequest{}
)

func (a *AnswerWebAppQueryRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.WebAppQueryID == "") {
		return false
	}
	if !(a.Result == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AnswerWebAppQueryRequest) String() string {
	if a == nil {
		return "AnswerWebAppQueryRequest(nil)"
	}
	type Alias AnswerWebAppQueryRequest
	return fmt.Sprintf("AnswerWebAppQueryRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AnswerWebAppQueryRequest) TypeID() uint32 {
	return AnswerWebAppQueryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AnswerWebAppQueryRequest) TypeName() string {
	return "answerWebAppQuery"
}

// TypeInfo returns info about TL type.
func (a *AnswerWebAppQueryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "answerWebAppQuery",
		ID:   AnswerWebAppQueryRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "WebAppQueryID",
			SchemaName: "web_app_query_id",
		},
		{
			Name:       "Result",
			SchemaName: "result",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AnswerWebAppQueryRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode answerWebAppQuery#a0b49cf1 as nil")
	}
	b.PutID(AnswerWebAppQueryRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AnswerWebAppQueryRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode answerWebAppQuery#a0b49cf1 as nil")
	}
	b.PutString(a.WebAppQueryID)
	if a.Result == nil {
		return fmt.Errorf("unable to encode answerWebAppQuery#a0b49cf1: field result is nil")
	}
	if err := a.Result.Encode(b); err != nil {
		return fmt.Errorf("unable to encode answerWebAppQuery#a0b49cf1: field result: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AnswerWebAppQueryRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode answerWebAppQuery#a0b49cf1 to nil")
	}
	if err := b.ConsumeID(AnswerWebAppQueryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode answerWebAppQuery#a0b49cf1: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AnswerWebAppQueryRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode answerWebAppQuery#a0b49cf1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode answerWebAppQuery#a0b49cf1: field web_app_query_id: %w", err)
		}
		a.WebAppQueryID = value
	}
	{
		value, err := DecodeInputInlineQueryResult(b)
		if err != nil {
			return fmt.Errorf("unable to decode answerWebAppQuery#a0b49cf1: field result: %w", err)
		}
		a.Result = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AnswerWebAppQueryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode answerWebAppQuery#a0b49cf1 as nil")
	}
	b.ObjStart()
	b.PutID("answerWebAppQuery")
	b.Comma()
	b.FieldStart("web_app_query_id")
	b.PutString(a.WebAppQueryID)
	b.Comma()
	b.FieldStart("result")
	if a.Result == nil {
		return fmt.Errorf("unable to encode answerWebAppQuery#a0b49cf1: field result is nil")
	}
	if err := a.Result.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode answerWebAppQuery#a0b49cf1: field result: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AnswerWebAppQueryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode answerWebAppQuery#a0b49cf1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("answerWebAppQuery"); err != nil {
				return fmt.Errorf("unable to decode answerWebAppQuery#a0b49cf1: %w", err)
			}
		case "web_app_query_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode answerWebAppQuery#a0b49cf1: field web_app_query_id: %w", err)
			}
			a.WebAppQueryID = value
		case "result":
			value, err := DecodeTDLibJSONInputInlineQueryResult(b)
			if err != nil {
				return fmt.Errorf("unable to decode answerWebAppQuery#a0b49cf1: field result: %w", err)
			}
			a.Result = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetWebAppQueryID returns value of WebAppQueryID field.
func (a *AnswerWebAppQueryRequest) GetWebAppQueryID() (value string) {
	if a == nil {
		return
	}
	return a.WebAppQueryID
}

// GetResult returns value of Result field.
func (a *AnswerWebAppQueryRequest) GetResult() (value InputInlineQueryResultClass) {
	if a == nil {
		return
	}
	return a.Result
}

// AnswerWebAppQuery invokes method answerWebAppQuery#a0b49cf1 returning error if any.
func (c *Client) AnswerWebAppQuery(ctx context.Context, request *AnswerWebAppQueryRequest) (*SentWebAppMessage, error) {
	var result SentWebAppMessage

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}