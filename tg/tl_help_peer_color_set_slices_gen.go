//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// HelpPeerColorSetClassArray is adapter for slice of HelpPeerColorSetClass.
type HelpPeerColorSetClassArray []HelpPeerColorSetClass

// Sort sorts slice of HelpPeerColorSetClass.
func (s HelpPeerColorSetClassArray) Sort(less func(a, b HelpPeerColorSetClass) bool) HelpPeerColorSetClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpPeerColorSetClass.
func (s HelpPeerColorSetClassArray) SortStable(less func(a, b HelpPeerColorSetClass) bool) HelpPeerColorSetClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpPeerColorSetClass.
func (s HelpPeerColorSetClassArray) Retain(keep func(x HelpPeerColorSetClass) bool) HelpPeerColorSetClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s HelpPeerColorSetClassArray) First() (v HelpPeerColorSetClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPeerColorSetClassArray) Last() (v HelpPeerColorSetClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPeerColorSetClassArray) PopFirst() (v HelpPeerColorSetClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpPeerColorSetClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpPeerColorSetClassArray) Pop() (v HelpPeerColorSetClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpPeerColorSet returns copy with only HelpPeerColorSet constructors.
func (s HelpPeerColorSetClassArray) AsHelpPeerColorSet() (to HelpPeerColorSetArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpPeerColorSet)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsHelpPeerColorProfileSet returns copy with only HelpPeerColorProfileSet constructors.
func (s HelpPeerColorSetClassArray) AsHelpPeerColorProfileSet() (to HelpPeerColorProfileSetArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpPeerColorProfileSet)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// HelpPeerColorSetArray is adapter for slice of HelpPeerColorSet.
type HelpPeerColorSetArray []HelpPeerColorSet

// Sort sorts slice of HelpPeerColorSet.
func (s HelpPeerColorSetArray) Sort(less func(a, b HelpPeerColorSet) bool) HelpPeerColorSetArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpPeerColorSet.
func (s HelpPeerColorSetArray) SortStable(less func(a, b HelpPeerColorSet) bool) HelpPeerColorSetArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpPeerColorSet.
func (s HelpPeerColorSetArray) Retain(keep func(x HelpPeerColorSet) bool) HelpPeerColorSetArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s HelpPeerColorSetArray) First() (v HelpPeerColorSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPeerColorSetArray) Last() (v HelpPeerColorSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPeerColorSetArray) PopFirst() (v HelpPeerColorSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpPeerColorSet
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpPeerColorSetArray) Pop() (v HelpPeerColorSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// HelpPeerColorProfileSetArray is adapter for slice of HelpPeerColorProfileSet.
type HelpPeerColorProfileSetArray []HelpPeerColorProfileSet

// Sort sorts slice of HelpPeerColorProfileSet.
func (s HelpPeerColorProfileSetArray) Sort(less func(a, b HelpPeerColorProfileSet) bool) HelpPeerColorProfileSetArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpPeerColorProfileSet.
func (s HelpPeerColorProfileSetArray) SortStable(less func(a, b HelpPeerColorProfileSet) bool) HelpPeerColorProfileSetArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpPeerColorProfileSet.
func (s HelpPeerColorProfileSetArray) Retain(keep func(x HelpPeerColorProfileSet) bool) HelpPeerColorProfileSetArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s HelpPeerColorProfileSetArray) First() (v HelpPeerColorProfileSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPeerColorProfileSetArray) Last() (v HelpPeerColorProfileSet, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPeerColorProfileSetArray) PopFirst() (v HelpPeerColorProfileSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpPeerColorProfileSet
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpPeerColorProfileSetArray) Pop() (v HelpPeerColorProfileSet, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
