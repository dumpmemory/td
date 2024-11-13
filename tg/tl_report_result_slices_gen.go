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

// ReportResultClassArray is adapter for slice of ReportResultClass.
type ReportResultClassArray []ReportResultClass

// Sort sorts slice of ReportResultClass.
func (s ReportResultClassArray) Sort(less func(a, b ReportResultClass) bool) ReportResultClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReportResultClass.
func (s ReportResultClassArray) SortStable(less func(a, b ReportResultClass) bool) ReportResultClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReportResultClass.
func (s ReportResultClassArray) Retain(keep func(x ReportResultClass) bool) ReportResultClassArray {
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
func (s ReportResultClassArray) First() (v ReportResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReportResultClassArray) Last() (v ReportResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReportResultClassArray) PopFirst() (v ReportResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReportResultClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReportResultClassArray) Pop() (v ReportResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsReportResultChooseOption returns copy with only ReportResultChooseOption constructors.
func (s ReportResultClassArray) AsReportResultChooseOption() (to ReportResultChooseOptionArray) {
	for _, elem := range s {
		value, ok := elem.(*ReportResultChooseOption)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsReportResultAddComment returns copy with only ReportResultAddComment constructors.
func (s ReportResultClassArray) AsReportResultAddComment() (to ReportResultAddCommentArray) {
	for _, elem := range s {
		value, ok := elem.(*ReportResultAddComment)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ReportResultChooseOptionArray is adapter for slice of ReportResultChooseOption.
type ReportResultChooseOptionArray []ReportResultChooseOption

// Sort sorts slice of ReportResultChooseOption.
func (s ReportResultChooseOptionArray) Sort(less func(a, b ReportResultChooseOption) bool) ReportResultChooseOptionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReportResultChooseOption.
func (s ReportResultChooseOptionArray) SortStable(less func(a, b ReportResultChooseOption) bool) ReportResultChooseOptionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReportResultChooseOption.
func (s ReportResultChooseOptionArray) Retain(keep func(x ReportResultChooseOption) bool) ReportResultChooseOptionArray {
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
func (s ReportResultChooseOptionArray) First() (v ReportResultChooseOption, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReportResultChooseOptionArray) Last() (v ReportResultChooseOption, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReportResultChooseOptionArray) PopFirst() (v ReportResultChooseOption, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReportResultChooseOption
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReportResultChooseOptionArray) Pop() (v ReportResultChooseOption, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ReportResultAddCommentArray is adapter for slice of ReportResultAddComment.
type ReportResultAddCommentArray []ReportResultAddComment

// Sort sorts slice of ReportResultAddComment.
func (s ReportResultAddCommentArray) Sort(less func(a, b ReportResultAddComment) bool) ReportResultAddCommentArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ReportResultAddComment.
func (s ReportResultAddCommentArray) SortStable(less func(a, b ReportResultAddComment) bool) ReportResultAddCommentArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ReportResultAddComment.
func (s ReportResultAddCommentArray) Retain(keep func(x ReportResultAddComment) bool) ReportResultAddCommentArray {
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
func (s ReportResultAddCommentArray) First() (v ReportResultAddComment, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReportResultAddCommentArray) Last() (v ReportResultAddComment, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReportResultAddCommentArray) PopFirst() (v ReportResultAddComment, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ReportResultAddComment
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ReportResultAddCommentArray) Pop() (v ReportResultAddComment, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}