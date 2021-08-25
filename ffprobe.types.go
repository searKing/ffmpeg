// Copyright 2021 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ffmpeg

import "encoding/json"

// See also https://github.com/FFmpeg/FFmpeg/blob/master/doc/ffprobe.xsd

type FFprobeType struct {
	ProgramVersion   *ProgramVersionType    `json:"program_version,omitempty"`
	LibraryVersions  []LibraryVersionType   `json:"library_versions,omitempty"`
	PixelFormats     []PixelFormatType      `json:"pixel_formats,omitempty"`
	Packets          []PacketType           `json:"packets,omitempty"`
	Frames           *FramesType            `json:"frames,omitempty"`
	PacketsAndFrames []PacketsAndFramesType `json:"packets_and_frames,omitempty"`
	Programs         []ProgramType          `json:"programs,omitempty"`
	Streams          []StreamType           `json:"streams,omitempty"`
	Chapters         []ChapterType          `json:"chapters,omitempty"`
	Format           *FormatType            `json:"format,omitempty"`
	Error            *ErrorType             `json:"error,omitempty"`
}

type PacketsType struct {
	Packet []PacketType `json:"packet,omitempty"`
}

type FramesType struct {
	Frame    FrameType    `json:"frame,omitempty"`
	Subtitle SubtitleType `json:"subtitle,omitempty"`
}

type PacketsAndFramesType struct {
	Packet   PacketType   `json:"packet,omitempty"`
	Frame    FrameType    `json:"frame,omitempty"`
	Subtitle SubtitleType `json:"subtitle,omitempty"`
}

type PacketType struct {
	Tag          []TagType            `json:"tag,omitempty"`
	SideDataList []PacketSideDataType `json:"side_data_list,omitempty"`
	CodecType    string               `json:"codec_type,omitempty"`
	StreamIndex  int                  `json:"stream_index,omitempty"`
	Pts          int64                `json:"pts,omitempty"`
	PtsTime      json.Number          `json:"pts_time,omitempty"` // float64
	Dts          int64                `json:"dts,omitempty"`
	DtsTime      json.Number          `json:"dts_time,omitempty"` // float64
	Duration     int64                `json:"duration,omitempty"`
	DurationTime json.Number          `json:"duration_time,omitempty"` // float64
	Size         int64                `json:"size,omitempty"`
	Pos          int64                `json:"pos,omitempty"`
	Flags        string               `json:"flags,omitempty"`
	Data         string               `json:"data,omitempty"`
	DataHash     string               `json:"data_hash,omitempty"`
}

type PacketSideDataListType struct {
	SideData []PacketSideDataType `json:"side_data,omitempty"`
}

type PacketSideDataType struct {
	SideDataType string `json:"side_data_type,omitempty"`
	SideDataSize int    `json:"side_data_size,omitempty"`
}

type FrameType struct {
	Tag          []TagType           `json:"tag,omitempty"`
	Logs         []LogType           `json:"logs,omitempty"`
	SideDataList []FrameSideDataType `json:"side_data_list,omitempty"`
}

type LogsType struct {
	Log []LogType `json:"log,omitempty"`
}

type LogType struct {
	Context        string `json:"context,omitempty"`
	Level          int    `json:"level,omitempty"`
	Category       int    `json:"category,omitempty"`
	ParentContext  string `json:"parent_context,omitempty"`
	ParentCategory int    `json:"parent_category,omitempty"`
	Message        string `json:"message,omitempty"`
}

type FrameSideDataListType struct {
	SideData []FrameSideDataType `json:"side_data,omitempty"`
}

type FrameSideDataType struct {
	Timecodes    []FrameSideDataTimecodeType `json:"timecodes,omitempty"`
	SideDataType string                      `json:"side_data_type,omitempty"`
	SideDataSize int                         `json:"side_data_size,omitempty"`
	Timecode     string                      `json:"timecode,omitempty"`
}

type FrameSideDataTimecodeList struct {
	Timecode []FrameSideDataTimecodeType `json:"timecode,omitempty"`
}

type FrameSideDataTimecodeType struct {
	Value string `json:"value,omitempty"`
}

type SubtitleType struct {
	// Deprecated: Use Subtitle instead.
	MediaType        string      `json:"media_type,omitempty"`
	Subtitle         string      `json:"subtitle,omitempty"`
	Pts              int64       `json:"pts,omitempty"`
	PtsTime          json.Number `json:"pts_time,omitempty"` // float64
	Format           int         `json:"format,omitempty"`
	StartDisplayTime int         `json:"start_display_time,omitempty"`
	EndDisplayTime   int         `json:"end_display_time,omitempty"`
	NumRects         int         `json:"num_rects,omitempty"`
}

