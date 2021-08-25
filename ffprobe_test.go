// Copyright 2021 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ffmpeg_test

import (
	"context"
	"testing"

	"github.com/searKing/ffmpeg"
)

const (
	TestInputFile1 = "./examples/sample_data/in1.mp4"
)

// {
//    "streams": [
//        {
//            "index": 0,
//            "codec_name": "h264",
//            "codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
//            "profile": "Main",
//            "codec_type": "video",
//            "codec_time_base": "1001/60000",
//            "codec_tag_string": "avc1",
//            "codec_tag": "0x31637661",
//            "width": 320,
//            "height": 240,
//            "coded_width": 320,
//            "coded_height": 240,
//            "closed_captions": 0,
//            "has_b_frames": 0,
//            "sample_aspect_ratio": "1:1",
//            "display_aspect_ratio": "4:3",
//            "pix_fmt": "yuv420p",
//            "level": 13,
//            "chroma_location": "left",
//            "refs": 1,
//            "is_avc": "true",
//            "nal_length_size": "4",
//            "r_frame_rate": "30000/1001",
//            "avg_frame_rate": "30000/1001",
//            "time_base": "1/90000",
//            "start_pts": 0,
//            "start_time": "0.000000",
//            "duration_ts": 627627,
//            "duration": "6.973633",
//            "bit_rate": "251413",
//            "bits_per_raw_sample": "8",
//            "nb_frames": "209",
//            "disposition": {
//                "default": 1,
//                "dub": 0,
//                "original": 0,
//                "comment": 0,
//                "lyrics": 0,
//                "karaoke": 0,
//                "forced": 0,
//                "hearing_impaired": 0,
//                "visual_impaired": 0,
//                "clean_effects": 0,
//                "attached_pic": 0,
//                "timed_thumbnails": 0
//            },
//            "tags": {
//                "language": "und",
//                "handler_name": "VideoHandler"
//            }
//        },
//        {
//            "index": 1,
//            "codec_name": "aac",
//            "codec_long_name": "AAC (Advanced Audio Coding)",
//            "profile": "LC",
//            "codec_type": "audio",
//            "codec_time_base": "1/44100",
//            "codec_tag_string": "mp4a",
//            "codec_tag": "0x6134706d",
//            "sample_fmt": "fltp",
//            "sample_rate": "44100",
//            "channels": 2,
//            "channel_layout": "stereo",
//            "bits_per_sample": 0,
//            "r_frame_rate": "0/0",
//            "avg_frame_rate": "0/0",
//            "time_base": "1/44100",
//            "start_pts": 0,
//            "start_time": "0.000000",
//            "duration_ts": 310272,
//            "duration": "7.035646",
//            "bit_rate": "125587",
//            "max_bit_rate": "125587",
//            "nb_frames": "303",
//            "disposition": {
//                "default": 1,
//                "dub": 0,
//                "original": 0,
//                "comment": 0,
//                "lyrics": 0,
//                "karaoke": 0,
//                "forced": 0,
//                "hearing_impaired": 0,
//                "visual_impaired": 0,
//                "clean_effects": 0,
//                "attached_pic": 0,
//                "timed_thumbnails": 0
//            },
//            "tags": {
//                "language": "und",
//                "handler_name": "SoundHandler"
//            }
//        }
//    ],
//    "format": {
//        "filename": "./examples/sample_data/in1.mp4",
//        "nb_streams": 2,
//        "nb_programs": 0,
//        "format_name": "mov,mp4,m4a,3gp,3g2,mj2",
//        "format_long_name": "QuickTime / MOV",
//        "start_time": "0.000000",
//        "duration": "7.036000",
//        "size": "336833",
//        "bit_rate": "382982",
//        "probe_score": 100,
//        "tags": {
//            "major_brand": "isom",
//            "minor_version": "512",
//            "compatible_brands": "isomiso2avc1mp41",
//            "encoder": "Lavf57.71.100"
//        }
//    }
//}
func TestProbeJson(t *testing.T) {
	got, err := ffmpeg.ProbeJson(context.Background(), TestInputFile1)
	if err != nil {
		t.Fatalf("ffmpeg probe failed: %s", err.Error())
	}
	t.Logf("%s\n%s", TestInputFile1, got)
}

func TestProbeFormatAndStream(t *testing.T) {
	probeInfo, err := ffmpeg.ProbeFormatAndStream(context.Background(), TestInputFile1)
	if err != nil {
		t.Fatalf("ffmpeg probe failed: %s", err.Error())
	}

	if probeInfo.Format.Duration.String() !="7.036000" {
		t.Fatalf("got %q, want %q", probeInfo.Format.Duration, "7.036000")
	}
}
