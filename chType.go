package util

import (
	"fmt"
	"strconv"
	"strings"
)

const CH_TYPE = chType(0)

type chType int

/*
*  replaceString,指定替换第几个字符串
*
 */
func (t chType)ReplaceOneString(s string, old string, news string, pos int) string {

	if !strings.Contains(s, old) {
		return ""
	}
	if pos > strings.Count(s, old) || pos <= 0 {
		return ""
	}
	list := strings.SplitAfter(s, old)
	list[pos-1] = strings.Replace(list[pos-1], old, news, 1)
	return strings.Join(list, "")
}

//根据空格转为list，没有空格就没有
func  (t chType)StringToList(s string) []string {
	return strings.Fields(s)
}

//字符串是否在list里面
func  (t chType)IsStringInList(list []string, sub string) bool {
	tmplist := strings.Join(list, " ")
	if strings.Contains(tmplist, sub) {
		return true
	}
	return false
}

func  (t chType)IsInt64InList(s []int64, i int64) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false

}

func  (t chType)StringListToInt64(l []string)[]int64{
	lints:=[]int64{}
	for _,v:=range l{
		i,_:=strconv.ParseInt(v,10,64)
		lints=append(lints,i)
	}
	return lints
}

func  (t chType)Int64ListToStringList(l []int64)[]string{
	lints:=[]string{}
	for _,v:=range l{
		lints=append(lints,fmt.Sprintf("%v",v))
	}
	return lints
}