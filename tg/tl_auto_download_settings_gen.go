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

// AutoDownloadSettings represents TL type `autoDownloadSettings#e04232f3`.
// Autodownload settings
//
// See https://core.telegram.org/constructor/autoDownloadSettings for reference.
type AutoDownloadSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable automatic media downloads?
	Disabled bool
	// Whether to preload the first seconds of videos larger than the specified limit
	VideoPreloadLarge bool
	// Whether to preload the next audio track when you're listening to music
	AudioPreloadNext bool
	// Whether to enable data saving mode in phone calls
	PhonecallsLessData bool
	// Maximum size of photos to preload
	PhotoSizeMax int
	// Maximum size of videos to preload
	VideoSizeMax int
	// Maximum size of other files to preload
	FileSizeMax int
	// Maximum suggested bitrate for uploading videos
	VideoUploadMaxbitrate int
}

// AutoDownloadSettingsTypeID is TL type id of AutoDownloadSettings.
const AutoDownloadSettingsTypeID = 0xe04232f3

func (a *AutoDownloadSettings) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Disabled == false) {
		return false
	}
	if !(a.VideoPreloadLarge == false) {
		return false
	}
	if !(a.AudioPreloadNext == false) {
		return false
	}
	if !(a.PhonecallsLessData == false) {
		return false
	}
	if !(a.PhotoSizeMax == 0) {
		return false
	}
	if !(a.VideoSizeMax == 0) {
		return false
	}
	if !(a.FileSizeMax == 0) {
		return false
	}
	if !(a.VideoUploadMaxbitrate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AutoDownloadSettings) String() string {
	if a == nil {
		return "AutoDownloadSettings(nil)"
	}
	type Alias AutoDownloadSettings
	return fmt.Sprintf("AutoDownloadSettings%+v", Alias(*a))
}

// FillFrom fills AutoDownloadSettings from given interface.
func (a *AutoDownloadSettings) FillFrom(from interface {
	GetDisabled() (value bool)
	GetVideoPreloadLarge() (value bool)
	GetAudioPreloadNext() (value bool)
	GetPhonecallsLessData() (value bool)
	GetPhotoSizeMax() (value int)
	GetVideoSizeMax() (value int)
	GetFileSizeMax() (value int)
	GetVideoUploadMaxbitrate() (value int)
}) {
	a.Disabled = from.GetDisabled()
	a.VideoPreloadLarge = from.GetVideoPreloadLarge()
	a.AudioPreloadNext = from.GetAudioPreloadNext()
	a.PhonecallsLessData = from.GetPhonecallsLessData()
	a.PhotoSizeMax = from.GetPhotoSizeMax()
	a.VideoSizeMax = from.GetVideoSizeMax()
	a.FileSizeMax = from.GetFileSizeMax()
	a.VideoUploadMaxbitrate = from.GetVideoUploadMaxbitrate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AutoDownloadSettings) TypeID() uint32 {
	return AutoDownloadSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*AutoDownloadSettings) TypeName() string {
	return "autoDownloadSettings"
}

// TypeInfo returns info about TL type.
func (a *AutoDownloadSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "autoDownloadSettings",
		ID:   AutoDownloadSettingsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Flags",
			SchemaName: "flags",
		},
		{
			Name:       "Disabled",
			SchemaName: "disabled",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "VideoPreloadLarge",
			SchemaName: "video_preload_large",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "AudioPreloadNext",
			SchemaName: "audio_preload_next",
			Null:       !a.Flags.Has(2),
		},
		{
			Name:       "PhonecallsLessData",
			SchemaName: "phonecalls_less_data",
			Null:       !a.Flags.Has(3),
		},
		{
			Name:       "PhotoSizeMax",
			SchemaName: "photo_size_max",
		},
		{
			Name:       "VideoSizeMax",
			SchemaName: "video_size_max",
		},
		{
			Name:       "FileSizeMax",
			SchemaName: "file_size_max",
		},
		{
			Name:       "VideoUploadMaxbitrate",
			SchemaName: "video_upload_maxbitrate",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AutoDownloadSettings) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoDownloadSettings#e04232f3 as nil")
	}
	b.PutID(AutoDownloadSettingsTypeID)
	if !(a.Disabled == false) {
		a.Flags.Set(0)
	}
	if !(a.VideoPreloadLarge == false) {
		a.Flags.Set(1)
	}
	if !(a.AudioPreloadNext == false) {
		a.Flags.Set(2)
	}
	if !(a.PhonecallsLessData == false) {
		a.Flags.Set(3)
	}
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autoDownloadSettings#e04232f3: field flags: %w", err)
	}
	b.PutInt(a.PhotoSizeMax)
	b.PutInt(a.VideoSizeMax)
	b.PutInt(a.FileSizeMax)
	b.PutInt(a.VideoUploadMaxbitrate)
	return nil
}

