// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// Invoker can invoke raw MTProto rpc calls.
type Invoker interface {
	InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error
}

// Client implement methods for calling functions from TL schema via Invoker.
type Client struct {
	rpc Invoker
}

func NewClient(invoker Invoker) *Client {
	return &Client{
		rpc: invoker,
	}
}
