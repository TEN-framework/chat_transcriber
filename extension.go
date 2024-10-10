/**
 *
 * Agora Real Time Engagement
 * Created by Wei Hu in 2022-10.
 * Copyright (c) 2024 Agora IO. All rights reserved.
 *
 */
package extension

import (
	"chat_transcriber/pb"
	"fmt"
	"log/slog"
	"time"

	"ten_framework/ten"

	"google.golang.org/protobuf/proto"
)

const (
	textDataTextField         = "text"
	textDataFinalField        = "is_final"
	textDataStreamIdField     = "stream_id"
	textDataEndOfSegmentField = "end_of_segment"
)

var (
	logTag = slog.String("extension", "CHAT_TRANSCRIBER_EXTENSION")
)

type chatTranscriberExtension struct {
	ten.DefaultExtension

	cachedTextMap map[uint32]string // record the cached text data for each stream id
}

func newExtension(name string) ten.Extension {
	return &chatTranscriberExtension{
		cachedTextMap: make(map[uint32]string),
	}
}

// OnData receives data from ten graph.
// current suppotend data:
//   - name: text_data
//     example:
//     {"name": "text_data", "properties": {"text": "hello", "is_final": true, "stream_id": 123, "end_of_segment": true}}
func (p *chatTranscriberExtension) OnData(
	tenEnv ten.TenEnv,
	data ten.Data,
) {
	// Get the text data from data.
	text, err := data.GetPropertyString(textDataTextField)
	if err != nil {
		slog.Warn(fmt.Sprintf("OnData GetProperty %s error: %v", textDataTextField, err), logTag)
		return
	}

	// Get the 'is_final' flag from data which indicates whether the text is final,
	// otherwise it could be overwritten by the next text.
	final, err := data.GetPropertyBool(textDataFinalField)
	if err != nil {
		slog.Warn(fmt.Sprintf("OnData GetProperty %s error: %v", textDataFinalField, err), logTag)
		return
	}

	// Get the stream id from data.
	var streamId uint32
	if id, err := data.GetPropertyUint32(textDataStreamIdField); err == nil {
		streamId = id
	}

	// Get the 'end_of_segment' flag from data which indicates whether a line break is needed.
	endOfSegment, err := data.GetPropertyBool(textDataEndOfSegmentField)
	if err != nil {
		slog.Warn(fmt.Sprintf("OnData GetProperty %s error: %v", textDataEndOfSegmentField, err), logTag)
		return
	}

	slog.Debug(fmt.Sprintf(
		"OnData %s: %s %s: %t %s: %d %s: %t",
		textDataTextField,
		text,
		textDataFinalField,
		final,
		textDataStreamIdField,
		streamId,
		textDataEndOfSegmentField,
		endOfSegment), logTag)

	// We cache all final text data and append the non-final text data to the cached data
	// until the end of the segment.
	if endOfSegment {
		if cachedText, ok := p.cachedTextMap[streamId]; ok {
			text = cachedText + text
			delete(p.cachedTextMap, streamId)
		}
	} else {
		if final {
			if cachedText, ok := p.cachedTextMap[streamId]; ok {
				text = cachedText + text
				p.cachedTextMap[streamId] = text
			} else {
				p.cachedTextMap[streamId] = text
			}
		}
	}

	pb := pb.Text{
		Uid:      int32(streamId),
		DataType: "transcribe",
		Texttime: time.Now().UnixMilli(),
		Words: []*pb.Word{
			{
				Text:    text,
				IsFinal: endOfSegment,
			},
		},
	}

	pbData, err := proto.Marshal(&pb)
	if err != nil {
		slog.Warn(fmt.Sprintf("OnData Marshal error: %v", err), logTag)
		return
	}

	// convert the origin text data to the protobuf data and send it to the graph.
	tenData, err := ten.NewData("data")
	tenData.SetPropertyBytes("data", pbData)
	if err != nil {
		slog.Warn(fmt.Sprintf("OnData NewData error: %v", err), logTag)
		return
	}

	tenEnv.SendData(tenData)
}

func init() {
	slog.Info("chat_transcriber extension init", logTag)

	// Register addon
	ten.RegisterAddonAsExtension(
		"chat_transcriber",
		ten.NewDefaultExtensionAddon(newExtension),
	)
}