// SetDisabled sets value of Disabled conditional field.
func (a *AutoDownloadSettings) SetDisabled(value bool) {
	if value {
		a.Flags.Set(0)
		a.Disabled = true
	} else {
		a.Flags.Unset(0)
		a.Disabled = false
	}
}

// GetDisabled returns value of Disabled conditional field.
func (a *AutoDownloadSettings) GetDisabled() (value bool) {
	return a.Flags.Has(0)
}

// SetVideoPreloadLarge sets value of VideoPreloadLarge conditional field.
func (a *AutoDownloadSettings) SetVideoPreloadLarge(value bool) {
	if value {
		a.Flags.Set(1)
		a.VideoPreloadLarge = true
	} else {
		a.Flags.Unset(1)
		a.VideoPreloadLarge = false
	}
}

// GetVideoPreloadLarge returns value of VideoPreloadLarge conditional field.
func (a *AutoDownloadSettings) GetVideoPreloadLarge() (value bool) {
	return a.Flags.Has(1)
}

// SetAudioPreloadNext sets value of AudioPreloadNext conditional field.
func (a *AutoDownloadSettings) SetAudioPreloadNext(value bool) {
	if value {
		a.Flags.Set(2)
		a.AudioPreloadNext = true
	} else {
		a.Flags.Unset(2)
		a.AudioPreloadNext = false
	}
}

// GetAudioPreloadNext returns value of AudioPreloadNext conditional field.
func (a *AutoDownloadSettings) GetAudioPreloadNext() (value bool) {
	return a.Flags.Has(2)
}

// SetPhonecallsLessData sets value of PhonecallsLessData conditional field.
func (a *AutoDownloadSettings) SetPhonecallsLessData(value bool) {
	if value {
		a.Flags.Set(3)
		a.PhonecallsLessData = true
	} else {
		a.Flags.Unset(3)
		a.PhonecallsLessData = false
	}
}

// GetPhonecallsLessData returns value of PhonecallsLessData conditional field.
func (a *AutoDownloadSettings) GetPhonecallsLessData() (value bool) {
	return a.Flags.Has(3)
}

// GetPhotoSizeMax returns value of PhotoSizeMax field.
func (a *AutoDownloadSettings) GetPhotoSizeMax() (value int) {
	return a.PhotoSizeMax
}

// GetVideoSizeMax returns value of VideoSizeMax field.
func (a *AutoDownloadSettings) GetVideoSizeMax() (value int) {
	return a.VideoSizeMax
}

// GetFileSizeMax returns value of FileSizeMax field.
func (a *AutoDownloadSettings) GetFileSizeMax() (value int) {
	return a.FileSizeMax
}

// GetVideoUploadMaxbitrate returns value of VideoUploadMaxbitrate field.
func (a *AutoDownloadSettings) GetVideoUploadMaxbitrate() (value int) {
	return a.VideoUploadMaxbitrate
}

// Decode implements bin.Decoder.
func (a *AutoDownloadSettings) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoDownloadSettings#e04232f3 to nil")
	}
	if err := b.ConsumeID(AutoDownloadSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field flags: %w", err)
		}
	}
	a.Disabled = a.Flags.Has(0)
	a.VideoPreloadLarge = a.Flags.Has(1)
	a.AudioPreloadNext = a.Flags.Has(2)
	a.PhonecallsLessData = a.Flags.Has(3)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field photo_size_max: %w", err)
		}
		a.PhotoSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field video_size_max: %w", err)
		}
		a.VideoSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field file_size_max: %w", err)
		}
		a.FileSizeMax = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode autoDownloadSettings#e04232f3: field video_upload_maxbitrate: %w", err)
		}
		a.VideoUploadMaxbitrate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AutoDownloadSettings.
var (
	_ bin.Encoder = &AutoDownloadSettings{}
	_ bin.Decoder = &AutoDownloadSettings{}
)
