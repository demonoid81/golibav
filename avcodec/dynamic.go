//go:build av_dynamic
// +build av_dynamic

package golibavcodec

import "github.com/demonoid81/golibav/avutil"

func init() {
	initFuncs = append(initFuncs, avutil.InitLogging)
}
