package utils

import "errors"

const cnPhoneLength = 11
const cnPhoneLengthAddOne = cnPhoneLength + 1

var ErrCNPhonesStringFormatWrong = errors.New("the format of china phones string wrong, must be mobile number or multiple mobile number divided by the comma")

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

// CheckCNPhonesFormat phones 多个手机号用','隔开
func CheckCNPhonesFormat(phones string) (err error) {

	var lenPhones = len(phones)
	if lenPhones%(cnPhoneLengthAddOne) != cnPhoneLength {
		err = ErrCNPhonesStringFormatWrong
		return
	}

	for i := 0; i < lenPhones; i++ {
		if phones[i] == ',' {
			if i%(cnPhoneLengthAddOne) != cnPhoneLength {
				err = ErrCNPhonesStringFormatWrong
				return
			} else {
				continue
			}
		}
		if phones[i] < '0' || phones[i] > '9' {
			err = ErrCNPhonesStringFormatWrong
			return
		}
	}
	return
}
