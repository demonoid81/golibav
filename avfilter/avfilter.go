//go:generate go run github.com/demonoid81/golibav/generate

package golibavfilter

// #include <libavfilter/buffersrc.h>
// #include <libavfilter/buffersink.h>
import "C"

// +gen convtype struct_AVDictionary github.com/demonoid81/golibav/avutil.Dictionary
// +gen convtype struct_AVFrame github.com/demonoid81/golibav/avutil.Frame
// +gen convtype struct_AVBufferRef github.com/demonoid81/golibav/avutil.BufferRef
// +gen convtype struct_AVRational github.com/demonoid81/golibav/avutil.Rational

// +gen convtype struct_AVBufferSrcParameters BufferSourceParameters
// +gen convtype struct_AVFilter Filter
// +gen convtype struct_AVFilterContext Context
// +gen convtype struct_AVFilterGraph Graph
// +gen convtype struct_AVFilterInOut InOut

// +gen fieldtype struct_AVBufferSrcParameters format github.com/demonoid81/golibav/avutil.PixelFormat

// +gen wrapfunc avfilter_graph_alloc NewGraph
// +gen wrapfunc avfilter_graph_free FreeGraph
// +gen wrapfunc avfilter_graph_create_filter CreateFilterGraph
// +gen wrapfunc avfilter_link Link
// +gen wrapfunc avfilter_get_by_name GetByName
// +gen wrapfunc avfilter_inout_free FreeInOut
// +gen wrapfunc avfilter_graph_parse2 ParseGraph
// +gen wrapfunc avfilter_graph_config ConfigGraph

// +gen wrapfunc av_buffersrc_parameters_alloc NewBufferSourceParameters
// +gen wrapfunc av_buffersrc_parameters_set SetBufferSourceParameters
// +gen wrapfunc av_buffersrc_write_frame WriteBufferSourceFrame

// +gen wrapfunc av_buffersink_get_frame GetBufferSinkFrame
