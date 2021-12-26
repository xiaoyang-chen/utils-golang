package utils

import "time"

// TimestampToLocalTime 将10/13/16/19位时间戳转换成本地时间字符串,
// 小于10位返回"timestamp wrong",
// 10: s级时间戳,13: ms级时间戳,16: us级时间戳,19: ns级时间戳,
// 10位格式为2006-01-02 15:04:05,
// 13位格式为2006-01-02 15:04:05.000,
// 16位格式为2006-01-02 15:04:05.000000,
// 19位格式为2006-01-02 15:04:05.000000000
func TimestampToLocalTimeStr(ts int64) (strlocalTime string) {

	var s, ns = int64(0), int64(0)
	switch {
	case ts >= 1e18: // ns
		s, ns = ts/1e9, ts%1e9
		strlocalTime = time.Unix(s, ns).
			Format("2006-01-02 15:04:05.000000000")
	case ts >= 1e15: // us
		s, ns = ts/1e6, ts%1e6*1e3
		strlocalTime = time.Unix(s, ns).Format("2006-01-02 15:04:05.000000")
	case ts >= 1e12: // ms
		s, ns = ts/1e3, ts%1e3*1e6
		strlocalTime = time.Unix(s, ns).Format("2006-01-02 15:04:05.000")
	case ts >= 1e9: // s
		strlocalTime = time.Unix(ts, 0).Format("2006-01-02 15:04:05")
	default: // wrong
		strlocalTime = "timestamp wrong"
	}
	return
}
