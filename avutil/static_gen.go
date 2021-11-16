// Code generated by robots; DO NOT EDIT.
//go:build !av_dynamic
// +build !av_dynamic

package golibavutil

import (
	common "github.com/demonoid81/golibav/internal/common"
	"runtime"
	"unsafe"
)

/*
#include <libavutil/avutil.h>
#include <libavutil/buffer.h>
#include <libavutil/dict.h>
#include <libavutil/frame.h>
#include <libavutil/pixdesc.h>
#include <libavutil/hwcontext.h>
#include <libavutil/log.h>
#include <libavutil/opt.h>
*/
import "C"

func CopyFrameProps(p0 *Frame, p1 *Frame) int32 {
	defer runtime.KeepAlive(p0)
	defer runtime.KeepAlive(p1)
	ret := C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(p0)), (*C.struct_AVFrame)(unsafe.Pointer(p1)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func DupeString(p0 string) *common.CChar {
	var s0 *C.char
	if p0 != "" {
		s0 = C.CString(p0)
		defer C.free(unsafe.Pointer(s0))
	}
	return (*common.CChar)(unsafe.Pointer(C.av_strdup(s0)))
}
func FindOpt(p0 unsafe.Pointer, p1 string, p2 string, p3 int32, p4 int32, p5 *unsafe.Pointer) *Option {
	var s1 *C.char
	if p1 != "" {
		s1 = C.CString(p1)
		defer C.free(unsafe.Pointer(s1))
	}
	var s2 *C.char
	if p2 != "" {
		s2 = C.CString(p2)
		defer C.free(unsafe.Pointer(s2))
	}
	defer runtime.KeepAlive(p3)
	defer runtime.KeepAlive(p4)
	return (*Option)(unsafe.Pointer(C.av_opt_find2(p0, s1, s2, *(*C.int)(unsafe.Pointer(&p3)), *(*C.int)(unsafe.Pointer(&p4)), p5)))
}
func Free(p0 unsafe.Pointer) {
	C.av_free(p0)
}
func FreeDict(p0 **Dictionary) {
	defer runtime.KeepAlive(p0)
	C.av_dict_free((**C.struct_AVDictionary)(unsafe.Pointer(p0)))
}
func FreeFrame(p0 **Frame) {
	defer runtime.KeepAlive(p0)
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(p0)))
}
func GetHWFrameBuffer(p0 *BufferRef, p1 *Frame, p2 int32) int32 {
	defer runtime.KeepAlive(p0)
	defer runtime.KeepAlive(p1)
	defer runtime.KeepAlive(p2)
	ret := C.av_hwframe_get_buffer((*C.struct_AVBufferRef)(unsafe.Pointer(p0)), (*C.struct_AVFrame)(unsafe.Pointer(p1)), *(*C.int)(unsafe.Pointer(&p2)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func InitHWFramesContext(p0 *BufferRef) int32 {
	defer runtime.KeepAlive(p0)
	ret := C.av_hwframe_ctx_init((*C.struct_AVBufferRef)(unsafe.Pointer(p0)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func Malloc(p0 uint64) unsafe.Pointer {
	defer runtime.KeepAlive(p0)
	return C.av_malloc(*(*C.size_t)(unsafe.Pointer(&p0)))
}
func MultiplyRational(p0 Rational, p1 Rational) Rational {
	defer runtime.KeepAlive(p0)
	defer runtime.KeepAlive(p1)
	ret := C.av_mul_q(*(*C.struct_AVRational)(unsafe.Pointer(&p0)), *(*C.struct_AVRational)(unsafe.Pointer(&p1)))
	return *(*Rational)(unsafe.Pointer(&ret))
}
func NewFrame() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}
func NewHWDeviceContext(p0 **BufferRef, p1 HWDeviceType, p2 string, p3 *Dictionary, p4 int32) int32 {
	defer runtime.KeepAlive(p0)
	var s2 *C.char
	if p2 != "" {
		s2 = C.CString(p2)
		defer C.free(unsafe.Pointer(s2))
	}
	defer runtime.KeepAlive(p3)
	defer runtime.KeepAlive(p4)
	ret := C.av_hwdevice_ctx_create((**C.struct_AVBufferRef)(unsafe.Pointer(p0)), (uint32)(p1), s2, (*C.struct_AVDictionary)(unsafe.Pointer(p3)), *(*C.int)(unsafe.Pointer(&p4)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func NewHWFramesContext(p0 *BufferRef) *BufferRef {
	defer runtime.KeepAlive(p0)
	return (*BufferRef)(unsafe.Pointer(C.av_hwframe_ctx_alloc((*C.struct_AVBufferRef)(unsafe.Pointer(p0)))))
}
func RefBuffer(p0 *BufferRef) *BufferRef {
	defer runtime.KeepAlive(p0)
	return (*BufferRef)(unsafe.Pointer(C.av_buffer_ref((*C.struct_AVBufferRef)(unsafe.Pointer(p0)))))
}
func RefFrame(p0 *Frame, p1 *Frame) int32 {
	defer runtime.KeepAlive(p0)
	defer runtime.KeepAlive(p1)
	ret := C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(p0)), (*C.struct_AVFrame)(unsafe.Pointer(p1)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func RescaleRound(p0 int64, p1 int64, p2 int64, p3 Rounding) int64 {
	defer runtime.KeepAlive(p0)
	defer runtime.KeepAlive(p1)
	defer runtime.KeepAlive(p2)
	ret := C.av_rescale_rnd(*(*C.int64_t)(unsafe.Pointer(&p0)), *(*C.int64_t)(unsafe.Pointer(&p1)), *(*C.int64_t)(unsafe.Pointer(&p2)), (uint32)(p3))
	return *(*int64)(unsafe.Pointer(&ret))
}
func SetDict(p0 **Dictionary, p1 string, p2 string, p3 int32) int32 {
	defer runtime.KeepAlive(p0)
	var s1 *C.char
	if p1 != "" {
		s1 = C.CString(p1)
		defer C.free(unsafe.Pointer(s1))
	}
	var s2 *C.char
	if p2 != "" {
		s2 = C.CString(p2)
		defer C.free(unsafe.Pointer(s2))
	}
	defer runtime.KeepAlive(p3)
	ret := C.av_dict_set((**C.struct_AVDictionary)(unsafe.Pointer(p0)), s1, s2, *(*C.int)(unsafe.Pointer(&p3)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func SetOpt(p0 unsafe.Pointer, p1 string, p2 string, p3 int32) int32 {
	var s1 *C.char
	if p1 != "" {
		s1 = C.CString(p1)
		defer C.free(unsafe.Pointer(s1))
	}
	var s2 *C.char
	if p2 != "" {
		s2 = C.CString(p2)
		defer C.free(unsafe.Pointer(s2))
	}
	defer runtime.KeepAlive(p3)
	ret := C.av_opt_set(p0, s1, s2, *(*C.int)(unsafe.Pointer(&p3)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func SetOptBin(p0 unsafe.Pointer, p1 string, p2 *uint8, p3 int32, p4 int32) int32 {
	var s1 *C.char
	if p1 != "" {
		s1 = C.CString(p1)
		defer C.free(unsafe.Pointer(s1))
	}
	defer runtime.KeepAlive(p2)
	defer runtime.KeepAlive(p3)
	defer runtime.KeepAlive(p4)
	ret := C.av_opt_set_bin(p0, s1, (*C.uint8_t)(unsafe.Pointer(p2)), *(*C.int)(unsafe.Pointer(&p3)), *(*C.int)(unsafe.Pointer(&p4)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func SetOptDouble(p0 unsafe.Pointer, p1 string, p2 float64, p3 int32) int32 {
	var s1 *C.char
	if p1 != "" {
		s1 = C.CString(p1)
		defer C.free(unsafe.Pointer(s1))
	}
	defer runtime.KeepAlive(p2)
	defer runtime.KeepAlive(p3)
	ret := C.av_opt_set_double(p0, s1, *(*C.double)(unsafe.Pointer(&p2)), *(*C.int)(unsafe.Pointer(&p3)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func SetOptInt(p0 unsafe.Pointer, p1 string, p2 int64, p3 int32) int32 {
	var s1 *C.char
	if p1 != "" {
		s1 = C.CString(p1)
		defer C.free(unsafe.Pointer(s1))
	}
	defer runtime.KeepAlive(p2)
	defer runtime.KeepAlive(p3)
	ret := C.av_opt_set_int(p0, s1, *(*C.int64_t)(unsafe.Pointer(&p2)), *(*C.int)(unsafe.Pointer(&p3)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func SetOptPixelFormat(p0 unsafe.Pointer, p1 string, p2 PixelFormat, p3 int32) int32 {
	var s1 *C.char
	if p1 != "" {
		s1 = C.CString(p1)
		defer C.free(unsafe.Pointer(s1))
	}
	defer runtime.KeepAlive(p3)
	ret := C.av_opt_set_pixel_fmt(p0, s1, (int32)(p2), *(*C.int)(unsafe.Pointer(&p3)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func SetOptRational(p0 unsafe.Pointer, p1 string, p2 Rational, p3 int32) int32 {
	var s1 *C.char
	if p1 != "" {
		s1 = C.CString(p1)
		defer C.free(unsafe.Pointer(s1))
	}
	defer runtime.KeepAlive(p2)
	defer runtime.KeepAlive(p3)
	ret := C.av_opt_set_q(p0, s1, *(*C.struct_AVRational)(unsafe.Pointer(&p2)), *(*C.int)(unsafe.Pointer(&p3)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func TransferHWFrameData(p0 *Frame, p1 *Frame, p2 int32) int32 {
	defer runtime.KeepAlive(p0)
	defer runtime.KeepAlive(p1)
	defer runtime.KeepAlive(p2)
	ret := C.av_hwframe_transfer_data((*C.struct_AVFrame)(unsafe.Pointer(p0)), (*C.struct_AVFrame)(unsafe.Pointer(p1)), *(*C.int)(unsafe.Pointer(&p2)))
	return *(*int32)(unsafe.Pointer(&ret))
}
func UnrefBuffer(p0 **BufferRef) {
	defer runtime.KeepAlive(p0)
	C.av_buffer_unref((**C.struct_AVBufferRef)(unsafe.Pointer(p0)))
}
func UnrefFrame(p0 *Frame) {
	defer runtime.KeepAlive(p0)
	C.av_frame_unref((*C.struct_AVFrame)(unsafe.Pointer(p0)))
}
func getHWDeviceTypeName(p0 HWDeviceType) *common.CChar {
	return (*common.CChar)(unsafe.Pointer(C.av_hwdevice_get_type_name((uint32)(p0))))
}
func getMediaTypeString(p0 MediaType) *common.CChar {
	return (*common.CChar)(unsafe.Pointer(C.av_get_media_type_string((int32)(p0))))
}
func getPixelFormatName(p0 PixelFormat) *common.CChar {
	return (*common.CChar)(unsafe.Pointer(C.av_get_pix_fmt_name((int32)(p0))))
}
func getSampleFormatName(p0 SampleFormat) *common.CChar {
	return (*common.CChar)(unsafe.Pointer(C.av_get_sample_fmt_name((int32)(p0))))
}