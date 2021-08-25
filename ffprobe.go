// Copyright 2021 The searKing Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ffmpeg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

var (

	// FFprobeOptionShowStream Show information about each media stream contained in the input multimedia stream.
	// Each media stream information is printed within a dedicated section with name "STREAM".
	FFprobeOptionShowStream = []string{"-show_streams"}
	// FFprobeOptionShowFormat Show information about the container format of the input multimedia stream.
	// All the container format information is printed within a section with name "FORMAT".
	FFprobeOptionShowFormat = []string{"-show_format"}

	// FFprobeOptionOutputPrintingFormatJson Set the output printing format.
	// writer_name specifies the name of the writer, and writer_options specifies the options to be passed to the writer.
	FFprobeOptionOutputPrintingFormatJson = []string{"-of", "json"}
)

func ProbeFormatAndStream(ctx context.Context, fileName string, args ...string) (FFprobeType, error) {
	args = append([]string{
		"-show_format",
		"-show_streams",
	}, args...)
	return ProbeFFprobeType(ctx, fileName, args...)
}

// ProbeFFprobeType Run ffprobe on the specified file and return a ProbeFFprobeType representation of the output.
// See also https://github.com/FFmpeg/FFmpeg/blob/master/doc/ffprobe.xsd
func ProbeFFprobeType(ctx context.Context, fileName string, args ...string) (FFprobeType, error) {
	data, err := ProbeJson(ctx, fileName, args...)
	if err != nil {
		return FFprobeType{}, err
	}
	var typ FFprobeType
	err = json.Unmarshal([]byte(data), &typ)
	if err != nil {
		return FFprobeType{}, err
	}
	return typ, nil
}

// ProbeJson Run ffprobe on the specified file and return a JSON representation of the output.
// See also https://ffmpeg.org/ffprobe-all.html#json
func ProbeJson(ctx context.Context, fileName string, args ...string) (string, error) {
	args = append([]string{
		"-of", "json",
	}, args...)
	return Probe(ctx, fileName, args...)
}

// Probe Run ffprobe on the specified file and return an output from stdout of ffprobe.
// See also https://ffmpeg.org/ffprobe-all.html
// See also https://github.com/FFmpeg/FFmpeg/blob/master/doc/ffprobe.xsd
func Probe(ctx context.Context, fileName string, args ...string) (string, error) {
	args = append(args, fileName)
	cmd := exec.CommandContext(ctx, "ffprobe", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	// check child exited gracefully, did not timeout
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%w, details:\n%s", err, stderr.String())
	}
	return string(stdout.Bytes()), nil
}

//
//err = cmd.Start()
//if err != nil {
//t.Fatalf("Start failed: %v", err)
//}
//go func() {
//	time.Sleep(1 * time.Second)
//	sendCtrlBreak(t, cmd.Process.Pid)
//}()
//err = cmd.Wait()
//if err != nil {
//t.Fatalf("Program exited with error: %v\n%v", err, string(b.Bytes()))
//}
