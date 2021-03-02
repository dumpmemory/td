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

// HelpGetCdnConfigRequest represents TL type `help.getCdnConfig#52029342`.
// Get configuration for CDN¹ file downloads.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/help.getCdnConfig for reference.
type HelpGetCdnConfigRequest struct {
}

// HelpGetCdnConfigRequestTypeID is TL type id of HelpGetCdnConfigRequest.
const HelpGetCdnConfigRequestTypeID = 0x52029342

func (g *HelpGetCdnConfigRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetCdnConfigRequest) String() string {
	if g == nil {
		return "HelpGetCdnConfigRequest(nil)"
	}
	type Alias HelpGetCdnConfigRequest
	return fmt.Sprintf("HelpGetCdnConfigRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetCdnConfigRequest) TypeID() uint32 {
	return HelpGetCdnConfigRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetCdnConfigRequest) TypeName() string {
	return "help.getCdnConfig"
}

// TypeInfo returns info about TL type.
func (g *HelpGetCdnConfigRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getCdnConfig",
		ID:   HelpGetCdnConfigRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetCdnConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getCdnConfig#52029342 as nil")
	}
	b.PutID(HelpGetCdnConfigRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetCdnConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getCdnConfig#52029342 to nil")
	}
	if err := b.ConsumeID(HelpGetCdnConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getCdnConfig#52029342: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetCdnConfigRequest.
var (
	_ bin.Encoder = &HelpGetCdnConfigRequest{}
	_ bin.Decoder = &HelpGetCdnConfigRequest{}
)

// HelpGetCdnConfig invokes method help.getCdnConfig#52029342 returning error if any.
// Get configuration for CDN¹ file downloads.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// Possible errors:
//  401 AUTH_KEY_PERM_EMPTY: The temporary auth key must be binded to the permanent auth key to use these methods.
//
// See https://core.telegram.org/method/help.getCdnConfig for reference.
// Can be used by bots.
func (c *Client) HelpGetCdnConfig(ctx context.Context) (*CdnConfig, error) {
	var result CdnConfig

	request := &HelpGetCdnConfigRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
