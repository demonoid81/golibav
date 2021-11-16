package golibav

import (
	"github.com/demonoid81/golibav/avcodec"
	"github.com/demonoid81/golibav/avformat"
	"github.com/demonoid81/golibav/avutil"
)

type _stream = avformat.Stream

type Stream struct {
	*_stream
	formatCtx *avformat.Context
}

func (s *Stream) GuessFramerate() avutil.Rational {
	return avformat.GuessFrameRate(s.formatCtx, s._stream, nil)
}

func (s *Stream) Codecpar() *CodecParameters {
	return &CodecParameters{
		_codecParameters: s._stream.Codecpar,
	}
}

func (s *Stream) SetCodecpar(params *CodecParameters) {
	if err := averror(avcodec.CopyParameters(s._stream.Codecpar, params._codecParameters)); err != nil {
		panic(err)
	}
}

func (s *Stream) FindDecoder() {

}
