//go:build !av_dynamic
// +build !av_dynamic

package golibavutil

func init() {
	initLogging()
}

func InitLogging() {
	// noop
}
