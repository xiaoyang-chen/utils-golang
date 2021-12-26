package utils

type Set interface {
	Exist(val interface{}) (exist bool)
	Add(val interface{})
	Delete(val interface{})
}

const capMin = 1
const capDefault = 8

var empty = struct{}{}

func StrSet(cap int) Set  { return make(strSet, checkCap(cap)) }
func Ui64Set(cap int) Set { return make(ui64Set, checkCap(cap)) }
func Ui32Set(cap int) Set { return make(ui32Set, checkCap(cap)) }
func Ui16Set(cap int) Set { return make(ui16Set, checkCap(cap)) }
func Ui8Set(cap int) Set  { return make(ui8Set, checkCap(cap)) }
func I64Set(cap int) Set  { return make(i64Set, checkCap(cap)) }
func I32Set(cap int) Set  { return make(i32Set, checkCap(cap)) }
func I16Set(cap int) Set  { return make(i16Set, checkCap(cap)) }
func I8Set(cap int) Set   { return make(i8Set, checkCap(cap)) }

func checkCap(cap int) (newCap int) {

	if newCap = cap; newCap < capMin {
		newCap = capDefault
	}
	return
}

type strSet map[string]struct{}

func (s strSet) Exist(val interface{}) (exist bool) {

	var key string
	if key, exist = val.(string); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s strSet) Add(val interface{}) {

	if key, ok := val.(string); ok {
		s[key] = empty
	}
}

func (s strSet) Delete(val interface{}) {

	if key, ok := val.(string); ok {
		delete(s, key)
	}
}

type ui64Set map[uint64]struct{}

func (s ui64Set) Exist(val interface{}) (exist bool) {

	var key uint64
	if key, exist = val.(uint64); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s ui64Set) Add(val interface{}) {

	if key, ok := val.(uint64); ok {
		s[key] = empty
	}
}

func (s ui64Set) Delete(val interface{}) {

	if key, ok := val.(uint64); ok {
		delete(s, key)
	}
}

type ui32Set map[uint32]struct{}

func (s ui32Set) Exist(val interface{}) (exist bool) {

	var key uint32
	if key, exist = val.(uint32); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s ui32Set) Add(val interface{}) {

	if key, ok := val.(uint32); ok {
		s[key] = empty
	}
}

func (s ui32Set) Delete(val interface{}) {

	if key, ok := val.(uint32); ok {
		delete(s, key)
	}
}

type ui16Set map[uint16]struct{}

func (s ui16Set) Exist(val interface{}) (exist bool) {

	var key uint16
	if key, exist = val.(uint16); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s ui16Set) Add(val interface{}) {

	if key, ok := val.(uint16); ok {
		s[key] = empty
	}
}

func (s ui16Set) Delete(val interface{}) {

	if key, ok := val.(uint16); ok {
		delete(s, key)
	}
}

type ui8Set map[uint8]struct{}

func (s ui8Set) Exist(val interface{}) (exist bool) {

	var key uint8
	if key, exist = val.(uint8); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s ui8Set) Add(val interface{}) {

	if key, ok := val.(uint8); ok {
		s[key] = empty
	}
}

func (s ui8Set) Delete(val interface{}) {

	if key, ok := val.(uint8); ok {
		delete(s, key)
	}
}

type i64Set map[int64]struct{}

func (s i64Set) Exist(val interface{}) (exist bool) {

	var key int64
	if key, exist = val.(int64); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s i64Set) Add(val interface{}) {

	if key, ok := val.(int64); ok {
		s[key] = empty
	}
}

func (s i64Set) Delete(val interface{}) {

	if key, ok := val.(int64); ok {
		delete(s, key)
	}
}

type i32Set map[int32]struct{}

func (s i32Set) Exist(val interface{}) (exist bool) {

	var key int32
	if key, exist = val.(int32); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s i32Set) Add(val interface{}) {

	if key, ok := val.(int32); ok {
		s[key] = empty
	}
}

func (s i32Set) Delete(val interface{}) {

	if key, ok := val.(int32); ok {
		delete(s, key)
	}
}

type i16Set map[int16]struct{}

func (s i16Set) Exist(val interface{}) (exist bool) {

	var key int16
	if key, exist = val.(int16); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s i16Set) Add(val interface{}) {

	if key, ok := val.(int16); ok {
		s[key] = empty
	}
}

func (s i16Set) Delete(val interface{}) {

	if key, ok := val.(int16); ok {
		delete(s, key)
	}
}

type i8Set map[int8]struct{}

func (s i8Set) Exist(val interface{}) (exist bool) {

	var key int8
	if key, exist = val.(int8); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s i8Set) Add(val interface{}) {

	if key, ok := val.(int8); ok {
		s[key] = empty
	}
}

func (s i8Set) Delete(val interface{}) {

	if key, ok := val.(int8); ok {
		delete(s, key)
	}
}

type iSet map[int]struct{}

func (s iSet) Exist(val interface{}) (exist bool) {

	var key int
	if key, exist = val.(int); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s iSet) Add(val interface{}) {

	if key, ok := val.(int); ok {
		s[key] = empty
	}
}

func (s iSet) Delete(val interface{}) {

	if key, ok := val.(int); ok {
		delete(s, key)
	}
}

type uiSet map[uint]struct{}

func (s uiSet) Exist(val interface{}) (exist bool) {

	var key uint
	if key, exist = val.(uint); !exist {
		return
	}
	_, exist = s[key]
	return
}

func (s uiSet) Add(val interface{}) {

	if key, ok := val.(uint); ok {
		s[key] = empty
	}
}

func (s uiSet) Delete(val interface{}) {

	if key, ok := val.(uint); ok {
		delete(s, key)
	}
}
