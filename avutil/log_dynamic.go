//go:build av_dynamic
// +build av_dynamic

package golibavutil

func init() {
	initFuncs = append(initFuncs, InitLogging)
}

func InitLogging() {
	dynamicInit2()
	initLogging()
}
