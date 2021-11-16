//go:build av_dynamic
// +build av_dynamic

package golibavfilter

import "github.com/demonoid81/golibav/avutil"

func init() {
	initFuncs = append(initFuncs, avutil.InitLogging)
}
