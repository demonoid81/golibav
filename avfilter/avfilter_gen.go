// Code generated by robots; DO NOT EDIT.

package golibavfilter

import (
	avutil "github.com/demonoid81/golibav/avutil"
	common "github.com/demonoid81/golibav/internal/common"
	"unsafe"
)

/*
#include <libavfilter/buffersrc.h>
#include <libavfilter/buffersink.h>
*/
import "C"

type BufferSourceParameters struct {
	Format            avutil.PixelFormat
	TimeBase          avutil.Rational
	Width             int32
	Height            int32
	SampleAspectRatio avutil.Rational
	FrameRate         avutil.Rational
	HwFramesCtx       *avutil.BufferRef
	SampleRate        int32
	ChannelLayout     uint64
}
type Context struct {
	AvClass       *C.struct_AVClass
	Filter        *Filter
	Name          *common.CChar
	InputPads     *C.struct_AVFilterPad
	Inputs        **C.struct_AVFilterLink
	NbInputs      uint32
	OutputPads    *C.struct_AVFilterPad
	Outputs       **C.struct_AVFilterLink
	NbOutputs     uint32
	Priv          unsafe.Pointer
	Graph         *Graph
	ThreadType    int32
	Internal      *C.struct_AVFilterInternal
	CommandQueue  *C.struct_AVFilterCommand
	EnableStr     *common.CChar
	Enable        unsafe.Pointer
	VarValues     *float64
	IsDisabled    int32
	HwDeviceCtx   *avutil.BufferRef
	NbThreads     int32
	Ready         uint32
	ExtraHwFrames int32
	_             [4]byte
}
type Filter struct {
	Name           *common.CChar
	Description    *common.CChar
	Inputs         *C.struct_AVFilterPad
	Outputs        *C.struct_AVFilterPad
	PrivClass      *C.struct_AVClass
	Flags          int32
	Preinit        *[0]byte
	Init           *[0]byte
	InitDict       *[0]byte
	Uninit         *[0]byte
	QueryFormats   *[0]byte
	PrivSize       int32
	FlagsInternal  int32
	Next           *Filter
	ProcessCommand *[0]byte
	InitOpaque     *[0]byte
	Activate       *[0]byte
}
type Graph struct {
	AvClass            *C.struct_AVClass
	Filters            **Context
	NbFilters          uint32
	ScaleSwsOpts       *common.CChar
	ResampleLavrOpts   *common.CChar
	ThreadType         int32
	NbThreads          int32
	Internal           *C.struct_AVFilterGraphInternal
	Opaque             unsafe.Pointer
	Execute            *C.avfilter_execute_func
	AresampleSwrOpts   *common.CChar
	SinkLinks          **C.struct_AVFilterLink
	SinkLinksCount     int32
	DisableAutoConvert uint32
}
type InOut struct {
	Name      *common.CChar
	FilterCtx *Context
	PadIdx    int32
	Next      *InOut
}
