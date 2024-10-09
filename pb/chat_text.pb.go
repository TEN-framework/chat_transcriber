// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: chat_text.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Text struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vendor       int32          `protobuf:"varint,1,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Version      int32          `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	Seqnum       int32          `protobuf:"varint,3,opt,name=seqnum,proto3" json:"seqnum,omitempty"`
	Uid          int32          `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag         int32          `protobuf:"varint,5,opt,name=flag,proto3" json:"flag,omitempty"`
	Time         int64          `protobuf:"varint,6,opt,name=time,proto3" json:"time,omitempty"` // final time =first nofinal time
	Lang         int32          `protobuf:"varint,7,opt,name=lang,proto3" json:"lang,omitempty"`
	Starttime    int32          `protobuf:"varint,8,opt,name=starttime,proto3" json:"starttime,omitempty"`
	Offtime      int32          `protobuf:"varint,9,opt,name=offtime,proto3" json:"offtime,omitempty"`
	Words        []*Word        `protobuf:"bytes,10,rep,name=words,proto3" json:"words,omitempty"`
	EndOfSegment bool           `protobuf:"varint,11,opt,name=end_of_segment,json=endOfSegment,proto3" json:"end_of_segment,omitempty"`
	DurationMs   int32          `protobuf:"varint,12,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	DataType     string         `protobuf:"bytes,13,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"` // transcribe ,translate
	Trans        []*Translation `protobuf:"bytes,14,rep,name=trans,proto3" json:"trans,omitempty"`
	Culture      string         `protobuf:"bytes,15,opt,name=culture,proto3" json:"culture,omitempty"`
	Texttime     int64          `protobuf:"varint,16,opt,name=texttime,proto3" json:"texttime,omitempty"` // pkg timestamp
}

func (x *Text) Reset() {
	*x = Text{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_text_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_chat_text_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_chat_text_proto_rawDescGZIP(), []int{0}
}

func (x *Text) GetVendor() int32 {
	if x != nil {
		return x.Vendor
	}
	return 0
}

func (x *Text) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Text) GetSeqnum() int32 {
	if x != nil {
		return x.Seqnum
	}
	return 0
}

func (x *Text) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Text) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *Text) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Text) GetLang() int32 {
	if x != nil {
		return x.Lang
	}
	return 0
}

func (x *Text) GetStarttime() int32 {
	if x != nil {
		return x.Starttime
	}
	return 0
}

func (x *Text) GetOfftime() int32 {
	if x != nil {
		return x.Offtime
	}
	return 0
}

func (x *Text) GetWords() []*Word {
	if x != nil {
		return x.Words
	}
	return nil
}

func (x *Text) GetEndOfSegment() bool {
	if x != nil {
		return x.EndOfSegment
	}
	return false
}

func (x *Text) GetDurationMs() int32 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *Text) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *Text) GetTrans() []*Translation {
	if x != nil {
		return x.Trans
	}
	return nil
}

func (x *Text) GetCulture() string {
	if x != nil {
		return x.Culture
	}
	return ""
}

func (x *Text) GetTexttime() int64 {
	if x != nil {
		return x.Texttime
	}
	return 0
}

type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string  `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	StartMs    int32   `protobuf:"varint,2,opt,name=start_ms,json=startMs,proto3" json:"start_ms,omitempty"`
	DurationMs int32   `protobuf:"varint,3,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	IsFinal    bool    `protobuf:"varint,4,opt,name=is_final,json=isFinal,proto3" json:"is_final,omitempty"`
	Confidence float64 `protobuf:"fixed64,5,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_text_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_chat_text_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_chat_text_proto_rawDescGZIP(), []int{1}
}

func (x *Word) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Word) GetStartMs() int32 {
	if x != nil {
		return x.StartMs
	}
	return 0
}

func (x *Word) GetDurationMs() int32 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *Word) GetIsFinal() bool {
	if x != nil {
		return x.IsFinal
	}
	return false
}

func (x *Word) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

type Translation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFinal bool     `protobuf:"varint,1,opt,name=is_final,json=isFinal,proto3" json:"is_final,omitempty"`
	Lang    string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Texts   []string `protobuf:"bytes,3,rep,name=texts,proto3" json:"texts,omitempty"`
}

func (x *Translation) Reset() {
	*x = Translation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_text_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Translation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Translation) ProtoMessage() {}

func (x *Translation) ProtoReflect() protoreflect.Message {
	mi := &file_chat_text_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Translation.ProtoReflect.Descriptor instead.
func (*Translation) Descriptor() ([]byte, []int) {
	return file_chat_text_proto_rawDescGZIP(), []int{2}
}

func (x *Translation) GetIsFinal() bool {
	if x != nil {
		return x.IsFinal
	}
	return false
}

func (x *Translation) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Translation) GetTexts() []string {
	if x != nil {
		return x.Texts
	}
	return nil
}

var File_chat_text_proto protoreflect.FileDescriptor

var file_chat_text_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x61, 0x67, 0x6f, 0x72, 0x61, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x22, 0xdf, 0x03, 0x0a, 0x04, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x6e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x71, 0x6e, 0x75, 0x6d, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c,
	0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x67, 0x6f, 0x72, 0x61, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x52,
	0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66,
	0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x65, 0x6e, 0x64, 0x4f, 0x66, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x67, 0x6f, 0x72,
	0x61, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x6c, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x75, 0x6c, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x65, 0x78, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x74, 0x65, 0x78, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x04,
	0x57, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x4d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22,
	0x52, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65,
	0x78, 0x74, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chat_text_proto_rawDescOnce sync.Once
	file_chat_text_proto_rawDescData = file_chat_text_proto_rawDesc
)

func file_chat_text_proto_rawDescGZIP() []byte {
	file_chat_text_proto_rawDescOnce.Do(func() {
		file_chat_text_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_text_proto_rawDescData)
	})
	return file_chat_text_proto_rawDescData
}

var file_chat_text_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_text_proto_goTypes = []interface{}{
	(*Text)(nil),        // 0: agora.chat_transcriber.Text
	(*Word)(nil),        // 1: agora.chat_transcriber.Word
	(*Translation)(nil), // 2: agora.chat_transcriber.Translation
}
var file_chat_text_proto_depIdxs = []int32{
	1, // 0: agora.chat_transcriber.Text.words:type_name -> agora.chat_transcriber.Word
	2, // 1: agora.chat_transcriber.Text.trans:type_name -> agora.chat_transcriber.Translation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chat_text_proto_init() }
func file_chat_text_proto_init() {
	if File_chat_text_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_text_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Text); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_text_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_text_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Translation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_text_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_text_proto_goTypes,
		DependencyIndexes: file_chat_text_proto_depIdxs,
		MessageInfos:      file_chat_text_proto_msgTypes,
	}.Build()
	File_chat_text_proto = out.File
	file_chat_text_proto_rawDesc = nil
	file_chat_text_proto_goTypes = nil
	file_chat_text_proto_depIdxs = nil
}
