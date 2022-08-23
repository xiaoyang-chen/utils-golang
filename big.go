package utils

import (
	"fmt"
	"math/big"
)

const errFmtBigIntSetString10 = "(*big.Int).SetString(%s, 10) fail"

var emptyBytes = []byte{} // emptyBytes as empty byte slice, can not change

func BigIntZero() *big.Int { return big.NewInt(0) }

func NewBigInt(i int64) *big.Int { return BigIntZero().SetInt64(i) }

func NewBigIntUnsigned(i uint64) *big.Int { return BigIntZero().SetUint64(i) }

// NewBigIntFromString 根据十进制对str进行解析
func NewBigIntFromString(str string) (b *big.Int, err error) {

	var ok bool
	if b, ok = BigIntZero().SetString(str, 10); !ok {
		err = fmt.Errorf(errFmtBigIntSetString10, str)
	}
	return
}

// NewBigIntMustFromString if parse str fail, will panic
func NewBigIntMustFromString(str string) (b *big.Int) {

	var err error
	if b, err = NewBigIntFromString(str); err != nil {
		panic(err)
	}
	return
}

// NewBigIntFromBytesWithNeg 将buf里的数据解析成b(*big.int), 当buf为空时, 返回0对应的*big.int, isNeg: 解析后的b是否需要转换为相反数(b = -b), 默认为false
func NewBigIntFromBytesWithNeg(buf []byte, isNeg bool) (b *big.Int) {

	var lenBuf = len(buf)
	if lenBuf == 0 {
		b = BigIntZero()
		return
	}

	b = BigIntZero().SetBytes(buf)
	if isNeg {
		b.Neg(b)
	}
	return
}

// BigInt2Bytes 将b转换成byte-slice, 当 b == nil 时返回nil, 当b的数值为0时, 返回空slice(非nil), 假设a为b转换成的byte-slice, 当 b > 0 时返回的slice为'+' + a, 当 b < 0 时返回的slice为'-' + a
func BigInt2Bytes(b *big.Int) (res []byte) {

	if b == nil {
		return
	}

	var sign = b.Sign()
	if sign == 0 {
		res = emptyBytes
		return
	}

	var bs = b.Bytes()
	res = make([]byte, len(bs)+1)
	copy(res[1:], bs)
	if sign == 1 {
		res[0] = '+'
	} else {
		res[0] = '-'
	}
	return
}

// BigIntIsZero 判断b是否为0, 当b为nil时返回true
func BigIntIsZero(b *big.Int) (isZero bool) { return b == nil || b.Sign() == 0 }

// BigIntCmp compares a and b and returns:
//
//	-1 if a <  b
//	 0 if a == b
//	+1 if a >  b
//
// if a == nil || b == nil, will panic
func BigIntCmp(a, b *big.Int) int { return a.Cmp(b) }

// BigIntEqual a和是否相等, if a == nil || b == nil, will panic
func BigIntEqual(a, b *big.Int) bool { return BigIntCmp(a, b) == 0 }

// BigIntLt returns true if a < b, if a == nil || b == nil, will panic
func BigIntLt(a, b *big.Int) bool { return BigIntCmp(a, b) == -1 }

// BigIntLte returns true if a <= b, if a == nil || b == nil, will panic
func BigIntLte(a, b *big.Int) bool { return BigIntCmp(a, b) != 1 }

// BigIntGt returns true if a > b, if a == nil || b == nil, will panic
func BigIntGt(a, b *big.Int) bool { return BigIntCmp(a, b) == 1 }

// BigIntGte returns true if a >= b, if a == nil || b == nil, will panic
func BigIntGte(a, b *big.Int) bool { return BigIntCmp(a, b) != -1 }

// BigIntNeg returns -x, if x == nil, will panic
func BigIntNeg(x *big.Int) (negX *big.Int) { return BigIntZero().Neg(x) }

// BigIntAbs returns |x|, if x == nil, will panic
func BigIntAbs(x *big.Int) (absX *big.Int) {

	switch x.Sign() {
	case 1:
		absX = BigIntCopy(x)
	case 0:
		absX = BigIntZero()
	default:
		absX = BigIntNeg(x)
	}
	return
}

// BigIntCopy returns copy of src, if src == nil, will panic
func BigIntCopy(src *big.Int) (newBi *big.Int) { return BigIntZero().Set(src) }

// BigIntMax returns max(a, b), if a == b, return a, if a == nil || b == nil, will panic
func BigIntMax(a, b *big.Int) (c *big.Int) {

	if BigIntGte(a, b) {
		c = a
	} else {
		c = b
	}
	return
}

// BigIntMin returns min(a, b), if a == b, return a, if a == nil || b == nil, will panic
func BigIntMin(a, b *big.Int) (c *big.Int) {

	if BigIntLte(a, b) {
		c = a
	} else {
		c = b
	}
	return
}
