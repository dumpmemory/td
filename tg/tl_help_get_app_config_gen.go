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

// HelpGetAppConfigRequest represents TL type `help.getAppConfig#98914110`.
// Get app-specific configuration, see client configuration¹ for more info on the result.
//
// Links:
//  1) https://core.telegram.org/api/config#client-configuration
//
// See https://core.telegram.org/method/help.getAppConfig for reference.
type HelpGetAppConfigRequest struct {
}

// HelpGetAppConfigRequestTypeID is TL type id of HelpGetAppConfigRequest.
const HelpGetAppConfigRequestTypeID = 0x98914110

func (g *HelpGetAppConfigRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetAppConfigRequest) String() string {
	if g == nil {
		return "HelpGetAppConfigRequest(nil)"
	}
	type Alias HelpGetAppConfigRequest
	return fmt.Sprintf("HelpGetAppConfigRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetAppConfigRequest) TypeID() uint32 {
	return HelpGetAppConfigRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetAppConfigRequest) TypeName() string {
	return "help.getAppConfig"
}

// TypeInfo returns info about TL type.
func (g *HelpGetAppConfigRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getAppConfig",
		ID:   HelpGetAppConfigRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetAppConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getAppConfig#98914110 as nil")
	}
	b.PutID(HelpGetAppConfigRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetAppConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getAppConfig#98914110 to nil")
	}
	if err := b.ConsumeID(HelpGetAppConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getAppConfig#98914110: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetAppConfigRequest.
var (
	_ bin.Encoder = &HelpGetAppConfigRequest{}
	_ bin.Decoder = &HelpGetAppConfigRequest{}
)

// HelpGetAppConfig invokes method help.getAppConfig#98914110 returning error if any.
// Get app-specific configuration, see client configuration¹ for more info on the result.
//
// Links:
//  1) https://core.telegram.org/api/config#client-configuration
//
// See https://core.telegram.org/method/help.getAppConfig for reference.
func (c *Client) HelpGetAppConfig(ctx context.Context) (JSONValueClass, error) {
	var result JSONValueBox

	request := &HelpGetAppConfigRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.JSONValue, nil
}
