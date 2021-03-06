//go:generate go run github.com/demonoid81/golibav/generate

package golibavformat

// #include <libavformat/avformat.h>
import "C"

// +gen convtype struct_AVDictionary github.com/demonoid81/golibav/avutil.Dictionary
// +gen convtype struct_AVPacket github.com/demonoid81/golibav/avcodec.Packet
// +gen convtype struct_AVCodec github.com/demonoid81/golibav/avcodec.Codec
// +gen convtype struct_AVCodecContext github.com/demonoid81/golibav/avcodec.Context
// +gen convtype struct_AVRational github.com/demonoid81/golibav/avutil.Rational
// +gen convtype struct_AVFrame github.com/demonoid81/golibav/avutil.Frame
// +gen convtype struct_AVClass github.com/demonoid81/golibav/avutil.Class

// +gen convtype struct_AVInputFormat InputFormat
// +gen convtype struct_AVStream Stream
// +gen convtype struct_AVFormatContext Context
// +gen convtype struct_AVIOContext IOContext

// +gen fieldtype struct_AVStream codecpar *github.com/demonoid81/golibav/avcodec.Parameters

// +gen wrapfunc avformat_alloc_context NewContext
// +gen wrapfunc avformat_free_context FreeContext
// +gen wrapfunc avformat_open_input OpenInput
// +gen wrapfunc avformat_find_stream_info FindStreamInfo
// +gen wrapfunc avformat_close_input CloseInput
// +gen wrapfunc avformat_new_stream NewStream
// +gen wrapfunc avformat_alloc_output_context2 NewOutputContext
// +gen wrapfunc avformat_write_header WriteHeader
// +gen wrapfunc avformat_seek_file SeekFile

// +gen wrapfunc avio_open OpenIO
// +gen wrapfunc avio_close CloseIO
// +gen wrapfunc avio_alloc_context NewIOContext
// +gen wrapfunc avio_context_free FreeIOContext

// +gen wrapfunc av_find_best_stream FindBestStream
// +gen wrapfunc av_guess_frame_rate GuessFrameRate
// +gen wrapfunc av_read_frame ReadFrame
// +gen wrapfunc av_interleaved_write_frame WriteInterleavedFrame
// +gen wrapfunc av_write_trailer WriteTrailer

// +gen paramtype avio_alloc_context 4 unsafe.Pointer
// +gen paramtype avio_alloc_context 5 unsafe.Pointer
// +gen paramtype avio_alloc_context 6 unsafe.Pointer

// +gen paramtype av_find_best_stream 1 github.com/demonoid81/golibav/avutil.MediaType
