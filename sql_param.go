package utils

import (
	"fmt"
	"strings"
)

var stringsReplacer = strings.NewReplacer(
	"[", "'", "]", "'", " ", "', '",
)

// 可能的优化：使用类似sync.Pool和bytes.Buffer进行内存优化, 使用底层指针强转的方式进行string转bytes, 和bytes转string, 优化内存使用执行效率

// GetSqlParamFromStrings 将 strings: [123asd a23fc c2rda2] 转化为 sqlParam: '123asd', 'a23fc', 'c2rda2', len(strings) == 0 => sqlParam: "", strings里的各个字符串不能存在字符'[', ']', ' ', 返回值 sqlParam 可以直接用于sql语句中, 不会被sql注入
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

// GetLikeSqlParamStrInTwoPercent 将 likeStr: "1d2fg3" 转化为 sqlParam: "'%1d2fg3%'", likeStr == "" => sqlParam: "", 返回值 sqlParam 可以直接用于sql语句中, 不会被sql注入
func GetLikeSqlParamStrInTwoPercent(likeStr string) (sqlParam string) {

	if likeStr == "" {
		return
	}
	sqlParam = fmt.Sprintf("'%%%s%%'", likeStr)
	return
}
