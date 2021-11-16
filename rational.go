package golibav

import (
	"github.com/demonoid81/golibav/avutil"
)

func Rat(num, den int32) avutil.Rational {
	return avutil.Rational{
		Num: num,
		Den: den,
	}
}