type StreamsType struct {
	Stream []StreamType `json:"stream,omitempty"`
}

type ProgramsType struct {
	Program []ProgramType `json:"program,omitempty"`
}

type StreamDispositionType struct {
	Default         int `json:"default,omitempty"`
	Dub             int `json:"dub,omitempty"`
	Original        int `json:"original,omitempty"`
	Comment         int `json:"comment,omitempty"`
	Lyrics          int `json:"lyrics,omitempty"`
	Karaoke         int `json:"karaoke,omitempty"`
	Forced          int `json:"forced,omitempty"`
	HearingImpaired int `json:"hearing_impaired,omitempty"`
	VisualImpaired  int `json:"visual_impaired,omitempty"`
	CleanEffects    int `json:"clean_effects,omitempty"`
	AttachedPic     int `json:"attached_pic,omitempty"`
	TimedThumbnails int `json:"timed_thumbnails,omitempty"`
	Captions        int `json:"captions,omitempty"`
	Descriptions    int `json:"descriptions,omitempty"`
	Metadata        int `json:"metadata,omitempty"`
	Dependent       int `json:"dependent,omitempty"`
	StillImage      int `json:"still_image,omitempty"`
}

type StreamType struct {
	Disposition   *StreamDispositionType `json:"disposition,omitempty"`
	Tag           []TagType              `json:"tag,omitempty"`
	SideDataList  []PacketSideDataType   `json:"side_data_list,omitempty"`
	Index         int                    `json:"index,omitempty"`
	CodecName     string                 `json:"codec_name,omitempty"`
	CodecLongName string                 `json:"codec_long_name,omitempty"`
	Profile       string                 `json:"profile,omitempty"`
	CodecType     string                 `json:"codec_type,omitempty"`
	CodecTag      string                 `json:"codec_tag,omitempty"`
	CodeTagString string                 `json:"code_tag_string,omitempty"`
	Extradata     string                 `json:"extradata,omitempty"`
	ExtradataHash string                 `json:"extradata_hash,omitempty"`
	// video attributes
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	CodeWidth          int    `json:"code_width,omitempty"`
	CodeHeight         int    `json:"code_height,omitempty"`
	ClosedCaptions     int    `json:"closed_captions,omitempty"` // bool
	FilmGrain          int    `json:"film_grain,omitempty"`      // bool
	HasBFrames         int    `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
	PixFmt             string `json:"pix_fmt,omitempty"`
	Level              int    `json:"level,omitempty"`
	ColorRange         string `json:"color_range,omitempty"`
	ColorSpace         string `json:"color_space,omitempty"`
	ColorTransfer      string `json:"color_transfer,omitempty"`
	ColorPrimaries     string `json:"color_primaries,omitempty"`
	ChromaLocation     string `json:"chroma_location,omitempty"`
	FieldOrder         string `json:"field_order,omitempty"`
	Refs               int    `json:"refs,omitempty"`
	// audio attributes
	SampleFmt     string      `json:"sample_fmt,omitempty"`
	SampleRate    json.Number `json:"sample_rate,omitempty"` // int
	Channels      int         `json:"channels,omitempty"`
	ChannelLayout string      `json:"channel_layout,omitempty"`
	BitsPerSample int         `json:"bits_per_sample,omitempty"`

	Id               string      `json:"id,omitempty"`
	RFrameRate       string      `json:"r_frame_rate,omitempty"`
	AvgFrameRate     string      `json:"avg_frame_rate,omitempty"`
	TimeBase         string      `json:"time_base,omitempty"`
	StartPts         int64       `json:"start_pts,omitempty"`
	StartTime        json.Number `json:"start_time,omitempty"` // float64
	DurationTs       int64       `json:"duration_ts,omitempty"`
	Duration         json.Number `json:"duration,omitempty"`            // float64
	BitRate          json.Number `json:"bit_rate,omitempty"`            // int
	MaxBitRate       json.Number `json:"max_bit_rate,omitempty"`        // int
	BitsPerRawSample json.Number `json:"bits_per_raw_sample,omitempty"` // int
	NbFrames         json.Number `json:"nb_frames,omitempty"`           // int
	NbReadFrames     json.Number `json:"nb_read_frames,omitempty"`      // int
	NbReadPackets    json.Number `json:"nb_read_packets,omitempty"`     // int
}

type ProgramType struct {
	Tag        []TagType    `json:"tag,omitempty"`
	Streams    []StreamType `json:"streams,omitempty"`
	ProgramId  int          `json:"program_id,omitempty"`
	ProgramNum int          `json:"program_num,omitempty"`
	NbStreams  int          `json:"nb_streams,omitempty"`
	StartTime  json.Number  `json:"start_time,omitempty"` // float64
	StartPts   int64        `json:"start_pts,omitempty"`
	EndTime    json.Number  `json:"end_time,omitempty"` // float64
	EndPts     int64        `json:"end_pts,omitempty"`
	PmyPid     int          `json:"pmy_pid,omitempty"`
	PcrPid     int          `json:"pcr_pid,omitempty"`
}

type FormatType struct {
	Tag            []TagType   `json:"tag,omitempty"`
	Filename       string      `json:"filename,omitempty"`
	NbStreams      int         `json:"nb_streams,omitempty"`
	NbPrograms     int         `json:"nb_programs,omitempty"`
	FormatName     string      `json:"format_name,omitempty"`
	FormatLongName string      `json:"format_long_name,omitempty"`
	StartTime      json.Number `json:"start_time,omitempty"` // float64
	Duration       json.Number `json:"duration,omitempty"`   // float64
	Size           json.Number `json:"size,omitempty"`       // int64
	BitRate        json.Number `json:"bit_rate,omitempty"`   // int64
	ProbeScore     int         `json:"probe_score,omitempty"`
}

type TagType struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type ErrorType struct {
	Code   int    `json:"code,omitempty"`
	String string `json:"string,omitempty"`
}

type ProgramVersionType struct {
	Version       string `json:"version,omitempty"`
	Copyright     string `json:"copy_right,omitempty"`
	BuildData     string `json:"build_data,omitempty"`
	BuildTime     string `json:"build_time,omitempty"`
	CompilerIdent string `json:"compiler_ident,omitempty"`
	Configuration string `json:"configuration,omitempty"`
}

type ChaptersType struct {
	Chapter []ChapterType `json:"chapter,omitempty"`
}

type ChapterType struct {
	Tag       []TagType   `json:"tag,omitempty"`
	Id        int         `json:"id,omitempty"`
	TimeBase  string      `json:"time_base,omitempty"`
	Start     int         `json:"start,omitempty"`
	StartTime json.Number `json:"start_time,omitempty"` // float64
	End       int         `json:"end,omitempty"`
	EndTime   json.Number `json:"end_time,omitempty"` // float64
}

type LibraryVersionType struct {
	Name    string `json:"name,omitempty"`
	Major   int    `json:"major,omitempty"`
	Minor   int    `json:"minor,omitempty"`
	Micro   int    `json:"micro,omitempty"`
	Version int    `json:"version,omitempty"`
	Ident   string `json:"ident,omitempty"`
}

type LibraryVersionsType struct {
	LibraryVersion []LibraryVersionType `json:"library_version,omitempty"`
}

type PixelFormatFlagsType struct {
	BigEndian int `json:"big_endian,omitempty"`
	Palette   int `json:"palette,omitempty"`
	Bitstream int `json:"bitstream,omitempty"`
	Hwaccel   int `json:"hwaccel,omitempty"`
	Planar    int `json:"planar,omitempty"`
	Rgb       int `json:"rgb,omitempty"`
	Alpha     int `json:"alpha,omitempty"`
}

type PixelFormatComponentType struct {
	Index    int `json:"index,omitempty"`
	BitDepth int `json:"bit_depth,omitempty"`
}

type PixelFormatComponentsType struct {
	Component []PixelFormatComponentType `json:"component,omitempty"`
}

type PixelFormatType struct {
	Flags        *PixelFormatFlagsType      `json:"flags,omitempty"`
	Components   []PixelFormatComponentType `json:"components,omitempty"`
	Name         string                     `json:"name,omitempty"`
	NbComponents int                        `json:"nb_components,omitempty"`
	Log2ChromaW  int                        `json:"log2_chroma_w,omitempty"`
	Log2ChromaH  int                        `json:"log2_chroma_h,omitempty"`
	BitsPerPixel int                        `json:"bits_per_pixel,omitempty"`
}

type PixelFormatsType struct {
	PixelFormat []PixelFormatType `json:"pixel_format,omitempty"`
}
