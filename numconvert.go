package utils

import "fmt"

// NumToFloat64 将数值转换成float64,为数值返回浮点数值和true,非数值返回0和false
func NumToFloat64(a interface{}) (f float64, ok bool) {

	switch x := a.(type) {
	case float64:
		f, ok = x, true
	case float32:
		f, ok = float64(x), true

	case int:
		f, ok = float64(x), true
	case int8:
		f, ok = float64(x), true
	case int16:
		f, ok = float64(x), true
	case int32:
		f, ok = float64(x), true
	case int64:
		f, ok = float64(x), true

	case uint:
		f, ok = float64(x), true
	case uint8:
		f, ok = float64(x), true
	case uint16:
		f, ok = float64(x), true
	case uint32:
		f, ok = float64(x), true
	case uint64:
		f, ok = float64(x), true
	}
	return
}

// Float64ToKMGT 大于1000时默认输出2位小数(小于1000时输出所有小数忽略末尾的0),
// 单位为K(kilo),M(Mega),G(Giga),T(Tera)的数值,
// 1K == 10^3, 1M == 10^6, 1G == 10^9, 1T == 10^12
func Float64ToKMGT(float float64) (kmgt string) {

	var format string
	switch {
	case float >= 1e12:
		format, float = "%.2fT", float/1e12
	case float >= 1e9:
		format, float = "%.2fG", float/1e9
	case float >= 1e6:
		format, float = "%.2fM", float/1e6
	case float >= 1e3:
		format, float = "%.2fK", float/1e3
	case float > -1e3:
		format = "%g"
	case float > -1e6:
		format, float = "%.2fK", float/1e3
	case float > -1e9:
		format, float = "%.2fM", float/1e6
	case float > -1e12:
		format, float = "%.2fG", float/1e9
	case float <= -1e12:
		format, float = "%.2fT", float/1e12
	default:
		format = "%g can't be represent by Float64ToKMGT"
	}
	kmgt = fmt.Sprintf(format, float)
	return
}

// ConvertPositiveIntegerToStringBase10, if in < 0, will return "-1"
func ConvertPositiveIntegerToStringBase10(in int64) (str string) {

	if in < 0 {
		return "-1"
	}
	var bs = make([]byte, 19) // len("9223372036854775807"), math.MaxInt64
	var i = 18
	for ; in >= 10; i-- {
		bs[i] = '0' + byte(in%10)
		in /= 10
	}
	bs[i] = '0' + byte(in)
	str = Bytes2Str(bs[i:])
	return
}
