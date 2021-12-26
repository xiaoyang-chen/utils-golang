package utils

import (
	"fmt"
	"strings"
)

var stringsReplacer = strings.NewReplacer(
	"'", "", "[", "'", "]", "'", " ", "', '",
)

// 可能的优化：使用类似sync.Pool和bytes.Buffer进行内存优化, 使用底层指针强转的方式进行string转bytes, 和bytes转string, 优化内存使用执行效率

// GetSqlParamFromStrings 将 strings: [123asd a23fc c2rda2] 转化为 sqlParam: '123asd', 'a23fc', 'c2rda2', len(strings) == 0 => sqlParam: "", strings里的各个字符串不能存在字符'''(单引号), '[', ']', ' '(空格符), 返回值 sqlParam 可以直接用于sql语句中, 不会被sql注入
func GetSqlParamFromStrings(strings []string) (sqlParam string) {

	if len(strings) == 0 {
		return
	}

	sqlParam = stringsReplacer.Replace(fmt.Sprintf("%v", strings))
	return
}

// GetSqlParamFromIntegers 将 integers: [123 23 22] 转化为 sqlParam: "123,23,22", 当integers不为整数类型的slice时sqlParam为"", 返回值 sqlParam 可以直接用于sql语句中, 不会被sql注入
func GetSqlParamFromIntegers(integers interface{}) (sqlParam string) {

	switch integers.(type) {
	case []int, []int8, []int16, []int32, []int64, []uint, []uint8, []uint16, []uint32, []uint64:
		sqlParam = fmt.Sprintf("%v", integers)
		sqlParam = strings.ReplaceAll(
			sqlParam[1:len(sqlParam)-1], " ", ",",
		)
	}
	return
}

// GetLikeSqlParamStrInTwoPercent 将 likeStr: "1d2fg3" 转化为 sqlParam: "'%1d2fg3%'", likeStr == ""或likeStr含有'''(单引号) => sqlParam: "", 返回值 sqlParam 可以直接用于sql语句中, 不会被sql注入
func GetLikeSqlParamStrInTwoPercent(likeStr string) (sqlParam string) {

	if likeStr == "" || strings.ContainsRune(likeStr, '\'') {
		return
	}

	sqlParam = fmt.Sprintf("'%%%s%%'", likeStr)
	return
}

// GetSqlParamStrInTwoSingleQuote 将 srcStr: "1d2fg3" 转化为 sqlParam: "'1d2fg3'", srcStr == "" => sqlParam: "''", srcStr含有'''(单引号) => sqlParam: "", 返回值 sqlParam 可以直接用于sql语句中, 不会被sql注入
func GetSqlParamStrInTwoSingleQuote(srcStr string) (sqlParam string) {

	if srcStr == "" {
		sqlParam = "''"
	} else if strings.ContainsRune(srcStr, '\'') {
	} else {
		sqlParam = fmt.Sprintf("'%s'", srcStr)
	}
	return
}

// GetMulQuestionMarkDividedByComma 返回用逗号隔开的length个'?', length <= 0 => sqlParam: "", length == 1 => sqlParam: "?", length > 1 => sqlParam: "?,...,?"
func GetMulQuestionMarkDividedByComma(length int) (sqlParam string) {

	if length <= 0 {
		return
	}

	length = length*2 - 1
	var bs = make([]byte, length)
	bs[0] = '?'
	for i := 1; i < length; i += 2 {
		bs[i], bs[i+1] = ',', '?'
	}
	sqlParam = Bytes2Str(bs)
	return
}
