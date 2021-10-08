package utils

// IsOnlyHasDigitalAndComma 判断字符串是否只有数字和逗号, 字符串为空或只有数字和逗号时返回true, 其他为false
func IsOnlyHasDigitalAndComma(str string) (isOnlyHas bool) {

	for _, v := range str {
		if (v < '0' || v > '9') && v != ',' {
			return
		}
	}
	isOnlyHas = true
	return
}
